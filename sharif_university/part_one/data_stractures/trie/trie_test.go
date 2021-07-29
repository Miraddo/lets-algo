package trie

import (
	"testing"
)

func TestTrie(t *testing.T) {

	test := InitTrie()
	test.InsertTrie("Word")
	test.InsertTrie("New")

	if test.SearchTrie("wor") != true {
		t.Errorf("Expected be true but is false")
	}

	if test.SearchTrie("ew") != false {
		t.Errorf("Expected be false but is true")
	}

	if test.SearchTrie("word") != true {
		t.Errorf("Expected be true but is false")
	}

	if test.SearchTrie("new") != true {
		t.Errorf("Expected be true but is false")
	}

	if test.SearchTrie("wob") != false {
		t.Errorf("Expected be false but is true")
	}

}
