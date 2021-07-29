package trie

// Node Struct
type Node struct {
	alpha [26]*Node
	isEnd bool
}


// Trie Struct
type Trie struct {
	root *Node
}


// InitTrie Init
func InitTrie() *Trie {
	result := &Trie{
		root:&Node{},
	}

	return result
}
