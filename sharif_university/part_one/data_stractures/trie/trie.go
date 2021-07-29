package trie

import (
	"strings"
)

// InsertTrie Insert Function
func (tr *Trie) InsertTrie(w string)  {
	word := strings.ToUpper(w)

	current :=  tr.root
	for i, x := range word{
		index := int(byte(x) - 65)

		if i == len(word) - 1{
			current.alpha[index] = &Node{}
			current.isEnd = true
		}else{
			current.alpha[index] = &Node{}
			current = current.alpha[index]
		}
	}
}

// SearchTrie Search Trie
func (tr *Trie) SearchTrie(w string) bool  {

	word := strings.ToUpper(w)
	current := tr.root
	if tr.root == nil{
		return false
	}
	for i, x := range word{
		index := int(byte(x) - 65)

		if current.alpha[index] != nil{
			if i == len(word) - 1{
				return true
			}else{
				current = current.alpha[index]
			}
		}else{
			return false
		}
	}

	return false
}

