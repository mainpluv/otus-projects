package hw03frequencyanalysis

import (
	"sort"
	"strings"
	"unicode"
)

type WF struct {
	Word string
	Freq int
}

func Top10(s string) []string {
	if len(s) == 0 {
		return []string{}
	}
	WFmap := MapOfWords(s)             // создание мапы {слово : частота}
	WFarr := make([]WF, 0, len(WFmap)) // создание слайса структур для последующей сортировки
	for word, freq := range WFmap {
		WFarr = append(WFarr, WF{Word: word, Freq: freq})
	}
	sort.Slice(WFarr, func(i, j int) bool { // сортировка
		if WFarr[i].Freq == WFarr[j].Freq {
			return WFarr[i].Word < WFarr[j].Word
		}
		return WFarr[i].Freq > WFarr[j].Freq
	})
	var res []string
	switch { // проверка наличия необходимого количества слов
	case len(WFarr) >= 10:
		for i := 0; i < 10; i++ {
			res = append(res, WFarr[i].Word)
		}
	case len(WFarr) == 0:
		return []string{}
	default:
		for i := 0; i < len(WFarr); i++ {
			res = append(res, WFarr[i].Word)
		}
	}
	return res
}

func MapOfWords(s string) map[string]int {
	mp := make(map[string]int)
	var word string
	for i, sign := range s {
		if sign != ' ' && sign != '\n' && sign != '\t' {
			word += string(sign)
			if i != len(s)-1 {
				continue
			}
		}
		if len(word) == 1 && word[0] == '-' {
			word = ""
		}
		if len(word) == 0 {
			continue
		}
		hasLetters := CheckForLetters(word) // проверка, не является ли слово "----", чтобы удалить знаки препинания по бокам
		if hasLetters {
			for unicode.IsPunct(rune(word[len(word)-1])) {
				word = word[:len(word)-1]
			}
			for unicode.IsPunct(rune(word[0])) {
				word = word[1:]
			}
			word = strings.ToLower(word)
		}
		mp[word]++
		word = ""
	}
	return mp
}

func CheckForLetters(word string) bool {
	for _, v := range word {
		if unicode.IsLetter(v) {
			return true
		}
	}
	return false
}
