package biz

type Message struct {
	Name string
}

type MessageRepo interface {
	GetMore(string) string
}

type MessageBiz struct {
	msgRepo MessageRepo
}

func NewMessageBiz(msgRepo MessageRepo) *MessageBiz {
	return &MessageBiz{msgRepo: msgRepo}
}

func (msgBiz *MessageBiz) GetMore(msg *Message) string {
	return msgBiz.msgRepo.GetMore(msg.Name)
}
