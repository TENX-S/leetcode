package others

import (
    "fmt"
    "sync"
)

// 判断奇偶数
func alternatePrint1(target int) {
    c := make(chan int)
    defer close(c)

    var wg sync.WaitGroup
    defer wg.Wait()

    wg.Add(1)
    go func() {
        defer wg.Done()
        for i := 0; i < target+1; i++ {
            c <- 1
            fmt.Println("hi")
            if i%2 == 1 {
                fmt.Println("A: ", i)
            }
        }
    }()

    wg.Add(1)
    go func() {
        defer wg.Done()
        for i := 0; i < target+1; i++ {
            <-c
            if i%2 == 0 {
                fmt.Println("B: ", i)
            }
        }
    }()
}

// 不判断奇偶数
func alternatePrint2(target int) {
    odd := make(chan int)
    even := make(chan int)
    var wg sync.WaitGroup
    curr := 0

    wg.Add(1)
    go func() {
        defer wg.Done()
        defer close(even)
        for k := range even {
            if k == -1 {
                return
            }
            if curr > target {
                odd <- -1
                return
            }
            fmt.Println("A: ", curr)
            curr++
            odd <- 0
        }
    }()

    wg.Add(1)
    go func() {
        defer wg.Done()
        defer close(odd)
        for k := range odd {
            if k == -1 {
                return
            }
            if curr > target {
                even <- -1
                return
            }
            fmt.Println("B: ", curr)
            curr++
            even <- 0
        }
    }()

    even <- 0
    wg.Wait()
}
