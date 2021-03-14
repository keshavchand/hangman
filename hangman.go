package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/eiannone/keyboard"
)

type thisString struct {
	old   []rune
	game  []rune
	ended bool
}

func generate(prob int, old string) thisString {
	t := thisString{}
	t.old = []rune(old)
	var partial []rune
	for _, i := range t.old {
		if rand.Intn(101) <= prob {
			partial = append(partial, i)
		} else {
			partial = append(partial, '_')
		}
	}
	t.game = partial
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

	return string(word)

}
func main() {
	rand.Seed(time.Now().UnixNano())

	word := getWord("words.txt", 300)
	game := generate(50, word)
	for attempt := 0; attempt < 10; attempt++ {
		fmt.Printf("%-4d %s\r", 10-attempt, string(game.game))
		i := readInputChar()
		if i == 0 {
			break
		}
		if game.iterate(i) {
			fmt.Printf("%-*s\n", len(word)+5, "Success")
			fmt.Println(string(game.game))
			break
		}
	}

	fmt.Printf("Out of attempts\nWord was %s", word)

}
