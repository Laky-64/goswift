package utils

import (
	"fmt"
	"github.com/Laky-64/goswift/demangling"
	"strings"
)

func recurseNodeTree(builder *strings.Builder, node *demangling.Node, depth int) {
	builder.WriteString(fmt.Sprintf("%skind=%s", strings.Repeat(" ", depth*2), strings.TrimSuffix(node.Kind.String(), "Kind")))
	if node.Text != "" {
		builder.WriteString(fmt.Sprintf(", text=%q", node.Text))
	}
	if node.Index > 0 {
		builder.WriteString(fmt.Sprintf(", index=%d", node.Index))
	}
	builder.WriteString("\n")
	for _, child := range node.Children {
		recurseNodeTree(builder, child, depth+1)
	}
}
