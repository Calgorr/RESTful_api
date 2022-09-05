package main

import (
	"fmt"
)

type album struct {
	ID     string  `json:"id"`
	Ttile  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

func main() {
	fmt.Println("Hello wolrd!")
}
