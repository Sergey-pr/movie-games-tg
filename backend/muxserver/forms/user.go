package forms

type LoginForm struct {
	Hash string   `json:"hash" validate:"required"`
	User UserForm `json:"user"`
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
