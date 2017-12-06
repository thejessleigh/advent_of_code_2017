package main

// Advent of Code Day 1
// Review a sequence of digits and find the sum of all the digits
// that match the next digit in the list. The list is circuluar.
// The digit after the last digit is the first digit in the list.

// 1122 produces a sum of 3 (1 + 2) because the first digit (1)
// matches the second digit and the third digit (2) matches the fourth digit.
// 1111 produces 4 because each digit (all 1) matches the next.
// 1234 produces 0 because no digit matches the next.
// 91212129 produces 9 because the only digit that matches the next one is the last digit, 9.

import (
    "fmt"
    "strconv"
    "strings"
    "container/ring"
)

func captcha(x int) (sum int){
    list := strings.Split(strconv.Itoa(x), "")
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
    fmt.Println(captcha(1122) == 3)
    fmt.Println(captcha(1111) == 4)
    fmt.Println(captcha(1234) == 0)
    fmt.Println(captcha(91212129) == 9)
}
