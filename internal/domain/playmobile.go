package domain

type PmbContent struct {
	Text string `json:"text"`
}

type PmbSMS struct {
	Originator string     `json:"originator"`
	Content    PmbContent `json:"content"`
}

type PmbMessage struct {
	Recipient string `json:"recipient"`
	MessageID string `json:"message-id"`
	Sms       PmbSMS `json:"sms"`
}

type PmbPayload struct {
	Messages []PmbMessage
}
