package main

import (
	"encoding/json"
	"html/template"
	"os"
)

type Generic map[string]any

func main() {
	tem := template.Must(template.New("template.html").ParseFiles("template.html"))
	j := []byte(`{"teste": {"teste": "vitodfdr"}, "teste2": "Souza"}`)
	var p Generic
	erro := json.Unmarshal(j, &p)
	if erro != nil {
		panic(erro)
	}
	erro = tem.Execute(os.Stdout, p)
	if erro != nil {
		panic(erro)
	}
}
