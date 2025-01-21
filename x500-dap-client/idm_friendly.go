package x500_dap_client

// Implementation of SchemaAwareDirectoryAccessClient for the IDMProtocolStack

import (
	"context"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/asn1"
	"time"

	"github.com/Wildboar-Software/x500-go/x500"
)

itu-t(0)
data(9)
pss(2342)
ucl(19200300)
pilot(100)
pilotAttributeType(1)

var Id_at_roomNumber            asn1.ObjectIdentifier = []int{0, 9, 2342, 19200300, 100, 1, 6}
var Id_at_manager               asn1.ObjectIdentifier = []int{0, 9, 2342, 19200300, 100, 1, 10}
var Id_at_homeTelephoneNumber   asn1.ObjectIdentifier = []int{0, 9, 2342, 19200300, 100, 1, 20}
var Id_at_secretary             asn1.ObjectIdentifier = []int{0, 9, 2342, 19200300, 100, 1, 21}
var Id_at_homePostalAddress     asn1.ObjectIdentifier = []int{0, 9, 2342, 19200300, 100, 1, 39}
var Id_at_personalTitle         asn1.ObjectIdentifier = []int{0, 9, 2342, 19200300, 100, 1, 40}
var Id_at_mobileTelephoneNumber asn1.ObjectIdentifier = []int{0, 9, 2342, 19200300, 100, 1, 41}
var Id_at_friendlyCountryName 	asn1.ObjectIdentifier = []int{0, 9, 2342, 19200300, 100, 1, 43}
var Id_at_carLicense            asn1.ObjectIdentifier = []int{2, 16, 840, 1, 113730, 3, 1, 1}
var Id_at_departmentNumber      asn1.ObjectIdentifier = []int{2, 16, 840, 1, 113730, 3, 1, 2}
var Id_at_employeeNumber        asn1.ObjectIdentifier = []int{2, 16, 840, 1, 113730, 3, 1, 3}
var Id_at_employeeType          asn1.ObjectIdentifier = []int{2, 16, 840, 1, 113730, 3, 1, 4}
var Id_at_preferredLanguage     asn1.ObjectIdentifier = []int{2, 16, 840, 1, 113730, 3, 1, 39}
var Id_at_displayName           asn1.ObjectIdentifier = []int{2, 16, 840, 1, 113730, 3, 1, 241}
var Id_at_labeledURI            asn1.ObjectIdentifier = []int{1, 3, 6, 1, 4, 1, 250, 1, 57}
var Id_at_mail                  asn1.ObjectIdentifier = []int{}
var Id_oc_inetOrgPerson         asn1.ObjectIdentifier = []int{2, 16, 840, 1, 113730, 3, 2, 2}
var Id_oc_pilotPerson           asn1.ObjectIdentifier = []int{}
var Id_oc_friendlyCountry       asn1.ObjectIdentifier = []int{}
var Id_oc_document              asn1.ObjectIdentifier = []int{}
var Id_oc_documentSeries        asn1.ObjectIdentifier = []int{}
var Id_oc_room                  asn1.ObjectIdentifier = []int{}
var Id_oc_posixGroup            asn1.ObjectIdentifier = []int{}
var Id_oc_ipService             asn1.ObjectIdentifier = []int{}
var Id_oc_ipNetwork             asn1.ObjectIdentifier = []int{}
var Id_oc_posixAccount          asn1.ObjectIdentifier = []int{}
var Id_oc_ipHost                asn1.ObjectIdentifier = []int{}
var Id_oc_naturalPerson         asn1.ObjectIdentifier = []int{}

func makeStringRDN(attrType asn1.ObjectIdentifier, name string) pkix.RelativeDistinguishedNameSET {
	return pkix.RelativeDistinguishedNameSET{
		pkix.AttributeTypeAndValue{
			Type: attrType,
			Value: asn1.RawValue{
				Class: asn1.ClassUniversal,
				Tag:   asn1.TagUTF8String,
				Bytes: []byte(name),
			},
		},
	}
}

func makeDN(superior DN, attrType asn1.ObjectIdentifier, name string) DN {
	dn := make([]pkix.RelativeDistinguishedNameSET, len(superior), len(superior)+1)
	copy(dn, superior)
	return append(dn, makeStringRDN(attrType, name))
}

func makeStringAttribute(attrType asn1.ObjectIdentifier, name string) x500.Attribute {
	return x500.Attribute{
		Type: attrType,
		Values: []asn1.RawValue{
			asn1.RawValue{
				Class: asn1.ClassUniversal,
				Tag:   asn1.TagUTF8String,
				Bytes: []byte(name),
			},
		},
	}
}

func makeStringsAttribute(attrType asn1.ObjectIdentifier, values []string) x500.Attribute {
	vs := make([]asn1.RawValue, 0, len(values))
	for _, v := range values {
		if len(v) == 0 {
			continue
		}
		value := asn1.RawValue{
			Class: asn1.ClassUniversal,
			Tag:   asn1.TagUTF8String,
			Bytes: []byte(v),
		}
		vs = append(vs, value)
	}
	return x500.Attribute{
		Type:   attrType,
		Values: vs,
	}
}

func makeDNAttribute(attrType asn1.ObjectIdentifier, values []DN) (x500.Attribute, error) {
	vs := make([]asn1.RawValue, 0, len(values))
	for _, v := range values {
		if len(v) == 0 {
			continue
		}
		dnBytes, err := asn1.Marshal(v)
		if err != nil {
			return x500.Attribute{}, err
		}
		value := asn1.RawValue{FullBytes: dnBytes}
		vs = append(vs, value)
	}
	attr := x500.Attribute{
		Type:   attrType,
		Values: vs,
	}
	return attr, nil
}

