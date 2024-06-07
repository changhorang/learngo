package exam

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

func canIDrink(age int) bool {
	if koreanAge := age - 2; koreanAge < 18 { // if-else에서만 사용하기 위한 variable expression
		return false
	} else {
		return true
	}
}

func canIDrink2(age int) bool {
	switch koreanAge := age - 2; {
	case koreanAge < 18:
		return false
	case koreanAge == 18:
		return true
	case koreanAge > 50:
		return false // if-else 남용 방지 가능
	}
	return false // 어떤 case도 아닌 경우에는 false
}
