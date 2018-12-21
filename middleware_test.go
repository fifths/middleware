package middleware

import (
	"github.com/julienschmidt/httprouter"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRoutesRequests_Use(t *testing.T) {
	mw := New()
	rmw := func(handle httprouter.Handle) httprouter.Handle {
		return handle
	}
	l := len(mw.routeMiddleware)
	mw.Use(rmw)
	if len(mw.routeMiddleware) != l+1 {
		t.Error("Error of RoutesRequests_Use")
	}
}

func TestRoutesRequests_Handle(t *testing.T) {
	routesRequests := New()
	var middlewareCalled bool
	middleware := func(handle httprouter.Handle) httprouter.Handle {
		return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
			middlewareCalled = true
			handle(w, r, p)
		}
	}
	routesRequests.Use(middleware)
	req := httptest.NewRequest("GET", "/middleware", nil)
	w := httptest.NewRecorder()
	returnMsg := `{"errcode": 0}`
	hn := func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		io.WriteString(w, returnMsg)
	}
	httprouterHandler := func(handle httprouter.Handle) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			handle(w, r, httprouter.Params{})
		}
	}
	handler := http.HandlerFunc(
		httprouterHandler(routesRequests.Handle(hn)),
	)
	handler.ServeHTTP(w, req)
	if !middlewareCalled {
		t.Error("Error of middleware")
	}
	if status := w.Code; status != http.StatusOK {
		t.Errorf("Errof of status code: %v", status)
	}
	if w.Body.String() != returnMsg {
		t.Errorf("Errof of body: %v ", w.Body.String())
	}
}
