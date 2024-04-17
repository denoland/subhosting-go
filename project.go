// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package subhosting

import (
	"context"
	"fmt"
	"net/http"

	"github.com/denoland/subhosting-go/internal/apijson"
	"github.com/denoland/subhosting-go/internal/param"
	"github.com/denoland/subhosting-go/internal/requestconfig"
	"github.com/denoland/subhosting-go/internal/shared"
	"github.com/denoland/subhosting-go/option"
)

// ProjectService contains methods and other services that help with interacting
// with the subhosting API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewProjectService] method instead.
type ProjectService struct {
	Options     []option.RequestOption
	Analytics   *ProjectAnalyticsService
	Deployments *ProjectDeploymentService
}

// NewProjectService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewProjectService(opts ...option.RequestOption) (r *ProjectService) {
	r = &ProjectService{}
	r.Options = opts
	r.Analytics = NewProjectAnalyticsService(opts...)
	r.Deployments = NewProjectDeploymentService(opts...)
	return
}

// Update project details
func (r *ProjectService) Update(ctx context.Context, projectID string, body ProjectUpdateParams, opts ...option.RequestOption) (res *shared.Project, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("projects/%s", projectID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Delete a project
func (r *ProjectService) Delete(ctx context.Context, projectID string, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := fmt.Sprintf("projects/%s", projectID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Get project details
func (r *ProjectService) Get(ctx context.Context, projectID string, opts ...option.RequestOption) (res *shared.Project, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("projects/%s", projectID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type ProjectUpdateParams struct {
	// The description of the project to be updated to. If this is `null`, no update
	// will be made to the project description.
	Description param.Field[string] `json:"description"`
	// The name of the project to be updated to. This must be globally unique. If this
	// is `null`, no update will be made to the project name.
	Name param.Field[string] `json:"name"`
}

func (r ProjectUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
