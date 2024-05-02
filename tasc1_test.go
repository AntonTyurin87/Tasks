package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSum(t *testing.T) {

	var in *bufio.Reader
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()

	file, err := ioutil.ReadFile("/home/anton/Go/OZON/1/1")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	answer, err := ioutil.ReadFile("/home/anton/Go/OZON/1/1a")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var countIn, countOut, a, b, c int

	fmt.Fscanln(file, &countIn)
	countOut = countIn

	Sum()

	fmt.Fprintln(countIn)

	for countIn > 0 {
		fmt.Fscanln(file, &a, &b)
		fmt.Fprintln(a, b)
		countIn--
	}

	for countOut > 0 {
		fmt.Fscan(in, &c)
		fmt.Fscanln(answer, &cA)
		countOut--
		assert.Equal(t, c, cA)
	}

}
