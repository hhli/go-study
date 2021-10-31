package main

import "github.com/hhli/go_study/compile"

// main
func main() {
	//sourceCode := "package main\n import \"fmt\"\n//comment\n func main() {\n  fmt.Println(\"Hello, world!\")\n}"
	//compile.Scan([]byte(sourceCode))
	//compile.Parse([]byte(sourceCode))
	filePath := "compile/example/eg1.go"
	compile.Inspect(filePath)
}
