package main

import (
	"container/ring"
	"fmt"
	"strconv"
	"strings"
	"io/ioutil"
	"os"
)

func captcha(x string) (sum int) {
	list := strings.Split(x, "")
	r := ring.New(len(list))
	for i := 0; i < r.Len(); i++ {
		num, err := strconv.Atoi(list[i])
		if err != nil {
			panic(err)
		}
		r.Value = num
		r = r.Next()
	}
	sum = 0
	steps := (r.Len() / 2)
	for i := 0; i < r.Len(); i++ {
		stepValue := r
		for i := 0; i < steps; i++ {
			stepValue = stepValue.Next()
		}
		if r.Value.(int) == stepValue.Value.(int) {
			sum += stepValue.Value.(int)
		}
		r = r.Next()
	}
	return
}

func main() {
	// testing input from problem description
	fmt.Println(captcha("1212") == 6)
	fmt.Println(captcha("1221") == 0)
	fmt.Println(captcha("123425") == 4)
	fmt.Println(captcha("123123") == 12)
	fmt.Println(captcha("12131415") == 4)

	// actual puzzle input
	raw_input, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}
	string_input := string(raw_input)
	fmt.Println(captcha(string_input))
}
