# async-exec-go

**async-exec-go** is a generic function to manage asynchronous calls in go applications.

### Install
```
go get -u github.com/kuckchuck96/async-exec-go
```

### How does it work?
- Create your function with the below format or pattern.

```
func incrementAge(age int, wg *sync.WaitGroup) int {
	defer wg.Done()
	return age + 3
}
```
> Please note you function should only accept 2 input params i.e. param1 of x type and param2 of *sync.WaitGroup type.

- You can call above function using **Executor** as follows:

```
async.Executor(incrementAge, age, result, wg)
```
> **incrementAge** is a function, **age** is input param, **result** is output channel, and **wg** is *sync.WaitGroup.
