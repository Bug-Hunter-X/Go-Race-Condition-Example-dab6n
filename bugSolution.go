```go
package main

import (

        "fmt"
        "sync"
)

func main() {
        var wg sync.WaitGroup
        var count int
        var mu sync.Mutex
        const numGoroutines = 1000

        for i := 0; i < numGoroutines; i++ {
                wg.Add(1)
                go func(i int) {
                        defer wg.Done()
                        mu.Lock()
                        count++
                        mu.Unlock()
                }(i)
        }

        wg.Wait()
        fmt.Println("Count:", count)
}
```