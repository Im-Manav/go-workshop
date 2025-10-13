# Channels

Once you've started a goroutine, you might want to communicate with it - to tell it to stop, to pass it data, or to get results back.

One way to stop a goroutine is to use the `context` package. You can create a `context.Context` with a timeout or a cancellation function, and pass it to the goroutine. The goroutine can then check the context to see if it should stop work.

```go
package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var wg sync.WaitGroup
	wg.Go(func() {
		i := 1
		for {
			fmt.Printf("Count: %d\n", i)
			time.Sleep(1 * time.Second)
			i++

			if ctx.Err() != nil {
				fmt.Println("Goroutine exiting due to context cancellation")
				return
			}
		}
	})

	wg.Wait()

	fmt.Println("Main function completed")
}
```

That's quite interesting, but it has a limitation - we can't pass data to the goroutine, or get results back from it.

Also, the `time.Sleep` operation won't be interrupted. If the goroutine is sleeping when the context is cancelled, it will only check the context after it wakes up. Maybe that's OK if the `time.Sleep` is a second, but what if it's a minute, or an hour? We wouldn't want to wait that long for the goroutine to stop.

This is where channels come in.

## What are channels?

Go channels are a little bit like an in-memory message queue. They allow you to send and receive values between goroutines. They're safe for concurrent use, so multiple goroutines can send and receive values on the same channel without additional synchronisation (i.e., you don't need mutexes).

They are type safe - a channel can only carry values of a specific type. You create a channel using the `make` function.

In this case, we're creating a channel that can carry `int` values:

```go
ch := make(chan int)
```

### Pushing data

You can push data into a channel using the `<-` operator, and you can read data from a channel using the same operator.

The position of the `<-` operator indicates the direction of data flow. If the channel is on the left of the operator, data is being sent into the channel:

```go
ch <- value
```

If the channel is full, the sending goroutine will block until there is space in the channel.

You can make a buffered channel by passing a second argument to `make`. This specifies the size of the buffer. A buffered channel allows you to send multiple values into the channel without blocking, up to the size of the buffer.

```go
ch := make(chan int, 5)
```

That is really useful if you want to populate a channel with some initial data, or if you want to make sure that there's always something in the channel for workers to process.

It's a rule in Go that you should always close a channel when you're done sending data into it. This is done using the `close` function:

```go
close(ch)
```

You should only close a channel from the sending side, and never from the receiving side. Closing a channel indicates that no more values will be sent on it.

Pushing data into a closed channel will cause a panic, which you don't want.

### Pulling data

If the `<-` is to the left of the channel, data is being received from the channel:

```go
value := <-ch
```

If the channel is empty, the receiving goroutine will block until there is data in the channel.

When reading from a channel, you can also check if the channel has been closed. This is done by using a second return value from the receive operation:

```go
value, ok := <-ch
```

But it's more common to use a `for range` loop to read from a channel until it is closed:

```go
for value := range ch {
    // process value
}
```

### Selecting on channels

Go also provides a `select` statement that allows a goroutine to wait on multiple communication operations. A `select` blocks until one of its cases can run, then it executes that case. If multiple cases are ready, one of them is chosen at random.

Making use of `select`, we can improve our previous example to use a channel that will be notified when the context is cancelled. This allows us to avoid using `time.Sleep`, and instead use `time.After` in a `select` statement, which can be interrupted by the context cancellation.

```go
package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var wg sync.WaitGroup
	wg.Go(func() {
		i := 1
		for {
			fmt.Printf("Count: %d\n", i)

			select {
			case <-ctx.Done():
				fmt.Println("Goroutine exiting due to context cancellation")
				return
			case <-time.After(1 * time.Second):
				i++
			}
		}
	})

	wg.Wait()

	fmt.Println("Main function completed")
}
```

This is also useful if a worker can process data from multiple channels. The `select` statement allows the worker to wait for data from any of the channels, and process it as soon as it arrives, or, if the context is cancelled, or a timeout occurs, to stop work.

## Patterns

Goroutines and channels can be combined to create a variety of concurrency patterns. Here are some common ones:

- Fan-out - multiple workers read from a single input channel
- Fan-in - a single worker reads from multiple input channels
- Worker pools - a fixed number of workers read from a single input channel and write to a single output channel
- Pipeline - a series of stages where each stage is a set of workers that read from an input channel and write to an output channel
- Orchestration - a single goroutine that coordinates multiple workers and channels
- Cancellation - using a done channel to signal goroutines to stop work
- Rate limiting - prevent too many goroutines from running at the same time, e.g. using a semaphore pattern with a buffered channel

## Example: Worker pool

It's relatively easy to create a worker pool using goroutines and channels. A worker pool is a fixed number of workers that read from a single input channel and write to a single output channel.

Start by making a channel for inputs, and one for outputs:

```go
inputs := make(chan int)
outputs := make(chan int)
```

Next, start a fixed number of worker goroutines. Each worker will read from the input channel, process the data, and write the result to the output channel.

```go
var wg sync.WaitGroup
for range 4 {
  wg.Go(func() {
    for input := range inputs {
      // Process the input (in this case, just square it).
      outputs <- input * input
    }
  })
}

// Close the outputs channel when all of the workers are done.
go func() {
  wg.Wait()
  close(outputs)
}()
```

Now, we can push some data into the input channel. In this case, we'll just push the numbers 1 to 10.

```go
go func() {
  for i := range 10 {
    inputs <- i
  }
  close(inputs)
}()
```

Finally, we can read the results from the output channel. We'll use a `for range` loop to read from the channel until it is closed.

```go
for output := range outputs {
  fmt.Println(output)
}
```

And we should see the squares of the numbers 1 to 10 printed to the console, however, not necessarily in order, because the workers are processing the inputs concurrently.

Putting it all together, we have:

```go
package main

import (
	"fmt"
	"sync"
)

func main() {
	inputs := make(chan int)
	outputs := make(chan int)

	// Start 4 workers.
	var wg sync.WaitGroup
	for range 4 {
		wg.Go(func() {
			for input := range inputs {
				// Process the input (in this case, just square it).
				outputs <- input * input
			}
		})
	}

	// Close the outputs channel when all of the workers are done.
	go func() {
		wg.Wait()
		close(outputs)
	}()

	// Push some data into the inputs channel.
	go func() {
		for i := range 10 {
			inputs <- i
		}
		close(inputs)
	}()

	// Read the results from the outputs channel.
	for output := range outputs {
		fmt.Println(output)
	}
}
```

You can try this out in your browser at https://go.dev/play/p/qOYxtz1MQec

It's not very useful to just square numbers, but you can imagine replacing the squaring operation with something more useful, like making an HTTP request, processing a file, performing a database query, carrying out image recognition, or any other CPU or IO bound operation.

## Task

Try converting the main.go example to use a worker pool to calculate primes.

If you graph the performance of calculating the primes against time, you'll find that with small numbers, the overhead of managing the worker pool outweighs the benefits of concurrency, but as the numbers get larger, the performance improves significantly.

This is why it's important to benchmark your code, and to understand the performance characteristics of your application. We'll touch on benchmarking in level 200.
