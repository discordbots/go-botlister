package botlister

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"
)

// API the struct of the API instance containing the tokens and http client.
type API struct {
	BotToken  string
	UserToken string
	Client    *http.Client
}

// NewAPI creates a new instance of API with the provided tokens.
func NewAPI(BotToken string, UserToken string) *API {
	return &API{
		BotToken:  BotToken,
		UserToken: UserToken,
		Client:    new(http.Client),
	}
}

func (a *API) doRequest(method string, user bool, endpoint string, body interface{}, output interface{}) error {

	var reader io.Reader = nil

	if body != nil {
		data, err := json.Marshal(body)
		if err != nil {
			return err
		}

		reader = bytes.NewBuffer(data)
	}

	req, err := http.NewRequest(method, endpoint, reader)
	if err != nil {
		return err
	}

	req.Header.Set("User-Agent", req.Header.Get("User-Agent")+" go-botlister")
	if user {
		req.Header.Set("Authorization", "Bearer "+a.UserToken)
	} else {
		req.Header.Set("Authorization", "Bot "+a.BotToken)
	}

	resp, err := a.Client.Do(req)
	if err != nil {
		return err
	} else if resp.StatusCode > 299 || resp.StatusCode < 200 {
		return errors.New("invalid response: " + strconv.Itoa(resp.StatusCode) + " (" + resp.Status + ")")
	}

	defer resp.Body.Close()
	if output != nil {
		decoder := json.NewDecoder(resp.Body)
		return decoder.Decode(output)
	}

	return nil
}
