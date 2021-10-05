package main

import (
	"strconv"
	"syscall/js"
)

//export fib
func Fib(N int) int {
	println("GO::EVENT Fib!")

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

//export update
func Update() {
	println("GO::EVENT Update!")

	document := js.Global().Get("document")
	numStr := document.Call("getElementById", "num").Get("value").String()
	num, _ := strconv.Atoi(numStr)
	result := Fib(num)
	document.Call("getElementById", "answer").Set("value", result)
}

func main() {
	println("GO:: Fib!")
}
