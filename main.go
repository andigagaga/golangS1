package main

import (
	
	"net/http"
	"strconv"
	"text/template"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	
	// e = echo package nyaa...
	// get = method yg akan djalankan...
	// npoin nya / routing...
	
	// untuk mengirim folder routing statis
	e.Static("/public", "public")

	e.GET ("/hello", helloWolrd)
	e.GET ("/index", home)
	e.GET ("/addproject", project)
	e.GET ("/testimonial", testimonial)
	e.GET ("/contact", contact)
	e.GET ("/project-detail/:id", projectDetail)

	// rout post nya
	e.POST("/", addFormProject)




	e.Logger.Fatal(e.Start(":1234"))
}
func helloWolrd (c echo.Context) error {
	return c.String(http.StatusOK, "hello, ibab")
}
func home (c echo.Context) error {
	var tmpl, err = template.ParseFiles("views/index.html")
	
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	return tmpl.Execute(c.Response(), nil)
}
func project (c echo.Context) error {
	var tmpl, err = template.ParseFiles("views/addproject.html")
	
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	return tmpl.Execute(c.Response(), nil)
}

// detail project
func projectDetail(c echo.Context) error {
	// strconcov/string converter = untuk conver tipe data lain jadi string
	// 
	id, _ := strconv.Atoi(c.Param("id"))

	data := map[string]interface{}{
		"id" : id,
		"Title" : "My Detail Project",
		"Content" : "Lorem ipsum dolor sit amet consectetur adipisicing elit. Consequuntur hic cupiditate neque? Fugiat saepe fuga labore vitae aperiam dolore cumque in assumenda consequuntur? Perferendis vero quibusdam quam eaque aliquam deleniti.",
	}
	var tmpl, err = template.ParseFiles("views/project-detail.html")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	return tmpl.Execute(c.Response(), data)
}
// func untuk rout post nyaaa
func addFormProject(c echo.Context) error {
			projectName := c.FormValue("input-project-title")
			startDate := c.FormValue("input-startdate")
			endDate := c.FormValue("input-enddate")
			message := c.FormValue("input-description")
			nodeJs := c.FormValue("nodejs")
			reactJs := c.FormValue("reactjs")
			nextJs := c.FormValue("nextjs")
			typeScript := c.FormValue("typescript")
			image := c.FormValue("input-project-image")

		// buat manggil variabel d atas dan agar value/nilainya dbtampilin d terminal
		println("Name Project: " + projectName)
		println("Start Date: " + startDate)
		println("End Date: " + endDate)
		println("Message: " + message)
		println("Technologies: " + nodeJs)
		println("Technologies: " + reactJs)
		println("Technologies: " + nextJs)
		println("Technologies: " + typeScript)
		println("Upload Images: " + image)

		return c.Redirect(http.StatusMovedPermanently, "/")

		


}



func testimonial (c echo.Context) error {
	var tmpl, err = template.ParseFiles("views/testimonial.html")
	
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	return tmpl.Execute(c.Response(), nil)
}
func contact (c echo.Context) error {
	var tmpl, err = template.ParseFiles("views/contact.html")
	
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	return tmpl.Execute(c.Response(), nil)
}

