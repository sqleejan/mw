package chaotic

import "net/http"

// Handler installs its own http routes, and returns
// http.Handler with a potentially chaotic behaviour
func Handler(url string) func(h http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		p := Policy{}
		p.mux = mux(url, h, p)
		return &p
	}
}

func mux(url string, h http.Handler, p Policy) http.Handler {
	mux := http.NewServeMux()
	mux.Handle(url+"/policy", &policyAPI{&p})
	mux.Handle(url+"/log", &p.Log)
	mux.Handle(url+"/", assets(url))
	mux.Handle("/", &p)

	return mux
}
