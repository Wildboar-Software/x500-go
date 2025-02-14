package x500

import (
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/asn1"
	"errors"
	"fmt"
	"math/big"
	"reflect"
	"strconv"
	"strings"
	"time"
)

const tagDate int = 31
const tagTimeOfDay int = 32
const tagDuration = 34

var (
	bitStringType        = reflect.TypeFor[asn1.BitString]()
	objectIdentifierType = reflect.TypeFor[asn1.ObjectIdentifier]()
	enumeratedType       = reflect.TypeFor[asn1.Enumerated]()
	timeType             = reflect.TypeFor[time.Time]()
	bigIntType           = reflect.TypeFor[*big.Int]()
	durationType         = reflect.TypeFor[time.Duration]()
	rawContentsType      = reflect.TypeFor[asn1.RawContent]()
	dnType               = reflect.TypeFor[DistinguishedName]()
  nameAndUidType       = reflect.TypeFor[NameAndOptionalUID]()
	rdnType              = reflect.TypeFor[pkix.RelativeDistinguishedNameSET]()
	certType             = reflect.TypeFor[x509.Certificate]()
	crlType              = reflect.TypeFor[x509.RevocationList]()
	csrType              = reflect.TypeFor[x509.CertificateRequest]()
)

type fieldParameters struct {
	oid       asn1.ObjectIdentifier
	must      bool // If true, return error when unmarshalling if a required attribute is missing.
	tag       int
	list      bool   // If true, a slice field is treated as being a single value (such as []string for PostalAddress)
	uselang   bool   // If true, take the language context value from the lang tag passed in. Only used when marshalling.
	lang      string // Only used when marshalling.
	omitempty bool   // If 0 or false, do not produce an attribute value
}

func addLanguageContext(value asn1.RawValue, lang string) Attribute_valuesWithContext_Item {
	return Attribute_valuesWithContext_Item{
		Value: value,
		ContextList: []Context{
			{
				ContextType: Id_avc_language,
				ContextValues: []asn1.RawValue{
					{
						Class: asn1.ClassUniversal,
						Tag:   asn1.TagPrintableString,
						Bytes: []byte(lang),
					},
				},
			},
		},
	}
}

func stringToOID(str string) (ret asn1.ObjectIdentifier, err error) {
	var part string
	for len(str) > 0 {
		part, str, _ = strings.Cut(str, ".")
		i, err := strconv.ParseInt(part, 10, 32)
		if err != nil {
			return ret, err
		}
		ret = append(ret, int(i))
	}
	return ret, nil
}

func parseFieldParameters(str string) (ret fieldParameters, err error) {
	var part string
	for len(str) > 0 {
		part, str, _ = strings.Cut(str, ",")
		switch {
		case strings.HasPrefix(part, "oid:"):
			oid, err := stringToOID(part[4:])
			if err != nil {
				return ret, err
			}
			ret.oid = oid
		case strings.HasPrefix(part, "lang:"):
			if len(part) > 8 || len(part) < 7 {
				return ret, errors.New("invalid lang tag in x500 tags: must be two or three character ISO 639-2 code")
			}
			ret.lang = part[5:]
		case part == "must":
			ret.must = true
		case part == "printable":
			ret.tag = asn1.TagPrintableString
		case part == "ia5":
			ret.tag = asn1.TagIA5String
		case part == "num":
			ret.tag = asn1.TagNumericString
		case part == "utf8":
			ret.tag = asn1.TagUTF8String
		case part == "null":
			// TODO: implement
			ret.tag = asn1.TagNull
		case part == "date":
			ret.tag = tagDate
		case part == "tod":
			ret.tag = tagTimeOfDay
		case part == "set":
			ret.tag = asn1.TagSet
		case part == "uselang":
			ret.uselang = true
    case part == "omitempty":
      ret.omitempty = true
    case part == "list":
      ret.list = true
		}
	}
	return
}

