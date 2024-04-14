package handler

import "fmt"

func errorParameterIsRequired(name, typeName string) error {
	return fmt.Errorf("Parameter: %s (type: %s) is required", name, typeName)
}

type CreateOpeningRequest struct {
	Role     string `json:"role"`
	Company  string `json:"company"`
	Location string `json:"location"`
	Remote   *bool  `json:"remote" `
	Link     string `json:"link"`
	Salary   int64  `json:"salary"`
}

func (request *CreateOpeningRequest) Validate() error {

	if request.Role == "" && request.Location == "" && request.Company == "" && request.Link == "" && request.Remote == nil && request.Salary <= 0 {
		return fmt.Errorf("Request body is empty or mal formade")
	}

	if request.Role == "" {
		return errorParameterIsRequired("role", "string")
	}

	if request.Location == "" {
		return errorParameterIsRequired("location", "string")
	}

	if request.Company == "" {
		return errorParameterIsRequired("company", "string")
	}
	if request.Link == "" {
		return errorParameterIsRequired("link", "string")
	}
	if request.Remote == nil {
		return errorParameterIsRequired("remote", "bool")
	}
	if request.Salary <= 0 {
		return errorParameterIsRequired("salary", "int64")
	}
	return nil
}

type UpdateRequestOpening struct {
	Role     string `json:"role"`
	Company  string `json:"company"`
	Location string `json:"location"`
	Remote   *bool  `json:"remote" `
	Link     string `json:"link"`
	Salary   int64  `json:"salary"`
}

func (request *UpdateRequestOpening) ValidateFields() error {
	if request.Role == "" && request.Location == "" && request.Company == "" && request.Link == "" && request.Remote == nil && request.Salary <= 0 {
		return fmt.Errorf("Request body is empty or mal formade")
	}

	if request.Role == "" {
		return errorParameterIsRequired("role", "string")
	}

	if request.Location == "" {
		return errorParameterIsRequired("location", "string")
	}

	if request.Company == "" {
		return errorParameterIsRequired("company", "string")
	}
	if request.Link == "" {
		return errorParameterIsRequired("link", "string")
	}
	if request.Remote == nil {
		return errorParameterIsRequired("remote", "bool")
	}
	if request.Salary <= 0 {
		return errorParameterIsRequired("salary", "int64")
	}
	return nil
}
