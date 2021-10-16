package main

import (
	"fmt"
	"strconv"
	"syscall/js"
)

// https://leetcode-cn.com/submissions/detail/41908622/
func fib(N int) int {
	if N < 2 {
		return N
	}
	cache := make([]int, N+1)
	cache[0], cache[1] = 0, 1
	for i := 2; i <= N; i++ {
		cache[i] = cache[i-1] + cache[i-2]
	}
	return cache[N]
}

// fibFunc 是 fib 的封装，从 args[0] 获取入参，计算结果用 js.ValueOf 转化并返回
// js.Value 将 JavaScript 的值转化为 Go 的值
// js.ValueOf 将 Go 的值转化为 JavaScript 的值
// this 是 JavaScript 中的 this
// args 是 JavaScript 中调用函数的参数列表
func fibFunc(this js.Value, args []js.Value) interface{} {
	s := args[0].String()
	fmt.Println("s is", s)
	num, _ := strconv.Atoi(s)
	answer := fib(num)
	fmt.Println("result is", answer)
	return js.ValueOf(answer)
}

func main() {
	done := make(chan int, 0)
	// Register Functions，通过 js.Global().Set() 方法，将 fibFunc 函数注册到 JavaScripts 中，以便在 Html 中能够调用
	// js.FuncOf 将 Go 的函数转化为可以在 JavaScript 中调用的函数
	js.Global().Set("fibFunc", js.FuncOf(fibFunc))
	<-done
}
