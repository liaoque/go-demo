package user

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)



func printChars(prefix string, group *sync.WaitGroup)  {
	defer group.Done()
	//defer 函数最后执行
	for ch := 'A'; ch <= 'Z'; ch++ {
		fmt.Printf("%s:%c\n", prefix, ch)
		runtime.Gosched()
	}
}

func printChars2(prefix string)  {
	//defer group.Done()
	//defer 函数最后执行
	for ch := 'A'; ch <= 'Z'; ch++ {
		fmt.Printf("%s:%c\n", prefix, ch)
		runtime.Gosched()
	}
}

func printChars3(group *sync.WaitGroup)  {
	defer group.Done()
}

func Test()  {
	group := &sync.WaitGroup{}
	n := 3

	group.Add(n)

	for i:=0; i< n;i++ {
		go printChars(fmt.Sprintf("gochars%02d", i), group)
	}

	group.Wait()
	fmt.Printf("over")

}

func Test2()  {
	go printChars2("ssssssssss")
	go printChars2("dddddddddd")
	printChars2("ffffffffff")
}

func Test3()  {
	group := &sync.WaitGroup{}
	n := 3

	group.Add(n)

	for i:=0; i< n;i++ {
		go func(i int) {
			fmt.Println(i)
			group.Done()
		}(i)
	}

	group.Wait()
	fmt.Printf("-----------------------")


	group.Add(n)


	for i:=0; i< n;i++ {
		go func(i int) {
			fmt.Println(i)
			group.Done()
		}(i)
	}

	group.Wait()

	fmt.Printf("over")
}

var counter int= 0
var counter32 int32= 0
func IncrMutex(mutex *sync.Mutex, group *sync.WaitGroup)  {
	mutex.Lock()
	defer func() {
		mutex.Unlock()
		group.Done()
	}()
	for i := 0; i < 1000; i++ {
		counter++
		runtime.Gosched()
	}
}

func DescMutex(mutex *sync.Mutex, group *sync.WaitGroup)  {
	mutex.Lock()
	defer func() {
		mutex.Unlock()
		group.Done()
	}()
	for i := 0; i < 1000; i++ {
		counter--
		runtime.Gosched()
	}
}

func TestLock()  {
	group := &sync.WaitGroup{}

	mutex := &sync.Mutex{}

	for i := 0; i < 5; i++ {
		group.Add(2)
		go IncrMutex(mutex, group)
		go DescMutex(mutex, group)
	}
	group.Wait()
	fmt.Println(counter)
}


func TestLockOne()  {
	once := &sync.Once{}

	for i := 0; i <= 10; i++ {
		func(){
			fmt.Println("11111111111", i)
		}()

		// 只会执行一次。 可用初始化之类的东西
		once.Do(func() {
			fmt.Println("2222222", i)
		})


	}

}

func incrAtomic(group *sync.WaitGroup)  {
	defer func() {
		group.Done()
	}()
	for i := 0; i < 1000; i++ {
		atomic.AddInt32(&counter32, 1)
		runtime.Gosched()
	}
}

func DescAtomic(group *sync.WaitGroup)  {
	defer func() {
		group.Done()
	}()
	for i := 0; i < 1000; i++ {
		atomic.AddInt32(&counter32, -1)
		runtime.Gosched()
	}
}

func TestAtomic()  {
	group := &sync.WaitGroup{}

	for i := 0; i < 5; i++ {
		group.Add(2)
		go incrAtomic(group)
		go DescAtomic(group)
	}
	group.Wait()
	fmt.Println(counter)
}

func printChars4(prefix string,channel chan string)  {
	//defer 函数最后执行
	for ch := 'A'; ch <= 'Z'; ch++ {
		fmt.Printf("%s:%c\n", prefix, ch)
		runtime.Gosched()
	}
	channel <- prefix
}

func TestPipeline()  {
	channel := make(chan string)
	channel2 := make(chan int, 10)

	fmt.Printf("%T, %v, %d\n", channel, channel, len(channel))
	fmt.Printf("%T, %v, %d\n", channel2, channel2, len(channel2))


	n:= 10
	//tag2 := <-channel
	//fmt.Println("over: ", tag2)


	for i :=0; i < n; i++ {
		go printChars4(fmt.Sprintf("gochars%02d", i), channel)
	}

	for i :=0; i < n; i++  {
		tag := <-channel
		fmt.Println("over: ", tag)
	}

	fmt.Println("over")





	ints := make(chan int)
	go func() {
		for e := range ints{
			fmt.Println("eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee")
			fmt.Println(e)
		}
		channel <- "0"
	}()

	go func() {
		for i := 0; i < 5; i++ {
			ints <- i
		}
		close(ints)
	}()

	<-channel

	//for  {
	//	tag := <-ints
	//	//if !ok {
	//	//	fmt.Println("over")
	//	//	break
	//	//}else {
	//		fmt.Println(tag)
	//	//}
	//}

	onlyWriteRead := make(chan int)
	go func(onlyRead <-chan int) {
		reads := <-onlyRead
		fmt.Println(reads)

	}(onlyWriteRead)

	go func(onlyWiter chan<- int) {
		onlyWiter<-10
	}(onlyWriteRead)



defer func() {
	time.Sleep(100)
	// 管道 写的时候要注意， 读写管道必须对应上 不然会出现异常退出
	close(onlyWriteRead)
	close(channel)
	close(channel2)

}()

	ch01 := make(chan int, 1)
	ch02 := make(chan int, 1)

	go func() {
		time.Sleep(5 * time.Second)
		ch01 <- 0
	}()

	go func() {
		duration := time.Duration(rand.Int31n(10)) * time.Second
		time.Sleep(duration)
		fmt.Println("sleep:", duration)
		ch02 <- 0
	}()

	// 通过 select case
	// 定义两个 chan, 其中一个做超时检测

	select {
	case <- ch02:
		fmt.Println("ch02")
	case <- ch01:
		fmt.Println("超时")
	default:
		fmt.Println("aaaa")
	}

	select {
	case <- ch02:
		fmt.Println("ch02")
	case <- time.After(5 * time.Second):
		fmt.Println("超时")
	default:
		fmt.Println("aaaa")
	}



}


