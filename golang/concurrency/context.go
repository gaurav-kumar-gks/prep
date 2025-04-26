package concurrency

/*

Contexts:
- Context is a way to pass immutable request-scoped values, cancellation signals, and deadlines across API boundaries.
- It is used to manage the lifecycle of requests in concurrent programming.
- Context is not meant to be used for passing optional parameters or configuration values.
- Context is typically passed as the first argument to functions and methods.


Various context types:

1. Background: The top-level context, typically used for main functions or tests.
ctx := context.Background()

2. TODO: A context that indicates that the caller doesn't care about the context.
ctx := context.TODO()

3. WithCancel: A context that can be cancelled.
ctx, cancel := context.WithCancel(parentCtx)
go func() {
	<- ctx.Done()
	fmt.Println("Context cancelled")
	// Perform cleanup here
}
cancel()

4. WithTimeout: A context that will be automatically cancelled after a specified duration.
ctx, cancel := context.WithTimeout(parentCtx, timeout)
defer cancel()

5. WithDeadline: A context that will be automatically cancelled at a specified time.
deadline := time.Now().Add(100 * time.Millisecond)
ctx, cancel := context.WithDeadline(parentCtx, deadline)
defer cancel()

6. WithValue: A context that carries a value.
valueCtx := context.WithValue(parentCtx, key, value)
if val := valueCtx.Value("key"); val != nil {
	fmt.Printf("Context value: %v\n", val)
}

*/
