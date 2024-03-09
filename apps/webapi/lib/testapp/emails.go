package testapp

import (
	"github.com/go-resty/resty/v2"
	"provar.se/webapi/lib/random"
)

type NameAddress struct {
	Name    string `json:"Name"`
	Address string `json:"Address"`
}

// MockEmail represents an email sent to a user by the system
type MockEmail struct {
	ID      string        `json:"ID"`
	Subject string        `json:"Subject"`
	From    NameAddress   `json:"From"`
	To      []NameAddress `json:"To"`
	HTML    string
	Text    string
}

// RandomEmailAddress returns a random email address used for testing
func RandomEmailAddress() string {
	rnd := random.String(5)
	return "test.user." + rnd + "@provar.se"
}

// getHTMLContent returns the HTML content of an email from mailpit
func getHTMLContent(ID string) string {
	req := resty.New().R()
	res, err := req.Get("http://localhost:8025/view/" + ID + ".html")
	if err != nil {
		panic(err)
	}
	if res.StatusCode() != 200 {
		panic("unexpected status code")
	}
	responseBody := res.String()
	return responseBody
}

// getTextContent returns the text content of an email from mailpit
func getTextContent(ID string) string {
	req := resty.New().R()
	res, err := req.Get("http://localhost:8025/view/" + ID + ".txt")
	if err != nil {
		panic(err)
	}
	if res.StatusCode() != 200 {
		panic("unexpected status code")
	}
	responseBody := res.String()
	return responseBody
}

// GetSentEmails returns emails sent to the target from mailpit
func GetSentEmails(target string) []MockEmail {
	responseBody := struct {
		Messages []MockEmail `json:"messages"`
	}{}
	req := resty.New().R()
	req.ForceContentType("application/json")
	req.SetResult(&responseBody)
	res, err := req.Get("http://localhost:8025/api/v1/search?query=to:" + target)
	if err != nil {
		panic(err)
	}
	if res.StatusCode() != 200 {
		panic("unexpected status code")
	}
	for idx := range responseBody.Messages {
		msg := &responseBody.Messages[idx]
		msg.HTML = getHTMLContent(msg.ID)
		msg.Text = getTextContent(msg.ID)
	}
	return responseBody.Messages
}
