package api

import (
	"errors"
	"net/http"
)

var (
	ErrInvalidStatusCode = errors.New("invalid status code")
)

// GetRedirect отправляет запрос и возвращает URL из заголовка Location (куда нас редиректят)
func GetRedirect(url string) (string, error) {
	const op = "api.GetRedirect"

	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse // Говорим клиенту НЕ идти по редиректу, а остановиться
		},
	}

	resp, err := client.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusFound && resp.StatusCode != http.StatusMovedPermanently {
		return "", ErrInvalidStatusCode
	}

	return resp.Header.Get("Location"), nil
}
