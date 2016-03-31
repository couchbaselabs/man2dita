package tags

type TopicNode struct {
	BaseNode
	topic string
}

func (t *TopicNode) Execute(out OutputWriter) {
	out.Write("<topic id=\""+t.topic+"\">", true)
	out.Indent()
	t.BaseNode.Execute(out)
	out.Unindent()
	out.Write("</topic>", true)
}

func CreateTopicNode(topic string) *TopicNode {
	return &TopicNode{
		topic: topic,
	}
}
