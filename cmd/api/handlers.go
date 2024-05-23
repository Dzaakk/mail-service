package main

import "net/http"

func (app *Config) SendMail(w http.ResponseWriter, r *http.Request) {
	type mailMessage struct {
		From    string `json:"from"`
		To      string `json:"to"`
		Subject string `json:"subject"`
		Message string `json:"message"`
	}

	var requestPaylooad mailMessage

	err := app.readJSON(w, r, &requestPaylooad)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	msg := Message{
		From:    requestPaylooad.From,
		To:      requestPaylooad.From,
		Subject: requestPaylooad.Subject,
		Data:    requestPaylooad.Message,
	}

	err = app.Mailer.SendSMTPMessage(msg)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	payload := jsonResponse{
		Error:   false,
		Message: "sent to " + requestPaylooad.To,
	}

	app.writeJSON(w, http.StatusAccepted, payload)
}
