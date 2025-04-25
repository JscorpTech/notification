package domain

type NotificationConsumerPort interface {
	Start()
	Handler(NotificationMsg)
}

type SMSServicePort interface {
	SendSMS(string, string) error
}
type NotifierPort interface {
	SendMessage([]string, string)
}

type NotificationMsg struct {
	Type    string   `json:"type"`
	Message string   `json:"message"`
	To      []string `json:"to"`
}
