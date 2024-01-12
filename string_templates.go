package main

import (
	"os"
	"text/template"
)

func main() {
	t1 := template.New("t1")
	t1, err := t1.Parse("Value is {{.}}\n")
	if err != nil {
		panic(err)
	}
	t1.Execute(os.Stdout, "some Text")
	t1.Execute(os.Stdout, 5)
	t1.Execute(os.Stdout, []string{"go", "rust"})
	Create := func(name, t string) *template.Template {
		return template.Must(template.New(name).Parse(t))
	}
	t2 := Create("t2", "Name {{.Name}}\n")
	t2.Execute(os.Stdout, struct{ Name string }{"jane doe"})
	t2.Execute(os.Stdout, map[string]string{"Name": "john doe"})

	t3 := Create("t3", "{{if . -}} yes {{else -}} no {{end}}\n")
	t3.Execute(os.Stdout, "foo")
	t3.Execute(os.Stdout, "")
	t3.Execute(os.Stdout, "  ")
	t3.Execute(os.Stdout, nil)
	t3.Execute(os.Stdout, false)
}
