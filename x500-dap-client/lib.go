// X.500 Directory Access Protocol (DAP) client
package x500

import (
	"context"
	"crypto/x509"
	"encoding/asn1"
	"math/big"
	"time"
  "github.com/Wildboar-Software/x500-go/x500"
)

type OutcomeType = int

const (
	OP_OUTCOME_FAILURE OutcomeType = 0
	OP_OUTCOME_RESULT  OutcomeType = 1
	OP_OUTCOME_ERROR   OutcomeType = 2
	OP_OUTCOME_REJECT  OutcomeType = 3
	OP_OUTCOME_ABORT   OutcomeType = 4
)

type RejectProblem = asn1.Enumerated

const (
	REJECT_PROBLEM_UNRECOGNIZED_PDU               RejectProblem = 0
	REJECT_PROBLEM_MISTYPED_PDU                   RejectProblem = 1
	REJECT_PROBLEM_BADLY_STRUCTURED_PDU           RejectProblem = 2
	REJECT_PROBLEM_DUPLICATE_INVOCATION           RejectProblem = 10
	REJECT_PROBLEM_UNRECOGNIZED_OPERATION         RejectProblem = 11
	REJECT_PROBLEM_MISTYPED_ARGUMENT              RejectProblem = 12
	REJECT_PROBLEM_RESOURCE_LIMITATION            RejectProblem = 13
	REJECT_PROBLEM_RELEASE_IN_PROGRESS            RejectProblem = 14
	REJECT_PROBLEM_UNRECOGNIZED_INVOCATION_RESULT RejectProblem = 20
	REJECT_PROBLEM_RESULT_RESPONSE_UNEXPECTED     RejectProblem = 21
	REJECT_PROBLEM_MISTYPED_RESULT                RejectProblem = 22
	REJECT_PROBLEM_UNRECOGNIZED_INVOCATION_ERROR  RejectProblem = 30
	REJECT_PROBLEM_ERROR_RESPONSE_UNEXPECTED      RejectProblem = 31
	REJECT_PROBLEM_UNRECOGNIZED_ERROR             RejectProblem = 32
	REJECT_PROBLEM_UNEXPECTED_ERROR               RejectProblem = 33
	REJECT_PROBLEM_MISTYPED_PARAMETER             RejectProblem = 34
)

type AttributeOrValueProblem struct {
	Problem int // Meaningless in updateError
	Type    asn1.ObjectIdentifier
	Value   asn1.RawValue // Meaningless in updateError
}

// TODO: Use this?
type X500OperationError struct {
	ApplicationContext asn1.ObjectIdentifier
	InvokeId           x500.InvokeId
	Code               int
	Problem            int                       // Present in many error types
	MatchedObject      x500.Name                      // Used by attributeError and nameError
	AbandonedOperation x500.InvokeId                  // Used by abandoned
	Candidate          x500.ContinuationReference     // Used by referral
	ContextPrefix      x500.DistinguishedName         // Used by dsaReferral
	AttributeProblems  []AttributeOrValueProblem // Used by updateError and attributeError
	SecurityParameters x500.SecurityParameters
	Performer          x500.DistinguishedName
	AliasDereferenced  bool
	Notification       []x500.Attribute
}

// The parameters of an X.500 directory bind, as well as those of the underlying
// Remote Operation Service Element (ROSE) and Association Control Service
// Element (ACSE), which are unused if not applicable to the underlying
// operation transport.
type X500AssociateArgument struct {
	// OSI Protocol Fields
	ModeSelector                      int // SHOULD always be 1
	OSIProtocolVersion1               bool
	CallingPresentationSelector       []byte
	CalledPresentationSelector        []byte
	PresentationContextDefinitionList x500.Context_list
	TransferSyntaxName                asn1.ObjectIdentifier
	PresentationContextIdentifier     int

	// AARQ-apdu Fields
	ACSEProtocolVersion1          bool
	ApplicationContext            asn1.ObjectIdentifier
	CalledAPTitle                 x500.DistinguishedName
	CalledAETitle                 x500.GeneralName
	CalledAPInvocationIdentifier  x500.AP_invocation_identifier
	CalledAEInvocationIdentifier  x500.AE_invocation_identifier
	CallingAPTitle                x500.DistinguishedName
	CallingAETitle                x500.GeneralName
	CallingAPInvocationIdentifier x500.AP_invocation_identifier
	CallingAEInvocationIdentifier x500.AE_invocation_identifier
	ImplementationInformation     string

	// Fields from DirectoryBindArgument
	V1          bool
	V2          bool
	Credentials *x500.Credentials
}

