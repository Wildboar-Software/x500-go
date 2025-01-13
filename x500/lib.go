package x500

import (
	"encoding/asn1"
	"time"
)

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

type CommonResultsInterface interface {
	GetSecurityParameters() SecurityParameters
	GetPerformer() DistinguishedName
	GetAliasDereferenced() bool
	GetNotification() []Attribute
}

type AccessPointInterface interface {
	GetAETitle() Name
	GetAddress() PresentationAddress
	GetProtocolInformation() []ProtocolInformation
}

type MasterOrShadowAccessPointInterface interface {
	AccessPointInterface
	GetCategory() MasterOrShadowAccessPoint_category
	GetChainingRequired() bool
}

type AVMPcommonComponentsInterface interface {
	GetVersion() AVMPversion
	GetTimestamp() time.Time
	GetAVMPSequence() AVMPsequence
}

type CASPcommonComponentsInterface interface {
	GetVersion() CASPversion
	GetSequence() CASPsequence
}

type SchemaElement interface {
	GetName() []UnboundedDirectoryString
	GetDescription() UnboundedDirectoryString
	GetObsolete() bool
}

type ObjectIdentifierIdentified interface {
	GetIdentifier() asn1.ObjectIdentifier
}

type WithSecurityParameters interface {
	GetSecurityParameters() SecurityParameters
}

type WithTargetObject interface {
	// Either Name or DistinguishedName will be populated.
	GetTargetObject() (*Name, *DistinguishedName)
}

type WithProblemCode interface {
	GetProblem() int
}
