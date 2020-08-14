package main

import (
	"fmt"
	"bufio"
	"os"
)

func main()  {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Masukan bilangan : ")
	num, _ := reader.ReadString('\n')

	if num % 2 == 0 {
		print("Bilangan yang diinputkan genap")
	} else {
		print("Bilangan yang diinputkan ganjil")
	}
}