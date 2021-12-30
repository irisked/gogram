package gogram

import 	(
	"time"
	"gogram/option"
	"gogram/types"
	"gogram/telegram"
)

// Poller interface.
type Poller interface {
	StartPolling(chan<- types.Update) chan<- interface{}
	IsPolling() bool
}

type pollingConfig struct {
	offset int
	limit int
	allowedUpdates []types.UpdateType
	started bool
}

type poller struct {
	config pollingConfig
	tg *telegram.Telegram
}
// NewPoller creates poller.
func NewPoller(tg *telegram.Telegram, limit int) Poller {
	p := poller{
		config: pollingConfig{
			offset: -1,
			limit: limit,
			started: false,
		},
		tg: tg,
	}
	return &p
}

func (p *poller) IsPolling() bool {
	return p.config.started
}

func (p *poller) StartPolling(updates chan<- types.Update) chan<- interface{} {
	return p.startPolling(updates)
}

func (p *poller) SetPolling(isPolling bool) {
	p.config.started = isPolling
}

func (p *poller) startPolling(updates chan<- types.Update) chan<- interface{} {
	if (!p.IsPolling()) {
		p.SetPolling(true)
		done := make(chan interface{})
		go func () {
			for {
				select{
				case <-done:
					p.SetPolling(false)
					return
				default:
				}
				p.fetch(updates)
				time.Sleep(2 * time.Second)
			}
		}()
		return done
	}
	panic("(p *poller) startPolling called more the one time.")
}

func (p *poller) fetch(updates chan<- types.Update) {
	incoming, err := p.tg.GetUpdates(
		option.Offset(p.config.offset),
		option.Limit(p.config.limit),
	)
	if err != nil {
		panic(err)
	}
	if len(incoming) > 0 {
		p.config.offset = incoming[len(incoming) - 1].UpdateID + 1
		for _, update := range incoming {
			updates <- update
		}
	}
}