# Lesson 1: Customizing a Web Server

In [lesson 0](../lesson-000-web-server/) we built a simple web server.
In this lesson, we will build on top of this web server.
We will add the ability for the user to run the web server in the command-line
and to configure it through flags and arguments.

This kind of component is something you will see done in multiple projects in
the Kubernetes community.
One great example is the
[Kubernetes API server](https://github.com/kubernetes/kubernetes/tree/master/cmd/kube-apiserver).

So let's get to it.

## Getting Started
In order to build our web server as a CLI we will use
https://github.com/spf13/cobra.

The main addition of our web server will be the definition of a cobbra command
to define flags, take in arguments, and all that.

## Testing It

Since we have a CLI, we now have useful descriptions and help messages
```
$ go build -o server && ./server --help
The server is an example application to mimick the organization of the
Kubernetes API server.

Usage:
  server [flags]

Flags:
  -a, --addr string   server's address (default "0.0.0.0:8080")
  -h, --help          help for server
  -v, --version       version for server
```

You can see we have some flags.
Let's see what version our app is
```
$ go build -o server && ./server -v
server version v1.0.0
```

And to run it
```
$ go build -o server && ./server
2020/05/17 14:42:57 version: v1.0.0
2020/05/17 14:42:57 args: []string{}
2020/05/17 14:42:57 starting server at 0.0.0.0:8080
```

And finally, if you want to specify a different listening address for the
server
```
$ go build -o server && ./server -a 0.0.0.0:9090
2020/05/17 14:43:40 version: v1.0.0
2020/05/17 14:43:40 args: []string{}
2020/05/17 14:43:40 starting server at 0.0.0.0:9090
```

You should be able to test it by going over to http://localhost:9090/.
