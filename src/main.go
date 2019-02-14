package main

import (
	"fmt"

	"./palen"
)

func main() {
	ispalen := palen.IsPalendrome("A man, a plan, a canal. Panama")
	fmt.Println(ispalen)
}
