package bin

import (
	"bufio"
	"bytes"
	"io"
	"log"
	"os"

	"github.com/vpaulo/wolf/evaluator"
	"github.com/vpaulo/wolf/lexer"
	"github.com/vpaulo/wolf/object"
	"github.com/vpaulo/wolf/parser"
)

func Run(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	env := object.NewEnvironment()

	var codes bytes.Buffer
	for {

		scanned := scanner.Scan()

		if !scanned {
			break
		}

		codes.WriteString(scanner.Text() + "\n")

	}

	l := lexer.New(codes.String())
	p := parser.New(l)
	program := p.ParseProgram()

	if len(p.Errors()) != 0 {
		printParserErrors(out, p.Errors())
	}

	evaluated := evaluator.Eval(program, env)
	if evaluated != nil {
		io.WriteString(out, evaluated.Inspect())
		io.WriteString(out, "\n")
	}

}

func printParserErrors(out io.Writer, errors []string) {
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}

func RunFile(path string) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatalf(err.Error())
	}
	defer file.Close()

	Run(file, os.Stdout)
}