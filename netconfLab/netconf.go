package main

import (
	"fmt"
	"github.com/gliderlabs/ssh"
	"io"
	"log"
)

func main() {
	ssh.Handle(func(s ssh.Session) {
		io.WriteString(s, fmt.Sprintf("Hello %s\n", s.User()))
	})

	log.Println("starting ssh server on port 2222...")
	log.Fatal(ssh.ListenAndServe(":2222", nil))
}
