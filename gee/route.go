package gee

import (
	"net/http"
	"strings"
)

// 路由器
type HandleFunc func(*Context)
type route struct {
	router map[string]HandleFunc
	roots  map[string]*node
}

func NewRoute() *route {
	return &route{
		router: make(map[string]HandleFunc, 10),
		roots:  make(map[string]*node, 4),
	}
}

func (r *route) parsePattern(pattern string) (parts []string) {
	vs := strings.Split(pattern, "/")
	for _, item := range vs {
		if item != "" {
			parts = append(parts, item)
			if item[0] == '*' {
				break
			}
		}
	}
	return parts
}

func (r *route) addRoute(method string, pattern string, handle HandleFunc) {
	key := method + "_" + pattern
	_, ok := r.roots[method]
	if !ok {
		r.roots[method] = &node{}
	}
	parts := r.parsePattern(pattern)
	r.roots[method].insert(pattern, parts, 0)
	r.router[key] = handle
}

func (r *route) getRouter(method string, path string) (*node, map[string]string) {
	searchParts := r.parsePattern(path)
	params := make(map[string]string, 1)
	root, ok := r.roots[method]
	if !ok {
		return nil, nil
	}
	n := root.search(searchParts, 0)
	if n != nil {
		parts := r.parsePattern(n.pattern)
		for index, part := range parts {
			if part[0] == ':' {
				params[part[1:]] = searchParts[index]
			}
			if part[0] == '*' && len(part) > 1 {
				params[part[1:]] = strings.Join(searchParts[index:], "/")
				break
			}
		}
	}
	return n, params
}

func (r *route) handle(c *Context) {
	n, params := r.getRouter(c.Method, c.Path)
	if n != nil {
		key := c.Method + "_" + n.pattern
		c.Params = params
		c.handlers = append(c.handlers, r.router[key])
	} else {
		c.handlers = append(c.handlers, func(ctx *Context) {
			c.String(http.StatusOK, "404 NOT FOUND")
		})

	}
	c.Next()
}
