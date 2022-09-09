package main

import (
	"fmt"
	"syscall/js"
)

func main() {
	// デバッグコンソールに出力
	println("Hello, WebAssembly!")
	println("Again, Hello, WebAssembly!")
	// fmtのインポートチェック
	fmt.Println("fmt is imported.")

	// jsのグローバルオブジェクトの取得
	window := js.Global()
	message := window.Get("document").Call("getElementById", "message")

	message.Set("innerText", "Hello, WebAssembly")

	message.Call("addEventListener", "click", js.FuncOf(func(js.Value, []js.Value) interface{} {
		result := window.Get("document").Call("getElementById", "result")
		result.Set("innerHTML", "Clicked!!")
		return nil
	}))

	select {}
}
