package main

import (
	"html/template"
	"log"
	"os"
)

func main()  {
	const templ = `<p>A: {{.A}}</p><p>B: {{.B}}</p>`

	t := template.Must(template.New("escape").Parse(templ))

	var data struct{
		A string		// 未信任的純文字
		B template.HTML // 受信任的 HTML
	}

	data.A = "<b>Hello!</b>"
	data.B = "<b>Hello!</b>"

	if err := t.Execute(os.Stdout, data); err != nil {
		log.Fatal(err)
	}
}
