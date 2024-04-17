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
)

// DomainCertificateService contains methods and other services that help with
// interacting with the subhosting API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewDomainCertificateService] method
// instead.
type DomainCertificateService struct {
	Options []option.RequestOption
}

// NewDomainCertificateService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewDomainCertificateService(opts ...option.RequestOption) (r *DomainCertificateService) {
	r = &DomainCertificateService{}
	r.Options = opts
	return
}

// Upload TLS certificate for a domain
//
// This API allows you to upload a TLS certificate for a domain.
//
// If the ownership of the domain is not verified yet, this API will trigger the
// verification process before storing the certificate.
func (r *DomainCertificateService) New(ctx context.Context, domainID string, body DomainCertificateNewParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := fmt.Sprintf("domains/%s/certificates", domainID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Provision TLS certificates for a domain
//
// This API begins the provisioning of TLS certificates for a domain.
//
// Note that a call to this API may take a while, up to a minute or so.
//
// If the ownership of the domain is not verified yet, this API will trigger the
// verification process before provisioning the certificate.
func (r *DomainCertificateService) Provision(ctx context.Context, domainID string, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := fmt.Sprintf("domains/%s/certificates/provision", domainID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, nil, opts...)
	return
}

type DomainCertificateNewParams struct {
	// The PRM encoded certificate chain for the TLS certificate
	CertificateChain param.Field[string] `json:"certificateChain,required"`
	// The PEM encoded private key for the TLS certificate
	PrivateKey param.Field[string] `json:"privateKey,required"`
}

func (r DomainCertificateNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
