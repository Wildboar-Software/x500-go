package main

import (
	"encoding/asn1"
	"time"
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
//   tcg-attribute                                   OBJECT IDENTIFIER ::= {tcg 2}
// 
// 
var Tcg_attribute asn1.ObjectIdentifier = []int{ 2, 23, 133, 2 }
/* END_OF_SYMBOL_DEFINITION Tcg_attribute *//* START_OF_SYMBOL_DEFINITION Tcg_platformClass */
// ### ASN.1 Definition:
// 
//   tcg-platformClass                               OBJECT IDENTIFIER ::= {tcg 5}
// 
// 
var Tcg_platformClass asn1.ObjectIdentifier = []int{ 2, 23, 133, 5 }
/* END_OF_SYMBOL_DEFINITION Tcg_platformClass *//* START_OF_SYMBOL_DEFINITION Tcg_kp */
// ### ASN.1 Definition:
// 
//   tcg-kp                                          OBJECT IDENTIFIER ::= {tcg 8}
// 
// 
var Tcg_kp asn1.ObjectIdentifier = []int{ 2, 23, 133, 8 }
/* END_OF_SYMBOL_DEFINITION Tcg_kp *//* START_OF_SYMBOL_DEFINITION Tcg_ca */
// ### ASN.1 Definition:
// 
//   tcg-ca                                          OBJECT IDENTIFIER ::= {tcg 11}
// 
// 
var Tcg_ca asn1.ObjectIdentifier = []int{ 2, 23, 133, 11 }
/* END_OF_SYMBOL_DEFINITION Tcg_ca *//* START_OF_SYMBOL_DEFINITION Tcg_address */
// ### ASN.1 Definition:
// 
//   tcg-address                                     OBJECT IDENTIFIER ::= {tcg 17}
// 
// 
var Tcg_address asn1.ObjectIdentifier = []int{ 2, 23, 133, 17 }
/* END_OF_SYMBOL_DEFINITION Tcg_address *//* START_OF_SYMBOL_DEFINITION Tcg_registry */
// ### ASN.1 Definition:
// 
//   tcg-registry                                    OBJECT IDENTIFIER ::= {tcg 18}
// 
// 
var Tcg_registry asn1.ObjectIdentifier = []int{ 2, 23, 133, 18 }
/* END_OF_SYMBOL_DEFINITION Tcg_registry *//* START_OF_SYMBOL_DEFINITION Tcg_common */
// ### ASN.1 Definition:
// 
//   tcg-common                                      OBJECT IDENTIFIER ::= {tcg-platformClass 1}
// 
// 
var Tcg_common asn1.ObjectIdentifier = []int{ 2, 23, 133, 5, 1 }
/* END_OF_SYMBOL_DEFINITION Tcg_common *//* START_OF_SYMBOL_DEFINITION Tcg_at_platformConfiguration */
// ### ASN.1 Definition:
// 
//   tcg-at-platformConfiguration                    OBJECT IDENTIFIER ::= {tcg-common 7}
// 
// 
var Tcg_at_platformConfiguration asn1.ObjectIdentifier = []int{ 2, 23, 133, 5, 1, 7 }
/* END_OF_SYMBOL_DEFINITION Tcg_at_platformConfiguration *//* START_OF_SYMBOL_DEFINITION Tcg_at_tcgPlatformSpecification */
// ### ASN.1 Definition:
// 
//   tcg-at-tcgPlatformSpecification                 OBJECT IDENTIFIER ::= {tcg-attribute 17}
// 
// 
var Tcg_at_tcgPlatformSpecification asn1.ObjectIdentifier = []int{ 2, 23, 133, 2, 17 }
/* END_OF_SYMBOL_DEFINITION Tcg_at_tcgPlatformSpecification *//* START_OF_SYMBOL_DEFINITION Tcg_at_tcgCredentialSpecification */
// ### ASN.1 Definition:
// 
//   tcg-at-tcgCredentialSpecification               OBJECT IDENTIFIER ::= {tcg-attribute 23}
// 
// 
var Tcg_at_tcgCredentialSpecification asn1.ObjectIdentifier = []int{ 2, 23, 133, 2, 23 }
/* END_OF_SYMBOL_DEFINITION Tcg_at_tcgCredentialSpecification *//* START_OF_SYMBOL_DEFINITION Tcg_at_tcgCredentialType */
// ### ASN.1 Definition:
// 
//   tcg-at-tcgCredentialType                        OBJECT IDENTIFIER ::= {tcg-attribute 25}
// 
// 
var Tcg_at_tcgCredentialType asn1.ObjectIdentifier = []int{ 2, 23, 133, 2, 25 }
/* END_OF_SYMBOL_DEFINITION Tcg_at_tcgCredentialType *//* START_OF_SYMBOL_DEFINITION Tcg_kp_PlatformAttributeCertificate */
// ### ASN.1 Definition:
// 
//   tcg-kp-PlatformAttributeCertificate             OBJECT IDENTIFIER ::= {tcg-kp 2}
// 
// 
var Tcg_kp_PlatformAttributeCertificate asn1.ObjectIdentifier = []int{ 2, 23, 133, 8, 2 }
/* END_OF_SYMBOL_DEFINITION Tcg_kp_PlatformAttributeCertificate *//* START_OF_SYMBOL_DEFINITION Tcg_kp_PlatformKeyCertificate */
// ### ASN.1 Definition:
// 
//   tcg-kp-PlatformKeyCertificate                   OBJECT IDENTIFIER ::= {tcg-kp 4}
// 
// 
var Tcg_kp_PlatformKeyCertificate asn1.ObjectIdentifier = []int{ 2, 23, 133, 8, 4 }
/* END_OF_SYMBOL_DEFINITION Tcg_kp_PlatformKeyCertificate *//* START_OF_SYMBOL_DEFINITION Tcg_kp_DeltaPlatformAttributeCertificate */
// ### ASN.1 Definition:
// 
//   tcg-kp-DeltaPlatformAttributeCertificate        OBJECT IDENTIFIER ::= {tcg-kp 5}
// 
// 
var Tcg_kp_DeltaPlatformAttributeCertificate asn1.ObjectIdentifier = []int{ 2, 23, 133, 8, 5 }
/* END_OF_SYMBOL_DEFINITION Tcg_kp_DeltaPlatformAttributeCertificate *//* START_OF_SYMBOL_DEFINITION Tcg_address_ethernetmac */
// ### ASN.1 Definition:
// 
//   tcg-address-ethernetmac                         OBJECT IDENTIFIER ::= {tcg-address 1}
// 
// 
var Tcg_address_ethernetmac asn1.ObjectIdentifier = []int{ 2, 23, 133, 17, 1 }
/* END_OF_SYMBOL_DEFINITION Tcg_address_ethernetmac *//* START_OF_SYMBOL_DEFINITION Tcg_address_wlanmac */
// ### ASN.1 Definition:
// 
//   tcg-address-wlanmac                             OBJECT IDENTIFIER ::= {tcg-address 2}
// 
// 
var Tcg_address_wlanmac asn1.ObjectIdentifier = []int{ 2, 23, 133, 17, 2 }
/* END_OF_SYMBOL_DEFINITION Tcg_address_wlanmac *//* START_OF_SYMBOL_DEFINITION Tcg_address_bluetoothmac */
// ### ASN.1 Definition:
// 
//   tcg-address-bluetoothmac                        OBJECT IDENTIFIER ::= {tcg-address 3}
// 
// 
var Tcg_address_bluetoothmac asn1.ObjectIdentifier = []int{ 2, 23, 133, 17, 3 }
/* END_OF_SYMBOL_DEFINITION Tcg_address_bluetoothmac *//* START_OF_SYMBOL_DEFINITION Tcg_registry_componentClass */
// ### ASN.1 Definition:
// 
//   tcg-registry-componentClass                     OBJECT IDENTIFIER ::= {tcg-registry 3}
// 
// 
var Tcg_registry_componentClass asn1.ObjectIdentifier = []int{ 2, 23, 133, 18, 3 }
/* END_OF_SYMBOL_DEFINITION Tcg_registry_componentClass *//* START_OF_SYMBOL_DEFINITION Tcg_registry_componentClass_pcie */
// ### ASN.1 Definition:
// 
//   tcg-registry-componentClass-pcie                OBJECT IDENTIFIER ::= {tcg-registry-componentClass 4}
// 
// 
var Tcg_registry_componentClass_pcie asn1.ObjectIdentifier = []int{ 2, 23, 133, 18, 3, 4 }
/* END_OF_SYMBOL_DEFINITION Tcg_registry_componentClass_pcie *//* START_OF_SYMBOL_DEFINITION Tcg_registry_componentClass_disk */
// ### ASN.1 Definition:
// 
//   tcg-registry-componentClass-disk                OBJECT IDENTIFIER ::= {tcg-registry-componentClass 5}
// 
// 
var Tcg_registry_componentClass_disk asn1.ObjectIdentifier = []int{ 2, 23, 133, 18, 3, 5 }
/* END_OF_SYMBOL_DEFINITION Tcg_registry_componentClass_disk *//* START_OF_SYMBOL_DEFINITION Tcg_traits */
// ### ASN.1 Definition:
// 
//   tcg-traits                                      OBJECT IDENTIFIER ::= {tcg 19}
// 
// 
var Tcg_traits asn1.ObjectIdentifier = []int{ 2, 23, 133, 19 }
/* END_OF_SYMBOL_DEFINITION Tcg_traits *//* START_OF_SYMBOL_DEFINITION Tcg_at_platformIdentifier */
// ### ASN.1 Definition:
// 
//   tcg-at-platformIdentifier                       OBJECT IDENTIFIER ::= {tcg-common 8}
// 
// 
var Tcg_at_platformIdentifier asn1.ObjectIdentifier = []int{ 2, 23, 133, 5, 1, 8 }
/* END_OF_SYMBOL_DEFINITION Tcg_at_platformIdentifier *//* START_OF_SYMBOL_DEFINITION Tcg_at_platformConfiguration_v3 */
// ### ASN.1 Definition:
// 
//   tcg-at-platformConfiguration-v3                 OBJECT IDENTIFIER ::= {tcg-at-platformConfiguration 3}
// 
// 
var Tcg_at_platformConfiguration_v3 asn1.ObjectIdentifier = []int{ 2, 23, 133, 5, 1, 7, 3 }
/* END_OF_SYMBOL_DEFINITION Tcg_at_platformConfiguration_v3 *//* START_OF_SYMBOL_DEFINITION Tcg_at_platformConfigUri_v3 */
// ### ASN.1 Definition:
// 
//   tcg-at-platformConfigUri-v3                     OBJECT IDENTIFIER ::= {tcg-at-platformConfiguration 4}
// 
// 
var Tcg_at_platformConfigUri_v3 asn1.ObjectIdentifier = []int{ 2, 23, 133, 5, 1, 7, 4 }
/* END_OF_SYMBOL_DEFINITION Tcg_at_platformConfigUri_v3 *//* START_OF_SYMBOL_DEFINITION Tcg_at_previousPlatformCertificates */
// ### ASN.1 Definition:
// 
//   tcg-at-previousPlatformCertificates             OBJECT IDENTIFIER ::= {tcg-attribute 26}
// 
// 
var Tcg_at_previousPlatformCertificates asn1.ObjectIdentifier = []int{ 2, 23, 133, 2, 26 }
/* END_OF_SYMBOL_DEFINITION Tcg_at_previousPlatformCertificates *//* START_OF_SYMBOL_DEFINITION Tcg_at_tbbSecurityAssertions_v3 */
// ### ASN.1 Definition:
// 
//   tcg-at-tbbSecurityAssertions-v3                 OBJECT IDENTIFIER ::= {tcg-attribute 27}
// 
// 
var Tcg_at_tbbSecurityAssertions_v3 asn1.ObjectIdentifier = []int{ 2, 23, 133, 2, 27 }
/* END_OF_SYMBOL_DEFINITION Tcg_at_tbbSecurityAssertions_v3 *//* START_OF_SYMBOL_DEFINITION Tcg_at_cryptographicAnchors */
// ### ASN.1 Definition:
// 
//   tcg-at-cryptographicAnchors                     OBJECT IDENTIFIER ::= {tcg-attribute 28}
// 
// 
var Tcg_at_cryptographicAnchors asn1.ObjectIdentifier = []int{ 2, 23, 133, 2, 28 }
/* END_OF_SYMBOL_DEFINITION Tcg_at_cryptographicAnchors *//* START_OF_SYMBOL_DEFINITION Tcg_kp_DeltaPlatformKeyCertificate */
// ### ASN.1 Definition:
// 
//   tcg-kp-DeltaPlatformKeyCertificate              OBJECT IDENTIFIER ::= {tcg-kp 6}
// 
// 
var Tcg_kp_DeltaPlatformKeyCertificate asn1.ObjectIdentifier = []int{ 2, 23, 133, 8, 6 }
/* END_OF_SYMBOL_DEFINITION Tcg_kp_DeltaPlatformKeyCertificate *//* START_OF_SYMBOL_DEFINITION Tcg_kp_AdditionalPlatformAttributeCertificate */
// ### ASN.1 Definition:
// 
//   tcg-kp-AdditionalPlatformAttributeCertificate   OBJECT IDENTIFIER ::= {tcg-kp 7}
// 
// 
var Tcg_kp_AdditionalPlatformAttributeCertificate asn1.ObjectIdentifier = []int{ 2, 23, 133, 8, 7 }
/* END_OF_SYMBOL_DEFINITION Tcg_kp_AdditionalPlatformAttributeCertificate *//* START_OF_SYMBOL_DEFINITION Tcg_kp_AdditionalPlatformKeyCertificate */
// ### ASN.1 Definition:
// 
//   tcg-kp-AdditionalPlatformKeyCertificate         OBJECT IDENTIFIER ::= {tcg-kp 8}
// 
// 
var Tcg_kp_AdditionalPlatformKeyCertificate asn1.ObjectIdentifier = []int{ 2, 23, 133, 8, 8 }
/* END_OF_SYMBOL_DEFINITION Tcg_kp_AdditionalPlatformKeyCertificate *//* START_OF_SYMBOL_DEFINITION Tcg_registry_componentClass_tcg */
// ### ASN.1 Definition:
// 
//   tcg-registry-componentClass-tcg                 OBJECT IDENTIFIER ::= {tcg-registry-componentClass 1}
// 
// 
var Tcg_registry_componentClass_tcg asn1.ObjectIdentifier = []int{ 2, 23, 133, 18, 3, 1 }
/* END_OF_SYMBOL_DEFINITION Tcg_registry_componentClass_tcg *//* START_OF_SYMBOL_DEFINITION Tcg_registry_componentClass_ietf */
// ### ASN.1 Definition:
// 
//   tcg-registry-componentClass-ietf                OBJECT IDENTIFIER ::= {tcg-registry-componentClass 2}
// 
// 
var Tcg_registry_componentClass_ietf asn1.ObjectIdentifier = []int{ 2, 23, 133, 18, 3, 2 }
/* END_OF_SYMBOL_DEFINITION Tcg_registry_componentClass_ietf *//* START_OF_SYMBOL_DEFINITION Tcg_registry_componentClass_dmtf */
// ### ASN.1 Definition:
// 
//   tcg-registry-componentClass-dmtf                OBJECT IDENTIFIER ::= {tcg-registry-componentClass 3}
// 
// 
var Tcg_registry_componentClass_dmtf asn1.ObjectIdentifier = []int{ 2, 23, 133, 18, 3, 3 }
/* END_OF_SYMBOL_DEFINITION Tcg_registry_componentClass_dmtf *//* START_OF_SYMBOL_DEFINITION Tcg_tr_ID */
// ### ASN.1 Definition:
// 
//   tcg-tr-ID                                       OBJECT IDENTIFIER ::= {tcg-traits 1}
// 
// 
var Tcg_tr_ID asn1.ObjectIdentifier = []int{ 2, 23, 133, 19, 1 }
/* END_OF_SYMBOL_DEFINITION Tcg_tr_ID *//* START_OF_SYMBOL_DEFINITION Tcg_tr_ID_Boolean */
// ### ASN.1 Definition:
// 
//   tcg-tr-ID-Boolean                               OBJECT IDENTIFIER ::= {tcg-tr-ID 1}
// 
// 
var Tcg_tr_ID_Boolean asn1.ObjectIdentifier = []int{ 2, 23, 133, 19, 1, 1 }
/* END_OF_SYMBOL_DEFINITION Tcg_tr_ID_Boolean *//* START_OF_SYMBOL_DEFINITION Tcg_tr_ID_certificateIdentifier */
// ### ASN.1 Definition:
// 
//   tcg-tr-ID-certificateIdentifier                 OBJECT IDENTIFIER ::= {tcg-tr-ID 2}
// 
// 
var Tcg_tr_ID_certificateIdentifier asn1.ObjectIdentifier = []int{ 2, 23, 133, 19, 1, 2 }
/* END_OF_SYMBOL_DEFINITION Tcg_tr_ID_certificateIdentifier *//* START_OF_SYMBOL_DEFINITION Tcg_tr_ID_CommonCriteria */
// ### ASN.1 Definition:
// 
//   tcg-tr-ID-CommonCriteria                        OBJECT IDENTIFIER ::= {tcg-tr-ID 3}
// 
// 
var Tcg_tr_ID_CommonCriteria asn1.ObjectIdentifier = []int{ 2, 23, 133, 19, 1, 3 }
/* END_OF_SYMBOL_DEFINITION Tcg_tr_ID_CommonCriteria *//* START_OF_SYMBOL_DEFINITION Tcg_tr_ID_componentClass */
// ### ASN.1 Definition:
// 
//   tcg-tr-ID-componentClass                        OBJECT IDENTIFIER ::= {tcg-tr-ID 4}
// 
// 
var Tcg_tr_ID_componentClass asn1.ObjectIdentifier = []int{ 2, 23, 133, 19, 1, 4 }
/* END_OF_SYMBOL_DEFINITION Tcg_tr_ID_componentClass *//* START_OF_SYMBOL_DEFINITION Tcg_tr_ID_componentIdentifierV11 */
// ### ASN.1 Definition:
// 
//   tcg-tr-ID-componentIdentifierV11                OBJECT IDENTIFIER ::= {tcg-tr-ID 5}
// 
// 
var Tcg_tr_ID_componentIdentifierV11 asn1.ObjectIdentifier = []int{ 2, 23, 133, 19, 1, 5 }
/* END_OF_SYMBOL_DEFINITION Tcg_tr_ID_componentIdentifierV11 *//* START_OF_SYMBOL_DEFINITION Tcg_tr_ID_FIPSLevel */
// ### ASN.1 Definition:
// 
//   tcg-tr-ID-FIPSLevel                             OBJECT IDENTIFIER ::= {tcg-tr-ID 6}
// 
// 
var Tcg_tr_ID_FIPSLevel asn1.ObjectIdentifier = []int{ 2, 23, 133, 19, 1, 6 }
/* END_OF_SYMBOL_DEFINITION Tcg_tr_ID_FIPSLevel *//* START_OF_SYMBOL_DEFINITION Tcg_tr_ID_ISO9000Level */
// ### ASN.1 Definition:
// 
//   tcg-tr-ID-ISO9000Level                          OBJECT IDENTIFIER ::= {tcg-tr-ID 7}
// 
// 
var Tcg_tr_ID_ISO9000Level asn1.ObjectIdentifier = []int{ 2, 23, 133, 19, 1, 7 }
/* END_OF_SYMBOL_DEFINITION Tcg_tr_ID_ISO9000Level *//* START_OF_SYMBOL_DEFINITION Tcg_tr_ID_networkMAC */
// ### ASN.1 Definition:
// 
//   tcg-tr-ID-networkMAC                            OBJECT IDENTIFIER ::= {tcg-tr-ID 8}
// 
// 
var Tcg_tr_ID_networkMAC asn1.ObjectIdentifier = []int{ 2, 23, 133, 19, 1, 8 }
/* END_OF_SYMBOL_DEFINITION Tcg_tr_ID_networkMAC *//* START_OF_SYMBOL_DEFINITION Tcg_tr_ID_OID */
// ### ASN.1 Definition:
// 
//   tcg-tr-ID-OID                                   OBJECT IDENTIFIER ::= {tcg-tr-ID 9}
// 
// 
var Tcg_tr_ID_OID asn1.ObjectIdentifier = []int{ 2, 23, 133, 19, 1, 9 }
/* END_OF_SYMBOL_DEFINITION Tcg_tr_ID_OID *//* START_OF_SYMBOL_DEFINITION Tcg_tr_ID_PEN */
// ### ASN.1 Definition:
// 
//   tcg-tr-ID-PEN                                   OBJECT IDENTIFIER ::= {tcg-tr-ID 10}
// 
// 
var Tcg_tr_ID_PEN asn1.ObjectIdentifier = []int{ 2, 23, 133, 19, 1, 10 }
/* END_OF_SYMBOL_DEFINITION Tcg_tr_ID_PEN *//* START_OF_SYMBOL_DEFINITION Tcg_tr_ID_platformFirmwareCapabilities */
// ### ASN.1 Definition:
// 
//   tcg-tr-ID-platformFirmwareCapabilities          OBJECT IDENTIFIER ::= {tcg-tr-ID 11}
// 
// 
var Tcg_tr_ID_platformFirmwareCapabilities asn1.ObjectIdentifier = []int{ 2, 23, 133, 19, 1, 11 }
/* END_OF_SYMBOL_DEFINITION Tcg_tr_ID_platformFirmwareCapabilities *//* START_OF_SYMBOL_DEFINITION Tcg_tr_ID_platformFirmwareSignatureVerification */
// ### ASN.1 Definition:
// 
//   tcg-tr-ID-platformFirmwareSignatureVerification OBJECT IDENTIFIER ::= {tcg-tr-ID 12}
// 
// 
var Tcg_tr_ID_platformFirmwareSignatureVerification asn1.ObjectIdentifier = []int{ 2, 23, 133, 19, 1, 12 }
/* END_OF_SYMBOL_DEFINITION Tcg_tr_ID_platformFirmwareSignatureVerification *//* START_OF_SYMBOL_DEFINITION Tcg_tr_ID_platformFirmwareUpdateCompliance */
// ### ASN.1 Definition:
// 
//   tcg-tr-ID-platformFirmwareUpdateCompliance      OBJECT IDENTIFIER ::= {tcg-tr-ID 13}
// 
// 
var Tcg_tr_ID_platformFirmwareUpdateCompliance asn1.ObjectIdentifier = []int{ 2, 23, 133, 19, 1, 13 }
/* END_OF_SYMBOL_DEFINITION Tcg_tr_ID_platformFirmwareUpdateCompliance *//* START_OF_SYMBOL_DEFINITION Tcg_tr_ID_platformHardwareCapabilities */
// ### ASN.1 Definition:
// 
//   tcg-tr-ID-platformHardwareCapabilities          OBJECT IDENTIFIER ::= {tcg-tr-ID 14}
// 
// 
var Tcg_tr_ID_platformHardwareCapabilities asn1.ObjectIdentifier = []int{ 2, 23, 133, 19, 1, 14 }
/* END_OF_SYMBOL_DEFINITION Tcg_tr_ID_platformHardwareCapabilities *//* START_OF_SYMBOL_DEFINITION Tcg_tr_ID_RTM */
// ### ASN.1 Definition:
// 
//   tcg-tr-ID-RTM                                   OBJECT IDENTIFIER ::= {tcg-tr-ID 15}
// 
// 
var Tcg_tr_ID_RTM asn1.ObjectIdentifier = []int{ 2, 23, 133, 19, 1, 15 }
/* END_OF_SYMBOL_DEFINITION Tcg_tr_ID_RTM *//* START_OF_SYMBOL_DEFINITION Tcg_tr_ID_status */
// ### ASN.1 Definition:
// 
//   tcg-tr-ID-status                                OBJECT IDENTIFIER ::= {tcg-tr-ID 16}
// 
// 
var Tcg_tr_ID_status asn1.ObjectIdentifier = []int{ 2, 23, 133, 19, 1, 16 }
/* END_OF_SYMBOL_DEFINITION Tcg_tr_ID_status *//* START_OF_SYMBOL_DEFINITION Tcg_tr_ID_URI */
// ### ASN.1 Definition:
// 
//   tcg-tr-ID-URI                                   OBJECT IDENTIFIER ::= {tcg-tr-ID 17}
// 
// 
var Tcg_tr_ID_URI asn1.ObjectIdentifier = []int{ 2, 23, 133, 19, 1, 17 }
/* END_OF_SYMBOL_DEFINITION Tcg_tr_ID_URI *//* START_OF_SYMBOL_DEFINITION Tcg_tr_ID_UTF8String */
// ### ASN.1 Definition:
// 
//   tcg-tr-ID-UTF8String                            OBJECT IDENTIFIER ::= {tcg-tr-ID 18}
// 
// 
var Tcg_tr_ID_UTF8String asn1.ObjectIdentifier = []int{ 2, 23, 133, 19, 1, 18 }
/* END_OF_SYMBOL_DEFINITION Tcg_tr_ID_UTF8String *//* START_OF_SYMBOL_DEFINITION Tcg_tr_ID_IA5String */
// ### ASN.1 Definition:
// 
//   tcg-tr-ID-IA5String                             OBJECT IDENTIFIER ::= {tcg-tr-ID 19}
// 
// 
var Tcg_tr_ID_IA5String asn1.ObjectIdentifier = []int{ 2, 23, 133, 19, 1, 19 }
/* END_OF_SYMBOL_DEFINITION Tcg_tr_ID_IA5String *//* START_OF_SYMBOL_DEFINITION Tcg_tr_ID_PEMCertString */
// ### ASN.1 Definition:
// 
//   tcg-tr-ID-PEMCertString                         OBJECT IDENTIFIER ::= {tcg-tr-ID 20}
// 
// 
var Tcg_tr_ID_PEMCertString asn1.ObjectIdentifier = []int{ 2, 23, 133, 19, 1, 20 }
/* END_OF_SYMBOL_DEFINITION Tcg_tr_ID_PEMCertString *//* START_OF_SYMBOL_DEFINITION Tcg_tr_ID_PublicKey */
// ### ASN.1 Definition:
// 
//   tcg-tr-ID-PublicKey                             OBJECT IDENTIFIER ::= {tcg-tr-ID 21}
// 
// 
var Tcg_tr_ID_PublicKey asn1.ObjectIdentifier = []int{ 2, 23, 133, 19, 1, 21 }
/* END_OF_SYMBOL_DEFINITION Tcg_tr_ID_PublicKey *//* START_OF_SYMBOL_DEFINITION Tcg_tr_category */
// ### ASN.1 Definition:
// 
//   tcg-tr-category                                     OBJECT IDENTIFIER ::= {tcg-traits 2}
// 
// 
var Tcg_tr_category asn1.ObjectIdentifier = []int{ 2, 23, 133, 19, 2 }
/* END_OF_SYMBOL_DEFINITION Tcg_tr_category *//* START_OF_SYMBOL_DEFINITION Tcg_tr_cat_platformManufacturer */
// ### ASN.1 Definition:
// 
//   tcg-tr-cat-platformManufacturer                     OBJECT IDENTIFIER ::= {tcg-tr-category 1}
// 
// 
var Tcg_tr_cat_platformManufacturer asn1.ObjectIdentifier = []int{ 2, 23, 133, 19, 2, 1 }
/* END_OF_SYMBOL_DEFINITION Tcg_tr_cat_platformManufacturer *//* START_OF_SYMBOL_DEFINITION Tcg_tr_cat_platformModel */
// ### ASN.1 Definition:
// 
//   tcg-tr-cat-platformModel                            OBJECT IDENTIFIER ::= {tcg-tr-category 2}
// 
// 
var Tcg_tr_cat_platformModel asn1.ObjectIdentifier = []int{ 2, 23, 133, 19, 2, 2 }
/* END_OF_SYMBOL_DEFINITION Tcg_tr_cat_platformModel *//* START_OF_SYMBOL_DEFINITION Tcg_tr_cat_platformVersion */
// ### ASN.1 Definition:
// 
//   tcg-tr-cat-platformVersion                          OBJECT IDENTIFIER ::= {tcg-tr-category 3}
// 
// 
var Tcg_tr_cat_platformVersion asn1.ObjectIdentifier = []int{ 2, 23, 133, 19, 2, 3 }
/* END_OF_SYMBOL_DEFINITION Tcg_tr_cat_platformVersion *//* START_OF_SYMBOL_DEFINITION Tcg_tr_cat_platformSerial */
// ### ASN.1 Definition:
// 
//   tcg-tr-cat-platformSerial                           OBJECT IDENTIFIER ::= {tcg-tr-category 4}
// 
// 
var Tcg_tr_cat_platformSerial asn1.ObjectIdentifier = []int{ 2, 23, 133, 19, 2, 4 }
/* END_OF_SYMBOL_DEFINITION Tcg_tr_cat_platformSerial *//* START_OF_SYMBOL_DEFINITION Tcg_tr_cat_platformManufacturerIdentifier */
// ### ASN.1 Definition:
// 
//   tcg-tr-cat-platformManufacturerIdentifier           OBJECT IDENTIFIER ::= {tcg-tr-category 5}
// 
// 
var Tcg_tr_cat_platformManufacturerIdentifier asn1.ObjectIdentifier = []int{ 2, 23, 133, 19, 2, 5 }
/* END_OF_SYMBOL_DEFINITION Tcg_tr_cat_platformManufacturerIdentifier *//* START_OF_SYMBOL_DEFINITION Tcg_tr_cat_platformOwnership */
// ### ASN.1 Definition:
// 
//   tcg-tr-cat-platformOwnership                        OBJECT IDENTIFIER ::= {tcg-tr-category 6}
// 
// 
var Tcg_tr_cat_platformOwnership asn1.ObjectIdentifier = []int{ 2, 23, 133, 19, 2, 6 }
/* END_OF_SYMBOL_DEFINITION Tcg_tr_cat_platformOwnership *//* START_OF_SYMBOL_DEFINITION Tcg_tr_cat_componentClass */
// ### ASN.1 Definition:
// 
//   tcg-tr-cat-componentClass                           OBJECT IDENTIFIER ::= {tcg-tr-category 7}
// 
// 
var Tcg_tr_cat_componentClass asn1.ObjectIdentifier = []int{ 2, 23, 133, 19, 2, 7 }
/* END_OF_SYMBOL_DEFINITION Tcg_tr_cat_componentClass *//* START_OF_SYMBOL_DEFINITION Tcg_tr_cat_componentManufacturer */
// ### ASN.1 Definition:
// 
//   tcg-tr-cat-componentManufacturer                    OBJECT IDENTIFIER ::= {tcg-tr-category 8}
// 
// 
var Tcg_tr_cat_componentManufacturer asn1.ObjectIdentifier = []int{ 2, 23, 133, 19, 2, 8 }
/* END_OF_SYMBOL_DEFINITION Tcg_tr_cat_componentManufacturer *//* START_OF_SYMBOL_DEFINITION Tcg_tr_cat_componentModel */
// ### ASN.1 Definition:
// 
//   tcg-tr-cat-componentModel                           OBJECT IDENTIFIER ::= {tcg-tr-category 9}
// 
// 
var Tcg_tr_cat_componentModel asn1.ObjectIdentifier = []int{ 2, 23, 133, 19, 2, 9 }
/* END_OF_SYMBOL_DEFINITION Tcg_tr_cat_componentModel *//* START_OF_SYMBOL_DEFINITION Tcg_tr_cat_componentSerial */
// ### ASN.1 Definition:
// 
//   tcg-tr-cat-componentSerial                          OBJECT IDENTIFIER ::= {tcg-tr-category 10}
// 
// 
var Tcg_tr_cat_componentSerial asn1.ObjectIdentifier = []int{ 2, 23, 133, 19, 2, 10 }
/* END_OF_SYMBOL_DEFINITION Tcg_tr_cat_componentSerial *//* START_OF_SYMBOL_DEFINITION Tcg_tr_cat_componentStatus */
// ### ASN.1 Definition:
// 
//   tcg-tr-cat-componentStatus                          OBJECT IDENTIFIER ::= {tcg-tr-category 11}
// 
// 
var Tcg_tr_cat_componentStatus asn1.ObjectIdentifier = []int{ 2, 23, 133, 19, 2, 11 }
/* END_OF_SYMBOL_DEFINITION Tcg_tr_cat_componentStatus *//* START_OF_SYMBOL_DEFINITION Tcg_tr_cat_componentLocation */
// ### ASN.1 Definition:
// 
//   tcg-tr-cat-componentLocation                        OBJECT IDENTIFIER ::= {tcg-tr-category 12}
// 
// 
var Tcg_tr_cat_componentLocation asn1.ObjectIdentifier = []int{ 2, 23, 133, 19, 2, 12 }
/* END_OF_SYMBOL_DEFINITION Tcg_tr_cat_componentLocation *//* START_OF_SYMBOL_DEFINITION Tcg_tr_cat_componentRevision */
// ### ASN.1 Definition:
// 
//   tcg-tr-cat-componentRevision                        OBJECT IDENTIFIER ::= {tcg-tr-category 13}
// 
// 
var Tcg_tr_cat_componentRevision asn1.ObjectIdentifier = []int{ 2, 23, 133, 19, 2, 13 }
/* END_OF_SYMBOL_DEFINITION Tcg_tr_cat_componentRevision *//* START_OF_SYMBOL_DEFINITION Tcg_tr_cat_componentFieldReplaceable */
// ### ASN.1 Definition:
// 
//   tcg-tr-cat-componentFieldReplaceable                OBJECT IDENTIFIER ::= {tcg-tr-category 14}
// 
// 
var Tcg_tr_cat_componentFieldReplaceable asn1.ObjectIdentifier = []int{ 2, 23, 133, 19, 2, 14 }
/* END_OF_SYMBOL_DEFINITION Tcg_tr_cat_componentFieldReplaceable *//* START_OF_SYMBOL_DEFINITION Tcg_tr_cat_EKCertificate */
// ### ASN.1 Definition:
// 
//   tcg-tr-cat-EKCertificate                            OBJECT IDENTIFIER ::= {tcg-tr-category 15}
// 
// 
var Tcg_tr_cat_EKCertificate asn1.ObjectIdentifier = []int{ 2, 23, 133, 19, 2, 15 }
/* END_OF_SYMBOL_DEFINITION Tcg_tr_cat_EKCertificate *//* START_OF_SYMBOL_DEFINITION Tcg_tr_cat_IAKCertificate */
// ### ASN.1 Definition:
// 
//   tcg-tr-cat-IAKCertificate                           OBJECT IDENTIFIER ::= {tcg-tr-category 16}
// 
// 
var Tcg_tr_cat_IAKCertificate asn1.ObjectIdentifier = []int{ 2, 23, 133, 19, 2, 16 }
/* END_OF_SYMBOL_DEFINITION Tcg_tr_cat_IAKCertificate *//* START_OF_SYMBOL_DEFINITION Tcg_tr_cat_IDevIDCertificate */
// ### ASN.1 Definition:
// 
//   tcg-tr-cat-IDevIDCertificate                        OBJECT IDENTIFIER ::= {tcg-tr-category 17}
// 
// 
var Tcg_tr_cat_IDevIDCertificate asn1.ObjectIdentifier = []int{ 2, 23, 133, 19, 2, 17 }
/* END_OF_SYMBOL_DEFINITION Tcg_tr_cat_IDevIDCertificate *//* START_OF_SYMBOL_DEFINITION Tcg_tr_cat_DICECertificate */
// ### ASN.1 Definition:
// 
//   tcg-tr-cat-DICECertificate                          OBJECT IDENTIFIER ::= {tcg-tr-category 18}
// 
// 
var Tcg_tr_cat_DICECertificate asn1.ObjectIdentifier = []int{ 2, 23, 133, 19, 2, 18 }
/* END_OF_SYMBOL_DEFINITION Tcg_tr_cat_DICECertificate *//* START_OF_SYMBOL_DEFINITION Tcg_tr_cat_SPDMCertificate */
// ### ASN.1 Definition:
// 
//   tcg-tr-cat-SPDMCertificate                          OBJECT IDENTIFIER ::= {tcg-tr-category 19}
// 
// 
var Tcg_tr_cat_SPDMCertificate asn1.ObjectIdentifier = []int{ 2, 23, 133, 19, 2, 19 }
/* END_OF_SYMBOL_DEFINITION Tcg_tr_cat_SPDMCertificate *//* START_OF_SYMBOL_DEFINITION Tcg_tr_cat_PEMCertificate */
// ### ASN.1 Definition:
// 
//   tcg-tr-cat-PEMCertificate                           OBJECT IDENTIFIER ::= {tcg-tr-category 20}
// 
// 
var Tcg_tr_cat_PEMCertificate asn1.ObjectIdentifier = []int{ 2, 23, 133, 19, 2, 20 }
/* END_OF_SYMBOL_DEFINITION Tcg_tr_cat_PEMCertificate *//* START_OF_SYMBOL_DEFINITION Tcg_tr_cat_PlatformCertificate */
// ### ASN.1 Definition:
// 
//   tcg-tr-cat-PlatformCertificate                      OBJECT IDENTIFIER ::= {tcg-tr-category 21}
// 
// 
var Tcg_tr_cat_PlatformCertificate asn1.ObjectIdentifier = []int{ 2, 23, 133, 19, 2, 21 }
/* END_OF_SYMBOL_DEFINITION Tcg_tr_cat_PlatformCertificate *//* START_OF_SYMBOL_DEFINITION Tcg_tr_cat_DeltaPlatformCertificate */
// ### ASN.1 Definition:
// 
//   tcg-tr-cat-DeltaPlatformCertificate                 OBJECT IDENTIFIER ::= {tcg-tr-category 22}
// 
// 
var Tcg_tr_cat_DeltaPlatformCertificate asn1.ObjectIdentifier = []int{ 2, 23, 133, 19, 2, 22 }
/* END_OF_SYMBOL_DEFINITION Tcg_tr_cat_DeltaPlatformCertificate *//* START_OF_SYMBOL_DEFINITION Tcg_tr_cat_RebasePlatformCertificate */
// ### ASN.1 Definition:
// 
//   tcg-tr-cat-RebasePlatformCertificate                OBJECT IDENTIFIER ::= {tcg-tr-category 23}
// 
// 
var Tcg_tr_cat_RebasePlatformCertificate asn1.ObjectIdentifier = []int{ 2, 23, 133, 19, 2, 23 }
/* END_OF_SYMBOL_DEFINITION Tcg_tr_cat_RebasePlatformCertificate *//* START_OF_SYMBOL_DEFINITION Tcg_tr_cat_genericCertificate */
// ### ASN.1 Definition:
// 
//   tcg-tr-cat-genericCertificate                       OBJECT IDENTIFIER ::= {tcg-tr-category 24}
// 
// 
var Tcg_tr_cat_genericCertificate asn1.ObjectIdentifier = []int{ 2, 23, 133, 19, 2, 24 }
/* END_OF_SYMBOL_DEFINITION Tcg_tr_cat_genericCertificate *//* START_OF_SYMBOL_DEFINITION Tcg_tr_cat_CommonCriteria */
// ### ASN.1 Definition:
// 
//   tcg-tr-cat-CommonCriteria                           OBJECT IDENTIFIER ::= {tcg-tr-category 25}
// 
// 
var Tcg_tr_cat_CommonCriteria asn1.ObjectIdentifier = []int{ 2, 23, 133, 19, 2, 25 }
/* END_OF_SYMBOL_DEFINITION Tcg_tr_cat_CommonCriteria *//* START_OF_SYMBOL_DEFINITION Tcg_tr_cat_componentIdentifierV11 */
// ### ASN.1 Definition:
// 
//   tcg-tr-cat-componentIdentifierV11                   OBJECT IDENTIFIER ::= {tcg-tr-category 26}
// 
// 
var Tcg_tr_cat_componentIdentifierV11 asn1.ObjectIdentifier = []int{ 2, 23, 133, 19, 2, 26 }
/* END_OF_SYMBOL_DEFINITION Tcg_tr_cat_componentIdentifierV11 *//* START_OF_SYMBOL_DEFINITION Tcg_tr_cat_FIPSLevel */
// ### ASN.1 Definition:
// 
//   tcg-tr-cat-FIPSLevel                                OBJECT IDENTIFIER ::= {tcg-tr-category 27}
// 
// 
var Tcg_tr_cat_FIPSLevel asn1.ObjectIdentifier = []int{ 2, 23, 133, 19, 2, 27 }
/* END_OF_SYMBOL_DEFINITION Tcg_tr_cat_FIPSLevel *//* START_OF_SYMBOL_DEFINITION Tcg_tr_cat_ISO9000 */
// ### ASN.1 Definition:
// 
//   tcg-tr-cat-ISO9000                                  OBJECT IDENTIFIER ::= {tcg-tr-category 28}
// 
// 
var Tcg_tr_cat_ISO9000 asn1.ObjectIdentifier = []int{ 2, 23, 133, 19, 2, 28 }
/* END_OF_SYMBOL_DEFINITION Tcg_tr_cat_ISO9000 *//* START_OF_SYMBOL_DEFINITION Tcg_tr_cat_networkMAC */
// ### ASN.1 Definition:
// 
//   tcg-tr-cat-networkMAC                               OBJECT IDENTIFIER ::= {tcg-tr-category 29}
// 
// 
var Tcg_tr_cat_networkMAC asn1.ObjectIdentifier = []int{ 2, 23, 133, 19, 2, 29 }
/* END_OF_SYMBOL_DEFINITION Tcg_tr_cat_networkMAC *//* START_OF_SYMBOL_DEFINITION Tcg_tr_cat_attestationProtocol */
// ### ASN.1 Definition:
// 
//   tcg-tr-cat-attestationProtocol                      OBJECT IDENTIFIER ::= {tcg-tr-category 30}
// 
// 
var Tcg_tr_cat_attestationProtocol asn1.ObjectIdentifier = []int{ 2, 23, 133, 19, 2, 30 }
/* END_OF_SYMBOL_DEFINITION Tcg_tr_cat_attestationProtocol *//* START_OF_SYMBOL_DEFINITION Tcg_tr_cat_PEN */
// ### ASN.1 Definition:
// 
//   tcg-tr-cat-PEN                                      OBJECT IDENTIFIER ::= {tcg-tr-category 31}
// 
// 
var Tcg_tr_cat_PEN asn1.ObjectIdentifier = []int{ 2, 23, 133, 19, 2, 31 }
/* END_OF_SYMBOL_DEFINITION Tcg_tr_cat_PEN *//* START_OF_SYMBOL_DEFINITION Tcg_tr_cat_platformFirmwareCapabilities */
// ### ASN.1 Definition:
// 
//   tcg-tr-cat-platformFirmwareCapabilities             OBJECT IDENTIFIER ::= {tcg-tr-category 32}
// 
// 
var Tcg_tr_cat_platformFirmwareCapabilities asn1.ObjectIdentifier = []int{ 2, 23, 133, 19, 2, 32 }
/* END_OF_SYMBOL_DEFINITION Tcg_tr_cat_platformFirmwareCapabilities *//* START_OF_SYMBOL_DEFINITION Tcg_tr_cat_platformHardwareCapabilities */
// ### ASN.1 Definition:
// 
//   tcg-tr-cat-platformHardwareCapabilities             OBJECT IDENTIFIER ::= {tcg-tr-category 33}
// 
// 
var Tcg_tr_cat_platformHardwareCapabilities asn1.ObjectIdentifier = []int{ 2, 23, 133, 19, 2, 33 }
/* END_OF_SYMBOL_DEFINITION Tcg_tr_cat_platformHardwareCapabilities *//* START_OF_SYMBOL_DEFINITION Tcg_tr_cat_platformFirmwareSignatureVerification */
// ### ASN.1 Definition:
// 
//   tcg-tr-cat-platformFirmwareSignatureVerification    OBJECT IDENTIFIER ::= {tcg-tr-category 34}
// 
// 
var Tcg_tr_cat_platformFirmwareSignatureVerification asn1.ObjectIdentifier = []int{ 2, 23, 133, 19, 2, 34 }
/* END_OF_SYMBOL_DEFINITION Tcg_tr_cat_platformFirmwareSignatureVerification *//* START_OF_SYMBOL_DEFINITION Tcg_tr_cat_platformFirmwareUpdateCompliance */
// ### ASN.1 Definition:
// 
//   tcg-tr-cat-platformFirmwareUpdateCompliance         OBJECT IDENTIFIER ::= {tcg-tr-category 35}
// 
// 
var Tcg_tr_cat_platformFirmwareUpdateCompliance asn1.ObjectIdentifier = []int{ 2, 23, 133, 19, 2, 35 }
/* END_OF_SYMBOL_DEFINITION Tcg_tr_cat_platformFirmwareUpdateCompliance *//* START_OF_SYMBOL_DEFINITION Tcg_tr_cat_RTM */
// ### ASN.1 Definition:
// 
//   tcg-tr-cat-RTM                                      OBJECT IDENTIFIER ::= {tcg-tr-category 36}
// 
// 
var Tcg_tr_cat_RTM asn1.ObjectIdentifier = []int{ 2, 23, 133, 19, 2, 36 }
/* END_OF_SYMBOL_DEFINITION Tcg_tr_cat_RTM *//* START_OF_SYMBOL_DEFINITION Tcg_tr_cat_PublicKey */
// ### ASN.1 Definition:
// 
//   tcg-tr-cat-PublicKey                                OBJECT IDENTIFIER ::= {tcg-tr-category 37}
// 
// 
var Tcg_tr_cat_PublicKey asn1.ObjectIdentifier = []int{ 2, 23, 133, 19, 2, 37 }
/* END_OF_SYMBOL_DEFINITION Tcg_tr_cat_PublicKey *//* START_OF_SYMBOL_DEFINITION Tcg_tr_registry */
// ### ASN.1 Definition:
// 
//   tcg-tr-registry                                     OBJECT IDENTIFIER ::= {tcg-traits 3}
// 
// 
var Tcg_tr_registry asn1.ObjectIdentifier = []int{ 2, 23, 133, 19, 3 }
/* END_OF_SYMBOL_DEFINITION Tcg_tr_registry *//* START_OF_SYMBOL_DEFINITION Tcg_tr_reg_none */
// ### ASN.1 Definition:
// 
//   tcg-tr-reg-none                                     OBJECT IDENTIFIER ::= {tcg-tr-registry 1}
// 
// 
var Tcg_tr_reg_none asn1.ObjectIdentifier = []int{ 2, 23, 133, 19, 3, 1 }
/* END_OF_SYMBOL_DEFINITION Tcg_tr_reg_none *//* START_OF_SYMBOL_DEFINITION Strmax */
// ### ASN.1 Definition:
// 
//   strmax          INTEGER ::= 256
// 
// 
const Strmax int = 256
/* END_OF_SYMBOL_DEFINITION Strmax *//* START_OF_SYMBOL_DEFINITION Urimax */
// ### ASN.1 Definition:
// 
//   urimax          INTEGER ::= 1024
// 
// 
const Urimax int = 1024
/* END_OF_SYMBOL_DEFINITION Urimax *//* START_OF_SYMBOL_DEFINITION Certstrmax */
// ### ASN.1 Definition:
// 
//   certstrmax      INTEGER ::= 100000
// 
// 
const Certstrmax int = 100000
/* END_OF_SYMBOL_DEFINITION Certstrmax *//* START_OF_SYMBOL_DEFINITION Trait */
// ### ASN.1 Definition:
// 
//   Trait ::= SEQUENCE {
//     traitId         TRAIT.&id({TraitSet}), -- Specifies the traitValue encoding
//     traitCategory   OBJECT IDENTIFIER, -- Identifies the information category contained in traitValue
//     traitRegistry   OBJECT IDENTIFIER, -- Identifies the registry used to match against the traitValue
//     description     [0] IMPLICIT UTF8String (SIZE (1..strmax)) OPTIONAL,
//     descriptionURI  [1] IMPLICIT IA5String (SIZE (1..urimax)) OPTIONAL,
//     traitValue      OCTET STRING (CONTAINING TRAIT.&TraitValueType({TraitSet}{@traitId}) ENCODED BY der)
// }
// 
// 
type Trait struct {
    TraitId asn1.ObjectIdentifier
    TraitCategory asn1.ObjectIdentifier
    TraitRegistry asn1.ObjectIdentifier
    Description string `asn1:"optional,tag:0"`
    DescriptionURI string `asn1:"optional,tag:1"`
    TraitValue []byte
}
/* END_OF_SYMBOL_DEFINITION Trait *//* START_OF_SYMBOL_DEFINITION CredentialType */
// ### ASN.1 Definition:
// 
//   CredentialType  ::=  OBJECT IDENTIFIER (
//     tcg-kp-PlatformAttributeCertificate
//     | tcg-kp-PlatformKeyCertificate
//     | tcg-kp-AdditionalPlatformAttributeCertificate
//     | tcg-kp-AdditionalPlatformKeyCertificate
//     | tcg-kp-DeltaPlatformAttributeCertificate
//     | tcg-kp-DeltaPlatformKeyCertificate
// )
// 
type CredentialType = asn1.ObjectIdentifier // ObjectIdentifierType
/* END_OF_SYMBOL_DEFINITION CredentialType *//* START_OF_SYMBOL_DEFINITION TCGCredentialType */
// ### ASN.1 Definition:
// 
//   TCGCredentialType ::= SEQUENCE {
//     certificateType CredentialType
// }
// 
// 
type TCGCredentialType struct {
    CertificateType CredentialType
}
/* END_OF_SYMBOL_DEFINITION TCGCredentialType *//* START_OF_SYMBOL_DEFINITION PreviousPlatformCertificates */
// ### ASN.1 Definition:
// 
//   PreviousPlatformCertificates  ::=  SEQUENCE (SIZE (1..MAX)) OF Trait
// 
type PreviousPlatformCertificates = [](Trait) // SequenceOfType
/* END_OF_SYMBOL_DEFINITION PreviousPlatformCertificates *//* START_OF_SYMBOL_DEFINITION TCGSpecificationVersion */
// ### ASN.1 Definition:
// 
//   TCGSpecificationVersion ::= SEQUENCE {
//     majorVersion    INTEGER,
//     minorVersion    INTEGER,
//     revision        INTEGER
// }
// 
// 
type TCGSpecificationVersion struct {
    MajorVersion int
    MinorVersion int
    Revision int
}
/* END_OF_SYMBOL_DEFINITION TCGSpecificationVersion *//* START_OF_SYMBOL_DEFINITION CryptographicAnchors */
// ### ASN.1 Definition:
// 
//   CryptographicAnchors  ::=  SEQUENCE (SIZE(1..MAX)) OF Trait
// 
type CryptographicAnchors = [](Trait) // SequenceOfType
/* END_OF_SYMBOL_DEFINITION CryptographicAnchors *//* START_OF_SYMBOL_DEFINITION TCGPlatformSpecification */
// ### ASN.1 Definition:
// 
//   TCGPlatformSpecification ::= SEQUENCE {
//     version         TCGSpecificationVersion,
//     platformClass   OCTET STRING (SIZE(4))
// }
// 
// 
type TCGPlatformSpecification struct {
    Version TCGSpecificationVersion
    PlatformClass []byte
}
/* END_OF_SYMBOL_DEFINITION TCGPlatformSpecification *//* START_OF_SYMBOL_DEFINITION TBBSecurityAssertions_v3 */
// ### ASN.1 Definition:
// 
//   TBBSecurityAssertions-v3  ::=  SEQUENCE (SIZE(1..MAX)) OF Trait
// 
type TBBSecurityAssertions_v3 = [](Trait) // SequenceOfType
/* END_OF_SYMBOL_DEFINITION TBBSecurityAssertions_v3 *//* START_OF_SYMBOL_DEFINITION PlatformConfiguration_v3 */
// ### ASN.1 Definition:
// 
//   PlatformConfiguration-v3 ::= SEQUENCE {
//     platformComponents [0] IMPLICIT SEQUENCE (SIZE(1..MAX)) OF ComponentIdentifier OPTIONAL,
//     platformProperties [1] IMPLICIT SEQUENCE (SIZE(1..MAX)) OF Property OPTIONAL
// }
// 
// 
type PlatformConfiguration_v3 struct {
    PlatformComponents [](ComponentIdentifier) `asn1:"optional,tag:0"`
    PlatformProperties [](Property) `asn1:"optional,tag:1"`
}
/* END_OF_SYMBOL_DEFINITION PlatformConfiguration_v3 *//* START_OF_SYMBOL_DEFINITION ComponentIdentifier */
// ### ASN.1 Definition:
// 
//   ComponentIdentifier  ::=  SEQUENCE (SIZE(1..MAX)) OF Trait
// 
type ComponentIdentifier = [](Trait) // SequenceOfType
/* END_OF_SYMBOL_DEFINITION ComponentIdentifier *//* START_OF_SYMBOL_DEFINITION Property */
// ### ASN.1 Definition:
// 
//   Property ::= SEQUENCE {
//     propertyName    UTF8String (SIZE (1..strmax)),
//     propertyValue   UTF8String (SIZE (1..strmax)),
//     status          [0] IMPLICIT AttributeStatus OPTIONAL
// }
// 
// 
type Property struct {
    PropertyName string
    PropertyValue string
    Status AttributeStatus `asn1:"optional,tag:0"`
}
/* END_OF_SYMBOL_DEFINITION Property *//* START_OF_SYMBOL_DEFINITION AttributeStatus */
// ### ASN.1 Definition:
// 
// AttributeStatus  ::=  ENUMERATED {
//     added (0),
//     modified (1),
//     removed (2)
// }
// 
type AttributeStatus = asn1.Enumerated
const (
    AttributeStatus_Added AttributeStatus = 0,
    AttributeStatus_Modified AttributeStatus = 1,
    AttributeStatus_Removed AttributeStatus = 2
)
/* END_OF_SYMBOL_DEFINITION AttributeStatus *//* START_OF_SYMBOL_DEFINITION PlatformConfigUri_v3 */
// ### ASN.1 Definition:
// 
//   PlatformConfigUri-v3  ::=  SEQUENCE (SIZE(1..MAX)) OF Trait
// 
type PlatformConfigUri_v3 = [](Trait) // SequenceOfType
/* END_OF_SYMBOL_DEFINITION PlatformConfigUri_v3 *//* START_OF_SYMBOL_DEFINITION PlatformOwnership */
// ### ASN.1 Definition:
// 
//   PlatformOwnership  ::=  SEQUENCE (SIZE(1..MAX)) OF Trait
// 
type PlatformOwnership = [](Trait) // SequenceOfType
/* END_OF_SYMBOL_DEFINITION PlatformOwnership *//* START_OF_SYMBOL_DEFINITION CertificateIdentifier */
// ### ASN.1 Definition:
// 
//   CertificateIdentifier ::= SEQUENCE {
//     hashedCertIdentifier        [0] IMPLICIT HashedCertificateIdentifier OPTIONAL,
//     genericCertIdentifier       [1] IMPLICIT IssuerSerial OPTIONAL
// }
// 
// 
type CertificateIdentifier struct {
    HashedCertIdentifier HashedCertificateIdentifier `asn1:"optional,tag:0"`
    GenericCertIdentifier IssuerSerial `asn1:"optional,tag:1"`
}
/* END_OF_SYMBOL_DEFINITION CertificateIdentifier *//* START_OF_SYMBOL_DEFINITION HashedCertificateIdentifier */
// ### ASN.1 Definition:
// 
//   HashedCertificateIdentifier ::= SEQUENCE {
//     hashAlgorithm               AlgorithmIdentifier,
//     hashOverSignatureValue      OCTET STRING
// }
// 
// 
type HashedCertificateIdentifier struct {
    HashAlgorithm AlgorithmIdentifier
    HashOverSignatureValue []byte
}
/* END_OF_SYMBOL_DEFINITION HashedCertificateIdentifier *//* START_OF_SYMBOL_DEFINITION CommonCriteriaEvaluation */
// ### ASN.1 Definition:
// 
//   CommonCriteriaEvaluation ::= SEQUENCE {
//     cCMeasures                  CommonCriteriaMeasures,
//     cCCertificateNumber         UTF8String (SIZE (1..strmax)),
//     cCCertificateAuthority      UTF8String (SIZE (1..strmax)),
//     evaluationScheme            [0] IMPLICIT UTF8String (SIZE (1..strmax)) OPTIONAL,
//     cCCertificateIssuanceDate   [1] IMPLICIT GeneralizedTime OPTIONAL,
//     cCCertificateExpiryDate     [2] IMPLICIT GeneralizedTime OPTIONAL
// }
// 
// 
type CommonCriteriaEvaluation struct {
    CCMeasures CommonCriteriaMeasures
    CCCertificateNumber string
    CCCertificateAuthority string
    EvaluationScheme string `asn1:"optional,tag:0"`
    CCCertificateIssuanceDate time.Time `asn1:"optional,tag:1"`
    CCCertificateExpiryDate time.Time `asn1:"optional,tag:2"`
}
/* END_OF_SYMBOL_DEFINITION CommonCriteriaEvaluation *//* START_OF_SYMBOL_DEFINITION ComponentIdentifierV11 */
// ### ASN.1 Definition:
// 
//   ComponentIdentifierV11 ::= SEQUENCE {
//     componentClass              ComponentClass,
//     componentManufacturer       UTF8String (SIZE (1..STRMAX)),
//     componentModel              UTF8String (SIZE (1..STRMAX)),
//     componentSerial             [0] IMPLICIT UTF8String (SIZE (1..STRMAX)) OPTIONAL,
//     componentRevision           [1] IMPLICIT UTF8String (SIZE (1..STRMAX)) OPTIONAL,
//     componentManufacturerId     [2] IMPLICIT PrivateEnterpriseNumber OPTIONAL,
//     fieldReplaceable            [3] IMPLICIT BOOLEAN OPTIONAL,
//     componentAddresses          [4] IMPLICIT SEQUENCE (SIZE(1.. MAX)) OF ComponentAddress OPTIONAL,
//     componentPlatformCert       [5] IMPLICIT CertificateIdentifier OPTIONAL,
//     componentPlatformCertUri    [6] IMPLICIT URIReference OPTIONAL,
//     status                      [7] IMPLICIT AttributeStatus OPTIONAL
// }
// 
// ComponentClass ::= SEQUENCE {
//     componentClassRegistry      ComponentClassRegistry,
//     componentClassValue         OCTET STRING (SIZE(4))
// }
// 
// ComponentClassRegistry ::= OBJECT IDENTIFIER (
//     tcg-registry-componentClass-tcg
//     | tcg-registry-componentClass-dmtf
//     | tcg-registry-componentClass-pcie
//     | tcg-registry-componentClass-storage
// )
// 
// PrivateEnterpriseNumber ::= OBJECT IDENTIFIER
// 
// ComponentAddress ::= SEQUENCE {
//     addressType         AddressType,
//     addressValue        UTF8String (SIZE (1..strmax))
// }
// 
// AddressType ::= OBJECT IDENTIFIER (
//     tcg-address-ethernetmac
//     | tcg-address-wlanmac
//     | tcg-address-bluetoothmac
// )
// 
// FIPSLevel ::= SEQUENCE {
//     version             IA5STRING (SIZE (1..strmax)), -- “140-1”, “140-2”, or “140-3”
//     level               SecurityLevel,
//     plus                BOOLEAN DEFAULT FALSE
// }
// 
// SecurityLevel ::= ENUMERATED {
//     level1 (1),
//     level2 (2),
//     level3 (3),
//     level4 (4)
// }
// 
// ISO9000Certification ::= SEQUENCE {
//     iso9000Certified    BOOLEAN DEFAULT FALSE,
//     iso9000Uri          IA5STRING (SIZE (1..URIMAX)) OPTIONAL
// }
// 
// PlatformFirmwareCapabilities ::= BIT STRING {
//     fwSetupAuthLocal (0),
//     fwSetupAuthRemote (1),
//     sMMProtection (2),
//     fwKernelDMAProtection (3)
// }
// 
// PlatformFirmwareSignatureVerification ::= BIT STRING {
//     hardwareSRTM (0),
//     secureBoot (1)
// }
// 
// PlatformFirmwareUpdateCompliance ::= BIT STRING {
//     sp800-147 (0),
//     sp800-147B (1),
//     sp800-193 (2)
// }
// 
// PlatformHardwareCapabilities ::= BIT STRING {
//     iOMMUSupport (0),
//     trustedExecutionEnvironment (1),
//     physicalTamperProtection (2),
//     physicalTamperDetection (3),
//     firmwareFlashWP (4),
//     externalDMASupport (5)
// }
// 
// RTMTypes ::= BIT STRING {
//     static (0),
//     dynamic (1),
//     nonHost (2),
//     virtual (3),
//     hardwareStatic (4),
//     bMC (5)
// }
// 
// -- Attributes
// 
// tCGCredentialType ATTRIBUTE ::= {
//     WITH SYNTAX     TCGCredentialType
//     ID              tcg-at-tcgCredentialType
// }
// 
// previousPlatformCertificates ATTRIBUTE ::= {
//     WITH SYNTAX     PreviousPlatformCertificates
//     ID              tcg-at-previousPlatformCertificates
// }
// 
// tCGCredentialSpecification ATTRIBUTE ::= {
//     WITH SYNTAX     TCGSpecificationVersion
//     ID              tcg-at-tcgCredentialSpecification
// }
// 
// cryptographicAnchors ATTRIBUTE ::= {
//     WITH SYNTAX     CryptographicAnchors
//     ID              tcg-at-cryptographicAnchors
// }
// 
// tCGPlatformSpecification ATTRIBUTE ::= {
//     WITH SYNTAX     TCGPlatformSpecification
//     ID              tcg-at-tcgPlatformSpecification
// }
// 
// tBBSecurityAssertions ATTRIBUTE ::= {
//     WITH SYNTAX     TBBSecurityAssertions-v3
//     ID              tcg-at-tbbSecurityAssertions-v3
// }
// 
// platformConfiguration ATTRIBUTE ::= {
//     WITH SYNTAX     PlatformConfiguration-v3
//     ID              tcg-at-platformConfiguration-v3
// }
// 
// platformConfigUri ATTRIBUTE ::= {
//     WITH SYNTAX     PlatformConfigUri-v3
//     ID              tcg-at-platformConfigUri-v3
// }
// 
// platformOwnership ATTRIBUTE ::= {
//     WITH SYNTAX     PlatformOwnership
//     ID              tcg-at-platformOwnership
// }
// 
// -- Traits
// 
// -- // FIXME: capitalized in the spec.
// booleanTrait TRAIT ::= {
//     SYNTAX          BOOLEAN
//     IDENTIFIED BY   tcg-tr-ID-Boolean
// }
// 
// certificateIdentifierTrait TRAIT ::= {
//     SYNTAX          CertificateIdentifier
//     IDENTIFIED BY   tcg-tr-ID-certificateIdentifier
// }
// 
// commonCriteriaTrait TRAIT ::= {
//     SYNTAX          CommonCriteriaEvaluation
//     IDENTIFIED BY   tcg-tr-ID-CommonCriteria
// }
// 
// componentClassTrait TRAIT ::= {
//     SYNTAX          OCTET STRING (SIZE(4))
//     IDENTIFIED BY   tcg-tr-ID-componentClass
// }
// 
// componentIdentifierV11Trait TRAIT ::= {
//     SYNTAX          ComponentIdentifierV11
//     IDENTIFIED BY   tcg-tr-ID-componentIdentifierV11
// }
// 
// fipsLevelTrait TRAIT ::= {
//     SYNTAX          FIPSLevel
//     IDENTIFIED BY   tcg-tr-ID-FIPSLevel
// }
// 
// iso9000Trait TRAIT ::= {
//     SYNTAX          ISO9000Certification
//     IDENTIFIED BY   tcg-tr-ID-ISO9000
// }
// 
// networkMACTrait TRAIT ::= {
//     SYNTAX          ComponentAddress
//     IDENTIFIED BY   tcg-tr-ID-networkMAC
// }
// 
// oidTrait TRAIT ::= {
//     SYNTAX          OBJECT IDENTIFIER
//     IDENTIFIED BY   tcg-tr-ID-OID
// }
// 
// penTrait TRAIT ::= {
//     SYNTAX          PrivateEnterpriseNumber
//     IDENTIFIED BY   tcg-tr-ID-PEN
// }
// 
// platformFirmwareCapabilitiesTrait TRAIT ::= {
//     SYNTAX          PlatformFirmwareCapabilities
//     IDENTIFIED BY   tcg-tr-ID-platformFirmwareCapabilities
// }
// 
// platformFirmwareSignatureVerificationTrait TRAIT ::= {
//     SYNTAX          PlatformFirmwareSignatureVerification
//     IDENTIFIED BY   tcg-tr-ID-platformFirmwareSignatureVerification
// }
// 
// platformFirmwareUpdateComplianceTrait TRAIT ::= {
//     SYNTAX          PlatformFirmwareUpdateCompliance
//     IDENTIFIED BY   tcg-tr-ID-platformFirmwareUpdateCompliance
// }
// 
// platformHardwareCapabilitiesTrait TRAIT ::= {
//     SYNTAX          PlatformHardwareCapabilities
//     IDENTIFIED BY   tcg-tr-ID-platformHardwareCapabilities
// }
// 
// rtmTrait TRAIT ::= {
//     SYNTAX          RTMTypes
//     IDENTIFIED BY   tcg-tr-ID-RTM
// }
// 
// statusTrait TRAIT ::= {
//     SYNTAX          AttributeStatus
//     IDENTIFIED BY   tcg-tr-ID-status
// }
// 
// uriTrait TRAIT ::= {
//     SYNTAX          URIReference
//     IDENTIFIED BY   tcg-tr-ID-URI
// }
// 
// utf8StringTrait TRAIT ::= {
//     SYNTAX          UTF8String (SIZE (1..strmax))
//     IDENTIFIED BY   tcg-tr-ID-UTF8String
// }
// 
// -- // TODO: This is UTF8StringTrait in the spec
// ia5StringTrait TRAIT ::= {
//     SYNTAX          IA5String (SIZE (1..strmax))
//     IDENTIFIED BY   tcg-tr-ID-IA5String
// }
// 
// pemCertStringTrait TRAIT ::= {
//     SYNTAX          UTF8String (SIZE (1..certstrmax))
//     IDENTIFIED BY   tcg-tr-ID-PEMCertString
// }
// 
// publicKeyTrait TRAIT ::= {
//     SYNTAX          SubjectPublicKeyInfo
//     IDENTIFIED BY   tcg-tr-ID-PublicKey
// }
// 
// END
// 
// 
type ComponentIdentifierV11 struct {
    ComponentClass ComponentClass
    ComponentManufacturer string
}
/* END_OF_SYMBOL_DEFINITION ComponentIdentifierV11 */