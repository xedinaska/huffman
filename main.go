package main

import (
	"fmt"
	"github.com/xedinaska/huffman/queue"
	"github.com/xedinaska/huffman/huffman"
)

func main()  {
	q, a := queue.NewWithPrio(
		&huffman.Node{5, "f", nil, nil, nil},
		&huffman.Node{9, "e", nil, nil, nil},
		&huffman.Node{12, "c", nil, nil, nil},
		&huffman.Node{13, "b", nil, nil, nil},
		&huffman.Node{16, "d", nil, nil, nil},
		&huffman.Node{45, "a", nil, nil, nil},
	)

	fmt.Println(huffman.Decode("001011101", huffman.Root(q)))
	fmt.Println(huffman.Encode("aabe", huffman.Alphabet(a.(map[string]queue.Elem))))
}
