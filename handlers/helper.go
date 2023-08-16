package handlers

import (
	"errors"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/nicksnyder/go-i18n/v2/i18n"
	. "github.com/unpWn4bl3/myChitchat/config"
	"github.com/unpWn4bl3/myChitchat/models"
)

// 日志组件
var logger *log.Logger
var config *Configuration
var localizer *i18n.Localizer

func init() {
	config = LoadConfig()
	localizer = i18n.NewLocalizer(config.LocaleBundle, config.App.Language)
	file, err := os.OpenFile("logs/chitchat.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Failed to open log file: ", err.Error())
	}
	logger = log.New(file, "INFO ", log.Ldate|log.Ltime|log.Lshortfile)
}

func info(args ...any) {
	logger.SetPrefix("INFO ")
	logger.Println(args...)
}

func warning(args ...any) {
	logger.SetPrefix("WARN ")
	logger.Println(args...)
}

func fatal(args ...any) {
	logger.SetPrefix("FATAL ")
	logger.Println(args...)
}

// 错误处理

// 重定向到404
func error_message(w http.ResponseWriter, r *http.Request, msg string) {
	url := fmt.Sprintf("/err?msg=%s", msg)
	http.Redirect(w, r, url, 302)
}

// 是否登录
func session(w http.ResponseWriter, r *http.Request) (sess models.Session, err error) {
	cookie, err := r.Cookie("_cookie")
	if err != nil {
		return
	}
	sess = models.Session{Uuid: cookie.Value}
	if ok, _ := sess.Check(); !ok {
		err = errors.New("Invalid session")
	}
	return
}

// 解析html模板
func parseTemplateFiles(filenames ...string) (t *template.Template) {
	files := make([]string, 0)
	t = template.New("layout")
	for _, file := range filenames {
		files = append(files, fmt.Sprintf("views/%s.html", file))
	}
	t = template.Must(t.ParseFiles(files...))
	return
}

func generateHTML(w http.ResponseWriter, data any, filenames ...string) {
	files := make([]string, 0)
	for _, file := range filenames {
		files = append(files, fmt.Sprintf("views/%s.html", file))
	}

	templates := template.Must(template.ParseFiles(files...))
	templates.ExecuteTemplate(w, "layout", data)
}

func Version() string {
	return "0.1"
}
