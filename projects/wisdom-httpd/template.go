package main

import (
	"bytes"
	"html/template"
)

// EchoTemplate 实现 echo.Renderer 接口
type EchoTemplate struct {
	templates *template.Template
}

// GetRenderTemplate 获取APP的渲染模板
func GetRenderTemplate() *EchoTemplate {
	// 创建模板对象并注册自定义模板函数
	tpl := template.New("index.tmpl").Funcs(template.FuncMap{
		"include": func(filename string, data interface{}) (template.HTML, error) {
			tmpl, err := template.ParseFiles(GetSpecialViewPath(filename))
			if err != nil {
				return "", err
			}
			var result bytes.Buffer
			err = tmpl.ExecuteTemplate(&result, filename, data)
			if err != nil {
				return "", err
			}
			return template.HTML(result.String()), nil
		},
	})
	// 解析模板文件
	tpl = template.Must(tpl.ParseFiles(GetSpecialViewPathList("index.tmpl", "wisdom.tmpl")...))
	tpl = template.Must(tpl.ParseGlob(GetSpecialViewPath("main/*.tmpl")))
	tpl = template.Must(tpl.ParseGlob(GetSpecialViewPath("partial/*.tmpl")))
	tplRender := &EchoTemplate{
		templates: tpl,
	}
	return tplRender
}
