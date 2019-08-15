package infobip

import (
	"bytes"
	"encoding/json"
	"github.com/pkg/errors"
	"net/http"
)

type authentication struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type senderRequest struct {
	Authentication authentication  `json:"authentication"`
	Messages       json.RawMessage `json:"messages"`
}

type client struct {
	url string
	userName string
	password string
	httpClient *http.Client
}

func New(url, userName, password string, httpClient *http.Client) Sender {
	return &client{
		url:        url,
		httpClient: httpClient,
		userName:   userName,
		password:   password,
	}
}

func (c *client) Send(messages []Message) (sendResponse *Response, err error) {
	messagesJSON, err := json.Marshal(messages)

	if err != nil {
		return nil, errors.Wrapf(err, "couldn't marshal messages")
	}

	body := senderRequest{
		Authentication: authentication{
			Username: c.userName,
			Password: c.password,
		},
		Messages: messagesJSON,
	}

	bodyJSON, err := json.Marshal(body)
	if err != nil {
		return nil, errors.Wrapf(err, "couldn't marshal request body")
	}

	req, err := http.NewRequest(http.MethodPost, c.url, bytes.NewReader(bodyJSON))
	if err != nil {
		return nil, errors.Wrapf(err, "could not init new request")
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, errors.Wrapf(err, "couldn't send request")
	}

	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&sendResponse)

	return
}