package main

import (
	"bufio"
	"fmt"
	"net"
	"os"

	"github.com/TRedzepagic/compositelogger/logs"
)

// Read autoreads UDP packets
func Read(connection *net.UDPConn, ReplyLog *logs.CompositeLog, TimeLog *logs.CompositeLog) {
	for {
		buffer := make([]byte, 1024)
		n, _, err := connection.ReadFromUDP(buffer)
		if err != nil {
			ReplyLog.Error(err)
			return
		}
		if string(buffer[0:n]) == "Timer Tick! : 10 seconds have elapsed" {
			TimeLog.Info("My address : " + connection.LocalAddr().String() + " " + string(buffer[0:n]))
		} else {
			ReplyLog.Info(string(buffer[0:n]))
			return
		}
	}
}

func main() {
	fileLog := logs.NewFileLogger("filelog")
	defer fileLog.Close()
	stdoutLog := logs.NewStdLogger()
	defer stdoutLog.Close()

	wantDebug := false

	ErrInfoLog := logs.NewCustomLogger(wantDebug, fileLog, stdoutLog)
	ReplyLog := logs.NewCustomLogger(wantDebug, stdoutLog)
	TimeLog := logs.NewCustomLogger(wantDebug, fileLog)

	//Expecting valid address host:port format
	arguments := os.Args
	if len(arguments) == 1 {
		ErrInfoLog.Info("No address specified... Exiting.....")
		return
	}
	addr := arguments[1]
	str, err := net.ResolveUDPAddr("udp4", addr) //server address
	connection, err := net.DialUDP("udp4", nil, str)
	if err != nil {
		ErrInfoLog.Error(err)
		return
	}

	fmt.Printf("Dialed address : %s\n", connection.RemoteAddr().String())
	defer connection.Close()

	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Text to send : ")
		text, _ := reader.ReadString('\n')
		data := []byte(text)
		_, err = connection.Write(data)

		if err != nil {
			ErrInfoLog.Error(err)
			return
		}
		go Read(connection, ReplyLog, TimeLog)

	}
}
