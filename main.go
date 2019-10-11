package readable

import (
	"io/ioutil"
	"math/rand"
	"path/filepath"
	"strings"
	"time"
)

var (
	adjectives []string
	nouns      []string
	adjCount   int
	nounCount  int
	vowels     []string
)

func init() {
	fileContent, err := ioutil.ReadFile(filepath.Join(".", "words", "adjectives.txt"))
	if err != nil {
		panic("adjectives.txt file not found")
	}
	adjectives = strings.Split(string(fileContent), " ")
	adjCount = len(adjectives)

	fileContent, err = ioutil.ReadFile(filepath.Join(".", "words", "nouns.txt"))
	if err != nil {
		panic("nouns.txt file not found")
	}
	nouns = strings.Split(string(fileContent), " ")
	nounCount = len(nouns)

	vowels = []string{"a", "e", "i", "o", "u"}

	rand.Seed(time.Now().Unix())
}

func isVowel(letter string) bool {
	for _, vowel := range vowels {
		if vowel == letter {
			return true
		}
	}
	return false
}

func toTitleCase(wordList []string) []string {
	var titleCaseWordList []string
	for _, word := range wordList {
		word = strings.ToLower(word)
		word = strings.Title(word)
		titleCaseWordList = append(titleCaseWordList, word)
	}
	return titleCaseWordList
}

// GenerateSpecial ...
func GenerateSpecial(capitalize bool, wordCount int, separator string) string {
	if wordCount < 3 || wordCount > 10 {
		panic("wordCount must be between 3 and 10")
	}

	var phrase []string
	phrase = append(phrase, adjectives[rand.Intn(adjCount)])
	phrase = append(phrase, nouns[rand.Intn(nounCount)])

	if wordCount > 5 {
		for i := 0; i < wordCount-2; i++ {
			phrase = append([]string{adjectives[rand.Intn(adjCount)]}, phrase...)
		}
	} else {
		if wordCount > 2 {
			phrase = append([]string{adjectives[rand.Intn(adjCount)]}, phrase...)
		}
		if wordCount > 3 {
			if isVowel(string(phrase[0][0])) {
				article := []string{"an", "the"}[rand.Intn(2)]
				phrase = append([]string{article}, phrase...)
			} else {
				article := []string{"a", "the"}[rand.Intn(2)]
				phrase = append([]string{article}, phrase...)
			}
		}
		if wordCount > 4 {
			phrase = append(phrase, "")
			copy(phrase[3:], phrase[2:])
			phrase[2] = "and"
		}
	}

	if capitalize {
		phrase = toTitleCase(phrase)
	}

	return strings.Join(phrase, separator)
}

// Generate ...
func Generate() string {
	return GenerateSpecial(true, 3, "")
}
