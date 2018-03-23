package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	var (
		resp    *http.Response
		cookies []*http.Cookie
		err     error
		body    []byte
	)
	if resp, err = http.Get("http://www.baidu.com"); err != nil {
		fmt.Printf("http get error and the error is  (%+v)", err)
		return
	}
	if resp.StatusCode != 200 {
		fmt.Printf(" resp.StatusCode is not ok, the code is %d", resp.StatusCode)
		return
	}

	// Cookies opts
	if cookies = resp.Cookies(); err != nil {
		fmt.Printf("http get error and the error is  (%+v)", err)
		return
	}

	for i := 0; i < len(cookies); i++ {
		ck := cookies[i]
		fmt.Printf("cookie is (%+v)\n", ck)
	}
	defer resp.Body.Close()
	if body, err = ioutil.ReadAll(resp.Body); err != nil {
		fmt.Printf("io read  error  (%+v)", err)
	}
	// this will may be marshall an unmarashall
	if len(body) == 0 {
		fmt.Println("read body is nil")
		return
	}
	fmt.Println("body is ", string(body))
}
