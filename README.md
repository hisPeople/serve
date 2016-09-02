# serve: serve files or folders

stolen from: https://github.com/itang/gohttp with a few additions. If the path you specify for serving is a file this will create a tmp folder and serve up a copy of just that single file. I find this useful for when I want to share a file with coworkers but don't want to share the other files in the directory.

### Usage

```
$ go get -u github.com/hisPeople/serve
$ go install github.com/hisPeople/serve

$ serve --help

$ serve
Serving HTTP on 192.168.1.103 port 8080 from "/home/hispeople/workspace/work" ...

$ serve -d=/home -p=9000
Serving HTTP on 192.168.1.128 port 9000 from "/home" ...
```

### License

Distributed under the [Apache License Version 2.0](http://www.apache.org/licenses/LICENSE-2.0.html).
