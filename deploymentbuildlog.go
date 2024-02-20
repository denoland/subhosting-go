// File generated from our OpenAPI spec by Stainless.

package subhosting

import (
	"context"
	"fmt"
	"net/http"

	"github.com/denoland/subhosting-go/internal/apijson"
	"github.com/denoland/subhosting-go/internal/requestconfig"
	"github.com/denoland/subhosting-go/option"
)

// DeploymentBuildLogService contains methods and other services that help with
// interacting with the subhosting API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewDeploymentBuildLogService] method
// instead.
type DeploymentBuildLogService struct {
	Options []option.RequestOption
}

// NewDeploymentBuildLogService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewDeploymentBuildLogService(opts ...option.RequestOption) (r *DeploymentBuildLogService) {
	r = &DeploymentBuildLogService{}
	r.Options = opts
	return
}

// Get build logs of a deployment
//
// This API returns build logs of the specified deployment. It's useful to watch
// the build progress, figure out what went wrong in case of a build failure, and
// so on.
//
// The response format can be controlled by the `Accept` header; if
// `application/x-ndjson` is specified, the response will be a stream of
// newline-delimited JSON objects. Otherwise it will be a JSON array of objects.
func (r *DeploymentBuildLogService) Get(ctx context.Context, deploymentID string, opts ...option.RequestOption) (res *[]DeploymentBuildLogGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("deployments/%s/build_logs", deploymentID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type DeploymentBuildLogGetResponse struct {
	Level   string                            `json:"level,required"`
	Message string                            `json:"message,required"`
	JSON    deploymentBuildLogGetResponseJSON `json:"-"`
}

// deploymentBuildLogGetResponseJSON contains the JSON metadata for the struct
// [DeploymentBuildLogGetResponse]
type deploymentBuildLogGetResponseJSON struct {
	Level       apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeploymentBuildLogGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
