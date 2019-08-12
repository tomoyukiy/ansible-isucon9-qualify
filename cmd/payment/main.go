package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/isucon/isucon9-qualify/bench/server"
)

func main() {
	liPayment, err := net.ListenTCP("tcp", &net.TCPAddr{Port: 5555})
	if err != nil {
		log.Fatal(err)
	}

	pay := server.NewPayment()

	serverPayment := &http.Server{
		Handler: pay,
	}

	pay.SetDelay(200 * time.Millisecond)

	fmt.Fprintln(os.Stderr, serverPayment.Serve(liPayment))
}
