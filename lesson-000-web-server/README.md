# Lesson 0: Building a Web Server

First step in our jourmey: we will build a web server.
Why?
Because it will teach us how to use a lot of basic building block for more complicated programs later on.

But before we dive into code, we will present some concepts we will be working on in this lesson.

## Knowledge bits

- In Go, we organize our code into modules. Go modules are in turn made up of packages.

  In our case we will only define the `main` package. This specific package name is used by the Go compiler to know that the package should compile as an executable program.

    The `main` **function** in the **package** `main` will be the entry point of our executable.

    More info on this topic can be found in this in https://blog.golang.org/using-go-modules.

- Functions in Go can return more than one value. It is a common pattern in Go to always return an error on operations that may fail. One should then check if the error variable is `nil`, which means there was no error, or to handle the error if not `nil`.

  For example:

    ```
    host, err := os.Hostname()
    if err != nil {
        fmt.Sprintf("we saw an error: %v\n", err)
    }

    fmt.Println("everything is awesome")
    ```

- pointers: no, don't worry about the naming (you can not do pointer arithmetic as in C or C++). Pointers in Go are used to pass variables by reference to functions. This means, we will be using the exact same object that was passed.

- when using `Sprintf`, `Fprintf` or `Printf` (or any printing function that ends with an `f`)  we can use several **printing verbs** to present the data. (official docs in https://golang.org/pkg/fmt/) In this lesson we will use:
  - `%v`: we will output the value in a default format
  - `%s`: the output would be the `string` format

## Lesson content

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

### Next steps (optional)

Couple interesting things we can do:
- add more routes to the existing server
- be able to read URL params (e.g http://localhost:8080/hello?user=alejandrox&lesson=0) and output
  ```
  Greeting from alejandrox1-machine1!
  Params:
  user: alejandrox
  lesson: 0
  ```
