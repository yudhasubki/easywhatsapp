package easywhatsapp

import (
	"github.com/pusher/pusher-http-go"
)

type PusherClient struct {
	Client    pusher.Client
	EventName string
	Channel   string
}

func (w *EasyWhatsapp) RenderQRCodeHTMLPusher(qrCode string) {
	if w.Streamer.Pusher.Channel != "" && w.Streamer.Pusher.EventName != "" {
		w.Streamer.Pusher.Client.Trigger(w.Streamer.Pusher.Channel, w.Streamer.Pusher.EventName, qrCode)
	}
}
