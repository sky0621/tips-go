package sample

func DoSomething() []string {
	var result []string
	for i := 0; i < 100000; i++ {
		result = append(result, "Hello")
	}
	return result
}
