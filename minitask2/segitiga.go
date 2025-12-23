package minitask2

import "fmt"

func SegitigaSikuSiku(num int) {
	for i := 0; i < num; i++ {
		fmt.Println("")
		for k := 0; k < i; k++ {
			fmt.Printf("*")
		}
	}

	i := 0
	for i < num {
		fmt.Println("")
		k := 0
		for k < i {
			fmt.Printf("*")
			k++
		}
		i++
	}

	for i := range num {
		fmt.Println("")
		for k := range i {
			if k < i {
				fmt.Printf("*")
			}
		}
	}
}
