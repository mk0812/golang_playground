package pratice

import (
	"fmt"
	"time"
)

func asyncFunc(s string){
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millsecond)
		fmt.Println(s)
	}
}

// メイン関数
func pratice() {
    // goルーチンの関数の実行
    for i := 0 ; i < 3; i ++ {
        str := fmt.Sprintf("Go routine (no: %v)", i)
        go asyncFunc(str)
    }

    // time.sleepしないとgoルーチンの実行完了よりも先にメイン関数の処理が完了してしまう。
    time.Sleep(1 * time.Second)
}