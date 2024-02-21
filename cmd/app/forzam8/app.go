package forzam8

import (
	"bufio"
	"encoding/base64"
	"log"
	"net"
	"os"
	"strings"
	"time"
)

func Run(address string, frequency int) {
	var err error
	addr, err := net.ResolveUDPAddr("udp", address)
	if err != nil {
		log.Println(err)
	}

	connection, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		log.Println(err)
	}
	defer connection.Close()

	filePath := "cmd/app/forzam8/forzamotorsport.udp.log"
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	content, err := os.ReadFile(filePath)
	lineCount := strings.Count(string(content), "\n")

	scanner := bufio.NewScanner(file)
	i := 0
	for scanner.Scan() {
		i++
		log.Printf("Sending %d/%d\n", i, lineCount)
		decodeBytes, _ := base64.StdEncoding.DecodeString(scanner.Text())
		_, err = connection.Write(decodeBytes)
		if err != nil {
			log.Println(err)
		}
		time.Sleep(time.Second / time.Duration(frequency))
	}
}
