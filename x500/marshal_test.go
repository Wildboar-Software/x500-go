package x500

import (
	"bytes"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/asn1"
	"encoding/pem"
	"math/big"
	"reflect"
	"testing"
	"time"
)

const keyPem string = `-----BEGIN PRIVATE KEY-----
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
yW0tzYNu6jXxgIzaaQ2T/ODhZR2iVd5B
-----END PRIVATE KEY-----
`

const certPem string = `-----BEGIN CERTIFICATE-----
MIIFfzCCA2egAwIBAgIUGmlo7xlQ5Bd08eBgD+MD28XzVMowDQYJKoZIhvcNAQEL
BQAwTzELMAkGA1UEBhMCVVMxCzAJBgNVBAgMAkZMMQ4wDAYDVQQHDAVUYW1wYTER
MA8GA1UECgwIV2lsZGJvYXIxEDAOBgNVBAMMB21lZXJrYXQwHhcNMjIwNzI1MjAy
NDM5WhcNMjQwNzI0MjAyNDM5WjBPMQswCQYDVQQGEwJVUzELMAkGA1UECAwCRkwx
DjAMBgNVBAcMBVRhbXBhMREwDwYDVQQKDAhXaWxkYm9hcjEQMA4GA1UEAwwHbWVl
cmthdDCCAiIwDQYJKoZIhvcNAQEBBQADggIPADCCAgoCggIBAPTK0uLlvhAg8e+6
5o0utIjhmTv9sYlbtOu0FRPBgY81Qfb0JTPmotM8/5inK7cwoUxqiHXO2oc/4i6d
U7KGHt1K66o/5xtAVjxTGuLI9WiMCT9485s7okxSk9tCpvwZ6i7y7cqxRe06QdDX
VXiE1f3EdkcV6gbVOfJcFJKatdhCKGPKsxMKtBCTYyJ2Cft7Cy953ceV5rOMYcta
VqlXr2dOyQYvwsGnZFPpe4mtyLR4wzWb84WIkGGmwFVaFUxWHn9B04l9G1eujGtQ
PT+UeCpEFs1wjaMAufcDfxrocBJ/VJG/vBOmNCUelnh8B5DywtXD3RJmi7fqJK/z
T8KgUTr6Ed8wLvnFfWi7iwsryNQv8CEeZeDbLIeDnEPc5XEXXXpcJm7Snf4nCfK/
ofP2spCP0PI+60DzQac1+GMVmmY5GFX405DBQAwDL5dcJ1RZCVEKFdGuaDYCItvI
7F9KRScbjUCbqJOD5B/GnmUoH0Xwxy14+LahMWCCKasA+9Y6QuUJ9fxMer3DhZ+U
oI2Gk1LVZ99Wx2xer1n4/4pIYL0ZYcrBm57lNdJ6eVr8T8ZMgiSAcdkY5iCuu/7X
e2zc1Km7RNSvJ6skF0YdYISCeImzecR8gNF8QU0h5ginaW2XPdMQsqOwwQaRyGcE
NDAZCVuSwn3pNbd0gEpR7UG4eq4RAgMBAAGjUzBRMB0GA1UdDgQWBBRyPT6EMjCl
cf4nE/u535EUzM3H1jAfBgNVHSMEGDAWgBRyPT6EMjClcf4nE/u535EUzM3H1jAP
BgNVHRMBAf8EBTADAQH/MA0GCSqGSIb3DQEBCwUAA4ICAQDb+icQ06S74Bpc0n8t
CXub0kNlENPdcVs/g1+nzbD9UgEi2PXuC0u9GLFuoFbe3i5Nn8Z4VlFo6/6rbvK2
gVc7MD/Nfo4/Ohe/Tv8YPw8klGi4f1h0O4/beqEayyfgXI/pIPUHqGlbD9xp27D9
f3stIjc7aGcoqX8N5tmamrKRM+/6p8IVKqqw+iVMXeLgTCZ/l1UkneJsRfz69m3z
YSDxPWn0LaIdFhvUZizkhiDJcPeZ05dlNfhscDoJyz5G8w/c7aGFwD/qYtjelBFP
hWV+9pfdSed6bnkStN6ztNPluJ7Vm2Mkw0aootc8HFxIfFPhgOElvv6XBMvZ8MzJ
xtlo6sE8BeCusPX4rU63ZeoadE3NFUyI1BRFyXUg896dXhAVvvm45FWkBJdUcDaL
F5bDAY41aZlmNi/nVtzwgk/ph6yLSp2NVkvT2wKPaimVcdsaXhlVtrRMLne6QiDm
1Tql5+gTmQfc2APqNnL3Hp+xsm5DgccOR0hJucf7BWL3evu9gCijGJ5cnGTHviWY
kczD5TXI9czgiEgEWID6qgCic3ZfNOXyhb5ZFrkVacWX2X76KVYffkW8NCFMkCjh
nePOV/Y2kXvUoFNpCjcpaG4FgbQyMxC45jRq9znmCBa3y44BWKs+M2+sKk/z1hWT
JDKTi1FovCNYrJ8YeKqWM+TMQA==
-----END CERTIFICATE-----
`

