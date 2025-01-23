package models

import (
	"terraform-provider-trocco/internal/client/parameter"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

func NewNullableBool(v types.Bool) *parameter.NullableBool {
	if v.IsUnknown() {
		return nil
	}

	return &parameter.NullableBool{Valid: !v.IsNull(), Value: v.ValueBool()}
}

func NewNullableInt64(v types.Int64) *parameter.NullableInt64 {
	if v.IsUnknown() {
		return nil
	}

	return &parameter.NullableInt64{Valid: !v.IsNull(), Value: v.ValueInt64()}
}

func NewNullableString(v types.String) *parameter.NullableString {
	if v.IsUnknown() {
		return nil
	}

	return &parameter.NullableString{Valid: !v.IsNull(), Value: v.ValueString()}
}
