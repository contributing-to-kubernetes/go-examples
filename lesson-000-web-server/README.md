# Lesson 0: Building a Web Server

First step in our jourmey: we will build a web server.
Why?
Because it will teach us how to use a lot of basic building block for more
complicated programs later on.

Our example program lives here in [`main.go`](./main.go).

But before we dive into code, we will present some concepts we will be working on in this lesson.

- Go functions can return more than one value. This is a quite common pattern and we will be using like in the following example.

    `host, err := os.Hostname()`

    Here we are expecting to assign the output of `os.Hostname()` into `host` variable.
    If something goes wrong, we would receive the error into `err` variable.

- pointers: no, don't worry about the naming (you can not do pointer arithmetic as in C or C++). Pointers in Go are used to pass variables by reference to functions. This means, we will be using the exact same object that was passed.

- when using `Sprintf`, `Fprintf` or `Printf`,  we can use several modifiers to present the data. In this lesson we will use:
  - `%v`: we will output the value in a default format
  - `%s`: the output would be the `string` format

The `go.mod` file indicates that in this directory we have a Go module.
We got that one from doing a `go mod init`, nothing fancy.

To better understand the mechanics of our example please go though the
documentation at https://godoc.org/net/http.

## Testing It

To run the app you can compile
```
$ go build -o app
```
and build it
```
$ ./app
starting web server at 0.0.0.0:8080
```

If you read through the code, you'll notice that we registered an HTTP handler
(a function that processes incoming requests) for the route `/`.
You can see this by sending a request to it
```
$ curl http://localhost:8080/
Greeting from alejandrox1-machine1!
```

Incidentally, since we only have one route registered, and that is the root,
all requests will be handled by it
```
$ curl http://localhost:8080/hello/
Greeting from alejandrox1-machine1!
```

And another one
```
$ curl http://localhost:8080/hello/again/
Greeting from alejandrox1-machine1!
```

### Next steps

As shown, we currently have one (root) route, which means: it works but not very practical.

We would like to extend the current lesson with:
- add more routes to the existing server
- be able to read URL params (e.g http://localhost:8080/hello?user=alejandrox&lesson=0) and output
  ```
  Greeting from alejandrox1-machine1!
  Params: 
  user: alejandrox
  lesson: 0
  ```
