package gitlab

import (
	gl "github.com/xanzy/go-gitlab"
	"log"
)

func HandlePutVariables(client *gl.Client, project int, filename string) {
	vars, err := importFromJson(filename)
	if err != nil {
		log.Fatalf(err.Error())
	}
	if err := updateProjectVariables(client, project, vars); err != nil {
		log.Fatalf(err.Error())
	}
}

func updateProjectVariables(client *gl.Client, project int, vars []gl.ProjectVariable) error {
	for _, v := range vars {
		//time.Sleep(time.Millisecond * 200)
		log.Printf("Creating %s variable with %s scope", v.Key, v.EnvironmentScope)
		if err := handleUpdateProjectVariables(client, project, v); err != nil {
			log.Printf("An error occured with %s variable with %s scope: %v", v.Key, v.EnvironmentScope, err)
			return err
		}
	}
	return nil
}

func handleUpdateProjectVariables(client *gl.Client, project int, v gl.ProjectVariable) error {
	_, rsp, err := client.ProjectVariables.UpdateVariable(project, v.Key, &gl.UpdateProjectVariableOptions{
		Value:            &v.Value,
		Description:      &v.Description,
		EnvironmentScope: &v.EnvironmentScope,
		Filter:           &gl.VariableFilter{EnvironmentScope: v.EnvironmentScope},
		Masked:           canBeMasked(&v),
		Protected:        &v.Protected,
		Raw:              &v.Raw,
		VariableType:     &v.VariableType,
	})

	if err != nil {
		if rsp.StatusCode == 404 {
			if err := handleCreateProjectVariables(client, project, v); err != nil {
				return err
			}
		} else {
			return err
		}
	}
	return nil
}

func handleCreateProjectVariables(client *gl.Client, project int, v gl.ProjectVariable) error {
	_, _, err := client.ProjectVariables.CreateVariable(project, &gl.CreateProjectVariableOptions{
		Key:              &v.Key,
		Value:            &v.Value,
		Description:      &v.Description,
		EnvironmentScope: &v.EnvironmentScope,
		Masked:           canBeMasked(&v),
		Protected:        &v.Protected,
		Raw:              &v.Raw,
		VariableType:     &v.VariableType,
	})
	if err != nil {
		return err
	}
	return nil
}
