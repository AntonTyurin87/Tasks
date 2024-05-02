package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// main для проверки
func main() {

	filename := "/home/anton/Go/OZON/7/1"

	var out *bufio.Writer
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var countAppointment, countPlaces, countVisitors, placeNumber int
	//var memberCount uint16
	//var time int

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	flag := true

	fmt.Fscan(file, &countAppointment)

	for countAppointment > 0 {

		var placeNumberSlice, result []int
		var placeAndIndexNumberSlice [][]int

		fmt.Fscan(file, &countPlaces, &countVisitors)
		
		/*
		if countPlaces < countVisitors {
			result = append[result, "x"]
			break
		}
		*/

		for countVisitors > 0 {

			flag := false

			fmt.Fscan(file, &placeNumber)
			placeNumberSlice = append(placeNumberLine, placeNumber)
			countVisitors--
		}

		for i, v := range placeNumberSlice {
			placeAndIndexNumberSlice = append(placeAndIndexNumberSlice, []int{i, v})
		}

		sort.Slice(placeAndIndexNumberSlice, func(i, j int) bool { return placeNumberSlice[i][1] < placeNumberSlice[j][1] })

		for i:= 1, i < len(placeAndIndexNumberSlice); i++ {
			if placeAndIndexNumberSlice[i-1] == placeAndIndexNumberSlice[i] {
				flag := true
				break
			}

		if flag {
			for i:= 0; i < countPlaces; i++ {
				result = append[result, "0"]
			}

		//функция разрешения конфликта
		


		}
		countVisitors--
		}
		countAppointment--
	}


/*
	fmt.Fscan(file, &countAppointment)

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
