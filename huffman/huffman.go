package huffman

func Decode(encoded string, root *Node) (dec string) {
	tr := root
	for _, s := range encoded {
		if string(s) == "0" {
			tr = tr.Left
		} else {
			tr = tr.Right
		}
		if tr.Symbol != ""{
			dec += tr.Symbol
			tr = root
		}
	}
	return
}

func Encode(msg string, alphabet map[string]*Node) (res string) {
	for _, s := range msg {
		n := alphabet[string(s)]
		t := ""
		for n.Parent != nil {
			if n.Parent.Left == n {
				t = "0" + t
			} else {
				t = "1" + t
			}
			n = n.Parent
		}
		res += t
	}
	return
}
