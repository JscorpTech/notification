package domain

type EskizLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type EskizLoginRes struct {
	Message   string
	TokenType string
	Data      struct {
		Token string
	}
}

type EskizMessage struct {
	Phone       string `json:"mobile_phone"`
	Message     string `json:"message"`
	From        string `json:"from"`
	CallbackURL string `json:"callback_url"`
}

type EskizMessageRes struct {
	ID      string `json:"id"`
	Message string `json:"message"`
	Status  string `json:"status"`
}
