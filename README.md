# gopherpc

`gopherpc` is a tool for automatically generating GopherJS code for calling
[RPC over HTTP](http://www.gorillatoolkit.org/pkg/rpc) services.  They are
defined very similarly to regular [net/rpc](https://golang.org/pkg/net/rpc/)
services, except they take an additional parameter, which is the
`*http.Request`, and therefore more useful for web applications.

## How it Works

First you need to define your RPC services and annotate them with a special
comment. In your web server package, you can do that like this:

```go
// rpc:gen
type TestService struct {}

func (s TestService) Ping(r *http.Request, _ *struct{}, reply *string) error {
	*reply = "pong"
	return nil
}
```

Given this implementation, you can generate GopherJS code to call it by using
the `gopherpc` command, which looks something like this:

```bash
$ go install github.com/dradtke/gopherpc/cmd/gopherpc
$ gopherpc -scan <pkg> -o <output>
```

In this example, `<pkg>` is the fully-qualified import path to the package
containing your RPC definitions, and `<output>` is the path that you want the
result to be written to. The package name of the output will be inferred from
the name of the directory it appears in, but that can be overridden with the
`-pkg` flag.

The output should be written to its own package, as it defines a new `Client`
type that can be used from other GopherJS code. That package can then be imported
by other GopherJS code to provide a fully type-safe handle to your RPC services.

See the `testdata` folder for more examples.
