// https://learn.microsoft.com/en-us/azure/azure-functions/create-first-function-vs-code-other?tabs=go%2Cmacos

package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func taskHandler(w http.ResponseWriter, r *http.Request) {
	message := "This HTTP triggered function executed successfully.\n"
	reqBody, err := ioutil.ReadAll(r.Body)
	if err == nil {
		wf, err := ParseWorkflow(reqBody)
		if err == nil {
			states := wf.States
			ExecuteStateAndReturnNext(wf.States, states[0])
		} else {
			message = "Failed to parse workflow!"
		}
	} else {
		message = "Missing body!"
	}
	fmt.Fprint(w, message)
}

func main() {
	listenAddr := ":8080"
	if val, ok := os.LookupEnv("FUNCTIONS_CUSTOMHANDLER_PORT"); ok {
		listenAddr = ":" + val
	}
	http.HandleFunc("/api/invoketask", taskHandler)
	http.ListenAndServe(listenAddr, nil)
}
