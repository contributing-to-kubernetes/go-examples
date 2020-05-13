# Lesson 0: Building a Web Server

First step in our jourmey: we will build a web server.
Why?
Because it will teach us how to use a lot of basic building block for more
complicated programs later on.

Our example program lives here in [`main.go`](./main.go).

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
