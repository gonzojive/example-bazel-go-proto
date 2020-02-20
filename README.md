## Example Bazel + Go project

Run the program:

```shell
bazel run //cmd/hello:hello
```

After adding a new import for `github.com/golang/glog`, run

```shell
bazel run //:gazelle -- update-repos github.com/golang/glog
```
