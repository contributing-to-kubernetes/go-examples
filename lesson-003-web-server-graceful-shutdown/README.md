# Lesson 3: Web server graceful shutdown

Taking as a foundation the webserver created in [lesson 0](../lesson-000-web-server/), we will be adding `graceful shutdown` functionality to it.

In order to achieve this, we will be using `channels` and `goroutines`

## Knowledge bits

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

In our case we want to use channels along with `goroutines`.

### Goroutines

A `goroutine` is a lightweight thread of execution managed by the Go runtime; it is the way we have to run a piece of code concurrently with the original calling code, as we can see explained in the [oficial documentation.](https://tour.golang.org/concurrency/1)

A small example could be:

```
func printMyString() {
    fmt.Println("my string")
}

func main() {
    go printMyString()
    time.Sleep(100 * time.Millisecond)
    fmt.Println("printMyString function")
}
```

In line 6 is where we start a new `goroutine`, calling our already defined `printMyString` function, while the `main` function runs in its own `goroutine` (called the _main goroutine_)


**_But how are we going to use both features together?_**


We need to point something about channels: they are **blocking by default**.

We can use this in our benefit and block our server from being closed until all pending requests have been served when we receive a specific signal.

## Lesson content

After getting a quick intro to `channels` and `goroutines`, we will dive into our lesson.

As we initially stated, the current implementation is based on our server from lesson0.

To achieve _graceful shutdown_ we added the following main changes:

- created a `sigint` channel; we will use it to notify our goroutine we have received a signal: `os.Interrupt` or `syscall.SIGTERM` in this case
- with `server.SetKeepAlivesEnabled(false)` we set the server to not keep alive any connection (which in fact is the desired effect of having a gracefull shutdown behavior)
- create the `done` channel; this one will be used to let the main goroutine we have finished the graceful shutdown.

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
