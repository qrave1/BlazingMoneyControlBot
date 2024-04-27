package main

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/qrave1/BlazingMoneyControlBot/internal/config"
	"github.com/qrave1/BlazingMoneyControlBot/internal/infrastructure/repositories"
	"log/slog"
	"os"
)

// todo fx di
// todo readme with scheme from drawio

//func main() {
//	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
//	defer cancel()
//
//	opts := []bot.Option{
//		bot.WithDebug(),
//		bot.WithDefaultHandler(defaultHandler),
//		bot.WithCallbackQueryDataHandler("button", bot.MatchTypePrefix, callbackHandler),
//	}
//
//	cfg := config.NewConfig()
//
//	b, err := bot.New(cfg.Telegram.TelegramToken, opts...)
//	if nil != err {
//		panic(err)
//	}
//
//	// Set default menu
//	b.SetMyCommands(ctx, &bot.SetMyCommandsParams{
//		Commands: []models.BotCommand{
//			{
//				Command:     "/start",
//				Description: "Старт",
//			},
//			{
//				Command:     "/help",
//				Description: "Помощь",
//			},
//		},
//		Scope:        &models.BotCommandScopeDefault{},
//		LanguageCode: "ru",
//	})
//
//	b.RegisterHandler(bot.HandlerTypeCallbackQueryData, "op", bot.MatchTypePrefix, opCallbackHandler)
//
//	_, err = b.SetWebhook(ctx, &bot.SetWebhookParams{
//		URL: cfg.Telegram.WebhookURL,
//	})
//	go http.ListenAndServe(":3000", b.WebhookHandler())
//	b.StartWebhook(ctx)
//}
//
//func callbackHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
//	// answering callback query first to let Telegram know that we received the callback query,
//	// and we're handling it. Otherwise, Telegram might retry sending the update repetitively
//	// as it thinks the callback query doesn't reach to our application. learn more by
//	// reading the footnote of the https://core.telegram.org/bots/api#callbackquery type.
//	b.AnswerCallbackQuery(ctx, &bot.AnswerCallbackQueryParams{
//		CallbackQueryID: update.CallbackQuery.ID,
//		ShowAlert:       false,
//	})
//	b.EditMessageText(ctx, &bot.EditMessageTextParams{
//		ChatID:    update.CallbackQuery.Message.Message.Chat.ID,
//		MessageID: update.CallbackQuery.Message.Message.ID,
//		Text:      "xxsm",
//		ReplyMarkup: &models.InlineKeyboardMarkup{
//			InlineKeyboard: [][]models.InlineKeyboardButton{
//				{
//					{Text: "Yes", CallbackData: "op_yes"},
//					{Text: "No", CallbackData: "op_no"},
//				},
//			},
//		},
//	})
//}
//
//func defaultHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
//	kb := &models.InlineKeyboardMarkup{
//		InlineKeyboard: [][]models.InlineKeyboardButton{
//			{
//				{Text: "Button 1", CallbackData: "button_1"},
//				{Text: "Button 2", CallbackData: "button_2"},
//			},
//			{
//				{Text: "Button 3", CallbackData: "button_3"},
//			},
//		},
//	}
//
//	b.SendMessage(ctx, &bot.SendMessageParams{
//		ChatID:      update.Message.Chat.ID,
//		Text:        "Click by button",
//		ReplyMarkup: kb,
//	})
//}
//
//func opCallbackHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
//	b.AnswerCallbackQuery(ctx, &bot.AnswerCallbackQueryParams{
//		CallbackQueryID: update.CallbackQuery.ID,
//		ShowAlert:       false,
//	})
//	b.EditMessageText(ctx, &bot.EditMessageTextParams{
//		ChatID:    update.CallbackQuery.Message.Message.Chat.ID,
//		MessageID: update.CallbackQuery.Message.Message.ID,
//		Text:      "success",
//		ReplyMarkup: &models.InlineKeyboardMarkup{
//			InlineKeyboard: [][]models.InlineKeyboardButton{},
//		},
//	})
//}

// todo this norm
//func main() {
//	app := &cli.App{
//		Name: "app",
//		Action: func(c *cli.Context) error {
//			//_, cleanup, err := factory.InitializeService()
//			//if err != nil {
//			//	log.Fatal(err)
//			//}
//			//defer cleanup()
//
//			sigCh := make(chan os.Signal, 1)
//			signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
//			<-sigCh
//
//			return nil
//		},
//		Commands: commands.Commands,
//	}
//
//	if err := app.Run(os.Args); err != nil {
//		panic(fmt.Errorf("error start app. %w", err))
//	}
//}

func main() {
	ctx := context.Background()

	cfg := config.NewConfig()

	log := slog.New(slog.NewTextHandler(os.Stdout, nil))

	db, err := sql.Open("postgres", cfg.DSN())
	if err != nil {
		panic(err)
	}

	wr := repositories.NewWalletPostgresRepository(db, log)

	err = wr.Create(ctx, 1, "cyxbebr1k", 1200)
	if err != nil {
		panic(err)
	}

	wallet, err := wr.Read(ctx, 1)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%#v\n", wallet)

	err = wr.UpdateBalance(ctx, 1, 1000)
	if err != nil {
		panic(err)
	}

	err = wr.Delete(ctx, 1)
	if err != nil {
		panic(err)
	}
}
