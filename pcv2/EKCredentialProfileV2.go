package main

import (
	"encoding/asn1"
)

/* START_OF_SYMBOL_DEFINITION Tcg */
// ### ASN.1 Definition:
// 
//   tcg OBJECT IDENTIFIER ::= {joint-iso-itu-t(2) international-organizations(23) tcg(133)}
// 
// 
var Tcg asn1.ObjectIdentifier = []int{ 2, 23, 133 }
/* END_OF_SYMBOL_DEFINITION Tcg *//* START_OF_SYMBOL_DEFINITION Tcg_attribute */
// ### ASN.1 Definition:
// 
//   tcg-attribute                   OBJECT IDENTIFIER ::= {tcg 2}
// 
// 
var Tcg_attribute asn1.ObjectIdentifier = []int{ 2, 23, 133, 2 }
/* END_OF_SYMBOL_DEFINITION Tcg_attribute *//* START_OF_SYMBOL_DEFINITION Tcg_kp */
// ### ASN.1 Definition:
// 
//   tcg-kp                          OBJECT IDENTIFIER ::= {tcg 8}
// 
// 
var Tcg_kp asn1.ObjectIdentifier = []int{ 2, 23, 133, 8 }
/* END_OF_SYMBOL_DEFINITION Tcg_kp *//* START_OF_SYMBOL_DEFINITION Tcg_at_tpmManufacturer */
// ### ASN.1 Definition:
// 
//   tcg-at-tpmManufacturer          OBJECT IDENTIFIER ::= {tcg-attribute 1}
// 
// 
var Tcg_at_tpmManufacturer asn1.ObjectIdentifier = []int{ 2, 23, 133, 2, 1 }
/* END_OF_SYMBOL_DEFINITION Tcg_at_tpmManufacturer *//* START_OF_SYMBOL_DEFINITION Tcg_at_tpmModel */
// ### ASN.1 Definition:
// 
//   tcg-at-tpmModel                 OBJECT IDENTIFIER ::= {tcg-attribute 2}
// 
// 
var Tcg_at_tpmModel asn1.ObjectIdentifier = []int{ 2, 23, 133, 2, 2 }
/* END_OF_SYMBOL_DEFINITION Tcg_at_tpmModel *//* START_OF_SYMBOL_DEFINITION Tcg_at_tpmVersion */
// ### ASN.1 Definition:
// 
//   tcg-at-tpmVersion               OBJECT IDENTIFIER ::= {tcg-attribute 3}
// 
// 
var Tcg_at_tpmVersion asn1.ObjectIdentifier = []int{ 2, 23, 133, 2, 3 }
/* END_OF_SYMBOL_DEFINITION Tcg_at_tpmVersion *//* START_OF_SYMBOL_DEFINITION Tcg_at_tpmSpecification */
// ### ASN.1 Definition:
// 
//   tcg-at-tpmSpecification         OBJECT IDENTIFIER ::= {tcg-attribute 16}
// 
// 
var Tcg_at_tpmSpecification asn1.ObjectIdentifier = []int{ 2, 23, 133, 2, 16 }
/* END_OF_SYMBOL_DEFINITION Tcg_at_tpmSpecification *//* START_OF_SYMBOL_DEFINITION Tcg_at_tpmSecurityAssertions */
// ### ASN.1 Definition:
// 
//   tcg-at-tpmSecurityAssertions    OBJECT IDENTIFIER ::= {tcg-attribute 18}
// 
// 
var Tcg_at_tpmSecurityAssertions asn1.ObjectIdentifier = []int{ 2, 23, 133, 2, 18 }
/* END_OF_SYMBOL_DEFINITION Tcg_at_tpmSecurityAssertions *//* START_OF_SYMBOL_DEFINITION Tcg_kp_EKCertificate */
// ### ASN.1 Definition:
// 
//   tcg-kp-EKCertificate            OBJECT IDENTIFIER ::= {tcg-kp 1}
// 
// 
var Tcg_kp_EKCertificate asn1.ObjectIdentifier = []int{ 2, 23, 133, 8, 1 }
/* END_OF_SYMBOL_DEFINITION Tcg_kp_EKCertificate *//* START_OF_SYMBOL_DEFINITION Strmax */
// ### ASN.1 Definition:
// 
//   strmax                          INTEGER ::= 256
// 
// 
const Strmax int = 256
/* END_OF_SYMBOL_DEFINITION Strmax *//* START_OF_SYMBOL_DEFINITION Urimax */
// ### ASN.1 Definition:
// 
//   urimax                          INTEGER ::= 1024
// 
// 
const Urimax int = 1024
/* END_OF_SYMBOL_DEFINITION Urimax *//* START_OF_SYMBOL_DEFINITION Certstrmax */
// ### ASN.1 Definition:
// 
//   certstrmax                      INTEGER ::= 100000
// 
// 
const Certstrmax int = 100000
/* END_OF_SYMBOL_DEFINITION Certstrmax *//* START_OF_SYMBOL_DEFINITION Version */
// ### ASN.1 Definition:
// 
//   Version  ::=  INTEGER { v1(0) }
// 
type Version = int
/* END_OF_SYMBOL_DEFINITION Version */

