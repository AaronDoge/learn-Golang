
Go语言并发
=====================

Go语言里的并发指的是能让某个函数独立于其他函数运行的能力。当一个函数创建为goroutine时，Go会将其视为一个独立的工作单元。

Go语言的的调度器是一个复杂的软件，能管理被创建的所有goroutine并为其分配执行时间。这个调度器在操作系统之上，将操作系统的线程与语言运行时的逻辑处理器绑定，并在逻辑处理器上运行。


1.并发与并行
-------------------------

 - 什么是操作系统的进程(thread)和线程(Process)? 

如图，展示了一个包含所有可能分配的常用资源的的进程。这些资源包括但不限于内存地址空间、文件和设备的句柄以及进程。一个线程是一个执行空间，这个空间会被操作系统调度来运行函数中所写的代码。每个进程 至少包含一个线程，每个进程的初始线性被称为主线程。因为执行这个线程的空间是应用程序的本身空间，所以当主线程终止时，应用程序也会终止。操作系统将线程调度到某个处理器上运行，这个处理器并不一定是进程所在的处理器。不同操作系统使用的线程调度算法一般都不一样，但是这种不同会被操作系统屏蔽，并不会展示给程序员。

操作系统会在物理处理器上调度线程来运行，而Go语言会在运行时逻辑处理器上调度goroutine来运行。

Go语言运行时会把goroutine调度到逻辑处理器上运行。这个逻辑处理器绑定到唯一的操作系统线程。当goroutine可以运行的时候，会被放入逻辑处理器的执行队列中。

当goroutine执行了一个阻塞的系统调用时，调度器会将这个线程与处理器分离，该线程会继续阻塞，等待系统调用的返回。与此同时，这个逻辑处理器就是去了用来运行的线程，此时调度器会创建一个新线程，并将其绑定在这个逻辑处理器上。之后，调度器会从本地运行队列中选择另一个goroutine来运行。一旦之前被阻塞的系统调用执行完成并返回，对应的goroutine会放回到本地运行队列，而之前的线程会保存好，以便之后可以继续使用。


2.Goroutine
-----------------------------



```go
//这个示例程序展示如何促进goroutine

package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	//分配一个逻辑处理器
	runtime.GOMAXPROCS(1)

	//
	var wg sync.WaitGroup
	wg.Add(2)

	fmt.Println("Start Goroutines")

	//声明一个匿名函数，并创建一个goroutine
	go func() {
		//在函数退出时调用Done来通知main函数工作已经完成
		defer wg.Done()

		//显示字母表3次
		for count := 0; count < 3; count++ {
			for char := 'a'; char < 'a'+26; char++ {
				fmt.Printf("%c ", char)
			}
		}
	}()

	//声明一个匿名函数，并创建一个goroutine
	go func() {
		//在函数退出时调用Done来通知main函数工作已经完成
		defer wg.Done()

		//显示字母3次
		for count := 0; count < 3; count++ {
			for char := 'A'; char < 'A'+26; char++ {
				fmt.Printf("%c ", char)
			}
		}
	}()

	//等待goroutine结束
	fmt.Println("Waiting To Finish...")
	wg.Wait()

	fmt.Println("\n Terminating Program")
}

```