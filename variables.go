package main

import "fmt"

// 4. 常量宣告 (Constants)
// 使用 const 關鍵字，一旦賦值就不可更改
const PI = 3.14159

// 5. 使用 iota 自動遞增列舉值
// iota 在 const 區塊中每出現一行就會自動加 1 (從 0 開始)
const (
	Monday    = iota // 0
	Tuesday          // 1
	Wednesday        // 2
	Thursday         // 3
	Friday           // 4
)

func main() {
	// 1. 最常用：短變數宣告 (只能用在 function 內部)
	// 自動推導型別，效率最高
	learningDays := 1

	// 2. 標準宣告：指定型別
	// 當你需要先宣告但稍後才賦值時使用
	var languageName string = "Golang"

	// 3. 多重宣告
	// Antigravity Agent 可以幫你快速生成這種成組的定義
	var (
		isFun = true
		level = "Beginner"
	)

	fmt.Printf("學習 %s 的第 %d 天，難度：%s，很有趣嗎？ %v\n", languageName, learningDays, level, isFun)

	// 打印常量與 iota
	fmt.Printf("圓周率: %f\n", PI)
	fmt.Printf("工作日列舉: 週一=%d, 週二=%d, 週三=%d\n", Monday, Tuesday, Wednesday)
}
