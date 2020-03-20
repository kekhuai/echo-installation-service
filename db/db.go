package db

import (
	"fmt"

	"github.com/Kamva/mgm/v2"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// New db instance
func New() error {
	err := mgm.SetDefaultConfig(nil, "echo_installation_service", options.Client().ApplyURI("mongodb://localhost:27017"))
	if nil != err {
		fmt.Println("storage error: ", err)
	}
	return nil
}
