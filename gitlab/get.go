package gitlab

import (
	gl "github.com/xanzy/go-gitlab"
	"log"
)

func HandleGetVariables(client *gl.Client, project int, filename string) {
	vars, err := listProjectVariables(client, project, &gl.ListProjectVariablesOptions{})
	if err != nil {
		log.Fatalf("Failed to list project variables: %v", err)
	}
	if err := exportToJson(filename, vars); err != nil {
		log.Fatalf(err.Error())
	}
}
