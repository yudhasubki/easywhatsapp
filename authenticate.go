package easywhatsapp

import (
	"errors"
	"fmt"
)

func (w *EasyWhatsapp) Login() error {
	session, err := w.Read()
	if err == nil {
		session, err = w.Connection.RestoreWithSession(w.Session)
		if err != nil {
			err = errors.New(fmt.Sprintf("Invalid Restore Session : %v", err.Error()))
			return err
		}
	} else {
		qr := w.GenerateQRCode()
		session, err = w.Connection.Login(qr)
		if err != nil {
			err = errors.New(fmt.Sprintf("Invalid Login QR Code : %v", err.Error()))
			return err
		}
	}

	w.Session = session
	err = w.Write()
	if err != nil {
		err = errors.New(fmt.Sprintf("Invalid Save Session : %v", err.Error()))
		return err
	}

	return nil
}
