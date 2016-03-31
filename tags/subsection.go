package tags

import (
	"strings"
)

type SubsectionNode struct {
	BaseNode
	name string
}

func CreateSubsectionNode(name string) *SubsectionNode {
	name = strings.Replace(name, "\"", "", -1)
	name = strings.TrimSpace(name)
	return &SubsectionNode{
		name: name,
	}
}

func (s *SubsectionNode) Execute(out OutputWriter) {
	out.Write("<b>"+s.name+"</b>", true)
	s.BaseNode.Execute(out)
}
