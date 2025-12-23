package minitask3

import "fmt"

func ManipulasiArray() {
	num := []int{50, 75, 66, 20, 32, 90}

	num = append(
		num[:3],
		append([]int{88}, num[3:]...)...,
	)

	for i, v := range num {
		fmt.Printf("Index %d: %d\n", i, v)
	}
}