func makeOIDAttribute(attrType, value asn1.ObjectIdentifier) (x500.Attribute, error) {
	valueBytes, err := asn1.Marshal(value)
	if err != nil {
		return x500.Attribute{}, err
	}
	attr := x500.Attribute{
		Type:   attrType,
		Values: []asn1.RawValue{asn1.RawValue{FullBytes: valueBytes}},
	}
	return attr, nil
}

func createObjectClassAttr(ocs []asn1.ObjectIdentifier) (x500.Attribute, error) {
	values := make([]asn1.RawValue, 0, len(ocs))
	for _, oc := range ocs {
		ocBytes, err := asn1.Marshal(oc)
		if err != nil {
			return x500.Attribute{}, err
		}
		value := asn1.RawValue{FullBytes: ocBytes}
		values = append(values, value)
	}
	attr := x500.Attribute{
		Type:   x500.Id_at_objectClass,
		Values: values,
	}
	return attr, nil
}

func (stack *IDMProtocolStack) CreateSubentry(ctx context.Context, superior DN, name string, subtreeSpec []x500.SubtreeSpecification) (resp X500OpOutcome, result *x500.AddEntryResultData, err error) {
	dn := makeDN(superior, x500.Id_at_commonName, name)
	attrs := make([]x500.Attribute, 0, 3)
	ocAttr, err := createObjectClassAttr([]asn1.ObjectIdentifier{x500.Id_sc_subentry})
	if err != nil {
		return X500OpOutcome{}, nil, err
	}
	cnAttr := makeStringAttribute(x500.Id_at_commonName, name)
	attrs = append(attrs, ocAttr)
	attrs = append(attrs, cnAttr)
	if len(subtreeSpec) > 0 {
		ssValues := make([]asn1.RawValue, len(subtreeSpec))
		for _, ss := range subtreeSpec {
			ssBytes, err := asn1.Marshal(ss)
			if err != nil {
				return X500OpOutcome{}, nil, err
			}
			ssValues = append(ssValues, asn1.RawValue{FullBytes: ssBytes})
		}
		ssAttr := x500.Attribute{
			Type:   x500.Id_oa_subtreeSpecification,
			Values: ssValues,
		}
		attrs = append(attrs, ssAttr)
	}
	return stack.AddEntrySimple(ctx, dn, attrs)
}

