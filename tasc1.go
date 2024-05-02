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

func Sum() {
	var in *bufio.Reader
	var out *bufio.Writer
	var count, j int
	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()

	fmt.Fscan(in, &count)

	for j <= count-1 {
		var a, b int
		fmt.Fscan(in, &a, &b)
		fmt.Fprintln(out, a+b)
		j++
	}
}

/*
func main() {
	Sum()
}
*/

func main() {

	filenameData := "/home/anton/Go/OZON/1/1"
	//filenameAnswer := "/home/anton/Go/OZON/1/1a"

	fileData, err := os.Open(filenameData)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer fileData.Close()

	var count, a, b int

	fmt.Fscanln(fileData, &count)

	Sum()

	fmt.Println(count)

	for count > 0 {
		fmt.Fscanln(fileData, &a, &b)
		fmt.Println(a, b)
		count--
	}

}

/*
func try(filename string) {

	var count, a, b int

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	fmt.Fscanln(file, &count)

	for count > 0 {
		fmt.Fscanln(file, &a, &b)
		fmt.Println(a + b)
		count--
	}

}
*/

/*
func readData(filename string){

    var name string
    var age int

    file, err := os.Open(filename)
    if err != nil{
        fmt.Println(err)
        os.Exit(1)
    }
    defer file.Close()

    fmt.Fscanln(file, &name)
    fmt.Fscanln(file, &age)
    fmt.Println(name, age)

*/
