# Example Bazel + Go + Protocol Buffers project


## Quickstart

Run the program:

```shell
bazel run //cmd/hello:hello
```

After adding a new import for `github.com/golang/glog`, run

```shell
bazel run //:gazelle -- update-repos github.com/golang/glog
```

After adding a new `.go` file:

```shell
bazel run //:gazelle
```

## Quickstart with protocol buffers

```shell
bazel run //cmd/hello_protobufs -- --name "Helloooooooo" --suffix ". :)"
```

Outputs

```
Helloooooooo, world. :)
```
