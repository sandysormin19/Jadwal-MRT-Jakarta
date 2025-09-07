package station

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/sandysormin19/Jadwal-MRT/common/response"
)

func Initiate(router *gin.RouterGroup) {
	stationService := NewService()
	//routing grup
	station := router.Group("/stations")
	//routing
	station.GET("/", func(c *gin.Context) {
		GetAllStation(c, stationService)
	})
	station.GET("/:id", func(c *gin.Context) {
		CheckScheduleByStationId(c, stationService) 
	})
}
func GetAllStation(ctx *gin.Context, service Service) {
	datas, err := service.GetAllStation()
	if err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			response.APIResponse{
				Success: false,
				Messege: err.Error(),
				Data:    nil,
			},
		)
		return
	}
	ctx.JSON(
		http.StatusOK,
		response.APIResponse{
			Success: true,
			Messege: "Success Get All Station",
			Data:    datas,
		},
	)
}

func CheckScheduleByStationId(ctx *gin.Context, service Service) {
	id := ctx.Param("id")
	datas, err := service.CheckScheduleByStationId(id)
	if err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			response.APIResponse{
				Success: false,
				Messege: err.Error(),
				Data:    nil,
			},
		)
		return
	}
	ctx.JSON(
		http.StatusOK,
		response.APIResponse{
			Success: true,
			Messege: "Success Get Schedule By Station ID",
			Data:    datas,
		},
	)
}