package station

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/sandysormin19/Jadwal-MRT/common/client"
)

type Service interface {
	GetAllStation() (response []StationResponse, err error)
	CheckScheduleByStationId(id string) (response []ScheduleResponse, err error)
}

type service struct {
	client *http.Client
}


func (s *service) GetAllStation() (response []StationResponse, err error) {
	// Implementasi logika untuk mendapatkan semua stasiun MRT
	url := "https://jakartamrt.co.id/id/val/stasiuns"

	//hit url
	byteResponse, err := client.DoRequest(s.client, url)
	if err != nil {
		return
	}

	//data station
	var stations []Station
	if err := json.Unmarshal(byteResponse, &stations); err != nil {
		return nil, err
	}

	// Loop Byte Mashal Json Data
	for _, i := range stations {
		response = append(response, StationResponse{
			Id:   i.Id,
			Name: i.Name,
		})
	}

	return
}

func (s *service) CheckScheduleByStationId(id string) (response []ScheduleResponse, err error) {
	// Implementasi logika untuk mendapatkan jadwal berdasarkan ID stasiun
	url := "https://jakartamrt.co.id/id/val/stasiuns"

	//hit url
	byteResponse, err := client.DoRequest(s.client, url)
	if err != nil {
		return
	}

	//data schedule
	var schedule []Schedule
	if err := json.Unmarshal(byteResponse, &schedule); err != nil {
		return nil, err
	}


var scheduleSelected Schedule
for _, item := range schedule{
	if item.StationId == id {
		scheduleSelected = item
		break
		}
	}

	if scheduleSelected.StationId == "" {
		err = errors.New("station id not found")
		return
	}

response, err = ConvertDataToResponse(scheduleSelected)
if err != nil {
	return

}

return
}

func ConvertDataToResponse (schedule Schedule) (response[]ScheduleResponse, err error){
	var (
		LebakBulusTripName = "Stasiun Lebak Bulus Grab"
		BundaranHITripName = "Stasiun Bundaran HI Bank DKI"
	)
 scheduleLebakBulus := schedule.ScheduleLebakBulus
 scheduleBundaranHI := schedule.ScheduleBundaranHI

 scheduleLebakBulusParsed, err := ConvertScheduleToTimeFormat(scheduleLebakBulus)
 if err != nil {
	return
 }

scheduleBundaranHIParsed, err := ConvertScheduleToTimeFormat(scheduleBundaranHI)
if err != nil {
	return
	}

for _, i := range scheduleLebakBulusParsed{
	if i.Format("15:04") > time.Now().Format("15:04"){
		response = append(response, ScheduleResponse{
			StationName: LebakBulusTripName,
			Time: i.Format("15:04"),
		})
	}
}

for _, i := range scheduleBundaranHIParsed{
	if i.Format("15:04") > time.Now().Format("15:04"){
		response = append(response, ScheduleResponse{
			StationName: BundaranHITripName,
			Time: i.Format("15:04"),
		})
	}
}

return

	}

// Create Data For Station
func ConvertScheduleToTimeFormat(schedule string)(response []time.Time, err error){
	var (
		parsedTime time.Time
		schedules = strings.Split(schedule, ",")
	)
	for _, i := range schedules{
		trimmedTime := strings.TrimSpace(i)
		if trimmedTime == ""{
			continue
		}
		parsedTime, err = time.Parse("15:04", trimmedTime)
		if err != nil {
			err = errors.New("error parsing time" + trimmedTime)
		return
		}
		response = append(response, parsedTime)	
	}
return
}
// Create To Take Data With Same ID Request
func NewService() Service {
	return &service{
		client: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}