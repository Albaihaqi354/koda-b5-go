package minitask11

import "fmt"

type Rekening struct {
	Balance int
}

func (r *Rekening) Deposit(dep int) {
	r.Balance += dep
	fmt.Printf("Deposit %d, Saldo Saat ini %d\n", dep, r.Balance)
}

func (r *Rekening) Transfer(dep int) {
	if r.Balance < dep {
		fmt.Printf("Transfer %d Gagal Saldo tidak cukup, Saldo Saat ini %d\n", dep, r.Balance)
		return
	}

	r.Balance -= dep
	fmt.Printf("Transfer %d berhasil, Saldo Saat ini %d\n", dep, r.Balance)
}
