package output

import (
	"github.com/fatih/color"
)

func PrintError(value any) {
	// intValue, ok := value.(int)
	// if ok {
	// 	color.Red("Код ошибки: %d", intValue)
	// }

	switch t := value.(type) {
	case string:
		color.Red(t)
	case error:
		color.Red(t.Error())
	case int:
		color.Red("Код ошибки: %d", t)
	default:
		color.Red("Неизвестная ошибка")
	}
}

//Generic

// func sum[T int | string](a, b T) T {
// 	switch d := any(a).(type) {
// 	case string:
// 		fmt.Println(d)
// 	}
// 	return a + b
// }

// type List[T any] struct {
// 	elements []T
// }

// func (l *List[T]) addElement() {
// }
