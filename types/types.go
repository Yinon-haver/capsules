package types

type Capsule struct {
	ID 			int
	FromPhone 	string
	ToPhones	[]string
	PostedOn 	string
	OpenedOn 	string
}

type Message struct {
	FromPhone	string
	Content		string
	Date		string
}

type User struct {
	Phone string
	Token string
}
