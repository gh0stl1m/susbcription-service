package healthcheks

import (
	"encoding/json"
	"net/http"
)

type LivenessResponse struct {
  Message string `json:"message,omitempty"`
}

func (hr *HealthCheckCtx) LivenessHandler(w http.ResponseWriter, r *http.Request) {

  responseMessage := LivenessResponse { Message: "Server Running" }
  w.Header().Set("Content-Type", "application/json")

  if err := hr.DB.Ping(); err != nil {

    hr.ErrorLog.Println("Something went wrong with database ping", err)

    responseMessage.Message = "Database connection failed"
  	w.WriteHeader(http.StatusServiceUnavailable)
    json.NewEncoder(w).Encode(responseMessage)

    return
  }

	w.WriteHeader(http.StatusOK)
  json.NewEncoder(w).Encode(responseMessage)

  return
}


