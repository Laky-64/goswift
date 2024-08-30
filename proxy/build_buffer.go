package proxy

import (
	"bytes"
	"fmt"
	"io"
)

func (ctx *Context) buildBuffer(addr uint64) ([]byte, error) {
	err := ctx.cr.SeekToAddr(addr)
	if err != nil {
		return nil, err
	}
	seqData := make([]uint8, 1)
	if _, err = ctx.cr.Read(seqData); err != nil {
		return nil, fmt.Errorf("failed to read to swift symbolic mangled name control data: %v", err)
	}
	var mangledData bytes.Buffer
	for {
		data := seqData[0]
		if data == 0x00 {
			break
		} else if data >= 0x01 && data <= 0x17 {
			mangledData.WriteByte(data)
			addrData := make([]byte, 4)
			_, err = ctx.cr.Read(addrData)
			if err != nil {
				return nil, err
			}
			mangledData.Write(addrData)
		} else if data >= 0x18 && data <= 0x1f {
			return nil, fmt.Errorf("unknown control character: 0x%02x", seqData[0])
		} else {
			mangledData.WriteByte(data)
		}
		if _, err = ctx.cr.Read(seqData); err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}
	}
	return mangledData.Bytes(), nil
}
