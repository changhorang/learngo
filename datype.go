package exam

import "fmt"

func ConVar() {
	const name string = "kim"
	fmt.Println(name)

	var name2 string = "lee"
	fmt.Println(name2)

	name3 := "kang" // GO 에서 data type 설정 (variable)
	name3 = "kkang"
	fmt.Println(name3)
}
