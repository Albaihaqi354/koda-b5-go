package minitask9

import (
	"fmt"
	"sync"
	"time"
)

func Worker() {
	var wg sync.WaitGroup
	defer fmt.Println("Berangkat Kerja")
	wg.Go(func() {
		mandi()
	})

	wg.Go(func() {
		coffee()
	})

	wg.Go(func() {
		sarapan()
	})

	wg.Go(func() {
		rapihKamar()
	})
	wg.Wait()
}

func mandi() {
	fmt.Println("Mandi")
	time.Sleep(time.Millisecond * 10000)
	fmt.Println("Mandi Selesai")
}

func coffee() {
	fmt.Println("Buat Kopi")
	time.Sleep(time.Millisecond * 5000)
	fmt.Println("Selesai Buat Kopi")
}

func sarapan() {
	fmt.Println("Sarapan")
	time.Sleep(time.Millisecond * 6000)
	fmt.Println("Sarapan Selesai")
}

func rapihKamar() {
	fmt.Println("Rapihkan kamar")
	time.Sleep(time.Microsecond * 15000)
	fmt.Println("Selesai Rapihkan Kamar")
}
