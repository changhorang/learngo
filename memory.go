package exam

import "fmt"

func mem_add() {
	a := 2
	b := &a // b의 값은 유지 (&: memory address)
	*b = 20 // b memory address 값을 변경 (*: 해당 address value)
	fmt.Println(a, b)

	// array
	names := [5]string{"a", "b", "c"} // array := [len]type{value}
	names[3] = "d"
	names[4] = "eee"
	fmt.Println(names)

	// slice
	_names := []string{"a", "b", "c"} // slice := []type{value}
	_names = append(_names, "ddd", "eee")
	fmt.Println(_names)
}