func marshalValue(v reflect.Value, params fieldParameters) (ret asn1.RawValue, err error) {
	var bytes []byte

	if !v.IsValid() {
		// I don't know how or where this would happen, but we handle it here.
		return asn1.RawValue{}, nil
	}

	t := v.Type()
	k := t.Kind()
	if k == reflect.Pointer {
		if v.IsNil() {
			return asn1.RawValue{}, nil
		}
	}

  // This switch statement deals with specially-handled types.
	switch t {
	case enumeratedType:
		enumValue := v.Interface().(asn1.Enumerated)
    if enumValue == 0 && params.omitempty {
      return asn1.RawValue{}, nil
    }
		bytes, err = asn1.Marshal(enumValue)
		contentOctets := []byte{}
		if enumValue < 127 {
			contentOctets = []byte{byte(enumValue)}
		}
		return asn1.RawValue{
			Class:      asn1.ClassUniversal,
			Tag:        asn1.TagEnum,
			IsCompound: false,
			Bytes:      contentOctets,
			FullBytes:  bytes,
		}, err
	case durationType:
		durValue := v.Interface().(time.Duration)
		dv := fmt.Sprintf("P%s", strings.ToUpper(durValue.String()))
		return asn1.RawValue{
			Class: asn1.ClassUniversal,
			Tag:   tagDuration,
			Bytes: []byte(dv),
		}, nil
	case timeType:
		timeValue := v.Interface().(time.Time)
    if timeValue.IsZero() && params.omitempty {
      return asn1.RawValue{}, nil
    }
		if params.tag == tagDate { // DATE
			y, m, d := timeValue.Date()
			date := fmt.Sprintf("%04d-%02d-%02d", y, m, d)
			return asn1.RawValue{
				Class: asn1.ClassUniversal,
				Tag:   tagDate,
				Bytes: []byte(date),
			}, nil
		}
		if params.tag == tagTimeOfDay { // TIME-OF-DAY
			h, m, s := timeValue.Clock()
			tod := fmt.Sprintf("%02d:%02d:%02d", h, m, s)
			return asn1.RawValue{
				Class: asn1.ClassUniversal,
				Tag:   tagTimeOfDay,
				Bytes: []byte(tod),
			}, nil
		}
		bytes, err = asn1.MarshalWithParams(timeValue, "generalized")
		return asn1.RawValue{FullBytes: bytes}, err
	case bitStringType:
		bsValue := v.Interface().(asn1.BitString)
    if bsValue.BitLength == 0 && params.omitempty {
      return asn1.RawValue{}, nil
    }
		bytes, err = asn1.Marshal(bsValue)
		return asn1.RawValue{FullBytes: bytes}, err
	case objectIdentifierType:
		oidValue := v.Interface().(asn1.ObjectIdentifier)
		if len(oidValue) == 0 {
			// If must, we let the error happen
			return asn1.RawValue{}, err
		}
		bytes, err = asn1.Marshal(oidValue)
		return asn1.RawValue{FullBytes: bytes}, err
	case bigIntType:
		bigintValue := v.Interface().(*big.Int)
		if bigintValue == nil && !params.must {
			return asn1.RawValue{}, err
		}
    if bigintValue == big.NewInt(0) && params.omitempty {
      return asn1.RawValue{}, err
    }
		bytes, err = asn1.Marshal(bigintValue)
		return asn1.RawValue{FullBytes: bytes}, err
	case certType:
		cValue := v.Interface().(x509.Certificate)
		if len(cValue.Raw) == 0 {
			return asn1.RawValue{}, err
		}
		return asn1.RawValue{FullBytes: cValue.Raw}, err
	case crlType:
		cValue := v.Interface().(x509.RevocationList)
		if len(cValue.Raw) == 0 {
			return asn1.RawValue{}, err
		}
		return asn1.RawValue{FullBytes: cValue.Raw}, err
	case csrType:
		cValue := v.Interface().(x509.CertificateRequest)
		if len(cValue.Raw) == 0 {
			return asn1.RawValue{}, err
		}
		return asn1.RawValue{FullBytes: cValue.Raw}, err
	case dnType:
		dnValue := v.Interface().(DistinguishedName)
    if len(dnValue) == 0 && params.omitempty {
      return asn1.RawValue{}, err
    }
		fullBytes, err := asn1.Marshal(dnValue)
		if err != nil {
			return ret, err
		}
		return asn1.RawValue{FullBytes: fullBytes}, err
  case nameAndUidType:
    naouid := v.Interface().(NameAndOptionalUID)
    if len(naouid.Dn) == 0 && params.omitempty {
      return asn1.RawValue{}, err
    }
	}

	switch v.Kind() {
	case reflect.Bool:
		fallthrough
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
    if v.IsZero() && params.omitempty {
      return asn1.RawValue{}, nil
    }
		bytes, err = asn1.Marshal(v.Interface())
	case reflect.String:
		s := v.String()
		if len(s) == 0 {
			return asn1.RawValue{}, err
		}
		switch params.tag {
		case asn1.TagIA5String:
			bytes, err = asn1.MarshalWithParams(v.String(), "ia5")
		case asn1.TagNumericString:
			bytes, err = asn1.MarshalWithParams(v.String(), "numeric")
		case asn1.TagPrintableString:
			bytes, err = asn1.MarshalWithParams(v.String(), "printable")
		default:
			bytes, err = asn1.Marshal(v.String())
		}
	// I think this will be used if you find a way to treat a slice as a
	// "list" value like PostalAddress
	case reflect.Slice:
		sliceType := t
		if sliceType.Elem().Kind() == reflect.Uint8 {
      b := v.Bytes()
      if len(b) == 0 && params.omitempty {
        return asn1.RawValue{}, nil
      }
			return asn1.RawValue{
				Class: asn1.ClassUniversal,
				Tag:   asn1.TagOctetString,
				Bytes: b,
			}, nil
		}
		l := v.Len()
    if l == 0 {
      return asn1.RawValue{}, nil
    }
		m := make([]asn1.RawValue, l)
		for i := 0; i < l; i++ {
			m[i], err = marshalValue(v.Index(i), params)
			if err != nil {
				return ret, err
			}
		}
		if params.tag == asn1.TagSet {
			bytes, err = asn1.MarshalWithParams(m, "set")
      ret.Tag = asn1.TagSet
		} else {
			bytes, err = asn1.Marshal(m)
      ret.Tag = asn1.TagSequence
		}
    ret.IsCompound = true
	case reflect.Struct:
		n := t.NumField()
		if n > 0 && t.Field(0).Type == rawContentsType {
			bytes = v.Field(0).Bytes()
		} else {
			bytes, err = asn1.Marshal(v.Interface())
		}
    ret.IsCompound = true
	default:
		// Any other type is unhandled. Just let asn1 generate the error.
		bytes, err = asn1.Marshal(v)
	}
  ret.FullBytes = bytes
	return ret, err
}

