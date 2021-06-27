package easywhatsapp

import (
	"log"

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
		err := w.Streamer.Pusher.Client.Trigger(w.Streamer.Pusher.Channel, w.Streamer.Pusher.EventName, qrCode)
		if err != nil {
			log.Printf("error streamer : %v", err.Error())
		}
	}
}

func (w *EasyWhatsapp) RenderQRCodeConsole(qrCode string) {
	if w.Streamer.Console.EnableConsole {
		obj := qrcodeTerminal.New()
		obj.Get(qrCode).Print()
	}
}
