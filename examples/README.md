## Example-1

Find all numbers between 1 and 100 that are divisible by 3 or 5, but not by 15 ?
```
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
Expected Output (Multiple of 15 are missing):
```
3
5
6
9
10
12
18
20
.....
```

## Example-2

Writing a simple Go program that uses goroutines (lightweight threads) to handle multiple tasks concurrently.
- Create a program to launch multiple tasks concurrently (e.g., downloading mock files).
- Measure and compare execution time with and without goroutines.

Execute the `main.go` file.
```bash
go run main.go
```
Sample Output
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
**Key Points:**
- Without goroutines, each task would take 2 seconds, totaling 10 seconds for 5 tasks.
- With goroutines, all tasks run simultaneously, completing in just over 2 seconds.
- This demonstrates Go's efficiency in handling concurrency using goroutines and WaitGroup.