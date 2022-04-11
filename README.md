# LanRen.Gin
LanRen.Gin (Golang)

[本地測試](http://127.0.0.1:8888)

# 對程式碼進行自動測試
```
go test

or

go test .\src\test

or

go test .\src\test -v
```
1. 當前目錄或指定目錄中檔名為 _test.go 結尾
2. 必須 import "testing"
3. 撰寫形式 func TestXxx(t *testing.T) 的函式
4. -v 可以看到 t.Error("error") 與 t.Log("log") 的內容
5. t.Skip()以下是不會執行的

# 效能測試
```
go test -bench.

or

go test -v -bench.
```
1. 函式必須是 func BenchmarkXxx(b *testing.B)
2. 目前與 test 寫一起會無法執行(暫時分開寫)
3. 進入到指定路徑內執行上面指令