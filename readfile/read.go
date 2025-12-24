package readfile

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func ReadFile() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Printf("Masukan file path:")
	scanner.Scan()
	path := scanner.Text()

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Ada Panic", r)
		}
	}()

	file, err := os.Open(path)
	if err != nil {
		fmt.Println("Error", err)
		return
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(data))
}
