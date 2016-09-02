package serve

import (
	"flag"
	"os"
	"runtime"
	"io/ioutil"
	"fmt"
	"os/exec"
	"os/signal"
	"github.com/hisPeople/serve"
)

var (
	port    = 8080
	webroot = "."
	tmpdir = "."
)

func init() {
	wd, _ := os.Getwd()
	flag.IntVar(&port, "port", port, "The port (default is 8080)")
	flag.StringVar(&webroot, "webroot", wd, "Web root directory (default is current work directory)")

	// shorthand version flags
	flag.IntVar(&port, "p", port, "The port (default is 8080)")
	flag.StringVar(&webroot, "d", wd, "Web root directory (default is current work directory)")

	flag.Parse()
}

func chkdir() {
	file, err := os.Open(webroot)
	if err != nil || os.IsNotExist(err) {
		fmt.Fprintf(os.Stderr, webroot + " does not exist", err)
		os.Exit(1)
	} else {
		stat, _ := file.Stat()
		if !stat.IsDir() {
			newtmpdir, _ := ioutil.TempDir("", "serve")
			tmpdir = newtmpdir
			cmd := exec.Command("cp", webroot, tmpdir + "/")
			cmd.Run()
			webroot = tmpdir
		}
	}
}

func cleanup() {
	os.RemoveAll(tmpdir)
	os.Exit(0)
}

func main() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for sig := range c {
			fmt.Printf(sig, " cleaning up and stopping server...\n")
			cleanup()
		}
	}()

	for {
		runtime.GOMAXPROCS(runtime.NumCPU())
		chkdir()
		server := &serve.FileServer{Port: port, Webroot: webroot}
		server.Start()
	}
}
