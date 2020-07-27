package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/go-mail/mail"
	"github.com/jung-kurt/gofpdf"
	"github.com/k3a/html2text"
)

func openFileByte(fileName string) []byte {
	fileMt, errMt := os.Open(fileName)
	if errMt != nil {
		fmt.Println(errMt)
		os.Exit(1)
	}
	defer fileMt.Close()

	statMt, errMt := fileMt.Stat()
	if errMt != nil {
		fmt.Println(errMt)
		os.Exit(1)
	}

	bs := make([]byte, statMt.Size())
	_, err := fileMt.Read(bs)
	if err != nil {
		os.Exit(1)
	}
	return bs
}

func openFile(fileName string) string {
	bs := openFileByte(fileName)
	fileContent := string(bs)

	return fileContent
}

func cleanSharp(incomeString string) string {
	str := strings.Split(incomeString, "\n")
	cleanedStr := []string{}
	for _, val := range str {
		if string(val[0]) != "#" {
			cleanedStr = append(cleanedStr, val)
		}
	}
	returnString := ""
	for _, val := range cleanedStr {
		returnString = returnString + val + "\n"
	}

	return returnString
}

func sendEmail(d *mail.Dialer, email string, from string, subj string, message string) error {
	m := mail.NewMessage()
	m.SetHeader("From", from)
	m.SetHeader("To", email)
	m.SetHeader("Subject", subj)
	m.SetBody("text/html", message)
	err := d.DialAndSend(m)

	return err
}

func sendEmailWithAttach(d *mail.Dialer, email string, from string, subj string, message string, attach string) error {
	m := mail.NewMessage()
	m.SetHeader("From", from)
	m.SetHeader("To", email)
	m.SetHeader("Subject", subj)
	m.SetBody("text/html", message)
	m.Attach(attach)
	err := d.DialAndSend(m)

	return err
}

func makeReport(errMap map[string]error, subject string, messageTextToReport string) (string, error) {
	pdf := gofpdf.New("P", "mm", "A4", "C:/Users/ksa/go/src/github.com/jung-kurt/gofpdf/font")
	pdf.AddFont("Helvetica", "", "helvetica_1251.json")
	pdf.AddPage()
	pdf.SetFont("Helvetica", "", 16)
	tr := pdf.UnicodeTranslatorFromDescriptor("cp1251")
	pdf.Cell(0, 20, tr(subject))
	pdf.SetXY(10, 40)
	pdf.SetFont("Helvetica", "", 12)
	currentTime := time.Now()
	pdf.CellFormat(0, 5, currentTime.Format("2006-01-02 15:04:05"), "", 1, "", false, 0, "")
	pdf.CellFormat(0, 5, " ", "", 1, "", false, 0, "")
	for keyEm, valEm := range errMap {
		cellText := ""
		if valEm == nil {
			cellText = keyEm + " отправка произведена"
		} else {
			cellText = "Отправка сообщения на почту " + keyEm + " завершилась с ошибкой: " + valEm.Error()
		}
		pdf.CellFormat(0, 5, tr(cellText), "", 1, "", false, 0, "")
		fmt.Println(keyEm, valEm)
	}
	pdf.CellFormat(0, 5, " ", "", 1, "", false, 0, "")
	pdf.CellFormat(0, 5, " ", "", 1, "", false, 0, "")
	pdf.CellFormat(0, 5, tr("Отправляемое сообщение"), "", 1, "", false, 0, "")
	plainText := html2text.HTML2Text(messageTextToReport)
	pdf.CellFormat(0, 5, "--------------------------------", "", 1, "", false, 0, "")
	pdf.SetFont("Helvetica", "", 10)
	pdf.MultiCell(200, 6, tr(plainText), "", "L", false)
	reportName := "report.pdf"
	err := pdf.OutputFileAndClose(reportName)

	return reportName, err
}

func getMalboxCredential(serderCredential string) (string, int, string, string, string, string) {
	in := openFile(serderCredential)
	in = cleanSharp(in)
	r := csv.NewReader(strings.NewReader(in))
	mailboxCredential := []string{}
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		mailboxCredential = record
	}
	smtpServer := strings.TrimSpace(mailboxCredential[0])
	smtpPort, _ := strconv.Atoi(strings.TrimSpace(mailboxCredential[1]))
	mailboxName := strings.TrimSpace(mailboxCredential[2])
	mailboxPass := strings.TrimSpace(mailboxCredential[3])
	subject := strings.TrimSpace(mailboxCredential[4])
	emailReporter := strings.TrimSpace(mailboxCredential[5])

	return smtpServer, smtpPort, mailboxName, mailboxPass, subject, emailReporter
}

func main() {
	fmt.Println("Hello, I`m Ready To Send Email")

	smtpServer, smtpPort, mailboxName, mailboxPass, subject, emailReporter := getMalboxCredential("senderCredential.csv")

	d := mail.NewDialer(smtpServer, smtpPort, mailboxName, mailboxPass)
	d.StartTLSPolicy = mail.MandatoryStartTLS

	yandexFormURL := openFile("yandexFormURL.txt")
	yandexFormURL = cleanSharp(yandexFormURL)
	yandexFormURL = strings.TrimSuffix(yandexFormURL, "\n")

	messageText := openFile("messageText.txt")
	messageText = cleanSharp(messageText)
	messageTextToReport := messageText

	formURLCounter := strings.Count(messageText, "formURL")
	messageText = strings.Replace(messageText, "formURL", yandexFormURL, formURLCounter)

	mailboxes := openFile("mailBoxList.txt")
	mailboxes = cleanSharp(mailboxes)
	mailboxes = strings.TrimSuffix(mailboxes, "\n")

	mb := map[string]string{}
	for _, val := range strings.Split(mailboxes, "\n") {
		mb[strings.Split(val, "|")[0]] = strings.Split(val, "|")[1]
	}

	errMap := map[string]error{}
	topicListCounter := strings.Count(messageText, "topicList")
	for email, topics := range mb {
		from := mailboxName
		subj := subject
		messageText1 := strings.Replace(messageText, "topicList", topics, topicListCounter)
		sendEmailErr := sendEmail(d, email, from, subj, messageText1)
		errMap[email] = sendEmailErr
	}

	reportFile, err := makeReport(errMap, subject, messageTextToReport)
	if err == nil {
		fmt.Println("Документ " + reportFile + " сформирован")
		_ = sendEmailWithAttach(d, emailReporter, mailboxName, "Отчет об отправке", "Отчет об отправке во вложении", reportFile)
	} else {
		_ = sendEmailWithAttach(d, emailReporter, mailboxName, "Отчет об отправке", "Произошла ошибка при формировании документа", reportFile)
		fmt.Println("Произошла ошибка при формировании документа:")
		fmt.Println(err)
	}
}
