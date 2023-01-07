package outsource

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

type Response struct {
	Data    Data    `json:"data"`
	Support Support `json:"support"`
}
type Data struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Year         int    `json:"year"`
	Color        string `json:"color"`
	PantoneValue string `json:"pantone_value"`
}
type Support struct {
	URL  string `json:"url"`
	Text string `json:"text"`
}

func GreetInSpanish(ctx context.Context, id string) (string, error) {
	base := "https://reqres.in/api/unknown/%s"
	url := fmt.Sprintf(base, id)
	fmt.Println(url)

	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	var res Response
	err = json.Unmarshal(body, &res)
	if err != nil {
		log.Fatal(err)
	}
	status := resp.StatusCode
	if status >= 400 {
		message := fmt.Sprintf("HTTP Error %d: %s", status, string(body))
		return "", errors.New(message)
	}
	fmt.Println(res)
	return res.Data.Name, nil
}

func GreetInSpanish2(ctx context.Context, name string) (string, error) {
	base := "http://localhost:9999/get-spanish-greeting?name=%s"
	url := fmt.Sprintf(base, url.QueryEscape(name))

	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	translation := string(body)
	status := resp.StatusCode
	if status >= 400 {
		message := fmt.Sprintf("HTTP Error %d: %s", status, translation)
		return "", errors.New(message)
	}

	return translation, nil
}
