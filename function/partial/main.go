package main

func main() {
	//p := pricer(1000, _)
	//fmt.Println(p)
}

type priceFunc func(total, discount int) int

func pricer(total, discount int) priceFunc {
	return func(total, discount int) int {
		return total - discount
	}
}

func price(total, discount int) int {
	return total - discount
}

func priceIf(total, discount interface{}) int {
	t, ok := total.(int)
	if !ok {
		return 0
	}
	d, ok := discount.(int)
	if !ok {
		return 0
	}
	return t - d
}
