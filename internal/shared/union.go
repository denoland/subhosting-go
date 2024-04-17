// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package shared

import (
	"time"
)

type UnionTime time.Time

func (UnionTime) ImplementsSharedAnalyticsValuesUnion() {}

type UnionString string

func (UnionString) ImplementsSharedAnalyticsValuesUnion() {}

type UnionBool bool

func (UnionBool) ImplementsSharedAnalyticsValuesUnion() {}

type UnionFloat float64

func (UnionFloat) ImplementsSharedAnalyticsValuesUnion() {}
