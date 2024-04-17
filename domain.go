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

// DomainService contains methods and other services that help with interacting
// with the subhosting API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewDomainService] method instead.
type DomainService struct {
	Options      []option.RequestOption
	Certificates *DomainCertificateService
}

// NewDomainService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewDomainService(opts ...option.RequestOption) (r *DomainService) {
	r = &DomainService{}
	r.Options = opts
	r.Certificates = NewDomainCertificateService(opts...)
	return
}

// Associate a domain with a deployment
//
// This API allows you to either:
//
// 1. associate a domain with a deployment, or
// 2. disassociate a domain from a deployment
//
// Domain association is required in order to serve the deployment on the domain.
//
// If the ownership of the domain is not verified yet, this API will trigger the
// verification process before associating the domain with the deployment.
func (r *DomainService) Update(ctx context.Context, domainID string, body DomainUpdateParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := fmt.Sprintf("domains/%s", domainID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, nil, opts...)
	return
}

// Delete a domain
func (r *DomainService) Delete(ctx context.Context, domainID string, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := fmt.Sprintf("domains/%s", domainID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Get domain details
func (r *DomainService) Get(ctx context.Context, domainID string, opts ...option.RequestOption) (res *shared.Domain, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("domains/%s", domainID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Verify ownership of a domain
//
// This API triggers the ownership verification of a domain. It should be called
// after necessary DNS records are properly set up.
func (r *DomainService) Verify(ctx context.Context, domainID string, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := fmt.Sprintf("domains/%s/verify", domainID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, nil, opts...)
	return
}

type DomainUpdateParams struct {
	// A deployment ID
	//
	// Note that this is not UUID v4, as opposed to organization ID and project ID.
	DeploymentID param.Field[string] `json:"deploymentId"`
}

func (r DomainUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
