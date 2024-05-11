package main

import (
	"fmt"
	"os"
)

func main() {
	months := []string{1: "Jan",2: "Feb", 3: "Mar", 4: "Apr", 5:"May", 6:"Jun", 7:"Jul", 8:"Aug", 9:"Sep", 10:"Oct", 11:"Nov", 12: "Dec"}

	fmt.Printf("len=%d, cap=%d, %v\n", len(months), cap(months), months)

	for i, v := range months {
		fmt.Fprintf(os.Stdout, "%d, %v\n", []any{i, v}...)
	}

	q2 := months[4:7]
	fmt.Println(q2)

	for i, v := range q2 {
		fmt.Fprintf(os.Stdout, "%d, %v\n", []any{i, v}...)
	}

	summer := months[6:9]
	fmt.Println(summer)

	for _, s := range summer {
		for _, q := range q2 {
			if s == q {
				fmt.Printf("%s appears in both\n", s)
			}
		}
	}
	endelessSummer := summer[:5]
	fmt.Println(endelessSummer)

	reverse(months[:])
	fmt.Println(months)

}

func reverse( s []string ) {
	for i, j := 0, len(s) - 1; i < j; i, j = i + 1, j - 1 {
		s[i], s[j] = s[j], s[i]
	}
}