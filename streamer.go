package easywhatsapp

import (
	"fmt"
	"log"

	"github.com/pusher/pusher-http-go"

	"github.com/nsqio/go-nsq"
)

type PusherClient struct {
	Client    pusher.Client
	EventName string
	Channel   string
}

type NSQClient struct {
	Host    string
	Port    string
	Topic   string
	Channel string
}

func (w *EasyWhatsapp) RenderQRCodeHTMLPusher(qrCode string) {
	if w.Streamer.Pusher.Channel != "" && w.Streamer.Pusher.EventName != "" {
		w.Streamer.Pusher.Client.Trigger(w.Streamer.Pusher.Channel, w.Streamer.Pusher.EventName, qrCode)
	}
}

func (w *EasyWhatsapp) RenderQRCodeHTMLNsq(qrCode string) {
	if w.Streamer.NSQ.Host != "" && w.Streamer.NSQ.Port != "" {
		config := nsq.NewConfig()
		producer, err := nsq.NewProducer(fmt.Sprintf("%s:%s", w.Streamer.NSQ.Host, w.Streamer.NSQ.Port), config)
		if err != nil {
			log.Fatalf("Error Produce Message : %v \n", err.Error())
		}

		err = producer.Publish(w.Streamer.NSQ.Topic, []byte(qrCode))
		if err != nil {
			log.Fatalf("Error Produce Message : %v \n", err.Error())
		}

		producer.Stop()
	}
}
