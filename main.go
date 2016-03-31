package main

import (
	"fmt"

	"github.com/buhe/go-hack/sub"
)

/**
angel
*/
type Rabbit struct {
	Name string
	Age  int
}

func (r *Rabbit) run() {
	fmt.Println("go go go ", r.Age)
}

func main() {
	sub.ExportedFuncion()
	printPrivate()
	var m = make(map[string]int)
	m["test1"] = 29
	fmt.Println("buhe is ", m["test1"])

	var a = make([]string, 1)
	a[0] = "buhe"
	fmt.Println(a[0], " hehe")

	hare := new(Rabbit)

	hare.Name = "hare"
	hare.Age = 3
	fmt.Println("My bady is ", hare.Name)
	hare.run()
}