// Describes the result of a directory bind, using either IDM or OSI protocols,
// and covering both bind success (result), bind error, and abort.
// (Reject is not a valid outcome for a bind operation.)
// To determine success, just check that OutcomeType == OPERATION_OUTCOME_TYPE_RESULT.
type X500AssociateOutcome struct {
	OutcomeType OutcomeType
	err         error         // Only set if OutcomeType == OPERATION_OUTCOME_TYPE_FAILURE
	Parameter   asn1.RawValue // Only set if OutcomeType != OPERATION_OUTCOME_TYPE_ABORT
	Abort       X500Abort     // Only set if OutcomeType == OPERATION_OUTCOME_TYPE_ABORT

	// OSI Protocol Fields
	ModeSelector                      int // SHOULD always be 1
	OSIProtocolVersion1               bool
	RespondingPresentationSelector    *big.Int
	PresentationContextDefinitionList x500.Result_list
	TransferSyntaxName                asn1.ObjectIdentifier
	PresentationContextIdentifier     int
	ProviderReason                    x500.Provider_reason

	// ACSE AARE-apdu and AAREerr-apdu Fields
	ACSEProtocolVersion1              bool
	ApplicationContext                asn1.ObjectIdentifier
	ACSEResult                        x500.Associate_result                                  // 0 if successful
	AssociateSourceDiagnosticUser     x500.Associate_source_diagnostic_acse_service_user     // Set to -1 if unset
	AssociateSourceDiagnosticProvider x500.Associate_source_diagnostic_acse_service_provider // Set to -1 if unset

	// ACSE uses separate responding AP-title and AE-qualifier fields, because
	// it is assumed that every AE has a directory name.
	// IDM uses GeneralName. Since a GeneralName is a superset, we use that
	// instead.
	RespondingAPTitle                x500.DistinguishedName
	RespondingAETitle                x500.GeneralName
	RespondingAPInvocationIdentifier int
	RespondingAEInvocationIdentifier int
	ImplementationInformation        string

	// Fields specific to the IDM protocol
	AETitleError x500.IdmBindError_aETitleError // Set to -1 if unset

	// Fields from DirectoryBindResult, directoryBindError
	V1                         bool
	V2                         bool
	Credentials                x500.Credentials     // NULL value if unset
	PwdResponseTimeLeft        int             // -1 if unset
	PwdResponseGracesRemaining int             // -1 if unset
	PwdResponseError           int             // -1 if unset
	ServiceError               x500.ServiceProblem  // Set to -1 if unset
	SecurityError              x500.SecurityProblem // Set to -1 if unset
	SecurityParameters         x500.SecurityParameters
}

// A Remote Operation Service Element (ROSE) request
type X500Request struct {
	PresentationContextIdentifier x500.Presentation_context_identifier
	InvokeId                      x500.InvokeId
	OpCode                        x500.Code
	Argument                      asn1.RawValue
}

// A Remote Operation Service Element (ROSE) abort
type X500Abort struct {
	PresentationContextIdentifierList x500.Presentation_context_identifier_list
	PresentationContextIdentifier     x500.Presentation_context_identifier
	AbortSource                       x500.ABRT_source
	ProviderReason                    x500.Abort_reason
	EventIdentifier                   x500.Event_identifier
	UserReason                        x500.Abort // Only used by IDM Abort
}

// A Remote Operation Service Element (ROSE) outcome: a generalization over
// results, errors, rejections, aborts, etc. In other words, a union of all
// possible outcomes to a ROSE operation other than a failure happening at the
// lower layers, such as a TCP socket closure.
type X500OpOutcome struct {
	OutcomeType   OutcomeType
	InvokeId      x500.InvokeId      // This is always set for any outcome type.
	OpCode        x500.Code          // Only set if OutcomeType == OPERATION_OUTCOME_TYPE_RESULT
	ErrCode       x500.Code          // Only set if OutcomeType == OPERATION_OUTCOME_TYPE_ERROR
	Parameter     asn1.RawValue // The result or error.
	RejectProblem RejectProblem // Only set if OutcomeType == OPERATION_OUTCOME_TYPE_REJECT
	Abort         X500Abort     // Only set if OutcomeType == OPERATION_OUTCOME_TYPE_ABORT
	err           error         // Only set if OutcomeType == OPERATION_OUTCOME_TYPE_FAILURE
}

// TODO: Rename to generalize over ROSE instead of X.500
// A Remote Operation Service Element (ROSE) unbind request
type X500UnbindRequest struct {
	PresentationContextIdentifier x500.Presentation_context_identifier
	Reason                        x500.Release_request_reason
}

// A Remote Operation Service Element (ROSE) unbind outcome
type X500UnbindOutcome struct {
	PresentationContextIdentifier x500.Presentation_context_identifier
	Reason                        x500.Release_response_reason
}

// An Internet Directly-Mapped `startTLS` message outcome, which may be a
// `tlsResponse` message or an error.
type StartTLSOutcome struct {
	err      error
	response x500.TLSResponse
}

// Remote Operation Service Element (ROSE) per ITU-T Recommendation X.880.
//
// For now, this is only implemented via the Internet Directly-Mapped (IDM)
// protocol defined in
// [ITU-T Recommendation X.519 (2019)](https://www.itu.int/itu-t/recommendations/rec.aspx?rec=X.519),
// but in the future, this may be implemented via
// [ISO Transport over TCP (ITOT)](https://www.rfc-editor.org/rfc/rfc1006),
// [Lightweight Presentation Protocol](https://www.rfc-editor.org/rfc/rfc1085),
// [DIXIE](https://www.rfc-editor.org/rfc/rfc1249), and others.
type RemoteOperationServiceElement interface {

	// Perform the Remote Operation Service Element (ROSE) Bind operation
	Bind(ctx context.Context, arg X500AssociateArgument) (response X500AssociateOutcome, err error)

	// Issue a Remote Operation Service Element (ROSE) request
	Request(ctx context.Context, req X500Request) (response X500OpOutcome, err error)

	// Unbind via the Remote Operation Service Element (ROSE)
	Unbind(ctx context.Context, req X500UnbindRequest) (response X500UnbindOutcome, err error)

	// Close the network layers / transport layers that facilitate the
	// Remote Operation Service Element (ROSE), such as by closing a TCP socket.
	// This is a separate operation from unbinding so that the underlying
	// transport may be re-used for a new session.
	CloseTransport() (err error)
}

