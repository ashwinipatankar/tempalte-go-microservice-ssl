package main

import (
"fmt"
	"log"
	"net/http"

	//"github.com/Stenny-io/stenny/services/auth/api"
"github.com/ashwinipatankar/tempalte-go-microservice-ssl/services/common/conf"
"github.com/ashwinipatankar/tempalte-go-microservice-ssl/services/hello-world/api"

)

func main() {
	httpPort := conf.CONF("HTTPPORT")
	fmt.Printf("Server hosted on port: %v\n", httpPort)
	// Cron job bindings here
	// TODO: delete old destroyed JWT cron job acc to date
	//Start server
log.Fatal(http.ListenAndServe("0.0.0.0:"+httpPort, api.Router()))
}

