package main

import "fmt"

func sort(is []int) {
	for i := 1; i < len(is); i++ {
		if is[i] > is[i-1] {
			x := is[i]
			is[i] = is[i-1]
			j := i - 1
			for {
				if j > 0 {
					if x > is[j-1] {
						is[j] = is[j-1]
						j -= 1

					} else {
						break

					}
				} else {
					break

				}

			}
			is[j] = x

		}

	}

}
func main() {
	mp := map[string]string{
		"k": "1",
		"a": "2",
		"n": "3",
		"e": "4",
	}
	mp["d"] = "5"
	mp["e"] = "6"
	fmt.Println(len(mp))
	fmt.Println(mp)
	for i, v := range mp {

		fmt.Println(i, "-------", v)
	}

	var s = "123"
	var s2 = "234"
	fmt.Println(s > s2)
	is := []int{1, 3, 2, 5, 7}
	sort(is)
	for i, v := range is {
		fmt.Println(i, "=====", v)
	}
	fmt.Println(is[:len(is)-1])
}
