package exam

import (
	"fmt"
	"strings"
)

func ConVar() {
	const name string = "kim"
	fmt.Println(name)

	var name2 string = "lee"
	fmt.Println(name2)

	name3 := "kang" // GO 에서 data type 설정 (variable)
	name3 = "kkang"
	fmt.Println(name3)
}

func multiply(a int, b int) int {
	return a * b
}

// return data type 지정
func lenAndUpper(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}

// many arguments
func repeatMe(words ...string) {
	fmt.Println(words)
}
