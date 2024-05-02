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
	"strconv"
	"strings"
)

func NotificationSystem() {

	var in *bufio.Reader
	var out *bufio.Writer
	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var userCount, requestCount int

	fmt.Fscan(in, &userCount, &requestCount)

	massageHistory := make(map[int]int, userCount)
	for i := 1; i <= userCount; i++ {
		massageHistory[i] = 0
	}

	var result []int

	//Циклы последовательных изменений истории сообщений
	massegeCounter := 1
	for 0 < requestCount {
		var requestCode, userId int

		fmt.Fscan(in, &requestCode, &userId)

		if requestCode == 1 && userId != 0 {
			massageHistory[userId] = massegeCounter
			massegeCounter++
		} else if requestCode == 2 {
			result = append(result, massageHistory[userId])
		} else if requestCode == 1 && userId == 0 {
			for k, _ := range massageHistory {
				massageHistory[k] = massegeCounter
			}
			massegeCounter++
		}
		requestCount--
	}

	var stringSlise []string

	for _, v := range result {
		stringSlise = append(stringSlise, strconv.Itoa(v))
	}

	resultString := strings.Join(stringSlise, "\n")

	fmt.Fprint(out, resultString)
}

func main() {
	NotificationSystem()
}

/*
package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func FirsLine(adres string) (int, int, [][]int) {
	file, err := ioutil.ReadFile(adres)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	data := strings.Split(string(file), "\n")

	var requestDescription [][]int

	for i := 0; i < len(data)-1; i++ {
		step := strings.Split(data[i], " ")

		//fmt.Println(step)

		first, _ := strconv.Atoi(step[0])
		second, _ := strconv.Atoi(step[1])

		requestStep := []int{first, second}

		requestDescription = append(requestDescription, requestStep)
	}

	userCount := requestDescription[0][0]
	requestCount := requestDescription[0][1]

	return userCount, requestCount, requestDescription[1:]
}

// Вариант функции для отработки на тестовых данных

func NotificationSystem(adres string) []int {

	userCount, requestCount, requestDescription := FirsLine(adres)

	fmt.Println(requestCount)

	massageHistory := make(map[int]int, userCount)


		//for i := 1; i <= userCount; i++ {
		//	massageHistory[i] = 0
		//}


	var result []int

	massegeCounter := 1

	for _, value := range requestDescription {

		//fmt.Println(massegeCounter)

		//fmt.Println(massageHistory)

		if value[0] == 1 && value[1] != 0 {
			massageHistory[value[1]] = massegeCounter
			massegeCounter++
		} else if value[0] == 2 {
			result = append(result, massageHistory[value[1]])
		} else if value[0] == 1 && value[1] == 0 {
			for i := 1; i <= userCount; i++ {
				massageHistory[i] = massegeCounter
			}
			massegeCounter++
		}

	}

	//fmt.Println(result)
	return result
}

func Check(adres_a string) []int {
	file, err := ioutil.ReadFile(adres_a)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	data := strings.Split(string(file), "\n")

	var result []int
	for _, value := range data {
		step, _ := strconv.Atoi(value)
		result = append(result, step)
	}
	return result
}

func main() {
	adres := "/home/anton/Go/OZON/3/30"
	// "/home/anton/Go/OZON/3/0"

	fmt.Println(NotificationSystem(adres))

	fmt.Println(len(NotificationSystem(adres)))

		adres_a := "/home/anton/Go/OZON/3/14.a"
		// "/home/anton/Go/OZON/3/0.a"

		for i, v := range NotificationSystem(adres) {
			if Check(adres_a)[i] != v {
				fmt.Println("!!!!!")
				return
			}
		}
		fmt.Println("ок")

}

*/
