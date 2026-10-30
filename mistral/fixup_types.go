// Hand-written gap fillers for types openapi-generator does not emit
// from this spec. Referenced by generated files; do not delete.

package mistral

import (
	"encoding/json"
	"fmt"
	"os"
)

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

// AnyOfmapAnyType is the inline anyOf [map, any] used for arbitrary
// nullable objects in request bodies.
type AnyOfmapAnyType = map[string]any

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

func (v *NullableAnyOf) Unset() {
	v.value = nil
	v.isSet = false
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

// NullableOsFile mirrors the generator's Nullable* wrappers for *os.File,
// which it emits with an invalid Nullable*os.File identifier.
type NullableOsFile struct {
	value *os.File
	isSet bool
}

func (v NullableOsFile) Get() *os.File { return v.value }
func (v NullableOsFile) IsSet() bool   { return v.isSet }
func (v *NullableOsFile) Set(val *os.File) {
	v.value = val
	v.isSet = true
}

func (v *NullableOsFile) Unset() {
	v.value = nil
	v.isSet = false
}

func (v NullableOsFile) MarshalJSON() ([]byte, error) {
	// os.File is not JSON-encodable; uploads go through multipart forms.
	return []byte("null"), nil
}

func (v *NullableOsFile) UnmarshalJSON(data []byte) error {
	v.isSet = true
	if string(data) == "null" {
		v.value = nil
		return nil
	}
	return fmt.Errorf("NullableOsFile only accepts null; binary uploads are not JSON-encodable")
}
