package services

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/JscorpTech/notification/internal/domain"
	"github.com/JscorpTech/notification/internal/redis"
	"github.com/k0kubun/pp/v3"
)

type eskizSMSService struct {
	BaseURL string
	Ctx     context.Context
}

// /broker-api/send
func NewEskizSMSService(ctx context.Context) domain.SMSServicePort {
	return &eskizSMSService{
		Ctx:     ctx,
		BaseURL: os.Getenv("ESKIZ_DOMAIN"),
	}
}

func (e *eskizSMSService) Request(payload any, path string, isAuth bool, retry bool) (*http.Response, error) {
	var buf bytes.Buffer
	_ = json.NewEncoder(&buf).Encode(payload)
	client := http.Client{
		Timeout: 60 * time.Second,
	}
	req, err := http.NewRequest("POST", e.BaseURL+path, &buf)
	req.Header.Add("Content-Type", "application/json")
	if isAuth {
		req.Header.Add("Authorization", "Bearer "+e.GetToken(true, true))
	}

	if err != nil {
		return nil, err
	}
	res, err := client.Do(req)
	if res.StatusCode == http.StatusUnauthorized && retry {
		time.Sleep(time.Second * 5)
		pp.Print("Qayta urunish")
		e.GetToken(false, false)
		return e.Request(payload, path, isAuth, false)
	}
	return res, err
}

func (e *eskizSMSService) GetToken(cache bool, retry bool) string {
	email := os.Getenv("ESKIZ_USER")
	password := os.Getenv("ESKIZ_PASSWORD")

	if email == "" || password == "" {
		log.Fatal("password or fmail not found")
	}

	token, err := redis.RDB.Get(e.Ctx, "eskiz_token").Result()
	if err == nil && cache {
		pp.Print("Eskiz token topildi 😁")
		return token
	}
	payload := domain.EskizLogin{
		Email:    email,
		Password: password,
	}
	res, err := e.Request(payload, "/auth/login", false, retry)
	if err != nil {
		pp.Print(err.Error())
	}
	var data domain.EskizLoginRes
	_ = json.NewDecoder(res.Body).Decode(&data)
	token = data.Data.Token
	redis.RDB.Set(e.Ctx, "eskiz_token", token, 30*24*time.Hour)
	pp.Print("Eskiz yangi token olindi 😔")
	return token
}

func (e *eskizSMSService) SendSMS(to, body string) error {
	payload := domain.EskizMessage{
		Phone:   to,
		Message: body,
		From:    os.Getenv("ESKIZ_FROM"),
	}
	res, err := e.Request(payload, "/message/sms/send", true, true)
	if err != nil {
		return err
	}
	var data domain.EskizMessageRes
	json.NewDecoder(res.Body).Decode(&data)
	pp.Print(data)
	return nil
}
