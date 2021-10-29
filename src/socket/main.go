package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/gorilla/websocket"
	"net/url"
	"os"
	"os/signal"
	"time"
)


type Socket struct {
	Type string `json:"type"`
	Content struct {
		Symbol string `json:"symbol"`
		TickType string `json:"tickType"`
		Date string `json:"Date"`
		OpenPrice string `json:"openPrice"`
	} `json:"content"`
}

var addr = flag.String("addr","pubwss.bithumb.com", "socket service address")

func main() {
	//flag.Parse()
	//log.SetFlags(0)

	interrupt := make(chan os.Signal,1)
	signal.Notify(interrupt, os.Interrupt)

	u := url.URL{
		Scheme: "wss",
		Host: *addr,
		Path: "/pub/ws",
	}
	fmt.Printf("connecting to %s\n", u.String())

	c, _ , err := websocket.DefaultDialer.Dial(u.String(),nil)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer c.Close()

	msg := []byte(`{"type":"ticker","symbols":["BTC_KRW","ETH_KRW"],"tickTypes":["1H"]}`)
	c.WriteMessage(websocket.TextMessage, msg)

	done := make(chan struct{})
	var socket Socket

	go func() {
		defer close(done)
		for {
			_ , message, err := c.ReadMessage()
			if err != nil {
				fmt.Println("read :", err)
				panic(err)
			}

			if err := json.Unmarshal(message, &socket); err != nil {
				fmt.Printf("recv: %s\n", message)
			}
			fmt.Println("socket : ", socket)
			fmt.Printf("recv: %s\n", message)
		}
	}()

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for {

		select {
		case <- done :
			return
		case t := <-ticker.C:
			err := c.WriteMessage(websocket.PingMessage, []byte(t.String()))
			if err != nil {
				fmt.Println("write:", err)
				panic(err)
			}
		case <-interrupt:
			fmt.Println("interrupt")
			err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				fmt.Println("write close:", err)
				return
			}
			select {
			case <-done:
			case <-time.After(time.Second):
			}
			return
		}
	}
}
