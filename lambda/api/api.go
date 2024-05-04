package api

import (
	"lambda-func/database"
)

type ApiHandler struct {
dbStore database.DynamoDBClient
}

funct NewApiHandler (dbStore string) ApiHandler(

)
