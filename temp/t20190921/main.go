package main

import "fmt"

func guide(t1 int) {
	guide := `
    /\
   /a \
   /----\
  /\ %d /\
 /b \  /c \


`
	fmt.Printf(guide, t1)
}

func main() {
	guide(14)

	nums := []int{1, 3, 4, 5, 7, 9, 10}
	a := 0
	b := 0
	c := 0
	d := 0
	e := 0
	f := 0
	g := 0

	for _, na := range nums {
		a = na
		for _, nb := range nums {
			if nb == na {
				continue
			}
			b = nb
			for _, nc := range nums {
				if nc == na || nc == nb {
					continue
				}
				c = nc
				for _, nd := range nums {
					if nd == na || nd == nb || nd == nc {
						continue
					}
					d = nd
					for _, ne := range nums {
						if ne == na || ne == nb || ne == nc || ne == nd {
							continue
						}
						e = ne
						for _, nf := range nums {
							if nf == na || nf == nb || nf == nc || nf == nd || nf == ne {
								continue
							}
							f = nf
							for _, ng := range nums {
								if ng == na || ng == nb || ng == nc || ng == nd || ng == ne || ng == nf {
									continue
								}
								g = ng

								if a+b == 6 &&
									b+c+d == 12 &&
									c+e == 17 &&
									d+e == 14 &&
									d+f == 7 &&
									f+g == 12 {
									fmt.Printf("a=%d, b=%d, c=%d, d=%d, e=%d, f=%d, g=%d\n", a, b, c, d, e, f, g)
								}
							}
						}
					}
				}
			}
		}
	}
}
