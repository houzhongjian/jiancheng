package utils

import (
	"fmt"
	"html/template"
)

//Display 分配视图.
func Display(name string) (*template.Template, error) {
	t, err := template.ParseFiles(fmt.Sprintf("./frontend/%s.html", name))
	return t, err
}
