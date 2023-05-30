package main

func main() {

}

func sum(numbers []int) int {
	result := 0
	for _, n := range numbers {
		result += n
	}
	return result
}
