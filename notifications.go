package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
)

func AddNotification(w http.ResponseWriter, r *http.Request, notificationMessage, notificationType string) {
	// create map for hx-trigger header value
	hxTriggerVal := make(map[string]interface{})

	// get current header value
	currentHxTriggerVal := w.Header().Get("Hx-Trigger")

	if strings.HasPrefix(currentHxTriggerVal, "{") {
		// convert current header value to map
		err := json.Unmarshal([]byte(currentHxTriggerVal), &hxTriggerVal)
		if err != nil {
			log.Printf("Error unmarshalling hx-trigger header: %s, header value: %s", err, currentHxTriggerVal)
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
		log.Printf("Error marshalling HxTriggerVal: %s, value: %s", err, hxTriggerVal)
	}

	// add hxTriggerValJson as header value
	w.Header().Set("Hx-Trigger", string(hxTriggerValJson))
	log.Printf("New Notification added: type: %s", notificationType)
}
