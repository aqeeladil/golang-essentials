# Go Programming Examples

## Example 1: Finding Numbers Divisible by 3 or 5 but Not 15

This Go program finds all numbers between 1 and 100 that are divisible by 3 or 5, but not by 15.

### Code:
```go
package main

import "fmt"

func main() {
    for i := 1; i <= 100; i++ {
        if (i%3 == 0 || i%5 == 0) && i%15 != 0 {
            fmt.Println(i)
        }
    }
}
```

### Expected Output:
Numbers that are multiples of 3 or 5 but not 15:
```
3
5
6
9
10
12
18
20
...
```

### Explanation:
- The program iterates from 1 to 100.
- It prints numbers divisible by either 3 or 5.
- Numbers that are multiples of 15 (divisible by both 3 and 5) are excluded.

---

## Example 2: Using Goroutines for Concurrent Execution

This example demonstrates how to use goroutines (lightweight threads) to execute multiple tasks concurrently.

### Features:
- Simultaneously launches multiple tasks (e.g., downloading mock files).
- Measures execution time to compare sequential vs. concurrent execution.

### Execution:
Run the `main.go` file using:
```bash
go run main.go
```

### Sample Output:
```yaml
Starting tasks concurrently using goroutines...
Task 1: Downloading file...
Task 2: Downloading file...
Task 3: Downloading file...
Task 4: Downloading file...
Task 5: Downloading file...
Task 1: Download complete!
Task 2: Download complete!
Task 3: Download complete!
Task 4: Download complete!
Task 5: Download complete!
All tasks completed in 2.003456 seconds
```

### Key Points:
- **Without goroutines**, each task would take 2 seconds sequentially, totaling **10 seconds** for 5 tasks.
- **With goroutines**, all tasks run concurrently, completing in just **over 2 seconds**.
- This demonstrates **Go's efficiency in handling concurrency** using goroutines and the `sync.WaitGroup` package.

---

