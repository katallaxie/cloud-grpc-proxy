package shared

import (
	"bytes"
	"text/template"

	"github.com/katallaxie/protoc-gen-cloud-proxy/api"

	pgs "github.com/lyft/protoc-gen-star"
)

type ProxyContext struct {
	Methods *api.Methods
	Method  pgs.Method

	Typ string
}

func Render(tpl *template.Template) func(ctx ProxyContext) (string, error) {
	return func(ctx ProxyContext) (string, error) {
		buf := &bytes.Buffer{}
		err := tpl.ExecuteTemplate(buf, ctx.Typ, ctx)
		return buf.String(), err
	}
}
