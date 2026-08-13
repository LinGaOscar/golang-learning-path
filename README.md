# golang-learning-path

個人 Go 語言學習練習集：一系列獨立、單檔案的小程式，用來練習 Go 的基礎語法、結構體/介面與併發模型。**不是一個整合式應用程式**，每個 `.go` 檔案都是獨立的 `package main`，各自有自己的 `func main()`。

## 目前內容

| 檔案 | 主題 |
|---|---|
| `main.go` | 環境確認 / 最小 hello world |
| `variables.go` | 變數宣告、常量、`iota` 列舉 |
| `fizzbuzz.go` | 控制流程（`for`、`switch`） |
| `slice_growth.go` | Slice 的 len/cap 成長與記憶體重新分配 |
| `shapes.go` | Struct、Method、Interface（`Shape`、`Rectangle` 等） |
| `concurrency.go` | Goroutine、`sync.WaitGroup`、`sync.Mutex` |
| `db_errors.go` | 自訂錯誤型別、錯誤包裝與判斷 |

`QuickStart.md` 是搭配練習內容整理的中文複習指南（觀念、術語、思考題），非程式碼文件。目前沒有任何 `_test.go` 測試檔。

## 學習階段規畫

### 第一階段：基礎語法與 AI 協作 (Week 1-2)
目標： 熟悉 Go 的簡潔邏輯，並學會讓 AI 協作工具幫你寫出符合規範的 Code。

重點內容：
- 環境安裝：`go mod` 管理。
- 基本型別：變數、Slices（切片）、Maps。
- 控制流程：`if`、`for`、`switch`（Go 只有 `for` 迴圈）。

對應練習：`variables.go`、`fizzbuzz.go`、`slice_growth.go`

### 第二階段：結構體與介面 (Week 3-4)
目標： 理解 Go 的組合設計哲學（Composition over Inheritance）。

重點內容：
- Structs 與 Methods。
- Interfaces（重點）：Go 靈魂所在，理解如何實現隱式介面。
- Error Handling：理解為何 Go 總是 `if err != nil`。

對應練習：`shapes.go`、`db_errors.go`

### 第三階段：併發程式設計 Concurrency (Week 5-6)
目標： 掌握 Go 的最強大武器。

重點內容：
- Goroutines（輕量級執行緒）。
- Channels（通訊機制）。
- Select 與 Mutex（鎖）。

對應練習：`concurrency.go`

## 如何執行單一練習

每個檔案都是獨立的 `package main`，**不能**用 `go run .` 或 `go build ./...` 執行整個目錄（會因為多個 `main` 重複宣告而編譯失敗）。請針對單一檔案執行：

```bash
go run main.go
go run variables.go
go run fizzbuzz.go
go run slice_growth.go
go run shapes.go
go run concurrency.go
go run db_errors.go
```

編譯成執行檔：

```bash
go build variables.go
./variables   # Windows 為 variables.exe
```

模組管理：

```bash
go mod tidy
```

格式化代碼：

```bash
go fmt ./...
```
