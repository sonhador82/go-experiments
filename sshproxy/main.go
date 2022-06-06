package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"

	"golang.org/x/crypto/ssh"
)

func main() {
	fmt.Println("Hello")

	// var hostKey ssh.PublicKey

	key, err := ioutil.ReadFile("/home/demo/.ssh/id_rsa")
	if err != nil {
		log.Fatalf("unable to read private key: %v", err)
	}

	// Create the Signer for this private key.
	signer, err := ssh.ParsePrivateKey(key)
	if err != nil {
		log.Fatalf("unable to parse private key: %v", err)
	}

	config := &ssh.ClientConfig{
		User: "demo",
		Auth: []ssh.AuthMethod{
			// Use the PublicKeys method for remote authentication.
			ssh.PublicKeys(signer),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	// Connect to the remote server and perform the SSH handshake.
	client, err := ssh.Dial("tcp", "127.0.0.1:22", config)
	if err != nil {
		log.Fatalf("unable to connect: %v", err)
	}
	defer client.Close()

	conn, err := client.Dial("tcp", "ifconfig.me:80")
	if err != nil {
		log.Fatalf("unable to connect to rabbitmq: %v", err)
	}

	fmt.Fprintf(conn, "GET / HTTP/1.1\r\nHost: ifconfig.me\r\n\r\n")

	fmt.Println("here")
	resp, err := bufio.NewReader(conn).ReadString('\n')
	fmt.Println("resp: ", resp)

}
