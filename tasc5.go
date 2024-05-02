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
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"strings"
)

/*
func filesCounter(string) {

}

func hascFinder(string) {

}
*/

// main - для ввода в консоль
func main() {
	var in *bufio.Reader
	var out *bufio.Writer
	var inJson *bufio.Scanner

	in = bufio.NewReader(os.Stdin)
	inJson = bufio.NewScanner(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var count int
	var lineCount int

	var ansver [][]int

	//var jsonInString []string

	fmt.Fscan(in, &count)

	//fmt.Println("Путей - ", count)

	//CountFiles := 0
	//Flag := false

	for count > 0 {
		fmt.Fscan(in, &lineCount)

		//fmt.Println("Строк в пути - ", lineCount)

		var jsonSlice []string

		for lineCount > 0 {

			//fmt.Println("путь начат")
			inJson.Scan()
			input := inJson.Text()

			jsonSlice = append(jsonSlice, input)

			lineCount--
		}

		//распаковка JSON

		jsonString := strings.Join(jsonSlice, "")

		//jsonInString = append(jsonInString, jsonString)

		fmt.Println("СТрока")
		fmt.Println(jsonString)

		var byt = []byte(jsonString)

		var dat map[string]interface{}

		err := json.Unmarshal(byt, &dat)
		if err != nil {
			fmt.Println("Не распаковался!!!!")
		}

		//json.Unmarshal(byt, &dat)

		CountFiles := 0
		Flag := false

		CountFiles, Flag = fileCounter(dat, CountFiles, Flag)

		//fmt.Println("Запускаем ответ")
		//fmt.Println("Файлов ", CountFiles)
		//fmt.Println("Флаг ", Flag)

		if !Flag {
			ansver = append(ansver, []int{CountFiles, 1})
		} else {
			ansver = append(ansver, []int{CountFiles, 0})
		}

		count--
	}

	//fmt.Println(jsonString)
	fmt.Println(ansver)
}

func fileCounter(jsonFloor map[string]interface{}, CountFiles int, Flag bool) (int, bool) {

	fmt.Println("Здесь работает fileCounter")

	if folderFiles, ok := jsonFloor["files"]; ok {
		files, _ := folderFiles.(string)

		fmt.Println(files)

		for _, values := range files {
			matched, _ := regexp.MatchString(`.hack\z`, string(values))
			if matched && !Flag {
				Flag = true
				CountFiles = len(files)
			}
		}
		CountFiles += len(files)
		fmt.Println("Здесь работает")
		fmt.Println("Файлов ", CountFiles)
		fmt.Println("Флаг ", Flag)
	}

	if _, ok := jsonFloor["folders"]; ok {

		folderIn, _ := jsonFloor["folders"].(string)

		var byt = []byte(folderIn)

		var dat map[string]interface{}

		err := json.Unmarshal(byt, &dat)
		if err != nil {
			fmt.Println("Не распаковался!!!!")
		}

		CountFiles, Flag = fileCounter(dat, CountFiles, Flag)
	}
	//fmt.Println("Файлов ", CountFiles)
	//fmt.Println("Флаг ", Flag)
	return CountFiles, Flag
}

/*
// для тестирования
func main() {
	filename := "/home/anton/Go/OZON/5/1"

	var out *bufio.Writer
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var count uint16
	var lineCount uint32
	//var time int

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	fmt.Fscan(file, &count)

	for count > 0 {
		fmt.Fscan(file, &lineCount)
		for lineCount > 0 {
			dat, _ := os.ReadFile(filename)
			fmt.Print(string(dat))
			lineCount--

		}
		//fmt.Fprintln(out, RancAssigment(times))
		count--
	}
}
*/
