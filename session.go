package easywhatsapp

import (
	"encoding/gob"
	"fmt"
	"os"

	"github.com/Rhymen/go-whatsapp"
)

const (
	FILE_NAME = "whatsapp-session.gob"
)

func (w *EasyWhatsapp) Read() (whatsapp.Session, error) {
	file, err := os.Open(path())
	if err != nil {
		return w.Session, err
	}
	defer file.Close()

	decoder := gob.NewDecoder(file)
	err = decoder.Decode(&w.Session)
	if err != nil {
		return w.Session, err
	}

	return w.Session, nil
}

func (w *EasyWhatsapp) Write() error {
	file, err := os.Create(path())
	if err != nil {
		return err
	}
	defer file.Close()
	encoder := gob.NewEncoder(file)
	err = encoder.Encode(w.Session)
	if err != nil {
		return err
	}
	return nil
}

func (w *EasyWhatsapp) Delete() error {
	err := os.Remove(path())
	if err != nil {
		return err
	}
	return nil
}

func (w *EasyWhatsapp) Exist() bool {
	_, err := w.Read()
	if err == nil {
		return true
	}
	return false
}

func path() string {
	return fmt.Sprintf("%s/%s", os.TempDir(), FILE_NAME)
}
