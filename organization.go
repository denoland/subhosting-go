// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package subhosting

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/denoland/subhosting-go/internal/apijson"
	"github.com/denoland/subhosting-go/internal/requestconfig"
	"github.com/denoland/subhosting-go/option"
)

// OrganizationService contains methods and other services that help with
// interacting with the subhosting API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewOrganizationService] method
// instead.
type OrganizationService struct {
	Options   []option.RequestOption
	Analytics *OrganizationAnalyticsService
	Projects  *OrganizationProjectService
	Databases *OrganizationDatabaseService
	Domains   *OrganizationDomainService
}

// NewOrganizationService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewOrganizationService(opts ...option.RequestOption) (r *OrganizationService) {
	r = &OrganizationService{}
	r.Options = opts
	r.Analytics = NewOrganizationAnalyticsService(opts...)
	r.Projects = NewOrganizationProjectService(opts...)
	r.Databases = NewOrganizationDatabaseService(opts...)
	r.Domains = NewOrganizationDomainService(opts...)
	return
}

// Get organization details
func (r *OrganizationService) Get(ctx context.Context, organizationID string, opts ...option.RequestOption) (res *Organization, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("organizations/%s", organizationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type Organization struct {
	ID        string           `json:"id,required" format:"uuid"`
	CreatedAt time.Time        `json:"createdAt,required" format:"date-time"`
	Name      string           `json:"name,required"`
	UpdatedAt time.Time        `json:"updatedAt,required" format:"date-time"`
	JSON      organizationJSON `json:"-"`
}

// organizationJSON contains the JSON metadata for the struct [Organization]
type organizationJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Name        apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Organization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r organizationJSON) RawJSON() string {
	return r.raw
}
