package http

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/antony/polling/polling/model"
	"github.com/antony/polling/polling/usecase"
	"github.com/antony/polling/util"
	"github.com/julienschmidt/httprouter"
)

type httpPollingHandler struct {
	PUsecase usecase.PollingUsecase
}

func (hph *httpPollingHandler) GetAll(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	util.SetWriterJSON(w)

	pollings, err := hph.PUsecase.GetAll()
	if err != nil {
		log.Println(err)
		w.Write([]byte("Internal server error"))
	}

	res, err := json.Marshal(pollings)
	if err != nil {
		log.Println(err)
		w.Write([]byte("Internal server error"))
	}
	w.Write(res)
}

func (hph *httpPollingHandler) GetByRoomID(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	util.SetWriterJSON(w)

	//extract request value
	sRoomID := p.ByName("id")
	roomID, err := strconv.Atoi(sRoomID)
	if err != nil {
		log.Println(err)
		w.Write([]byte("Internal server error"))
	}

	//get polling by room id by its usecase
	polling, err := hph.PUsecase.GetByRoomID(int64(roomID))
	if err != nil {
		log.Println(err)
		w.Write([]byte("Internal server error"))
	}

	res, err := json.Marshal(polling)
	if err != nil {
		log.Println(err)
		w.Write([]byte("Internal server error"))
	}
	w.Write(res)
}

func (hph *httpPollingHandler) Insert(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	util.SetWriterJSON(w)

	//decode request body to it's model
	var polling model.Polling

	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&polling)
}

func NewPollingHttpHandler(router *httprouter.Router, pu usecase.PollingUsecase) {
	handler := httpPollingHandler{
		pu,
	}
	router.GET("/polling", handler.GetAll)
	router.GET("/room/:id/polling", handler.GetByRoomID)
	router.POST("/polling", handler.Insert)
}
