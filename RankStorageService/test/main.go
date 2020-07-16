package main

import (
	"fmt"
	"log"
)

func main() {
	num := 12
	t := fmt.Sprintf("%"+num+"v", "aaaaaaa")
	log.Println(t)
}