const csrPem = `-----BEGIN CERTIFICATE REQUEST-----
MIICiTCCAXECAQAwRDELMAkGA1UEBhMCVVMxCzAJBgNVBAgMAkZMMRUwEwYDVQQH
DAxKYWNrc29udmlsbGUxETAPBgNVBAoMCFdpbGRib2FyMIIBIjANBgkqhkiG9w0B
AQEFAAOCAQ8AMIIBCgKCAQEA4HWy2WvFGEOy06yFprR/bQMx0rOqfQEndc5ez0B2
OMkWUPKKT6CSD7A5Vse9JvKWcVBW56yYKd4XMv1SAGZe/68UuY6DsRzzPgFGT/yr
pYSWmaz86NGKU2sTaFSYc2tgvAzBozJjpRGtg1i3DdnrpsvjriBn2eklLugzJ7hN
7165HLnZ5Anapf2md64CwzNQDMHoN9qSH0qnHFblWl/KUGvcUBzdga4AenGLvmBP
MEG4Xtw4S0IPZH1HQnaSf8rpSLRrSXjMZ8YjORFNfuMkROgb1FmgO1qal4AHMO//
mAVyyZpj4DiXL/F/P+RvopbWYYQtjqAao8zTw32U9CEVIQIDAQABoAAwDQYJKoZI
hvcNAQELBQADggEBALj9PrFyIwyvE/AGsq4tNmhz3hMCBy7Rss/R3T46DM8kaeml
/yKFpJfcQmUIuyPwS0SYJ77QNQjwk9xbhjxJzIV5cGfQ9Ci5IkC7fPFxxzeV/miI
dPH45hQrNFF6BHAV2vdKbvYUXqiZvM2cryEPAeBMmjORJ9qtJllIAmngcVhBmFaV
OhhJbU+cITbUyStTaoqNnG+W31WfYyV6aRztF1DlwyFtAaWFKg1gZUoF999VSXv5
XWOSIVog/fBUnBRcJ3FJgHfbZDeJfvKp1b8AP+k4SHXAnOqSYkOMugdy6fod0ufq
GGTLwSLIpAE/OC0ScV8e+Kkhi92hqCQEePyz8/o=
-----END CERTIFICATE REQUEST-----
`

