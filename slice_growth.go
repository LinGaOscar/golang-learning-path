package main

import (
	"fmt"
)

func main() {
	// 建立一個初始長度為 0，容量為 2 的 slice
	s := make([]int, 0, 2)

	fmt.Printf("初始狀態: len=%d, cap=%d, pointer=%p\n", len(s), cap(s), s)

	// 連續新增元素，直到觸發擴容
	for i := 1; i <= 10; i++ {
		oldPtr := fmt.Sprintf("%p", s)
		s = append(s, i)
		newPtr := fmt.Sprintf("%p", s)

		fmt.Printf("新增 %2d: len=%2d, cap=%2d, pointer=%s", i, len(s), cap(s), newPtr)

		if oldPtr != newPtr && i > 1 {
			fmt.Printf(" <-- ⚠️ 底層陣列已更換 (擴容)\n")
		} else {
			fmt.Printf("\n")
		}
	}

	fmt.Println("\n知識補充：")
	fmt.Println("- len(s): Slice 目前包含的元素個數（杯子裡裝了多少水）")
	fmt.Println("- cap(s): Slice 底層陣列總共可容納的個數（杯子本身的大小）")

	fmt.Println("\n結論：")
	fmt.Println("1. 當 append 超過 cap 時，Go 會配置一塊新的、更大的記憶體。")
	fmt.Println("2. 舊陣列的資料會被複製到新陣列中（這會造成 CPU 效能損耗）。")
	fmt.Println("3. Slice 的指標會更新為指向新陣列。")
	fmt.Println("4. 為了效能，若已知資料量，建議使用 make([]int, 0, 預期容量) 預分配空間。")
}
