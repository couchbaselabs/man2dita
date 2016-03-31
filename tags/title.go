package tags

type TitleNode struct {
	BaseNode
	title string
}

func (t *TitleNode) Execute(out OutputWriter) {
	out.Write("<title>"+t.title+"</title>", true)
	t.BaseNode.Execute(out)
}

func CreateTitleNode(title string) *TitleNode {
	return &TitleNode{
		title: title,
	}
}
