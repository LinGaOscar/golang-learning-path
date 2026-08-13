# dev.md

本機開發筆記。這是個人 Go 學習練習集，不是產品專案，以下僅記錄實際需要的環境與跑法。

## 環境需求

- Go 1.26+（`go.mod` 宣告 `go 1.26.2`；開發機上實測 `go1.26.4` 可正常編譯）
- 無外部相依套件（`go.mod` 目前沒有 `require` 項目），不需要額外安裝套件或跑 `go mod download`

## 執行單一練習

檔案彼此獨立、都是 `package main` + 自己的 `func main()`，**無法**整包一起 build/run。針對想跑的檔案執行：

```bash
go run <檔名>.go
# 例如
go run shapes.go
```

若需要編譯出執行檔：

```bash
go build <檔名>.go
./<檔名>
```

## 結構備註

- 檔案全部平放在 repo 根目錄，沒有子資料夾、沒有共用套件、沒有 `internal/` 或 `cmd/` 結構。
- 目前沒有任何 `_test.go`，因此沒有 `go test` 可跑。
- `QuickStart.md` 是複習用的中文筆記文件，與程式碼無直接關聯，修改程式碼不需要同步更新它。

## 已知限制

- `go build ./...`、`go vet ./...`、`go run .` 都會因為多個檔案重複宣告 `main` 而失敗，這是刻意的練習集結構，不是 bug。
