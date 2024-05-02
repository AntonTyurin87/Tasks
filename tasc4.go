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
	"sort"
	"strconv"
	"strings"
)

// RancAssigment - функция назначения места
func RancAssigment(times []int) string {

	var result string

	var timeSlice [][]int

	long := len(times)

	lastResult := -1
	lastPlace := -1

	place := make(map[int]int, long)

	for i, v := range times {
		timeSlice = append(timeSlice, []int{i, v})
	}

	sort.Slice(timeSlice, func(i, j int) bool { return timeSlice[i][1] < timeSlice[j][1] })

	for i, v := range timeSlice {
		if v[1] == lastResult || v[1] == lastResult+1 || v[1] == lastResult-1 {
			place[v[0]] = lastPlace
			lastResult = v[1]
		} else {
			place[v[0]] = i
			lastResult = v[1]
			lastPlace = i
		}
	}

	var resultSlice []string

	for i := 0; i < len(times); i++ {
		resultSlice = append(resultSlice, strconv.Itoa(place[i]+1)) //!!!
	}

	result = strings.Join(resultSlice, " ")
	return result
}

// main как должно быть
func main() {
	var in *bufio.Reader
	var out *bufio.Writer
	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var count uint16
	var memberCount uint32
	var time int

	fmt.Fscan(in, &count)

	for count > 0 {

		fmt.Fscan(in, &memberCount)

		var times []int

		for memberCount > 0 {
			fmt.Fscan(in, &time)
			times = append(times, time)
			memberCount--
		}

		fmt.Fprintln(out, RancAssigment(times))
		count--
	}
}

/*
// main для проверки
func main() {

	filename := "/home/anton/Go/OZON/4/1"

	var out *bufio.Writer
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var count uint16
	var memberCount uint32
	var time int

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	fmt.Fscan(file, &count)

	for count > 0 {
		fmt.Fscan(file, &memberCount)
		var times []int

		for memberCount > 0 {
			fmt.Fscan(file, &time)
			times = append(times, time)
			memberCount--
		}
		fmt.Fprintln(out, RancAssigment(times))

		count--
	}
}
*/
