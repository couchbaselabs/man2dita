package tags

type BodyNode struct {
	BaseNode
}

func (b *BodyNode) Execute(out OutputWriter) {
	out.Write("<body>", true)
	out.Indent()
	b.BaseNode.Execute(out)
	out.Unindent()
	out.Write("</body>", true)
}

func CreateBodyNode() *BodyNode {
	return &BodyNode{}
}
