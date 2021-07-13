package easywhatsapp

import (
	"fmt"
	"time"
)

func (w *EasyWhatsapp) Login() error {
	session, err := w.Read()
	if err == nil {
		session, err = w.Connection.RestoreWithSession(w.Session)
		if err != nil {
			for i := 0; i < w.RetryConnection; i++ {
				session, err = w.Connection.RestoreWithSession(w.Session)
				if i == w.RetryConnection && err != nil {
					return err
				}

				fmt.Println("Retrying restore session...")
				time.Sleep(time.Duration(w.RetryAfterFailed) * time.Second)
			}
		}
	} else {
		qr := w.GenerateQRCode()
		session, err = w.Connection.Login(qr)
		if err != nil {
			err = fmt.Errorf("invalid Login QR Code : %v", err.Error())
			return err
		}
	}

	w.Session = session
	err = w.Write()
	if err != nil {
		err = fmt.Errorf("invalid Save Session : %v", err.Error())
		return err
	}

	return nil
}
