package e2

const (
	GET = iota
	POST
	PUT
	DELETE
	CONNECTIBNG
	HEAD
	OPTIONS
	PATCH
	TRACE
)

type MethodMaps []handler

type handler map[string]HandlerFunc

func NewRouter() MethodMaps {
	return []handler{
		GET:    make(handler),
		POST:   make(handler),
		PUT:    make(handler),
		DELETE: make(handler),
	}
}

func (m MethodMaps) POST(path string, fn HandlerFunc) {
	if _, ok := m[POST][path]; ok {
		panic("duplicate url with Post method")
	}
	m[POST][path] = fn
}
