package tags

import ()

type ListNode struct {
	BaseNode
}

func CreateListNode() *ListNode {
	return &ListNode{}
}

func (s *ListNode) Execute(out OutputWriter) {
	out.Write("<ul>", true)
	out.Indent()
	out.Write("<li>", true)
	out.Indent()
	s.BaseNode.Execute(out)
	out.Unindent()
	out.Write("</li>", true)
	out.Unindent()
	out.Write("</ul>", true)
}
