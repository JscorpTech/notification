package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/JscorpTech/notification/internal/domain"
)

type pmbSMSService struct {
	BaseURL string
}

// /broker-api/send
func NewPmbSMSService() domain.SMSServicePort {
	return &pmbSMSService{
		BaseURL: "https://send.smsxabar.uz",
	}
}

func (e *pmbSMSService) SendSMS(to, body string) error {
	client := http.Client{
		Timeout: 60 * time.Second,
	}
	payload := domain.PmbPayload{
		Messages: []domain.PmbMessage{
			{
				Recipient: "+998888112309",
				MessageID: "salomsdfs",
				Sms: domain.PmbSMS{
					Originator: "3600",
					Content: domain.PmbContent{
						Text: "salom",
					},
				},
			},
		},
	}
	var buf bytes.Buffer
	_ = json.NewEncoder(&buf).Encode(payload)
	req, _ := http.NewRequest("POST", e.BaseURL+"/broker-api/send", &buf)
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	var data map[string]interface{}
	json.NewDecoder(res.Body).Decode(&data)
	fmt.Print(data)
	return nil
}
