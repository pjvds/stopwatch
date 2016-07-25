# stopwatch

Time things in Go.

## Time a function

``` go
func doWork() {
	// do something
}

elapsed := stopwatch.Time(func() {
	doWork()
})

fmt.Printf("doWork took %v", elapsed)
```

## Time a function with result

``` go
func doWork() (int, error){
	// do something
}

var result int
elapsed := stopwatch.Time(func() {
	result = doWork()
})

fmt.Printf("doWork result=%v and took %v", result, elapsed)
```
