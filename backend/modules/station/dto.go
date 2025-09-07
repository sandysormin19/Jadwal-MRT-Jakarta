package station

import "time"

type Station struct{
	Id string `json:"nid"`
	Name string `json:"title"`
}
type StationResponse struct{
	Id string `json:"nid"`
	Name string `json:"title"`
}

type Schedule struct {
	StationId          string `json:"nid"`
	StationName        string `json:"title"`
	ScheduleBundaranHI string `json:"jadwal_hi_biasa"`
	ScheduleLebakBulus string `json:"jadwal_lb_biasa"`
}

type ScheduleResponse struct {
	StationName string `json:"station"`
	Time        string `json:"time"`
}

type ScheduleStationTime struct {
	StationName string
	Time        []time.Time
}
