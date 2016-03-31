package tags

import ()

type IndentNode struct {
	BaseNode
}

func CreateIndentNode() *IndentNode {
	return &IndentNode{}
}

func (s *IndentNode) Execute(out OutputWriter) {
	out.Write("<sl>", false)
	out.Indent()
	out.Write("<sli>", true)
	out.Indent()
	s.BaseNode.Execute(out)
	out.Unindent()
	out.Write("</sli>", true)
	out.Unindent()
	out.Write("</sl>", false)
}
