🚀 Golang 學習階段規畫
第一階段：基礎語法與 AI 協作 (Week 1-2)
目標： 熟悉 Go 的簡潔邏輯，並學會讓 Antigravity 幫你寫出符合規範的 Code。

重點內容：

環境安裝：go mod 管理。

基本型別：變數、Slices (切片)、Maps。

控制流程：if, for, switch（Go 只有 for 迴圈）。

GitHub 任務：

建立一個 go-basics 倉庫。

Antigravity 應用： 使用 "Agent Manager" 視窗下指令：「幫我寫一個處理學生名單的 CLI 工具」，觀察它如何處理 Slice。

第二階段：結構體與介面 (Week 3-4)
目標： 理解 Go 的組合設計哲學（Composition over Inheritance）。

重點內容：

Structs 與 Methods。

Interfaces (重點)： 這是 Go 靈魂，理解如何實現隱式介面。

Error Handling：理解為何 Go 總是 if err != nil。

GitHub 任務：

實作一個「簡單銀行帳戶系統」。

Antigravity 應用： 讓 Agent 幫你為 Struct 寫單元測試 (_test.go)，並觀察它如何利用 Artifacts 顯示測試覆蓋率。

第三階段：併發程式設計 Concurrency (Week 5-6)
目標： 掌握 Go 的最強大武器。

重點內容：

Goroutines (輕量級執行緒)。

Channels (通訊機制)。

Select 與 Mutex (鎖)。

GitHub 任務：

實作一個「多執行緒網頁爬蟲」。

Antigravity 應用： 讓 Agent 幫你檢查 Race Condition（競態競爭）。