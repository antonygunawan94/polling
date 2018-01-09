package http

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/antony/polling/httputil"
	"github.com/antony/polling/polling/model"
	"github.com/antony/polling/polling/usecase"
	"github.com/julienschmidt/httprouter"
)

type httpPollingHandler struct {
	PUsecase *usecase.PollingUsecase
}

func (hph *httpPollingHandler) GetAll(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	httputil.SetWriterJSON(w)

	pollings, err := hph.PUsecase.GetAll()
	if err != nil {
		log.Println(err)
		httputil.WriteErrorResponse(w, 500, "Internal server error")
		return
	}

	res, err := json.Marshal(pollings)
	if err != nil {
		log.Println(err)
		httputil.WriteErrorResponse(w, 500, "Internal server error")
		return
	}
	w.Write(res)
}

func (hph *httpPollingHandler) GetByID(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	httputil.SetWriterJSON(w)

	//extract request value
	sID := p.ByName("id")
	ID, err := strconv.Atoi(sID)
	if err != nil {
		log.Println(err)
		httputil.WriteErrorResponse(w, 500, "Internal server error")
		return
	}

	//get polling by room id by its usecase
	polling, err := hph.PUsecase.GetByID(int64(ID))
	if err != nil {
		log.Println(err)
		httputil.WriteErrorResponse(w, 500, "Internal server error")
		return
	}
	httputil.WriteSuccessResponse(w, fmt.Sprintf("Polling with id %v", ID), polling)
}

func (hph *httpPollingHandler) GetByRoomID(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	httputil.SetWriterJSON(w)

	//extract request value
	sRoomID := p.ByName("id")
	roomID, err := strconv.Atoi(sRoomID)
	if err != nil {
		log.Println(err)
		httputil.WriteErrorResponse(w, 500, "Internal server error")
		return
	}

	//get polling by room id by its usecase
	polling, err := hph.PUsecase.GetByRoomID(int64(roomID))
	if err != nil {
		log.Println(err)
		httputil.WriteErrorResponse(w, 500, "Internal server error")
		return
	}

	/* res, err := json.Marshal(polling)
	if err != nil {
		log.Println(err)
		httputil.WriteErrorResponse(w, 500, "Internal server error")
		return
	} */
	httputil.WriteSuccessResponse(w, fmt.Sprintf("Polling with room id %v", roomID), polling)
}

func (hph *httpPollingHandler) GetPollingDetailByID(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	httputil.SetWriterJSON(w)

	//extract request value
	sID := p.ByName("id")
	ID, err := strconv.Atoi(sID)
	if err != nil {
		log.Println(err)
		httputil.WriteErrorResponse(w, 500, "Internal server error")
		return
	}

	polling, puas, err := hph.PUsecase.GetPollingDetailByID(int64(ID))
	if err != nil {
		log.Println(err)
		httputil.WriteErrorResponse(w, 500, "Internal server error")
		return
	}

	data := make(map[string]interface{})
	data["polling"] = polling
	data["polling_user_answers"] = puas
	httputil.WriteSuccessResponse(w, "Success getting detail of polling ID "+sID, data)
}

func (hph *httpPollingHandler) Create(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	httputil.SetWriterJSON(w)

	//decode request body to it's model
	var polling model.Polling

	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&polling)

	err := hph.PUsecase.Insert(&polling)
	if err != nil {
		log.Println(err)
		httputil.WriteErrorResponse(w, 500, "Internal server error")
		return
	}

	httputil.WriteSuccessResponse(w, "Success inserting new polling", nil)
}

func NewPollingHttpHandler(router *httprouter.Router, pu *usecase.PollingUsecase) {
	handler := httpPollingHandler{
		pu,
	}
	router.GET("/polling", handler.GetAll)
	router.GET("/polling/:id", handler.GetPollingDetailByID)
	router.GET("/room/:id/polling", handler.GetByRoomID)
	router.POST("/polling", handler.Create)
}
