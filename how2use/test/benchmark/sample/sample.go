package sample

func DoSomething() []string {
	var result []string
	for i := 0; i < 1000000; i++ {
		result = append(result, "Hello")
	}
	return result
}
