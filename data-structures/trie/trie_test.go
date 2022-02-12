package trie

import (
	"testing"
)

func TestInit(t *testing.T) {
	// Arrange
	trie := &Trie{}

	// Act
	trie.init()

	// Assert
	if trie.root == nil {
		t.Errorf("expected not %v, got %v", nil, trie.root)
	}
}

func TestGetRuneNumber(t *testing.T) {
	// Arrange
	testsValues := []struct {
		input rune
		want  int
	}{
		{'a', 0},
		{'b', 1},
		{'c', 2},
		{'y', 24},
		{'z', 25},
		{'$', -1},
	}

	// Act
	for _, testValue := range testsValues {
		output := getRuneNumber(testValue.input)

		// Assert
		if output != testValue.want {
			t.Errorf("when the input is %v, it should return %v, but got %v", testValue.input, testValue.want, output)
		}
	}
}

func TestInsert(t *testing.T) {
	// Arrange
	trie := &Trie{}
	trie.init()

	// Act
	trie.insert("go")

	// Assert
	if trie.root.children[6] == nil {
		t.Errorf("expected not %v, got %v", nil, trie.root.children[6])
	}
	if trie.root.children[6].isEnd != false {
		t.Errorf("expected %v, got %v", false, trie.root.children[6].isEnd)
	}

	if trie.root.children[6].children[14] == nil {
		t.Errorf("expected not %v, got %v", nil, trie.root.children[6].children[14])
	}
	if trie.root.children[6].children[14].isEnd != true {
		t.Errorf("expected %v, got %v", true, trie.root.children[6].children[14].isEnd)
	}
}

func TestInsertTwoWords(t *testing.T) {
	// Arrange
	trie := &Trie{}
	trie.init()

	// Act
	trie.insert("day")
	trie.insert("date")

	// Assert
	if trie.root.children[3] == nil {
		t.Errorf("expected not %v, got %v", nil, trie.root.children[3])
	}
	if trie.root.children[3].isEnd != false {
		t.Errorf("expected %v, got %v", false, trie.root.children[3].isEnd)
	}
	if trie.root.children[3].children[0] == nil {
		t.Errorf("expected not %v, got %v", nil, trie.root.children[3].children[0])
	}
	if trie.root.children[3].children[0].isEnd != false {
		t.Errorf("expected %v, got %v", false, trie.root.children[3].children[0].isEnd)
	}
	if trie.root.children[3].children[0].children[24] == nil {
		t.Errorf("expected not %v, got %v", nil, trie.root.children[3].children[0].children[24])
	}
	if trie.root.children[3].children[0].children[24].isEnd != true {
		t.Errorf("expected %v, got %v", true, trie.root.children[3].children[0].children[24].isEnd)
	}

	if trie.root.children[3].children[0].children[19] == nil {
		t.Errorf("expected not %v, got %v", nil, trie.root.children[3].children[0].children[19])
	}
	if trie.root.children[3].children[0].children[19].isEnd != false {
		t.Errorf("expected %v, got %v", false, trie.root.children[3].children[0].children[19].isEnd)
	}
	if trie.root.children[3].children[0].children[19].children[4] == nil {
		t.Errorf("expected not %v, got %v", nil, trie.root.children[3].children[0].children[19].children[4])
	}
	if trie.root.children[3].children[0].children[19].children[4].isEnd != true {
		t.Errorf("expected %v, got %v", true, trie.root.children[3].children[0].children[19].children[4].isEnd)
	}
}

func TestFindNotFound(t *testing.T) {
	// Arrange
	trie := &Trie{}
	trie.init()

	// Act
	output := trie.find("go")

	// Assert
	if output != false {
		t.Errorf("expected %v, got %v", false, output)
	}
}

func TestFind(t *testing.T) {
	// Arrange
	trie := &Trie{}
	trie.init()
	trie.insert("go")

	// Act
	output1 := trie.find("g")
	output2 := trie.find("go")

	// Assert
	if output1 != false {
		t.Errorf("expected %v, got %v", false, output1)
	}
	if output2 != true {
		t.Errorf("expected %v, got %v", true, output2)
	}
}

func TestFindMultipleWords(t *testing.T) {
	// Arrange
	trie := &Trie{}
	trie.init()
	trie.insert("day")
	trie.insert("date")

	// Act
	output1 := trie.find("da")
	output2 := trie.find("date")
	output3 := trie.find("date")
	output4 := trie.find("dating")

	// Assert
	if output1 != false {
		t.Errorf("expected %v, got %v", false, output1)
	}
	if output2 != true {
		t.Errorf("expected %v, got %v", true, output2)
	}
	if output3 != true {
		t.Errorf("expected %v, got %v", true, output3)
	}
	if output4 != false {
		t.Errorf("expected %v, got %v", false, output4)
	}
}
