# Lesson 3: Web server graceful shutdown

Taking as a foundation the webserver created in [lesson 000](../lesson-000-web-server/), we will be adding `graceful shutdown` functionality to it.

In order to achieve this, we will be using `channels` and `goroutines`

## Knowledge bits


### Goroutines

A `goroutine` is a lightweight thread of execution managed by the Go runtime; it is the way we have to run a piece of code concurrently with the original calling code, as we can see explained in the [oficial documentation.](https://tour.golang.org/concurrency/1)

A small example could be:

```
func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	go say("k8s")
	say("contribute")
}
```

Calling `go say()` is where we start a new `goroutine`, calling our already defined `say` function, while the `main` function runs in its own `goroutine` (called the _main goroutine_)


### Channels

Channels are a concurrency synchronization technique we can use in Go.

We can create a channel using the keyword `chan`, and we can transport data of only one data type.

From the [official documention](https://tour.golang.org/concurrency/2) we can see how data can be sent and received from the channel.

```
ch := make(chan int)

ch <- v    // Send v to channel ch.
v := <-ch  // Receive from ch, and
           // assign value to v.
```

**_Unbuffered_ vs _buffered_ channels**

Channels are created **blocking by default** (or _unbuffered_)

If we try to send a resource to an unbuffered channel, the channel will lock (preventing us from sending anything else into the channel) until someone else reads from it.
This essentially means that an unbuffered channel is restricted to never having more than 1 element inside of it.
And viceversa.

On the other hand, buffered channels have capacity and they are able to keep a number of resources.
The only times buffered channels will lock goroutines are when a sender tries to send a resource and the channel is full or when a goroutine tries to get a resource and the channel is empty.

We can use this in our benefit and block our server from being closed until all pending requests have been served when we receive a specific signal.


### Signals

Signals are software interrupts sent to a program to indicate that an event has happened.

In this example we will take care of two specific signals: `os.Interrupt` and `syscall.SIGTERM`:

- `os.Interrupt` this is tipically the signal sent when we type Control-C, which normally causes the program to exit.
- `syscall.SIGTERM` is usually sent when you want to give the process an opportunity to clean up before termination.


## Lesson content

After getting a quick intro to `channels` and `goroutines`, we will dive into our lesson.

As we initially stated, the current implementation is based on our server from [lesson 000](../lesson-000-web-server/).


To achieve _graceful shutdown_ we added the following changes:

- created a `sigint` channel; we will use it to notify our goroutine we have received a signal: `os.Interrupt` or `syscall.SIGTERM` in this case
- with `server.SetKeepAlivesEnabled(false)` we set the server to not keep alive any connection (which in fact is the desired effect of having a gracefull shutdown behavior) to allow our server to finish processing any requests already received before the app received a termination signal.
The first step is to disable "keep alive" TCP connections [https://godoc.org/net/http#Server.SetKeepAlivesEnabled](https://godoc.org/net/http#Server.SetKeepAlivesEnabled) before proceeding with the graceful shutdown of the server [https://godoc.org/net/http#Server.Shutdown](https://godoc.org/net/http#Server.Shutdown)
- create the `done` channel; this one will be used to let the main goroutine we have finished the graceful shutdown.
Adding this in the `func main()` allows us to run code only after the app is shut down - if needed


## Testing It

In this case we would need to run our server in one terminal and have an additional one, which we will use to `curl` our server.

Terminal 1 will have our server running after compiling our source code.

In Terminal 2 we will execute the following line:

```
for (( ; ; )); do curl http://localhost:8080;done
```

This is an infinite **for** loop in `bash`, which will be requesting our server using `curl`. Everytime it hits our server, it will print our `Greeting from _x y z_` message.

To see our example working as we expect (having a graceful shutdown) we would need to hit `ctrl + c` in Terminal 1 where our server is running.

_Hint: we would also need to `ctrl + c` our **for loop** in Terminal 2 to prevent our script to continue sending requests to our stopped server_

We could also open a new terminal and first find the PID of our process e.g. `pgrep -i main`. This command will output the our main process ID. We can then call `kill -SIGTERM _processid_` to see how our server behaves when it receives the other signal we defined in our server.
