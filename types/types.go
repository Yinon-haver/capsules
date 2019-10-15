package types

type Capsule struct {
	ID 			int
	FromPhone 	string
	ToPhones	[]string
	PostedOn 	string
	OpenedOn 	string
}

type MessageWithDate struct {
	FromPhone	string
	Content		string
	Date		string
}
