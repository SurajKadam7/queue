# queue

### how to use

```go
package main

import (
   "fmt"
   "github.com/SurajKadam7/queue"
)

func main() {
	queue := queue.New[int64]()
	queue.Push(10)
	queue.Push(20)
	queue.Push(30)
	queue.Push(40)

	for queue.lst.Len() > 0 {
		fmt.Println(queue.Pop())
	}
}
```