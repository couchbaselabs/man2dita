package tags

import ()

type ParagraphNode struct {
	BaseNode
}

func CreateParagraphNode() *ParagraphNode {
	return &ParagraphNode{}
}

func (s *ParagraphNode) Execute(out OutputWriter) {
	out.Write("<p>", true)
	out.Indent()
	s.BaseNode.Execute(out)
	out.Unindent()
	out.Write("</p>", false)
}