type DirectoryAccessClient interface {
	RemoteOperationServiceElement

	// DAP Operations

	// Perform the `read` Directory Access Protocol (DAP) operation. The `result` returned may be `nil`.
	Read(ctx context.Context, arg_data x500.ReadArgumentData) (response X500OpOutcome, result *x500.ReadResultData, err error)

	// Perform the `compare` Directory Access Protocol (DAP) operation. The `result` returned may be `nil`.
	Compare(ctx context.Context, arg_data x500.CompareArgumentData) (resp X500OpOutcome, result *x500.CompareResultData, err error)

	// Perform the `abandon` Directory Access Protocol (DAP) operation. The `result` returned may be `nil`.
	Abandon(ctx context.Context, arg_data x500.AbandonArgumentData) (resp X500OpOutcome, result *x500.AbandonResultData, err error)

	// Perform the `list` Directory Access Protocol (DAP) operation. The `info` returned may be `nil`.
	List(ctx context.Context, arg_data x500.ListArgumentData) (resp X500OpOutcome, info *x500.ListResultData_listInfo, err error)

	// Perform the `search` Directory Access Protocol (DAP) operation. The `info` returned may be `nil`.
	Search(ctx context.Context, arg_data x500.SearchArgumentData) (resp X500OpOutcome, info *x500.SearchResultData_searchInfo, err error)

	// Perform the `addEntry` Directory Access Protocol (DAP) operation. The `result` returned may be `nil`.
	AddEntry(ctx context.Context, arg_data x500.AddEntryArgumentData) (resp X500OpOutcome, result *x500.AddEntryResultData, err error)

	// Perform the `removeEntry` Directory Access Protocol (DAP) operation. The `result` returned may be `nil`.
	RemoveEntry(ctx context.Context, arg_data x500.RemoveEntryArgumentData) (resp X500OpOutcome, result *x500.RemoveEntryResultData, err error)

	// Perform the `modifyEntry` Directory Access Protocol (DAP) operation. The `result` returned may be `nil`.
	ModifyEntry(ctx context.Context, arg_data x500.ModifyEntryArgumentData) (resp X500OpOutcome, result *x500.ModifyEntryResultData, err error)

	// Perform the `modifyDN` Directory Access Protocol (DAP) operation. The `result` returned may be `nil`.
	ModifyDN(ctx context.Context, arg_data x500.ModifyDNArgumentData) (resp X500OpOutcome, result *x500.ModifyDNResultData, err error)

	// Perform the `changePassword` Directory Access Protocol (DAP) operation. The `result` returned may be `nil`.
	ChangePassword(ctx context.Context, arg_data x500.ChangePasswordArgumentData) (resp X500OpOutcome, result *x500.ChangePasswordResultData, err error)

	// Perform the `administerPassword` Directory Access Protocol (DAP) operation. The `result` returned may be `nil`.
	AdministerPassword(ctx context.Context, arg_data x500.AdministerPasswordArgumentData) (resp X500OpOutcome, result *x500.AdministerPasswordResultData, err error)

	// Higher-Level DAP API

	// Bind using simple authentication: your distinguished name and password
	BindSimply(ctx context.Context, dn DN, password string) (resp X500AssociateOutcome, err error)

	// Bind by signing a token with your configured signing key.
	// An error will be returned if no signing key or no signing cert is configured.
	// The `requesterDN` is _your_ DN. The `recipientDN` is the application entity
	// title of the DSA.
	BindStrongly(ctx context.Context, requesterdn DN, recipientdn DN, acPath *x500.AttributeCertificationPath) (resp X500AssociateOutcome, err error)

	// Bind using the `PLAIN` Simple Authentication and Security Layer (SASL) Mechanism
	BindPlainly(ctx context.Context, username string, password string) (resp X500AssociateOutcome, err error)

	// Read selected user attributes from an entry named by a distinguished name.
	// `result` may be `nil`
	ReadSimple(ctx context.Context, dn DN, userAttributes []asn1.ObjectIdentifier) (response X500OpOutcome, result *x500.ReadResultData, err error)

	// Remove an entry named by the distinguished name `dn` using the
	// `removeEntry` X.500 Directory Access Protocol (DAP) operation.
	RemoveEntryByDN(ctx context.Context, dn DN) (resp X500OpOutcome, result *x500.RemoveEntryResultData, err error)

	// Abandon an X.500 Directory operation by its invoke ID.
	AbandonById(ctx context.Context, invokeId int) (resp X500OpOutcome, result *x500.AbandonResultData, err error)

	// List up to `limit` subordinates underneath the entry named by the
	// distinguished name `dn`.
	ListByDN(ctx context.Context, dn DN, limit int) (resp X500OpOutcome, info *x500.ListResultData_listInfo, err error)

	// Add a new entry having distinguished name `dn` and composed of attributes
	// `attrs`. If you need to use the `targetSystem` field to add an entry in
	// another DSA and establish a hierarchical operational binding, you will
	// need to use the lower-level [AddEntry] method.
	AddEntrySimple(ctx context.Context, dn DN, attrs []x500.Attribute) (resp X500OpOutcome, result *x500.AddEntryResultData, err error)

	// Invoke the `changePassword` operation on the entry named by distinguished
	// name `dn` using unencrypted values `old` and `new` for `oldPwd` and
	// `newPwd` respectively.
	ChangePasswordSimple(ctx context.Context, dn DN, old string, new string) (resp X500OpOutcome, result *x500.ChangePasswordResultData, err error)

	// Invoke the `administerPassword` operation on the entry named by
	// distinguished name `dn` using the unencrypted value `new` for `newPwd`.
	AdministerPasswordSimple(ctx context.Context, dn DN, new string) (resp X500OpOutcome, result *x500.AdministerPasswordResultData, err error)
}

