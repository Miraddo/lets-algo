package trie

// Node Struct is just create 26 nil index of *Node array
type Node struct {
	alpha [26]*Node
	isEnd bool
}


// Trie Struct to know where is our root
type Trie struct {
	root *Node
}


// InitTrie Init create new trie root
func InitTrie() *Trie {
	result := &Trie{
		root:&Node{},
	}

	return result
}
