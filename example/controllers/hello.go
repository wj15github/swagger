package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/leesper/holmes"
	"github.com/wj15github/swagger/example/models"
)

// @Title Hello1
// @Description test API - 1
// @Accept  json
// @Success 200 {object}	models.CodeMessage
// @Resource /1.0
// @Router /1.0/hello1 [get]
func Hello1(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	message := &models.CodeMessage{
		Code: 0,
		Msg:  "hello1",
	}
	if err := json.NewEncoder(w).Encode(message); err != nil {
		holmes.Errorln(err)
	}
}

// @Title Hello2
// @Description test API - 2
// @Accept  json
// @Success 200 {object}	models.CodeMessage
// @Resource /1.0
// @Router /1.0/hello2 [get]
func Hello2(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	message := &models.CodeMessage{
		Code: 0,
		Msg:  "hello2",
	}
	if err := json.NewEncoder(w).Encode(message); err != nil {
		holmes.Errorln(err)
	}
}
