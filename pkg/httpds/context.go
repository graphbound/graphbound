package httpds

import (
	"context"
	"net/http"
)

type Context struct {
	Request  *http.Request
	Response *http.Response
	Error    error
	handlers []Plugin
	index    int8
	Ctx      context.Context
}

func NewContext(handlers ...Plugin) *Context {
	return &Context{
		handlers: handlers,
		index:    -1,
	}
}

func (c *Context) Use(handlers ...Plugin) {
	c.handlers = append(c.handlers, handlers...)
}

func (c *Context) Next() {
	c.index++
	if c.index < int8(len(c.handlers)) {
		c.handlers[c.index](c)
		c.index++
	}
}

func (c *Context) Handle() {
	if length := len(c.handlers); length > 0 {
		c.handlers[length-1](c)
	}
}
