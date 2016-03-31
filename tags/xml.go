package tags

type XMLNode struct {
	BaseNode
}

func (h *XMLNode) Execute(out OutputWriter) {
	out.Write("<?xml version=\"1.0\" encoding=\"UTF-8\"?>", false)
	out.Write("<!DOCTYPE topic PUBLIC \"-//OASIS//DTD DITA Topic//EN\" \"topic.dtd\">", true)
	h.BaseNode.Execute(out)
}

func CreateXMLNode() *XMLNode {
	return &XMLNode{}
}
