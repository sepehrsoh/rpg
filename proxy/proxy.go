package proxy

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
)

var maxConn int = 0

type ReversArgs struct {
	LocalPort int
	HostPort  int
	Target    string
}

func NewReversArgs(localPort int, hostPort int, target string) *ReversArgs {
	return &ReversArgs{LocalPort: localPort, HostPort: hostPort, Target: target}
}

func ReverseProxy(args *ReversArgs) {
	logrus.Infof("start server with target %v:%v \n Listen on port : %v", args.Target, args.HostPort, args.LocalPort)
	incoming, err := net.Listen("tcp", fmt.Sprintf(":%d", args.LocalPort))
	if err != nil {
		log.Fatalf("could not start server on %d: %v", args.LocalPort, err)
	}
	fmt.Printf("server running on %d\n", args.LocalPort)
	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL)
	for {
		if maxConn < 10 {
			maxConn += 1
			defer func() {
				maxConn -= 1
			}()
			go func() {
				select {
				case <-c:
					logrus.Infoln("Shutting Down...")
					return
				default:
					client, err := incoming.Accept()
					if err != nil {
						log.Fatal("could not accept client connection", err)
					}
					defer client.Close()

					target, err := net.Dial("tcp", fmt.Sprintf("%v:%v", args.Target, args.HostPort))
					if err != nil {
						log.Fatal("could not connect to target", err)
					}
					defer target.Close()
					go func() { io.Copy(target, client) }()
					go func() { io.Copy(client, target) }()

					<-c
				}
			}()
		}

	}

}
