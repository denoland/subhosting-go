// File generated from our OpenAPI spec by Stainless.

package subhosting

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/denoland/subhosting-go/internal/apijson"
	"github.com/denoland/subhosting-go/internal/apiquery"
	"github.com/denoland/subhosting-go/internal/param"
	"github.com/denoland/subhosting-go/internal/requestconfig"
	"github.com/denoland/subhosting-go/option"
)

// DeploymentAppLogService contains methods and other services that help with
// interacting with the subhosting API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewDeploymentAppLogService] method
// instead.
type DeploymentAppLogService struct {
	Options []option.RequestOption
}

// NewDeploymentAppLogService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewDeploymentAppLogService(opts ...option.RequestOption) (r *DeploymentAppLogService) {
	r = &DeploymentAppLogService{}
	r.Options = opts
	return
}

// Get execution logs of a deployment
//
// This API can return either past logs or real-time logs depending on the presence
// of the since and until query parameters; if at least one of them is provided,
// past logs are returned, otherwise real-time logs are returned.
//
// Also, the response format can be controlled by the `Accept` header; if
// `application/x-ndjson` is specified, the response will be a stream of
// newline-delimited JSON objects. Otherwise it will be a JSON array of objects.
func (r *DeploymentAppLogService) Get(ctx context.Context, deploymentID string, query DeploymentAppLogGetParams, opts ...option.RequestOption) (res *[]DeploymentAppLogGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("deployments/%s/app_logs", deploymentID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type DeploymentAppLogGetResponse struct {
	Level   DeploymentAppLogGetResponseLevel  `json:"level,required"`
	Message string                            `json:"message,required"`
	Region  DeploymentAppLogGetResponseRegion `json:"region,required"`
	// Log timestamp
	Time time.Time                       `json:"time,required" format:"date-time"`
	JSON deploymentAppLogGetResponseJSON `json:"-"`
}

// deploymentAppLogGetResponseJSON contains the JSON metadata for the struct
// [DeploymentAppLogGetResponse]
type deploymentAppLogGetResponseJSON struct {
	Level       apijson.Field
	Message     apijson.Field
	Region      apijson.Field
	Time        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeploymentAppLogGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DeploymentAppLogGetResponseLevel string

const (
	DeploymentAppLogGetResponseLevelError   DeploymentAppLogGetResponseLevel = "error"
	DeploymentAppLogGetResponseLevelWarning DeploymentAppLogGetResponseLevel = "warning"
	DeploymentAppLogGetResponseLevelInfo    DeploymentAppLogGetResponseLevel = "info"
	DeploymentAppLogGetResponseLevelDebug   DeploymentAppLogGetResponseLevel = "debug"
)

type DeploymentAppLogGetResponseRegion string

const (
	DeploymentAppLogGetResponseRegionGcpAsiaEast1              DeploymentAppLogGetResponseRegion = "gcp-asia-east1"
	DeploymentAppLogGetResponseRegionGcpAsiaEast2              DeploymentAppLogGetResponseRegion = "gcp-asia-east2"
	DeploymentAppLogGetResponseRegionGcpAsiaNortheast1         DeploymentAppLogGetResponseRegion = "gcp-asia-northeast1"
	DeploymentAppLogGetResponseRegionGcpAsiaNortheast2         DeploymentAppLogGetResponseRegion = "gcp-asia-northeast2"
	DeploymentAppLogGetResponseRegionGcpAsiaNortheast3         DeploymentAppLogGetResponseRegion = "gcp-asia-northeast3"
	DeploymentAppLogGetResponseRegionGcpAsiaSouth1             DeploymentAppLogGetResponseRegion = "gcp-asia-south1"
	DeploymentAppLogGetResponseRegionGcpAsiaSouth2             DeploymentAppLogGetResponseRegion = "gcp-asia-south2"
	DeploymentAppLogGetResponseRegionGcpAsiaSoutheast1         DeploymentAppLogGetResponseRegion = "gcp-asia-southeast1"
	DeploymentAppLogGetResponseRegionGcpAsiaSoutheast2         DeploymentAppLogGetResponseRegion = "gcp-asia-southeast2"
	DeploymentAppLogGetResponseRegionGcpAustraliaSoutheast1    DeploymentAppLogGetResponseRegion = "gcp-australia-southeast1"
	DeploymentAppLogGetResponseRegionGcpAustraliaSoutheast2    DeploymentAppLogGetResponseRegion = "gcp-australia-southeast2"
	DeploymentAppLogGetResponseRegionGcpEuropeCentral2         DeploymentAppLogGetResponseRegion = "gcp-europe-central2"
	DeploymentAppLogGetResponseRegionGcpEuropeNorth1           DeploymentAppLogGetResponseRegion = "gcp-europe-north1"
	DeploymentAppLogGetResponseRegionGcpEuropeSouthwest1       DeploymentAppLogGetResponseRegion = "gcp-europe-southwest1"
	DeploymentAppLogGetResponseRegionGcpEuropeWest1            DeploymentAppLogGetResponseRegion = "gcp-europe-west1"
	DeploymentAppLogGetResponseRegionGcpEuropeWest2            DeploymentAppLogGetResponseRegion = "gcp-europe-west2"
	DeploymentAppLogGetResponseRegionGcpEuropeWest3            DeploymentAppLogGetResponseRegion = "gcp-europe-west3"
	DeploymentAppLogGetResponseRegionGcpEuropeWest4            DeploymentAppLogGetResponseRegion = "gcp-europe-west4"
	DeploymentAppLogGetResponseRegionGcpEuropeWest6            DeploymentAppLogGetResponseRegion = "gcp-europe-west6"
	DeploymentAppLogGetResponseRegionGcpEuropeWest8            DeploymentAppLogGetResponseRegion = "gcp-europe-west8"
	DeploymentAppLogGetResponseRegionGcpMeWest1                DeploymentAppLogGetResponseRegion = "gcp-me-west1"
	DeploymentAppLogGetResponseRegionGcpNorthamericaNortheast1 DeploymentAppLogGetResponseRegion = "gcp-northamerica-northeast1"
	DeploymentAppLogGetResponseRegionGcpNorthamericaNortheast2 DeploymentAppLogGetResponseRegion = "gcp-northamerica-northeast2"
	DeploymentAppLogGetResponseRegionGcpSouthamericaEast1      DeploymentAppLogGetResponseRegion = "gcp-southamerica-east1"
	DeploymentAppLogGetResponseRegionGcpSouthamericaWest1      DeploymentAppLogGetResponseRegion = "gcp-southamerica-west1"
	DeploymentAppLogGetResponseRegionGcpUsCentral1             DeploymentAppLogGetResponseRegion = "gcp-us-central1"
	DeploymentAppLogGetResponseRegionGcpUsEast1                DeploymentAppLogGetResponseRegion = "gcp-us-east1"
	DeploymentAppLogGetResponseRegionGcpUsEast4                DeploymentAppLogGetResponseRegion = "gcp-us-east4"
	DeploymentAppLogGetResponseRegionGcpUsEast5                DeploymentAppLogGetResponseRegion = "gcp-us-east5"
	DeploymentAppLogGetResponseRegionGcpUsSouth1               DeploymentAppLogGetResponseRegion = "gcp-us-south1"
	DeploymentAppLogGetResponseRegionGcpUsWest1                DeploymentAppLogGetResponseRegion = "gcp-us-west1"
	DeploymentAppLogGetResponseRegionGcpUsWest2                DeploymentAppLogGetResponseRegion = "gcp-us-west2"
	DeploymentAppLogGetResponseRegionGcpUsWest3                DeploymentAppLogGetResponseRegion = "gcp-us-west3"
	DeploymentAppLogGetResponseRegionGcpUsWest4                DeploymentAppLogGetResponseRegion = "gcp-us-west4"
)

type DeploymentAppLogGetParams struct {
	// Opaque value that represents the cursor of the last log returned in the previous
	// request.
	//
	// This is only effective for the past log mode.
	Cursor param.Field[string] `query:"cursor"`
	// Log level(s) to filter logs by.
	//
	// Defaults to all levels (i.e. no filter applied).
	//
	// Multiple levels can be specified using comma-separated format.
	Level param.Field[DeploymentAppLogGetParamsLevel] `query:"level"`
	// Maximum number of logs to return in one request.
	//
	// This is only effective for the past log mode.
	Limit param.Field[int64] `query:"limit"`
	// Sort order, either `asc` or `desc`. Defaults to `desc`.
	//
	// For backward compatibility, `timeAsc` and `timeDesc` are also supported, but
	// deprecated.
	//
	// This is only effective for the past log mode.
	Order param.Field[string] `query:"order"`
	// Text to search for in log message.
	Q param.Field[string] `query:"q"`
	// Region(s) to filter logs by.
	//
	// Defaults to all regions (i.e. no filter applied).
	//
	// Multiple regions can be specified using comma-separated format.
	Region param.Field[DeploymentAppLogGetParamsRegion] `query:"region"`
	// Start time of the time range to filter logs by.
	//
	// Defaults to the Unix Epoch (though the log retention period is 2 weeks as of
	// now).
	//
	// If neither `since` nor `until` is specified, real-time logs are returned.
	Since param.Field[time.Time] `query:"since" format:"date-time"`
	// The field to sort by. Currently only `time` is supported.
	//
	// This is only effective for the past log mode.
	Sort param.Field[string] `query:"sort"`
	// End time of the time range to filter logs by.
	//
	// Defaults to the current time.
	//
	// If neither `since` nor `until` is specified, real-time logs are returned.
	Until param.Field[time.Time] `query:"until" format:"date-time"`
}

// URLQuery serializes [DeploymentAppLogGetParams]'s query parameters as
// `url.Values`.
func (r DeploymentAppLogGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Log level(s) to filter logs by.
//
// Defaults to all levels (i.e. no filter applied).
//
// Multiple levels can be specified using comma-separated format.
type DeploymentAppLogGetParamsLevel string

const (
	DeploymentAppLogGetParamsLevelError   DeploymentAppLogGetParamsLevel = "error"
	DeploymentAppLogGetParamsLevelWarning DeploymentAppLogGetParamsLevel = "warning"
	DeploymentAppLogGetParamsLevelInfo    DeploymentAppLogGetParamsLevel = "info"
	DeploymentAppLogGetParamsLevelDebug   DeploymentAppLogGetParamsLevel = "debug"
)

// Region(s) to filter logs by.
//
// Defaults to all regions (i.e. no filter applied).
//
// Multiple regions can be specified using comma-separated format.
type DeploymentAppLogGetParamsRegion string

const (
	DeploymentAppLogGetParamsRegionGcpAsiaEast1              DeploymentAppLogGetParamsRegion = "gcp-asia-east1"
	DeploymentAppLogGetParamsRegionGcpAsiaEast2              DeploymentAppLogGetParamsRegion = "gcp-asia-east2"
	DeploymentAppLogGetParamsRegionGcpAsiaNortheast1         DeploymentAppLogGetParamsRegion = "gcp-asia-northeast1"
	DeploymentAppLogGetParamsRegionGcpAsiaNortheast2         DeploymentAppLogGetParamsRegion = "gcp-asia-northeast2"
	DeploymentAppLogGetParamsRegionGcpAsiaNortheast3         DeploymentAppLogGetParamsRegion = "gcp-asia-northeast3"
	DeploymentAppLogGetParamsRegionGcpAsiaSouth1             DeploymentAppLogGetParamsRegion = "gcp-asia-south1"
	DeploymentAppLogGetParamsRegionGcpAsiaSouth2             DeploymentAppLogGetParamsRegion = "gcp-asia-south2"
	DeploymentAppLogGetParamsRegionGcpAsiaSoutheast1         DeploymentAppLogGetParamsRegion = "gcp-asia-southeast1"
	DeploymentAppLogGetParamsRegionGcpAsiaSoutheast2         DeploymentAppLogGetParamsRegion = "gcp-asia-southeast2"
	DeploymentAppLogGetParamsRegionGcpAustraliaSoutheast1    DeploymentAppLogGetParamsRegion = "gcp-australia-southeast1"
	DeploymentAppLogGetParamsRegionGcpAustraliaSoutheast2    DeploymentAppLogGetParamsRegion = "gcp-australia-southeast2"
	DeploymentAppLogGetParamsRegionGcpEuropeCentral2         DeploymentAppLogGetParamsRegion = "gcp-europe-central2"
	DeploymentAppLogGetParamsRegionGcpEuropeNorth1           DeploymentAppLogGetParamsRegion = "gcp-europe-north1"
	DeploymentAppLogGetParamsRegionGcpEuropeSouthwest1       DeploymentAppLogGetParamsRegion = "gcp-europe-southwest1"
	DeploymentAppLogGetParamsRegionGcpEuropeWest1            DeploymentAppLogGetParamsRegion = "gcp-europe-west1"
	DeploymentAppLogGetParamsRegionGcpEuropeWest2            DeploymentAppLogGetParamsRegion = "gcp-europe-west2"
	DeploymentAppLogGetParamsRegionGcpEuropeWest3            DeploymentAppLogGetParamsRegion = "gcp-europe-west3"
	DeploymentAppLogGetParamsRegionGcpEuropeWest4            DeploymentAppLogGetParamsRegion = "gcp-europe-west4"
	DeploymentAppLogGetParamsRegionGcpEuropeWest6            DeploymentAppLogGetParamsRegion = "gcp-europe-west6"
	DeploymentAppLogGetParamsRegionGcpEuropeWest8            DeploymentAppLogGetParamsRegion = "gcp-europe-west8"
	DeploymentAppLogGetParamsRegionGcpMeWest1                DeploymentAppLogGetParamsRegion = "gcp-me-west1"
	DeploymentAppLogGetParamsRegionGcpNorthamericaNortheast1 DeploymentAppLogGetParamsRegion = "gcp-northamerica-northeast1"
	DeploymentAppLogGetParamsRegionGcpNorthamericaNortheast2 DeploymentAppLogGetParamsRegion = "gcp-northamerica-northeast2"
	DeploymentAppLogGetParamsRegionGcpSouthamericaEast1      DeploymentAppLogGetParamsRegion = "gcp-southamerica-east1"
	DeploymentAppLogGetParamsRegionGcpSouthamericaWest1      DeploymentAppLogGetParamsRegion = "gcp-southamerica-west1"
	DeploymentAppLogGetParamsRegionGcpUsCentral1             DeploymentAppLogGetParamsRegion = "gcp-us-central1"
	DeploymentAppLogGetParamsRegionGcpUsEast1                DeploymentAppLogGetParamsRegion = "gcp-us-east1"
	DeploymentAppLogGetParamsRegionGcpUsEast4                DeploymentAppLogGetParamsRegion = "gcp-us-east4"
	DeploymentAppLogGetParamsRegionGcpUsEast5                DeploymentAppLogGetParamsRegion = "gcp-us-east5"
	DeploymentAppLogGetParamsRegionGcpUsSouth1               DeploymentAppLogGetParamsRegion = "gcp-us-south1"
	DeploymentAppLogGetParamsRegionGcpUsWest1                DeploymentAppLogGetParamsRegion = "gcp-us-west1"
	DeploymentAppLogGetParamsRegionGcpUsWest2                DeploymentAppLogGetParamsRegion = "gcp-us-west2"
	DeploymentAppLogGetParamsRegionGcpUsWest3                DeploymentAppLogGetParamsRegion = "gcp-us-west3"
	DeploymentAppLogGetParamsRegionGcpUsWest4                DeploymentAppLogGetParamsRegion = "gcp-us-west4"
)
