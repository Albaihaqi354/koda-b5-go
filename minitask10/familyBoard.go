package minitask10

import "fmt"

type Message struct {
	Sender string
	Text   string
}

func NewMessage(sender, text string) Message {
	return Message{
		Sender: sender,
		Text:   text,
	}
}

func Blackboard(mc chan Message) {
	for msg := range mc {
		fmt.Printf("%s: %s\n", msg.Sender, msg.Text)
	}
}

func MsgSender(mc chan Message, msg Message) {
	mc <- msg
}
