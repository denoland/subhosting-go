// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package shared

import (
	"reflect"
	"time"

	"github.com/denoland/subhosting-go/internal/apijson"
	"github.com/tidwall/gjson"
)

// Project analytics data
type Analytics struct {
	Fields []AnalyticsField         `json:"fields,required"`
	Values [][]AnalyticsValuesUnion `json:"values,required" format:"date-time"`
	JSON   analyticsJSON            `json:"-"`
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

func (r analyticsJSON) RawJSON() string {
	return r.raw
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

func (r analyticsFieldJSON) RawJSON() string {
	return r.raw
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

func (r AnalyticsFieldsType) IsKnown() bool {
	switch r {
	case AnalyticsFieldsTypeTime, AnalyticsFieldsTypeNumber, AnalyticsFieldsTypeString, AnalyticsFieldsTypeBoolean, AnalyticsFieldsTypeOther:
		return true
	}
	return false
}

// Union satisfied by [shared.UnionTime], [shared.UnionFloat],
// [shared.UnionString], [shared.UnionBool] or [shared.AnalyticsValuesUnknown].
type AnalyticsValuesUnion interface {
	ImplementsSharedAnalyticsValuesUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AnalyticsValuesUnion)(nil)).Elem(),
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

func (r deploymentJSON) RawJSON() string {
	return r.raw
}

// The status of a deployment.
type DeploymentStatus string

const (
	DeploymentStatusFailed  DeploymentStatus = "failed"
	DeploymentStatusPending DeploymentStatus = "pending"
	DeploymentStatusSuccess DeploymentStatus = "success"
)

func (r DeploymentStatus) IsKnown() bool {
	switch r {
	case DeploymentStatusFailed, DeploymentStatusPending, DeploymentStatusSuccess:
		return true
	}
	return false
}

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
	DeploymentID string `json:"deploymentId,nullable"`
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

func (r domainJSON) RawJSON() string {
	return r.raw
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

func (r domainCertificateJSON) RawJSON() string {
	return r.raw
}

type DomainCertificatesCipher string

const (
	DomainCertificatesCipherRsa DomainCertificatesCipher = "rsa"
	DomainCertificatesCipherEc  DomainCertificatesCipher = "ec"
)

func (r DomainCertificatesCipher) IsKnown() bool {
	switch r {
	case DomainCertificatesCipherRsa, DomainCertificatesCipherEc:
		return true
	}
	return false
}

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

func (r domainDNSRecordJSON) RawJSON() string {
	return r.raw
}

type DomainProvisioningStatus struct {
	Code    DomainProvisioningStatusCode `json:"code,required"`
	Message string                       `json:"message"`
	JSON    domainProvisioningStatusJSON `json:"-"`
	union   DomainProvisioningStatusUnion
}

// domainProvisioningStatusJSON contains the JSON metadata for the struct
// [DomainProvisioningStatus]
type domainProvisioningStatusJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r domainProvisioningStatusJSON) RawJSON() string {
	return r.raw
}

func (r *DomainProvisioningStatus) UnmarshalJSON(data []byte) (err error) {
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

func (r DomainProvisioningStatus) AsUnion() DomainProvisioningStatusUnion {
	return r.union
}

// Union satisfied by [shared.DomainProvisioningStatusObject],
// [shared.DomainProvisioningStatusObject], [shared.DomainProvisioningStatusObject]
// or [shared.DomainProvisioningStatusObject].
type DomainProvisioningStatusUnion interface {
	implementsSharedDomainProvisioningStatus()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DomainProvisioningStatusUnion)(nil)).Elem(),
		"code",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DomainProvisioningStatusObject{}),
			DiscriminatorValue: "success",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DomainProvisioningStatusObject{}),
			DiscriminatorValue: "failed",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DomainProvisioningStatusObject{}),
			DiscriminatorValue: "pending",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DomainProvisioningStatusObject{}),
			DiscriminatorValue: "manual",
		},
	)
}

type DomainProvisioningStatusObject struct {
	Code DomainProvisioningStatusObjectCode `json:"code,required"`
	JSON domainProvisioningStatusObjectJSON `json:"-"`
}

// domainProvisioningStatusObjectJSON contains the JSON metadata for the struct
// [DomainProvisioningStatusObject]
type domainProvisioningStatusObjectJSON struct {
	Code        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DomainProvisioningStatusObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r domainProvisioningStatusObjectJSON) RawJSON() string {
	return r.raw
}

func (r DomainProvisioningStatusObject) implementsSharedDomainProvisioningStatus() {}

type DomainProvisioningStatusObjectCode string

const (
	DomainProvisioningStatusObjectCodeSuccess DomainProvisioningStatusObjectCode = "success"
)

func (r DomainProvisioningStatusObjectCode) IsKnown() bool {
	switch r {
	case DomainProvisioningStatusObjectCodeSuccess:
		return true
	}
	return false
}

type DomainProvisioningStatusCode string

const (
	DomainProvisioningStatusCodeSuccess DomainProvisioningStatusCode = "success"
	DomainProvisioningStatusCodeFailed  DomainProvisioningStatusCode = "failed"
	DomainProvisioningStatusCodePending DomainProvisioningStatusCode = "pending"
	DomainProvisioningStatusCodeManual  DomainProvisioningStatusCode = "manual"
)

func (r DomainProvisioningStatusCode) IsKnown() bool {
	switch r {
	case DomainProvisioningStatusCodeSuccess, DomainProvisioningStatusCodeFailed, DomainProvisioningStatusCodePending, DomainProvisioningStatusCodeManual:
		return true
	}
	return false
}

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

func (r kvDatabaseJSON) RawJSON() string {
	return r.raw
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

func (r projectJSON) RawJSON() string {
	return r.raw
}
