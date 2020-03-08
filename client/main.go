package main

import (
	"log"

	"github.com/didintri196/easyssh"
	"golang.org/x/crypto/ssh"
)

func main() {
	config := &ssh.ClientConfig{
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	hostString := "localhost:9999"
	conn, err := easyssh.Dial("tcp", hostString, config)
	if err != nil {
		log.Fatalf("unable to connect: %s", err)
	}
	defer conn.Close()

	err = conn.RemoteForward("localhost:8080", "localhost:27017")
	if err != nil {
		log.Fatalf("unable to forward local port: %s", err)
	}

}
