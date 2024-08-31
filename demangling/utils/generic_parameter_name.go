package utils

func genericParameterName(depth rune, index rune) string {
	var name string
	for {
		name += string('A' + (index % 26))
		index /= 26
		if index == 0 {
			break
		}
	}
	if depth != 0 {
		name += string(depth)
	}
	return name
}
