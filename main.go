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
	// 	Name:      "Project 1",
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
	// 	Name:      "Project 2",
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
	e.POST("/addproject", saveProject)

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
	var tmpl, err = template.ParseFiles("views/addproject.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	Projects := map[string]interface{}{
		"Projects": dataProject,
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

// detail project
func projectDetail(c echo.Context) error {
	// strconcov/string converter = untuk conver tipe data lain jadi string
	id, _ := strconv.Atoi(c.Param("id"))

	var ProjectDetail = Project{}

	for i, data := range dataProject {
		if id == i {
			ProjectDetail = Project{
				Name:      data.Name,
				StarDate:  data.StarDate,
				EndDate:   data.EndDate,
				Duration:  data.Duration,
				Detail:    data.Detail,
				Playstore: data.Playstore,
				Android:   data.Android,
				Java:      data.Java,
				React:     data.React,
			}
		}
	}
	data := map[string]interface{}{
		"Project": ProjectDetail,
	}

	var tmpl, err = template.ParseFiles("views/project-detail.html")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return tmpl.Execute(c.Response(), data)
}
func editProject(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	var ProjectDetail = Project{}

	for i, data := range dataProject {
		if id == i {
			ProjectDetail = Project{
				Name:      data.Name,
				StarDate:  data.StarDate,
				EndDate:   data.EndDate,
				Duration:  data.Duration,
				Detail:    data.Detail,
				Playstore: data.Playstore,
				Android:   data.Android,
				Java:      data.Java,
				React:     data.React,
			}
		}
	}
	data := map[string]interface{}{
		"Project": ProjectDetail,
		"Id":      id,
	}

	var tmpl, err = template.ParseFiles("views/editproject.html")
	if err != nil {
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

	var editedProject = Project{
		Name:      name,
		StarDate:  starDate,
		EndDate:   endDate,
		Duration:  countDuration(starDate, endDate),
		Detail:    detail,
		Playstore: nodejs,
		Android:   reactjs,
		Java:      nextjs,
		React:     typescript,
	}

	dataProject[id] = editedProject
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

// func untuk rout post nyaaa
func saveProject(c echo.Context) error {
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

	var newProject = Project{
		Name:      name,
		StarDate:  starDate,
		EndDate:   endDate,
		Duration:  diffUse,
		Detail:    detail,
		Playstore: nodejs,
		Android:   reactjs,
		Java:      nextjs,
		React:     typescript,
	}

	dataProject = append(dataProject, newProject)

	fmt.Println(dataProject)

	return c.Redirect(http.StatusMovedPermanently, "/addproject-data")
}
func deleteProject(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	dataProject = append(dataProject[:id], dataProject[id+1:]...)
	return c.Redirect(http.StatusMovedPermanently, "/addproject-data")
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
