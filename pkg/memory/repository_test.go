package memory

import (
	"chatserver/pkg/chat"
	"reflect"
	"testing"
)

func TestChatRepository_GetLastMessages(t *testing.T) {
	type fields struct {
		messages []chat.Message
	}
	type args struct {
		n int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []chat.Message
	}{
		{
			name: "want 1 last message",
			fields: fields{
				messages: []chat.Message{
					{Text: "1"},
					{Text: "2"},
					{Text: "3"},
					{Text: "4"},
				},
			},
			args: args{
				n: 1,
			},
			want: []chat.Message{{Text: "4"}},
		},
		{
			name: "want 4 last message",
			fields: fields{
				messages: []chat.Message{
					{Text: "1"},
					{Text: "2"},
					{Text: "3"},
					{Text: "4"},
				},
			},
			args: args{
				n: 4,
			},
			want: []chat.Message{
				{Text: "1"},
				{Text: "2"},
				{Text: "3"},
				{Text: "4"},
			},
		},
		{
			name: "want 1 last message but got no messages,get no messages",
			fields: fields{
				messages: []chat.Message{},
			},
			args: args{
				n: 1,
			},
			want: []chat.Message{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := ChatRepository{
				messages: tt.fields.messages,
			}
			got := []chat.Message{}
			for m := range c.GetLastMessages(tt.args.n) {
				got = append(got, m)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetLastMessages() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestChatRepository_Put(t *testing.T) {
	type fields struct {
		messages []chat.Message
	}
	type args struct {
		msg chat.Message
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []chat.Message
	}{
		{
			name:   "put message",
			fields: fields{[]chat.Message{}},
			args:   args{chat.Message{Text: "1"}},
			want:   []chat.Message{{Text: "1"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &ChatRepository{
				messages: tt.fields.messages,
			}
			c.Put(tt.args.msg)
			if !reflect.DeepEqual(c.messages, tt.want) {
				t.Errorf("Put() = %v, want %v", c.messages, tt.want)
			}

		})
	}
}
