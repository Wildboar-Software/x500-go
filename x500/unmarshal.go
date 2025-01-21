package x500

import (
	"encoding/asn1"
	"errors"
	"fmt"
	"reflect"
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

func unmarshalValue(v reflect.Value, encoded asn1.RawValue, params fieldParameters) (err error) {
	k := v.Kind()
	switch k {
	case reflect.Bool:
		fallthrough
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		fallthrough
	case reflect.String:
		rest, err := asn1.Unmarshal(encoded.FullBytes, v.Addr().Interface())
		if err != nil {
			return err
		}
		if len(rest) > 0 {
			return errors.New("trailing data")
		}
	}
	return nil
}

func unmarshalField(v reflect.Value, attrs map[string]Attribute, params fieldParameters) (err error) {
	attr, present := attrs[params.oid.String()]
	if !present || attr.IsEmpty() {
		if params.must {
			return errors.New(fmt.Sprintf("missing required attribute %s", params.oid))
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
