package easywhatsapp

func (w *EasyWhatsapp) GenerateQRCode() chan string {
	qr := make(chan string)
	go func() {
		qrCode := <-qr
		w.RenderQRCodeHTMLPusher(qrCode)
	}()
	return qr
}
