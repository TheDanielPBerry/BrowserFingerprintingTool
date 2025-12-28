package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
	"tracker/models"
)

type Result struct {
	Success bool
	Message string
};

//True if the string is a valid sha256 hash
func isSHA256(hash string) bool {
	count := 0
	for _, char := range hash {
		count++
		if (char >= '0' && char <= '9') {
			continue
		} else if (char >= 'A' && char <= 'F') {
			continue
		} else if (char >= 'a' && char <= 'f') {
			continue
		} else {
			return false
		}
	}
	if count != 64 {
		return false
	}
	return true;
}

func isValidTimeZone(tz string) bool {
	_, err := time.LoadLocation(tz)
	return err == nil
}

func validate(event *models.Event) {
	if !isSHA256(event.CanvasHash) {
		event.CanvasHash = ""
	}
	if !isValidTimeZone(event.Timezone) {
		event.Timezone = ""
	}
	if len(event.Language) > 8 {
		event.Language = ""
	}
}

func Stat(w http.ResponseWriter, req *http.Request) {
	var event models.Event
	decoder := json.NewDecoder(req.Body)
	decoder.Decode(&event)
	log.Print(event)
	validate(&event)

	event.IPv4 = GetClientIP(req)
	event.IPv6 = GetClientIP(req)


	var apiResult *Result = new(Result);

	id, err := models.InsertTracking(event)
	if err != nil {
		apiResult.Success = false
		apiResult.Message = "Issue with stat"
	} else {
		apiResult.Success = true
		apiResult.Message = fmt.Sprintf("Record %d successfully created", id)
	}
	json.NewEncoder(w).Encode(apiResult)
}
