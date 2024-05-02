/*
 Рекомендуется использовать быстрый (буферизованный) ввод и вывод
var in *bufio.Reader
var out *bufio.Writer
in = bufio.NewReader(os.Stdin)
out = bufio.NewWriter(os.Stdout)
defer out.Flush()

var a, b int
fmt.Fscan(in, &a, &b)
fmt.Fprint(out, a + b)
*/

package main

import (
	"bufio"
	"fmt"
	"os"
)

// LableChenger - изменяет входящую строку в соответствии с вводимыми данными
func LableChenger() {
	var in *bufio.Reader
	var out *bufio.Writer
	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var startString string
	var chengeCount int

	fmt.Fscan(in, &startString)
	stringResult := toRune(startString)

	fmt.Fscan(in, &chengeCount)

	//Циклы последовательных изменений строки
	k := 0
	for k <= chengeCount-1 {
		var start, finish int
		var text string

		fmt.Fscan(in, &start, &finish, &text)

		textRune := toRune(text)

		for i := start - 1; i <= (finish - 1); i++ {
			stringResult[i] = textRune[i-start+1]
		}
		k++
	}

	//слайс рун преобразуем в строку
	result := string(stringResult)

	fmt.Println(result)

	fmt.Fprint(out, result)
}

// toRune - преобразование строки в слайс рун
func toRune(textIn string) []rune {
	var textRune []rune
	for _, v := range textIn {
		textRune = append(textRune, v)
	}
	return textRune
}

func main() {
	LableChenger()
}
