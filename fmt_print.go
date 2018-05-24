package main

import (
	"fmt"
)

func main(){
	//PrintlnとPrintfは標準出力へと出力する
	fmt.Println("Hello World")
	fmt.Println("Hello", "New", "World")  //複数の引数をとってスペースを空けて表示
	fmt.Println("")  //改行

	fmt.Printf("数値は%d\n", 5)
	fmt.Printf("10進数 : %d  2進数 : %b  8進数 : %o  16進数 : %x\n", 17, 17, 17 , 17)
	fmt.Printf("%d年%d月%d日%s曜日\n", 2018, 5, 23)		//足りない時も
	fmt.Printf("%d年%d月%d日%s曜日\n", 2018, 5, 23, "水", "わっしょい")		//余ってる時も
	fmt.Println("")
	fmt.Printf("数値=%v  文字列=%v  配列=%v\n", 5, "Golang", [...]int{1,2,3}) //%vはなんでも埋め込む %#vはリテラル表現をする
	fmt.Printf("数値=%T  文字列=%T  配列=%T\n", 5, "Golang", [...]int{1,2,3}) //%Tは型の表示をする

	//print, printlnは標準エラー出力へと出力する
	print("Hello World")  //改行なし
	println("Hello World")  //改行あり
	//これらはfmtをつけない！！
}