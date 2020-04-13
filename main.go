package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	//Expecting valid address host:port format
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("No port number specified... Exiting.....")
		return
	}

	PORT := arguments[1]
	str, err := net.ResolveUDPAddr("udp4", PORT)
	connection, err := net.DialUDP("udp4", nil, str)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Connected to : %s\n", connection.RemoteAddr().String())
	defer connection.Close()

	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Text to send : ")
		text, _ := reader.ReadString('\n')
		data := []byte(text)
		_, err = connection.Write(data)

		if strings.TrimSpace(string(data)) == "STOP" {
			fmt.Println("Exiting UDP client!")
			return
		}

		if err != nil {
			fmt.Println(err)
			return
		}

		buffer := make([]byte, 1024)
		n, _, err := connection.ReadFromUDP(buffer)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Printf("Server replies : %s\n", string(buffer[0:n]))
	}
}
