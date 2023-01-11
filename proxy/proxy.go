package proxy

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"os/signal"
)

type ReversArgs struct {
	LocalPort int
	HostPort  int
	Target    string
}

func NewReversArgs(localPort int, hostPort int, target string) *ReversArgs {
	return &ReversArgs{LocalPort: localPort, HostPort: hostPort, Target: target}
}

func ReverseProxy(args *ReversArgs) {
	incoming, err := net.Listen("tcp", fmt.Sprintf(":%d", args.LocalPort))
	if err != nil {
		log.Fatalf("could not start server on %d: %v", args.LocalPort, err)
	}
	fmt.Printf("server running on %d\n", args.LocalPort)
	signals := make(chan os.Signal, 1)
	stop := make(chan bool)
	signal.Notify(signals, os.Interrupt)
	signal.Notify(signals, os.Kill)
	go func() {
		for _ = range signals {
			fmt.Println("\nStopping...")
			stop <- true
		}
	}()

	client, err := incoming.Accept()
	if err != nil {
		log.Fatal("could not accept client connection", err)
	}
	defer client.Close()
	fmt.Printf("client '%v' connected!\n", client.RemoteAddr())

	target, err := net.Dial("tcp", fmt.Sprintf("%v:%v", args.Target, args.HostPort))
	if err != nil {
		log.Fatal("could not connect to target", err)
	}
	defer target.Close()
	fmt.Printf("connection to server %v established!\n", target.RemoteAddr())
	go func() { io.Copy(target, client) }()
	go func() { io.Copy(client, target) }()
	<-stop
}

func init() {

}
