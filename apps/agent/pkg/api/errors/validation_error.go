package errors

import (
	"context"
	"net/http"

	"github.com/Southclaws/fault/fmsg"
	"github.com/getghost/ghost/apps/agent/pkg/api/ctxutil"
	"github.com/getghost/ghost/apps/agent/pkg/openapi"
)

func HandleValidationError(ctx context.Context, err error) openapi.ValidationError {

	issues := fmsg.GetIssues(err)
	details := make([]openapi.ValidationErrorDetail, len(issues))
	for i, issue := range issues {
		details[i] = openapi.ValidationErrorDetail{
			Message: issue,
		}
	}

	return openapi.ValidationError{
		Title:     "Internal Server Error",
		Detail:    "An internal server error occurred",
		Errors:    details,
		Instance:  "https://errors.ghost.com/todo",
		Status:    http.StatusBadRequest,
		RequestId: ctxutil.GetRequestId(ctx),
		Type:      "TODO docs link",
	}

}
