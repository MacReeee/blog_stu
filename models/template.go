package models

import (
	"html/template"
	"io"
	"log"
	"time"
)

type TemplateBlog struct {
	*template.Template
}

type HtmlTemplate struct {
	Index      TemplateBlog
	Category   TemplateBlog
	Custom     TemplateBlog
	Detail     TemplateBlog
	Login      TemplateBlog
	Pigeonhole TemplateBlog
	Writing    TemplateBlog
}

func (t *TemplateBlog) WriteData(w io.Writer, date any) {
	err := t.Execute(w, date)
	if err != nil {
		_, err := w.Write([]byte(err.Error()))
		if err != nil {
			log.Println(err)
			return
		}
	}
}

func (t *TemplateBlog) WriteError(w io.Writer, err error) {
	if err != nil {
		_, err := w.Write([]byte(err.Error()))
		if err != nil {
			log.Println(err)
			return
		}
	}
}

func InitTemplate(dir string) (HtmlTemplate, error) {
	tp, err := readTemplate(
		[]string{"index", "category", "custom", "detail", "login", "pigeonhole", "writing"},
		dir,
	)
	var htmlTemplate HtmlTemplate

	if err != nil {
		return htmlTemplate, err
	}
	htmlTemplate.Index = tp[0]
	htmlTemplate.Category = tp[1]
	htmlTemplate.Custom = tp[2]
	htmlTemplate.Detail = tp[3]
	htmlTemplate.Login = tp[4]
	htmlTemplate.Pigeonhole = tp[5]
	htmlTemplate.Writing = tp[6]
	return htmlTemplate, nil
}

func readTemplate(templates []string, path string) ([]TemplateBlog, error) {
	var tbs []TemplateBlog
	for _, view := range templates {
		viewName := view + ".html"
		t := template.New(viewName)
		//2.拿到各个页面需要的前端页面
		home := path + "home.html"
		header := path + "layout/header.html"
		footer := path + "layout/footer.html"
		personal := path + "layout/personal.html"
		post := path + "layout/post-list.html"
		pagination := path + "layout/pagination.html"
		t.Funcs(template.FuncMap{"isODD": isODD, "getNextName": getNextName, "date": date, "dateDay": dateDay})
		t, err := t.ParseFiles(path+viewName, home, header, footer, personal, post, pagination)
		if err != nil {
			log.Println("错误:", err)
			return nil, err
		}
		tbs = append(tbs, TemplateBlog{Template: t})
	}
	return tbs, nil
}

func isODD(num int) bool {
	return num%2 == 0
}

func getNextName(strs []string, index int) string {
	return strs[index+1]
}

func date(layout string) string {
	return time.Now().Format(layout)
}

func dateDay(date time.Time) string {
	return date.Format("2006-01-02 15:04:05")
}
