package main

import (
	// "os"
	//
	"bufio"
	"fmt"
	"os"

	checkout "github.com/Albaihaqi354/koda5-b5-go/checkeout"
	"github.com/Albaihaqi354/koda5-b5-go/createperson"
	"github.com/Albaihaqi354/koda5-b5-go/readfile"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	// for {
	// 	fmt.Println("\n===========================================")
	// 	fmt.Println("|| 1. Hitung Luas dan Keliling Lingkaran ||")
	// 	fmt.Println("|| 2. Segitiga Siku-Siku                 ||")
	// 	fmt.Println("|| 3. Array                              ||")
	// 	fmt.Println("|| 4. Biodata                            ||")
	// 	fmt.Println("|| 0. Keluar                             ||")
	// 	fmt.Println("===========================================")
	// 	fmt.Print("Pilih menu: ")

	// 	scanner.Scan()
	// 	inputMenu := scanner.Text()

	// 	pilihan, err := strconv.Atoi(inputMenu)
	// 	if err != nil {
	// 		fmt.Println("Input harus angka!")
	// 		continue
	// 	}

	// 	switch pilihan {
	// 	case 1:
	// 		radius := float64(20)

	// 		fmt.Println("Radius:", radius)
	// 		fmt.Printf("Luas: %2f", minitask1.HitungLuas(radius))
	// 		fmt.Printf("Keliling: %2f", minitask1.HitungKeliling(radius))

	// 	case 2:
	// 		minitask2.SegitigaSikuSiku(10)

	// 	case 3:
	// 		minitask3.ManipulasiArray()

	// 	case 4:
	// 		minitask4.TampilkanBiodata()

	// 	case 0:
	// 		fmt.Println("Selesai")
	// 		return

	// 	default:
	// 		fmt.Println("Pilihan tidak valid")
	// 	}
	// }

	// Minitask 6
	readfile.ReadFile()

	//Minitask 7
	person := createperson.NewPerson(
		"Bian",
		"Bandung",
		"087727726316",
	)

	fmt.Println(person.Print())
	fmt.Println(person.Greet())

	fmt.Println("Masukan Nama Baru: ")
	scanner.Scan()
	inputName := scanner.Text()
	person.SetName(inputName)
	fmt.Println(person.Greet())

	// Minitask 8
	prices := []int{100000, 20000, 30000}

	bank := checkout.BankPayment{}
	online := checkout.OnlinePayment{}
	fiktif := checkout.FiktifPayment{}

	bank.Pay(prices)
	online.Pay(prices)
	fiktif.Pay(prices)

	fmt.Println("pembayaran fiktif:", fiktif.Data)

}
