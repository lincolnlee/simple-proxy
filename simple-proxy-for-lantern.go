package main

import (
	"flag"
	"fmt"
	"github.com/lincolnlee/simple-proxy/sp"
	"io"
	"log"
	"net"
	"os"
	"strconv"
)

func handleClientRequest(client net.Conn) {
	if client == nil {
		log.Println("client == nil")
		return
	}
	defer client.Close()

	log.Println("ll:net.Dial:", config.Proxy+":"+strconv.Itoa(config.ProxyPort))
	//获得了请求的host和port，就开始拨号吧
	server, err := net.Dial("tcp", config.Proxy+":"+strconv.Itoa(config.ProxyPort))
	if err != nil {
		log.Println("ll:net.Dial:err:", err)
		return
	}

	//进行转发
	go io.Copy(server, client)
	io.Copy(client, server)
}

var configFile string
var config *simple_proxy.Config

func main() {
	log.Println("Reading config...")

	var cmdConfig simple_proxy.Config

	flag.StringVar(&configFile, "c", "config.json", "specify config file")
	flag.StringVar(&cmdConfig.Server, "s", "0.0.0.0", "server ip")
	flag.IntVar(&cmdConfig.ServerPort, "p", 10086, "server port")
	flag.StringVar(&cmdConfig.Client, "b", "", "client ip")
	flag.StringVar(&cmdConfig.Password, "k", "", "password")
	flag.StringVar(&cmdConfig.Proxy, "l", "127.0.0.1", "lantern ip")
	flag.IntVar(&cmdConfig.ProxyPort, "x", 51769, "lantern port")
	flag.Parse()

	var err error
	config, err = simple_proxy.ParseConfig(configFile)
	if err != nil {
		if !os.IsNotExist(err) {
			fmt.Fprintf(os.Stderr, "error reading %s: %v\n", configFile, err)
			os.Exit(1)
		}
		config = &cmdConfig
		simple_proxy.UpdateConfig(config, config)
	} else {
		simple_proxy.UpdateConfig(config, &cmdConfig)
	}

	log.Println("net.Listen:", config.Server, ":", config.ServerPort)
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	l, err := net.Listen("tcp", config.Server+":"+strconv.Itoa(config.ServerPort))
	if err != nil {
		log.Panic(err)
	}

	for {
		client, err := l.Accept()
		if err != nil {
			log.Panic(err)
		}

		log.Println("ll:handleClientRequest:", client.RemoteAddr())
		go handleClientRequest(client)
	}
}
