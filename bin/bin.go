package bin

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/vpaulo/wolf/lexer"
	"github.com/vpaulo/wolf/token"
)

func Run(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	var codes bytes.Buffer
	for {

		scanned := scanner.Scan()

		if !scanned {
			break
		}

		codes.WriteString(scanner.Text() + "\n")

	}

	l := lexer.New(codes.String())

	for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
		fmt.Printf("%+v\n", tok)
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