func marshalField(v reflect.Value, params fieldParameters) (attr Attribute, err error) {
	if !v.IsValid() {
		return attr, fmt.Errorf("x500: cannot marshal nil value")
	}
	k := v.Kind()
	// If the field is an interface{} then recurse into it.
	if k == reflect.Interface && v.Type().NumMethod() == 0 {
		return marshalField(v.Elem(), params)
	}
	t := v.Type()
	attr.Type = params.oid
	if k == reflect.Slice && !params.list && t.Elem().Kind() != reflect.Uint8 && t != objectIdentifierType && t != dnType && t != rdnType {
		l := v.Len()
		values := make([]asn1.RawValue, 0, l)
		for i := 0; i < l; i++ {
			innerv, err := marshalValue(v.Index(i), params)
			if err != nil {
				return attr, err
			}
			if reflect.ValueOf(innerv).IsZero() {
				// Returning asn1.RawValue{} means there was no error:
				// the value just shouldn't be encoded.
				continue
			}
			values = append(values, innerv)
		}
		attr.Values = values
		if params.uselang && len(params.lang) == 2 {
			attr.ValuesWithContext = make([]Attribute_valuesWithContext_Item, len(attr.Values))
			for i, plainValue := range attr.Values {
				attr.ValuesWithContext[i] = addLanguageContext(plainValue, params.lang)
			}
			attr.Values = make([]asn1.RawValue, 0)
		}
		return attr, nil
	}
	value, err := marshalValue(v, params)
	if err != nil {
		return attr, err
	}
	if reflect.ValueOf(value).IsZero() {
		// Returning an empty attribute means "nothing to encode"
		return attr, nil
	}
	if params.uselang && len(params.lang) == 2 {
		attr.ValuesWithContext = []Attribute_valuesWithContext_Item{addLanguageContext(value, params.lang)}
	} else {
		attr.Values = []asn1.RawValue{value}
	}
	return attr, nil
}

func MarshalWithParams(val any, lang string) (attrs []Attribute, err error) {
	v := reflect.ValueOf(val)
	if v.Kind() != reflect.Struct {
		return attrs, errors.New("cannot marshal non-struct to directory attributes")
	}
	t := v.Type()
	n := t.NumField()
	if n == 0 {
		return attrs, nil
	}
	for i := 0; i < n; i++ {
		if !t.Field(i).IsExported() {
			return nil, errors.New("struct contains unexported fields")
		}
	}

	attrs = make([]Attribute, 0, n+1)

	for i := 0; i < n; i++ {
		tag := t.Field(i).Tag.Get("x500")
		if len(tag) == 0 {
			// We simply don't marshal non-X.500 directory values.
			continue
		}
		fieldParams, err := parseFieldParameters(tag)
		if fieldParams.lang == "" {
			fieldParams.lang = lang
		}
		if err != nil {
			return nil, err
		}
		attr, err := marshalField(v.Field(i), fieldParams)
		if err != nil {
			return nil, err
		}
		if attr.Len() == 0 {
			// No values means there was no error:
			// this attribute just isn't meant to be encoded.
			continue
		}
		attrs = append(attrs, attr)
	}
	return
}

// TODO: Encoding for pkix.DistinguishedName
// TODO: Encoding for pkix.RelativeDistinguishedName

// Might already be done.
// TODO: Encoding for pkix.Extension
// TODO: Encoding for pkix.AlgorithmIdentifier
// TODO: Encoding for NameAndOptionalUID
// TODO: Fax

// TODO: Encoding for uint types
// TODO: Schema elements
// TODO: Guide
// TODO: Substring Assertion
// TODO: EnhancedGuide
// TODO: SubtreeSpecification
// TODO: WGS84Coordinates
// TODO: WGS84Position
// TODO: WGS84Line
// TODO: WGS84Polygon
// TODO: ComplexNumber
// TODO: Currency
// TODO: TextualKeyValue
// TODO: CIDR
// TODO: HostPort
// TODO: TransportAddress
// TODO: NameAndString
// TODO: Crypto
// TODO: Does this work with just plain asn1.RawValue?
