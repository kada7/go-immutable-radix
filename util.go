package iradix

import (
	"fmt"
)


func formatNode(deep int, n *Node) string{
	s := "\n"+ multiT(deep) +"[node] "

	if n == nil {
		s += "nil"
		return s
	}

	s += fmt.Sprintf("prefix: %s, [leaf] %s\n", 
		string(n.prefix), 
		formatLeaf(n.leaf),
	)
	for _, e := range n.edges {
		s += multiT(deep)
		s += fmt.Sprintf("edge: label: %s, node: %s \n", 
			string(e.label), 
			formatNode(deep + 1, e.node))
	}
	return s
}

func multiT(count int) string{
	s := ""
	for i := 0 ; i< count; i++ {
		s+= "\t"
	}
	return s
}

func formatLeaf(ln *leafNode)string{
	if ln == nil {
		return "nil"
	}
	return fmt.Sprintf("key: %s, val: %v", ln.key, ln.val)
}
