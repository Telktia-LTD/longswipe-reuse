package telegramx

import (
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type TelegramService struct {
	botToken string
	bot      *tgbotapi.BotAPI
}

func NewTelegramService(telegramBotToken string) *TelegramService {
	botToken := os.Getenv(telegramBotToken)
	return &TelegramService{
		botToken: botToken,
	}
}

func (s *TelegramService) InitBot() error {
	bot, err := tgbotapi.NewBotAPI(s.botToken)
	if err != nil {
		return err
	}
	bot.Debug = true
	s.bot = bot
	log.Printf("Authorized on account %s", bot.Self.UserName)
	return nil
}

func (s *TelegramService) HandleUpdates() {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := s.bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil { // ignore any non-Message updates
			continue
		}

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		// Respond to the message using the SendMessage method
		responseText := "You said: " + update.Message.Text
		if err := s.SendMessage(update.Message.Chat.ID, responseText); err != nil {
			log.Printf("Failed to send message: %v", err)
		}
	}
}

func (s *TelegramService) SendMessage(chatID int64, text string) error {
	msg := tgbotapi.NewMessage(chatID, text)
	_, err := s.bot.Send(msg)
	if err != nil {
		log.Printf("Failed to send message: %v", err)
		return err
	}
	return nil
}
