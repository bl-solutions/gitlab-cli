# GitLab CLI

This CLI tool make the interactions with GitLab project variables easier.

## Installation

You must have Golang v1.22 or higher installed.

```bash
go install github.com/bl-solutions/gitlab-cli@latest
```

## Usage

You can display help with `-h` flag.

### Configuration

This CLI tool used a configuration file to interact with GitLab API.

The configuration file must be named `.gitlab.yaml` and must be stored under your `$HOME`
directory or under the current directory.

This configuration file must contain the following structure:
```yaml
gitlab:
  url: https://gitlab.com/api/v4
  token: <personal-access-token>
```

Please refer to the GitLab documentation to generate a personal access token.

### Get project variables

To get all project variables of a project, you can use the command:
`gitlab-cli get <project-id>`

You can specify a custom filename with the `-f` or `--filename` flags.

The project ID can be found on GitLab UI.

### Put project variables

> Be careful, this command is destructive and overwrite existing variables.

To put project variables from a file, you can use the command:
`gitlab-cli put <project-id>`

You can specify a custom filename with the `-f` or `--filename` flags.

The project ID can be found on GitLab UI.

### Flush project variables

> Be careful, this command erase every variable of the project.

To flush all project variables, you can use the command:
`gitlab-cli flush <project-id>`

The project ID can be found on GitLab UI.

### JSON schema

The project variables are provided following this schema:

```json
[
  {
    "key": "string",
    "value": "string",
    "variable_type": "string",
    "protected": "bool",
    "masked": "bool",
    "raw": "bool",
    "environment_scope": "string",
    "description": "string"
  }
]
```

For details, refer to the [GitLab official documentation about project variables](https://docs.gitlab.com/ee/api/project_level_variables.html).