// This does NOT create the objectClass attribute.
func createPersonAttributes(opts *CreatePersonOptions) ([]x500.Attribute, error) {
	attrs := make([]x500.Attribute, 0, 32)
	if len(opts.commonName) > 0 {
		attrs = append(attrs, makeStringAttribute(x500.Id_at_commonName, opts.commonName))
	}
	if len(opts.givenName) > 0 {
		attrs = append(attrs, makeStringAttribute(x500.Id_at_givenName, opts.givenName))
	}
	if len(opts.givenName) > 0 {
		attrs = append(attrs, makeStringAttribute(x500.Id_at_givenName, opts.givenName))
	}
	if len(opts.initials) > 0 {
		attrs = append(attrs, makeStringAttribute(x500.Id_at_initials, opts.initials))
	}
	if len(opts.surname) > 0 {
		attrs = append(attrs, makeStringAttribute(x500.Id_at_surname, opts.surname))
	}
	if len(opts.generationQualifier) > 0 {
		attrs = append(attrs, makeStringAttribute(x500.Id_at_generationQualifier, opts.generationQualifier))
	}
	if len(opts.displayName) > 0 {
		attrs = append(attrs, makeStringAttribute(Id_at_displayName, opts.displayName))
	}
	if len(opts.stateOrProvinceName) > 0 {
		attrs = append(attrs, makeStringAttribute(x500.Id_at_stateOrProvinceName, opts.stateOrProvinceName))
	}
	if len(opts.streetAddress) > 0 {
		attrs = append(attrs, makeStringAttribute(x500.Id_at_streetAddress, opts.streetAddress))
	}
	if len(opts.employeeNumber) > 0 {
		attrs = append(attrs, makeStringAttribute(Id_at_employeeNumber, opts.employeeNumber))
	}
	if len(opts.preferredLanguage) > 0 {
		attrs = append(attrs, makeStringAttribute(Id_at_preferredLanguage, opts.preferredLanguage))
	}
	if len(opts.postalCode) > 0 {
		attrs = append(attrs, makeStringAttribute(x500.Id_at_postalCode, opts.postalCode))
	}
	if len(opts.poBox) > 0 {
		attrs = append(attrs, makeStringAttribute(x500.Id_at_postOfficeBox, opts.poBox))
	}
	if len(opts.description) > 0 {
		attrs = append(attrs, makeStringsAttribute(x500.Id_at_description, opts.description))
	}
	if len(opts.localityName) > 0 {
		attrs = append(attrs, makeStringsAttribute(x500.Id_at_localityName, opts.localityName))
	}
	if len(opts.businessCategory) > 0 {
		attrs = append(attrs, makeStringsAttribute(x500.Id_at_businessCategory, opts.businessCategory))
	}
	if len(opts.organizationName) > 0 {
		attrs = append(attrs, makeStringsAttribute(x500.Id_at_organizationName, opts.organizationName))
	}
	if len(opts.title) > 0 {
		attrs = append(attrs, makeStringsAttribute(x500.Id_at_title, opts.title))
	}
	if len(opts.personalTitle) > 0 {
		attrs = append(attrs, makeStringsAttribute(Id_at_personalTitle, opts.personalTitle))
	}
	if len(opts.uid) > 0 {
		attrs = append(attrs, makeStringsAttribute(x500.Id_coat_uid, opts.uid))
	}
	if len(opts.carLicense) > 0 {
		attrs = append(attrs, makeStringsAttribute(Id_at_carLicense, opts.carLicense))
	}
	if len(opts.departmentNumber) > 0 {
		attrs = append(attrs, makeStringsAttribute(Id_at_departmentNumber, opts.departmentNumber))
	}
	if len(opts.employeeType) > 0 {
		attrs = append(attrs, makeStringsAttribute(Id_at_employeeType, opts.employeeType))
	}
	if len(opts.labeledURI) > 0 {
		attrs = append(attrs, makeStringsAttribute(Id_at_labeledURI, opts.labeledURI))
	}
	if len(opts.roomNumber) > 0 {
		attrs = append(attrs, makeStringsAttribute(Id_at_roomNumber, opts.roomNumber))
	}
	if len(opts.telephoneNumber) > 0 {
		vs := make([]asn1.RawValue, 0, len(opts.telephoneNumber))
		for _, v := range opts.telephoneNumber {
			if len(v) == 0 {
				continue
			}
			value := asn1.RawValue{
				Class: asn1.ClassUniversal,
				Tag:   asn1.TagPrintableString,
				Bytes: []byte(v),
			}
			vs = append(vs, value)
		}
		attr := x500.Attribute{
			Type:   x500.Id_at_telephoneNumber,
			Values: vs,
		}
		attrs = append(attrs, attr)
	}
	if len(opts.homeTelephoneNumber) > 0 {
		vs := make([]asn1.RawValue, 0, len(opts.homeTelephoneNumber))
		for _, v := range opts.homeTelephoneNumber {
			if len(v) == 0 {
				continue
			}
			value := asn1.RawValue{
				Class: asn1.ClassUniversal,
				Tag:   asn1.TagPrintableString,
				Bytes: []byte(v),
			}
			vs = append(vs, value)
		}
		attr := x500.Attribute{
			Type:   Id_at_homeTelephoneNumber,
			Values: vs,
		}
		attrs = append(attrs, attr)
	}
	if len(opts.mobileTelephoneNumber) > 0 {
		vs := make([]asn1.RawValue, 0, len(opts.mobileTelephoneNumber))
		for _, v := range opts.mobileTelephoneNumber {
			if len(v) == 0 {
				continue
			}
			value := asn1.RawValue{
				Class: asn1.ClassUniversal,
				Tag:   asn1.TagPrintableString,
				Bytes: []byte(v),
			}
			vs = append(vs, value)
		}
		attr := x500.Attribute{
			Type:   Id_at_mobileTelephoneNumber,
			Values: vs,
		}
		attrs = append(attrs, attr)
	}
	if len(opts.faxNumber) > 0 {
		vs := make([]asn1.RawValue, 0, len(opts.faxNumber))
		for _, v := range opts.faxNumber {
			if len(v) == 0 {
				continue
			}
			fax := x500.FacsimileTelephoneNumber{TelephoneNumber: v}
			faxBytes, err := asn1.Marshal(fax)
			if err != nil {
				return attrs, err
			}
			value := asn1.RawValue{FullBytes: faxBytes}
			vs = append(vs, value)
		}
		attr := x500.Attribute{
			Type:   x500.Id_at_facsimileTelephoneNumber,
			Values: vs,
		}
		attrs = append(attrs, attr)
	}
	if len(opts.userPassword) > 0 {
		value := asn1.RawValue{
			Class: asn1.ClassUniversal,
			Tag:   asn1.TagOctetString,
			Bytes: []byte(opts.userPassword),
		}
		attr := x500.Attribute{
			Type:   x500.Id_at_userPassword,
			Values: []asn1.RawValue{value},
		}
		attrs = append(attrs, attr)
	}

	if len(opts.manager) > 0 {
		attr, err := makeDNAttribute(Id_at_manager, opts.manager)
		if err != nil {
			return attrs, err
		}
		attrs = append(attrs, attr)
	}
	if len(opts.secretary) > 0 {
		attr, err := makeDNAttribute(Id_at_secretary, opts.secretary)
		if err != nil {
			return attrs, err
		}
		attrs = append(attrs, attr)
	}
	if len(opts.seeAlso) > 0 {
		attr, err := makeDNAttribute(x500.Id_at_seeAlso, opts.secretary)
		if err != nil {
			return attrs, err
		}
		attrs = append(attrs, attr)
	}
	if len(opts.postalAddress) > 0 {
		addr, err := asn1.Marshal(opts.postalAddress)
		if err != nil {
			return attrs, err
		}
		attr := x500.Attribute{
			Type:   x500.Id_at_postalAddress,
			Values: []asn1.RawValue{asn1.RawValue{FullBytes: addr}},
		}
		attrs = append(attrs, attr)
	}
	if len(opts.homePostalAddress) > 0 {
		addr, err := asn1.Marshal(opts.homePostalAddress)
		if err != nil {
			return attrs, err
		}
		attr := x500.Attribute{
			Type:   Id_at_homePostalAddress,
			Values: []asn1.RawValue{asn1.RawValue{FullBytes: addr}},
		}
		attrs = append(attrs, attr)
	}
	if len(opts.mail) > 0 {
		vs := make([]asn1.RawValue, 0, len(opts.mail))
		for _, v := range opts.mail {
			if len(v) == 0 {
				continue
			}
			value := asn1.RawValue{
				Class: asn1.ClassUniversal,
				Tag:   asn1.TagIA5String,
				Bytes: []byte(v),
			}
			vs = append(vs, value)
		}
		attr := x500.Attribute{
			Type:   Id_at_mail,
			Values: vs,
		}
		attrs = append(attrs, attr)
	}
	if len(opts.userCertificate) > 0 {
		vs := make([]asn1.RawValue, 0, len(opts.userCertificate))
		for _, v := range opts.userCertificate {
			value := asn1.RawValue{FullBytes: v.Raw}
			vs = append(vs, value)
		}
		attr := x500.Attribute{
			Type:   x500.Id_at_userCertificate,
			Values: vs,
		}
		attrs = append(attrs, attr)
	}

	if len(opts.uniqueIdentifier) > 0 {
		vs := make([]asn1.RawValue, 0, len(opts.uniqueIdentifier))
		for _, v := range opts.uniqueIdentifier {
			if v.BitLength == 0 {
				continue
			}
			bitsBytes, err := asn1.Marshal(v)
			if err != nil {
				return attrs, err
			}
			value := asn1.RawValue{FullBytes: bitsBytes}
			vs = append(vs, value)
		}
		attr := x500.Attribute{
			Type:   x500.Id_at_uniqueIdentifier,
			Values: vs,
		}
		attrs = append(attrs, attr)
	}
	return attrs, nil
}

