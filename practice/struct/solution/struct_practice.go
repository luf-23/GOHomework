package solution

import (
	"fmt"
	"time"
)

/*
通道通信：
在通过通道进行通信时，空结构体可以作为简单的信号来传递，用于通知某个事件的发生，而不
需要传递任何具体的数据。例如，在一个并发程序中，一个 goroutine 完成任务后向另一个 goroutine 发
送信号。
*/

func StructPractice01() {
	work := func(d chan struct{}) {
		// 模拟工作
		time.Sleep(2 * time.Second)
		fmt.Println("Worker finished")
		// 发送完成信号
		d <- struct{}{}
	}
	done := make(chan struct{})
	fmt.Println("Starting worker")
	go work(done)
	<-done // 等待工作完成信号
	fmt.Println("Done")
}

/*
集合成员表示：
当你只关心某个元素是否存在于集合中，而不需要存储关于该元素的任何额外数据时，空结
构体非常有用。例如，在一个需要判断一组唯一标识符是否存在的场景中，可以使用 map 结合空结构体
来实现。
*/

func StructPractice02() {
	uniqueIds := make(map[string]struct{})
	uniqueIds["number001"] = struct{}{}
	uniqueIds["number002"] = struct{}{}
	uniqueIds["number004"] = struct{}{}
	_, exists1 := uniqueIds["number003"]
	_, exists2 := uniqueIds["number004"]
	fmt.Println("number003 exists:", exists1)
	fmt.Println("number004 exists:", exists2)

}