type DN = x500.DistinguishedName

type Subschema struct {
	AttributeTypes   *[]x500.AttributeTypeDescription
	ObjectClasses    *[]x500.ObjectClassDescription
	NameForms        *[]x500.NameFormDescription
	StructureRules   *[]x500.DITStructureRuleDescription
	ContentRules     *[]x500.DITContentRuleDescription
	Friendships      *[]x500.FriendsDescription
	MatchingRules    *[]x500.MatchingRuleDescription
	MatchingRuleUses *[]x500.MatchingRuleUseDescription
	SyntaxNames      *[]x500.LdapSyntaxDescription
	ContextUses      *[]x500.DITContextUseDescription
	ContextTypes     *[]x500.ContextDescription
}

type DSAInfo struct {
	VendorName              string
	VendorVersion           string
	FullVendorVersion       string
	MyAccessPoint           x500.AccessPoint
	SuperiorKnowledge       []x500.AccessPoint
	DITBridgeKnowledge      []x500.DitBridgeKnowledge
	SubschemaSubentry       x500.DistinguishedName
	SupportedControls       asn1.ObjectIdentifier
	SupportedExtensions     asn1.ObjectIdentifier
	SupportedFeatures       asn1.ObjectIdentifier
	SupportedLDAPVersion    int
	DynamicSubtrees         []x500.DistinguishedName
	AltServers              []string
	SupportedSASLMechanisms []string
	NamingContexts          []x500.DistinguishedName

	// A catch-all, which includes all attributes (user and operational)
	RootDSEAttributes []x500.Attribute
}

// This type describes an entry of object classes person, residentialPerson,
// organizationalPerson, inetOrgPerson, and pilotPerson
type CreatePersonOptions struct {
	commonName            string
	givenName             string
	initials              string
	surname               string
	generationQualifier   string
	description           []string
	telephoneNumber       []string
	faxNumber             []string
	userPassword          string
	manager               []DN
	secretary             []DN
	seeAlso               []DN
	localityName          []string
	stateOrProvinceName   string
	streetAddress         string
	postalAddress         []string
	postalCode            string
	poBox                 string
	businessCategory      []string
	carLicense            []string
	departmentNumber      []string
	displayName           string
	employeeNumber        string
	employeeType          []string
	preferredLanguage     string
	homeTelephoneNumber   []string
	homePostalAddress     []string
	mail                  []string
	labeledURI            []string
	mobileTelephoneNumber []string
	organizationName      []string
	roomNumber            []string
	uid                   []string
	userCertificate       []x509.Certificate
	uniqueIdentifier      []asn1.BitString
	title                 []string
	personalTitle         []string
}

type CreateAccountOptions struct {
	description            []string
	seeAlso                []DN
	localityName           []string
	organizationName       []string
	organizationalUnitName []string
	host                   []string
}

type CreateDocumentOptions struct {
	commonName             []string
	description            []string
	seeAlso                []DN
	localityName           []string
	organizationName       []string
	organizationalUnitName []string
	documentTitle          []string
	documentVersion        []string
	documentAuthor         []DN
	documentLocation       []string
	documentPublisher      []string
}

type CreateDocumentSeriesOptions struct {
	description            []string
	seeAlso                []DN
	telephoneNumber        []string
	localityName           []string
	organizationName       []string
	organizationalUnitName []string
}

type CreateDomainOptions struct {
	associatedName             []DN
	organizationName           []string
	description                []string
	businessCategory           []string
	seeAlso                    []DN
	searchGuide                []x500.Guide
	userPassword               []string
	localityName               []string
	stateOrProvinceName        []string
	postalAddress              []string
	postalCode                 string
	physicalDeliveryOfficeName string
	postOfficeBox              string
	streetAddress              string
	fax                        []string
	telephoneNumber            []string
}

type CreateOrganizationOptions struct {
	description                []string
	businessCategory           []string
	seeAlso                    []DN
	searchGuide                []x500.Guide
	userPassword               string
	stateOrProvinceName        string
	localityName               string
	streetAddress              string
	postalCode                 string
	physicalDeliveryOfficeName string
	postOfficeBox              string
	postalAddress              []string
	telephoneNumber            []string
	faxNumber                  []string
}