// TODO: Just simplify the signature.
func (stack *IDMProtocolStack) CreatePerson(ctx context.Context, superior DN, cn, surname string, opts *CreatePersonOptions) (resp X500OpOutcome, result *x500.AddEntryResultData, err error) {
	dn := makeDN(superior, x500.Id_at_commonName, cn)
	attrs := make([]x500.Attribute, 0, 32)
	cnAttr := makeStringAttribute(x500.Id_at_commonName, cn)
	attrs = append(attrs, cnAttr)
	// FIXME:
	return stack.AddEntrySimple(ctx, dn, attrs)
}

// TODO: Change the signature. Just use opts alone.
func (stack *IDMProtocolStack) CreateOrgPerson(ctx context.Context, superior DN, cn, surname string, opts *CreatePersonOptions) (resp X500OpOutcome, result *x500.AddEntryResultData, err error) {
	dn := makeDN(superior, x500.Id_at_commonName, cn)
	attrs, err := createPersonAttributes(opts)
	if err != nil {
		return X500OpOutcome{}, nil, err
	}
	ocAttr, err := createObjectClassAttr([]asn1.ObjectIdentifier{x500.Id_oc_person, x500.Id_oc_organizationalPerson})
	if err != nil {
		return X500OpOutcome{}, nil, err
	}
	attrs = append(attrs, ocAttr)
	return stack.AddEntrySimple(ctx, dn, attrs)
}

// TODO: DN is not called "superior" in the interface
// TODO: Change the signature. Just use opts alone.
func (stack *IDMProtocolStack) CreateResPerson(ctx context.Context, superior DN, cn, surname string, opts *CreatePersonOptions) (resp X500OpOutcome, result *x500.AddEntryResultData, err error) {
	dn := makeDN(superior, x500.Id_at_commonName, cn)
	attrs, err := createPersonAttributes(opts)
	if err != nil {
		return X500OpOutcome{}, nil, err
	}
	ocAttr, err := createObjectClassAttr([]asn1.ObjectIdentifier{x500.Id_oc_person, x500.Id_oc_residentialPerson})
	if err != nil {
		return X500OpOutcome{}, nil, err
	}
	attrs = append(attrs, ocAttr)
	return stack.AddEntrySimple(ctx, dn, attrs)
}

// TODO: DN is not called "superior" in the interface
// TODO: Change the signature. Just use opts alone.
func (stack *IDMProtocolStack) CreateInetOrgPerson(ctx context.Context, superior DN, cn, surname string, opts *CreatePersonOptions) (resp X500OpOutcome, result *x500.AddEntryResultData, err error) {
	dn := makeDN(superior, x500.Id_at_commonName, cn)
	attrs, err := createPersonAttributes(opts)
	if err != nil {
		return X500OpOutcome{}, nil, err
	}
	ocAttr, err := createObjectClassAttr([]asn1.ObjectIdentifier{x500.Id_oc_person, x500.Id_oc_organizationalPerson, Id_oc_inetOrgPerson})
	if err != nil {
		return X500OpOutcome{}, nil, err
	}
	attrs = append(attrs, ocAttr)
	return stack.AddEntrySimple(ctx, dn, attrs)
}

// TODO: DN is not called "superior" in the interface
// TODO: Change the signature. Just use opts alone.
func (stack *IDMProtocolStack) CreatePilotPerson(ctx context.Context, superior DN, cn, surname string, opts *CreatePersonOptions) (resp X500OpOutcome, result *x500.AddEntryResultData, err error) {
	dn := makeDN(superior, x500.Id_at_commonName, cn)
	attrs, err := createPersonAttributes(opts)
	if err != nil {
		return X500OpOutcome{}, nil, err
	}
	ocAttr, err := createObjectClassAttr([]asn1.ObjectIdentifier{x500.Id_oc_person, Id_oc_pilotPerson})
	if err != nil {
		return X500OpOutcome{}, nil, err
	}
	attrs = append(attrs, ocAttr)
	return stack.AddEntrySimple(ctx, dn, attrs)
}

// If name is not "", the friendlyCountry object class is used.
func (stack *IDMProtocolStack) CreateCountry(ctx context.Context, superior DN, iso2Code, desc, name string) (resp X500OpOutcome, result *x500.AddEntryResultData, err error) {
	countryNameValue := asn1.RawValue{
		Class: asn1.ClassUniversal,
		Tag:   asn1.TagPrintableString,
		Bytes: []byte(iso2Code),
	}
	dn := make([]pkix.RelativeDistinguishedNameSET, len(superior), len(superior)+1)
	copy(dn, superior)
	rdn := pkix.RelativeDistinguishedNameSET{
		pkix.AttributeTypeAndValue{
			Type: x500.Id_at_countryName,
			Value: countryNameValue,
		},
	}
	dn = append(dn, rdn)
	attrs := make([]x500.Attribute, 0, 4)
	attrs = append(attrs, x500.Attribute{
		Type: x500.Id_at_countryName,
		Values: []asn1.RawValue{countryNameValue},
	})
	objectClasses := []asn1.ObjectIdentifier{x500.Id_oc_country}
	if len(name) > 0 {
		objectClasses = append(objectClasses, Id_oc_friendlyCountry)
	}
	ocAttr, err := createObjectClassAttr(objectClasses)
	if err != nil {
		return X500OpOutcome{}, nil, err
	}
	attrs = append(attrs, ocAttr)
	if len(desc) > 0 {
		attrs = append(attrs, makeStringAttribute(x500.Id_at_description, desc))
	}
	if len(name) > 0 {
		attrs = append(attrs, makeStringAttribute(Id_at_friendlyCountryName, name))
	}
	return stack.AddEntrySimple(ctx, dn, attrs)
}

