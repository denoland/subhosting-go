// File generated from our OpenAPI spec by Stainless.

package shared

import (
	"reflect"
	"time"

	"github.com/denoland/subhosting-go/internal/apijson"
	"github.com/tidwall/gjson"
)

// Project analytics data
type Analytics struct {
	Fields []AnalyticsField   `json:"fields,required"`
	Values [][]AnalyticsValue `json:"values,required" format:"date-time"`
	JSON   analyticsJSON      `json:"-"`
}

// analyticsJSON contains the JSON metadata for the struct [Analytics]
type analyticsJSON struct {
	Fields      apijson.Field
	Values      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Analytics) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AnalyticsField struct {
	Name string `json:"name,required"`
	// A data type that analytic data can be represented in.
	//
	// Inspired by Grafana's data types defined at:
	// https://github.com/grafana/grafana/blob/e3288834b37b9aac10c1f43f0e621b35874c1f8a/packages/grafana-data/src/types/dataFrame.ts#L11-L23
	Type AnalyticsFieldsType `json:"type,required"`
	JSON analyticsFieldJSON  `json:"-"`
}

// analyticsFieldJSON contains the JSON metadata for the struct [AnalyticsField]
type analyticsFieldJSON struct {
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AnalyticsField) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A data type that analytic data can be represented in.
//
// Inspired by Grafana's data types defined at:
// https://github.com/grafana/grafana/blob/e3288834b37b9aac10c1f43f0e621b35874c1f8a/packages/grafana-data/src/types/dataFrame.ts#L11-L23
type AnalyticsFieldsType string

const (
	AnalyticsFieldsTypeTime    AnalyticsFieldsType = "time"
	AnalyticsFieldsTypeNumber  AnalyticsFieldsType = "number"
	AnalyticsFieldsTypeString  AnalyticsFieldsType = "string"
	AnalyticsFieldsTypeBoolean AnalyticsFieldsType = "boolean"
	AnalyticsFieldsTypeOther   AnalyticsFieldsType = "other"
)

// Union satisfied by [shared.UnionTime], [shared.UnionFloat],
// [shared.UnionString], [shared.UnionBool] or [shared.AnalyticsValuesUnknown].
type AnalyticsValue interface {
	ImplementsSharedAnalyticsValue()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AnalyticsValue)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(UnionTime(UnionTime{})),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(UnionString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.True,
			Type:       reflect.TypeOf(UnionBool(false)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.False,
			Type:       reflect.TypeOf(UnionBool(false)),
		},
	)
}

type Deployment struct {
	// A deployment ID
	//
	// Note that this is not UUID v4, as opposed to organization ID and project ID.
	ID        string    `json:"id,required"`
	CreatedAt time.Time `json:"createdAt,required" format:"date-time"`
	// The KV databases that this deployment has access to. Currently, only `"default"`
	// database is supported.
	Databases map[string]string `json:"databases,required" format:"uuid"`
	ProjectID string            `json:"projectId,required" format:"uuid"`
	// The status of a deployment.
	Status    DeploymentStatus `json:"status,required"`
	UpdatedAt time.Time        `json:"updatedAt,required" format:"date-time"`
	// The description of this deployment. This is present only when the `status` is
	// `success`.
	Description string         `json:"description,nullable"`
	Domains     []string       `json:"domains,nullable"`
	JSON        deploymentJSON `json:"-"`
}

// deploymentJSON contains the JSON metadata for the struct [Deployment]
type deploymentJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Databases   apijson.Field
	ProjectID   apijson.Field
	Status      apijson.Field
	UpdatedAt   apijson.Field
	Description apijson.Field
	Domains     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Deployment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of a deployment.
type DeploymentStatus string

const (
	DeploymentStatusFailed  DeploymentStatus = "failed"
	DeploymentStatusPending DeploymentStatus = "pending"
	DeploymentStatusSuccess DeploymentStatus = "success"
)

