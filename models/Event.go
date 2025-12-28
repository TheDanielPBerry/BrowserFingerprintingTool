package models

import (
	"errors"
	"log"
	"time"
)

type Event struct {
	ID int64
	IPv4 string 
	IPv6 string
	Timezone string `json:"timezone"`
	ScreenWidth int `json:"screen_width"`
	ScreenHeight int `json:"screen_height"`
	NumCPUCores int `json:"num_cpu_cores"`
	Language string `json:"language"`
	UserAgent string `json:"user_agent"`
	CanvasHash string `json:"canvas_hash"`
	Genesis time.Time
	Mutated time.Time
};

func InsertTracking(event Event) (int, error) {
	db := GetDBContext()
	
	sql:= `INSERT INTO Event
		(
			IPv4,
			IPv6,
			Timezone,
			ScreenWidth,
			ScreenHeight,
			NumCPUCores,
			Language,
			UserAgent,
			CanvasHash
		)
		VALUES
		(?,?,?,?,?,?,?,?,?)`
	
	stmt, err := db.Preparex(sql)
	if err != nil {
		log.Panic(err)
		return 0, errors.New("Fatal Error")
	}

	result, err := stmt.Exec(
		event.IPv4,
		event.IPv6,
		event.Timezone,
		event.ScreenWidth,
		event.ScreenHeight,
		event.NumCPUCores,
		event.Language,
		event.UserAgent,
		event.CanvasHash,
	)

	if err != nil {
		log.Print(err)
		return 0, errors.New("Could not insert event")
	}
	id, err := result.LastInsertId()
	return int(id), err

}
