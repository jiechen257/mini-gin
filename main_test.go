package main

import (
	"fmt"
	"io/ioutil"
	"log"
	gin "mini-gin/gin"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-playground/assert/v2"
)

func pong(c *gin.Context) {
	fmt.Println("Hello Gin")
	c.String("%s", "pong")
}

func TestPingRoute(t *testing.T) {
	r := gin.New()
	r.AddRoute("GET", "/ping", pong)

	ts := httptest.NewServer(r)
	defer ts.Close()

	{
		res, err := http.Get(fmt.Sprintf("%s/ping", ts.URL))
		if err != nil {
			log.Println(err)
		}
		resp, _ := ioutil.ReadAll(res.Body)
		assert.Equal(t, http.StatusOK, res.StatusCode)
		assert.Equal(t, "pong", string(resp))
	}

	// w := httptest.NewRecorder()
	// req, _ := http.NewRequest("GET", "/ping", nil)
	// router.ServeHTTP(w, req)

	// assert.Equal(t, http.StatusOK, w.Code)
	// assert.Equal(t, "pong", w.Body.String())
}
