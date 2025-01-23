package input_options

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	"terraform-provider-trocco/internal/client/entities/job_definitions"
	parmas "terraform-provider-trocco/internal/client/parameters/job_definitions"
)

type Decoder struct {
	MatchName types.String `tfsdk:"match_name"`
}

func NewDecoder(decoder *job_definitions.Decoder) *Decoder {
	if decoder == nil {
		return nil
	}
	return &Decoder{
		MatchName: types.StringValue(decoder.MatchName),
	}
}

func (decoder *Decoder) ToDecoderInput() *parmas.DecoderInput {
	if decoder == nil {
		return nil
	}
	return &parmas.DecoderInput{
		MatchName: decoder.MatchName.ValueString(),
	}
}
