package easywhatsapp

import (
	"regexp"
)

const (
	GROUP_IDENTIFICATION = "@g.us"
)

func (w *MessageHandler) GetGroupJID() map[string]string {
	rex := regexp.MustCompile(GROUP_IDENTIFICATION)
	groups := make(map[string]string)
	for _, j := range w.RemoteJID {
		if rex.Match([]byte(j)) {
			groups[j] = j
		}
	}
	return groups
}
