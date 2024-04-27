package tg_bot

import (
	"context"
	"fmt"
	"github.com/go-telegram/bot"
	"github.com/qrave1/BlazingMoneyControlBot/internal/config"
	"net/http"
	"time"
)

type Server struct {
	b *bot.Botu
}

// todo add controllers and handlers
func NewServer(cfg *config.Config) (*Server, error) {
	b, err := bot.New(cfg.Telegram.TelegramToken)
	if err != nil {
		return nil, fmt.Errorf("error creating tg bot: %v", err)
	}

	return &Server{b: b}, nil
}

func (s *Server) Run(ctx context.Context) {
	errCh := make(chan error)
	go func() {
		errCh <- http.ListenAndServe(":3000", s.b.WebhookHandler())
	}()
	select {
	case <-time.After(100 * time.Millisecond):
		break
	case err := <-errCh:
		panic(fmt.Errorf("error start bot via webhooks. %w", err))
	}

	s.b.StartWebhook(ctx)
}
