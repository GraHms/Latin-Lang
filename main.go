package main

import (
	"flag"
	"fmt"
	"github.com/grahms/samoralang/evaluator"
	"github.com/grahms/samoralang/lexer"
	"github.com/grahms/samoralang/object"
	"github.com/grahms/samoralang/parser"
	"io"
	"os"
)

//import (
//	"fmt"
//	"latin/repl"
//	"os"
//	user2 "os/user"
//)
//
//func main() {
//	user, err := user2.Current()
//	if err != nil {
//		panic(err)
//	}
//	fmt.Printf("Hi %s! This is the Latin Program", user.Username)
//	fmt.Printf("Feel free to type in commands\n")
//	repl.Start(os.Stdin, os.Stdout)
//}

func main() {

	flag.Parse()

	var input []byte
	var err error

	if len(flag.Args()) > 0 {
		input, err = os.ReadFile(os.Args[1])
	} else {
		input, err = io.ReadAll(os.Stdin)
	}

	if err != nil {
		fmt.Printf("Error reading: %s\n", err.Error())
	}

	Execute(string(input))
}

func Execute(input string) int {
	out := os.Stdout
	env := object.NewEnvironment()
	l := lexer.New(input)
	p := parser.New(l)

	program := p.ParseProgram()
	if len(p.Errors()) != 0 {
		fmt.Printf("Error parsing: %s\n", p.Errors())
		os.Exit(1)
	}
	evaluated := evaluator.Eval(program, env)
	if evaluated != nil {
		_, _ = io.WriteString(out, evaluated.Inspect())
		_, _ = io.WriteString(out, "\n")
	}
	return 0
}
