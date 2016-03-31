package tags

import (
	"strings"
)

type SectionNode struct {
	BaseNode
	name string
}

func CreateSectionNode(name string) *SectionNode {
	name = strings.Replace(name, "\"", "", -1)
	name = strings.TrimSpace(name)
	name = strings.Title(strings.ToLower(name))
	return &SectionNode{
		name: name,
	}
}

func (s *SectionNode) Execute(out OutputWriter) {
	out.Write("<section>", true)
	out.Indent()
	out.Write("<title>"+s.name+"</title>", true)
	s.BaseNode.Execute(out)
	out.Unindent()
	out.Write("</section>", true)
}
