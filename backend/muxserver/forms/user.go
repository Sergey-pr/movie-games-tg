package forms

type UserForm struct {
	TelegramId int    `json:"telegram_id" validate:"required"`
	Name       string `json:"name" validate:"required"`
	Language   string `json:"language" validate:"required"`
}
