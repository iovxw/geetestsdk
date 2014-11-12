// Copyright 2014 Bluek404. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package geetestsdk

import (
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

var GeeAPIURL = "http://api.geetest.com/validate.php"

func New(k string) *GeetestSDK { return &GeetestSDK{key: k} }

type GeetestSDK struct {
	key string
}

func (g *GeetestSDK) Validate(challenge, validate, seccode string) (bool, error) {
	hash := fmt.Sprintf("%x", md5.Sum([]byte(g.key+"geetest"+challenge)))
	if validate == hash {
		resp, err := http.PostForm(GeeAPIURL, url.Values{"seccode": []string{seccode}})
		if err != nil {
			return false, err
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return false, err
		}

		if string(body) != fmt.Sprintf("%x", md5.Sum([]byte(seccode))) {
			return false, nil
		}
	} else {
		return false, nil
	}
	return true, nil
}
