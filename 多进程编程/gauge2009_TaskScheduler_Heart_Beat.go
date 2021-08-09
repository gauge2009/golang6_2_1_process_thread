package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/url"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/websocket"
)
var addr_taskscheduler = flag.String("addr", "localhost:10088", "TaskScheduler service address")
type SocketModel struct {
	//Index int16
	FunctionName string
	SiteCode string
	EmpID string	`json:"empid"`
	Token string	`json:"token"`
	Data string
	StructureID string
	//Info int8
}
func BuildSocketModel(FunctionName string, SiteCode string, EmpID string, Token string, Data string, StructureID string) *SocketModel {
	return &SocketModel{
		FunctionName:FunctionName,
		SiteCode:SiteCode,
		EmpID:EmpID,
		Token:Token,
		Data:Data,
		StructureID:StructureID,
	}
}
func main() {
	flag.Parse()
	log.SetFlags(0)
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)
	u := url.URL{Scheme: "ws", Host: *addr_taskscheduler, Path: "/echo"}
	log.Printf("connecting to %s", u.String())
	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		/// 心跳异常情形1）server端在建立连接前已然宕机， client 连接创建时就会err —— Err 10061
		/// TODO : 发送email


		log.Fatal("dial:", err)
	}
	defer c.Close()
	done := make(chan struct{})
	go func() {
		defer close(done)
		for {
			_, message, err := c.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				return
			}
			log.Printf("recv: %s", message)
		}
	}()
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()
	for {
		select {
		case <-done:
			return
		case t := <-ticker.C:
			/// DEMO1:
			socketModel1 := []byte(t.String())
			fmt.Println(socketModel1)
			/// DEMO2:
			//var socketModel SocketModel
			//socketModel := new(SocketModel)
			//socketModel :=  &SocketModel{}
			socketModel:= BuildSocketModel("ping_from_crystal_beacon","Company1","N/A","N/A","","N/A")
			socketModelbytes, err0 := json.Marshal(socketModel)
			fmt.Println(socketModel)
			fmt.Println(string(socketModelbytes))
			//err = json.Unmarshal(buf, mainConfig)
			if err0 != nil {
				log.Panicln("decode  failed:", string(socketModelbytes), err0)
			}
			err := c.WriteMessage(websocket.TextMessage, socketModelbytes)
			if err != nil {
				/// 心跳异常情形2）如果client 在成功连接server后，server端宕机，client 在后续发送消息时也会err——WSAECONNABORTED (10053)【WSAE CONN ABORTED】
				/// TODO : 发送email


				log.Println("write:", err)
				return
			}
		case <-interrupt:
			log.Println("interrupt")

		   // I/O阻塞
			err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				log.Println("write close:", err)
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
