package main

import "fmt"

func main() {
	fmt.Println("--- FizzBuzz 1 到 20 (使用 Switch) ---")

	for i := 1; i <= 20; i++ {
		// 在 Go 中，switch 後面可以不接變數，直接在 case 寫布林判斷
		switch {
		case i%3 == 0 && i%5 == 0:
			fmt.Println("FizzBuzz")
		case i%3 == 0:
			fmt.Println("Fizz")
		case i%5 == 0:
			fmt.Println("Buzz")
		default:
			fmt.Println(i)
		}
	}

	nums := []int{2, 4, 6}
	for index, value := range nums {
		fmt.Printf("索引：%d, 數值：%d\n", index, value)
	}

	switch day := "Friday"; day {
	case "Monday", "Tuesday":
		fmt.Println("工作日")
	case "Wednesday":
		fmt.Println("小週末")
	default:
		fmt.Println("其他")
	}
}