type PasswordPolicyOptions struct {
	PwdModifyEntryAllowed      bool
	PwdChangeAllowed           bool
	PwdMaxAge                  int
	PwdExpiryAge               int
	PwdMinLength               int
	PwdVocabulary              asn1.BitString
	PwdAlphabet                []string
	PwdDictionaries            []string
	PwdExpiryWarning           int
	PwdGraces                  int
	PwdFailureDuration         int
	PwdLockoutDuration         int
	PwdMaxFailures             int
	PwdMaxTimeInHistory        int
	PwdMinTimeInHistory        int
	PwdHistorySlots            int
	PwdRecentlyExpiredDuration int
	PwdEncAlg                  int
}

type CRLOptions struct {
	crl       []x509.RevocationList // attributeCertificateRevocationList
	endEntity []x509.RevocationList // eeAttrCertificateRevocationList
	authority []x509.RevocationList // attributeAuthorityRevocationList
}

// ISO Tag = ISO/IEC 15693
type BecomeISOTagInfoOptions struct {
	tagOid      asn1.ObjectIdentifier
	tagAfi      byte // Each is one-byte
	contentUrl  []string
	uii         asn1.BitString
	uiiInUrn    string
	tagLocation *x500.UtmCoordinates
}

type BecomeISOTagTypeOptions struct {
	tagOid    asn1.ObjectIdentifier
	tagAfi    byte
	uiiFormat *x500.UiiFormat
}

type BecomeEPCTaginfoOptions struct {
	epc         asn1.BitString
	epcInUrn    string
	contentUrl  []string
	tagLocation *x500.UtmCoordinates
}

type BecomePosixAccountOptions struct {
	commonName    []string
	uid           string
	uidNumber     int
	gidNumber     int
	homeDirectory string
	userPassword  string
	loginShell    string
	gecos         string
	description   []string
}

type BecomeShadowAccount struct {
	userPassword     string
	shadowLastChange int
	shadowMin        int
	shadowMax        int
	shadowWarning    int
	shadowInactive   int
	shadowExpire     int
	shadowFlag       int
	description      []string
}

type BecomeNaturalPersonOptions struct {
	emailAddress         []string
	unstructuredName     []string
	unstructuredAddress  []string
	dateOfBirth          time.Time
	placeOfBirth         string
	gender               string
	countryOfCitizenship []string
	countryOfResidence   []string
	pseudonym            []string
	serialNumber         []string
}

