package x500

import (
	"encoding/asn1"
	"time"
)

// Anything that uses the components of CommonArguments or CommonArgumentsSeq
type CommonArgumentsInterface interface {
	GetServiceControls() ServiceControls
	GetSecurityParameters() SecurityParameters
	GetRequestor() DistinguishedName
	GetOperationProgress() OperationProgress
	GetAliasedRDNs() int
	GetCriticalExtensions() asn1.BitString
	GetReferenceType() ReferenceType
	GetEntryOnly() bool
	GetExclusions() Exclusions
	GetNameResolveOnMaster() bool
	GetOperationContexts() ContextSelection
	GetFamilyGrouping() FamilyGrouping
}

// Anything that uses the components of CommonResults or CommonResultsSeq
type CommonResultsInterface interface {
	GetSecurityParameters() SecurityParameters
	GetPerformer() DistinguishedName
	GetAliasDereferenced() bool
	GetNotification() []Attribute
}

// Anything that uses the components of AccessPoint
type AccessPointInterface interface {
	GetAETitle() Name
	GetAddress() PresentationAddress
	GetProtocolInformation() []ProtocolInformation
}

// Anything that uses the components of MasterOrShadowAccessPoint
type MasterOrShadowAccessPointInterface interface {
	AccessPointInterface
	GetCategory() MasterOrShadowAccessPoint_category
	GetChainingRequired() bool
}

// Anything that uses the components of AVMPcommonComponents
type AVMPcommonComponentsInterface interface {
	GetVersion() AVMPversion
	GetTimestamp() time.Time
	GetAVMPSequence() AVMPsequence
}

// Anything that uses the components of CASPcommonComponents
type CASPcommonComponentsInterface interface {
	GetVersion() CASPversion
	GetSequence() CASPsequence
}

// Any type of X.500 directory schema element, such as an attribute type,
// object class, or name form.
type SchemaElement interface {
	GetName() []UnboundedDirectoryString
	GetDescription() string
	GetObsolete() bool
}

// Anything identified by an object identifier
type ObjectIdentifierIdentified interface {
	GetIdentifier() asn1.ObjectIdentifier
}

// Anything that contains SecurityParameters
type WithSecurityParameters interface {
	GetSecurityParameters() SecurityParameters
}

// Anything that addresses an entry in the directory, either by its
// distinguished name or by one of the other variants of Name
// defined in Amendment 1 to ITU-T Recommendation X.501 (2019).
type WithTargetObject interface {
	// Either Name or DistinguishedName will be populated.
  // But you should still check that both pointers aren't nil anyway.
	GetTargetObject() (*Name, *DistinguishedName)
}

// Anything (a directory error) with a problem code
type WithProblemCode interface {
	GetProblem() int
}

