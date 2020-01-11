package easywhatsapp

import (
	"time"

	"github.com/Rhymen/go-whatsapp"
	"github.com/pusher/pusher-http-go"
)

type EasyWhatsapp struct {
	Timeout    time.Duration
	Connection *whatsapp.Conn
	Session    whatsapp.Session
	Streamer   Streamer
}

type Streamer struct {
	Pusher PusherClient
	NSQ    NSQClient
}

// New instantiate of EasyWhatsapp Struct
// PusherClient with key : "APP_ID", "APP_KEY", "APP_SECRET", "APP_CLUSTER"
func (w *EasyWhatsapp) New(timeLimit int) (*EasyWhatsapp, error) {
	t := timeout(timeLimit)
	conn, err := whatsapp.NewConn(t)
	if err != nil {
		return &EasyWhatsapp{}, err
	}

	return &EasyWhatsapp{
		Connection: conn,
		Timeout:    timeout(timeLimit),
	}, nil
}

func timeout(timeout int) time.Duration {
	return time.Duration(timeout) * time.Second
}

func client(client map[string]string) pusher.Client {
	return pusher.Client{
		AppID:   client["APP_ID"],
		Key:     client["APP_KEY"],
		Secret:  client["APP_SECRET"],
		Cluster: client["APP_CLUSTER"],
	}
}
