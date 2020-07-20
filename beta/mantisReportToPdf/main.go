package main

import (
	"database/sql"
	"fmt"
	"strconv"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jung-kurt/gofpdf"
)

const (
	//имя пользователя базы
	dbUserName = "mantis"
	//пароль для пользователя базы
	dbPassword = "mantis"
	//имя базы данных
	dbName = "mantis"
	//сетевые параметры подключения к sql серверу
	sqlServerCredential = "tcp(192.168.0.1:3306)"
	//DSN строка для подключения
	dsn = dbUserName + ":" + dbPassword + "@" + sqlServerCredential + "/" + dbName
	//URL формы опроса для проведения актуализации контента
	yandexFormURL = "https://forms.yandex.ru/u/yourFormID/"
	//Служебный почтовый ящик, используемый для отправки сообщений из Яндекс.Форм и принятия сообщений в MantisBT
	emailBoxName = "get@site.loc"
	//ID проекта "Наименование проекта" в таблице mantis.mantis_project_table (описание схемы БД https://mantisbt.org/docs/erd/latest.pdf )
	projectID = 1
	//ID категории внутри проекта projectID, в которую помещается принятое сообщение с почтового ящика
	categoryID = 1
)

//описания статусов внутри конфигурационного файла mantis - значения по умолчанию
var statusEnumString = map[int]string{
	10: "новая",
	20: "обратная связь",
	30: "рассматривается",
	40: "подтверждена",
	50: "назначена",
	80: "решена",
	90: "закрыта",
}
var resolutionEnumString = map[int]string{
	10: "открыта",
	20: "решена",
	30: "переоткрыта",
	40: "не удается воспроизвести",
	50: "нерешаема",
	60: "повтор",
	70: "изменения не нужны",
	80: "отложена",
	90: "отказ в исправлении",
}

/*
type mantisBugTable struct {
	id            int
	projectID     int
	status        int
	resolution    int
	summary       string
	categoryID    int
	dateSubmitted int64
	lastUpdated   int64
}
*/

type contentValidationTable struct {
	id            int
	status        int
	resolution    int
	summary       string
	dateSubmitted int64
	lastUpdated   int64
	projectName   string
	categoryName  string
	description   string
}

func main() {
	fmt.Println("Hello, i`m ready to import")

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	sqlRequest := "SELECT mantis.mantis_bug_table.id, " +
		"mantis.mantis_bug_table.status, " +
		"mantis.mantis_bug_table.resolution, " +
		"mantis.mantis_bug_table.summary, " +
		"mantis.mantis_bug_table.date_submitted, " +
		"mantis.mantis_bug_table.last_updated, " +
		"mantis.mantis_project_table.name as project_name, " +
		"mantis.mantis_category_table.name as category_name, " +
		"mantis.mantis_bug_text_table.description " +
		"FROM mantis.mantis_bug_table " +
		"left join mantis.mantis_project_table on mantis.mantis_project_table.id = mantis.mantis_bug_table.project_id " +
		"left join mantis.mantis_category_table on mantis.mantis_category_table.id = mantis.mantis_bug_table.category_id " +
		"left join mantis.mantis_bug_text_table on mantis.mantis_bug_text_table.id = mantis.mantis_bug_table.id " +
		"where mantis.mantis_bug_table.project_id = " + strconv.Itoa(projectID) + " and mantis.mantis_bug_table.category_id = " + strconv.Itoa(categoryID) + ";"

	rows, err := db.Query(sqlRequest)
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()

	mapRows := map[int]map[string]string{}
	for rows.Next() {
		mapRowContent := map[string]string{}
		s := contentValidationTable{}
		err = rows.Scan(&s.id, &s.status, &s.resolution, &s.summary, &s.dateSubmitted, &s.lastUpdated, &s.projectName, &s.categoryName, &s.description)
		if err != nil {
			fmt.Println(err)
			continue
		}
		mapRowContent["status"] = statusEnumString[s.status]
		mapRowContent["resolution"] = resolutionEnumString[s.resolution]
		mapRowContent["summary"] = s.summary
		mapRowContent["dateSubmitted"] = time.Unix(s.dateSubmitted, 0).String()
		mapRowContent["lastUpdated"] = time.Unix(s.lastUpdated, 0).String()
		mapRowContent["projectName"] = s.projectName
		mapRowContent["categoryName"] = s.categoryName
		mapRowContent["description"] = s.description
		mapRows[s.id] = mapRowContent

	}

	pdf := gofpdf.New("P", "mm", "A4", "C:/Users/ksa/go/src/github.com/jung-kurt/gofpdf/font")
	pdf.AddFont("Helvetica", "", "helvetica_1251.json")
	pdf.AddPage()
	pdf.SetFont("Helvetica", "", 16)
	tr := pdf.UnicodeTranslatorFromDescriptor("cp1251")
	pdf.Cell(0, 20, tr("Валидация контента на сайте"))
	pdf.SetXY(10, 40)
	pdf.SetFont("Helvetica", "", 12)

	for key, val := range mapRows {
		sD := strings.Split(val["description"], "#")
		sliceToPdf := make([]string, 0)
		sliceToPdf = append(sliceToPdf, "Номер задачи:"+strconv.Itoa(key))
		sliceToPdf = append(sliceToPdf, "Состояние задачи: "+val["status"])
		sliceToPdf = append(sliceToPdf, "Решение задачи: "+val["resolution"])
		sliceToPdf = append(sliceToPdf, "Есть изменения: "+sD[4])
		sliceToPdf = append(sliceToPdf, "Факультет: "+sD[5])
		sliceToPdf = append(sliceToPdf, "Отправитель: "+sD[3])
		sliceToPdf = append(sliceToPdf, "Email отправителя: "+sD[1])
		sliceToPdf = append(sliceToPdf, "Дата получения информации: "+val["dateSubmitted"])
		sliceToPdf = append(sliceToPdf, "Дата последнего изменения: "+val["lastUpdated"])

		for idx := range sliceToPdf {
			pdf.CellFormat(0, 5, tr(sliceToPdf[idx]), "", 1, "", false, 0, "")
		}

		pdf.CellFormat(0, 5, "- -- --- ----- -------- -------------", "", 1, "", false, 0, "")
		pdf.CellFormat(0, 5, " ", "", 1, "", false, 0, "")
	}

	reportName := "report.pdf"
	err = pdf.OutputFileAndClose(reportName)
	if err == nil {
		fmt.Println("Документ " + reportName + " сформирован")
	} else {
		fmt.Println("Произошла ошибка при формировании документа:")
		fmt.Println(err)
	}

}
