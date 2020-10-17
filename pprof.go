package pprof

import (
	"net/http"

	"github.com/arl/statsviz"
	"github.com/vmihailenco/treemux"
)

const (
	defaultPrefix string = "/debug/statsviz"
)

// RouterstatsvizRegister for treemux
func RouterStatsvizRegister(rg *treemux.TreeMux) {
	prefixRouter := rg.NewGroup(defaultPrefix)
	prefixRouter.GET("/", pprofHandler(statsviz.Index.ServeHTTP))
	prefixRouter.GET("/*path", pprofHandler(statsviz.Index.ServeHTTP))
	prefixRouter.GET("/ws", pprofHandler(statsviz.Ws))
}

func pprofHandler(h http.HandlerFunc) treemux.HandlerFunc {
	handler := http.HandlerFunc(h)
	return func(w http.ResponseWriter, req treemux.Request) error {
		handler.ServeHTTP(w, req.Request)
		return nil
	}
}
