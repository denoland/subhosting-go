// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package subhosting

import (
	"github.com/denoland/subhosting-go/internal/apierror"
	"github.com/denoland/subhosting-go/shared"
)

type Error = apierror.Error

// Project analytics data
//
// This is an alias to an internal type.
type Analytics = shared.Analytics

// This is an alias to an internal type.
type AnalyticsField = shared.AnalyticsField

// A data type that analytic data can be represented in.
//
// Inspired by Grafana's data types defined at:
// https://github.com/grafana/grafana/blob/e3288834b37b9aac10c1f43f0e621b35874c1f8a/packages/grafana-data/src/types/dataFrame.ts#L11-L23
//
// This is an alias to an internal type.
type AnalyticsFieldsType = shared.AnalyticsFieldsType

// This is an alias to an internal value.
const AnalyticsFieldsTypeTime = shared.AnalyticsFieldsTypeTime

// This is an alias to an internal value.
const AnalyticsFieldsTypeNumber = shared.AnalyticsFieldsTypeNumber

// This is an alias to an internal value.
const AnalyticsFieldsTypeString = shared.AnalyticsFieldsTypeString

// This is an alias to an internal value.
const AnalyticsFieldsTypeBoolean = shared.AnalyticsFieldsTypeBoolean

// This is an alias to an internal value.
const AnalyticsFieldsTypeOther = shared.AnalyticsFieldsTypeOther

// This is an alias to an internal type.
type AnalyticsValuesUnion = shared.AnalyticsValuesUnion

// This is an alias to an internal type.
type Deployment = shared.Deployment

// The status of a deployment.
//
// This is an alias to an internal type.
type DeploymentStatus = shared.DeploymentStatus

// This is an alias to an internal value.
const DeploymentStatusFailed = shared.DeploymentStatusFailed

// This is an alias to an internal value.
const DeploymentStatusPending = shared.DeploymentStatusPending

// This is an alias to an internal value.
const DeploymentStatusSuccess = shared.DeploymentStatusSuccess

// This is an alias to an internal type.
type Domain = shared.Domain

// This is an alias to an internal type.
type DomainCertificate = shared.DomainCertificate

// This is an alias to an internal type.
type DomainCertificatesCipher = shared.DomainCertificatesCipher

// This is an alias to an internal value.
const DomainCertificatesCipherRsa = shared.DomainCertificatesCipherRsa

// This is an alias to an internal value.
const DomainCertificatesCipherEc = shared.DomainCertificatesCipherEc

// This is an alias to an internal type.
type DomainDNSRecord = shared.DomainDNSRecord

// This is an alias to an internal type.
type DomainProvisioningStatus = shared.DomainProvisioningStatus

// This is an alias to an internal type.
type DomainProvisioningStatusObject = shared.DomainProvisioningStatusObject

// This is an alias to an internal type.
type DomainProvisioningStatusObjectCode = shared.DomainProvisioningStatusObjectCode

// This is an alias to an internal value.
const DomainProvisioningStatusObjectCodeSuccess = shared.DomainProvisioningStatusObjectCodeSuccess

// This is an alias to an internal type.
type DomainProvisioningStatusCode = shared.DomainProvisioningStatusCode

// This is an alias to an internal value.
const DomainProvisioningStatusCodeSuccess = shared.DomainProvisioningStatusCodeSuccess

// This is an alias to an internal value.
const DomainProvisioningStatusCodeFailed = shared.DomainProvisioningStatusCodeFailed

// This is an alias to an internal value.
const DomainProvisioningStatusCodePending = shared.DomainProvisioningStatusCodePending

// This is an alias to an internal value.
const DomainProvisioningStatusCodeManual = shared.DomainProvisioningStatusCodeManual

// This is an alias to an internal type.
type KvDatabase = shared.KvDatabase

// This is an alias to an internal type.
type Project = shared.Project
