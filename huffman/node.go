package huffman

import "github.com/xedinaska/huffman/queue"

type Node struct {
	Freq   int
	Symbol string
	Left   *Node
	Right  *Node
	Parent *Node
}

func (n *Node) GetKey() string {
	return n.Symbol
}

func (n *Node) GetValue() int {
	return n.Freq
}

func Root(q queue.Queue) *Node {
	n := len(q)
	for i:=0; i<(n-1); i++{
		z := &Node{}

		left := q.Pop().(*Node)
		left.Parent = z

		right := q.Pop().(*Node)
		right.Parent = z

		z.Left = left
		z.Right = right

		z.Freq = left.Freq + right.Freq

		q.Insert(z)
	}
	return q.Pop().(*Node)
}

func Alphabet(a map[string]queue.Elem) map[string]*Node {
	alphabet := make(map[string]*Node, 0)
	for k, v := range a {
		alphabet[k] = v.(*Node)
	}
	return alphabet
}