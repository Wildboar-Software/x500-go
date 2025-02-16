package x500

import (
	"encoding/asn1"
	"errors"
	"fmt"
	"reflect"
	"strings"
	"unicode"
	"unicode/utf8"

	"github.com/Wildboar-Software/x500-go/teletex"
)

// An invalidUnmarshalError describes an invalid argument passed to Unmarshal.
// (The argument to Unmarshal must be a non-nil pointer.)
type invalidUnmarshalError struct {
	Type reflect.Type
}

func (e *invalidUnmarshalError) Error() string {
	if e.Type == nil {
		return "asn1: Unmarshal recipient value is nil"
	}

	if e.Type.Kind() != reflect.Pointer {
		return "asn1: Unmarshal recipient value is non-pointer " + e.Type.String()
	}
	return "asn1: Unmarshal recipient value is nil " + e.Type.String()
}

func isASCII(s []byte) bool {
	for i := 0; i < len(s); i++ {
		if s[i] > unicode.MaxASCII {
			return false
		}
	}
	return true
}

func isASN1TimeChar(b byte) bool {
	return (b >= '0' && b <= '9') ||
		b == '+' || b == ',' || b == '-' || b == '.' ||
		b == '/' || b == ':' || strings.ContainsRune("CDHMRPSTWYZ", rune(b))
}

func unmarshalTimeString(octets []byte) (string, error) {
	for i, c := range octets {
		if !isASN1TimeChar(c) {
			return "", fmt.Errorf("malformed asn.1 time string: invalid char '%c' at index %d", c, i)
		}
	}
	return string(octets), nil
}

func unmarshalTimeOfDayAsString(octets []byte) (string, error) {
	if len(octets) != 8 {
		return "", fmt.Errorf("malformed asn.1 time-of-day: %d characters instead of 8", len(octets))
	}
	for i, c := range octets {
		if i == 2 || i == 5 {
			if c != ':' {
				return "", fmt.Errorf("malformed asn.1 time-of-day: expected a colon but got '%c'", c)
			} else {
				continue
			}
		}
		if !unicode.IsDigit(rune(c)) {
			return "", fmt.Errorf("malformed asn.1 time-of-day: expected a digit but got '%c' at index %d", c, i)
		}
	}
	return string(octets), nil
}

func unmarshalDateAsString(octets []byte) (string, error) {
	if len(octets) != 10 {
		return "", fmt.Errorf("malformed asn.1 date: %d characters instead of 10", len(octets))
	}
	for i, c := range octets {
		if i == 4 || i == 7 {
			if c != '-' {
				return "", fmt.Errorf("malformed asn.1 date: expected a hyphen but got '%c'", c)
			} else {
				continue
			}
		}
		if !unicode.IsDigit(rune(c)) {
			return "", fmt.Errorf("malformed asn.1 date: expected a digit but got '%c' at index %d", c, i)
		}
	}
	return string(octets), nil
}

// 1951-10-14T15:30:00
func unmarshalDateTimeAsString(octets []byte) (string, error) {
	if len(octets) != 19 {
		return "", fmt.Errorf("malformed asn.1 datetime: %d characters instead of 19", len(octets))
	}
	for i, c := range octets {
		if i == 4 || i == 7 {
			if c != '-' {
				return "", fmt.Errorf("malformed asn.1 datetime: expected a hyphen but got '%c'", c)
			} else {
				continue
			}
		}
		if i == 10 {
			if c != 'T' {
				return "", fmt.Errorf("malformed asn.1 datetime: expected a 'T' but got '%c'", c)
			} else {
				continue
			}
		}
		if i == 13 || i == 16 {
			if c != ':' {
				return "", fmt.Errorf("malformed asn.1 datetime: expected a colon but got '%c'", c)
			} else {
				continue
			}
		}
		if !unicode.IsDigit(rune(c)) {
			return "", fmt.Errorf("malformed asn.1 datetime: expected a digit but got '%c' at index %d", c, i)
		}
	}
	return string(octets), nil
}

