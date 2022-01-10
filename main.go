package main

import "fmt"

// main
func main() {
	pojo := Pojo{}
	fmt.Println(pojo.id2String)
	var i interface{}
	i = 0
	temp, ok := i.(string)
	if ok {
		println(temp)
	} else {
		println("fail")
	}

}

type Pojo struct {
	id2String map[string]string
}