type Domain struct {
	// The ID of the domain.
	ID    string `json:"id,required" format:"uuid"`
	Token string `json:"token,required"`
	// TLS certificates for the domain.
	Certificates []DomainCertificate `json:"certificates,required"`
	CreatedAt    time.Time           `json:"createdAt,required" format:"date-time"`
	// These records are used to verify the ownership of the domain.
	DNSRecords []DomainDNSRecord `json:"dnsRecords,required"`
	// The domain value.
	Domain string `json:"domain,required"`
	// Whether the domain's ownership is validated or not.
	IsValidated bool `json:"isValidated,required"`
	// The ID of the organization that the domain is associated with.
	OrganizationID     string                   `json:"organizationId,required" format:"uuid"`
	ProvisioningStatus DomainProvisioningStatus `json:"provisioningStatus,required"`
	UpdatedAt          time.Time                `json:"updatedAt,required" format:"date-time"`
	// A deployment ID
	//
	// Note that this is not UUID v4, as opposed to organization ID and project ID.
	DeploymentID string `json:"deploymentId"`
	// The ID of the project that the domain is associated with.
	//
	// If the domain is not associated with any project, this field is omitted.
	ProjectID string     `json:"projectId,nullable" format:"uuid"`
	JSON      domainJSON `json:"-"`
}

// domainJSON contains the JSON metadata for the struct [Domain]
type domainJSON struct {
	ID                 apijson.Field
	Token              apijson.Field
	Certificates       apijson.Field
	CreatedAt          apijson.Field
	DNSRecords         apijson.Field
	Domain             apijson.Field
	IsValidated        apijson.Field
	OrganizationID     apijson.Field
	ProvisioningStatus apijson.Field
	UpdatedAt          apijson.Field
	DeploymentID       apijson.Field
	ProjectID          apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *Domain) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DomainCertificate struct {
	Cipher    DomainCertificatesCipher `json:"cipher,required"`
	CreatedAt time.Time                `json:"createdAt,required" format:"date-time"`
	ExpiresAt time.Time                `json:"expiresAt,required" format:"date-time"`
	UpdatedAt time.Time                `json:"updatedAt,required" format:"date-time"`
	JSON      domainCertificateJSON    `json:"-"`
}

// domainCertificateJSON contains the JSON metadata for the struct
// [DomainCertificate]
type domainCertificateJSON struct {
	Cipher      apijson.Field
	CreatedAt   apijson.Field
	ExpiresAt   apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DomainCertificate) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DomainCertificatesCipher string

const (
	DomainCertificatesCipherRsa DomainCertificatesCipher = "rsa"
	DomainCertificatesCipherEc  DomainCertificatesCipher = "ec"
)

type DomainDNSRecord struct {
	Content string              `json:"content,required"`
	Name    string              `json:"name,required"`
	Type    string              `json:"type,required"`
	JSON    domainDNSRecordJSON `json:"-"`
}

// domainDNSRecordJSON contains the JSON metadata for the struct [DomainDNSRecord]
type domainDNSRecordJSON struct {
	Content     apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DomainDNSRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Union satisfied by [shared.DomainProvisioningStatusSuccess],
// [shared.DomainProvisioningStatusFailed],
// [shared.DomainProvisioningStatusPending] or
// [shared.DomainProvisioningStatusManual].
type DomainProvisioningStatus interface {
	implementsSharedDomainProvisioningStatus()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DomainProvisioningStatus)(nil)).Elem(),
		"code",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DomainProvisioningStatusSuccess{}),
			DiscriminatorValue: "success",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DomainProvisioningStatusFailed{}),
			DiscriminatorValue: "failed",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DomainProvisioningStatusPending{}),
			DiscriminatorValue: "pending",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DomainProvisioningStatusManual{}),
			DiscriminatorValue: "manual",
		},
	)
}

type DomainProvisioningStatusSuccess struct {
	Code DomainProvisioningStatusSuccessCode `json:"code,required"`
	JSON domainProvisioningStatusSuccessJSON `json:"-"`
}

