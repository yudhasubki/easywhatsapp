package easywhatsapp

import (
	"regexp"
	"strings"
)

const (
	GROUP_IDENTIFICATION = "@g.us"
)

func (w *MessageHandler) GetGroupJID() map[string]string {
	rex := regexp.MustCompile(GROUP_IDENTIFICATION)
	groups := make(map[string]string)
	for _, j := range w.RemoteJID {
		if rex.Match([]byte(j)) {
			idx := strings.Index(j, "-")
			groups[j] = j[idx+1 : len(j)]
		}
	}
	return groups
}
