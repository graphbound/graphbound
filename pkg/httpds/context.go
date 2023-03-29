package httpds

import (
	"net/http"
)

// Context is the core of data sources. It allows us to pass variables between
// handlers and manage the flow of a request
type Context struct {
	fullPath string
	Request  *http.Request
	Response *http.Response
	Error    error
	handlers []Plugin
	index    int8
}

// NewContext creates a new context and initializes the handler chain
func NewContext(handlers ...Plugin) *Context {
	return &Context{
		handlers: handlers,
		index:    0,
	}
}

// FullPath returns the request templated path
func (c Context) FullPath() string {
	return c.fullPath
}

// Use adds a new handler to the handler chain
func (c *Context) Use(handlers ...Plugin) {
	c.handlers = append(c.handlers, handlers...)
}

// Next runs the next handler in the handler chain
func (c *Context) Next() {
	c.index++
	if c.index < int8(len(c.handlers)) {
		c.handlers[c.index](c)
	}
}

// Handle runs the main handler in the handler chain
func (c *Context) Handle() {
	if length := len(c.handlers); length > 0 {
		c.handlers[0](c)
	}
}
