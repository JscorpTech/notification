package domain

type EmailServicePort interface {
	SendMail([]string, []byte)
}
