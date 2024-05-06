// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package subhosting

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/denoland/subhosting-go/internal/apijson"
	"github.com/denoland/subhosting-go/internal/apiquery"
	"github.com/denoland/subhosting-go/internal/param"
	"github.com/denoland/subhosting-go/internal/requestconfig"
	"github.com/denoland/subhosting-go/option"
	"github.com/denoland/subhosting-go/shared"
)

// OrganizationProjectService contains methods and other services that help with
// interacting with the subhosting API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewOrganizationProjectService]
// method instead.
type OrganizationProjectService struct {
	Options []option.RequestOption
}

// NewOrganizationProjectService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewOrganizationProjectService(opts ...option.RequestOption) (r *OrganizationProjectService) {
	r = &OrganizationProjectService{}
	r.Options = opts
	return
}

// Create a project
//
// This API allows you to create a new project under the specified organization.
// The project name is optional; if not provided, a random name will be generated.
func (r *OrganizationProjectService) New(ctx context.Context, organizationID string, body OrganizationProjectNewParams, opts ...option.RequestOption) (res *shared.Project, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("organizations/%s/projects", organizationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// List projects of an organization
//
// This API returns a list of projects belonging to the specified organization in a
// pagenated manner. The URLs for the next, previous, first, and last page are
// returned in the `Link` header of the response, if any.
func (r *OrganizationProjectService) List(ctx context.Context, organizationID string, query OrganizationProjectListParams, opts ...option.RequestOption) (res *[]shared.Project, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("organizations/%s/projects", organizationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type OrganizationProjectNewParams struct {
	// The description of the project. If this is `null`, an empty string will be set.
	Description param.Field[string] `json:"description"`
	// The name of the project. This must be globally unique. If this is `null`, a
	// random unique name will be generated.
	Name param.Field[string] `json:"name"`
}

func (r OrganizationProjectNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type OrganizationProjectListParams struct {
	// The maximum number of items to return per page.
	Limit param.Field[int64] `query:"limit"`
	// Sort order, either `asc` or `desc`. Defaults to `asc`.
	Order param.Field[string] `query:"order"`
	// The page number to return.
	Page param.Field[int64] `query:"page"`
	// Query by project name or project ID
	Q param.Field[string] `query:"q"`
	// The field to sort by, either `name` or `updated_at`. Defaults to `updated_at`.
	Sort param.Field[string] `query:"sort"`
}

// URLQuery serializes [OrganizationProjectListParams]'s query parameters as
// `url.Values`.
func (r OrganizationProjectListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
