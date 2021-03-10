package events

import (
	"flag"
	"log"
	"net"
	"net/http"
)

func GetOutboundIP() net.IP {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP
}

func StartHttpServer(subcommand string, args []string) {
	serverCommand := flag.NewFlagSet(subcommand, flag.ExitOnError)

	serverCommand.Parse(args)

	http.HandleFunc("/add", AddHandler)
	machineIp := GetOutboundIP()
	log.Printf("Starting server on port http://%s:12345", machineIp.String())
	log.Fatal(http.ListenAndServe(":12345", nil))
}
