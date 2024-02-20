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

// ProjectAnalyticsService contains methods and other services that help with
// interacting with the subhosting API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewProjectAnalyticsService] method
// instead.
type ProjectAnalyticsService struct {
	Options []option.RequestOption
}

// NewProjectAnalyticsService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewProjectAnalyticsService(opts ...option.RequestOption) (r *ProjectAnalyticsService) {
	r = &ProjectAnalyticsService{}
	r.Options = opts
	return
}

// Retrieve project analytics
//
// This API returns analytics for the specified project. The analytics are returned
// as time series data in 15 minute intervals, with the `time` field representing
// the start of the interval.
func (r *ProjectAnalyticsService) Get(ctx context.Context, projectID string, query ProjectAnalyticsGetParams, opts ...option.RequestOption) (res *shared.Analytics, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("projects/%s/analytics", projectID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type ProjectAnalyticsGetParams struct {
	// Start of the time range in RFC3339 format.
	//
	// Defaults to 24 hours ago.
	Since param.Field[time.Time] `query:"since,required" format:"date-time"`
	// End of the time range in RFC3339 format.
	//
	// Defaults to the current time.
	Until param.Field[time.Time] `query:"until,required" format:"date-time"`
}

// URLQuery serializes [ProjectAnalyticsGetParams]'s query parameters as
// `url.Values`.
func (r ProjectAnalyticsGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
