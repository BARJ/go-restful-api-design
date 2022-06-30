package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

type RouterConfig struct {
	Host string
	Port int
}

type Route struct {
	Path   string
	Method string
	handle httprouter.Handle
}

type Handler interface {
	GetRoutes() []Route
}

func newRouter(handlers ...Handler) http.Handler {
	router := httprouter.New()
	for _, h := range handlers {
		for _, r := range h.GetRoutes() {
			router.Handle(r.Method, r.Path, r.handle)
		}
	}
	return router
}

type QueryParameters struct {
	values url.Values
}

func (qp QueryParameters) GetString(key string) string {
	return strings.TrimSpace(qp.values.Get(key))
}

func (qp QueryParameters) GetStringSlice(key string) []string {
	var ss []string
	if qp.values == nil {
		return ss
	}
	for _, v := range qp.values[key] {
		s := strings.TrimSpace(v)
		if len(s) > 0 {
			ss = append(ss, s)
		}
	}
	return ss
}

func (qp QueryParameters) GetInt(key string) (int, error) {
	s := qp.GetString(key)
	if len(s) == 0 {
		return 0, nil
	}
	i, err := strconv.Atoi(s)
	if err != nil {
		return 0, fmt.Errorf("QueryParameters.GetInt: %w", err)
	}
	return i, nil
}

func (qp QueryParameters) GetSortBy(key string) (string, SortOrder, error) {
	s := qp.GetString(key)
	if len(s) == 0 {
		return "", OrderUndefined, nil
	}
	ss := strings.Split(s, ".")
	if len(ss) != 2 || len(ss[0]) == 0 || (ss[1] != "asc" && ss[1] != "desc") {
		return "", OrderUndefined, fmt.Errorf("QueryParameters.GetSortBy: invalid format: expect \"{field name}.{asc}/{desc}\"")
	}
	sortOrder := OrderAscending
	if ss[1] == "desc" {
		sortOrder = OrderDescending
	}
	return ss[0], sortOrder, nil
}
