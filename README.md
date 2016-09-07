# serve: serve files or folders

stolen from: https://github.com/itang/gohttp with a few additions. 
If the path you specify for serving is a file this will create a tmp folder and serve up a copy of just that single file. I find this useful for when I want to share a file with co-workers but don't want to share the other files in the directory.

It also generates a url to your machine [http://your.ip:port] and copies it your clipboard for easy pasting in your favorite chat client or whatever.

- Linux requires `xclip` or `xsel` to be installed 

### Usage

```
$ go get -u github.com/hisPeople/serve
$ go install github.com/hisPeople/serve

$ serve --help

$ serve
Serving HTTP on 192.168.1.103 port 8080 from "/home/hispeople/workspace/work" ...

$ serve -d /home -p 9000
Serving HTTP on 192.168.1.128 port 9000 from "/home" ...
```

### License

Distributed under the [Apache License Version 2.0](http://www.apache.org/licenses/LICENSE-2.0.html).
