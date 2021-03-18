package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/eiannone/keyboard"
)

type thisString struct {
	old   []rune
	game  []rune
	ended bool
}

func genSetRune(inp []rune) []rune {
	digits := make([]int, 26)
	for _, i := range inp {
		digits[int(i)-'a'] = 1
	}

	digitSet := make([]rune, 0)
	for _, i := range digits {
		digitSet = append(digitSet, rune(i))
	}
	return digitSet
}
func findAndReplace(word *[]rune, needle rune, placeHolder rune) {
	for idx, i := range *word {
		if i == needle {
			(*word)[idx] = placeHolder
		}
	}

}

//make sure that old string is all of small case
func generate(prob int, old string) thisString {
	t := thisString{}
	t.old = []rune(old)
	setOfDigits := genSetRune(t.old)
	t.game = make([]rune, len(t.old))
	copy(t.game, t.old)
	for idx, i := range setOfDigits {
		if i != 0 {
			if rand.Intn(101) <= prob {
				findAndReplace(&t.game, rune(idx)+'a', '_')
			}
		}
	}
	return t
}

func (t thisString) iterate(char rune) bool {
	for idx, i := range t.old {
		if i == char {
			t.game[idx] = char
			t.old[idx] = '_'
		}
	}

	for _, i := range t.game {
		if i == '_' {
			return false
		}
	}

	return true
}

func readInputChar() rune {
	char, _, err := keyboard.GetSingleKey()
	if err != nil {
		panic(err)
	}
	return char

}

func getWords(fileName string) []string {
	var words = make([]string, 0)
	_f, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	file := bufio.NewReader(_f)
	for {
		line, _, err1 := file.ReadLine()
		if err1 != nil {
			break
		}
		words = append(words, strings.ToLower(string(line)))
	}
	return words
}
func getWord(fileName string, noOfWords int) string {
	nth := rand.Intn(noOfWords)

	_f, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	file := bufio.NewReader(_f)
	for i := 0; i < nth; i++ {
		_, _, err := file.ReadLine()
		if err != nil {
			panic(err)
		}
	}

	word, _, err := file.ReadLine()
	if err != nil {
		panic(err)
	}

	return strings.ToLower(string(word))

}
func main() {
	rand.Seed(time.Now().UnixNano())

	word := getWord("words.txt", 300)
	println(word)
	game := generate(50, word)
	for attempt := 0; attempt < 10; attempt++ {
		fmt.Printf("%-4d %s\r", 10-attempt, string(game.game))
		i := readInputChar()
		if i == 0 {
			return
		}
		if game.iterate(i) {
			fmt.Printf("%-*s\n", len(word)+5, "Success")
			fmt.Printf("The word was %s\n", string(game.game))
			return
		}
	}
	fmt.Printf("Out of attempts\nWord was %s", word)

}
