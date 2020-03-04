package easywhatsapp

import (
	"log"
	"time"

	"github.com/Rhymen/go-whatsapp"
	"github.com/Rhymen/go-whatsapp/binary/proto"
	"github.com/pusher/pusher-http-go"
)

type EasyWhatsapp struct {
	Timeout         time.Duration
	Connection      *whatsapp.Conn
	Session         whatsapp.Session
	Streamer        Streamer
	Message         MessageHandler
	Synchronously   bool
	SessionPath     *string
	SessionFileName *string
}

type Streamer struct {
	Pusher PusherClient
}

func (m *EasyWhatsapp) ShouldCallSynchronously() bool {
	return m.Synchronously
}

func (m *EasyWhatsapp) HandleError(err error) {
	log.Printf("Error retrieving chat history : %s", err)
}

func (w *EasyWhatsapp) HandleRawMessage(message *proto.WebMessageInfo) {
	if message != nil && message.Key.RemoteJid != nil {
		w.Message.RemoteJID[*message.Key.Id] = *message.Key.RemoteJid
	}
}

func New(timeLimit int, synchronously bool) (*EasyWhatsapp, error) {
	t := timeout(timeLimit)
	conn, err := whatsapp.NewConn(t)
	if err != nil {
		return &EasyWhatsapp{}, err
	}

	return &EasyWhatsapp{
		Connection:    conn,
		Synchronously: synchronously,
		Timeout:       timeout(timeLimit),
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
