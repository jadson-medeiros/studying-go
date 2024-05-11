package main

type MyNumber int

type Number interface {
	~int | ~float64
}

func Sum[T Number](m map[string]T) T {
	var sum T
	for _, v := range m {
		sum += v
	}

	return sum
}

func Compare[T comparable](a T, b T) bool {
	if a == b {
		return true
	}

	return false
}

func main() {
	m := map[string]int{"Jadson": 1000, "John": 2000}
	mfloat := map[string]float64{"Jadson": 1000.20, "John": 2000.30}
	myNumber :=  map[string]MyNumber{"Jadson": 1000, "John": 2000}
	println(Sum(m))
	println(Sum(mfloat))
	println(Sum(myNumber))
	println(Compare(10, 10))
}