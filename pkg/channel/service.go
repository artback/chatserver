package channel

import "chatserver/pkg/chat"

type service struct {
	channels map[string]map[int64]chan chat.Message
	repo     chat.Repository
}

func NewService(repo chat.Repository) chat.Service {
	return &service{repo: repo, channels: make(map[string]map[int64]chan chat.Message)}
}
func (s service) Disconnect(topic string, id int64) {
	delete(s.channels[topic], id)
}

func (s service) Broadcast(topic string, msg chat.Message) {
	s.repo.Put(msg)
	for _, c := range s.channels[topic] {
		go func(c chan chat.Message) {
			c <- msg
		}(c)
	}
}

func (s service) Connect(topic string, id int64) chan chat.Message {
	if s.channels[topic][id] == nil {
		s.channels[topic] = make(map[int64]chan chat.Message)
	}
	c := make(chan chat.Message)
	s.channels[topic][id] = c

	return s.repo.GetLastMessages(50)
}
