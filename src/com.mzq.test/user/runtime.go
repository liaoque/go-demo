package user

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func TestRunTime()  {
	// go安装路径
	fmt.Println(runtime.GOROOT())
	// 获取可使用cpu逻辑核数
	fmt.Println(runtime.NumCPU())
	// 设置可使用cpu逻辑核数
	fmt.Println(runtime.GOMAXPROCS(1))

	fmt.Println(runtime.NumCPU())
	//获取例程数量
	fmt.Println(runtime.NumGoroutine())

	group := &sync.WaitGroup{}
	group.Add(10)

	i := 0
	for ; i < 10; i++ {
		go func() {
			time.Sleep(10)
			group.Done()
		}()
	}

	fmt.Println(runtime.NumGoroutine())

	group.Wait()
	fmt.Println(runtime.NumGoroutine())


}