// domainProvisioningStatusSuccessJSON contains the JSON metadata for the struct
// [DomainProvisioningStatusSuccess]
type domainProvisioningStatusSuccessJSON struct {
	Code        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DomainProvisioningStatusSuccess) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DomainProvisioningStatusSuccess) implementsSharedDomainProvisioningStatus() {}

type DomainProvisioningStatusSuccessCode string

const (
	DomainProvisioningStatusSuccessCodeSuccess DomainProvisioningStatusSuccessCode = "success"
)

type DomainProvisioningStatusFailed struct {
	Code    DomainProvisioningStatusFailedCode `json:"code,required"`
	Message string                             `json:"message,required"`
	JSON    domainProvisioningStatusFailedJSON `json:"-"`
}

// domainProvisioningStatusFailedJSON contains the JSON metadata for the struct
// [DomainProvisioningStatusFailed]
type domainProvisioningStatusFailedJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DomainProvisioningStatusFailed) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DomainProvisioningStatusFailed) implementsSharedDomainProvisioningStatus() {}

type DomainProvisioningStatusFailedCode string

const (
	DomainProvisioningStatusFailedCodeFailed DomainProvisioningStatusFailedCode = "failed"
)

type DomainProvisioningStatusPending struct {
	Code DomainProvisioningStatusPendingCode `json:"code,required"`
	JSON domainProvisioningStatusPendingJSON `json:"-"`
}

// domainProvisioningStatusPendingJSON contains the JSON metadata for the struct
// [DomainProvisioningStatusPending]
type domainProvisioningStatusPendingJSON struct {
	Code        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DomainProvisioningStatusPending) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DomainProvisioningStatusPending) implementsSharedDomainProvisioningStatus() {}

type DomainProvisioningStatusPendingCode string

const (
	DomainProvisioningStatusPendingCodePending DomainProvisioningStatusPendingCode = "pending"
)

type DomainProvisioningStatusManual struct {
	Code DomainProvisioningStatusManualCode `json:"code,required"`
	JSON domainProvisioningStatusManualJSON `json:"-"`
}

// domainProvisioningStatusManualJSON contains the JSON metadata for the struct
// [DomainProvisioningStatusManual]
type domainProvisioningStatusManualJSON struct {
	Code        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DomainProvisioningStatusManual) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DomainProvisioningStatusManual) implementsSharedDomainProvisioningStatus() {}

type DomainProvisioningStatusManualCode string

const (
	DomainProvisioningStatusManualCodeManual DomainProvisioningStatusManualCode = "manual"
)

type KvDatabase struct {
	// A KV database ID
	ID        string    `json:"id,required" format:"uuid"`
	CreatedAt time.Time `json:"createdAt,required" format:"date-time"`
	// A description of this KV database
	Description string `json:"description,required"`
	// An organization ID that this KV database belongs to
	OrganizationID string         `json:"organizationId,required" format:"uuid"`
	UpdatedAt      time.Time      `json:"updatedAt,required" format:"date-time"`
	JSON           kvDatabaseJSON `json:"-"`
}

// kvDatabaseJSON contains the JSON metadata for the struct [KvDatabase]
type kvDatabaseJSON struct {
	ID             apijson.Field
	CreatedAt      apijson.Field
	Description    apijson.Field
	OrganizationID apijson.Field
	UpdatedAt      apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *KvDatabase) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type Project struct {
	ID          string      `json:"id,required" format:"uuid"`
	CreatedAt   time.Time   `json:"createdAt,required" format:"date-time"`
	Description string      `json:"description,required"`
	Name        string      `json:"name,required"`
	UpdatedAt   time.Time   `json:"updatedAt,required" format:"date-time"`
	JSON        projectJSON `json:"-"`
}

// projectJSON contains the JSON metadata for the struct [Project]
type projectJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Description apijson.Field
	Name        apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Project) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
