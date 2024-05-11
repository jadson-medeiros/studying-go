package main

import "fmt"

func main() {
	sal := map[string]int {"Job": 1000, "James": 2000, "John": 3000}
	fmt.Println(sal["Job"])

	delete(sal, "Job")

	fmt.Println(sal["Job"])

	sal["Job"] = 6000
	fmt.Println(sal["Job"])

	s := make(map[string]int)
	s["Job"] = 300

	fmt.Println(s["Job"])

	for name, salary := range sal {
		fmt.Printf("%s's salary is %d\n", name, salary)
	}
}