package main

import (
	"fmt"
	"log"
	"time"

	"golang.org/x/net/websocket"
)

type Client struct {
	URL    string
	Origin string
	Jwt    string
}

func NewWebsocketClient(origin, url, jwt string) *Client {
	return &Client{
		Origin: origin,
		URL:    url,
		Jwt:    jwt,
	}
}

func (c *Client) ReceiveMessage() {
	cfg, err := websocket.NewConfig(c.URL, c.Origin)
	if err != nil {
		log.Println("cfg:", err)
	}
	cfg.Header.Add("Jwt", c.Jwt)
	cfg.Header.Add("Sec-WebSocket-Protocol", "protoo")

	ws, err := websocket.DialConfig(cfg)
	if err != nil {
		log.Fatal("ws:", err)
	}

	go func() {
		// Reading
		flag := 1
		for {
			var msg = make([]byte, 512)
			m, err := ws.Read(msg)
			if err != nil {
				flag++
				log.Println(err)
				time.Sleep(time.Duration(flag) * 200 * time.Millisecond)
				continue
			}

			log.Println("Receive:", string(msg[:m]))
			time.Sleep(50 * time.Millisecond)
			// log.Println("Read thread after sleep...")
		}

	}()

	// defer ws.Close() //关闭连接
}

func (c *Client) SendMessage(user string, body string) {
	url := "ws://192.168.0.201:8083/dispatch"
	cfg, err := websocket.NewConfig(url, c.Origin)
	if err != nil {
		log.Println("cfg:", err)
	}
	cfg.Header.Add("Jwt", "password")
	cfg.Header.Add("Sec-WebSocket-Protocol", "protoo")

	ws, err := websocket.DialConfig(cfg)
	if err != nil {
		log.Fatal("ws:", err)
	}

	go func() {
		for i := range [30]string{} {
			mj := fmt.Sprintf("{\"to\":%s, \"data\":\"%s %d\"}", string(user[5]), body, i)
			_, err = ws.Write([]byte(mj))
			if err != nil {
				log.Println(err)
			}
			log.Printf("Send to %s: %s\n", user, mj)
			time.Sleep(3 * time.Second)
		}
	}()

	// defer ws.Close() //关闭连接
}

func main() {
	origin := "http://192.168.0.201:8083"
	url := "ws://192.168.0.201:8083/ws"
	whh := "test/1"
	lsd := "test/2"

	c := NewWebsocketClient(origin, url, whh)
	msg := "hello,I'm wanghuahua, I like golang!"
	c.ReceiveMessage()
	c.SendMessage(whh, msg)

	c2 := NewWebsocketClient(origin, url, lsd)
	msg2 := "hi,I'm lishuandan, I hate golang!"
	c2.ReceiveMessage()
	c2.SendMessage(lsd, msg2)
	select {}
}
