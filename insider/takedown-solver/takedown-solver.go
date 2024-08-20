package main

import (
	"fmt"
	"math/big"
	"net"
	"strconv"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", ":8877")
	if err != nil {
		fmt.Println(err)
		return
	}

	data := make([]byte, 1048576)
	_, err = conn.Read(data)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Fprintf(conn, "initiate_takedown\n")

	_, err = conn.Read(data)
	if err != nil {
		fmt.Println(err)
		return
	}

	complete_data := ""
	fmt.Fprintf(conn, "y\n")

	for level := 1; true; level++ {
		count := 0
		complete_data = ""

		for {
			count, err = conn.Read(data)
			if err != nil {
				fmt.Print("End of message, hopefully all challenges are completed\nPrinting received message:\n\n")
				fmt.Println(complete_data)
				conn.Close()
				return
			}

			complete_data += string(data[:count])
			length := len(complete_data)
			if strings.Contains(complete_data[length-count:length], "END") {
				break
			} else {
				fmt.Println("Last message didn't contain \"END\", reading more")
			}
		}

		fmt.Print(complete_data)
		message := strings.Split(complete_data, "\n")
		started := false
		num := big.NewInt(0)
		temp := big.NewInt(0)
		for i := 0; message[i] != "END"; i++ {
			msg := message[i]

			if started {
				switch msg[0:4] {
				case "ADD_":
					temp.SetString(msg[5:], 10)
					num.Add(num, temp)
				case "SUB_":
					temp.SetString(msg[5:], 10)
					num.Sub(num, temp)
				case "MULT":
					temp.SetString(msg[5:], 10)
					num.Mul(num, temp)
				case "LSFT":
					lNum, err := strconv.Atoi(msg[5:])
					if err != nil {
						panic(err)
					}
					num.Lsh(num, uint(lNum))
				}
			}

			if msg == "START" {
				started = true
			}
		}

		fmt.Printf("%s (%d)\n", num.String(), level)
		fmt.Fprintf(conn, "%s\n", num.String())
	}
}
