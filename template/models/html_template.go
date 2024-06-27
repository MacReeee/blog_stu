package models

import (
	"fmt"
	"goblog/config"
	"html/template"
	"io"
	"net/http"
	"time"
)

type TemplatePointer struct {
	*template.Template
}

type HtmlTemplate struct {
	Index    TemplatePointer
	Category TemplatePointer
	Pigeonhole  TemplatePointer
	Login   TemplatePointer
	Detail      TemplatePointer
	Writing      TemplatePointer
}

func (t TemplatePointer) WriteData(w io.Writer, data interface{}) {

	err := t.Execute(w, data)
	if err != nil {
		if _, e := w.Write([]byte(err.Error())); e != nil {
			fmt.Println(e)
		}
	}
}

func (t TemplatePointer) WriteError(w http.ResponseWriter, err error) {
	if _, e := w.Write([]byte(err.Error())); e != nil {
		fmt.Println(e)
	}
}

func BuildViewData(title string, data interface{}) map[string]interface{} {
	return map[string]interface{}{
		"Title":  title,
		"Data":   data,
		"Config": config.Cfg,
		"Navs":    config.Cfg.Viewer.Navigation,
	}
}

func InitHtmlTemplate(viewDir string) (HtmlTemplate, error) {
	var htmlTemplate HtmlTemplate

	tp, err := readHtmlTemplate(
		[]string{"index","category","pigeonhole","login","detail","writing"},
		viewDir)
	if err != nil {
		return htmlTemplate, err
	}

	htmlTemplate.Index = tp[0]
	htmlTemplate.Category = tp[1]
	htmlTemplate.Pigeonhole = tp[2]
	htmlTemplate.Login = tp[3]
	htmlTemplate.Detail = tp[4]
	htmlTemplate.Writing = tp[5]
	return htmlTemplate, nil
}

func SpreadDigit(n int) []int {
	var r []int
	for i := 1; i <= n; i++ {
		r = append(r,i)
	}
	return r
}
func IsODD(num int) bool {
	return num%2 == 0
}
func Date(layout string) string {
	return time.Now().Format(layout)
}
func DateDay(date time.Time) string {
	return date.Format("2006-01-02 15:04:05")
}
func GetNextName(strs []string, i int) interface{} {
	return strs[i+1]
}
func readHtmlTemplate(htmlFileName []string, viewDir string) ([]TemplatePointer, error) {
	var htmlTemplate []TemplatePointer

	head := viewDir + "/layout/header.html"
	footer := viewDir + "/layout/footer.html"

	for _, name := range htmlFileName {
		filePath := viewDir + "/" + name + ".html"
		tp := template.New(name + ".html")
		tp.Funcs(template.FuncMap{"isODD":IsODD,"date":Date,"dateDay":DateDay,"getNextName":GetNextName})
		var err error
		//if name == "index" {
		home := viewDir + "/home.html"
		personal := viewDir + "/layout/personal.html"
		postList := viewDir + "/layout/post-list.html"
		pagination := viewDir + "/layout/pagination.html"
		tp,err = tp.ParseFiles(filePath,home,personal,postList,pagination,head,footer)
		//}else if name == "category" {
		//	personal := viewDir + "/layout/personal.html"
		//	postList := viewDir + "/layout/post-list.html"
		//	pagination := viewDir + "/layout/pagination.html"
		//	tp,err = tp.ParseFiles(filePath,personal,postList,pagination,head,footer)
		//} else{
		//	tp,err = tp.ParseFiles(filePath,head,footer)
		//}
		if err != nil {
			return htmlTemplate, err
		}
		htmlTemplate = append(htmlTemplate, TemplatePointer{tp})
	}
	return htmlTemplate, nil
}
