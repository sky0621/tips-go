package main

func main() {
	box1 := "3"
	print(box1)
	{
		box2 := "5"
		print(box2)
	}
	for i := 0; i < 1000; i++ {
		createBox()
	}
}

func createBox() {
	x := "test"
	print(x)
}
