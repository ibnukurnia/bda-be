package app

import (
	"fmt"
	"msbda/pkg/e"
	"strconv"

	"github.com/gin-gonic/gin"
)

type g struct {
	ginCtx  *gin.Context
	queries map[string]string
}

func Context(ctx *gin.Context) *g {
	return &g{
		ginCtx: ctx,
	}
}

func (g *g) WithQueries(keys ...string) *g {
	if g.queries == nil {
		g.queries = make(map[string]string)
	}

	for _, key := range keys {
		g.queries[key] = g.ginCtx.Query(key)
	}

	return g
}

func (g *g) Response(status int, data any) *response {
	return &response{
		Status: status,
		Data:   data,
		ctx:    g.ginCtx,
	}
}

func (g *g) GetQueryStringOrDefault(key string, d string) string {
	if val, exist := g.queries[key]; exist && val != "" {
		return val
	}

	return d
}

func (g *g) GetQueryIntOrDefault(key string, d int) int {
	if val, exist := g.queries[key]; exist && val != "" {
		n, err := strconv.Atoi(val)

		if err == nil {
			return n
		}
	}

	return d
}

func (g *g) GetParamInt(key string) (int, bool) {
	p := g.ginCtx.Param(key)

	n, err := strconv.Atoi(p)
	if err != nil {
		g.Response(500, nil).
			WithMessage(fmt.Sprintf("failed to parse %s to int", p)).
			Send()

		return 0, false
	}

	return n, true
}

func (g *g) BindJson(request RequestValidated) bool {
	return g.ParseError(validateJson(g.ginCtx, request))
}

func (g *g) ParseError(err error) bool {
	if err != nil {
		if errSchema, ok := err.(*e.SchemaError); ok {
			g.Response(errSchema.Status, nil).
				WithMessage(err.Error()).
				Send()

		} else {
			g.Response(500, nil).
				WithMessage(err.Error()).
				Send()
		}

		return false
	}

	return true
}