func (stack *IDMProtocolStack) CreateStateOrProvince(ctx context.Context, superior DN, name, desc string) (resp X500OpOutcome, result *x500.AddEntryResultData, err error) {
	dn := makeDN(superior, x500.Id_at_stateOrProvinceName, name)
	ocAttr, err := createObjectClassAttr([]asn1.ObjectIdentifier{x500.Id_oc_locality})
	if err != nil {
		return X500OpOutcome{}, nil, err
	}
	attrs = append(attrs, ocAttr)
	attrs = append(attrs, makeStringAttribute(x500.Id_at_stateOrProvinceName, name))
	if len(desc) > 0 {
		attrs = append(attrs, makeStringAttribute(x500.Id_at_description, desc))
	}
	return stack.AddEntrySimple(ctx, dn, attrs)
}

func (stack *IDMProtocolStack) CreateLocality(ctx context.Context, superior DN, name, desc string) (resp X500OpOutcome, result *x500.AddEntryResultData, err error) {
	dn := makeDN(superior, x500.Id_at_localityName, name)
	ocAttr, err := createObjectClassAttr([]asn1.ObjectIdentifier{x500.Id_oc_locality})
	if err != nil {
		return X500OpOutcome{}, nil, err
	}
	attrs = append(attrs, ocAttr)
	attrs = append(attrs, makeStringAttribute(x500.Id_at_localityName, name))
	if len(desc) > 0 {
		attrs = append(attrs, makeStringAttribute(x500.Id_at_description, desc))
	}
	return stack.AddEntrySimple(ctx, dn, attrs)
}

func (stack *IDMProtocolStack) CreateOrganization(ctx context.Context, superior DN, name string, opts *CreateOrganizationOptions) (resp X500OpOutcome, result *x500.AddEntryResultData, err error) {
	return stack.AddEntrySimple(ctx, dn, attrs)
}

func (stack *IDMProtocolStack) CreateOrgUnit(ctx context.Context, superior DN, name string, opts *CreateOrganizationOptions) (resp X500OpOutcome, result *x500.AddEntryResultData, err error) {
	return stack.AddEntrySimple(ctx, dn, attrs)
}

func (stack *IDMProtocolStack) CreateOrgRole(ctx context.Context, superior DN, name, desc string, members []DN) (resp X500OpOutcome, result *x500.AddEntryResultData, err error) {
	return stack.AddEntrySimple(ctx, dn, attrs)
}

func (stack *IDMProtocolStack) CreateProcess(ctx context.Context, superior DN, name, desc string) (resp X500OpOutcome, result *x500.AddEntryResultData, err error) {
	return stack.AddEntrySimple(ctx, dn, attrs)
}

func (stack *IDMProtocolStack) CreateAppEntity(ctx context.Context, superior DN, name string, addr x500.PresentationAddress, desc string) (resp X500OpOutcome, result *x500.AddEntryResultData, err error) {
	return stack.AddEntrySimple(ctx, dn, attrs)
}

func (stack *IDMProtocolStack) CreateDSA(ctx context.Context, superior DN, name string, addr x500.PresentationAddress, desc string) (resp X500OpOutcome, result *x500.AddEntryResultData, err error) {
	return stack.AddEntrySimple(ctx, dn, attrs)
}

func (stack *IDMProtocolStack) CreateDevice(ctx context.Context, superior DN, name, serial, desc string, owner DN) (resp X500OpOutcome, result *x500.AddEntryResultData, err error) {
	return stack.AddEntrySimple(ctx, dn, attrs)
}

func (stack *IDMProtocolStack) CreateDMD(ctx context.Context, superior DN, name, desc string) (resp X500OpOutcome, result *x500.AddEntryResultData, err error) {
	return stack.AddEntrySimple(ctx, dn, attrs)
}

func (stack *IDMProtocolStack) CreateOIDRoot(ctx context.Context, superior DN, arc1, arc2, arc3 int, aliasedEntry DN) (resp X500OpOutcome, result *x500.AddEntryResultData, err error) {
	return stack.AddEntrySimple(ctx, dn, attrs)
}

func (stack *IDMProtocolStack) CreateOIDArc(ctx context.Context, superior DN, arc int, aliasedEntry DN) (resp X500OpOutcome, result *x500.AddEntryResultData, err error) {
	return stack.AddEntrySimple(ctx, dn, attrs)
}

func (stack *IDMProtocolStack) CreateGroupOfUrlsEntry(ctx context.Context, superior DN, name, desc, org string, members []string, owner DN) (resp X500OpOutcome, result *x500.AddEntryResultData, err error) {
	return stack.AddEntrySimple(ctx, dn, attrs)
}

func (stack *IDMProtocolStack) CreateAliasEntry(ctx context.Context, dn, aliasedEntry DN) (resp X500OpOutcome, result *x500.AddEntryResultData, err error) {
	return stack.AddEntrySimple(ctx, dn, attrs)
}

func (stack *IDMProtocolStack) CreateCRLDPEntry(ctx context.Context, superior DN, name string, crl, eepk, authority, delta *x509.RevocationList) (resp X500OpOutcome, result *x500.AddEntryResultData, err error) {
	return stack.AddEntrySimple(ctx, dn, attrs)
}

