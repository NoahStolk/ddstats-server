package api

import (
	"log"
	"net/http"

	"github.com/alexwilkerson/ddstats-api/pkg/ddapi"
	"github.com/alexwilkerson/ddstats-api/pkg/models/postgres"

	"github.com/alexwilkerson/ddstats-api/pkg/websocket"
)

const (
	oldestValidClientVersion = "0.3.1"
	currentClientVersion     = "0.4.5"
)

type API struct {
	client       *http.Client
	db           *postgres.Postgres
	websocketHub *websocket.Hub
	ddAPI        *ddapi.API
	infoLog      *log.Logger
	errorLog     *log.Logger
}

func NewAPI(client *http.Client, db *postgres.Postgres, websocketHub *websocket.Hub, ddapi *ddapi.API, infoLog, errorLog *log.Logger) *API {
	return &API{
		client:       client,
		db:           db,
		websocketHub: websocketHub,
		ddAPI:        ddapi,
		infoLog:      infoLog,
		errorLog:     errorLog,
	}
}