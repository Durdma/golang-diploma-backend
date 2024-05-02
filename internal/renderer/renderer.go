package renderer

import (
	"context"
	"github.com/a-h/templ"
	"github.com/gin-gonic/gin/render"
	"net/http"
)

var Default = &HTMLTemplRenderer{}

type HTMLTemplRenderer struct {
	FallbackHTMLRenderer render.HTMLRender
}

func (r *HTMLTemplRenderer) Instance(s string, d any) render.Render {
	templData, ok := d.(templ.Component)
	if !ok {
		if r.FallbackHTMLRenderer != nil {
			return r.FallbackHTMLRenderer.Instance(s, d)
		}
	}

	return &Renderer{
		Ctx:       context.Background(),
		Status:    -1,
		Component: templData,
	}
}

func New(ctx context.Context, status int, component templ.Component) *Renderer {
	return &Renderer{
		Ctx:       ctx,
		Status:    status,
		Component: component,
	}
}

type Renderer struct {
	Ctx       context.Context
	Status    int
	Component templ.Component
}

func (r Renderer) Render(w http.ResponseWriter) error {
	r.WriteContentType(w)
	if r.Status != -1 {
		w.WriteHeader(r.Status)
	}

	if r.Component != nil {
		return r.Component.Render(r.Ctx, w)
	}

	return nil
}

func (r Renderer) WriteContentType(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
}
