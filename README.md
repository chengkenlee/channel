# channel例子

## Support

```
package main

import (
	"fmt"
	"sync"
)

var w sync.WaitGroup

func main() {
	done := make(chan int)

	w.Add(3)
	go t1("one", done)
	go t2("two", done)
	//go t3("three", done)

	done <- 1
	w.Wait()
}

func t1(name string, done chan int) {
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
func t2(name string, done chan int) {
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

```


## Best Regards

* Hat tip to anyone whose code was used
* Inspiration
* etc
