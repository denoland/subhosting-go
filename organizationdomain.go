// File generated from our OpenAPI spec by Stainless.

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

// OrganizationDomainService contains methods and other services that help with
// interacting with the subhosting API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewOrganizationDomainService] method
// instead.
type OrganizationDomainService struct {
	Options []option.RequestOption
}

// NewOrganizationDomainService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewOrganizationDomainService(opts ...option.RequestOption) (r *OrganizationDomainService) {
	r = &OrganizationDomainService{}
	r.Options = opts
	return
}

// Add a domain to an organization
//
// This API allows you to add a new domain to the specified organization.
//
// Before use, added domain needs to be verified, and also TLS certificates for the
// domain need to be provisioned.
func (r *OrganizationDomainService) New(ctx context.Context, organizationID string, body OrganizationDomainNewParams, opts ...option.RequestOption) (res *shared.Domain, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("organizations/%s/domains", organizationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// List domains of an organization
//
// This API returns a list of domains belonging to the specified organization in a
// pagenated manner.
//
// The URLs for the next, previous, first, and last page are returned in the `Link`
// header of the response, if any.
func (r *OrganizationDomainService) List(ctx context.Context, organizationID string, query OrganizationDomainListParams, opts ...option.RequestOption) (res *[]shared.Domain, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("organizations/%s/domains", organizationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type OrganizationDomainNewParams struct {
	Domain param.Field[string] `json:"domain,required"`
}

func (r OrganizationDomainNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type OrganizationDomainListParams struct {
	// The maximum number of items to return per page.
	Limit param.Field[int64] `query:"limit"`
	// Sort order, either `asc` or `desc`. Defaults to `asc`.
	Order param.Field[string] `query:"order"`
	// The page number to return.
	Page param.Field[int64] `query:"page"`
	// Query by domain
	Q param.Field[string] `query:"q"`
	// The field to sort by, `domain`, `created_at`, or `updated_at`. Defaults to
	// `updated_at`.
	Sort param.Field[string] `query:"sort"`
}

// URLQuery serializes [OrganizationDomainListParams]'s query parameters as
// `url.Values`.
func (r OrganizationDomainListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
