package main

import (
	"fmt"
	"io/ioutil"
	"log"
	gin "mini-gin/gin"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/go-playground/assert/v2"
)

func pong(c *gin.Context) {
	fmt.Println("Response successful！", time.Now().Format("2006-01-02 15:04:05"))
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
}