const crlPem = `-----BEGIN X509 CRL-----
MIIDFDCCAfwCAQEwDQYJKoZIhvcNAQEFBQAwXzEjMCEGA1UEChMaU2FtcGxlIFNp
Z25lciBPcmdhbml6YXRpb24xGzAZBgNVBAsTElNhbXBsZSBTaWduZXIgVW5pdDEb
MBkGA1UEAxMSU2FtcGxlIFNpZ25lciBDZXJ0Fw0xMzAyMTgxMDMyMDBaFw0xMzAy
MTgxMDQyMDBaMIIBNjA8AgMUeUcXDTEzMDIxODEwMjIxMlowJjAKBgNVHRUEAwoB
AzAYBgNVHRgEERgPMjAxMzAyMTgxMDIyMDBaMDwCAxR5SBcNMTMwMjE4MTAyMjIy
WjAmMAoGA1UdFQQDCgEGMBgGA1UdGAQRGA8yMDEzMDIxODEwMjIwMFowPAIDFHlJ
Fw0xMzAyMTgxMDIyMzJaMCYwCgYDVR0VBAMKAQQwGAYDVR0YBBEYDzIwMTMwMjE4
MTAyMjAwWjA8AgMUeUoXDTEzMDIxODEwMjI0MlowJjAKBgNVHRUEAwoBATAYBgNV
HRgEERgPMjAxMzAyMTgxMDIyMDBaMDwCAxR5SxcNMTMwMjE4MTAyMjUxWjAmMAoG
A1UdFQQDCgEFMBgGA1UdGAQRGA8yMDEzMDIxODEwMjIwMFqgLzAtMB8GA1UdIwQY
MBaAFL4SAcyq6hGA2i6tsurHtfuf+a00MAoGA1UdFAQDAgEDMA0GCSqGSIb3DQEB
BQUAA4IBAQBCIb6B8cN5dmZbziETimiotDy+FsOvS93LeDWSkNjXTG/+bGgnrm3a
QpgB7heT8L2o7s2QtjX2DaTOSYL3nZ/Ibn/R8S0g+EbNQxdk5/la6CERxiRp+E2T
UG8LDb14YVMhRGKvCguSIyUG0MwGW6waqVtd6K71u7vhIU/Tidf6ZSdsTMhpPPFu
PUid4j29U3q10SGFF6cCt1DzjvUcCwHGhHA02Men70EgZFADPLWmLg0HglKUh1iZ
WcBGtev/8VsUijyjsM072C6Ut5TwNyrrthb952+eKlmxLNgT0o5hVYxjXhtwLQsL
7QZhrypAM1DLYqQjkiDI7hlvt7QuDGTJ
-----END X509 CRL-----`

func getCert() x509.Certificate {
	block, _ := pem.Decode([]byte(certPem))
	if block == nil || block.Type != "CERTIFICATE" {
		panic("not certificate")
	}
	cert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		panic("invalid test certificate")
	}
	return *cert
}

func getCSR() x509.CertificateRequest {
	block, _ := pem.Decode([]byte(csrPem))
	if block == nil || block.Type != "CERTIFICATE REQUEST" {
		panic("not csr")
	}
	cert, err := x509.ParseCertificateRequest(block.Bytes)
	if err != nil {
		panic("invalid test csr")
	}
	return *cert
}

func getCRL() x509.RevocationList {
	block, _ := pem.Decode([]byte(crlPem))
	if block == nil || block.Type != "X509 CRL" {
		panic("not crl")
	}
	crl, err := x509.ParseRevocationList(block.Bytes)
	if err != nil {
		panic("invalid test crl")
	}
	return *crl
}

func getDN() DistinguishedName {
	return DistinguishedName{
		pkix.RelativeDistinguishedNameSET{
			{
				Type: Id_at_countryName,
				Value: asn1.RawValue{
					Tag:        asn1.TagPrintableString,
					Class:      asn1.ClassUniversal,
					IsCompound: false,
					Bytes:      []byte("US"),
				},
			},
		},
	}
}

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
	cert := getCert()
	csr := getCSR()
	crl := getCRL()
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
		SomeName: NameAndOptionalUID{
			Dn:  getDN(),
			Uid: asn1.BitString{Bytes: []byte{5}, BitLength: 8},
		},
		SomeCRL:     crl,
		Cert:        cert,
		CertReq:     csr,
		SomeDN:      getDN(),
		NonX500Type: 5,
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
	// moreEnums
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
	if len(nameDecoded.Dn) != 1 {
		t.Error("Invalid NameAndOptionalUID")
		return
	}

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
	if len(dnDecoded) != 1 {
		t.Error("Invalid DistinguishedName")
		return
	}

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
	if len(rdnDecoded) != 1 {
		t.Error("Invalid RelativeDistinguishedName")
		return
	}
	_, err = x509.ParseCertificate(certEncoded.FullBytes)
	if err != nil {
		t.Error(err)
		return
	}
	_, err = x509.ParseRevocationList(crlEncoded.FullBytes)
	if err != nil {
		t.Error(err)
		return
	}
	_, err = x509.ParseCertificateRequest(certReqEncoded.FullBytes)
	if err != nil {
		t.Error(err)
		return
	}
}

// TODO: Test zero values
// TODO: Test encoding with language contexts
// TODO: Test encoding list values
// TODO: Test encoding different string types
// TODO: Make RDN always encode and decode as a SET, despite tag or not
