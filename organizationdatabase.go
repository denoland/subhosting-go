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
	"github.com/denoland/subhosting-go/internal/shared"
	"github.com/denoland/subhosting-go/option"
)

// OrganizationDatabaseService contains methods and other services that help with
// interacting with the subhosting API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewOrganizationDatabaseService]
// method instead.
type OrganizationDatabaseService struct {
	Options []option.RequestOption
}

// NewOrganizationDatabaseService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewOrganizationDatabaseService(opts ...option.RequestOption) (r *OrganizationDatabaseService) {
	r = &OrganizationDatabaseService{}
	r.Options = opts
	return
}

// Create a KV database
//
// This API allows you to create a new KV database under the specified
// organization. You will then be able to associate the created KV database with a
// new deployment by specifying the KV database ID in the "Create a deployment" API
// call.
func (r *OrganizationDatabaseService) New(ctx context.Context, organizationID string, body OrganizationDatabaseNewParams, opts ...option.RequestOption) (res *shared.KvDatabase, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("organizations/%s/databases", organizationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// List KV databases of an organization
//
// This API returns a list of KV databases belonging to the specified organization
// in a pagenated manner. The URLs for the next, previous, first, and last page are
// returned in the `Link` header of the response, if any.
func (r *OrganizationDatabaseService) List(ctx context.Context, organizationID string, query OrganizationDatabaseListParams, opts ...option.RequestOption) (res *[]shared.KvDatabase, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("organizations/%s/databases", organizationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type OrganizationDatabaseNewParams struct {
	// The description of the KV database. If this is `null`, an empty string will be
	// set.
	Description param.Field[string] `json:"description"`
}

func (r OrganizationDatabaseNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type OrganizationDatabaseListParams struct {
	// The maximum number of items to return per page.
	Limit param.Field[int64] `query:"limit"`
	// Sort order, either `asc` or `desc`. Defaults to `asc`.
	Order param.Field[string] `query:"order"`
	// The page number to return.
	Page param.Field[int64] `query:"page"`
	// Query by KV database ID
	Q param.Field[string] `query:"q"`
	// The field to sort by. Currently only `created_at` is supported.
	Sort param.Field[string] `query:"sort"`
}

// URLQuery serializes [OrganizationDatabaseListParams]'s query parameters as
// `url.Values`.
func (r OrganizationDatabaseListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
