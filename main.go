package main

import "fmt"

// main
func main() {
	//  ==有问题结束
	var source []string
	s1 := "1"
	s2 := "2"

	source = append(source, s1)
	source = append(source, s2)
	var dest []*string

	for _, s3 := range source {
		dest = append(dest, &s3)
	}

	for _, s := range dest {
		fmt.Println(*s)
	}
	//  ==有问题结束

	// ==没有问题开始
	//var source []*string
	//s1 := "1"
	//s2 := "2"
	//
	//source = append(source, &s1)
	//source = append(source, &s2)
	//var dest []*string
	//
	//for _, s3 := range source {
	//	dest = append(dest, s3)
	//}
	//
	//for _, s := range dest {
	//	fmt.Println(*s)
	//}
	// ==没有问题结束
}
