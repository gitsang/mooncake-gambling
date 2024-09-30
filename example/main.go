package main

import (
	"fmt"
	"strings"

	mooncake_gambling "github.com/gitsang/mooncake-gambling"
)

func main() {
	_, results := mooncake_gambling.Gamble()
	fmt.Println(strings.Join(results[0:6], " "))
	fmt.Println(strings.Join(results[6:], " "))
}
