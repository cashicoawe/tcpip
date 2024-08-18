package main

import "fmt"

func _pinfo(s *[]int, name string) {
	fmt.Printf("\n%s = %d\n", name, *s)
	fmt.Printf("len = %d\n", len(*s))
	fmt.Printf("cap = %d\n", cap(*s))
	fmt.Printf("%s address =\t%p\n", name, s)
	fmt.Printf("%s[0] address =\t%p\n\n", name, &((*s)[0]))
}

func arg_is_slice(s []int, name string) {
	fmt.Println("arg is slice---------------------------------")
	// 関数に slice を渡した場合、 関数内で参照できる slice は元の slice をコピーしたもの
	// 関数と引数について考えれば、当たり前
	// slice が持つ array への参照もコピーされるので、見ている array は共通

	_pinfo(&s, name)

	// append の仕様により、 array 自体が複製される
	s2 := append(s, 100)
	_pinfo(&s2, name+"_apnd")

	// 関数は浅いコピー、 append は深いコピーと考えるといいかも
}

func main() {
	s := []int{1, 2}
	_pinfo(&s, "s")

	arg_is_slice(s, "s")
}
