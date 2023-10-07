package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type responseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (w *responseWriter) WriteHeader(statusCode int) {
	w.statusCode = statusCode
	w.ResponseWriter.WriteHeader(statusCode)
}

type Notification struct {
	Message string `json:"message"`
	Type    string `json:"type"`
}

func AddNotification(w http.ResponseWriter, r *http.Request, notificationMessage, notificationType string) responseWriter {
	// create a new ResponseWriter that wraps the existing one
	newWritter := &responseWriter{w, http.StatusOK}
	// get previous notifications
	// TODO
	newNotification := map[string][]Notification{
		"messages": []Notification{
			Notification{
				Message: notificationMessage,
				Type:    notificationType,
			},
		},
	}
	// convert new notification to json
	jsonNotification, err := json.Marshal(newNotification)
	if err != nil {
		log.Printf("Error marshalling notification: %s", err)
	}
	// add new notification as header value
	newWritter.Header().Set("Hx-Trigger", string(jsonNotification))

	return *newWritter

}
