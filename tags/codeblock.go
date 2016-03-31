package tags

import ()

type CodeblockNode struct {
	BaseNode
}

func CreateCodeblockNode() *CodeblockNode {
	return &CodeblockNode{}
}

func (s *CodeblockNode) Execute(out OutputWriter) {
	out.Write("<codeblock>", true)
	out.Indent()
	s.BaseNode.Execute(out)
	out.Unindent()
	out.Write("</codeblock>", false)
}
