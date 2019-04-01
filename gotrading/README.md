# はじめに

思いつき記事です。
GO言語を勉強して、はや半年くらいたちました。

やったこととしては

+ GO言語の公式ドキュメントを読む。
+ A Tour of Goを２周した。
+ Udemyの「Go入門＋応用」を一通りやった。

がしかし！！！！！！！
全く覚えられない！！！！！

なので、絶賛アウトプットして知識を深めている最中です。
(GO言語 + Vue(Elm)で個人的に何か作ろうプロジェクト進行中ですが)

この記事ではGO言語で株価を取得してみようというタイトルで進めていきます。

# 経緯
友達が「来月から株始めるわ！」と言ってきたので、「じゃあ、株取得してみるわ（APIで）。」となりました。

PythonでAPI取得して、それをグラフ化するというのもいいかなと思いましたが、Qiitaに色々ありましたので、GO言語で挑戦してみます。

# 手順
GO言語って便利ですね。
サードパーティで「quandl」（株価取得ライブラリ）がありました。

[参考：package quandl](https://godoc.org/github.com/DannyBen/quandl)

最初は上記の参考URLの手順で進めていきます。

## インストール
まずはquandlパッケージのインストールから始めましょう。

```
$ go get github.com/DannyBen/quandl
```

## コード
とりあえず簡単なサンプルコードからみていきましょう。
正味、下のURL（Example）みたほうが早い説あります。
[https://github.com/DannyBen/quandl/blob/master/quandl_test.go](https://github.com/DannyBen/quandl/blob/master/quandl_test.go)

```go
package main

import (
	"os"

	"github.com/DannyBen/quandl"
)

func main() {
	quandl.APIKey = os.Getenv("QUANDL_KEY") //APIKeyを取得

	v := quandl.Options{} //Option
	//Set registers a key=value pair to be sent in the Quandl request
	v.Set("trim_start", "2017-01-01")
	v.Set("trim_end", "2017-02-02")
	v.Set("column_index", "4")

	data, _ := quandl.GetSymbol("WIKI/AAPL", v)             //第１引数に指定した銘柄のシンボルを取得
	data2, _ := quandl.GetSymbolRaw("WIKI/AAPL", "json", v) //第１引数に指定した銘柄のシンボルをjson形式で取得

//ioutil.WriteFileで保存
}
```
### 結果
```
[[2017-02-02 128.53] [2017-02-01 128.75] [2017-01-31 121.35] [2017-01-30 121.63] [2017-01-27 121.95] [2017-01-26 121.94] [2017-01-25 121.88] [2017-01-24 119.97] [2017-01-23 120.08] [2017-01-20 120] [2017-01-19 119.78] [2017-01-18 119.99] [2017-01-17 120] [2017-01-13 119.04] [2017-01-12 119.25] [2017-01-11 119.75] [2017-01-10 119.11] [2017-01-09 118.99] [2017-01-06 117.91] [2017-01-05 116.61] [2017-01-04 116.02] [2017-01-03 116.15]]
```
色々試して実行したコードなんですが、とりあえず日付と株価が取得できました。
他にも、```GetSearch```関数（文字列検索→銘柄の検索かな？）や```ToColumns```関数()など色々なメソッドがあるので、参考にしてみてください（僕も頑張って色々形になるように頑張ります）。

## メソッド追記
●```GetSymbolRaw``` ... 取得したデータをJSONで返す
●```GetListRaw``` ... 銘柄の値（リスト型）を返す
●```GetSearchRaw``` ... サーチで当てはまった銘柄の各値をJSON(XML)の型で返す
●```StringColumn``` ... インターフェース型のコラムを文字列型のコラムに変更

# まとめ
事前にPython(+Pandas)でやってみたんですが、もちろんPythonのほうが便利かもしれません。
GO言語でこのような挑戦（大した挑戦はしていない）して、Godocも徐々に読めるようになって、これからだなと思いました（白目）。
英語読めてデメリットなし。
