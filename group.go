package easywhatsapp

import (
	"encoding/json"
	"fmt"
)

type Group struct {
	Status  int    `json:"code"`
	GroupID string `json:"gid"`
}

func (e *EasyWhatsapp) CreateGroup() (Group, error) {
	var group Group
	js, err := e.Connection.CreateGroup("Bongkeng", []string{"6281992156001@s.whatsapp.net"})
	if err != nil {
		fmt.Println(err)
	}

	groupInfo := <-js
	err = json.Unmarshal([]byte(groupInfo), &group)
	if err != nil {
		return Group{}, err
	}
	return group, nil
}
