package x500

import (
	"bytes"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/asn1"
	"math/big"
	"reflect"
	"testing"
	"time"
)

func TestMarshal(t *testing.T) {
	type Person struct {
		CommonName []string `x500:"oid:2.5.4.3"`
		Surname    []string `x500:"oid:2.5.4.4"`
	}
	p := Person{
		CommonName: []string{"Spongebob"},
		Surname:    []string{"Squarepants"},
	}
	attrs, err := MarshalWithParams(p, "")
	if err != nil {
		t.Error(err)
		return
	}
	if len(attrs) != 2 {
		t.Errorf("Marshing produced %d attributes instead of %d", len(attrs), 2)
		return
	}
	if !attrs[0].Type.Equal(Id_at_commonName) {
		t.Errorf("First attribute was not commonName: %+v", attrs[0].Type)
		return
	}
	if !attrs[1].Type.Equal(Id_at_surname) {
		t.Errorf("Second attribute was not surname: %+v", attrs[1].Type)
		return
	}
	if len(attrs[0].Values) != 1 {
		t.Errorf("commonName contained %d values", len(attrs[0].Values))
		return
	}
	if len(attrs[1].Values) != 1 {
		t.Errorf("surname contained %d values", len(attrs[1].Values))
		return
	}
	if len(attrs[0].ValuesWithContext) > 0 {
		t.Errorf("commonName contained %d values with context", len(attrs[0].ValuesWithContext))
		return
	}
	if len(attrs[1].ValuesWithContext) > 0 {
		t.Errorf("surname contained %d values with context", len(attrs[1].ValuesWithContext))
		return
	}
	cn0 := attrs[0].Values[0]
	sn0 := attrs[1].Values[0]
	var cn string
	var sn string
	rest, err := asn1.Unmarshal(cn0.FullBytes, &cn)
	if err != nil {
		t.Error(err)
		return
	}
	if len(rest) > 0 {
		t.Error("trailing data")
		return
	}
	rest, err = asn1.Unmarshal(sn0.FullBytes, &sn)
	if err != nil {
		t.Error(err)
		return
	}
	if len(rest) > 0 {
		t.Error("trailing data")
		return
	}
	if cn != "Spongebob" {
		t.Errorf("commonName did not have the right value: %s", cn)
		return
	}
	if sn != "Squarepants" {
		t.Errorf("surname did not have the right value: %s", sn)
		return
	}
}

func TestMarshalSingleValued(t *testing.T) {
	type Person struct {
		CommonName string `x500:"oid:2.5.4.3"`
		Surname    string `x500:"oid:2.5.4.4"`
	}
	p := Person{
		CommonName: "Spongebob",
		Surname:    "Squarepants",
	}
	attrs, err := MarshalWithParams(p, "")
	if err != nil {
		t.Error(err)
		return
	}
	if len(attrs) != 2 {
		t.Errorf("Marshing produced %d attributes instead of %d", len(attrs), 2)
		return
	}
	if !attrs[0].Type.Equal(Id_at_commonName) {
		t.Errorf("First attribute was not commonName: %+v", attrs[0].Type)
		return
	}
	if !attrs[1].Type.Equal(Id_at_surname) {
		t.Errorf("Second attribute was not surname: %+v", attrs[1].Type)
		return
	}
	if len(attrs[0].Values) != 1 {
		t.Errorf("commonName contained %d values", len(attrs[0].Values))
		return
	}
	if len(attrs[1].Values) != 1 {
		t.Errorf("surname contained %d values", len(attrs[1].Values))
		return
	}
	if len(attrs[0].ValuesWithContext) > 0 {
		t.Errorf("commonName contained %d values with context", len(attrs[0].ValuesWithContext))
		return
	}
	if len(attrs[1].ValuesWithContext) > 0 {
		t.Errorf("surname contained %d values with context", len(attrs[1].ValuesWithContext))
		return
	}
	cn0 := attrs[0].Values[0]
	sn0 := attrs[1].Values[0]
	var cn string
	var sn string
	rest, err := asn1.Unmarshal(cn0.FullBytes, &cn)
	if err != nil {
		t.Error(err)
		return
	}
	if len(rest) > 0 {
		t.Error("trailing data")
		return
	}
	rest, err = asn1.Unmarshal(sn0.FullBytes, &sn)
	if err != nil {
		t.Error(err)
		return
	}
	if len(rest) > 0 {
		t.Error("trailing data")
		return
	}
	if cn != "Spongebob" {
		t.Errorf("commonName did not have the right value: %s", cn)
		return
	}
	if sn != "Squarepants" {
		t.Errorf("surname did not have the right value: %s", sn)
		return
	}
}

