package memory

import (
	"chatserver/pkg/chat"
	"math"
	"sync"
)

type ChatRepository struct {
	sync.RWMutex
	messages []chat.Message
}

func (c *ChatRepository) GetLastMessages(n int) chan chat.Message {
	ch := make(chan chat.Message)
	go func(mux *sync.RWMutex) {
		mux.RLock()
		start := int(math.Max(float64(len(c.messages)-n), 0))
		msgs := c.messages[start:]
		for _, e := range msgs {
			ch <- e
		}
		close(ch)
		mux.RUnlock()
	}(&c.RWMutex)
	return ch
}

func (c *ChatRepository) Put(msg chat.Message) {
	c.Lock()
	c.messages = append(c.messages, msg)
	c.Unlock()
}
