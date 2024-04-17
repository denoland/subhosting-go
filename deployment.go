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

// DeploymentService contains methods and other services that help with interacting
// with the subhosting API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewDeploymentService] method instead.
type DeploymentService struct {
	Options   []option.RequestOption
	BuildLogs *DeploymentBuildLogService
	AppLogs   *DeploymentAppLogService
}

// NewDeploymentService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewDeploymentService(opts ...option.RequestOption) (r *DeploymentService) {
	r = &DeploymentService{}
	r.Options = opts
	r.BuildLogs = NewDeploymentBuildLogService(opts...)
	r.AppLogs = NewDeploymentAppLogService(opts...)
	return
}

// Delete a deployment
func (r *DeploymentService) Delete(ctx context.Context, deploymentID string, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := fmt.Sprintf("deployments/%s", deploymentID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Get deployment details
func (r *DeploymentService) Get(ctx context.Context, deploymentID string, opts ...option.RequestOption) (res *shared.Deployment, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("deployments/%s", deploymentID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Redeploy a deployment with different configuration
func (r *DeploymentService) Redeploy(ctx context.Context, deploymentID string, body DeploymentRedeployParams, opts ...option.RequestOption) (res *shared.Deployment, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("deployments/%s/redeploy", deploymentID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type DeploymentRedeployParams struct {
	// KV database ID mappings to associate with the deployment.
	//
	// A key represents a KV database name (e.g. `"default"`), and a value is a KV
	// database ID.
	//
	// Currently, only `"default"` database is supported. If any other database name is
	// specified, that will be rejected.
	//
	// The provided KV database mappings will be _merged_ with the existing one, just
	// like `env_vars`.
	//
	// If `databases` itself is not provided, no update will happen to the existing KV
	// database mappings.
	Databases param.Field[map[string]string] `json:"databases" format:"uuid"`
	// A description of the created deployment. If not provided, no update will happen
	// to the description.
	Description param.Field[string] `json:"description"`
	// A dictionary of environment variables to be set in the runtime environment of
	// the deployment.
	//
	// The provided environment variables will be _merged_ with the existing one. For
	// example, if the existing environment variables are:
	//
	// ```json
	// {
	// "a": "alice",
	// "b": "bob"
	// "c": "charlie"
	// }
	// ```
	//
	// and you pass the following environment variables in your redeploy request:
	//
	// ```json
	//
	//	{
	//	  "a": "alice2",
	//	  "b": null,
	//	  "d": "david"
	//	}
	//
	// ```
	//
	// then the result will be:
	//
	// ```json
	//
	//	{
	//	  "a": "alice2",
	//	  "c": "charlie",
	//	  "d": "david"
	//	}
	//
	// ```
	//
	// If `env_vars` itself is not provided, no update will happen to the existing
	// environment variables.
	EnvVars param.Field[map[string]string] `json:"env_vars"`
}

func (r DeploymentRedeployParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
