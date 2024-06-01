package golangstudy

import (
	"fmt"
	"strings"
)

// defer : func return 후 동작
func lenAndUpper2(name string) (leng int, uppercs string) {
	defer fmt.Println("I'm done")
	leng = len(name)
	uppercs = strings.ToUpper(name)
	return // leng, uppercs 작성 필요 X
}

func supperAdd(numbers ...int) int {
	total := 0
	for index, number := range numbers {
		fmt.Println(index, number)
		total += number
	}
	return total
}
