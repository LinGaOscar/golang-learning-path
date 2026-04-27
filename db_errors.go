package main

import (
	"errors"
	"fmt"
	"time"
)

// ==========================================
// 1. 定義自定義錯誤型別 (Custom Error Type)
// ==========================================

// DatabaseError 包含錯誤發生的上下文細節
type DatabaseError struct {
	Op      string    // 發生錯誤的操作 (例如: Connect, Query)
	User    string    // 嘗試操作的使用者
	When    time.Time // 發生時間
	Message string    // 錯誤說明內容
	Err     error     // 原始錯誤 (用於示範 errors.Is 判斷的核心錯誤)
}

// 實作 error 介面 (這使 DatabaseError 成為一個 error)
func (e *DatabaseError) Error() string {
	return fmt.Sprintf("[%s] 使用者 %s 在執行 %s 時發生錯誤: %s", 
		e.When.Format("15:04:05"), e.User, e.Op, e.Message)
}

// 實作 Unwrap 介面
// 這非常重要！它讓 errors.Is 和 errors.As 能夠「看穿」外層包裝，找到內層的原始錯誤。
// 語法拆解：
// func          -> 宣告函式的關鍵字
// (e *DatabaseError) -> 接收者 (Receiver)，表示此為掛載在 DatabaseError 指標型別下的「方法」
// Unwrap        -> 方法名稱，Go 標準庫約定使用此名稱來進行「錯誤解包」
// ()            -> 此方法不需傳入任何參數
// error         -> 回傳型別，必須回傳一個符合 error 介面的值
func (e *DatabaseError) Unwrap() error {
	return e.Err // 回傳內層包裝的原始錯誤
}

// ==========================================
// 2. 定義哨兵錯誤 (Sentinel Errors)
// ==========================================
// 這些通常用於 errors.Is 的比對，代表錯誤的「種類」。

var (
	ErrTimeout    = errors.New("連線超時 (Connection Timeout)")
	ErrPermission = errors.New("權限不足 (Permission Denied)")
)

// ==========================================
// 3. 模擬資料庫連線函式
// ==========================================

func ConnectDB(user string, scenario string) error {
	// 基礎的包裝結構
	dbErr := &DatabaseError{
		Op:   "Connect",
		User: user,
		When: time.Now(),
	}

	switch scenario {
	case "timeout":
		dbErr.Message = "伺服器無回應，已達到最大等待時間"
		dbErr.Err = ErrTimeout // 將哨兵錯誤包裝進去
		return dbErr
	case "permission":
		dbErr.Message = "無效的存取權限，請確認 API Key"
		dbErr.Err = ErrPermission // 將哨兵錯誤包裝進去
		return dbErr
	default:
		// 模擬成功連線
		return nil
	}
}

// ==========================================
// 4. 主程式：示範如何判斷錯誤
// ==========================================

func main() {
	scenarios := []string{"timeout", "permission", "ok"}

	for _, scene := range scenarios {
		fmt.Printf(">> [情境測試: %s]\n", scene)
		
		err := ConnectDB("Oscar_Lin", scene)

		if err != nil {
			// --- 使用 errors.Is ---
			// 判斷該錯誤「是否包含」某個特定的哨兵錯誤
			if errors.Is(err, ErrTimeout) {
				fmt.Println("   [errors.Is] ⚠️  偵測到連線超時！建議：檢查防火牆或重試。")
			} else if errors.Is(err, ErrPermission) {
				fmt.Println("   [errors.Is] 🚫 偵測到權限錯誤！建議：檢查帳號權限。")
			}

			// --- 使用 errors.As ---
			// 嘗試將 err 轉換為特定的「結構體指標」，以提取詳細欄位
			var targetErr *DatabaseError
			if errors.As(err, &targetErr) {
				fmt.Printf("   [errors.As] 🔍 提取詳細資訊 -> 操作: %s, 使用者: %s, 時間: %v\n", 
					targetErr.Op, targetErr.User, targetErr.When)
				fmt.Printf("   原始訊息: %v\n", targetErr)
			}
		} else {
			fmt.Println("   ✅ 連線成功！沒有發生錯誤。")
		}
		fmt.Println("-------------------------------------------")
	}
}
