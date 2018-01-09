package httputil

import (
	"encoding/json"
	"log"
	"net/http"
)

type response struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func generateJSONResponse(message string, data interface{}) []byte {
	resp := response{
		message,
		data,
	}
	bs, err := json.Marshal(resp)
	if err != nil {
		log.Println(err)
		return nil
	}
	return bs
}

func SetWriterJSON(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
}

func WriteErrorResponse(w http.ResponseWriter, code int, message string) {
	w.WriteHeader(code)
	w.Write(generateJSONResponse(message, nil))
}

func WriteSuccessResponse(w http.ResponseWriter, message string, data interface{}) {
	w.WriteHeader(200)
	w.Write(generateJSONResponse(message, data))
}
