package main

import (
	"fmt"
	"io"
	"os"
	"runtime"
	"time"
)

func main() {
	speaker(os.Stdin, os.Stdout)

	time.Sleep(5 * time.Second)
	runtime.Gosched()
}

func speaker(in io.Reader, out io.Writer) {
	fmt.Println("in speaker")
	io.Copy(out, in)


}
