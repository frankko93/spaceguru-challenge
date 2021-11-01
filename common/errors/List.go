package errors

// CategorizedError ...
type categorizedError struct {
	Code        string `json:"code,omitempty"`
	Reference   string `json:"reference,omitempty"`
	Description string `json:"description"`
}

var (
	InvalidID = categorizedError{
		Code:        "1000",
		Description: "Invalid ID",
	}

	SameUserPermissions = categorizedError{
		Code:        "4001",
		Description: "Involved users cannot share permissions or be the same",
	}

	InvalidPermissions = categorizedError{
		Code:        "4002",
		Description: "Invalid permissions for user",
	}

	NotFoundUserPermissions = categorizedError{
		Code:        "4003",
		Description: "Approver users or roles do not exist",
	}

	NotApprover = categorizedError{
		Code:        "4004",
		Description: "Approvers must be sent",
	}

	FieldValidation = categorizedError{
		Code:        "4010",
		Description: "Invalid field",
	}

	StructureFormatError = categorizedError{
		Code:        "4020",
		Description: "Invalid structure format",
	}

	InvalidStatus = categorizedError{
		Code:        "4030",
		Description: "Invalid status",
	}

	InvalidCallbackStatus = categorizedError{
		Code:        "4040",
		Description: "Invalid callback status",
	}

	InvalidAction = categorizedError{
		Code:        "4050",
		Description: "Invalid action",
	}

	InvalidStrategyType = categorizedError{
		Code:        "4060",
		Description: "Invalid strategy type",
	}

	InvalidStrategyPeerRAList = categorizedError{
		Code:        "4070",
		Description: "Invalid levels for strategy peer",
	}

	InvalidStrategyPeerRE = categorizedError{
		Code:        "4071",
		Description: "Invalid requested_approver for strategy peer",
	}

	InvalidStrategyEscalationRAList = categorizedError{
		Code:        "4080",
		Description: "Invalid levels for strategy escalation",
	}

	InvalidStrategyEscalationRAListRE = categorizedError{
		Code:        "4081",
		Description: "Invalid levels for strategy escalation - ID & ROLE from requested_approver can't be empty",
	}

	InvalidStrategyEscalationRAListRELvl = categorizedError{
		Code:        "4082",
		Description: "Invalid levels for strategy escalation - invalid requested_approver.level",
	}

	InvalidStrategyApprovalsQuantity = categorizedError{
		Code:        "4090",
		Description: "Invalid approvals_quantity",
	}

	InvalidRoleOrID = categorizedError{
		Code:        "4100",
		Description: "Invalid ROLE or ID - it cannot be empty",
	}
)

// CustomFieldValidation ...
func CustomFieldValidationError(field, description string) (err categorizedError) {
	err = categorizedError{
		Code:        FieldValidation.Code,
		Reference:   field,
		Description: description,
	}

	return
}
