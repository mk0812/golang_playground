package main

import (
    "fmt"
    "sync"
    "time"
)

func asyncFunc(s string,wg *sync.WaitGroup) {
	defer wg.Done() // WaitGroupを最後に完了しないといけない。
    for i := 0; i < 5; i ++ {
        time.Sleep(100 * time.Millisecond)
        fmt.Println(s)
    }
}

func main(){
	 var wg sync.WaitGroup

    // goルーチンの関数の実行
    for i := 0 ; i < 3; i ++ {
        wg.Add(1) // goルーチンを実行する関数分だけAddする。
        str := fmt.Sprintf("Go routine (no: %v)", i)
        go asyncFunc(str, &wg)
    }

    // goルーチンで実行される関数が終了するまで待つ。
    wg.Wait()
}