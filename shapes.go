package main

import (
	"fmt"
	"math"
)

// Shape 介面定義了一個方法：Area()
// 只要任何型別擁有這個方法，它就自動成為一個 Shape
type Shape interface {
	Area() float64
}

// Rectangle 結構體
type Rectangle struct {
	Width  float64
	Height float64
}

// Rectangle 實作 Area 方法
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Circle 結構體
type Circle struct {
	Radius float64
}

// Circle 實作 Area 方法 (Value Receiver)
// c 是副本，修改 c 不會影響原物件
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Scale 方法 (Pointer Receiver)
// 使用 *Circle 直接操作記憶體位址，會修改原物件的值
func (c *Circle) Scale(factor float64) {
	c.Radius = c.Radius * factor
}

// ColoredCircle (Composition 組合)
// Go 不使用繼承，而是將 Circle 嵌入 (Embed) 到新的結構體中
type ColoredCircle struct {
	Circle // 組合：ColoredCircle 自動擁有 Circle 的所有欄位與方法
	Color  string
}

// 這裡的參數型別是 Shape 介面
func printArea(s Shape) {
	fmt.Printf("圖形面積為: %.2f\n", s.Area())
}

func main() {
	rect := Rectangle{Width: 10, Height: 5}
	circ := Circle{Radius: 7}

	fmt.Println("--- 1. 隱式介面實作 ---")
	printArea(rect)
	printArea(circ)

	fmt.Println("\n--- 2. Pointer Receiver (縮放圓形) ---")
	fmt.Printf("縮放前半徑: %.2f\n", circ.Radius)
	circ.Scale(2) // 直接修改原物件
	fmt.Printf("縮放後半徑: %.2f\n", circ.Radius)
	printArea(circ)

	fmt.Println("\n--- 3. Composition (組合) ---")
	redCircle := ColoredCircle{
		Circle: Circle{Radius: 5},
		Color:  "Red",
	}
	// ColoredCircle 直接繼承了 Area() 方法
	fmt.Printf("顏色: %s, ", redCircle.Color)
	printArea(redCircle)
}
