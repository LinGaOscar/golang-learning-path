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

---

## 🛠️ 如何執行 Go 程式

在終端機 (Terminal) 中，你可以使用以下指令來執行、編譯或管理你的 Go 專案：

### 1. 直接執行 (開發常用)
如果你只想快速看到結果，不需要產生執行檔：
```bash
go run main.go
# 或者執行專案中的其他檔案
go run variables.go
go run fizzbuzz.go
```

### 2. 編譯成執行檔 (部署常用)
將程式碼編譯成二進位執行檔：
```bash
go build main.go
# 在 Windows 會產生 main.exe，在 macOS/Linux 會產生 main
# 執行編譯後的程式：
./main
```

### 3. 模組管理
如果你有引用外部套件或是需要整理 `go.mod`：
```bash
# 自動下載缺少的套件並移除不使用的套件
go mod tidy
```

### 4. 格式化代碼
讓你的程式碼符合 Go 官方規範：
```bash
go fmt ./...
```