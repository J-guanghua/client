package http

import (
	"net/http"
	"regexp"
)

type Response struct {
	body []byte
	*http.Response
}

func toSnakeCase(camelCase string) string {
	re := regexp.MustCompile(`([a-z0-9])([A-Z])`)
	snakeCase := re.ReplaceAllString(camelCase, "${1}_${2}")
	return snakeCase
}

func NewResquest(method,url string,msg interface{}) (*http.Request,error) {
	return http.NewRequest(method,url,nil)
}