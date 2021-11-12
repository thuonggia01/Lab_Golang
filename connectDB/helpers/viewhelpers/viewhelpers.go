package viewhelpers

import (
	"html/template"
	"time"
)

func Funcs() template.FuncMap {
	return template.FuncMap{
		"FormatTime": formatTime,
	}
}
func formatTime(t time.Time) string {
	return t.Format("2006/01/02 15:04:05")
}
