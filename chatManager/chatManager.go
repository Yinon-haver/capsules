package chatManager

import (
	"github.com/gorilla/websocket"
)

//import (
//	"github.com/capsules-web-server/config"
//	"github.com/capsules-web-server/logger"
//	"github.com/capsules-web-server/utils"
//	"github.com/gorilla/websocket"
//	"log"
//)
//
//type CapsuleID 		int
//type Phone 			string
//
//type usersSockets struct {
//	UsersSockets	map[Phone]*websocket.Conn
//	ConnectedNum	int
//}
//
//type messageType 	bool
//const (
//	Connection 		messageType = true
//	Chat			messageType = false
//)
//
//type broadcastMessage interface {
//	getType() 		messageType
//	getCapsuleID()	CapsuleID
//	getContent()	string
//	getFromPhone()	Phone
//}
//
//type chatMessage struct {
//	Type			messageType
//	FromPhone		Phone
//	CapsuleID		CapsuleID
//	Content 		string
//}
//
//type connectMessage struct {
//	Type			messageType
//	FromPhone		Phone
//	CapsuleID		CapsuleID
//	Participants	[]Phone
//	Content 		string
//}
//
//func (msg chatMessage) getType() messageType {
//	return msg.Type
//}
//func (msg chatMessage) getCapsuleID() CapsuleID {
//	return msg.CapsuleID
//}
//func (msg chatMessage) getContent() string {
//	return msg.Content
//}
//func (msg chatMessage) getFromPhone() Phone {
//	return msg.FromPhone
//}
//
//func (msg connectMessage) getType() messageType {
//	return msg.Type
//}
//func (msg connectMessage) getCapsuleID() CapsuleID {
//	return msg.CapsuleID
//}
//func (msg connectMessage) getContent() string {
//	return msg.Content
//}
//func (msg connectMessage) getFromPhone() Phone {
//	return msg.FromPhone
//}
//
//
//var broadcastChannel chan broadcastMessage
//var users map[Phone]CapsuleID
//var capsulesUsers map[CapsuleID]usersSockets
//
func Init() {
	//TODO

	//broadcastChannel = make(chan broadcastMessage, config.GetBrodcastChannelSize())
	//capsulesUsers = make(map[CapsuleID]usersSockets)
	//
	//go func() {
	//	for {
	//		msg := <-broadcastChannel
	//
	//		if msg.getType() == Connection {
	//			users[msg.getFromPhone()] = msg.getCapsuleID()
	//			usersSockets, ok := capsulesUsers[msg.getCapsuleID()]
	//			if ok {
	//				usersSockets.UsersSockets[msg.getFromPhone()] =
	//				usersSockets.ConnectedNum++
	//			}
	//			return
	//		}
	//
	//		usersSockets := capsulesUsers[msg.getCapsuleID()]
	//
	//		for phone, socket := range usersSockets.UsersSockets {
	//			if socket == nil {
	//				//TODO
	//				return
	//			}
	//
	//			err := socket.WriteJSON(msg.getContent())
	//			if err != nil {
	//				//TODO
	//				socket.Close()
	//				usersSockets.UsersSockets[phone] = nil
	//				usersSockets.ConnectedNum--
	//				if usersSockets.ConnectedNum == 0 {
	//					delete(capsulesUsers, msg.getCapsuleID())
	//				}
	//
	//				delete(users, phone)
	//			}
	//		}
	//	}
	//}()
}

func RunChat(ws *websocket.Conn, phone string, capsuleID int) error {

	//TODO
	return nil

	//// Register our new client
	//clients[ws] = true
	//
	//for {
	//	var msg broadcastMessage
	//
	//	err := ws.ReadJSON(&msg.Content)
	//	if err != nil {
	//		logger.Warning("fail to read from socket:", err)
	//		delete(clients, ws)
	//		break
	//	}
	//
	//	msg.FromPhone = phone
	//	msg.CapsuleID = capsuleID
	//	msg.Date = utils.GetTimestampString()
	//
	//	broadcast <- msg
	//}
}
