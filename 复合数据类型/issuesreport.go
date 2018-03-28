package main

import (
	"html/template"
	"log"
	"os"
	"time"

	"github.com/google/go-github/github"
)

const temp1 = `{{.TotalCount}} issues:
{{range .Items}}---------------------------------
Numbers: {{.Numbers}}
User: {{.User.Login}}
Title: {{.Title | printf "%.64s"}}
Age: {{.CreateAt | daysAgo}} days
{{end}}
`

func daysAgo(t time.Time) int {
	return int(time.Since(t).Hours() / 24)
}

var report = template.Must(template.New("report").
	Funcs(template.FuncMap{"daysAgo": daysAgo}).
	Parse(temp1))

func main() {
	result, err := github.SearchIssue(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	if err := report.Execute(os.Stdout, result); err != nil {
		log.Fatal(err)
	}
}
