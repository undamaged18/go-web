package engine

import (
	"ecommerce/services/config"
	"fmt"
	"html/template"
	"io/fs"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

var templateName string
var conf = config.New()

const templateExt = ".template.gohtml"
const layoutExt = ".layout.gohtml"
const componentExt = ".component.gohtml"

func init() {
	getCache()
}

func getCache() {
		tc, err := ParseTemplates()
		if err != nil {
			panic(err)
		}
		conf.Templates = tc
}

func RenderTemplate(w http.ResponseWriter, tmpl string, data map[string]interface{}) {
	if conf.Environment != "production" || conf.Templates == nil {
	getCache()
	}
	t, _ := conf.Templates[tmpl]
	_ = t.Execute(w, data)
}

// ParseTemplates runs through the templates directory and parses the HTML pages and returns
// *template.Template in a map[string]*template.Template format
func ParseTemplates() (map[string]*template.Template, error) {
	root := filepath.Clean(conf.Paths.Templates)
	var ts *template.Template
	cache := map[string]*template.Template{}
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			pages, err := filepath.Glob(fmt.Sprintf("%s/*%s", path, templateExt))
			if err != nil {
				return err
			}
			for _, page := range pages {
				var name string
				name = filepath.Base(page)
				ts, err = template.New(name).Funcs(funcMap()).ParseFiles(page)
				if err != nil {
					return err
				}

				layoutFiles := fmt.Sprintf("%s/layouts/*%s", root, layoutExt)
				layouts, err := filepath.Glob(layoutFiles)
				if err != nil {
					return err
				}

				if len(layouts) > 0 {
					ts, err = ts.ParseGlob(layoutFiles)
					if err != nil {
						return err
					}
				}

				componentFiles := fmt.Sprintf("%s/components/", conf.Paths.Templates)
				err = filepath.Walk(componentFiles, func(path string, info fs.FileInfo, err error) error {
					if info.IsDir() {
						componentSubFiles := fmt.Sprintf("%s/components/%s/*%s", root, info.Name(), componentExt)
						components, err := filepath.Glob(componentSubFiles)
						if err != nil {
							return err
						}
						if len(components) > 0 {
							ts, err = ts.ParseGlob(componentSubFiles)
							if err != nil {
								return err
							}
						}
					}
					return err
				})

				if info.Name() == "templates" {
					templateName = filepath.Base(page)
				} else {
					templateName = fmt.Sprintf("%s:%s", info.Name(), filepath.Base(page))
				}
				templateName = strings.Replace(templateName, templateExt, "", -1)
				cache[templateName] = ts
			}
		}
		return nil
	})
	return cache, err
}
