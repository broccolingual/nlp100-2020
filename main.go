package main

import (
	"fmt"
	"strings"
	"unicode"
)

type Set map[interface{}]struct{} // Set型の定義

func NewSet() *Set { // 新しい集合を作成
	return &Set{}
}

func (s Set) Add(key interface{}) { // 集合に値を追加
	s[key] = struct{}{}
}

func (s Set) Delete(key interface{}) { // 集合から値を削除
	_, ok := s[key] // 値の存在判定
	if ok {
		delete(s, key)
	}
}

func ReverseText(text string) string {
	var newText string
	for i := len(text) - 1; i >= 0; i-- {
		newText += string(text[i])
	}
	return newText
}

func CountEnglishWords(text string) []int {
	wordList := []int{}
	text = strings.ReplaceAll(text, ",", "")
	text = strings.ReplaceAll(text, ".", "")
	for _, word := range strings.Split(text, " ") {
		wordList = append(wordList, len(word))
	}
	return wordList
}

func Ngram(text string, n int) []string {
	gramList := []string{}
	for i := 0; i < len(text)-n+1; i++ {
		tmpText := ""
		for j := 0; j < n; j++ {
			tmpText += string(text[i+j])
		}
		gramList = append(gramList, tmpText)
	}
	return gramList
}

func Cipher(text string) string {
	var newText string
	for i := 0; i < len(text); i++ {
		if unicode.IsLower(rune(text[i])) {
			newText += string(219 - text[i])
		} else {
			newText += string(text[i])
		}
	}
	return newText
}

func Typoglycemia(text string) string {
	var wordList string
	return wordList
}

func main() {
	// p1
	fmt.Println("p1")
	fmt.Println(ReverseText("stressed"))

	// p3
	fmt.Println("p3")
	fmt.Println(CountEnglishWords("Now I need a drink, alcoholic of course, after the heavy lectures involving quantum mechanics."))

	// p5 (1 = unigram, 2 = bigram, 3 = trigram)
	fmt.Println("p5")
	fmt.Println(Ngram("I am an NLPer", 2))

	// p6
	fmt.Println("p6")
	set1 := NewSet()
	x := Ngram("paraparaparadise", 2)
	y := Ngram("paragraph", 2)
	fmt.Println(x)
	fmt.Println(y)
	for _, k := range x {
		set1.Add(k)
	}
	for _, k := range y {
		set1.Add(k)
	}
	fmt.Println("和集合: ", set1)
	set2 := NewSet()
	for _, k := range x {
		set2.Add(k)
	}
	for _, k := range y {
		set2.Delete(k)
	}
	fmt.Println("差集合: ", set2)
	for k, _ := range *set2 {
		set1.Delete(k)
	}
	fmt.Println("積集合: ", set1)

	// p7
	fmt.Println("p7")
	fmt.Printf("%d時の%sは%.1f\n", 12, "気温", 22.4)

	// p8
	fmt.Println("p8")
	text := "Paragraph"
	fmt.Println(text)
	fmt.Println(Cipher(text))
	fmt.Println(Cipher(Cipher(text)))

	// p9
	fmt.Println("p9")
}
