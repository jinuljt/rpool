<p align='right'>
<a href='https://drone.io/github.com/jinuljt/rpool/latest'><img width='80' align='absmiddle' height='18' src='https://drone.io/github.com/jinuljt/rpool/status.png' /></a>
<a href='http://godoc.org/github.com/jinuljt/rpool'><img width='94' align='absmiddle' height='18' alt='GoDoc' src='https://godoc.org/github.com/jinuljt/rpool?status.png' /></a>
</p>


## rpool  ##

rpool 是一个go协程的资源池。它保证同时只有指定个数的协程在工作。


## 为什么会存在这个库 ##

目的是想要控制运行的goroutine的数量，同时又能够达到WaitGroup的效果。

## 如何使用rpool ##

```
//初始化一个rpool，指定最多同时执行的100个协程
rp = rpool.NewRPool(100)

//运行1000个协程
for i := 0; i < 1000; i ++ {
   rp.Add() //当执行rp.Add() 时，如果pool中运行的协程数量已经到达指定数量（100），会block，直到其他协程释放资源。
   go test()
}

rp.Wait() //等待所有的协程执行完毕，如果不关心是否都执行完毕，这个函数可以不调用
fmt.Println("Done")

func test() {
    time.Sleep(1 * time.Millisecond)
    rp.Done() //当一个协程执行完毕，必须要调用Done() 否则无法释放资源，也就是会导致Add() 函数永远block
}
```
