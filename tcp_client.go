package main

import (
	"github.com/buseliiik/is105sem03/mycrypt"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "172.17.0.3:8080")
	if err != nil {
		log.Fatal(err)
	}

	log.Println("os.Args[1] = ", os.Args[1])

	kryptertMelding := mycrypt.Krypter([]rune(os.Args[1]), mycrypt.ALF_SEM03, len(mycrypt.ALF_SEM03))
	log.Println("Kryptert melding: ", string(kryptertMelding))

	_, err = conn.Write([]byte(string(kryptertMelding)))
	if err != nil {
		log.Fatal(err)
	}

	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		log.Fatal(err)
	}

	dekryptertMelding := mycrypt.DeKrypter([]rune(string(buf[:n])), mycrypt.ALF_SEM03, len(mycrypt.ALF_SEM03))
	log.Println("Dekrypter melding: ", string(dekryptertMelding))

	response := string(dekryptertMelding)
	log.Printf("Reply from proxy: %s", response)

}
