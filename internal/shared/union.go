// File generated from our OpenAPI spec by Stainless.

package shared

import (
	"time"
)

type UnionTime time.Time

func (UnionTime) ImplementsSharedAnalyticsValue() {}

type UnionString string

func (UnionString) ImplementsSharedAnalyticsValue() {}

type UnionBool bool

func (UnionBool) ImplementsSharedAnalyticsValue() {}

type UnionFloat float64

func (UnionFloat) ImplementsSharedAnalyticsValue() {}
