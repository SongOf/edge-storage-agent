package handler

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestVolumeHandler(t *testing.T) {
	r, err := http.NewRequest("GET", "http://localhost:36000/volume/stat?seconds=1", nil)
	assert.Nil(t, err)

	w := httptest.NewRecorder()
	SystemHandler(w, r)

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	stats := string(body)
	fmt.Println(stats)
}

func TestVolumeHandlerHttp(t *testing.T) {
	params := url.Values{}
	Url, _ := url.Parse("http://localhost:36000/volume/stat")
	params.Set("seconds", "1")
	params.Set("path", "/etc")
	//如果参数中有中文参数,这个方法会进行URLEncode
	Url.RawQuery = params.Encode()
	urlPath := Url.String()
	fmt.Println(urlPath)
	resp, _ := http.Get(urlPath)
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}
