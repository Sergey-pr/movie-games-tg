package models

import (
	"context"
	"fmt"
	"github.com/Sergey-pr/movie-games-tg/muxserver/forms"
	"github.com/Sergey-pr/movie-games-tg/persist"
	"github.com/Sergey-pr/movie-games-tg/utils"
	"github.com/doug-martin/goqu/v9"
	"github.com/ssoroka/slice"
	"strings"
)

const (
	CardProcessorsTableName = "card_processors"
)

var NegativeAnswersSlice = []string{"NO", "No", "no", "НЕТ", "Нет", "нет"}

type CardProcessor struct {
	Id     int   `db:"id" goqu:"skipupdate,skipinsert"`
	UserId int   `db:"user_id"`
	CardId *int  `db:"card_id"`
	State  int   `db:"state"`
	User   *User `db:"-"`
	Card   *Card `db:"-"`
	ChatId int   `db:"-"`
}

func (obj *CardProcessor) ProcessMsg(ctx context.Context, form *forms.BotUpdate) error {
	if obj.State != 0 && form.Message.Text != "/stop" {
		err := obj.addCard(ctx, form)
		if err != nil {
			return err
		}
		return nil
	}
	switch form.Message.Text {
	case "/add":
		err := obj.addCard(ctx, form)
		if err != nil {
			return err
		}
	case "/stop":
		err := obj.sendStopMessage()
		if err != nil {
			return err
		}
		obj.State = 0
		err = obj.Save(ctx)
		if err != nil {
			return err
		}
	default:
		err := obj.sendDefaultMessage()
		if err != nil {
			return err
		}
	}
	return nil
}

func (obj *CardProcessor) addCard(ctx context.Context, form *forms.BotUpdate) error {
	var err error
	switch obj.State {
	case 0:
		err = obj.processAddCard(ctx)
		if err != nil {
			return err
		}
	case 1:
		err = obj.processAddEnglishName(ctx, form.Message.Text)
		if err != nil {
			return err
		}
	case 2:
		err = obj.processAddRussianName(ctx, form.Message.Text)
		if err != nil {
			return err
		}
	case 3:
		err = obj.processAddEnglishDesc(ctx, form.Message.Text)
		if err != nil {
			return err
		}
	case 4:
		err = obj.processAddRussianDesc(ctx, form.Message.Text)
		if err != nil {
			return err
		}
	case 5:
		err = obj.processAddEnglishQuote(ctx, form.Message.Text)
		if err != nil {
			return err
		}
	case 6:
		err = obj.processAddRussianQuote(ctx, form.Message.Text)
		if err != nil {
			return err
		}
	case 7:
		err = obj.processAddEnglishAnswers(ctx, form.Message.Text)
		if err != nil {
			return err
		}
	case 8:
		err = obj.processAddRussianAnswers(ctx, form.Message.Text)
		if err != nil {
			return err
		}
	case 9:
		err = obj.processAddEnglishFacts(ctx, form.Message.Text)
		if err != nil {
			return err
		}
	case 10:
		err = obj.processAddRussianFacts(ctx, form.Message.Text)
		if err != nil {
			return err
		}
	case 11:
		err = obj.processAddDrawing(ctx, form)
		if err != nil {
			return err
		}
	case 12:
		err = obj.processAddPixelated(ctx, form)
		if err != nil {
			return err
		}
	case 13:
		err = obj.processAddScreenshot(ctx, form)
		if err != nil {
			return err
		}
	case 14:
		err = obj.processAddBackgroundColor1(ctx, form.Message.Text)
		if err != nil {
			return err
		}
	case 15:
		err = obj.processAddBackgroundColor2(ctx, form.Message.Text)
		if err != nil {
			return err
		}
	case 16:
		err = obj.processAddTextColor(ctx, form.Message.Text)
		if err != nil {
			return err
		}
	default:
		err = obj.processError()
		if err != nil {
			return err
		}
	}
	return nil
}

