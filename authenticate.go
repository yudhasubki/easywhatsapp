package easywhatsapp

import (
	"fmt"
)

func (w *EasyWhatsapp) Login() error {
	session, err := w.Read()
	if err == nil {
		err = w.Connection.Restore()
		if err != nil {
			return fmt.Errorf("restoring failed: %v", err)
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
