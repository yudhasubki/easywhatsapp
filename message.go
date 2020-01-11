package easywhatsapp

import (
	"errors"
	"fmt"

	"github.com/Rhymen/go-whatsapp"
)

type History struct {
	messages []string
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
