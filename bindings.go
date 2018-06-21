package yoo

import (
	"strconv"
	"fmt"
	"net/http"
)

func newConsole() Any {
	var log TSFunction = func (args *[]Any) Any {
		for _, v := range *args {
			fmt.Print(v, " ")
		}
		fmt.Println()
		return nil
	}
	return CreateObject(
		[]string{ "log" },
		[]Any{ &log },
	)
}

func newHttp() Any {
	var newContent ProxyFunction = func (object Any, t int, key string, value Any) Any {
		obj := *object.(*HttpContent)
		switch key {
		case "method":
			return obj.r.Method
		default:
			return nil
		}
	}
	newCtxPtr := &newContent
	var Server TSFunction = func (argsPtr *[]Any) Any {
		args := *argsPtr
		length := len(args)
		var server *http.Server
		if length == 0 {
			server = &http.Server {}
		} else if handlerPtr, ok := args[0].(*TSFunction); ok {
			handle := *handlerPtr
			var h http.HandlerFunc = func (w http.ResponseWriter, r *http.Request) {
				handle(&[]Any { &Proxy{ &HttpContent { &w, r }, newCtxPtr } })
			}
			port := 80
			if length > 1 {
				if p, ok := args[1].(float64); ok {
					port = int(p)
				}
			}
			server = &http.Server {
				Addr: "127.0.0.1:" + strconv.Itoa(port),
				Handler: h,
			}
		}
		var obj Any
		var listen TSFunction = func (args *[]Any) Any {
			server.ListenAndServe()
			return obj
		}
		obj = &Object4 {
			k0: "listen",
			v0: &listen,
		}
		return obj
	}
	return CreateObject(
		[]string{ "Server" },
		[]Any{ &Server },
	)
}

func GetBindings() *Variables {
	return &Variables {
		"net/http": newHttp(),
		"console": newConsole(),
	}
}

type HttpContent struct {
	w *http.ResponseWriter
	r *http.Request
}
