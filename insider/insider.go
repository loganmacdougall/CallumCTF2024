package main

import (
	"bytes"
	"fmt"
	"net"
	"net/http"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

const unlockBody = `{
	"key": "ruRNZcHJu59BZXuAP24N9Z4zqAN6GmUJ",
	"action": "%s"
}`

var unlockPostUrl string

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	unlockPostUrl = os.Getenv("UNLOCK_POST_URL")

	ln, err := net.Listen("tcp", ":8877")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Unlock URL: %s\n", unlockPostUrl)
	fmt.Println("Server Started on port 8877")

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	conn.Write([]byte("Enter: "))

	buf := make([]byte, 1024)
	_, err := conn.Read(buf)
	if err != nil {
		fmt.Println(err)
		return
	}

	message := strings.TrimRight(string(buf), "\x00\n\b\r\t ")
	fmt.Printf("Received: %s\n", message)

	switch message {
	case "truth{example}":
		conn.Write([]byte("That's just an example truth, give me something real next time\n"))
	case "truth{all_your_db_is_belong_to_us}":
		conn.Write([]byte("https://youtube.com\n"))
		go unlockWebserverFeature("promo")
	case "truth{Str@Ck_CRacK3d_R3Ad_Ev3ry^h1ng!$}":
		conn.Write([]byte("https://youtube.com\n"))
		go unlockWebserverFeature("guns")
	case "truth{L@dies_and_Gentl3men_We_G0t_HiM}":
		conn.Write([]byte("https://youtube.com\n"))
		go unlockWebserverFeature("seized")
	case "minus10", "minus50":
		conn.Write([]byte("I don't need your coupons\n"))
	default:
		conn.Write([]byte("Come back when you have something\n"))
	}
}

func unlockWebserverFeature(feature string) {
	fmt.Printf("Unlocking feature: %s\n", feature)

	var body = []byte(fmt.Sprintf(unlockBody, feature))

	r, err := http.NewRequest("PATCH", unlockPostUrl, bytes.NewBuffer(body))
	if err != nil {
		panic(err)
	}

	r.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	res, err := client.Do(r)
	if err != nil {
		panic(err)
	}

	res.Body.Close()
}
