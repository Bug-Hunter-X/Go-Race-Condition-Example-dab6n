```go
package main

import (

        "fmt"
        "sync"
)

func main() {
        var wg sync.WaitGroup
        var count int
        const numGoroutines = 1000

        for i := 0; i < numGoroutines; i++ {
                wg.Add(1)
                go func(i int) {
                        defer wg.Done()
                        count++
                }(i)
        }

        wg.Wait()
        fmt.Println("Count:", count)
}
```