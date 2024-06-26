package gitlab

import (
	"context"
	"encoding/json"
	"fmt"
	gl "github.com/xanzy/go-gitlab"
	"log"
	"os"
	"regexp"
	"time"
)

const (
	timeoutMs = 1000
)

func InitClient(url string, token string, projectId int) (*gl.Client, error) {
	ctx := context.Background()

	// Create a GitLab client with personal access token and custom URL
	client, err := gl.NewClient(
		token,
		gl.WithBaseURL(url),
		gl.WithCustomRetryMax(3),
		gl.WithCustomRetryWaitMinMax(time.Millisecond*100, time.Millisecond*1000),
	)

	if err != nil {
		return nil, err
	}

	// Test the connection to GitLab with the client
	err = tryConnection(ctx, client)

	if err != nil {
		return nil, err
	}

	// Verify if the project exists
	if err := projectExists(ctx, client, projectId); err != nil {
		return nil, err
	}

	return client, nil
}

func tryConnection(ctx context.Context, client *gl.Client) error {
	ctx, cancel := context.WithTimeout(ctx, time.Millisecond*timeoutMs)
	defer cancel()
	ch := make(chan error)

	go func() {
		_, _, err := client.Metadata.GetMetadata(gl.WithContext(ctx))
		if err != nil {
			ch <- err
		} else {
			ch <- nil
		}
	}()

	for {
		select {
		case <-ctx.Done():
			return fmt.Errorf("connection timeout")
		case err := <-ch:
			return err
		}
	}
}

func projectExists(ctx context.Context, client *gl.Client, pid int) error {
	ctx, cancel := context.WithTimeout(ctx, time.Millisecond*timeoutMs)
	defer cancel()
	ch := make(chan error)

	go func() {
		_, _, err := client.Projects.GetProject(pid, &gl.GetProjectOptions{})
		if err != nil {
			ch <- err
		} else {
			ch <- nil
		}
	}()

	for {
		select {
		case <-ctx.Done():
			return fmt.Errorf("connection timeout")
		case err := <-ch:
			return err
		}
	}
}

func canBeMasked(v *gl.ProjectVariable) *bool {
	isSpacePresent := regexp.MustCompile(`\s`).MatchString(v.Value)
	isQuestionMarkPresent := regexp.MustCompile(`\?`).MatchString(v.Value)
	result := len(v.Value) > 8 && !isSpacePresent && !isQuestionMarkPresent
	return &result
}

func listProjectVariables(client *gl.Client, pid int, opt *gl.ListProjectVariablesOptions) ([]*gl.ProjectVariable, error) {
	variables, response, err := client.ProjectVariables.ListVariables(pid, opt)

	if err != nil {
		log.Printf("Failed to list variables: %v", err)
		return nil, err
	}

	if response.NextPage != 0 {
		nextVariables, err := listProjectVariables(
			client, pid,
			&gl.ListProjectVariablesOptions{
				Page:       response.NextPage,
				PerPage:    opt.PerPage,
				OrderBy:    opt.OrderBy,
				Pagination: opt.Pagination,
				Sort:       opt.Sort,
			})
		if err != nil {
			return nil, err
		}
		variables = append(variables, nextVariables...)
	}

	return variables, nil
}

func exportToJson(filename string, data any) error {
	file, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}

	if err := os.WriteFile(filename, file, 0644); err != nil {
		return err
	}

	return nil
}

func importFromJson(filename string) ([]gl.ProjectVariable, error) {
	file, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var data []gl.ProjectVariable
	if err := json.Unmarshal(file, &data); err != nil {
		return nil, err
	}

	return data, nil
}
