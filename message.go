package easywhatsapp

import (
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/Rhymen/go-whatsapp"
	"github.com/Rhymen/go-whatsapp/binary"
	"github.com/Rhymen/go-whatsapp/binary/proto"
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

type SearchInfo struct {
	JID    string
	FromMe bool
}

func (m *MessageHandler) ShouldCallSynchronously() bool {
	return true
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

func (m *MessageHandler) HandleError(err error) {
	log.Printf("error occured while retrieving chat history: %s", err.Error())
}

func (w *EasyWhatsapp) AddHandler() {
	w.Message = MessageHandler{
		Connection: w.Connection,
		RemoteJID:  make(map[string]string),
	}
	w.Connection.AddHandler(w)
}

func (w *EasyWhatsapp) GetHistory(remoteJid string, chunkSize int) []History {
	w.Message = MessageHandler{Connection: w.Connection}
	w.Connection.LoadFullChatHistory(remoteJid, chunkSize, time.Millisecond*0, &w.Message)
	return w.Message.Messages
}

func (w *EasyWhatsapp) GetChatByJID(remoteJid string, chunkSize int) []History {
	messages := w.GetHistory(remoteJid, chunkSize)
	histories := make([]History, 0)

	for _, message := range messages {
		histories = append(histories, History{
			AuthorID:   message.AuthorID,
			Message:    message.Message,
			Date:       message.Date,
			ScreenName: message.ScreenName,
		})
	}
	return histories
}

func (w *EasyWhatsapp) GetChats(chunkSize int) (histories []History) {
	var jids []string
	histories = make([]History, 0)
	for _, jid := range w.Message.RemoteJID {
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

func (w *EasyWhatsapp) GetGroupChats(chunkSize int) []History {
	histories := make([]History, 0)
	for _, g := range w.Message.GetGroupJID() {
		messages := w.GetHistory(g, chunkSize)
		history := make(chan History)
		go appendMessage(messages, history)
		histories = append(histories, <-history)
	}
	return histories
}

func appendMessage(messages []History, history chan History) {
	for _, message := range messages {
		hs := History{
			AuthorID:   message.AuthorID,
			Message:    message.Message,
			Date:       message.Date,
			ScreenName: message.ScreenName,
		}
		history <- hs
	}
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

func (e *EasyWhatsapp) SearchMessage(keyMessage string) (bool, SearchInfo, error) {
	info := SearchInfo{}
	query, err := e.Connection.Search(keyMessage, 100, 1)
	if err != nil {
		return false, SearchInfo{}, err
	}

	msg := decodeMessages(query)
	for _, m := range msg {
		if m.Message.Conversation != nil && m.Key.FromMe != nil {
			if keyMessage == *m.Message.Conversation && *m.Key.FromMe {
				info.JID = *m.Key.RemoteJid
				info.FromMe = *m.Key.FromMe
			}
		}
	}
	return true, info, nil
}

func decodeMessages(n *binary.Node) []*proto.WebMessageInfo {
	var messages = make([]*proto.WebMessageInfo, 0)
	if n == nil || n.Attributes == nil || n.Content == nil {
		return messages
	}

	for _, msg := range n.Content.([]interface{}) {
		switch msg.(type) {
		case *proto.WebMessageInfo:
			messages = append(messages, msg.(*proto.WebMessageInfo))
		default:
			log.Println("decodeMessages: Non WebMessage encountered")
		}
	}

	return messages
}
