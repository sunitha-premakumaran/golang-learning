## Golang
- A fast, high performance, open-source complied programming language
- Static typed
- Portability and Multi-Platform Nature
    - dependencies complies to binary

### Why Golang?
- Memory Safety 
- Structural Typing
- [CSP-Style](https://golang.org/doc/effective_go.html#concurrency) concurrency

### Terminologies
- Goroutines
    - Functions that is capable of executing concurrently with other functions
    - Use the keyword `go` followed by function invocation to create gorountine

- Channels
    - the way in which two gorountines communicate to synchronize their execution and share data

- Select Statement
    - The `select` statement lets a goroutine wait on multiple communication operations.
    - A select blocks until one of its cases can run, then it executes that case. It chooses one at random if multiple are ready.

### Solutions
- [Quiz based on csv file](./gophercises-quiz/main.go)

