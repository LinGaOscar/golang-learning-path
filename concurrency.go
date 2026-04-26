package main

import (
	"fmt"
	"sync"
)

/*
【 併發控制執行流程說明 】
1. 初始化：設定共用資源 (counter) 與同步工具 (WaitGroup, Mutex)。
2. 任務分發：使用 Add(1000) 告訴 WaitGroup 我們有 1000 個任務要跑。
3. 併發執行：啟動 1000 個 Goroutine，各跑各的。
4. 互斥存取：Goroutine 內部使用 Lock/Unlock 確保 counter 每次只能被一個人加 1。
5. 完成回報：每個 Goroutine 結束前呼叫 Done()。
6. 阻塞等待：主執行緒 Wait() 停在原地，直到 1000 個 Done() 全部入帳。
7. 結案：列印最終結果。
*/

func main() {
	// [步驟 1] 初始化資源
	counter := 0
	const iterations = 1000

	// [步驟 1] 初始化同步工具
	var wg sync.WaitGroup
	var mu sync.Mutex

	// [步驟 2] 註冊任務總數
	wg.Add(iterations)

	for i := 0; i < iterations; i++ {
		// [步驟 3] 啟動併發任務 (Goroutine)
		go func() {
			// [步驟 5] 使用 defer 確保任務完成後一定會回報 Done()
			defer wg.Done()
			
			// [步驟 4] 進入臨界區 (排隊等廁所)
			mu.Lock()   
			
			// 修改共用資源
			counter++   
			
			// [步驟 4] 離開臨界區 (出廁所開門)
			mu.Unlock() 
		}()
	}

	// [步驟 6] 阻塞等待，直到 WaitGroup 計數器歸零
	fmt.Println("等待所有任務完成中...")
	wg.Wait()
	
	// [步驟 7] 所有任務完成，列印結果
	fmt.Printf("最終計數值: %d (預期值: %d)\n", counter, iterations)
}

/*
【 核心概念圖解補充 】

1. 為什麼沒加鎖會低於 1000？ (Race Condition)
   - 假設 Counter = 10
   - Goroutine A 讀取到 10
   - Goroutine B 同時也讀取到 10 (此時 A 還沒寫回)
   - A 計算 10+1 = 11 並寫回
   - B 也計算 10+1 = 11 並寫回
   - 結果：明明加了兩次，但 Counter 卻只變成 11，其中一次加法消失了！

2. Mutex (鎖) 的運作
   - Lock()：就像進廁所鎖門。門鎖了，後面的 Goroutine 必須排隊。
   - Unlock()：就像出廁所開門。讓下一個排隊的人可以進來。
   - 確保了「讀取 -> 加 1 -> 寫回」這三個動作是一口氣完成的 (原子性)。

3. WaitGroup 的運作
   - Add(1000)：門口櫃檯放了 1000 枚代幣。
   - Done()：員工完成工作，交回一枚代幣。
   - Wait()：老闆 (Main) 在門口等，直到 1000 枚代幣全部收齊才准下班。
*/
