// File generated from our OpenAPI spec by Stainless.

package subhosting

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/denoland/subhosting-go/internal/apiquery"
	"github.com/denoland/subhosting-go/internal/param"
	"github.com/denoland/subhosting-go/internal/requestconfig"
	"github.com/denoland/subhosting-go/internal/shared"
	"github.com/denoland/subhosting-go/option"
)

// OrganizationAnalyticsService contains methods and other services that help with
// interacting with the subhosting API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewOrganizationAnalyticsService]
// method instead.
type OrganizationAnalyticsService struct {
	Options []option.RequestOption
}

// NewOrganizationAnalyticsService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewOrganizationAnalyticsService(opts ...option.RequestOption) (r *OrganizationAnalyticsService) {
	r = &OrganizationAnalyticsService{}
	r.Options = opts
	return
}

// Retrieve organization analytics
//
// This API returns analytics for the specified organization. The analytics are
// returned as time series data in 15 minute intervals, with the `time` field
// representing the start of the interval.
func (r *OrganizationAnalyticsService) Get(ctx context.Context, organizationID string, query OrganizationAnalyticsGetParams, opts ...option.RequestOption) (res *shared.Analytics, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("organizations/%s/analytics", organizationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type OrganizationAnalyticsGetParams struct {
	// Start of the time range in RFC3339 format.
	//
	// Defaults to 24 hours ago.
	//
	// Note that the maximum allowed time range is 24 hours.
	Since param.Field[time.Time] `query:"since,required" format:"date-time"`
	// End of the time range in RFC3339 format.
	//
	// Defaults to the current time.
	//
	// Note that the maximum allowed time range is 24 hours.
	Until param.Field[time.Time] `query:"until,required" format:"date-time"`
}

// URLQuery serializes [OrganizationAnalyticsGetParams]'s query parameters as
// `url.Values`.
func (r OrganizationAnalyticsGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
