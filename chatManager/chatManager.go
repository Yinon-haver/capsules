package chatManager

import (
	"fmt"
	"github.com/capsules-web-server/db"
	"github.com/capsules-web-server/logger"
	"github.com/capsules-web-server/types"
	"github.com/capsules-web-server/utils"
	"github.com/gorilla/websocket"
	expo "github.com/oliveroneill/exponent-server-sdk-golang/sdk"
	"strconv"
)

/*************************************  sockets table structs  *************************************/

type userSocketStruct struct {
	Token		string
	Socket 		*websocket.Conn
}

type socketsTableEntryStruct struct {
	UsersSockets map[string]userSocketStruct
	ConnectedNumber int
}

type socketTableStruct map[int]socketsTableEntryStruct

/*************************************  channels structs  *************************************/

type addCapsuleStruct struct {
	CapsuleID	int
	Users 		socketsTableEntryStruct
}

type connectUserStruct struct {
	CapsuleID	int
	Phone		string
	Socket 		*websocket.Conn
}

type disconnectUserStruct struct {
	CapsuleID	int
	Phone		string
}

type receiveMessagesStruct struct {
	CapsuleID	int
	Message types.Message
}

type sendMessagesStruct struct {
	CapsuleID	int
	Token 		string
	Socket		*websocket.Conn
	Message 	types.Message
	ToPhone		string
}

type notificationStruct struct {
	CapsuleID	int
	Token		string
	Message types.Message
}

type getCapsuleUsersStruct struct {
	CapsuleID	int
	Phone		string
	Socket		*websocket.Conn
}

type getCapsuleTokensStruct struct {
	CapsuleID	int
	Message types.Message
}

type saveMessagesStruct struct {
	CapsuleID	int
	Message types.Message
}

/**********************************************************************************************/

var addCapsuleChannel 		chan addCapsuleStruct
var connectUserChannel 		chan connectUserStruct
var disconnectUserChannel	chan disconnectUserStruct
var receiveMessagesChannel	chan receiveMessagesStruct
var sendMessagesChannel		chan sendMessagesStruct
var notificationChannel		chan notificationStruct
var getCapsuleUsersChannel	chan getCapsuleUsersStruct
var getCapsuleTokensChannel	chan getCapsuleTokensStruct
var saveMessagesChannel		chan saveMessagesStruct
var socketsTable			socketTableStruct

func Init() {
	utils.RunProcess(requestsHandlerProcess, 1)
	utils.RunProcess(getCapsuleTokensProcess, 1)
	utils.RunProcess(getCapsuleUsersProcess, 1)
	utils.RunProcess(sendMessagesProcess, 1)
	utils.RunProcess(notificationsProcess, 1)
	utils.RunProcess(saveMessagesProcess, 1)
}

func disconnectUser(disconnectUserRequest disconnectUserStruct) {
	socketsTableEntry, ok := socketsTable[disconnectUserRequest.CapsuleID]
	if ok {
		userSocket, ok := socketsTableEntry.UsersSockets[disconnectUserRequest.Phone]
		if !ok {
			logger.Error("some user not found in sockets table entry")
			return
		}

		if userSocket.Socket != nil {
			userSocket.Socket = nil
			socketsTableEntry.ConnectedNumber--
			if socketsTableEntry.ConnectedNumber == 0 {
				delete(socketsTable, disconnectUserRequest.CapsuleID)
			} else {
				socketsTableEntry.UsersSockets[disconnectUserRequest.Phone] = userSocket
				socketsTable[disconnectUserRequest.CapsuleID] = socketsTableEntry
			}
		}
	}
}

func connectUser(connectUserRequest connectUserStruct) {
	socketsTableEntry, ok := socketsTable[connectUserRequest.CapsuleID]
	if ok {
		userSocket, ok := socketsTableEntry.UsersSockets[connectUserRequest.Phone]
		if !ok {
			logger.Error("some user not found in sockets table entry")
			return
		}

		if userSocket.Socket != nil {
			logger.Error("try to connect user that already connected")
			return
		}

		userSocket.Socket = connectUserRequest.Socket
		socketsTableEntry.UsersSockets[connectUserRequest.Phone] = userSocket
		socketsTableEntry.ConnectedNumber++
		socketsTable[connectUserRequest.CapsuleID] = socketsTableEntry
	} else {
		getCapsuleUsersRequest := getCapsuleUsersStruct{CapsuleID: connectUserRequest.CapsuleID,
														Phone: connectUserRequest.Phone,
														Socket: connectUserRequest.Socket}
		getCapsuleUsersChannel <- getCapsuleUsersRequest
	}
}

func addCapsule(addCapsuleRequest addCapsuleStruct) {
	socketsTable[addCapsuleRequest.CapsuleID] = addCapsuleRequest.Users
}

