package types

type IMessage interface {
	IMessageHeader
	MsgBody() IMessageBody
	ToRlp() IRlpMessage
	Check() error
}

type IMessageIndex interface {
	GetHeight() uint64
}

type IWorks interface {
	GetCycle() uint64
	GetEndTime() uint64
	GetWorkLoad() uint64
}
