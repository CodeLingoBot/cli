package api

import (
	"net/http"
)

func AppsList() (*http.Response, error) {
	req := map[string]interface{}{
		"method":   "GET",
		"endpoint": "/api/apps",
	}
	return Do(req)
}

func AppsShow(app string) (*http.Response, error) {
	req := map[string]interface{}{
		"method":   "GET",
		"endpoint": "/api/apps/" + app,
	}
	return Do(req)
}

func AppsCreate(app string) (*http.Response, error) {
	req := map[string]interface{}{
		"method":   "POST",
		"endpoint": "/api/apps",
		"params": map[string]interface{}{
			"app": map[string]interface{}{
				"fullname": app,
			},
		},
	}
	return Do(req)
}
