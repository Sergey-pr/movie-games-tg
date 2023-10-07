package forms

type LoginForm struct {
	InitData string `json:"init_data"`
}

type UserForm struct {
	TelegramId int    `json:"id"`
	Name       string `json:"first_name"`
	UserName   string `json:"username"`
	Language   string `json:"language_code"`
}

type UserLang struct {
	Language string `json:"language_code"`
}

type UserAnswer struct {
	CardId int `json:"card_id"`
	Points int `json:"points"`
}
