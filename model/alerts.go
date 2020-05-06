package model

import "time"

type AlertsDetails struct {
	AlertId           int `json:"alert_id"`
	ClientId          int `json:"client_id"`
	ConseillerId int    `json:"conseiller_id"`
	Date        time.Time `json:"date"`
	Motif          string    `json:"motif"`
}