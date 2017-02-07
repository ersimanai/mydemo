package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	_"net/http"
	"github.com/gorilla/websocket"
	"fmt"
	"strings"
	//"mydemo/mydemo/models"
	//"code.google.com/p/go.net/websocket"
	_"golang.org/x/net/websocket"
)

type WebSocketController struct {
	beego.Controller
}

type Message struct {
	User string
	Content string
}
type Connection struct {
	ws *websocket.Conn
	send chan []byte
	username string
}
type hub struct {
	connections map[*Connection]bool
	braodcast chan []byte
	register chan *Connection
	unregister chan *Connection
}

var h = hub {
	braodcast: make(chan []byte),
	register: make(chan *Connection),
	unregister: make(chan *Connection),
	connections: make(map[*Connection]bool),
}

func init() {

	fmt.Println("websocket.go  init")
	go h.run()
}

func (this *WebSocketController) Post() {

	username := this.GetString("uname")
	fmt.Println("websocket Post===",username)
	this.Data["UserName"] = username
	this.Data["btn"] = "发送"
	this.TplName = "websocket.tpl"



}

func (this *WebSocketController) Join() {
	username := this.GetString("uname")

	if len(username) == 0 {
		this.Redirect("/", 302)
		return
	}
	fmt.Println("name = ",username)
	this.Data["UserName"] = username
	this.Data["btn"] = "发送"
	fmt.Println("websocket.go _____Join")
		//c.Ctx.WriteString("Join")
	ws, err := websocket.Upgrade(this.Ctx.ResponseWriter, this.Ctx.Request, nil, 1024, 1024)

	if err != nil {
		fmt.Println("websocket faild")
		return
	}
	fmt.Println("err =",err)

	c := &Connection{send:make(chan []byte, 256), ws: ws, username:username}
	h.register <- c
	defer func() {
		h.unregister <- c
	}()

	go c.writer()
	c.reader()
	

}

func (h *hub) run() {
	for {
		select {
		case c := <-h.register:
			h.connections[c] = true
		case c:= <-h.unregister:
			if _, ok := h.connections[c]; ok {
				delete(h.connections, c)
				close(c.send)
			}
		case m:= <-h.braodcast:
			for c := range h.connections {
				select {
				case c.send <-m:
				default:
					delete(h.connections, c)
					close(c.send)
				}
			}
		}
	} 
}

func (c *Connection) reader() {
	for {
		fmt.Println("server_reader")
		_, message, err := c.ws.ReadMessage()
		if err != nil {
			break
		}
		h.braodcast <- message
	}
	c.ws.Close()

}

func (c *Connection) writer() {
	for result := range c.send {
		name := getName(string(result),":")
		content := getContent(string(result),":")

		message := Message{User:name,Content:content}
		data, err := json.Marshal(message)

		fmt.Println("server_writer:=====",data)

		if err != nil {
			beego.Error("Fail to marshal event:", err)
			return
		}

		err = c.ws.WriteMessage(websocket.TextMessage, data)

		if err != nil {
			break
		}
	}

	c.ws.Close()
}

func getName(str,substr string) string {  
 
  	result := strings.Index(str,substr)    
 
 
    prefix := []byte(str)[0:result]    
   
    rs := []rune(string(prefix))    
     
    
  
    
  return string(rs)  
}

func getContent(str,substr string) string {  
 
  	result := strings.Index(str,substr)    
 
 
    prefix := []byte(str)[result+1:]    
   
    rs := []rune(string(prefix))    
     
  
  
    
  return string(rs)  
}



