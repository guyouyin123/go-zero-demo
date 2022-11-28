package middleware

import (
	"fmt"
	"net/http"
)

type GlobalMiddleware struct {
}

//全局中间件
func NewGlobalMiddleware() *GlobalMiddleware {
	return &GlobalMiddleware{}
}

func (m *GlobalMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO generate middleware implement function, delete after code implementation

		fmt.Println("=======全局中间件之前=======")
		next(w, r)
		fmt.Println("=======全局中间件之后=======")
	}
}
