package x500

import (
	"encoding/asn1"
)

// # ASN.1 Definition:
//
//	DSEType ::= BIT STRING {
//	  root           (0),  -- root DSE
//	  glue           (1),  -- represents knowledge of a name only
//	  cp             (2),  -- context prefix
//	  entry          (3),  -- object entry
//	  alias          (4),  -- alias entry
//	  subr           (5),  -- subordinate reference
//	  nssr           (6),  -- non-specific subordinate reference
//	  supr           (7),  -- superior reference
//	  xr             (8),  -- cross reference
//	  admPoint       (9),  -- administrative point
//	  subentry       (10), -- subentry
//	  shadow         (11), -- shadow copy
//	  immSupr        (13), -- immediate superior reference
//	  rhob           (14), -- rhob information
//	  sa             (15), -- subordinate reference to alias entry
//	  dsSubentry     (16), -- DSA Specific subentry
//	  familyMember   (17), -- family member
//	  ditBridge      (18)}
type DSEType = asn1.BitString

const DSEType_Root int = 0

const DSEType_Glue int = 1

const DSEType_Cp int = 2

const DSEType_Entry int = 3

const DSEType_Alias int = 4

const DSEType_Subr int = 5

const DSEType_Nssr int = 6

const DSEType_Supr int = 7

const DSEType_Xr int = 8

const DSEType_AdmPoint int = 9

const DSEType_Subentry int = 10

const DSEType_Shadow int = 11

const DSEType_ImmSupr int = 13

const DSEType_Rhob int = 14

const DSEType_Sa int = 15

const DSEType_DsSubentry int = 16

const DSEType_FamilyMember int = 17

const DSEType_DitBridge int = 18

// # ASN.1 Definition:
//
//	SupplierOrConsumer ::= SET {
//	  COMPONENTS OF              AccessPoint, -- supplier or consumer
//	  agreementID           [3]  OperationalBindingID,
//	  ... }
type SupplierOrConsumer struct {
	Ae_title            Name                    `asn1:"explicit,tag:0"`
	Address             PresentationAddress     `asn1:"explicit,tag:1"`
	ProtocolInformation [](ProtocolInformation) `asn1:"optional,explicit,tag:2,set"`
	AgreementID         OperationalBindingID    `asn1:"explicit,tag:3"`
}

func (ap *SupplierOrConsumer) GetAETitle() Name {
	return ap.Ae_title
}

func (ap *SupplierOrConsumer) GetAddress() PresentationAddress {
	return ap.Address
}

func (ap *SupplierOrConsumer) GetProtocolInformation() []ProtocolInformation {
	return ap.ProtocolInformation
}

// NOTE: Supplier_is_master is represented as an `asn1.RawValue` because the
// boolean defaults to true, but Go only defaults booleans to false, so we
// have to preserve the raw element.
//
// # ASN.1 Definition:
//
//	SupplierInformation ::= SET {
//	  COMPONENTS OF              SupplierOrConsumer, -- supplier
//	  supplier-is-master    [4]  BOOLEAN DEFAULT TRUE,
//	  non-supplying-master  [5]  AccessPoint OPTIONAL,
//	  ... }
type SupplierInformation struct {
	Ae_title             Name                    `asn1:"explicit,tag:0"`
	Address              PresentationAddress     `asn1:"explicit,tag:1"`
	ProtocolInformation  [](ProtocolInformation) `asn1:"optional,explicit,tag:2,set"`
	AgreementID          OperationalBindingID    `asn1:"explicit,tag:3"`
	Supplier_is_master   asn1.RawValue           `asn1:"optional,explicit,tag:4"`
	Non_supplying_master AccessPoint             `asn1:"optional,explicit,tag:5,set"`
}

// # ASN.1 Definition:
//
//	ConsumerInformation ::= SupplierOrConsumer
type ConsumerInformation = SupplierOrConsumer

// # ASN.1 Definition:
//
//	SupplierAndConsumers ::= SET {
//	  COMPONENTS OF         AccessPoint, -- supplier
//	  consumers        [3]  SET OF AccessPoint,
//	  ... }
type SupplierAndConsumers struct {
	Ae_title            Name                    `asn1:"explicit,tag:0"`
	Address             PresentationAddress     `asn1:"explicit,tag:1"`
	ProtocolInformation [](ProtocolInformation) `asn1:"optional,explicit,tag:2"`
	Consumers           [](AccessPoint)         `asn1:"explicit,tag:3,set"`
}

