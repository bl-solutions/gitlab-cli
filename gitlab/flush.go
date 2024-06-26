package gitlab

import (
	gl "github.com/xanzy/go-gitlab"
	"log"
)

func HandleFlushVariables(client *gl.Client, project int) {
	vars, err := listProjectVariables(client, project, &gl.ListProjectVariablesOptions{})
	if err != nil {
		log.Fatalln(err)
	}
	for _, v := range vars {
		log.Printf("Remove %s variable with %s scope on %d project", v.Key, v.EnvironmentScope, project)
		_, err = client.ProjectVariables.RemoveVariable(project, v.Key, &gl.RemoveProjectVariableOptions{Filter: &gl.VariableFilter{EnvironmentScope: v.EnvironmentScope}})
		if err != nil {
			log.Fatalln(err)
		}
	}
}
