package main

import (
	"fmt"
)
var pk_x = 2		//パッケージ変数の定義方法	:=だとエラー


func hoge() {
	var x1 int
	x1 = 1
	x2 := 1		//一個ならこれを一番使う
	var x3 = 1	//１個ならこれは冗長だが
	var (	//複数なら見やすくていい
		x4_1 = 1
		x4_2 = 1
	)
	fmt.Printf("%d,%d,%d,%d,%d,\n", x1, x2, x3, x4_1, x4_2)		//定義したんなら使わないとエラー

}