# async-exec-go

**async-exec-go** is a generic function to manage asynchronous calls in go applications.

## How does it work?
- Create your function with the below format or pattern.

    `func incrementAge(age int, wg *sync.WaitGroup) int {
		defer wg.Done()
	    return age + 3
	}`

- You can call above function using **Executor** as follows:

	`async.Executor(incrementAge, age, result, wg)`