func unmarshalValue(v reflect.Value, encoded asn1.RawValue, params fieldParameters) (err error) {
	if params.tag != 0 && params.tag != encoded.Tag {
		return fmt.Errorf("unexpected tag: expected %d but got %d", params.tag, encoded.Tag)
	}
	k := v.Kind()
	switch k {
	case reflect.Bool:
		fallthrough
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		rest, err := asn1.Unmarshal(encoded.FullBytes, v.Addr().Interface())
		if err != nil {
			return err
		}
		if len(rest) > 0 {
			return errors.New("trailing data")
		}
	case reflect.String:
		var s string
		switch encoded.Tag {
		case tagTime:
			s, err = unmarshalTimeString(encoded.Bytes)
		case tagDateTime:
			s, err = unmarshalDateTimeAsString(encoded.Bytes)
		case tagDate:
			s, err = unmarshalDateAsString(encoded.Bytes)
		case tagTimeOfDay:
			s, err = unmarshalTimeOfDayAsString(encoded.Bytes)
		case tagDuration:
			if len(encoded.Bytes) < 3 {
				return fmt.Errorf("asn.1 duration too short: encoded on %d bytes", len(encoded.Bytes))
			}
			fallthrough
		case tagOidIri:
			s = string(encoded.Bytes)
			if !utf8.ValidString(s) {
				return errors.New("malformed relative oid-iri string")
			}
		case tagRelativeOidIri:
			s = string(encoded.Bytes)
			if !utf8.ValidString(s) {
				return errors.New("malformed oid-iri string")
			}
		case asn1.TagT61String:
			s = teletex.TeletexToUTF8(encoded.Bytes)
		case tagUniversalString:
			s, err = universalStringFromBytes(encoded.Bytes)
		case tagGraphicString:
			fallthrough
		case tagGeneralString:
			fallthrough
		case tagVisibleString:
			fallthrough
		case tagVideotexString:
			if !isASCII(encoded.Bytes) {
				return fmt.Errorf("cannot decode non-ASCII string with tag %d", encoded.Tag)
			}
			s = string(encoded.Bytes)
		default: // All other string types are already handled by golang's asn1
			// https://cs.opensource.google/go/go/+/refs/tags/go1.24.0:src/encoding/asn1/asn1.go;l=699
			rest, err := asn1.Unmarshal(encoded.FullBytes, v.Addr().Interface())
			if err != nil {
				return err
			}
			if len(rest) > 0 {
				return errors.New("trailing data")
			}
			return nil
		}
		if err != nil {
			return err
		}
		v.SetString(s)
	case reflect.Slice:
		sliceType := v.Type()
		if sliceType.Elem().Kind() == reflect.Uint8 {
			b := v.Bytes()
			if len(b) == 0 && params.omitempty {
				// FIXME:
				return nil
			}
			return nil
		}
		// This handles list types.
		subelements := []asn1.RawValue{}
		unmarshalParams := ""
		if params.tag == asn1.TagSet {
			unmarshalParams = "set"
		}
		rest, err := asn1.UnmarshalWithParams(encoded.FullBytes, &subelements, unmarshalParams)
		if err != nil {
			return err
		}
		if len(rest) > 0 {
			return errors.New("trailing data")
		}
		slice := reflect.MakeSlice(sliceType, len(subelements), len(subelements))
		for i, subel := range subelements {
			err = unmarshalValue(slice.Index(i), subel, params)
			if err != nil {
				return err
			}
		}
		v.Set(slice)
	}
	return nil
}

func unmarshalField(v reflect.Value, attrs map[string]Attribute, params fieldParameters) (err error) {
	attr, present := attrs[params.oid.String()]
	if !present || attr.IsEmpty() {
		if params.must {
			return fmt.Errorf("missing required attribute %s", params.oid)
		}
		// Otherwise, if the attribute is missing, there's nothing to do.
		return nil
	}
	k := v.Kind()
	if k == reflect.Slice && !params.list {
		l := attr.Len()
		values := reflect.MakeSlice(v.Type(), l, l)
		for i := 0; i < l; i++ {
			attrValue := attr.Get(i)
			if attrValue == nil {
				// This shouldn't happen.
				continue
			}
			err := unmarshalValue(values.Index(i), *attrValue, params)
			if err != nil {
				return err
			}
			v.Set(values)
		}
		return nil
	}
	encoded := attr.GetSingleValue()
	return unmarshalValue(v, *encoded, params)
}

func unmarshalStruct(v reflect.Value, attrs map[string]Attribute, params fieldParameters) (err error) {
	if v.Kind() != reflect.Struct {
		return errors.New("cannot unmarshal x.500 attributes to non-struct")
	}
	t := v.Type()
	n := t.NumField()
	if n == 0 {
		return nil
	}
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if !field.IsExported() {
			return errors.New("struct contains unexported fields")
		}
		p, err := parseFieldParameters(field.Tag.Get("x500"))
		if err != nil {
			return err
		}
		err = unmarshalField(v.Field(i), attrs, p)
		if err != nil {
			return err
		}
	}
	return nil
}

func UnmarshalWithParams(attrs []Attribute, val any, params string) error {
	p, err := parseFieldParameters(params)
	if err != nil {
		return err
	}
	v := reflect.ValueOf(val)
	if v.Kind() != reflect.Pointer || v.IsNil() {
		return &invalidUnmarshalError{reflect.TypeOf(val)}
	}
	attrMap := make(map[string]Attribute, len(attrs))
	for _, attr := range attrs {
		attrMap[attr.Type.String()] = attr
	}
	err = unmarshalStruct(v.Elem(), attrMap, p)
	if err != nil {
		return err
	}
	return nil
}
