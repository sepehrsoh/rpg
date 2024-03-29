package proxy

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
)

type ReversArgs struct {
	LocalPort *[]int
	HostPort  int
	Target    string
}

func NewReversArgs(localPort *[]int, hostPort int, target string) *ReversArgs {
	return &ReversArgs{LocalPort: localPort, HostPort: hostPort, Target: target}
}

func ReverseProxy(args *ReversArgs) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGQUIT)
	signal.Notify(c, syscall.SIGKILL)
	signal.Notify(c, syscall.SIGTERM)
	signal.Notify(c, syscall.SIGINT)
	ctx, canFu := context.WithCancel(context.Background())
	go func() {
		<-c
		canFu()
	}()
	for _, port := range *args.LocalPort {
		var targetPort = args.HostPort
		if targetPort == 0 {
			targetPort = port
		}
		logrus.Infof("start server with target %v:%v \n Listen on port : %v", args.Target, targetPort, port)
		incoming, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
		if err != nil {
			logrus.Errorf("could not start server on %d: %v", args.LocalPort, err)
		}
		go func(incoming net.Listener, targetPort int) {
			for {
				select {
				case <-ctx.Done():
					logrus.Infoln("Shutting Down...")
					break
				default:
					client, err := incoming.Accept()
					if err != nil {
						log.Println("could not accept client connection", err)
						break
					}

					target, err := net.Dial("tcp", fmt.Sprintf("%v:%v", args.Target, targetPort))
					if err != nil {
						log.Println("could not connect to target", err)
						break
					}
					go func() {
						_, err = io.Copy(target, client)
						if err != nil {
							logrus.Errorln("err client to target : ", err)
						}
					}()
					go func() {
						_, err = io.Copy(client, target)
						if err != nil {
							logrus.Errorln("err target to client : ", err)
						}
					}()
				}
			}
		}(incoming, targetPort)
	}
	<-ctx.Done()

}
