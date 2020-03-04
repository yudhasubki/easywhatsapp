package easywhatsapp

import (
	"encoding/json"
	"fmt"
)

type Group struct {
	Status  int    `json:"code"`
	GroupID string `json:"gid"`
}

func (e *EasyWhatsapp) CreateGroup(groupName string, participants []string) (Group, error) {
	var group Group
	js, err := e.Connection.CreateGroup(groupName, participants)
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