func (obj *CardProcessor) processAddTextColor(ctx context.Context, msg string) error {
	var answer string
	if obj.User.Language == "ru" {
		answer = fmt.Sprintf("Цвет: %s\n\nКарточка готова!", msg)
	} else {
		answer = fmt.Sprintf("Color: %s\n\nThe card is ready!", msg)
	}
	err := utils.SendBotMessage(obj.ChatId, answer)
	if err != nil {
		return err
	}

	obj.Card.TextColor = msg
	obj.Card.Completed = true
	err = obj.Card.Save(ctx)
	if err != nil {
		return err
	}

	obj.State = 0
	obj.Card = nil
	obj.CardId = nil
	err = obj.Save(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (obj *CardProcessor) processAddBackgroundColor2(ctx context.Context, msg string) error {
	var answer string
	if obj.User.Language == "ru" {
		answer = fmt.Sprintf("Цвет: %s\n\nТеперь добавим цвет текста!", msg)
	} else {
		answer = fmt.Sprintf("Color: %s\n\nNow let's add the text color!", msg)
	}
	err := utils.SendBotMessage(obj.ChatId, answer)
	if err != nil {
		return err
	}

	obj.Card.BackgroundColor2 = msg
	err = obj.Card.Save(ctx)
	if err != nil {
		return err
	}

	obj.State = 16
	err = obj.Save(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (obj *CardProcessor) processAddBackgroundColor1(ctx context.Context, msg string) error {
	var answer string
	if obj.User.Language == "ru" {
		answer = fmt.Sprintf("Цвет: %s\n\nТеперь введите второй цвет градиента.", msg)
	} else {
		answer = fmt.Sprintf("Color: %s\n\nNow type the second background color.", msg)
	}
	err := utils.SendBotMessage(obj.ChatId, answer)
	if err != nil {
		return err
	}

	obj.Card.BackgroundColor1 = msg
	err = obj.Card.Save(ctx)
	if err != nil {
		return err
	}

	obj.State = 15
	err = obj.Save(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (obj *CardProcessor) processAddScreenshot(ctx context.Context, form *forms.BotUpdate) error {
	var imageId string
	if form.Message.Document.FileId != "" {
		imageId = form.Message.Document.FileId
	} else if len(form.Message.Photo) > 0 {
		imageId = form.Message.Photo[0].FileId
	} else {
		var answer string
		if obj.User.Language == "ru" {
			answer = "Приложите изображение к сообщению, желательно как файл."
		} else {
			answer = "Please attach image to the message as a file."
		}
		err := utils.SendBotMessage(obj.ChatId, answer)
		if err != nil {
			return err
		}
		return nil
	}
	var answer string
	if obj.User.Language == "ru" {
		answer = "Изображение сохранено.\n\nТеперь напишите первый цвет для градиента фона. (#FF00FF)"
	} else {
		answer = "Image saved.\n\nNow let's add a first color for background gradient. (#FF00FF)"
	}
	err := utils.SendBotMessage(obj.ChatId, answer)
	if err != nil {
		return err
	}

	obj.Card.ScreenshotId = imageId
	err = obj.Card.Save(ctx)
	if err != nil {
		return err
	}

	err = utils.DownloadBotImage(imageId)
	if err != nil {
		return err
	}

	obj.State = 14
	err = obj.Save(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (obj *CardProcessor) processAddPixelated(ctx context.Context, form *forms.BotUpdate) error {
	var imageId string
	if form.Message.Document.FileId != "" {
		imageId = form.Message.Document.FileId
	} else if len(form.Message.Photo) > 0 {
		imageId = form.Message.Photo[0].FileId
	} else {
		var answer string
		if obj.User.Language == "ru" {
			answer = "Приложите изображение к сообщению, желательно как файл."
		} else {
			answer = "Please attach image to the message as a file."
		}
		err := utils.SendBotMessage(obj.ChatId, answer)
		if err != nil {
			return err
		}
		return nil
	}
	var answer string
	if obj.User.Language == "ru" {
		answer = "Изображение сохранено.\n\nТеперь добавьте кадр из фильма."
	} else {
		answer = "Image saved.\n\nNow let's add a screenshot from movie."
	}
	err := utils.SendBotMessage(obj.ChatId, answer)
	if err != nil {
		return err
	}

	obj.Card.PixelatedId = imageId
	err = obj.Card.Save(ctx)
	if err != nil {
		return err
	}

	err = utils.DownloadBotImage(imageId)
	if err != nil {
		return err
	}

	obj.State = 13
	err = obj.Save(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (obj *CardProcessor) processAddDrawing(ctx context.Context, form *forms.BotUpdate) error {
	var imageId string
	if form.Message.Document.FileId != "" {
		imageId = form.Message.Document.FileId
	} else if len(form.Message.Photo) > 0 {
		imageId = form.Message.Photo[0].FileId
	} else {
		var answer string
		if obj.User.Language == "ru" {
			answer = "Приложите изображение к сообщению, желательно как файл."
		} else {
			answer = "Please attach image to the message as a file."
		}
		err := utils.SendBotMessage(obj.ChatId, answer)
		if err != nil {
			return err
		}
		return nil
	}
	var answer string
	if obj.User.Language == "ru" {
		answer = "Изображение сохранено.\n\nТеперь добавьте пикселизированное изображение."
	} else {
		answer = "Image saved.\n\nNow let's add a pixelated image."
	}
	err := utils.SendBotMessage(obj.ChatId, answer)
	if err != nil {
		return err
	}

	obj.Card.DrawingId = imageId
	err = obj.Card.Save(ctx)
	if err != nil {
		return err
	}

	err = utils.DownloadBotImage(imageId)
	if err != nil {
		return err
	}

	obj.State = 12
	err = obj.Save(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (obj *CardProcessor) processAddRussianFacts(ctx context.Context, msg string) error {
	var answer string
	if slice.Contains(NegativeAnswersSlice, msg) {
		if obj.User.Language == "ru" {
			answer = "Теперь добавим рисунок с прозрачным фоном. Приложите его к сообщению, желательно как файл."
		} else {
			answer = "Now let's add drawing with transparent background. Please attach it to the message as a file."
		}
	} else {
		facts := strings.Split(msg, ";")
		if obj.User.Language == "ru" {
			answer = fmt.Sprintf("Русские факты:\n%s\n\nТеперь добавим изображение рисунка. Оно должно быть с прозрачным фоном.", strings.Join(facts, "\n"))
		} else {
			answer = fmt.Sprintf("Russian facts are:\n%s\n\nNow let's add a drawing image. It should be with a transparent background.", strings.Join(facts, "\n"))
		}
		obj.Card.FactsRu = utils.ToGenericArray(facts)
		err := obj.Card.Save(ctx)
		if err != nil {
			return err
		}
	}

	err := utils.SendBotMessage(obj.ChatId, answer)
	if err != nil {
		return err
	}

	obj.State = 11
	err = obj.Save(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (obj *CardProcessor) processAddEnglishFacts(ctx context.Context, msg string) error {
	facts := strings.Split(msg, ";")

	var answer string
	if obj.User.Language == "ru" {
		answer = fmt.Sprintf("Факты:\n%s\n\nДобавим факты на русском?\n\nНапишите \"Нет\" если не нужно, или напишите факты на русском.", strings.Join(facts, "\n"))
	} else {
		answer = fmt.Sprintf("Intereting facts:\n%s\n\nDo you want to add facts in russian?\n\nAnswer \"No\" if you don't need russian answers or type russian facts.", strings.Join(facts, "\n"))
	}
	err := utils.SendBotMessage(obj.ChatId, answer)
	if err != nil {
		return err
	}

	obj.Card.FactsEn = utils.ToGenericArray(facts)
	err = obj.Card.Save(ctx)
	if err != nil {
		return err
	}

	obj.State = 10
	err = obj.Save(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (obj *CardProcessor) processAddRussianAnswers(ctx context.Context, msg string) error {
	var answer string
	if slice.Contains(NegativeAnswersSlice, msg) {
		if obj.User.Language == "ru" {
			answer = "Теперь добавим факты о фильме. Их нужно перечислить через \";\"."
		} else {
			answer = "Now let's add interesting facts about the movie. Type them with \";\" as delimiter."
		}
	} else {
		answers := strings.Split(msg, ",")
		if len(answers) != 5 {
			var errMsg string
			if obj.User.Language == "ru" {
				errMsg = fmt.Sprintf("Нужно указать 5 вариантов через запятую, вы указали %d.", len(answers))
			} else {
				errMsg = fmt.Sprintf("You need to write 5 variants with comma between them. You've written %d variants.", len(answers))
			}
			err := utils.SendBotMessage(obj.ChatId, errMsg)
			if err != nil {
				return err
			}
		}
		if obj.User.Language == "ru" {
			answer = fmt.Sprintf("Русские варианты ответов:\n%s\n\nТеперь добавим интересные факты на английском. Их нужно перечислить через \";\".", strings.Join(answers, "\n"))
		} else {
			answer = fmt.Sprintf("Russian answers are:\n%s\n\nNow let's add interesting facts. Type them with \";\" as delimiter.", strings.Join(answers, "\n"))
		}
		obj.Card.AnswersRu = utils.ToGenericArray(answers)
		err := obj.Card.Save(ctx)
		if err != nil {
			return err
		}
	}

	err := utils.SendBotMessage(obj.ChatId, answer)
	if err != nil {
		return err
	}

	obj.State = 9
	err = obj.Save(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (obj *CardProcessor) processAddEnglishAnswers(ctx context.Context, msg string) error {
	answers := strings.Split(msg, ",")
	if len(answers) != 5 {
		var errMsg string
		if obj.User.Language == "ru" {
			errMsg = fmt.Sprintf("Нужно указать 5 вариантов через запятую, вы указали %d.", len(answers))
		} else {
			errMsg = fmt.Sprintf("You need to write 5 variants with comma between them. You've written %d variants.", len(answers))
		}
		err := utils.SendBotMessage(obj.ChatId, errMsg)
		if err != nil {
			return err
		}
	}

	var answer string
	if obj.User.Language == "ru" {
		answer = fmt.Sprintf("Варианты ответов:\n%s\n\nДобавим ответы на русском?\n\nНапишите \"Нет\" если не нужно, или напишите ответы на русском.", strings.Join(answers, "\n"))
	} else {
		answer = fmt.Sprintf("Answer variants:\n%s\n\nDo you want to add answers in russian?\n\nAnswer \"No\" if you don't need russian answers or type russian answers.", strings.Join(answers, "\n"))
	}
	err := utils.SendBotMessage(obj.ChatId, answer)
	if err != nil {
		return err
	}

	obj.Card.AnswersEn = utils.ToGenericArray(answers)
	err = obj.Card.Save(ctx)
	if err != nil {
		return err
	}

	obj.State = 8
	err = obj.Save(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (obj *CardProcessor) processAddRussianQuote(ctx context.Context, msg string) error {
	var answer string
	if slice.Contains(NegativeAnswersSlice, msg) {
		if obj.User.Language == "ru" {
			answer = "Теперь добавим варианты ответов. Нужно указать 5 вариантов через запятую. Название фильма должно быть среди них."
		} else {
			answer = "Now let's add possible answers to the card. You need to write 5 variants with comma as delimiter. One of the answers should be the Movie's name."
		}
	} else {
		if obj.User.Language == "ru" {
			answer = fmt.Sprintf("Русская цитата: %s\n\nТеперь добавим варианты ответов. Нужно указать 5 вариантов через запятую. Название фильма должно быть среди них.", msg)
		} else {
			answer = fmt.Sprintf("The Russian quote is: %s\n\nNow let's add possible answers to the card. You need to write 5 variants with comma as delimiter. One of the answers should be the Movie's name.", msg)
		}
		obj.Card.QuoteRu = msg
		err := obj.Card.Save(ctx)
		if err != nil {
			return err
		}
	}

	err := utils.SendBotMessage(obj.ChatId, answer)
	if err != nil {
		return err
	}

	obj.State = 7
	err = obj.Save(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (obj *CardProcessor) processAddEnglishQuote(ctx context.Context, msg string) error {
	var answer string
	if obj.User.Language == "ru" {
		answer = fmt.Sprintf("Цитата: %s\n\nДобавим цитату на русском?\n\nНапишите \"Нет\" если не нужно, или напишите цитату на русском.", msg)
	} else {
		answer = fmt.Sprintf("The quote is: %s\n\nDo you want to add quote in russian?\n\nAnswer \"No\" if you don't need russian quote or answer with the russian quote.", msg)
	}
	err := utils.SendBotMessage(obj.ChatId, answer)
	if err != nil {
		return err
	}

	obj.Card.QuoteEn = msg
	err = obj.Card.Save(ctx)
	if err != nil {
		return err
	}

	obj.State = 6
	err = obj.Save(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (obj *CardProcessor) processAddRussianDesc(ctx context.Context, msg string) error {
	var answer string
	if slice.Contains(NegativeAnswersSlice, msg) {
		if obj.User.Language == "ru" {
			answer = "Теперь добавим цитату на ангийском"
		} else {
			answer = "Now let's add a quote from the movie in english"
		}
	} else {
		if obj.User.Language == "ru" {
			answer = fmt.Sprintf("Русское описание: %s\n\nТеперь добавим цитату на английском.", msg)
		} else {
			answer = fmt.Sprintf("The Russian desc is: %s\n\nNow let's add a quote from the movie in english.", msg)
		}
		obj.Card.DescRu = msg
		err := obj.Card.Save(ctx)
		if err != nil {
			return err
		}
	}

	err := utils.SendBotMessage(obj.ChatId, answer)
	if err != nil {
		return err
	}

	obj.State = 5
	err = obj.Save(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (obj *CardProcessor) processAddEnglishDesc(ctx context.Context, msg string) error {
	var answer string
	if obj.User.Language == "ru" {
		answer = fmt.Sprintf("Описание: %s\n\nДобавим русское описание?\n\nНапишите \"Нет\" если не нужно, или напишите русское описание фильма.", msg)
	} else {
		answer = fmt.Sprintf("The description is: %s\n\nDo you want to add russian movie description?\n\nAnswer \"No\" if you don't need russian movie description or answer with the russian movie description.", msg)
	}
	err := utils.SendBotMessage(obj.ChatId, answer)
	if err != nil {
		return err
	}

	obj.Card.DescEn = msg
	err = obj.Card.Save(ctx)
	if err != nil {
		return err
	}

	obj.State = 4
	err = obj.Save(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (obj *CardProcessor) processAddRussianName(ctx context.Context, msg string) error {
	var answer string
	if slice.Contains(NegativeAnswersSlice, msg) {
		if obj.User.Language == "ru" {
			answer = "Теперь добавим описание на английском"
		} else {
			answer = "Now let's add the movie's description in english"
		}
	} else {
		if obj.User.Language == "ru" {
			answer = fmt.Sprintf("Русское название: %s\n\nТеперь добавим описание на английском.", msg)
		} else {
			answer = fmt.Sprintf("The Russian name is: %s\n\nNow let's add the movie's description in english.", msg)
		}
		obj.Card.NameRu = msg
		err := obj.Card.Save(ctx)
		if err != nil {
			return err
		}
	}

	err := utils.SendBotMessage(obj.ChatId, answer)
	if err != nil {
		return err
	}

	obj.State = 3
	err = obj.Save(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (obj *CardProcessor) processAddEnglishName(ctx context.Context, msg string) error {
	var answer string
	if obj.User.Language == "ru" {
		answer = fmt.Sprintf("Название: %s\n\nДобавим русское название?\n\nНапишите \"Нет\" если не нужно, или напишите русское название фильма.", msg)
	} else {
		answer = fmt.Sprintf("The name is: %s\n\nDo you want to add russian movie name?\n\nAnswer \"No\" if you don't need russian movie name or answer with the russian movie name.", msg)
	}
	err := utils.SendBotMessage(obj.ChatId, answer)
	if err != nil {
		return err
	}

	if obj.Card == nil {
		card := &Card{
			Category: "movie",
			NameEn:   msg,
		}
		obj.Card = card
		err = obj.Card.Save(ctx)
		if err != nil {
			return err
		}
	}

	obj.CardId = &obj.Card.Id
	obj.State = 2
	err = obj.Save(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (obj *CardProcessor) processAddCard(ctx context.Context) error {
	var answer string
	if obj.User.Language == "ru" {
		answer = "Давай создадим новую карточку, если хочешь прерваться напиши /stop\n\nКакое английское название фильма?"
	} else {
		answer = "Let's create a new card, to stop send /stop\n\nWhat's the card's english movie name?"
	}
	err := utils.SendBotMessage(obj.ChatId, answer)
	if err != nil {
		return err
	}

	obj.State = 1
	err = obj.Save(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (obj *CardProcessor) sendStopMessage() error {
	var answer string
	if obj.User.Language == "ru" {
		answer = "Создание карточки остановленно"
	} else {
		answer = "Creating of a new card is stopped"
	}
	err := utils.SendBotMessage(obj.ChatId, answer)
	if err != nil {
		return err
	}
	return nil
}

func (obj *CardProcessor) sendDefaultMessage() error {
	var answer string
	if obj.User.Language == "ru" {
		answer = "Нажмите Start чтобы запустить игру"
	} else {
		answer = "Click \"Start\" button to open the game"
	}
	err := utils.SendBotMessage(obj.ChatId, answer)
	if err != nil {
		return err
	}
	return nil
}

func (obj *CardProcessor) processError() error {
	var answer string
	if obj.User.Language == "ru" {
		answer = "Невозможно понять ответ"
	} else {
		answer = "Can't understand answer"
	}
	err := utils.SendBotMessage(obj.ChatId, answer)
	if err != nil {
		return err
	}
	return nil
}

// Save card instance in DB
func (obj *CardProcessor) Save(ctx context.Context) error {
	var err error
	if obj.Id == 0 {
		err = obj.create(ctx)
	} else {
		err = obj.update(ctx)
	}
	if err != nil {
		return err
	}

	return nil
}

// createCard private method for create new cards DB record
func (obj *CardProcessor) create(ctx context.Context) error {
	insert := persist.Db.Insert(CardProcessorsTableName).
		Rows(obj).
		Returning("*").Executor()

	if _, err := insert.ScanStructContext(ctx, obj); err != nil {
		return err
	}
	return nil
}

// updateCard private method for update card record in DB
func (obj *CardProcessor) update(ctx context.Context) error {
	update := persist.Db.From(CardProcessorsTableName).
		Where(goqu.C("id").Eq(obj.Id)).Update().Set(obj).
		Executor()
	_, err := update.ExecContext(ctx)
	if err != nil {
		return err
	}
	return nil
}

// Delete delete card from DB
func (obj *CardProcessor) Delete(ctx context.Context) error {
	_, err := persist.Db.From(CardProcessorsTableName).
		Where(goqu.Ex{"id": obj.Id}).
		Delete().
		Executor().ExecContext(ctx)
	if err != nil {
		return err
	}

	return nil
}