func TestMarshalWeirdValues(t *testing.T) {
	type WeirdType struct {
		SomeInt     int                               `x500:"oid:1.2.3.4"`
		SomeBytes   []byte                            `x500:"oid:1.2.3.5"`
		SomeBool    bool                              `x500:"oid:1.2.3.6"`
		SomeOID     asn1.ObjectIdentifier             `x500:"oid:1.2.3.7"`
		MoreOIDs    []asn1.ObjectIdentifier           `x500:"oid:1.2.3.8"`
		SomeEnum    asn1.Enumerated                   `x500:"oid:1.2.3.9"`
		SomeBigInt  *big.Int                          `x500:"oid:1.2.3.10"`
		SomeBits    asn1.BitString                    `x500:"oid:1.2.3.11"`
		MoreEnums   []asn1.Enumerated                 `x500:"oid:1.2.3.12"`
		SomeTime    time.Time                         `x500:"oid:1.2.3.13"`
		SomeName    NameAndOptionalUID                `x500:"oid:1.2.3.14"`
		SomeDN      DistinguishedName                 `x500:"oid:1.2.3.15"`
		SomeRDN     pkix.RelativeDistinguishedNameSET `x500:"oid:1.2.3.16,set"`
		Cert        x509.Certificate                  `x500:"oid:1.2.3.17"`
		SomeCRL     x509.RevocationList               `x500:"oid:1.2.3.18"`
		CertReq     x509.CertificateRequest           `x500:"oid:1.2.3.19"`
		NonX500Type int
	}
	moreOIDs := []asn1.ObjectIdentifier{
		{1, 3, 4, 6, 1, 56940},
		{1, 3, 4, 6, 1, 56941},
		{1, 3, 4, 6, 1, 56942},
	}
	bits := asn1.BitString{
		Bytes:     []byte{0xFF, 0xF0},
		BitLength: 12,
	}
	nowTime := time.Now()
	v := WeirdType{
		SomeInt:    5,
		SomeBytes:  []byte{1, 3, 5},
		SomeBool:   true,
		SomeOID:    asn1.ObjectIdentifier{1, 3, 4, 6, 1, 56940},
		MoreOIDs:   moreOIDs,
		SomeEnum:   asn1.Enumerated(5),
		SomeBigInt: big.NewInt(295057002850720),
		SomeBits:   bits,
		MoreEnums:  []asn1.Enumerated{1, 5, 8},
		SomeTime:   nowTime,
		SomeRDN: []pkix.AttributeTypeAndValue{
			{
				Type: Id_at_commonName,
				Value: asn1.RawValue{
					Class: asn1.ClassUniversal,
					Tag:   asn1.TagUTF8String,
					Bytes: []byte("hi mom"),
				},
			},
		},
	}
	attrs, err := MarshalWithParams(v, "")
	if err != nil {
		t.Error(err)
		return
	}
	if len(attrs) != 16 {
		t.Errorf("number of attributes: %d", len(attrs))
		return
	}
	correctOIDS := []asn1.ObjectIdentifier{
		{1, 2, 3, 4},
		{1, 2, 3, 5},
		{1, 2, 3, 6},
		{1, 2, 3, 7},
		{1, 2, 3, 8},
		{1, 2, 3, 9},
		{1, 2, 3, 10},
		{1, 2, 3, 11},
		{1, 2, 3, 12},
		{1, 2, 3, 13},
		{1, 2, 3, 14},
		{1, 2, 3, 15},
		{1, 2, 3, 16},
		{1, 2, 3, 17},
		{1, 2, 3, 18},
		{1, 2, 3, 19},
	}
	for i, attr := range attrs {
		if !attr.Type.Equal(correctOIDS[i]) {
			t.Errorf("wrong oid in attribute: %s", attr.Type)
			return
		}
		if i == 4 {
			if attr.Len() != 3 {
				t.Errorf("incorrect number of attribute values (%d) for type %s", attr.Len(), attr.Type)
				return
			}
		} else if i == 8 {
			if attr.Len() != 3 {
				t.Errorf("incorrect number of attribute values (%d) for type %s", attr.Len(), attr.Type)
				return
			}
		} else if attr.Len() != 1 {
			t.Errorf("incorrect number of attribute values (%d) for type %s", attr.Len(), attr.Type)
			return
		}
	}

	someIntEncoded := attrs[0].GetSingleValue()
	someBytesEncoded := attrs[1].GetSingleValue()
	someBoolEncoded := attrs[2].GetSingleValue()
	someOIDEncoded := attrs[3].GetSingleValue()
	// moreOIDsEncoded 4
	someEnumEncoded := attrs[5].GetSingleValue()
	someBigIntEncoded := attrs[6].GetSingleValue()
	someBitsEncoded := attrs[7].GetSingleValue()
	someTimeEncoded := attrs[9].GetSingleValue()
	nameEncoded := attrs[10].GetSingleValue()
	dnEncoded := attrs[11].GetSingleValue()
	rdnEncoded := attrs[12].GetSingleValue()
	certEncoded := attrs[13].GetSingleValue()
	crlEncoded := attrs[14].GetSingleValue()
	certReqEncoded := attrs[15].GetSingleValue()

	if !reflect.DeepEqual(someIntEncoded.FullBytes, []byte{2, 1, 5}) {
		t.Error("int encoded incorrectly")
		return
	}
	if !reflect.DeepEqual(someBytesEncoded.Bytes, []byte{1, 3, 5}) {
		t.Error("bytes encoded incorrectly")
		return
	}
	if !reflect.DeepEqual(someEnumEncoded.FullBytes, []byte{10, 1, 5}) {
		t.Error("enum encoded incorrectly")
		return
	}
	if !reflect.DeepEqual(someBoolEncoded.FullBytes, []byte{1, 1, 0xFF}) {
		t.Error("bool encoded incorrectly")
		return
	}
	someOID := asn1.ObjectIdentifier{}
	rest, err := asn1.Unmarshal(someOIDEncoded.FullBytes, &someOID)
	if err != nil {
		t.Error(err)
		return
	}
	if len(rest) > 0 {
		t.Error("trailing data")
		return
	}
	if !someOID.Equal(asn1.ObjectIdentifier{1, 3, 4, 6, 1, 56940}) {
		t.Error("oid encoded incorrectly")
		return
	}
	someBigint := big.NewInt(1)
	rest, err = asn1.Unmarshal(someBigIntEncoded.FullBytes, &someBigint)
	if err != nil {
		t.Error(err)
		return
	}
	if len(rest) > 0 {
		t.Error("trailing data")
		return
	}
	if someBigint.Int64() != 295057002850720 {
		t.Error("bigint encoded incorrectly")
		return
	}

	for i, encodedOID := range attrs[4].Values {
		var oid asn1.ObjectIdentifier
		rest, err = asn1.Unmarshal(encodedOID.FullBytes, &oid)
		if err != nil {
			t.Error(err)
			return
		}
		if len(rest) > 0 {
			t.Error("trailing data")
			return
		}
		if !oid.Equal(moreOIDs[i]) {
			t.Errorf("multi-valued OIDs did not match: %s != %s", oid, moreOIDs[i])
			return
		}
	}

	var bitsDecoded asn1.BitString
	rest, err = asn1.Unmarshal(someBitsEncoded.FullBytes, &bitsDecoded)
	if err != nil {
		t.Error(err)
		return
	}
	if len(rest) > 0 {
		t.Error("trailing data")
		return
	}
	if !bytes.Equal(bitsDecoded.Bytes, bits.Bytes) || bitsDecoded.BitLength != bits.BitLength {
		t.Error("bits encoded incorrectly")
		return
	}

	expectedEnums := []byte{1, 5, 8}
	for i, v := range attrs[8].Values {
		if v.Tag != asn1.TagEnum || v.FullBytes[2] != expectedEnums[i] {
			t.Errorf("enums encoded incorrectly: %d != %d (index %d)", v.FullBytes[2], expectedEnums[i], i)
			return
		}
	}

	var timeDecoded time.Time
	rest, err = asn1.Unmarshal(someTimeEncoded.FullBytes, &timeDecoded)
	if err != nil {
		t.Error(err)
		return
	}
	if len(rest) > 0 {
		t.Error("trailing data")
		return
	}
	if timeDecoded.String()[:16] != nowTime.String()[:16] {
		t.Error("time encoded incorrectly")
		return
	}

	var nameDecoded NameAndOptionalUID
	rest, err = asn1.Unmarshal(nameEncoded.FullBytes, &nameDecoded)
	if err != nil {
		t.Error(err)
		return
	}
	if len(rest) > 0 {
		t.Error("trailing data")
		return
	}

	// TODO: Compare

	var dnDecoded DistinguishedName
	rest, err = asn1.Unmarshal(dnEncoded.FullBytes, &dnDecoded)
	if err != nil {
		t.Error(err)
		return
	}
	if len(rest) > 0 {
		t.Error("trailing data")
		return
	}

	// TODO: Compare

	var rdnDecoded pkix.RelativeDistinguishedNameSET
	rest, err = asn1.UnmarshalWithParams(rdnEncoded.FullBytes, &rdnDecoded, "set")
	if err != nil {
		t.Error(err)
		return
	}
	if len(rest) > 0 {
		t.Error("trailing data")
		return
	}

	// TODO: Compare

	_, err = x509.ParseCertificate(certEncoded.FullBytes)
	if err != nil {
		t.Error(err)
		return
	}
	// TODO: Compare

	_, err = x509.ParseRevocationList(crlEncoded.FullBytes)
	if err != nil {
		t.Error(err)
		return
	}
	// TODO: Compare

	_, err = x509.ParseCertificateRequest(certReqEncoded.FullBytes)
	if err != nil {
		t.Error(err)
		return
	}
	// TODO: Compare

}

// TODO: Test zero values
// TODO: Test encoding with language contexts
// TODO: Test encoding list values
// TODO: Test encoding different string types
// TODO: Make RDN always encode and decode as a SET, despite tag or not
