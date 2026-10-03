// Hand-written gap fillers for types openapi-generator does not emit
// from this spec. Referenced by generated files; do not delete.

package mistral

import "encoding/json"

// ImageURL is components.schemas.ImageURL, skipped by the generator
// (case-insensitive collision with the ImageUrl anyOf wrapper).
type ImageURL struct {
	Url    string       `json:"url"`
	Detail *ImageDetail `json:"detail,omitempty"`
}

// AnyOfmapnull is the inline anyOf [object, null] used for metadata filters.
type AnyOfmapnull = map[string]any

// AnyOf is the inline anyOf [{}, null]: an arbitrary nullable JSON value.
type AnyOf = any

// NullableAnyOf mirrors the generator's Nullable* wrappers for AnyOf.
type NullableAnyOf struct {
	value *AnyOf
	isSet bool
}

func (v NullableAnyOf) Get() *AnyOf { return v.value }
func (v NullableAnyOf) IsSet() bool { return v.isSet }
func (v *NullableAnyOf) Set(val *AnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableAnyOf) MarshalJSON() ([]byte, error) {
	if !v.isSet || v.value == nil {
		return []byte("null"), nil
	}
	return json.Marshal(*v.value)
}

func (v *NullableAnyOf) UnmarshalJSON(data []byte) error {
	v.isSet = true
	var val AnyOf
	if err := json.Unmarshal(data, &val); err != nil {
		return err
	}
	v.value = &val
	return nil
}