func (stack *IDMProtocolStack) CreateDomain(ctx context.Context, superior DN, dc string, opts *CreateDomainOptions) (resp X500OpOutcome, result *x500.AddEntryResultData, err error) {
	return stack.AddEntrySimple(ctx, dn, attrs)
}

func (stack *IDMProtocolStack) CreateAccount(ctx context.Context, superior DN, uid string, opts *CreateAccountOptions) (resp X500OpOutcome, result *x500.AddEntryResultData, err error) {
	return stack.AddEntrySimple(ctx, dn, attrs)
}

func (stack *IDMProtocolStack) CreateDocument(ctx context.Context, superior DN, docid string, opts *CreateDocumentOptions) (resp X500OpOutcome, result *x500.AddEntryResultData, err error) {
	return stack.AddEntrySimple(ctx, dn, attrs)
}

func (stack *IDMProtocolStack) CreateDocumentSeries(ctx context.Context, superior DN, name string, opts *CreateDocumentSeriesOptions) (resp X500OpOutcome, result *x500.AddEntryResultData, err error) {
	return stack.AddEntrySimple(ctx, dn, attrs)
}

func (stack *IDMProtocolStack) CreateRoom(ctx context.Context, superior DN, name, number, desc, phone string, seeAlso []DN) (resp X500OpOutcome, result *x500.AddEntryResultData, err error) {
	return stack.AddEntrySimple(ctx, dn, attrs)
}

func (stack *IDMProtocolStack) CreatePosixGroup(ctx context.Context, superior DN, gid int, name, desc, userpwd string, members []string) (resp X500OpOutcome, result *x500.AddEntryResultData, err error) {
	return stack.AddEntrySimple(ctx, dn, attrs)
}

func (stack *IDMProtocolStack) CreateIPService(ctx context.Context, superior DN, name, transport string, port int, manager DN) (resp X500OpOutcome, result *x500.AddEntryResultData, err error) {
	return stack.AddEntrySimple(ctx, dn, attrs)
}

func (stack *IDMProtocolStack) CreateIPNetwork(ctx context.Context, superior DN, name, ip, netmask, loc, desc string) (resp X500OpOutcome, result *x500.AddEntryResultData, err error) {
	return stack.AddEntrySimple(ctx, dn, attrs)
}

func (stack *IDMProtocolStack) BecomePasswordAdminSubentry(ctx context.Context, dn DN, pwdAttribute x500.AttributeType, opts *PasswordPolicyOptions) (resp X500OpOutcome, result *x500.ModifyEntryResultData, err error) {
	return stack.ModifyEntry(ctx, arg_data)
}

func (stack *IDMProtocolStack) BecomeSubschemaSubentry(ctx context.Context, dn DN, schema *Subschema) (resp X500OpOutcome, result *x500.ModifyEntryResultData, err error) {
	return stack.ModifyEntry(ctx, arg_data)
}

func (stack *IDMProtocolStack) BecomeServiceAdminSubentry(ctx context.Context, dn DN, searchRules []x500.SearchRuleDescription) (resp X500OpOutcome, result *x500.ModifyEntryResultData, err error) {
	return stack.ModifyEntry(ctx, arg_data)
}

func (stack *IDMProtocolStack) BecomeContextAssertionDefaultSubentry(ctx context.Context, dn DN, cads []x500.TypeAndContextAssertion) (resp X500OpOutcome, result *x500.ModifyEntryResultData, err error) {
	return stack.ModifyEntry(ctx, arg_data)
}

func (stack *IDMProtocolStack) BecomePMIUser(ctx context.Context, dn DN, ac *x500.AttributeCertificate) (resp X500OpOutcome, result *x500.ModifyEntryResultData, err error) {
	return stack.ModifyEntry(ctx, arg_data)
}

func (stack *IDMProtocolStack) BecomePMIAA(ctx context.Context, dn DN, aaCert *x500.AttributeCertificate, crls *CRLOptions) (resp X500OpOutcome, result *x500.ModifyEntryResultData, err error) {
	return stack.ModifyEntry(ctx, arg_data)
}

func (stack *IDMProtocolStack) BecomePMISOA(ctx context.Context, dn DN, adc *x500.AttributeCertificate, crls *CRLOptions) (resp X500OpOutcome, result *x500.ModifyEntryResultData, err error) {
	return stack.ModifyEntry(ctx, arg_data)
}

func (stack *IDMProtocolStack) BecomeAttCertCRLDistPoint(ctx context.Context, dn DN, crls *CRLOptions) (resp X500OpOutcome, result *x500.ModifyEntryResultData, err error) {
	return stack.ModifyEntry(ctx, arg_data)
}

func (stack *IDMProtocolStack) BecomePKIUser(ctx context.Context, dn DN, userCert *x509.Certificate) (resp X500OpOutcome, result *x500.ModifyEntryResultData, err error) {
	return stack.ModifyEntry(ctx, arg_data)
}

func (stack *IDMProtocolStack) BecomePKICA(ctx context.Context, dn DN, cacert *x509.Certificate, opts *CRLOptions) (resp X500OpOutcome, result *x500.ModifyEntryResultData, err error) {
	return stack.ModifyEntry(ctx, arg_data)
}

func (stack *IDMProtocolStack) BecomePMIDelegationPath(ctx context.Context, dn DN, path []x500.AttCertPath) (resp X500OpOutcome, result *x500.ModifyEntryResultData, err error) {
	return stack.ModifyEntry(ctx, arg_data)
}

func (stack *IDMProtocolStack) BecomePrivilegePolicy(ctx context.Context, dn DN, policies []x500.PolicySyntax) (resp X500OpOutcome, result *x500.ModifyEntryResultData, err error) {
	return stack.ModifyEntry(ctx, arg_data)
}

func (stack *IDMProtocolStack) BecomeProtectedPrivilegePolicy(ctx context.Context, dn DN, policies []x500.AttributeCertificate) (resp X500OpOutcome, result *x500.ModifyEntryResultData, err error) {
	return stack.ModifyEntry(ctx, arg_data)
}

