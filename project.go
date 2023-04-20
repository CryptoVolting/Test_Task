package testProject

import "errors"

type Project struct {
	Id          string `json:"id" db:"id"`
	Name        string `json:"name" db:"name" binding:"required"`
	TypeProject string `json:"typeProject" db:"typeproject" binding:"required"`
}

type UpdateProjectInput struct {
	Id          *string `json:"id"`
	Name        *string `json:"name"`
	TypeProject *string `json:"typeProject"`
}

func (i UpdateProjectInput) Validate() error {
	if i.Name == nil && i.TypeProject == nil {
		return errors.New("update structure has no values")
	}

	return nil
}

type OperatorProject struct {
	Id         string
	IdOperator string
	IdProject  string
}

type IdOperatorAndProject struct {
	IdOperator *string `json:"idOperator"`
	IdProject  *string `json:"idProject"`
}
