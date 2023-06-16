package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
)

var protocol = flag.String("protocol", "tcp", "ECHO server use tcp/udp/all.")
var listenHost = flag.String("host", "0.0.0.0", "ECHO server listen on this address.")
var listenPort = flag.Int("port", 7890, "ECHO server listen on this port.")
var bufferSize = flag.Uint("buf-size", 1024, "ECHO server receive buffer size.")

func handleTCPConn(conn net.Conn) {
	protocol := "tcp"
	r := bufio.NewReader(conn)
	w := bufio.NewWriter(conn)
	defer conn.Close()
	for {
		var data = make([]byte, *bufferSize)
		n, err := r.Read(data)
		if err == io.EOF || err != nil {
			break
		}
		log.Printf("ECHO server receive new data <%s> %s<->%s, length: %d", protocol, conn.LocalAddr(), conn.RemoteAddr(), n)
		w.Write(data[:n])
		w.Flush()
	}
	log.Printf("ECHO server close old connection <%s> %s<->%s", protocol, conn.LocalAddr(), conn.RemoteAddr())
}

func handleUDPConn(conn *net.UDPConn, data []byte, addr *net.UDPAddr) {
	protocol := "udp"
	log.Printf("ECHO server receive new data <%s> %s<->%s, length: %d", protocol, conn.LocalAddr(), addr, len(data))
	conn.WriteTo(data, addr)
}

func serveTCP(addr string) {
	protocol := "tcp"
	log.Printf("ECHO server listen on <%s> %s", protocol, addr)
	sock, err := net.Listen(protocol, addr)
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := sock.Accept()
		log.Printf("ECHO server accept new connection <%s> %s<->%s", protocol, conn.LocalAddr(), conn.RemoteAddr())
		if err != nil {
			log.Print(err)
		}
		go handleTCPConn(conn)
	}
}

func serveUDP(addr string) {
	protocol := "udp"
	log.Printf("ECHO server listen on <%s> %s", protocol, addr)
	udpAddr, err := net.ResolveUDPAddr(protocol, addr)
	if err != nil {
		log.Fatal(err)
	}
	sock, err := net.ListenUDP(protocol, udpAddr)
	if err != nil {
		log.Fatal(err)
	}
	for {
		var data = make([]byte, *bufferSize)
		n, addr, err := sock.ReadFromUDP(data)
		if err == io.EOF || err != nil {
			log.Print(err)
		}
		go handleUDPConn(sock, data[:n], addr)
	}
}

func main() {
	flag.Parse()
	addr := fmt.Sprintf("%s:%d", *listenHost, *listenPort)
	switch *protocol {
	case "tcp":
		serveTCP(addr)
	case "udp":
		serveUDP(addr)
	case "all":
		go serveTCP(addr)
		serveUDP(addr)
	default:
		log.Fatalf("Error protocol %s, please use tcp/udp/all", *protocol)
	}
}
