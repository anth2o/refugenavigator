package tests

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/anth2o/refugenavigator/internal/scrapper"
)

type TestStruct struct {
	Value scrapper.FlexibleInt `json:"value"`
}

func TestFlexibleInt_UnmarshalJSON(t *testing.T) {
	tests := []struct {
		name     string
		jsonStr  string
		want     int
		wantErr  bool
		errorMsg string
	}{
		{
			name:    "integer value",
			jsonStr: `{"value": 42}`,
			want:    42,
		},
		{
			name:    "string value",
			jsonStr: `{"value": "42"}`,
			want:    42,
		},
		{
			name:    "negative integer",
			jsonStr: `{"value": -100}`,
			want:    -100,
		},
		{
			name:    "negative string",
			jsonStr: `{"value": "-100"}`,
			want:    -100,
		},
		{
			name:     "invalid string",
			jsonStr:  `{"value": "not a number"}`,
			wantErr:  true,
			errorMsg: "value should be a number or a string representation of a number",
		},
		{
			name:     "float value",
			jsonStr:  `{"value": 3.14}`,
			wantErr:  true,
			errorMsg: "json: cannot unmarshal number into Go struct field TestStruct.value of type string",
		},
		{
			name:     "invalid json",
			jsonStr:  `invalid json`,
			wantErr:  true,
			errorMsg: "invalid character 'i' looking for beginning of value",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var testStruct TestStruct
			unmarshalErr := json.Unmarshal([]byte(tt.jsonStr), &testStruct)
			if unmarshalErr != nil {
				if !tt.wantErr {
					t.Errorf("UnmarshalJSON() unexpected error: %v", unmarshalErr)
					return
				}
			} else {
				if tt.wantErr {
					t.Errorf("UnmarshalJSON() expected error but got none")
					return
				}
			}
			if tt.wantErr && tt.errorMsg != "" {
				if !strings.Contains(unmarshalErr.Error(), tt.errorMsg) {
					t.Errorf("UnmarshalJSON() error message = %v, want %v", unmarshalErr.Error(), tt.errorMsg)
				}
			}
			if testStruct.Value != scrapper.FlexibleInt(tt.want) {
				t.Errorf("UnmarshalJSON() got = %v, want %v", testStruct.Value, tt.want)
			}
		})
	}
}