func (stack *IDMProtocolStack) BecomeDeltaCRL(ctx context.Context, dn DN, deltaCRLs []x509.RevocationList) (resp X500OpOutcome, result *x500.ModifyEntryResultData, err error) {
	return stack.ModifyEntry(ctx, arg_data)
}

func (stack *IDMProtocolStack) BecomeCPCPS(ctx context.Context, dn DN, statements []x500.InfoSyntax, policies []x500.PolicySyntax) (resp X500OpOutcome, result *x500.ModifyEntryResultData, err error) {
	return stack.ModifyEntry(ctx, arg_data)
}

func (stack *IDMProtocolStack) BecomePKICertPath(ctx context.Context, dn DN, paths []x500.PkiPath) (resp X500OpOutcome, result *x500.ModifyEntryResultData, err error) {
	return stack.ModifyEntry(ctx, arg_data)
}

func (stack *IDMProtocolStack) BecomeStrongAuthenticationUser(ctx context.Context, dn DN, usercerts []x509.Certificate) (resp X500OpOutcome, result *x500.ModifyEntryResultData, err error) {
	return stack.ModifyEntry(ctx, arg_data)
}

func (stack *IDMProtocolStack) BecomeUserSecurityInfo(ctx context.Context, dn DN, supportedAlgs []x500.SupportedAlgorithm) (resp X500OpOutcome, result *x500.ModifyEntryResultData, err error) {
	return stack.ModifyEntry(ctx, arg_data)
}

func (stack *IDMProtocolStack) BecomeUserPwdClass(ctx context.Context, dn DN, pwd x500.UserPwd) (resp X500OpOutcome, result *x500.ModifyEntryResultData, err error) {
	return stack.ModifyEntry(ctx, arg_data)
}

func (stack *IDMProtocolStack) BecomeCertificationAuthority(ctx context.Context, dn DN, cert x509.Certificate, crls *CRLOptions) (resp X500OpOutcome, result *x500.ModifyEntryResultData, err error) {
	return stack.ModifyEntry(ctx, arg_data)
}

func (stack *IDMProtocolStack) BecomeISOTagInfo(ctx context.Context, dn DN, opts *BecomeISOTagInfoOptions) (resp X500OpOutcome, result *x500.ModifyEntryResultData, err error) {
	return stack.ModifyEntry(ctx, arg_data)
}

func (stack *IDMProtocolStack) BecomeISOTagType(ctx context.Context, dn DN, opts *BecomeISOTagTypeOptions) (resp X500OpOutcome, result *x500.ModifyEntryResultData, err error) {
	return stack.ModifyEntry(ctx, arg_data)
}

func (stack *IDMProtocolStack) BecomeEPCTagInfo(ctx context.Context, dn DN, opts *BecomeEPCTaginfoOptions) (resp X500OpOutcome, result *x500.ModifyEntryResultData, err error) {
	return stack.ModifyEntry(ctx, arg_data)
}

func (stack *IDMProtocolStack) BecomeEPCTagType(ctx context.Context, dn DN, format x500.UiiFormat) (resp X500OpOutcome, result *x500.ModifyEntryResultData, err error) {
	return stack.ModifyEntry(ctx, arg_data)
}

func (stack *IDMProtocolStack) BecomePosixAccount(ctx context.Context, dn DN, opts *BecomePosixAccountOptions) (resp X500OpOutcome, result *x500.ModifyEntryResultData, err error) {
	return stack.ModifyEntry(ctx, arg_data)
}

func (stack *IDMProtocolStack) BecomeShadowAccount(ctx context.Context, dn DN, opts *BecomeShadowAccount) (resp X500OpOutcome, result *x500.ModifyEntryResultData, err error) {
	return stack.ModifyEntry(ctx, arg_data)
}

func (stack *IDMProtocolStack) BecomeIEEE802Device(ctx context.Context, dn DN, macAddress string) (resp X500OpOutcome, result *x500.ModifyEntryResultData, err error) {
	return stack.ModifyEntry(ctx, arg_data)
}

func (stack *IDMProtocolStack) BecomeIPHost(ctx context.Context, dn DN, cn, ip, desc, loc string, manager DN) (resp X500OpOutcome, result *x500.ModifyEntryResultData, err error) {
	return stack.ModifyEntry(ctx, arg_data)
}

func (stack *IDMProtocolStack) BecomeNaturalPerson(ctx context.Context, dn DN, opts *BecomeNaturalPersonOptions) (resp X500OpOutcome, result *x500.ModifyEntryResultData, err error) {
	return stack.ModifyEntry(ctx, arg_data)
}

func (stack *IDMProtocolStack) GetSubschema(ctx context.Context, dn DN) (resp X500OpOutcome, result *Subschema, err error) {

}

func (stack *IDMProtocolStack) ListAttributeTypes(ctx context.Context, dn DN) (resp X500OpOutcome, result *[]x500.AttributeTypeDescription, err error) {

}

func (stack *IDMProtocolStack) ListObjectClasses(ctx context.Context, dn DN) (resp X500OpOutcome, result *[]x500.ObjectClassDescription, err error) {

}

func (stack *IDMProtocolStack) ListNameForms(ctx context.Context, dn DN) (resp X500OpOutcome, result *[]x500.NameFormDescription, err error) {

}

func (stack *IDMProtocolStack) ListStructureRules(ctx context.Context, dn DN) (resp X500OpOutcome, result *[]x500.DITStructureRuleDescription, err error) {

}

func (stack *IDMProtocolStack) ListContentRules(ctx context.Context, dn DN) (resp X500OpOutcome, result *[]x500.DITContentRuleDescription, err error) {

}

