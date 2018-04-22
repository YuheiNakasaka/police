# Police

Police limits the count of goroutine processed at same time.

# Usage

Read this example.

```
package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/YuheiNakasaka/police"
)

func main() {
	// initialize
	police := &police.Arrival{}
	// set the count of goroutine processed at the same time
	police.Limit(3)

	wg := &sync.WaitGroup{}
	arr := []int{1, 1, 1, 1, 1, 1}
	for i := 0; i < len(arr); i++ {
		wg.Add(1)
		go func(j int) {
			// start block
			police.Block()

			defer wg.Done()
			time.Sleep(time.Duration(arr[j]) * time.Second)
			fmt.Printf("(arr[%d] = %d) is done.\n", j, arr[j])

			// finish block
			police.Release()
		}(i)
	}
	wg.Wait()
}
```

Above code is the same at here.

```
package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	ch := make(chan struct{}, 3)

	wg := &sync.WaitGroup{}
	arr := []int{1, 1, 1, 1, 1, 1}
	for i := 0; i < len(arr); i++ {
		wg.Add(1)
		go func(j int) {
			ch <- struct{}{}

			defer wg.Done()
			time.Sleep(time.Duration(arr[j]) * time.Second)
			fmt.Printf("(arr[%d] = %d) is done.\n", j, arr[j])

			<-ch
		}(i)
	}
	wg.Wait()
}
```

# License
MIT

# Author
Yuhei Nakasaka
