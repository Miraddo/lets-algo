package trie

import (
	"testing"
)

// TestTrie is our test function to be sure our code is word correctly
func TestTrie(t *testing.T) {

	test := InitTrie() // create new Init Trie

	test.InsertTrie("Word")
	test.InsertTrie("New")

	test.InsertTrie("Bat")
	test.InsertTrie("Battle")


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

	if test.SearchTrie("bat") != true {
		t.Errorf("Expected be true but is false")
	}

	if test.SearchTrie("battle") != true {
		t.Errorf("Expected be true but is false")
	}


	if test.SearchTrie("battlee") != false {
		t.Errorf("Expected be false but is true")
	}

	if test.SearchTrie("batta") != false {
		t.Errorf("Expected be false but is true")
	}
}
