package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
	"os"
	"os/exec"
	"runtime"
	"time"

	"github.com/kavu/go_reuseport"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	listener, err := reuseport.NewReusablePortListener("tcp4", "localhost:8881")
	if err != nil {
		panic(err)

	}
	defer listener.Close()

	server := &http.Server{}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(os.Getgid())
		go func() {
			cmd := exec.Command("./server")
			cmd.Start()
			log.Println("-------start")
			time.Sleep(18 * time.Second)
			log.Println("-------end")
			cmd.Process.Kill()
			cmd.Wait()
		}()
		fmt.Fprintf(w, "Hello, %q\n", html.EscapeString(r.URL.Path))
	})

	panic(server.Serve(listener))

}
