// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package subhosting

import (
	"context"
	"fmt"
	"net/http"

	"github.com/denoland/subhosting-go/internal/apijson"
	"github.com/denoland/subhosting-go/internal/param"
	"github.com/denoland/subhosting-go/internal/requestconfig"
	"github.com/denoland/subhosting-go/option"
	"github.com/denoland/subhosting-go/shared"
)

// DatabaseService contains methods and other services that help with interacting
// with the subhosting API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewDatabaseService] method instead.
type DatabaseService struct {
	Options []option.RequestOption
}

// NewDatabaseService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewDatabaseService(opts ...option.RequestOption) (r *DatabaseService) {
	r = &DatabaseService{}
	r.Options = opts
	return
}

// Update KV database details
func (r *DatabaseService) Update(ctx context.Context, databaseID string, body DatabaseUpdateParams, opts ...option.RequestOption) (res *shared.KvDatabase, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("databases/%s", databaseID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

type DatabaseUpdateParams struct {
	// The description of the KV database to be updated to. If this is `null`, no
	// update will be made to the KV database description.
	Description param.Field[string] `json:"description"`
}

func (r DatabaseUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
