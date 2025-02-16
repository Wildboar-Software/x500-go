package x500

import (
	"encoding/asn1"
	"testing"
)

func encodeString(s string) asn1.RawValue {
	bytes, _ := asn1.Marshal(s)
	return asn1.RawValue{FullBytes: bytes}
}

func TestUnmarshal(t *testing.T) {
	type Person struct {
		CommonName string `x500:"oid:2.5.4.3"`
		Surname    string `x500:"oid:2.5.4.4"`
	}
	p := Person{}
	attrs := []Attribute{
		{
			Type:   Id_at_surname,
			Values: []asn1.RawValue{encodeString("Squarepants")},
		},
		{
			Type:   Id_at_commonName,
			Values: []asn1.RawValue{encodeString("Spongebob")},
		},
	}
	err := UnmarshalWithParams(attrs, &p, "")
	if err != nil {
		t.Error(err)
		return
	}
	if p.CommonName != "Spongebob" {
		t.Errorf("commonName was %s", p.CommonName)
		return
	}
	if p.Surname != "Squarepants" {
		t.Errorf("surname was %s", p.Surname)
		return
	}
}

func TestUnmarshalMultivalued(t *testing.T) {
	type Person struct {
		CommonName []string `x500:"oid:2.5.4.3"`
		Surname    []string `x500:"oid:2.5.4.4"`
	}
	p := Person{}
	attrs := []Attribute{
		{
			Type:   Id_at_surname,
			Values: []asn1.RawValue{encodeString("Squarepants")},
		},
		{
			Type:   Id_at_commonName,
			Values: []asn1.RawValue{encodeString("Spongebob")},
		},
	}
	err := UnmarshalWithParams(attrs, &p, "")
	if err != nil {
		t.Error(err)
		return
	}
	if p.CommonName[0] != "Spongebob" {
		t.Errorf("commonName was %s", p.CommonName)
		return
	}
	if p.Surname[0] != "Squarepants" {
		t.Errorf("surname was %s", p.Surname)
		return
	}
}

func TestListUnmarshaling(t *testing.T) {
	type PostalThing struct {
		PostalAddress []string `x500:"oid:2.5.4.43,list"`
	}
	p := PostalThing{
		PostalAddress: []string{
			"P. Sherman",
			"42 Wallaby Way",
			"Sydney, AU",
		},
	}
	attrs, err := MarshalWithParams(p, "")
	if err != nil {
		t.Error(err)
		return
	}
	p.PostalAddress = []string{}
	err = UnmarshalWithParams(attrs, &p, "")
	if err != nil {
		t.Error(err)
		return
	}
	if len(p.PostalAddress) != 3 {
		t.Error("not unmarshaled correctly")
		return
	}
}

// TODO: Pointer-typed fields
// TODO: Test unmarshalling special fields like certs.
