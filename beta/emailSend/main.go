package main

import (
	"fmt"

	"github.com/go-mail/mail"
)

func main() {
	emailFrom := "box@test.loc"
	emailTo := "karandin_sa@test.loc"
	emailSubject := "Актуализация контента на сайте test.loc"
	smtpServerName := "smtp.yourMail.server"
	smtpServerPort := 465
	smtpAccountName := "box@yourDomain.mail"
	smtpAccountPass := "yourPass"
	formURL := "https://forms.yandex.ru/u/yourIdForm/"
	messageString := "<p><b>Добрый день.</b></p>" +
		"<p>Прошу Вас проверить актуальность информации на сайте test.loc в закрепленных за Вами разделах: " +
		"<br>В случае необходимости внесения изменений на сайте Вам необходимо заполнить данную <a href=\"" + formURL + "\">форму</a></br></P>" +
		"<p>Это письмо сформировано автоматически и отвечать на него не нужно.</p>" +
		"<p><b>Прошу заполнить указанную <a href=\"" + formURL + "\">форму</a></b></p>" +
		"<p><b>С уважением, <br>почтовый робот test.loc</br></b></p>"

	fmt.Println("Hello, I`m Ready To Send Email")

	m := mail.NewMessage()
	m.SetHeader("From", emailFrom)
	m.SetHeader("To", emailTo)
	m.SetHeader("Subject", emailSubject)
	m.SetBody("text/html", messageString)

	d := mail.NewDialer(smtpServerName, smtpServerPort, smtpAccountName, smtpAccountPass)
	d.StartTLSPolicy = mail.MandatoryStartTLS

	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
}
