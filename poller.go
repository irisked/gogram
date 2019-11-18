package gogram

import 	(
	"time"
	"github.com/irisked/gogram/option"
	"github.com/irisked/gogram/types"
	"github.com/irisked/gogram/telegram"
)

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

func newPoller(tg *telegram.Telegram, limit int) poller {
	p := poller{
		config: pollingConfig{
			offset: -1,
			limit: limit,
			started: false,
		},
		tg: tg,
	}
	return p
}

func (p *poller) isPolling() bool {
	return p.config.started
}

func (p *poller) startPolling(updates chan types.Update) chan<- interface{} {
	if (!p.isPolling()) {
		p.config.started = true
		exit := make(chan interface{})
		go func () {
			for {
				select{
				case <- exit:
					p.config.started = false
					close(exit)
					return
				default:
					p.fetch(updates)
					time.Sleep(2 * time.Second)
				}
			}
		}()
		return exit
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