package checkout

import "fmt"

type Payment interface {
	Pay(price []int) error
}

type BankPayment struct{}

func (b BankPayment) Pay(price []int) error {
	total := 0
	for _, p := range price {
		if p <= 0 {
			return fmt.Errorf("harga tidak valid")
		}
		total += p
	}

	fmt.Println("Pembayaran BANK:", total)
	return nil
}

type OnlinePayment struct{}

func (o OnlinePayment) Pay(price []int) error {
	total := 0
	for _, o := range price {
		if o <= 0 {
			return fmt.Errorf("harga tidak valid")
		}
		total += o
	}

	fmt.Println("Pembayaran Online:", total)
	return nil
}

type FiktifPayment struct {
	Data []int
}

func (f *FiktifPayment) Pay(price []int) error {
	total := 0
	for _, p := range price {
		if p <= 0 {
			return fmt.Errorf("harga tidak valid")
		}
		total += p
	}

	f.Data = append(f.Data, total)
	return nil
}
