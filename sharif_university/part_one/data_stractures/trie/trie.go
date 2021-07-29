package trie

import (
	"strings"
)

// InsertTrie Insert Function we convert the words to UpperCase to put words in the write place of the index
func (tr *Trie) InsertTrie(w string)  {
	word := strings.ToUpper(w) // make words uppercase

	current :=  tr.root // get first place of root

	for i, x := range word{
		// find index of the word
		index := int(byte(x) - 65)

		if i == len(word) - 1{ // search and put the word in right place
			if current.alpha[index] == nil{ // if is nil so we add new array
				current.alpha[index] = &Node{}
			}
			current.alpha[index].isEnd = true // at the end we added true
		}else{
			if current.alpha[index] == nil{
				current.alpha[index] = &Node{}
			}
			current = current.alpha[index]
		}
	}
}

// SearchTrie Search Function is return boolean that the word is exist or not => true or false
func (tr *Trie) SearchTrie(w string) bool  {

	word := strings.ToUpper(w) // make words uppercase
	current := tr.root // get first place of root

	// if root is already empty we will return false
	if tr.root == nil{
		return false
	}

	// search part
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