type SchemaAwareDirectoryAccessClient interface {
	// Create Specific Entries

	CreateSubentry(ctx context.Context, superior DN, name string, subtreeSpec []x500.SubtreeSpecification) (resp X500OpOutcome, result *x500.AddEntryResultData, err error)
	// If name is not "", the friendlyCountry object class is used.
	CreatePerson(ctx context.Context, superior DN, cn, surname string, opts *CreatePersonOptions) (resp X500OpOutcome, result *x500.AddEntryResultData, err error)
	CreateOrgPerson(ctx context.Context, superior DN, cn, surname string, opts *CreatePersonOptions) (resp X500OpOutcome, result *x500.AddEntryResultData, err error)
	CreateResPerson(ctx context.Context, dn DN, cn, surname string, opts *CreatePersonOptions) (resp X500OpOutcome, result *x500.AddEntryResultData, err error)
	CreateInetOrgPerson(ctx context.Context, dn DN, cn, surname string, opts *CreatePersonOptions) (resp X500OpOutcome, result *x500.AddEntryResultData, err error)
	CreatePilotPerson(ctx context.Context, dn DN, cn, surname string, opts *CreatePersonOptions) (resp X500OpOutcome, result *x500.AddEntryResultData, err error)
	CreateCountry(ctx context.Context, superior DN, iso2Code, desc, name string) (resp X500OpOutcome, result *x500.AddEntryResultData, err error)
	CreateStateOrProvince(ctx context.Context, superior DN, name, desc string) (resp X500OpOutcome, result *x500.AddEntryResultData, err error)
	CreateLocality(ctx context.Context, superior DN, name, desc string) (resp X500OpOutcome, result *x500.AddEntryResultData, err error)
	CreateOrganization(ctx context.Context, superior DN, name string, opts *CreateOrganizationOptions) (resp X500OpOutcome, result *x500.AddEntryResultData, err error)
	CreateOrgUnit(ctx context.Context, superior DN, name string, opts *CreateOrganizationOptions) (resp X500OpOutcome, result *x500.AddEntryResultData, err error)
	CreateOrgRole(ctx context.Context, superior DN, name, desc string, members []DN) (resp X500OpOutcome, result *x500.AddEntryResultData, err error)
	CreateProcess(ctx context.Context, superior DN, name, desc string) (resp X500OpOutcome, result *x500.AddEntryResultData, err error)
	CreateAppEntity(ctx context.Context, superior DN, name string, addr x500.PresentationAddress, desc string) (resp X500OpOutcome, result *x500.AddEntryResultData, err error)
	CreateDSA(ctx context.Context, superior DN, name string, addr x500.PresentationAddress, desc string) (resp X500OpOutcome, result *x500.AddEntryResultData, err error)
	CreateDevice(ctx context.Context, superior DN, name, serial, desc string, owner DN) (resp X500OpOutcome, result *x500.AddEntryResultData, err error)
	CreateDMD(ctx context.Context, superior DN, name, desc string) (resp X500OpOutcome, result *x500.AddEntryResultData, err error)
	CreateOIDRoot(ctx context.Context, superior DN, arc1, arc2, arc3 int, aliasedEntry DN) (resp X500OpOutcome, result *x500.AddEntryResultData, err error)
	CreateOIDArc(ctx context.Context, superior DN, arc int, aliasedEntry DN) (resp X500OpOutcome, result *x500.AddEntryResultData, err error)
	CreateGroupOfUrlsEntry(ctx context.Context, superior DN, name, desc, org string, members []string, owner DN) (resp X500OpOutcome, result *x500.AddEntryResultData, err error)
	CreateAliasEntry(ctx context.Context, dn, aliasedEntry DN) (resp X500OpOutcome, result *x500.AddEntryResultData, err error)
	CreateCRLDPEntry(ctx context.Context, superior DN, name string, crl, eepk, authority, delta *x509.RevocationList) (resp X500OpOutcome, result *x500.AddEntryResultData, err error)
	CreateDomain(ctx context.Context, superior DN, dc string, opts *CreateDomainOptions) (resp X500OpOutcome, result *x500.AddEntryResultData, err error)
	CreateAccount(ctx context.Context, superior DN, uid string, opts *CreateAccountOptions) (resp X500OpOutcome, result *x500.AddEntryResultData, err error)
	CreateDocument(ctx context.Context, superior DN, docid string, opts *CreateDocumentOptions) (resp X500OpOutcome, result *x500.AddEntryResultData, err error)
	CreateDocumentSeries(ctx context.Context, superior DN, name string, opts *CreateDocumentSeriesOptions) (resp X500OpOutcome, result *x500.AddEntryResultData, err error)
	CreateRoom(ctx context.Context, superior DN, name, number, desc, phone string, seeAlso []DN) (resp X500OpOutcome, result *x500.AddEntryResultData, err error)
	CreatePosixGroup(ctx context.Context, superior DN, gid int, name, desc, userpwd string, members []string) (resp X500OpOutcome, result *x500.AddEntryResultData, err error)
	CreateIPService(ctx context.Context, superior DN, name, transport string, port int, manager DN) (resp X500OpOutcome, result *x500.AddEntryResultData, err error)
	CreateIPNetwork(ctx context.Context, superior DN, name, ip, netmask, loc, desc string) (resp X500OpOutcome, result *x500.AddEntryResultData, err error)

	// Add Specific Auxiliary Object Classes

	BecomePasswordAdminSubentry(ctx context.Context, dn DN, pwdAttribute x500.AttributeType, opts *PasswordPolicyOptions) (resp X500OpOutcome, result *x500.ModifyEntryResultData, err error)
	BecomeSubschemaSubentry(ctx context.Context, dn DN, schema *Subschema) (resp X500OpOutcome, result *x500.ModifyEntryResultData, err error)
	BecomeServiceAdminSubentry(ctx context.Context, dn DN, searchRules []x500.SearchRuleDescription) (resp X500OpOutcome, result *x500.ModifyEntryResultData, err error)
	BecomeContextAssertionDefaultSubentry(ctx context.Context, dn DN, cads []x500.TypeAndContextAssertion) (resp X500OpOutcome, result *x500.ModifyEntryResultData, err error)
	BecomePMIUser(ctx context.Context, dn DN, ac *x500.AttributeCertificate) (resp X500OpOutcome, result *x500.ModifyEntryResultData, err error)
	BecomePMIAA(ctx context.Context, dn DN, aaCert *x500.AttributeCertificate, crls *CRLOptions) (resp X500OpOutcome, result *x500.ModifyEntryResultData, err error)
	BecomePMISOA(ctx context.Context, dn DN, adc *x500.AttributeCertificate, crls *CRLOptions) (resp X500OpOutcome, result *x500.ModifyEntryResultData, err error)
	BecomeAttCertCRLDistPoint(ctx context.Context, dn DN, crls *CRLOptions) (resp X500OpOutcome, result *x500.ModifyEntryResultData, err error)
	BecomePKIUser(ctx context.Context, dn DN, userCert *x509.Certificate) (resp X500OpOutcome, result *x500.ModifyEntryResultData, err error)
	BecomePKICA(ctx context.Context, dn DN, cacert *x509.Certificate, opts *CRLOptions) (resp X500OpOutcome, result *x500.ModifyEntryResultData, err error)
	BecomePMIDelegationPath(ctx context.Context, dn DN, path []x500.AttCertPath) (resp X500OpOutcome, result *x500.ModifyEntryResultData, err error)
	BecomePrivilegePolicy(ctx context.Context, dn DN, policies []x500.PolicySyntax) (resp X500OpOutcome, result *x500.ModifyEntryResultData, err error)
	BecomeProtectedPrivilegePolicy(ctx context.Context, dn DN, policies []x500.AttributeCertificate) (resp X500OpOutcome, result *x500.ModifyEntryResultData, err error)
	BecomeDeltaCRL(ctx context.Context, dn DN, deltaCRLs []x509.RevocationList) (resp X500OpOutcome, result *x500.ModifyEntryResultData, err error)
	BecomeCPCPS(ctx context.Context, dn DN, statements []x500.InfoSyntax, policies []x500.PolicySyntax) (resp X500OpOutcome, result *x500.ModifyEntryResultData, err error)
	BecomePKICertPath(ctx context.Context, dn DN, paths []x500.PkiPath) (resp X500OpOutcome, result *x500.ModifyEntryResultData, err error)
	BecomeStrongAuthenticationUser(ctx context.Context, dn DN, usercerts []x509.Certificate) (resp X500OpOutcome, result *x500.ModifyEntryResultData, err error)
	BecomeUserSecurityInfo(ctx context.Context, dn DN, supportedAlgs []x500.SupportedAlgorithm) (resp X500OpOutcome, result *x500.ModifyEntryResultData, err error)
	BecomeUserPwdClass(ctx context.Context, dn DN, pwd x500.UserPwd) (resp X500OpOutcome, result *x500.ModifyEntryResultData, err error)
	BecomeCertificationAuthority(ctx context.Context, dn DN, cert x509.Certificate, crls *CRLOptions) (resp X500OpOutcome, result *x500.ModifyEntryResultData, err error)
	BecomeISOTagInfo(ctx context.Context, dn DN, opts *BecomeISOTagInfoOptions) (resp X500OpOutcome, result *x500.ModifyEntryResultData, err error)
	BecomeISOTagType(ctx context.Context, dn DN, opts *BecomeISOTagTypeOptions) (resp X500OpOutcome, result *x500.ModifyEntryResultData, err error)
	BecomeEPCTagInfo(ctx context.Context, dn DN, opts *BecomeEPCTaginfoOptions) (resp X500OpOutcome, result *x500.ModifyEntryResultData, err error)
	BecomeEPCTagType(ctx context.Context, dn DN, format x500.UiiFormat) (resp X500OpOutcome, result *x500.ModifyEntryResultData, err error)
	BecomePosixAccount(ctx context.Context, dn DN, opts *BecomePosixAccountOptions) (resp X500OpOutcome, result *x500.ModifyEntryResultData, err error)
	BecomeShadowAccount(ctx context.Context, dn DN, opts *BecomeShadowAccount) (resp X500OpOutcome, result *x500.ModifyEntryResultData, err error)
	BecomeIEEE802Device(ctx context.Context, dn DN, macAddress string) (resp X500OpOutcome, result *x500.ModifyEntryResultData, err error)
	BecomeIPHost(ctx context.Context, dn DN, cn, ip, desc, loc string, manager DN) (resp X500OpOutcome, result *x500.ModifyEntryResultData, err error)
	BecomeNaturalPerson(ctx context.Context, dn DN, opts *BecomeNaturalPersonOptions) (resp X500OpOutcome, result *x500.ModifyEntryResultData, err error)

	// Methods for Viewing Schema

	GetSubschema(ctx context.Context, dn DN) (resp X500OpOutcome, result *Subschema, err error)
	ListAttributeTypes(ctx context.Context, dn DN) (resp X500OpOutcome, result *[]x500.AttributeTypeDescription, err error)
	ListObjectClasses(ctx context.Context, dn DN) (resp X500OpOutcome, result *[]x500.ObjectClassDescription, err error)
	ListNameForms(ctx context.Context, dn DN) (resp X500OpOutcome, result *[]x500.NameFormDescription, err error)
	ListStructureRules(ctx context.Context, dn DN) (resp X500OpOutcome, result *[]x500.DITStructureRuleDescription, err error)
	ListContentRules(ctx context.Context, dn DN) (resp X500OpOutcome, result *[]x500.DITContentRuleDescription, err error)
	ListFriendships(ctx context.Context, dn DN) (resp X500OpOutcome, result *[]x500.FriendsDescription, err error)
	ListMatchingRules(ctx context.Context, dn DN) (resp X500OpOutcome, result *[]x500.MatchingRuleDescription, err error)
	ListMatchingRuleUses(ctx context.Context, dn DN) (resp X500OpOutcome, result *[]x500.MatchingRuleUseDescription, err error)
	ListSyntaxNames(ctx context.Context, dn DN) (resp X500OpOutcome, result *[]x500.LdapSyntaxDescription, err error)
	ListContextUses(ctx context.Context, dn DN) (resp X500OpOutcome, result *[]x500.DITContextUseDescription, err error)
	ListSearchRules(ctx context.Context, dn DN) (resp X500OpOutcome, result *[]x500.SearchRuleDescription, err error)
	ListSubentries(ctx context.Context, dn DN) (resp X500OpOutcome, result *x500.ListResultData, err error)

	// Group Management

	GroupAdd(ctx context.Context, group, member DN, uid *asn1.BitString) (resp X500OpOutcome, result *x500.ModifyEntryResultData, err error)
	GroupRemove(ctx context.Context, group, member DN, uid *asn1.BitString) (resp X500OpOutcome, result *x500.ModifyEntryResultData, err error)
	GroupCheckMember(ctx context.Context, group, member DN, uid *asn1.BitString) (resp X500OpOutcome, result *x500.ModifyEntryResultData, err error)

	// PKI and PMI Management

	AddPublicKeyCertificate(ctx context.Context, dn DN, cert x509.Certificate) (resp X500OpOutcome, result *x500.ModifyEntryResultData, err error)
	AddAttributeCertificate(ctx context.Context, dn DN, acert x500.AttributeCertificate) (resp X500OpOutcome, result *x500.ModifyEntryResultData, err error)

	// Administrative Role Management

	AddAdministrativeRole(ctx context.Context, dn DN, role asn1.ObjectIdentifier) (resp X500OpOutcome, result *x500.ModifyEntryResultData, err error)
	RemoveAdministrativeRole(ctx context.Context, dn DN, role asn1.ObjectIdentifier) (resp X500OpOutcome, result *x500.ModifyEntryResultData, err error)

	// Entry Modifications

	AddAttribute(ctx context.Context, dn DN, attr x500.Attribute) (resp X500OpOutcome, result *x500.ModifyEntryResultData, err error)
	RemoveAttribute(ctx context.Context, dn DN, attr x500.AttributeType) (resp X500OpOutcome, result *x500.ModifyEntryResultData, err error)
	AddValues(ctx context.Context, dn DN, values x500.Attribute) (resp X500OpOutcome, result *x500.ModifyEntryResultData, err error)
	RemoveValues(ctx context.Context, dn DN, values x500.Attribute) (resp X500OpOutcome, result *x500.ModifyEntryResultData, err error)
	AlterValues(ctx context.Context, dn DN, attrtype x500.AttributeType, addend int) (resp X500OpOutcome, result *x500.ModifyEntryResultData, err error)
	ResetValue(ctx context.Context, dn DN, attr x500.AttributeType) (resp X500OpOutcome, result *x500.ModifyEntryResultData, err error)
	ReplaceValues(ctx context.Context, dn DN, attr x500.Attribute) (resp X500OpOutcome, result *x500.ModifyEntryResultData, err error)

	// Specific Modification Types

	AddCollectiveExclusions(ctx context.Context, dn DN, attr x500.AttributeType) (resp X500OpOutcome, result *x500.ModifyEntryResultData, err error)
	SetAliasedEntryName(ctx context.Context, src, dest DN) (resp X500OpOutcome, result *x500.ModifyEntryResultData, err error)
	AddPermission(ctx context.Context, dn DN, operation, object string) (resp X500OpOutcome, result *x500.ModifyEntryResultData, err error)
	AddHierarchicalParent(ctx context.Context, dn, parent x500.DistinguishedName) (resp X500OpOutcome, result *x500.ModifyEntryResultData, err error)

	// General Entry Modifications

	AddObjectClass(ctx context.Context, dn DN, oc asn1.ObjectIdentifier) (resp X500OpOutcome, result *x500.ModifyEntryResultData, err error)
	RemoveObjectClass(ctx context.Context, dn DN, oc asn1.ObjectIdentifier) (resp X500OpOutcome, result *x500.ModifyEntryResultData, err error)

	// Operational Attribute Modifications
	SetCreateTimestamp(ctx context.Context, dn DN, ts time.Time) (resp X500OpOutcome, result *x500.ModifyEntryResultData, err error)
	SetModifyTimestamp(ctx context.Context, dn DN, ts time.Time) (resp X500OpOutcome, result *x500.ModifyEntryResultData, err error)
	SetCreatorsName(ctx context.Context, dn, creator DN) (resp X500OpOutcome, result *x500.ModifyEntryResultData, err error)
	SetModifiersName(ctx context.Context, dn, modifier DN) (resp X500OpOutcome, result *x500.ModifyEntryResultData, err error)

	// Miscellaneous Utilities
	GetDSAinfo(ctx context.Context) (info DSAInfo, err error)
}

