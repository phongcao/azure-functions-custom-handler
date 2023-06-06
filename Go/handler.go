// https://learn.microsoft.com/en-us/azure/azure-functions/create-first-function-vs-code-other?tabs=go%2Cmacos

package main

import (
	"fmt"
	"net/http"
	"os"
	"reflect"
)

func taskHandler(w http.ResponseWriter, r *http.Request) {
	message := "This HTTP triggered function executed successfully. Pass a task in the query string for a personalized response.\n"
	taskName := r.URL.Query().Get("task")
	if taskName != "" {
		method := reflect.ValueOf(TaskType(0)).MethodByName(taskName)
		if !method.IsValid() {
			message = "Task not found!"
		} else {
			res := method.Call(nil)
			message = res[0].String()
		}
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
