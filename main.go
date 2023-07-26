package main

import (
	"batch48/connection"
	"context"
	"fmt"
	"net/http"
	"strconv"
	"text/template"
	"time"

	"github.com/labstack/echo/v4"
)

type Project struct {
	Id        int
	Name      string
	StarDate  string
	EndDate   string
	Duration  string
	Detail    string
	Playstore bool
	Android   bool
	Java      bool
	React     bool
}

var dataProject = []Project{
	// {
	// 	Id:        0,
	// 	Name:      "Contoh Project-1",
	// 	StarDate:  "15-05-2023",
	// 	EndDate:   "15-06-2023",
	// 	Duration:  "1 bulan",
	// 	Detail:    "Bootcamp sebulan gaes",
	// 	Playstore: true,
	// 	Android:   true,
	// 	Java:      true,
	// 	React:     true,
	// },

	// {
	// 	Id:        1,
	// 	Name:      "Contoh Projec-2",
	// 	StarDate:  "15-05-2023",
	// 	EndDate:   "15-06-2023",
	// 	Duration:  "1 bulan",
	// 	Detail:    "Bootcamp sebulan gaes hehe",
	// 	Playstore: true,
	// 	Android:   true,
	// 	Java:      true,
	// 	React:     true,
	// },
}

func main() {
	connection.DataBaseConnect()
	e := echo.New()

	// e = echo package nyaa...
	// get = method yg akan djalankan...
	// npoin nya / routing...

	// untuk mengirim folder routing statis
	e.Static("/public", "public")

	e.GET("/hello", helloWolrd)
	e.GET("/index", home)
	e.GET("/addproject-data", project)
	e.GET("/formaddproject", formproject)

	e.GET("/testimonial", testimonial)
	e.GET("/contact", contact)
	e.GET("/project-detail/:id", projectDetail)
	e.GET("/edit-addproject/:id", editProject)

	// rout post nya
	e.POST("/delete-addproject/:id", deleteProject)
	e.POST("/edit-addproject/:id", submitEditedProject)
	e.POST("/addproject", submitProject)

	e.Logger.Fatal(e.Start(":1234"))
}
func helloWolrd(c echo.Context) error {
	return c.String(http.StatusOK, "hello, ibab")
}
func home(c echo.Context) error {
	var tmpl, err = template.ParseFiles("views/index.html")

	data, _ := connection.Conn.Query(context.Background(), "SELECT id, name, start_date, end_date, duration, detail, playstore, android, java, react FROM tb_project")

	var result []Project
	for data.Next() {
		var each = Project{}
		err := data.Scan(&each.Id, &each.Name, &each.StarDate, &each.EndDate, &each.Duration, &each.Detail, &each.Playstore, &each.Android, &each.Java, &each.React)

		if err != nil {
			fmt.Println(err.Error())
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		result = append(result, each)

	}


	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	dataIndex := map[string]interface{}{
		"Blogs": result,
	}
	fmt.Println("ini data index", dataIndex)

	return tmpl.Execute(c.Response(), dataIndex)
}
func project(c echo.Context) error {

	data, _ := connection.Conn.Query(context.Background(), "SELECT id, name, start_date, end_date, duration, detail, playstore, android, java, react FROM tb_project")

	var result []Project
	for data.Next() {
		var each = Project{}
		err := data.Scan(&each.Id, &each.Name, &each.StarDate, &each.EndDate, &each.Duration, &each.Detail, &each.Playstore, &each.Android, &each.Java, &each.React)

		if err != nil {
			fmt.Println(err.Error())
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		result = append(result, each)

	}
	var tmpl, err = template.ParseFiles("views/addproject.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	Projects := map[string]interface{}{
		"Projects": result,
	}

	return tmpl.Execute(c.Response(), Projects)
}

func formproject(c echo.Context) error {
	var tmpl, err = template.ParseFiles("views/formaddproject.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return tmpl.Execute(c.Response(), nil)
}
// func untuk rout post nyaaa
func submitProject(c echo.Context) error {
	name := c.FormValue("input-project-title")
	detail := c.FormValue("input-description")

	// ambil date input
	date1 := c.FormValue("input-startdate")
	date2 := c.FormValue("input-enddate")

	// parse date input dan formatting
	uDate1, _ := time.Parse("2006-01-02", date1)
	starDate := uDate1.Format("2 Jan 2006")

	uDate2, _ := time.Parse("2006-01-02", date2)
	endDate := uDate2.Format("2 Jan 2006")

	// perhitungan selisih
	var diffUse string
	timeDiff := uDate2.Sub(uDate1)

	if timeDiff.Hours()/24 < 30 {
		tampil := strconv.FormatFloat(timeDiff.Hours()/24, 'f', 0, 64)
		diffUse = "Duration : " + tampil + " hari"
	} else if timeDiff.Hours()/24/30 < 12 {
		tampil := strconv.FormatFloat(timeDiff.Hours()/24/30, 'f', 0, 64)
		diffUse = "Duration : " + tampil + " Bulan"
	} else {

	}
	// checkbox
	var nodejs bool
	if c.FormValue("nodejs") == "checked" {
		nodejs = true
	}

	var reactjs bool
	if c.FormValue("reactjs") == "checked" {
		reactjs = true
	}

	var nextjs bool
	if c.FormValue("nextjs") == "checked" {
		nextjs = true
	}

	var typescript bool
	if c.FormValue("typescript") == "checked" {
		typescript = true
	}

	
	_, err := connection.Conn.Exec(context.Background(), "INSERT INTO tb_project (name, detail, start_date, end_date, duration, playstore, android, java, react) VALUES ($1 , $2, $3, $4, $5, $6, $7, $8, $9)", name, detail, starDate, endDate, diffUse, nodejs, reactjs, nextjs, typescript,)

	if err != nil{
		c.JSON(http.StatusInternalServerError, err.Error())
	}

	fmt.Println(dataProject)

	return c.Redirect(http.StatusMovedPermanently, "/addproject-data")
}

// detail project
func projectDetail(c echo.Context) error {
	// strconcov/string converter = untuk conver tipe data lain jadi string
	id := c.Param("id")
	var tmpl, errTmp= template.ParseFiles("views/project-detail.html")

	idToInt, _ := strconv.Atoi(id)

	var ProjectDetail = Project{}

	// query row untuk nge get satu datanya
	err := connection.Conn.QueryRow(context.Background(), "SELECT id, name, start_date, end_date, duration, detail, playstore, android, java, react FROM tb_project WHERE id=$1", idToInt).Scan(&ProjectDetail.Id, &ProjectDetail.Name, &ProjectDetail.StarDate, &ProjectDetail.EndDate, &ProjectDetail.Duration, &ProjectDetail.Detail, &ProjectDetail.Playstore, &ProjectDetail.Android, &ProjectDetail.Java, &ProjectDetail.React)

	fmt.Println("ini data blog detail", err)

	if err != nil{
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	data := map[string]interface{}{
		"Id":   id,
		"Project": ProjectDetail,
	}

	
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": errTmp.Error()})
	}

	return tmpl.Execute(c.Response(), data)
}
func editProject(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	var ProjectDetail = Project{}

	err := connection.Conn.QueryRow(context.Background(), "SELECT * FROM tb_project WHERE id=$1", id).Scan(&ProjectDetail.Name, &ProjectDetail.StarDate, &ProjectDetail.EndDate, &ProjectDetail.Duration, &ProjectDetail.Detail, &ProjectDetail.Playstore, &ProjectDetail.Android, &ProjectDetail.Java, &ProjectDetail.React)

	if err != nil{
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	data := map[string]interface{}{
		"Project": ProjectDetail,
		"Id":      id,
	}

	var tmpl, errTmp = template.ParseFiles("views/editproject.html")
	if errTmp != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return tmpl.Execute(c.Response(), data)

}
func submitEditedProject(c echo.Context) error {

	// Menangkap Id dari Query Params
	id, _ := strconv.Atoi(c.Param("id"))

	name := c.FormValue("input-project-title")
	starDate := c.FormValue("input-startdate")
	endDate := c.FormValue("input-enddate")
	detail := c.FormValue("input-description")
	// checkbox
	var nodejs bool
	if c.FormValue("nodejs") == "checked" {
		nodejs = true
	}

	var reactjs bool
	if c.FormValue("reactjs") == "checked" {
		reactjs = true
	}

	var nextjs bool
	if c.FormValue("nextjs") == "checked" {
		nextjs = true
	}

	var typescript bool
	if c.FormValue("typescript") == "checked" {
		typescript = true
	}

	// dataProject[id] = editedProject
	_, err := connection.Conn.Exec(context.Background(), "UPDATE tb_project SET name=$1, start_date=$2, end_date=$3, duration=$4, detail=$5, playstore=$6, android=$7, java=$8, react=$9, WHERE id=$10", name, starDate, endDate, detail, nodejs, reactjs, nextjs, typescript, id)

	if err != nil{
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}


	return c.Redirect(http.StatusMovedPermanently, "/addproject-data")
}
func deleteProject(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	// dataProject = append(dataProject[:id], dataProject[id+1:]...)

	// query delete to data base
	connection.Conn.Exec(context.Background(), "DELETE FROM tb_project WHERE id=$1", id)
	return c.Redirect(http.StatusMovedPermanently, "/addproject-data")
}
func countDuration(d1 string, d2 string) string {
	date1, _ := time.Parse("2006-01-02", d1)
	date2, _ := time.Parse("2006-01-02", d2)

	diff := date2.Sub(date1)
	days := int(diff.Hours() / 24)
	weeks := days / 7
	months := days / 30

	if months > 12 {
		return strconv.Itoa(months/12) + " tahun"
	}
	if months > 0 {
		return strconv.Itoa(months) + " bulan"
	}
	if weeks > 0 {
		return strconv.Itoa(weeks) + " minggu"
	}
	return strconv.Itoa(days) + " hari"
}
func testimonial(c echo.Context) error {
	var tmpl, err = template.ParseFiles("views/testimonial.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	return tmpl.Execute(c.Response(), nil)
}
func contact(c echo.Context) error {
	var tmpl, err = template.ParseFiles("views/contact.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	return tmpl.Execute(c.Response(), nil)
}
