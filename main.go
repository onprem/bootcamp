package main

import (
	"fmt"

	"github.com/onprem/bootcamp/example"

	myfmt "fmt"

	"github.com/google/uuid"
)

const (
	a string = "haha"
)

func main() {
	// s := store{
	// 	name: "b",
	// 	lat:  1.22,
	// 	long: 2.333,
	// }

	snew := store{
		name: "baaa",
		lat:  1.2111,
		long: 3.333,
	}

	myfmt.Print(snew.address())

	var nn name = "Prem"

	myfmt.Print(nn.hello())

	myfmt.Printf("\n %d\n", example.Sub(4, 5))

	id, _ := uuid.NewRandom()

	arr := []string{"a", "b", "c", id.String()}

	for i, v := range arr {
		fmt.Println(i, v)
	}
}
