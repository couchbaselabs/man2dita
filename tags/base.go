package tags

type AbstractNode interface {
	Add(a AbstractNode)
	Execute(out OutputWriter)
}

type BaseNode struct {
	nodes []AbstractNode
}

func (b *BaseNode) Add(a AbstractNode) {
	b.nodes = append(b.nodes, a)
}

func (b *BaseNode) Execute(out OutputWriter) {
	for i := 0; i < len(b.nodes); i++ {
		b.nodes[i].Execute(out)
	}
}

type OutputWriter interface {
	Write(string, bool)
	Indent()
	Unindent()
}
