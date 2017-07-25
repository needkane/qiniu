package main

import (
	"io"
	"log"
	"os"
	"sync"
	"time"
)

var Window = 50 * time.Millisecond

// Controller can limit multiple io.Reader(or io.Writer) within specific rate.
type Controller struct {
	capacity  int
	threshold int
	cond      *sync.Cond
	done      chan struct{}
}
type rateWriter struct {
	underlying io.Writer
	c          *Controller
}

func (self *Controller) run(capacity int) {
	t := time.NewTicker(Window)
	for {
		select {
		case <-t.C:
			self.cond.L.Lock()
			self.capacity = capacity
			self.cond.L.Unlock()
			self.cond.Broadcast()
		case <-self.done:
			t.Stop()
			// 认为Close一定发生在所有Read之后
			return
		}
	}

}

func (self *Controller) Close() error {
	self.done <- struct{}{}
	return nil
}

func (self *Controller) Writer(underlying io.Writer) io.Writer {
	return &rateWriter{
		underlying: underlying,
		c:          self,
	}
}
func (self *Controller) assgin(size int) int {
	self.cond.L.Lock()
	for self.capacity == 0 {
		self.cond.Wait()
	}
	if size > self.capacity {
		size = self.capacity
	}
	self.capacity -= size
	self.cond.L.Unlock()
	return size
}

func (self *Controller) fill(size int) {
	if size <= 0 {
		return
	}

	self.cond.L.Lock()
	self.capacity += size
	if self.capacity > self.threshold {
		self.capacity = self.threshold
	}
	self.cond.L.Unlock()
	self.cond.Broadcast()
}
func (self *rateWriter) Write(p []byte) (written int, err error) {
write:
	size := len(p)
	size = self.c.assgin(size)

	n, err := self.underlying.Write(p[:size])
	self.c.fill(size - n)
	written += n
	if err != nil {
		return
	}
	if size != len(p) {
		p = p[size:]
		goto write
	}
	return
}
func NewController(ratePerSecond int) *Controller {

	capacity := ratePerSecond * int(Window) / int(time.Second)
	log.Println(ratePerSecond, capacity)
	self := &Controller{
		threshold: capacity,
		capacity:  capacity,
		cond:      sync.NewCond(new(sync.Mutex)),
		done:      make(chan struct{}, 1),
	}
	go self.run(capacity)
	return self
}
func NewRateWriter(w io.Writer, ratePerSecond int) io.WriteCloser {
	c := NewController(ratePerSecond)
	w = c.Writer(w)
	return struct {
		io.Writer
		io.Closer
	}{
		Writer: w,
		Closer: c,
	}
}

func main() {
	var src, dst *os.File
	src, _ = os.Open("base64")
	defer func() {
		src.Close()
		dst.Close()
	}()

	dst, _ = os.OpenFile("base64_2", os.O_WRONLY|os.O_CREATE, 0644)

	wc := NewRateWriter(dst, 1024)
	defer wc.Close()
	begin_output := time.Now()
	_, err := io.Copy(wc, src)
	log.Println(err)
	log.Printf("use time %d", int(time.Since(begin_output)/time.Microsecond))

}
