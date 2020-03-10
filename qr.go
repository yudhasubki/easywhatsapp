package easywhatsapp

func (w *EasyWhatsapp) GenerateQRCode() chan string {
	qr := make(chan string)
	go func() {
		qrCode := <-qr
		w.RenderQRCodeHTMLPusher(qrCode)
		w.RenderQRCodeConsole(qrCode)
	}()
	return qr
}
