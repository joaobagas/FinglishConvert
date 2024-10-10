package main

import (
	"bufio"
	"fmt"
	"os"

	"finglishconvert/src"
)

func main() {
	persianText := "سلام دنیا" // "Hello, World" in Persian
	fmt.Println(persianText)

	reader := bufio.NewReader(os.Stdin)
	toTranslate, err := reader.ReadString('\n')
	if err != nil {
	}
	fmt.Printf(src.TranslateEngFar(toTranslate))
}
