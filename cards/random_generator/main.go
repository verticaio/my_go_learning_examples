package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	fmt.Println(r.Intn(56 - 1))
	d := rand.New(rand.NewSource(1000))
	e :=  rand.New(rand.NewSource(1000))
	fmt.Println(d.Intn(56 - 1))
	fmt.Println(e.Intn(56 - 1))
}
