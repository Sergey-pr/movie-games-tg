package forms

type LoginForm struct {
	InitData string `json:"init_data"`
}

type UserForm struct {
	TelegramId int    `json:"id" validate:"required"`
	Name       string `json:"first_name" validate:"required"`
	UserName   string `json:"username"`
	Language   string `json:"language_code" validate:"required"`
}

type UserLang struct {
	Language string `json:"language_code" validate:"required"`
}

type UserAnswer struct {
	CardId int `json:"card_id" validate:"required"`
	Points int `json:"points" validate:"required"`
}
