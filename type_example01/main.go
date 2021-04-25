package main

import "fmt"

/*
Burada custom type color yaradir ve main func da bunu deyer verir. Daha sonra color yeni c  class nin
describe func tunu cagirib ona argument oturub get edir

'describe' is a function with a receiver of type 'color' that requires an argument of type 'string', then returns a value of type 'string'

 */

func main() {
	c := color("Red")

	fmt.Println(c.describe("is an awesome color"))
}

type color string

func (c color) describe(description string) (string) {
	return string(c) + " " + description
}

