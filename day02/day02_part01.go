package main

// from a series of rows, get the difference between
// that row's largest and smallest values.
// return the sum of all of the row differences

import (
    "fmt"
    // "io/ioutil"
    // "os"
    "strconv"
    "strings"
)

func checksum(spreadsheet string) (sum int) {
    rows := strings.Split(spreadsheet, "\n")
    sum = 0
    for i := 0; i < len(rows); i++ {
        row := strings.Split(rows[i], "")
        greatest := 0
        least := 10
        for i := 0; i < len(row); i++ {
            num, err := strconv.Atoi(row[i])
            if err != nil {
                panic(err)
            }
            if num > greatest {
                greatest = num
            }
            if num < least {
                least = num
            }
        }
        sum += greatest - least
    }
    return
}

func main() {
    // test data
    s := "5195\n753\n2468"
    fmt.Println(checksum(s))
    fmt.Println(checksum(s) == 18)

    // actual data
    // raw_input, err := ioutil.ReadAll(os.Stdin)
    // if err != nil {
    //     panic(err)
    // }
    // string_input := string(raw_input)
    // fmt.Println(checksum(string_input))
}
