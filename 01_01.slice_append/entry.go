package main

import "fmt"

func _pinfo(s *[]int, name string) {
	fmt.Printf("\n%s = %d\n", name, *s)
	fmt.Printf("len = %d\n", len(*s))
	fmt.Printf("cap = %d\n", cap(*s))
	fmt.Printf("%s address = %p\n\n", name, s)
	fmt.Printf("%s[0] address = %p\n\n", name, &((*s)[0]))
}

func main() {
	s := []int{1, 2}

	_pinfo(&s, "s")

	fmt.Println("append-----------------------------------")

	// append は、元の参照先 array とは別の array を生成する
	// &s[0] と &s2[0] は異なることが確認できる

	s2 := append(s, 10)

	_pinfo(&s, "s")
	_pinfo(&s2, "s2")

	fmt.Println("s[0] = 100-----------------------------------")

	s[0] = 100

	_pinfo(&s, "s")
	_pinfo(&s2, "s2")
}
