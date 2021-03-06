package db

import (
	"database/sql"
	"fmt"
	"github.com/capsules-web-server/config"
	"github.com/capsules-web-server/logger"
	"github.com/capsules-web-server/types"
	"github.com/capsules-web-server/utils"
	"github.com/lib/pq"
	_ "github.com/lib/pq"
)

var db *sql.DB

func Init() {
	var err error
	connection, err := pq.ParseURL(config.GetDBUrl())
	db, err = sql.Open("postgres", connection)
	//db, err = sql.Open("postgres", "dbname=capsules sslmode=disable")
	if err != nil {
		logger.Error("fail to connect to sql:", err)
	}
}

func CreateUser(phone string, token string) (err error) {
	_, err = db.Exec(fmt.Sprintf("INSERT INTO users VALUES('%s', '%s')", phone, token))

	if err != nil {
		logger.Error("fail to insert user to users table in db:", err)
	}

	return
}

func CreateCapsule(phone string, toPhones []string, content string, openDate string) (err error) {
	var capsuleID int
	err = db.QueryRow(fmt.Sprintf("INSERT INTO capsules(from_phone, posted_on, opened_on) VALUES('%s','%s','%s') RETURNING id", phone, utils.GetTimestampString(), openDate)).Scan(&capsuleID)
	if err != nil {
		logger.Error("fail to insert capsule to capsules table in db:", err)
		return
	}

	for _, phone := range toPhones {
		_, err = db.Exec(fmt.Sprintf("INSERT INTO capsules_users(capsule_id, user_phone, is_watched) VALUES(%d,'%s',%t)", capsuleID, phone, false))
		if err != nil {
			logger.Error("fail to insert (capsule, user) to capsule_user table in db:", err)
		}
	}
	_, err = db.Exec(fmt.Sprintf("INSERT INTO capsules_users(capsule_id, user_phone, is_watched) VALUES(%d,'%s',%t)", capsuleID, phone, false))
	if err != nil {
		logger.Error("fail to insert (capsule, user) to capsule_user table in db:", err)
		return
	}

	_, err = db.Exec(fmt.Sprintf("INSERT INTO messages(capsule_id, message_date, from_user, content) VALUES(%d,'%s','%s','%s')", capsuleID, utils.GetTimestampString(), phone, content))
	if err != nil {
		logger.Error("fail to insert message to messages table in db:", err)
	}

	return
}

func GetCapsules(phone string, offset int, amount int, isWatched bool) (capsules []types.Capsule, err error) {
	capsulesRows, err := db.Query(fmt.Sprintf(
		`SELECT id,from_phone,posted_on,opened_on, array_agg(user_phone)
				FROM
				(SELECT id,from_phone,posted_on,opened_on
					FROM capsules, capsules_users
					WHERE capsules.id = capsules_users.capsule_id AND user_phone = '%s' AND is_watched = %v) AS capsules,
				capsules_users
				WHERE capsules_users.capsule_id = capsules.id AND capsules.from_phone <> capsules_users.user_phone
				GROUP BY id,from_phone,posted_on,opened_on
				ORDER BY opened_on ASC
				LIMIT %d
				OFFSET %d`, phone, isWatched, amount, offset))
	if err != nil {
		logger.Error("fail to get the capsules of user from the db:", err)
		return
	}

	for capsulesRows.Next() {
		var capsule types.Capsule

		err = capsulesRows.Scan(&capsule.ID, &capsule.FromPhone, &capsule.PostedOn, &capsule.OpenedOn, pq.Array(&capsule.ToPhones))
		if err != nil {
			logger.Error("fail to get capsule from capsules table row:", err)
			continue
		}

		capsules = append(capsules, capsule)
	}

	return
}

func GetMessages(phone string, capsuleID int, offset int, amount int) (messages []types.Message, err error) {
	rows, err := db.Query(fmt.Sprintf("SELECT from_user, content, message_date FROM messages WHERE capsule_id = %d ORDER BY message_date DESC LIMIT %d OFFSET %d", capsuleID, amount, offset))
	if err != nil {
		logger.Error("fail to get the messages from messages table in the db:", err)
		return
	}

	for rows.Next() {
		var message types.Message

		err = rows.Scan(&message.FromPhone, &message.Content, &message.Date)
		if err != nil {
			logger.Error("fail to get message from messages table row:", err)
			continue
		}

		messages = append(messages, message)
	}

	_, err = db.Exec(fmt.Sprintf("UPDATE capsules_users SET is_watched = true WHERE capsule_id = %d AND user_phone = '%s'", capsuleID, phone))
	if err != nil {
		logger.Error("fail to insert message to messages table in db:", err)
	}

	return
}

func GetCapsuleUsers(capsuleID int) (users []types.User, err error) {
	rows, err := db.Query(fmt.Sprintf("SELECT phone, token"+
		" FROM users, (SELECT user_phone FROM capsules_users WHERE capsules_users.capsule_id = %d) as phones"+
		" WHERE phone = user_phone", capsuleID))
	if err != nil {
		logger.Error("fail to get the users capsules from users, capsules_users tables in the db:", err)
		return
	}

	for rows.Next() {
		var user types.User

		err = rows.Scan(&user.Phone, &user.Token)
		if err != nil {
			logger.Error("fail to get user from users table row:", err)
			return
		}

		users = append(users, user)
	}

	return
}

func SaveMessage(capsuleID int, message types.Message) (err error) {
	_, err = db.Exec(fmt.Sprintf("INSERT INTO messages(capsule_id, from_user, content, message_date) "+
		"VALUES(%d,'%s','%s','%s')", capsuleID, message.FromPhone, utils.MultiplyQuote(message.Content), utils.GetTimestampString()))
	if err != nil {
		logger.Error("fail to insert message to messages table in db:", err)
	}

	return
}

//func GetToken(phone string) (token string, err error) {
//	rows, err := db.Query(fmt.Sprintf("SELECT token FROM users WHERE phone = '%s'", phone))
//	if err != nil {
//		logger.Error("fail to get the token of user from users table in the db:", err)
//	}
//
//	for rows.Next() {
//		err = rows.Scan(&token)
//		if err != nil {
//			logger.Error("fail to get token from users table row:", err)
//		}
//	}
//
//	return
//}