func receiveMessages(receiveMessagesRequest receiveMessagesStruct) {
	saveMessagesChannel <- saveMessagesStruct{CapsuleID: receiveMessagesRequest.CapsuleID,
												Message: receiveMessagesRequest.Message}

	socketsTableEntry, ok := socketsTable[receiveMessagesRequest.CapsuleID]
	if ok {
		for toPhone, userSocket := range socketsTableEntry.UsersSockets {
			if userSocket.Socket != nil {

				sendMessagesChannel <- sendMessagesStruct{Message: receiveMessagesRequest.Message,
															Token: userSocket.Token,
															CapsuleID: receiveMessagesRequest.CapsuleID,
															Socket: userSocket.Socket,
															ToPhone: toPhone}
			} else {
				notificationChannel <- notificationStruct{CapsuleID: receiveMessagesRequest.CapsuleID,
															Message: receiveMessagesRequest.Message,
															Token: userSocket.Token}
			}
		}
	} else {
		getCapsuleTokensChannel <- getCapsuleTokensStruct{Message:receiveMessagesRequest.Message,
															CapsuleID:receiveMessagesRequest.CapsuleID}
	}
}

func requestsHandlerProcess() {
	for {
		select {
		case addCapsuleRequest := <-addCapsuleChannel:
			addCapsule(addCapsuleRequest)
		case connectUserRequest := <-connectUserChannel:
			connectUser(connectUserRequest)
		case messageRequest := <-receiveMessagesChannel:
			receiveMessages(messageRequest)
		case disconnectUserRequest := <-disconnectUserChannel:
			disconnectUser(disconnectUserRequest)
		}
	}
}

func getCapsuleUsersProcess() {
	for {
		getCapsuleUsersRequest := <- getCapsuleUsersChannel

		users, err := db.GetCapsuleUsers(getCapsuleUsersRequest.CapsuleID)
		if err != nil {
			logger.Error("get capsule users failed")
			continue
		}

		var socketsTableEntry socketsTableEntryStruct

		for _, user := range users {
			socketsTableEntry.UsersSockets[user.Phone] = userSocketStruct{Token: user.Token, Socket: nil}
		}

		senderUser, _ := socketsTableEntry.UsersSockets[getCapsuleUsersRequest.Phone]
		senderUser.Socket = getCapsuleUsersRequest.Socket
		socketsTableEntry.UsersSockets[getCapsuleUsersRequest.Phone] = senderUser
		socketsTableEntry.ConnectedNumber = 1

		addCapsuleChannel <- addCapsuleStruct{CapsuleID:getCapsuleUsersRequest.CapsuleID,
												Users: socketsTableEntry}
	}
}

func getCapsuleTokensProcess() {
	for {
		getCapsuleTokensRequest := <- getCapsuleTokensChannel

		users, err := db.GetCapsuleUsers(getCapsuleTokensRequest.CapsuleID)
		if err != nil {
			logger.Error("get capsule users failed")
			continue
		}

		for _, user := range users {
			notificationChannel <- notificationStruct{CapsuleID: getCapsuleTokensRequest.CapsuleID,
														Message: getCapsuleTokensRequest.Message,
														Token: user.Token}
		}
	}
}

func saveMessagesProcess() {
	for {
		saveMessageRequest := <- saveMessagesChannel

		err := db.SaveMessage(saveMessageRequest.CapsuleID, saveMessageRequest.Message)
		if err != nil {
			logger.Error("save message failed")
		}
	}
}

func sendMessagesProcess() {
	for {
		sendMessageRequest := <-sendMessagesChannel
		err := sendMessageRequest.Socket.WriteJSON(sendMessageRequest.Message)
		if err != nil {
			notificationChannel <- notificationStruct{CapsuleID: sendMessageRequest.CapsuleID,
				Token:   sendMessageRequest.Token,
				Message: sendMessageRequest.Message}
			disconnectUserChannel <- disconnectUserStruct{CapsuleID: sendMessageRequest.CapsuleID,
				Phone: sendMessageRequest.ToPhone}
		}
	}
}

func notificationsProcess() {
	for {
		notificationRequest := <- notificationChannel

		// To check the token is valid
		pushToken, err := expo.NewExponentPushToken(fmt.Sprintf("ExponentPushToken[%s]", notificationRequest.Token))
		if err != nil {
			logger.Info("token not valid", err)
			continue
		}

		// Create a new Expo SDK client
		client := expo.NewPushClient(nil)

		// Publish message
		response, err := client.Publish(
			&expo.PushMessage{
				To: pushToken,
				//Body: "This is a test notification",
				Data: map[string]string{"CapsuleID": strconv.Itoa(notificationRequest.CapsuleID),
										"FromPhone": notificationRequest.Message.FromPhone,
										"Content": notificationRequest.Message.Content,
										"Date": notificationRequest.Message.Date},
				//Sound: "default",
				//Title: "Notification Title",
				Priority: expo.DefaultPriority,
			},
		)
		// Check errors
		if err != nil {
			logger.Error("notification sending failed", err)
			continue
		}
		// Validate responses
		if response.ValidateResponse() != nil {
			logger.Error(response.PushMessage.To, "failed")
		}
	}
}

func RunChat(ws *websocket.Conn, phone string, capsuleID int) {
	var receiveMessagesRequest receiveMessagesStruct
	receiveMessagesRequest.CapsuleID = capsuleID

	for {
		err := ws.ReadJSON(&receiveMessagesRequest.Message.Content)

		if err != nil {
			logger.Info("read from socket fail:", err)
			disconnectUserChannel <- disconnectUserStruct{CapsuleID:capsuleID, Phone:phone}
			return
		}

		receiveMessagesRequest.Message.Date = utils.GetTimestampString()

		receiveMessagesChannel <- receiveMessagesRequest
	}
}
