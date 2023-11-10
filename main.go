package main

import (
	"fmt"
	"math/rand"
)

type ID int

var ()

func main() {

	var mat [6][6]int

	for i,j := range mat {
		
		for k, _ := range j {

			mat[i][k] = rand.Intn(10)
		}
		fmt.Println()
	}


}
