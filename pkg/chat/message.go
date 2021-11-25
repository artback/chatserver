package chat

type Message struct {
	Text string `json:"text" validate:"required"`
}

type Repository interface {
	GetLastMessages(n int) chan Message
	Put(msg Message)
}

type Service interface {
	Connect(topic string, id int64) chan Message
	Broadcast(topic string, msg Message)
	Disconnect(topic string, id int64)
}
