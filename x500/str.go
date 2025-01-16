// X.500 Directory Data Structures with ASN.1 Encoding and Decoding
package x500

import (
	"encoding/asn1"
	"errors"
	"fmt"

	"github.com/Wildboar-Software/x500-go/teletex"
)

func ASN1RawValueToStr(value asn1.RawValue) (output string, err error) {
	if value.Class != asn1.ClassUniversal {
		// TODO: Something better than this.
		return fmt.Sprintf("% x", value.FullBytes), nil
	}
	switch value.Tag {
	case asn1.TagBoolean:
		{
			if len(value.Bytes) == 0 {
				return "ERROR", errors.New("BOOLEAN encoded on zero bytes")
			}
			if value.Bytes[0] > 0 {
				return "TRUE", nil
			} else {
				return "FALSE", nil
			}
		}
	case asn1.TagInteger:
		fallthrough
	case asn1.TagEnum:
		{
			var decoded int = 0
			rest, err := asn1.Unmarshal(value.FullBytes, &decoded)
			if err != nil {
				return "?", err
			}
			if len(rest) > 0 {
				return "?", err
			}
			return fmt.Sprint(decoded), nil
		}
	case asn1.TagOctetString:
		{
			return fmt.Sprintf("% x", value.Bytes), nil
		}
	case asn1.TagNull:
		{
			return "NULL", nil
		}
	case asn1.TagOID:
		{
			oid := asn1.ObjectIdentifier{}
			rest, err := asn1.Unmarshal(value.FullBytes, &oid)
			if err != nil {
				return "?", err
			}
			if len(rest) > 0 {
				return "?", err
			}
			return oid.String(), nil
		}
	case asn1.TagUTF8String:
		fallthrough
	case asn1.TagNumericString:
		fallthrough
	case asn1.TagPrintableString:
		fallthrough
	case asn1.TagT61String:
		fallthrough
	case asn1.TagIA5String:
		fallthrough
	case asn1.TagGeneralString:
		fallthrough
	case asn1.TagGeneralizedTime:
		fallthrough
	case asn1.TagUTCTime:
		{
			return string(value.Bytes[:]), nil
		}
	default:
		{
			return fmt.Sprintf("% x", value.FullBytes), nil
		}
	}
}

func NewDirectoryString(s string) asn1.RawValue {
	return asn1.RawValue{
		Class:      asn1.ClassUniversal,
		Tag:        asn1.TagUTF8String,
		IsCompound: false,
		Bytes:      []byte(s),
	}
}

func universalStringFromBytes(bytes []byte) (s string, err error) {
	if len(bytes)%4 != 0 {
		return "", errors.New("invalid universalstring length")
	}
	if len(bytes) == 0 {
		return "", nil
	}
	runes := make([]rune, len(bytes)/4)

	l := len(bytes)
	hasNullTerminator := bytes[l-1] == 0 && bytes[l-2] == 0 && bytes[l-3] == 0 && bytes[l-4] == 0
	if hasNullTerminator {
		bytes = bytes[:l-4] // strip it
	}

	for len(bytes) > 0 {
		nextRune := rune(bytes[0])<<24 + rune(bytes[1])<<16 + rune(bytes[2])<<8 + rune(bytes[3])
		runes = append(runes, nextRune)
		bytes = bytes[4:]
	}

	return string(runes), nil
}

func DirectoryStringToString(ds asn1.RawValue) (s string, err error) {
	if ds.Class != asn1.ClassUniversal {
		return "", errors.New("non-universal tag directory string")
	}
	if ds.Tag == 28 {
		return universalStringFromBytes(ds.Bytes)
	}
	if len(ds.FullBytes) > 0 {
		rest, err := asn1.Unmarshal(ds.FullBytes, &s)
		if err != nil {
			return "", err
		}
		if len(rest) > 0 {
			return s, errors.New("trailing bytes after directorystring")
		}
	}
	if len(ds.Bytes) == 0 {
		return "", nil
	}
	// It is an edge case and a bug for an asn1.RawValue to have .Bytes
	// populated, but not .FullBytes, but we still attempt to handle it.
	switch ds.Tag {
	case asn1.TagUTF8String:
		fallthrough
	case asn1.TagPrintableString:
		return string(ds.Bytes), nil
	case asn1.TagT61String:
		return teletex.TeletexToUTF8(ds.Bytes), nil
	case asn1.TagBMPString:
		// Why, yes, I would LOVE to marshal a value just to immediately
		// unmarshal it to a string, because I do not want to re-implement the
		// same exact BMPString decoding the Go standard library already has
		// for this silly little edge case where FullBytes is not populated and
		// the string happens to be a BMPString. Blame Golang's ASN.1 library.
		encodedbmp, err := asn1.Marshal(ds)
		if err != nil {
			return "", err
		}
		rest, err := asn1.Unmarshal(encodedbmp, &s)
		if err != nil {
			return "", err
		}
		if len(rest) > 0 {
			return s, errors.New("trailing bytes after bmpstring")
		}
		return s, nil
	default:
		return "", errors.New(fmt.Sprintf("invalid tag for directorystring: %d", ds.Tag))
	}
}