func (stack *IDMProtocolStack) ListFriendships(ctx context.Context, dn DN) (resp X500OpOutcome, result *[]x500.FriendsDescription, err error) {

}

func (stack *IDMProtocolStack) ListMatchingRules(ctx context.Context, dn DN) (resp X500OpOutcome, result *[]x500.MatchingRuleDescription, err error) {

}

func (stack *IDMProtocolStack) ListMatchingRuleUses(ctx context.Context, dn DN) (resp X500OpOutcome, result *[]x500.MatchingRuleUseDescription, err error) {

}

func (stack *IDMProtocolStack) ListSyntaxNames(ctx context.Context, dn DN) (resp X500OpOutcome, result *[]x500.LdapSyntaxDescription, err error) {

}

func (stack *IDMProtocolStack) ListContextUses(ctx context.Context, dn DN) (resp X500OpOutcome, result *[]x500.DITContextUseDescription, err error) {

}

func (stack *IDMProtocolStack) ListSearchRules(ctx context.Context, dn DN) (resp X500OpOutcome, result *[]x500.SearchRuleDescription, err error) {

}

func (stack *IDMProtocolStack) ListSubentries(ctx context.Context, dn DN) (resp X500OpOutcome, result *x500.ListResultData, err error) {

}

func (stack *IDMProtocolStack) GroupAdd(ctx context.Context, group, member DN, uid *asn1.BitString) (resp X500OpOutcome, result *x500.ModifyEntryResultData, err error) {
	return stack.ModifyEntry(ctx, arg_data)
}

func (stack *IDMProtocolStack) GroupRemove(ctx context.Context, group, member DN, uid *asn1.BitString) (resp X500OpOutcome, result *x500.ModifyEntryResultData, err error) {
	return stack.ModifyEntry(ctx, arg_data)
}

func (stack *IDMProtocolStack) GroupCheckMember(ctx context.Context, group, member DN, uid *asn1.BitString) (resp X500OpOutcome, result *x500.ModifyEntryResultData, err error) {
	return stack.ModifyEntry(ctx, arg_data)
}

func (stack *IDMProtocolStack) AddPublicKeyCertificate(ctx context.Context, dn DN, cert x509.Certificate) (resp X500OpOutcome, result *x500.ModifyEntryResultData, err error) {
	return stack.ModifyEntry(ctx, arg_data)
}

func (stack *IDMProtocolStack) AddAttributeCertificate(ctx context.Context, dn DN, acert x500.AttributeCertificate) (resp X500OpOutcome, result *x500.ModifyEntryResultData, err error) {
	return stack.ModifyEntry(ctx, arg_data)
}

func (stack *IDMProtocolStack) AddAdministrativeRole(ctx context.Context, dn DN, role asn1.ObjectIdentifier) (resp X500OpOutcome, result *x500.ModifyEntryResultData, err error) {
	return stack.AddValues(ctx, dn, attr)
}

func (stack *IDMProtocolStack) RemoveAdministrativeRole(ctx context.Context, dn DN, role asn1.ObjectIdentifier) (resp X500OpOutcome, result *x500.ModifyEntryResultData, err error) {
	return stack.RemoveValues(ctx, dn, attr)
}

func (stack *IDMProtocolStack) AddCollectiveExclusions(ctx context.Context, dn DN, attr x500.AttributeType) (resp X500OpOutcome, result *x500.ModifyEntryResultData, err error) {
	return stack.AddValues(ctx, dn, attr)
}

func (stack *IDMProtocolStack) SetAliasedEntryName(ctx context.Context, src, dest DN) (resp X500OpOutcome, result *x500.ModifyEntryResultData, err error) {
	return stack.AddValues(ctx, dn, attr)
}

func (stack *IDMProtocolStack) AddPermission(ctx context.Context, dn DN, operation, object string) (resp X500OpOutcome, result *x500.ModifyEntryResultData, err error) {
	return stack.AddValues(ctx, dn, attr)
}

func (stack *IDMProtocolStack) AddHierarchicalParent(ctx context.Context, dn, parent x500.DistinguishedName) (resp X500OpOutcome, result *x500.ModifyEntryResultData, err error) {
	return stack.AddValues(ctx, dn, attr)
}

func (stack *IDMProtocolStack) AddObjectClass(ctx context.Context, dn DN, oc asn1.ObjectIdentifier) (resp X500OpOutcome, result *x500.ModifyEntryResultData, err error) {
	return stack.AddValues(ctx, dn, attr)
}

func (stack *IDMProtocolStack) RemoveObjectClass(ctx context.Context, dn DN, oc asn1.ObjectIdentifier) (resp X500OpOutcome, result *x500.ModifyEntryResultData, err error) {
	return stack.RemoveValues(ctx, dn, attr)
}

func (stack *IDMProtocolStack) SetCreateTimestamp(ctx context.Context, dn DN, ts time.Time) (resp X500OpOutcome, result *x500.ModifyEntryResultData, err error) {
	return stack.ModifyEntry(ctx, arg_data)
}

func (stack *IDMProtocolStack) SetModifyTimestamp(ctx context.Context, dn DN, ts time.Time) (resp X500OpOutcome, result *x500.ModifyEntryResultData, err error) {
	return stack.ModifyEntry(ctx, arg_data)
}

func (stack *IDMProtocolStack) SetCreatorsName(ctx context.Context, dn, creator DN) (resp X500OpOutcome, result *x500.ModifyEntryResultData, err error) {
	return stack.ModifyEntry(ctx, arg_data)
}

func (stack *IDMProtocolStack) SetModifiersName(ctx context.Context, dn, modifier DN) (resp X500OpOutcome, result *x500.ModifyEntryResultData, err error) {
	return stack.ModifyEntry(ctx, arg_data)
}

func (stack *IDMProtocolStack) GetDSAinfo(ctx context.Context) (info DSAInfo, err error) {
	return DSAInfo{}, nil
}
