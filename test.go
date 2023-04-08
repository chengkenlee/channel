package main

import "fmt"

func t3(name string, done chan int) {
	defer w.Done()
	for {
		value, ok := <-done
		if !ok {
			fmt.Println(name, "关闭", done)
			return
		}
		if value == 200 {
			close(done)
			fmt.Println(name, "我们赢了，关闭通道", value, done)
			return
		}
		fmt.Println(name, "完成", value, done)
		value++
		done <- value
	}
}
