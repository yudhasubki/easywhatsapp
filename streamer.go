package easywhatsapp

import (
	qrcodeTerminal "github.com/Baozisoftware/qrcode-terminal-go"
	"github.com/pusher/pusher-http-go"
)

type PusherClient struct {
	Client    pusher.Client
	EventName string
	Channel   string
}

type ConsoleClient struct {
	EnableConsole bool
}

func (w *EasyWhatsapp) RenderQRCodeHTMLPusher(qrCode string) {
	if w.Streamer.Pusher.Channel != "" && w.Streamer.Pusher.EventName != "" {
		w.Streamer.Pusher.Client.Trigger(w.Streamer.Pusher.Channel, w.Streamer.Pusher.EventName, qrCode)
	}
}

func (w *EasyWhatsapp) RenderQRCodeConsole(qrCode string) {
	if w.Streamer.Console.EnableConsole {
		obj := qrcodeTerminal.New()
		obj.Get(qrCode).Print()
	}
}
