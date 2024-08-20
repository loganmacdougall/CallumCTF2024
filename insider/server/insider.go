package main

import (
	"bytes"
	"fmt"
	"math"
	"math/big"
	"math/rand"
	"net"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

type Challenge struct {
	message string
	answer  *big.Int
}

const unlockBody = `{
	"key": "ruRNZcHJu59BZXuAP24N9Z4zqAN6GmUJ",
	"action": "%s"
}`

var unlockPostUrl string
var canTakedown = false

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
		conn.Write([]byte("https://youtu.be/evOJ5JZVmj0\n"))
		go unlockWebserverFeature("promo")
	case "truth{Str@Ck_CRacK3d_R3Ad_Ev3ry^h1ng!$}":
		conn.Write([]byte("https://youtu.be/Bp6PFaw7XC8\n"))
		canTakedown = true
		go unlockWebserverFeature("guns")
	case "truth{L@dIes_anD_g3nTlemEn_w3_g0T_h1m}":
		conn.Write([]byte("The website has been seized. thank you for your service\n"))
		go unlockWebserverFeature("seized")
	case "minus10", "minus50":
		conn.Write([]byte("I don't need your coupons\n"))
	case "initiate_takedown":
		if canTakedown {
			handleTakedown(conn)
		} else {
			conn.Write([]byte("I don't takedown anything with evidence\n"))
		}
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

func handleTakedown(conn net.Conn) {
	fmt.Println("Initiating takedown")

	conn.Write([]byte("Alright, I'll help you take them down, but don't think it's going " +
		"to be easy. You're going to help me out. There's lots of math involved.\n\n" +
		"Example, if you see:\n\nSTART\nADD_ 1\nADD_ 2\nEND\n: \n\nyou type '3'\nAre you ready? (y/n) "))

	buf := make([]byte, 65536)
	_, err := conn.Read(buf)
	if err != nil {
		fmt.Println(err)
		return
	}

	message := strings.ToLower(strings.TrimRight(string(buf), "\x00\n\b\r\t "))

	if message != "y" && message != "yes" {
		conn.Write([]byte("Come back when you're ready\n"))
		return
	}

	passed := true
	for i := 1; i <= 25 && passed; i++ {
		message = ""
		challenge := takedownChallenge(i)
		conn.Write([]byte("\nSTART\n"))
		conn.Write([]byte(challenge.message))
		conn.Write([]byte("END\n: "))

		for {
			count, err := conn.Read(buf)
			if err != nil {
				fmt.Println(err)
				conn.Write([]byte("Error occurred reading your answer, try again\n"))
				return
			}

			message += string(buf[:count])
			if strings.Contains(message[len(message)-count:], "\n") {
				break
			}
		}

		message = strings.ToLower(strings.TrimRight(message, "\x00\n\b\r\t "))

		if message != challenge.answer.String() {
			passed = false
			break
		}
	}

	if passed {
		conn.Write([]byte("Good work - we now have everything we need to take them down\nenter this truth and finish the job\n\ntruth{L@dIes_anD_g3nTlemEn_w3_g0T_h1m}\n"))

	} else {
		conn.Write([]byte("Inncorect, feel free to come back and try again\n"))
	}
}

func takedownChallenge(level int) Challenge {

	const (
		Add int = 1
		Sub     = 2
		Mul     = 3
		Sft     = 4
	)

	challenge_msg := ""
	challenge_ans := big.NewInt(0)
	length := int(math.Pow(2, float64(level*3)/5.0)) + 1
	max_num := int64(math.Pow(3.0, float64(level)))

	num := big.NewInt(0)
	for i := 0; i < length; i++ {
		op := Add
		if level > 5 && rand.Int31n(4) == 0 {
			op = Sub
		} else if level > 10 && rand.Int31n(5) == 0 {
			op = Mul
		} else if level > 15 && rand.Int31n(5) == 0 {
			op = Sft
		}

		opStr := ""
		switch op {
		case Add:
			opStr = "ADD_ "
			num.SetInt64(rand.Int63n(max_num))
			challenge_ans.Add(challenge_ans, num)
		case Sub:
			opStr = "SUB_ "
			num.SetInt64(rand.Int63n(max_num))
			challenge_ans.Sub(challenge_ans, num)
		case Mul:
			opStr = "MULT "
			num.SetInt64(rand.Int63n(max_num))
			challenge_ans.Mul(challenge_ans, num)
		case Sft:
			opStr = "LSFT "
			randLShft := uint(rand.Int31n(15) + 1)
			challenge_ans.Lsh(challenge_ans, randLShft)
			challenge_msg += opStr + strconv.Itoa(int(randLShft)) + "\n"
		}

		if opStr != "LSFT " {
			challenge_msg += opStr + num.String() + "\n"
		}
	}

	return Challenge{challenge_msg, challenge_ans}
}
