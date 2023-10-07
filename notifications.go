package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
)

type responseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (w *responseWriter) WriteHeader(statusCode int) {
	w.statusCode = statusCode
	w.ResponseWriter.WriteHeader(statusCode)
}

func (w *responseWriter) Write(b []byte) (int, error) {
	if w.statusCode == 0 {
		w.statusCode = http.StatusOK
	}
	return w.ResponseWriter.Write(b)
}

func AddNotification(w http.ResponseWriter, r *http.Request, notificationMessage, notificationType string) http.ResponseWriter {
	// create a new ResponseWriter that wraps the existing one
	newWritter := &responseWriter{
		ResponseWriter: w,
		statusCode:     http.StatusOK,
	}

	hxTriggerVal := make(map[string]interface{})

	// get current header value
	currentHxTriggerVal := w.Header().Get("Hx-Trigger")

	if strings.HasPrefix(currentHxTriggerVal, "{") {
		// convert current header value to map
		err := json.Unmarshal([]byte(currentHxTriggerVal), &hxTriggerVal)
		if err != nil {
			log.Printf("Error unmarshalling hx-trigger header: %s", err)
		}
	} else if currentHxTriggerVal != "" {
		// convert current header value to map
		hxTriggerVal = map[string]interface{}{
			currentHxTriggerVal: true,
		}
	}

	// check if notifications key exists
	notifications, ok := hxTriggerVal["notifications"].([]interface{})
	// if not, create it
	if !ok {
		notifications = []interface{}{}
	}
	// append new notification
	notifications = append(notifications, map[string]string{
		"message": notificationMessage,
		"type":    notificationType,
	})
	// add notifications to hxTriggerVal
	hxTriggerVal["notifications"] = notifications

	// convert hxTriggerVal to json
	hxTriggerValJson, err := json.Marshal(hxTriggerVal)
	if err != nil {
		log.Printf("Error marshalling hx-trigger header: %s", err)
	}

	// add hxTriggerVal as header value
	newWritter.Header().Set("Hx-Trigger", string(hxTriggerValJson))
	log.Printf("New Notification added: type: %s", notificationType)
	return newWritter
}