func (ap *SupplierAndConsumers) GetAETitle() Name {
	return ap.Ae_title
}

func (ap *SupplierAndConsumers) GetAddress() PresentationAddress {
	return ap.Address
}

func (ap *SupplierAndConsumers) GetProtocolInformation() []ProtocolInformation {
	return ap.ProtocolInformation
}

// # ASN.1 Definition:
//
//	id-doa-dseType                  OBJECT IDENTIFIER ::= {id-doa 0}
var Id_doa_dseType asn1.ObjectIdentifier = []int{2, 5, 12, 0}

// # ASN.1 Definition:
//
//	id-doa-myAccessPoint            OBJECT IDENTIFIER ::= {id-doa 1}
var Id_doa_myAccessPoint asn1.ObjectIdentifier = []int{2, 5, 12, 1}

// # ASN.1 Definition:
//
//	id-doa-superiorKnowledge        OBJECT IDENTIFIER ::= {id-doa 2}
var Id_doa_superiorKnowledge asn1.ObjectIdentifier = []int{2, 5, 12, 2}

// # ASN.1 Definition:
//
//	id-doa-specificKnowledge        OBJECT IDENTIFIER ::= {id-doa 3}
var Id_doa_specificKnowledge asn1.ObjectIdentifier = []int{2, 5, 12, 3}

// # ASN.1 Definition:
//
//	id-doa-nonSpecificKnowledge     OBJECT IDENTIFIER ::= {id-doa 4}
var Id_doa_nonSpecificKnowledge asn1.ObjectIdentifier = []int{2, 5, 12, 4}

// # ASN.1 Definition:
//
//	id-doa-supplierKnowledge        OBJECT IDENTIFIER ::= {id-doa 5}
var Id_doa_supplierKnowledge asn1.ObjectIdentifier = []int{2, 5, 12, 5}

// # ASN.1 Definition:
//
//	id-doa-consumerKnowledge        OBJECT IDENTIFIER ::= {id-doa 6}
var Id_doa_consumerKnowledge asn1.ObjectIdentifier = []int{2, 5, 12, 6}

// # ASN.1 Definition:
//
//	id-doa-secondaryShadows         OBJECT IDENTIFIER ::= {id-doa 7}
var Id_doa_secondaryShadows asn1.ObjectIdentifier = []int{2, 5, 12, 7}

// # ASN.1 Definition:
//
//	id-doa-ditBridgeKnowledge       OBJECT IDENTIFIER ::= {id-doa 8}
var Id_doa_ditBridgeKnowledge asn1.ObjectIdentifier = []int{2, 5, 12, 8}

// # ASN.1 Definition:
//
//	id-kmr-accessPointMatch         OBJECT IDENTIFIER ::= {id-kmr 0}
var Id_kmr_accessPointMatch asn1.ObjectIdentifier = []int{2, 5, 14, 0}

// # ASN.1 Definition:
//
//	id-kmr-masterShadowMatch        OBJECT IDENTIFIER ::= {id-kmr 1}
var Id_kmr_masterShadowMatch asn1.ObjectIdentifier = []int{2, 5, 14, 1}

// # ASN.1 Definition:
//
//	id-kmr-supplierConsumerMatch    OBJECT IDENTIFIER ::= {id-kmr 2}
var Id_kmr_supplierConsumerMatch asn1.ObjectIdentifier = []int{2, 5, 14, 2}

// # ASN.1 Definition:
//
//	id-kmr-supplierConsumersMatch   OBJECT IDENTIFIER ::= {id-kmr 3}
var Id_kmr_supplierConsumersMatch asn1.ObjectIdentifier = []int{2, 5, 14, 3}

// # ASN.1 Definition:
//
// supplierOrConsumerInformationMatch-AssertionType ::= SEQUENCE { -- REMOVED_FROM_UNNESTING -- }
type SupplierOrConsumerInformationMatch_AssertionType struct {
	Ae_title             Name `asn1:"explicit,tag:0"`
	Agreement_identifier int  `asn1:"explicit,tag:2"`
}
