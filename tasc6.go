package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	var in *bufio.Reader
	var out *bufio.Writer
	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var countFriends, countCards, cardNumber int

	fmt.Fscan(in, &countFriends, &countCards)

	cards := make(map[int]bool, countCards)

	for i := 1; i <= countCards; i++ {
		cards[i] = true
	}

	var friendsCard [][]int

	for i := 0; i < countFriends; i++ {
		fmt.Fscan(in, &cardNumber)
		friendsCard = append(friendsCard, []int{i, cardNumber})
	}

	sort.Slice(friendsCard, func(i, j int) bool { return friendsCard[i][1] < friendsCard[j][1] })

	var cardSlice [][]int
	var resultSlice []string
	result := true

	lastCardNumber := friendsCard[0][1]

	for _, v := range friendsCard {
		//fmt.Println("друг ", v)
		flag := false
		for i := lastCardNumber + 1; i <= countCards; i++ {
			if cards[i] && i > v[1] {
				cards[i] = false
				lastCardNumber = i
				cardSlice = append(cardSlice, []int{v[0], i})
				//fmt.Println("карта ", i)
				flag = true
				break
			}
		}

		if !flag {
			result = false
			break
		}
	}
	if result {
		sort.Slice(cardSlice, func(i, j int) bool { return cardSlice[i][0] < cardSlice[j][0] })
		for i := 0; i < len(cardSlice); i++ {
			resultSlice = append(resultSlice, strconv.Itoa(cardSlice[i][1]))
		}
		fmt.Fprintln(out, strings.Join(resultSlice, " "))
	} else {
		fmt.Fprintln(out, -1)
	}
}

/*
// main для проверки
func main() {

	filename := "/home/anton/Go/OZON/6/34"

	var out *bufio.Writer
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var countFriends, countCards, cardNumber int
	//var memberCount uint16
	//var time int

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	fmt.Fscan(file, &countFriends, &countCards)

	cards := make(map[int]bool, countCards)

	for i := 1; i <= countCards; i++ {
		cards[i] = true
	}

	var friendsCard [][]int

	for i := 0; i < countFriends; i++ {
		fmt.Fscan(file, &cardNumber)
		friendsCard = append(friendsCard, []int{i, cardNumber})
	}

	sort.Slice(friendsCard, func(i, j int) bool { return friendsCard[i][1] < friendsCard[j][1] })

	//fmt.Println(friendsCard)

	var cardSlice [][]int
	var resultSlice []string
	result := true

	lastCardNumber := friendsCard[0][1]

	//fmt.Println(countCards)

	for _, v := range friendsCard {
		//fmt.Println("друг ", v)
		flag := false
		for i := lastCardNumber + 1; i <= countCards; i++ {
			if cards[i] && i > v[1] {
				cards[i] = false
				lastCardNumber = i
				cardSlice = append(cardSlice, []int{v[0], i})
				//fmt.Println("карта ", i)
				flag = true
				break
			}
		}

		if !flag {
			result = false
			break
		}

	}

	if result {
		sort.Slice(cardSlice, func(i, j int) bool { return cardSlice[i][0] < cardSlice[j][0] })
		for i := 0; i < len(cardSlice); i++ {
			resultSlice = append(resultSlice, strconv.Itoa(cardSlice[i][1]))
		}
		fmt.Println(strings.Join(resultSlice, " "))
	} else {
		fmt.Println(-1)
	}
}
*/

// main для запуска
