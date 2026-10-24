package gogram

import (
	"crypto/subtle"
	"encoding/json"
	"net/http"

	"github.com/irisked/gogram/telegram"
)

// maxWebhookBody bounds the size of an update accepted by WebhookHandler.
const maxWebhookBody = 1 << 20

// WebhookHandler returns an http.Handler that receives updates sent by
// Telegram to a webhook (see telegram.Client.SetWebhook) and handles
// them before responding. Handler errors go to the ErrorHandler; the
// response is a success either way, since Telegram would otherwise send
// the update again. On its first update, the handler looks up the bot's
// username with Init unless WithUsername was given.
//
// If secretToken is not empty, requests must carry it in the
// X-Telegram-Bot-Api-Secret-Token header, as set with
// telegram.WithSecretToken; others are rejected.
//
//	bot.Client().SetWebhook(ctx, "https://example.com/bot",
//		telegram.WithSecretToken(secret))
//	http.Handle("/bot", bot.WebhookHandler(secret))
func (b *Bot) WebhookHandler(secretToken string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.Header().Set("Allow", http.MethodPost)
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}
		if secretToken != "" {
			got := r.Header.Get("X-Telegram-Bot-Api-Secret-Token")
			if subtle.ConstantTimeCompare([]byte(got), []byte(secretToken)) != 1 {
				http.Error(w, "forbidden", http.StatusForbidden)
				return
			}
		}
		var update telegram.Update
		if err := json.NewDecoder(http.MaxBytesReader(w, r.Body, maxWebhookBody)).Decode(&update); err != nil {
			http.Error(w, "bad request", http.StatusBadRequest)
			return
		}
		if b.Username() == "" {
			if err := b.Init(r.Context()); err != nil {
				b.logger.WarnContext(r.Context(), "gogram: getMe failed; accepting commands for any bot", "error", err)
			}
		}
		b.handle(r.Context(), update)
		// Telegram retries updates answered with an error, so a failed
		// handler still gets 200: the ErrorHandler has seen the error.
		w.WriteHeader(http.StatusOK)
	})
}
