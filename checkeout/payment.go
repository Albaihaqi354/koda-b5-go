package checkout

import "fmt"

type Payment interface {
	Pay(price []int) error
}

type BankPayment struct{}

func (b BankPayment) Pay(price []int) error {
	if price[0] <= 0 {
		return fmt.Errorf("harga tidak valid")
	}
	fmt.Printf("Pembayaran Bank BCA: Rp.%d\n", price[0])
	return nil
}

type OnlinePayment struct{}

func (o OnlinePayment) Pay(price []int) error {
	if price[0] <= 0 {
		return fmt.Errorf("harga tidak valid")
	}
	fmt.Printf("Pembayaran Online QRIS: Rp.%d\n", price[0])
	return nil
}

type FiktifPayment struct {
	Data []int
}

func (f *FiktifPayment) Pay(price []int) error {
	for _, p := range price {
		if p <= 0 {
			return fmt.Errorf("harga tidak valid")
		}
		f.Data = append(f.Data, p)
	}
	return nil
}
