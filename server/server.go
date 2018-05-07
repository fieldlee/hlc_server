// center
package server

import (
//	"net/http"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"io/ioutil"
	"encoding/json"
	"log"
)

type Host struct {
	Name	string	`json:"name"`
	Domain_port	string	`json:"domain_port"`
}

type Config struct {
	Hosts	[]Host	`json:"hosts"`
	Global	Host	`json:global"`
}

type Remote struct {}

var domain = "127.0.0.1"
var port = "10000"
var global *rpc.Client

var config Config

func ConnectGlobalServer(domain_port string) *rpc.Client {
	global, err := jsonrpc.Dial("tcp", domain_port)
	if err != nil {
		log.Fatalln(err.Error())
	}

	return global
}

func Run(args []string) {
	for i := 1; i < len(args); i++ {
		switch args[i] {
		case "-D":
			if i == len(args) {
				log.Fatalln("invalid domain")
			}
			domain = args[i + 1]
		case "-p":
			if i == len(args) {
				log.Fatalln("invalid port")
			}
			port = args[i + 1]
		}
	}

	data, err := ioutil.ReadFile("/etc/hlc/hlc-server.conf.json")
	if err != nil {
		log.Fatalln(err.Error(), "cannot find the file: lcserver.conf.json")
	}

	err = json.Unmarshal(data, &config)
	if err != nil {
		log.Fatalln(err.Error(), "cannot parse the file: lcserver.conf.json")
	}

	global = ConnectGlobalServer(config.Global.Domain_port)
	defer global.Close()

	log.Println("listening @ " + domain + ":" + port)

	listen, err := net.Listen("tcp", ":" + port)
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer listen.Close()

	host := rpc.NewServer()
	err = host.Register(Remote{})
	if err != nil {
		log.Println(err.Error())
	}

	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Fatalln(err.Error())
			continue
		}
		go host.ServeCodec(jsonrpc.NewServerCodec(conn))
	}
}
