package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
)

type person struct{
	nama, nim string
	ipk int64
}

var data []person

func main() {
	fmt.Println(":::::Menu:::::")
	fmt.Println("1. Input")
	fmt.Println("2. Output")
	fmt.Println("3. Keluar")
	fmt.Print("Pilih : ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	pilih := scanner.Text()

	if pilih == "1" {
		inputData()
	} else if pilih == "2" {
		outputData()
	} else {
		os.Exit(0) 
	}
}

func inputData()  {
	fmt.Println(":::::Input Data:::::")
	fmt.Print("Banyak data : ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	banyak, _:= strconv.ParseInt(scanner.Text(), 10, 64)

	for i := int64(0); i < banyak; i++ {
		fmt.Println("\nData ", i + 1)
		fmt.Print("Nama\t : ")
		scanner.Scan()
		nama := scanner.Text()

		fmt.Print("Nim\t : ")
		scanner.Scan()
		nim := scanner.Text()

		fmt.Print("IPK\t : ")
		scanner.Scan()
		ipk, _ := strconv.ParseInt(scanner.Text(), 10, 64)

		data = append(data, person{nama, nim, ipk})
	}

	fmt.Print("Kembali? (y/n) : ")
	scanner.Scan()
	if scanner.Text() == "y" {
		main()
	} else {
		inputData()
	}
}

func outputData()  {
	fmt.Println(":::::Output Data:::::")

	for i := 0; i < len(data); i++ {
		fmt.Println("\nData ", i + 1)
		fmt.Println("Nama\t : ", data[i].nama)
		fmt.Println("NIM\t : ", data[i].nim)
		fmt.Println("IPK\t : ", data[i].ipk)
	}

	fmt.Print("Kembali? (y/n) : ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	if scanner.Text() == "y" {
		main()
	} else {
		outputData()
	}	
}
