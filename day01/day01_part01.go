package main

import (
	"container/ring"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
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
	for i := 0; i < r.Len(); i++ {
		if r.Value == r.Prev().Value {
			sum += r.Value.(int)
		}
		r = r.Next()
	}
	return
}

func main() {
	// testing input from problem description
	fmt.Println(captcha("1122") == 3)
	fmt.Println(captcha("1111") == 4)
	fmt.Println(captcha("1234") == 0)
	fmt.Println(captcha("91212129") == 9)

	// actual puzzle input
	raw_input, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}
	string_input := string(raw_input)
	fmt.Println(captcha(string_input))
}
