package easywhatsapp

import (
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/Rhymen/go-whatsapp"
)

type MessageHandler struct {
	Connection *whatsapp.Conn
	Messages   []History
	RemoteJID  map[string]string
}

type History struct {
	AuthorID   string
	Message    string
	Date       int64
	ScreenName string
}

func (m *MessageHandler) ShouldCallSynchronously() bool {
	return true
}

func (m *MessageHandler) HandleError(err error) {
	log.Printf("Error occured while retrieving chat history: %s", err.Error())
}

func (m *MessageHandler) HandleTextMessage(message whatsapp.TextMessage) {
	authorID := "-"
	screenName := "-"
	if message.Info.FromMe {
		authorID = m.Connection.Info.Wid
		screenName = ""
	} else {
		if message.Info.Source.Participant != nil {
			authorID = *message.Info.Source.Participant
		} else {
			authorID = message.Info.RemoteJid
		}
		if message.Info.Source.PushName != nil {
			screenName = *message.Info.Source.PushName
		}
	}

	date := time.Unix(int64(message.Info.Timestamp), 0)
	m.Messages = append(m.Messages, History{
		AuthorID:   authorID,
		Message:    message.Text,
		Date:       date.Unix(),
		ScreenName: screenName,
	})
}

func (w *EasyWhatsapp) AddHandler() *EasyWhatsapp {
	w.Connection.AddHandler(w)
	w.Message = MessageHandler{
		Connection: w.Connection,
		RemoteJID:  make(map[string]string),
	}
	return w
}

func (w *MessageHandler) GetHistory(remoteJid string, chunkSize int) []History {
	handler := &MessageHandler{Connection: w.Connection}
	w.Connection.LoadFullChatHistory(remoteJid, chunkSize, time.Millisecond*0, handler)
	return handler.Messages
}

func (w *MessageHandler) GetChatByJID(remoteJid string, chunkSize int) (histories []History) {
	messages := w.GetHistory(remoteJid, chunkSize)
	for _, message := range messages {
		histories = append(histories, History{
			AuthorID:   message.AuthorID,
			Message:    message.Message,
			Date:       message.Date,
			ScreenName: message.ScreenName,
		})
	}
	return
}

func (w *MessageHandler) GetChats(chunkSize int) (histories []History) {
	var jids []string
	for _, jid := range w.RemoteJID {
		jids = append(jids, jid)
		messages := w.GetHistory(jid, chunkSize)
		for _, message := range messages {
			histories = append(histories, History{
				AuthorID:   message.AuthorID,
				Message:    message.Message,
				Date:       message.Date,
				ScreenName: message.ScreenName,
			})
		}
	}

	return
}

func (w *EasyWhatsapp) SendText(remoteJid string, message string) (string, error) {
	msg := whatsapp.TextMessage{
		Info: whatsapp.MessageInfo{
			RemoteJid: remoteJid,
		},
		Text: message,
	}

	msgId, err := w.Connection.Send(msg)
	if err != nil {
		err = errors.New(fmt.Sprintf("Error sending message : %v", err.Error()))
		return "", err
	}

	return msgId, nil
}
