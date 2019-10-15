package server

import (
	"encoding/json"
	"github.com/capsules-web-server/config"
	"github.com/capsules-web-server/db"
	"github.com/capsules-web-server/logger"
	"github.com/capsules-web-server/types"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"strconv"
)

type userSocket struct {
	Phone   	string
	Socket 		*websocket.Conn
}

type capsuleChat struct {
	CapsuleChan		chan types.MessageWithDate
	UsersSockets	[]userSocket
}

//var capsulesChat = make(map[int]capsuleChat)

func usersHandler(_ http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case http.MethodPost:
		decoder := json.NewDecoder(request.Body)

		type Params struct {
			Phone	string
		}

		var params Params

		err := decoder.Decode(&params)
		if err != nil {
			logger.Error("parse users post request failed:", err)
			return
		}

		err = db.CreateUser(params.Phone)
		if err != nil {
			logger.Error("create user failed:", err)
			return
		}
	default:
		logger.Error("illegal http method for users")
		return
	}
}

func capsulesHandler(writer http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case http.MethodGet:
		vals := request.URL.Query()

		phone := vals.Get("phone")

		offset, err := strconv.Atoi(vals.Get("offset"))
		if err != nil {
			logger.Error("illegal parameter offset for capsules-web-server get request:", err)
			return
		}

		amount, err := strconv.Atoi(vals.Get("amount"))
		if err != nil {
			logger.Error("illegal parameter amount for capsules-web-server get request:", err)
			return
		}

		isWatched, err := strconv.ParseBool(vals.Get("isWatched"))
		if err != nil {
			logger.Error("illegal parameter isWatched for capsules-web-server get request:", err)
			return
		}

		capsules, err := db.GetCapsules(phone, offset, amount, isWatched)
		if err != nil {
			logger.Error("get last capsules failed:", err)
			return
		}

		err = json.NewEncoder(writer).Encode(capsules)
		if err != nil {
			logger.Error("sending capsules failed:", err)
			return
		}
	case http.MethodPost:
		decoder := json.NewDecoder(request.Body)

		type Params struct {
			Phone		string
			ToPhones	[]string
			Content 	string
			OpenDate	string
		}

		var params Params
		err := decoder.Decode(&params)
		if err != nil {
			logger.Error("parse capsules post request failed:", err)
			return
		}

		err = db.CreateCapsule(params.Phone, params.ToPhones, params.Content, params.OpenDate)
		if err != nil {
			logger.Error("create capsules-web-server failed:", err)
			return
		}
	default:
		logger.Error("illegal http method for capsules")
	}
}

func  openChatConnection(phone string, capsuleID int) (err error) {
	//TODO implement this function
	return nil
}

func chatHandler(writer http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case http.MethodGet:
		vals := request.URL.Query()

		phone := vals.Get("phone")

		capsuleID, err := strconv.Atoi(vals.Get("capsuleID"))
		if err != nil {
			logger.Error("illegal parameter capsuleID for chat get request:", err)
			return
		}

		offset, err := strconv.Atoi(vals.Get("offset"))
		if err != nil {
			logger.Error("illegal parameter offset for chat get request:", err)
			return
		}

		amount, err := strconv.Atoi(vals.Get("amount"))
		if err != nil {
			logger.Error("illegal parameter amount for chat get request:", err)
			return
		}

		messages, err := db.GetMessages(phone, capsuleID, offset, amount)
		if err != nil {
			logger.Error("get last Messages failed:", err)
			return
		}

		err = json.NewEncoder(writer).Encode(messages)
		if err != nil {
			logger.Error("sending messages failed:", err)
			return
		}
	case http.MethodPost:
		decoder := json.NewDecoder(request.Body)

		type Params struct {
			Phone     string
			CapsuleID int
		}

		var params Params

		err := decoder.Decode(&params)
		if err != nil {
			logger.Error("parse chat post request failed:", err)
			return
		}

		err = openChatConnection(params.Phone, params.CapsuleID)
		if err != nil {
			logger.Error("open chat connection failed:", err)
		}
	default:
		logger.Error("illegal http method for chat")
	}
}

func Run() {
	mux := http.NewServeMux()

	mux.HandleFunc("/users", usersHandler)
	mux.HandleFunc("/capsules", capsulesHandler)
	mux.HandleFunc("/chat", chatHandler)

	err := http.ListenAndServe(":" + strconv.Itoa(config.GetPort()), mux)
	if err != nil {
		log.Fatal("listening fail:", err)
	}
}