/* START_OF_SYMBOL_DEFINITION Version_V1 */
const Version_V1 Version = 0
/* END_OF_SYMBOL_DEFINITION Version_V1 *//* START_OF_SYMBOL_DEFINITION TPMSecurityAssertions */
// ### ASN.1 Definition:
// 
//   TPMSecurityAssertions ::= SEQUENCE {
//     version                         Version DEFAULT v1,
//     fieldUpgradable                 BOOLEAN DEFAULT FALSE,
//     ekGenerationType                [0] IMPLICIT EKGenerationType OPTIONAL,
//     ekGenerationLocation            [1] IMPLICIT EKGenerationLocation OPTIONAL,
//     ekCertificateGenerationLocation [2] IMPLICIT EKCertificateGenerationLocation OPTIONAL,
//     ccInfo                          [3] IMPLICIT CommonCriteriaMeasures OPTIONAL,
//     fipsLevel                       [4] IMPLICIT FIPSLevel OPTIONAL,
//     iso9000Certified                [5] IMPLICIT BOOLEAN DEFAULT FALSE,
//     iso9000Uri                      IA5STRING (SIZE (1..urimax)) OPTIONAL -- // FIXME: Missing last paren
// }
// 
// 
type TPMSecurityAssertions struct {
    Version Version `asn1:"optional"`
    FieldUpgradable bool `asn1:"optional"`
    EkGenerationType EKGenerationType `asn1:"optional,tag:0"`
    EkGenerationLocation EKGenerationLocation `asn1:"optional,tag:1"`
    EkCertificateGenerationLocation EKCertificateGenerationLocation `asn1:"optional,tag:2"`
    CcInfo CommonCriteriaMeasures `asn1:"optional,tag:3"`
    FipsLevel FIPSLevel `asn1:"optional,tag:4"`
    Iso9000Certified bool `asn1:"optional,tag:5"`
    Iso9000Uri IA5STRING `asn1:"optional"`
}
/* END_OF_SYMBOL_DEFINITION TPMSecurityAssertions *//* START_OF_SYMBOL_DEFINITION EKGenerationType */
// ### ASN.1 Definition:
// 
// EKGenerationType  ::=  ENUMERATED {
//     internal (0),
//     injected (1),
//     internalRevocable(2),
//     injectedRevocable(3),
//     firmwareLimited(4)
// }
// 
type EKGenerationType = asn1.Enumerated
const (
    EKGenerationType_Internal EKGenerationType = 0,
    EKGenerationType_Injected EKGenerationType = 1,
    EKGenerationType_InternalRevocable EKGenerationType = 2,
    EKGenerationType_InjectedRevocable EKGenerationType = 3,
    EKGenerationType_FirmwareLimited EKGenerationType = 4
)
/* END_OF_SYMBOL_DEFINITION EKGenerationType *//* START_OF_SYMBOL_DEFINITION EKGenerationLocation */
// ### ASN.1 Definition:
// 
// EKGenerationLocation  ::=  ENUMERATED {
//     tpmManufacturer (0),
//     platformManufacturer (1),
//     ekCertSigner (2)
// }
// 
type EKGenerationLocation = asn1.Enumerated
const (
    EKGenerationLocation_TpmManufacturer EKGenerationLocation = 0,
    EKGenerationLocation_PlatformManufacturer EKGenerationLocation = 1,
    EKGenerationLocation_EkCertSigner EKGenerationLocation = 2
)
/* END_OF_SYMBOL_DEFINITION EKGenerationLocation *//* START_OF_SYMBOL_DEFINITION EKCertificateGenerationLocation */
// ### ASN.1 Definition:
// 
// EKCertificateGenerationLocation  ::=  ENUMERATED {
//     tpmManufacturer (0),
//     platformManufacturer (1),
//     ekCertSigner (2)
// }
// 
type EKCertificateGenerationLocation = asn1.Enumerated
const (
    EKCertificateGenerationLocation_TpmManufacturer EKCertificateGenerationLocation = 0,
    EKCertificateGenerationLocation_PlatformManufacturer EKCertificateGenerationLocation = 1,
    EKCertificateGenerationLocation_EkCertSigner EKCertificateGenerationLocation = 2
)
/* END_OF_SYMBOL_DEFINITION EKCertificateGenerationLocation *//* START_OF_SYMBOL_DEFINITION CommonCriteriaMeasures */
// ### ASN.1 Definition:
// 
//   CommonCriteriaMeasures ::= SEQUENCE {
//     version                 IA5STRING (SIZE (1..strmax)), -- “2.2” or “3.1”; future syntax defined by CC
//     assurancelevel          EvaluationAssuranceLevel,
//     evaluationStatus        EvalutionStatus,
//     plus                    BOOLEAN DEFAULT FALSE,
//     strengthOfFunction      [0] IMPLICIT StrengthOfFunction OPTIONAL,
//     profileOid              [1] IMPLICIT OBJECT IDENTIFIER OPTIONAL,
//     profileUri              [2] IMPLICIT URIReference OPTIONAL,
//     targetOid               [3] IMPLICIT OBJECT IDENTIFIER OPTIONAL,
//     targetUri               [4] IMPLICIT URIReference OPTIONAL
// }
// 
// 
type CommonCriteriaMeasures struct {
    Version IA5STRING
    Assurancelevel EvaluationAssuranceLevel
    EvaluationStatus EvalutionStatus
    Plus bool `asn1:"optional"`
    StrengthOfFunction StrengthOfFunction `asn1:"optional,tag:0"`
    ProfileOid asn1.ObjectIdentifier `asn1:"optional,tag:1"`
    ProfileUri URIReference `asn1:"optional,tag:2"`
    TargetOid asn1.ObjectIdentifier `asn1:"optional,tag:3"`
    TargetUri URIReference `asn1:"optional,tag:4"`
}
/* END_OF_SYMBOL_DEFINITION CommonCriteriaMeasures *//* START_OF_SYMBOL_DEFINITION EvaluationAssuranceLevel */
// ### ASN.1 Definition:
// 
// EvaluationAssuranceLevel  ::=  ENUMERATED {
//     levell (1),
//     level2 (2),
//     level3 (3),
//     level4 (4),
//     level5 (5),
//     level6 (6),
//     level7 (7)
// }
// 
type EvaluationAssuranceLevel = asn1.Enumerated
const (
    EvaluationAssuranceLevel_Levell EvaluationAssuranceLevel = 1,
    EvaluationAssuranceLevel_Level2 EvaluationAssuranceLevel = 2,
    EvaluationAssuranceLevel_Level3 EvaluationAssuranceLevel = 3,
    EvaluationAssuranceLevel_Level4 EvaluationAssuranceLevel = 4,
    EvaluationAssuranceLevel_Level5 EvaluationAssuranceLevel = 5,
    EvaluationAssuranceLevel_Level6 EvaluationAssuranceLevel = 6,
    EvaluationAssuranceLevel_Level7 EvaluationAssuranceLevel = 7
)
/* END_OF_SYMBOL_DEFINITION EvaluationAssuranceLevel *//* START_OF_SYMBOL_DEFINITION StrengthOfFunction */
// ### ASN.1 Definition:
// 
// StrengthOfFunction  ::=  ENUMERATED {
//     basic (0),
//     medium (1),
//     high (2)
// }
// 
type StrengthOfFunction = asn1.Enumerated
const (
    StrengthOfFunction_Basic StrengthOfFunction = 0,
    StrengthOfFunction_Medium StrengthOfFunction = 1,
    StrengthOfFunction_High StrengthOfFunction = 2
)
/* END_OF_SYMBOL_DEFINITION StrengthOfFunction *//* START_OF_SYMBOL_DEFINITION URIReference */
// ### ASN.1 Definition:
// 
//   URIReference ::= SEQUENCE {
//     uniformResourceIdentifier   IA5String (SIZE (1..urimax)), -- // FIXME: Missing last paren
//     hashAlgorithm               AlgorithmIdentifier{{SupportedAlgorithms}} OPTIONAL,
//     hashValue                   BIT STRING OPTIONAL
// }
// 
// 
type URIReference struct {
    UniformResourceIdentifier string
    HashAlgorithm AlgorithmIdentifier `asn1:"optional"`
    HashValue asn1.BitString `asn1:"optional"`
}
/* END_OF_SYMBOL_DEFINITION URIReference *//* START_OF_SYMBOL_DEFINITION EvaluationStatus */
// ### ASN.1 Definition:
// 
// EvaluationStatus  ::=  ENUMERATED {
//     designedToMeet (0),
//     evaluationInProgress (1),
//     evaluationCompleted (2)
// }
// 
type EvaluationStatus = asn1.Enumerated
const (
    EvaluationStatus_DesignedToMeet EvaluationStatus = 0,
    EvaluationStatus_EvaluationInProgress EvaluationStatus = 1,
    EvaluationStatus_EvaluationCompleted EvaluationStatus = 2
)
/* END_OF_SYMBOL_DEFINITION EvaluationStatus *//* START_OF_SYMBOL_DEFINITION FIPSLevel */
// ### ASN.1 Definition:
// 
//   FIPSLevel ::= SEQUENCE {
//     version     IA5STRING (SIZE (1..strmax)), -- “140-1” or “140-2”
//     level       SecurityLevel,
//     plus        BOOLEAN DEFAULT FALSE
// }
// 
// 
type FIPSLevel struct {
    Version IA5STRING
    Level SecurityLevel
    Plus bool `asn1:"optional"`
}
/* END_OF_SYMBOL_DEFINITION FIPSLevel *//* START_OF_SYMBOL_DEFINITION SecurityLevel */
// ### ASN.1 Definition:
// 
// SecurityLevel  ::=  ENUMERATED {
//     level1 (1),
//     level2 (2),
//     level3 (3),
//     level4 (4)
// }
// 
type SecurityLevel = asn1.Enumerated
const (
    SecurityLevel_Level1 SecurityLevel = 1,
    SecurityLevel_Level2 SecurityLevel = 2,
    SecurityLevel_Level3 SecurityLevel = 3,
    SecurityLevel_Level4 SecurityLevel = 4
)
/* END_OF_SYMBOL_DEFINITION SecurityLevel *//* START_OF_SYMBOL_DEFINITION TPMSpecification */
// ### ASN.1 Definition:
// 
//   TPMSpecification ::= SEQUENCE {
//     family      UTF8String (SIZE (1..strmax)),
//     level       INTEGER,
//     revision    INTEGER
// }
// 
// 
type TPMSpecification struct {
    Family string
    Level int
    Revision int
}
/* END_OF_SYMBOL_DEFINITION TPMSpecification */