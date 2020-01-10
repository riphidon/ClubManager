package utils

import (
	"fmt"
	"net/http"
	"net/url"
	"path"
	"strings"
)

func CatchURLData(r *http.Request, d string) string {
	data := r.URL.Query().Get(d)
	return data
}

func DoURL(path string, segment, param string) string {
	baseURL, err := url.Parse(path)
	if err != nil {
		fmt.Println("Non valid URL: ", err.Error())
	}
	baseURL.Path += segment

	params := url.Values{}
	params.Add("q", param)
	baseURL.RawQuery = params.Encode()

	return baseURL.String()
}

//returns the part of the url at the requested depth, 0 = first part after "/""
func SplitPath(url string, depth int) string {
	parts := strings.Split(path.Clean(url)[1:], "/")
	if len(parts) >= depth+1 {
		return parts[depth]
	}
	return ""
}

func CatchUserErr(code int) string {
	errMessage := ""
	switch code {
	case 1:
		errMessage = "Invalid passphrase"
		return errMessage
	case 2:
		errMessage = "Email already in use"
		return errMessage
	case 3:
		errMessage = "Unknown User"
		return errMessage
	case 4:
		errMessage = "Invalid password"
		return errMessage
	default:
		return ""

	}
}

func CheckState(id string) bool {
	if id == "" {
		return false
	}
	return true
}
