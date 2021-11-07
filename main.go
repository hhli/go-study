package main

import (
	"log"

	"github.com/hhli/go_study/compile"
)

// main
func main() {
	compile.Walk()

	files, err := compile.WalkDir("./compile/example", ".go")

	if err != nil {
		log.Printf("遍历目录出现错误:%v", err)
		return
	}

	log.Printf("file length:%d", len(files))

	//compile.DoFind(files)
}
