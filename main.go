package main

import "github.com/Albaihaqi354/koda5-b5-go/minitask10"

// "os"
//
// "bufio"
// "os"

// "github.com/Albaihaqi354/koda5-b5-go/minitask9"

func main() {
	// scanner := bufio.NewScanner(os.Stdin)

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

	// minitask9.Worker()

	// Minitask 6
	// readfile.ReadFile()

	// //Minitask 7
	// person := createperson.NewPerson(
	// 	"Bian",
	// 	"Bandung",
	// 	"087727726316",
	// )

	// fmt.Println(person.Print())
	// fmt.Println(person.Greet())

	// fmt.Println("Masukan Nama Baru: ")
	// scanner.Scan()
	// inputName := scanner.Text()
	// person.SetName(inputName)
	// fmt.Println(person.Greet())

	// // Minitask 8
	// fmt.Println("Pilih metode pembayaran:")
	// fmt.Println("1. Bank Bca")
	// fmt.Println("2. Online Qris")
	// fmt.Println("3. Fiktif")
	// fmt.Print("Pilihan: ")

	// scanner.Scan()
	// pilihan := scanner.Text()

	// var prices []int

	// switch pilihan {
	// case "1", "2":
	// 	fmt.Print("Masukkan nominal: Rp.")
	// 	scanner.Scan()
	// 	input := scanner.Text()

	// 	nominal, err := strconv.Atoi(input)
	// 	if err != nil || nominal <= 0 {
	// 		fmt.Println("Nominal tidak valid")
	// 		return
	// 	}

	// 	prices = append(prices, nominal)

	// case "3":
	// 	fmt.Println("Masukkan nominal (0 untuk selesai):")
	// 	for {
	// 		fmt.Print("Nominal: Rp.")
	// 		scanner.Scan()
	// 		input := scanner.Text()

	// 		nominal, err := strconv.Atoi(input)
	// 		if err != nil {
	// 			fmt.Println("Harus angka")
	// 			continue
	// 		}

	// 		if nominal == 0 {
	// 			break
	// 		}

	// 		prices = append(prices, nominal)
	// 	}

	// default:
	// 	fmt.Println("Pilihan tidak valid")
	// 	return
	// }

	// bank := checkout.BankPayment{}
	// online := checkout.OnlinePayment{}
	// fiktif := checkout.FiktifPayment{}

	// switch pilihan {
	// case "1":
	// 	bank.Pay(prices)

	// case "2":
	// 	online.Pay(prices)

	// case "3":
	// 	fiktif.Pay(prices)

	// 	fmt.Println("Nominal pembayaran fiktif:")
	// 	for i, v := range fiktif.Data {
	// 		fmt.Printf("%d. Rp. %d\n", i+1, v)
	// 	}
	// }

	msgChanel := make(chan minitask10.Message)

	go minitask10.Blackboard(msgChanel)

	minitask10.MsgSender(
		msgChanel, minitask10.NewMessage("Bian", "Sudah Makan?"),
	)

	minitask10.MsgSender(
		msgChanel, minitask10.NewMessage("Bunda", "Bantu Belanja"),
	)

	minitask10.MsgSender(
		msgChanel, minitask10.NewMessage("Adik", "Bantu Belajar"),
	)

	close(msgChanel)
}