type CommonArgumentsInterface interface {
	GetServiceControls() x500.ServiceControls
	GetSecurityParameters() x500.SecurityParameters
	GetRequestor() DN
	GetOperationProgress() x500.OperationProgress
	GetAliasedRDNs() int
	GetCriticalExtensions() asn1.BitString
	GetReferenceType() x500.ReferenceType
	GetEntryOnly() bool
	GetExclusions() x500.Exclusions
	GetNameResolveOnMaster() bool
	GetOperationContexts() x500.ContextSelection
	GetFamilyGrouping() x500.FamilyGrouping
}

type CommonResultsInterface interface {
	GetSecurityParameters() *x500.SecurityParameters
	GetPerformer() *DN
	GetAliasDereferenced() bool
	GetNotification() []x500.Attribute
}

type AccessPointInterface interface {
	GetAETitle() x500.Name
	GetAddress() x500.PresentationAddress
	GetProtocolInformation() []x500.ProtocolInformation
}

type AVMPcommonComponentsInterface interface {
	GetVersion() x500.AVMPversion
	GetTimestamp() time.Time
	GetAVMPSequence() x500.AVMPsequence
}

type SchemaElement interface {
	GetName() string
	GetDescription() string
	GetObsolete() bool
}

type ObjectIdentifierIdentified interface {
	GetOIDIdentifier() asn1.ObjectIdentifier
}

type DirectoryStringIdentified interface {
	GetStringIdentifier() asn1.RawValue
}

type IntIdentified interface {
	GetIntIdentifier() int
}

