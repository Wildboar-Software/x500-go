// X.500 Directory Access Protocol (DAP) client
package x500_dap_client

import (
	"context"
	"encoding/asn1"
	"math/big"

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
	PwdResponseTimeLeft        int                  // -1 if unset
	PwdResponseGracesRemaining int                  // -1 if unset
	PwdResponseError           int                  // -1 if unset
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
	InvokeId      x500.InvokeId // This is always set for any outcome type.
	OpCode        x500.Code     // Only set if OutcomeType == OPERATION_OUTCOME_TYPE_RESULT
	ErrCode       x500.Code     // Only set if OutcomeType == OPERATION_OUTCOME_TYPE_ERROR
	Parameter     asn1.RawValue // The result or error.
	RejectProblem RejectProblem // Only set if OutcomeType == OPERATION_OUTCOME_TYPE_REJECT
	Abort         X500Abort     // Only set if OutcomeType == OPERATION_OUTCOME_TYPE_ABORT
	err           error         // Only set if OutcomeType == OPERATION_OUTCOME_TYPE_FAILURE
}

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

// X.500 Directory Access Protocol (DAP) Client
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
}

// X.500 Directory Access Protocol (DAP) Client with convenience functions that
// provide a simpler interface.
type SimpleDirectoryAccessClient interface {
	DirectoryAccessClient

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

	CompareSimple(ctx context.Context, dn DN, ava x500.AttributeValueAssertion) (resp X500OpOutcome, result *x500.CompareResultData, err error)

	// Invoke the `changePassword` operation on the entry named by distinguished
	// name `dn` using unencrypted values `old` and `new` for `oldPwd` and
	// `newPwd` respectively.
	ChangePasswordSimple(ctx context.Context, dn DN, old string, new string) (resp X500OpOutcome, result *x500.ChangePasswordResultData, err error)

	// Invoke the `administerPassword` operation on the entry named by
	// distinguished name `dn` using the unencrypted value `new` for `newPwd`.
	AdministerPasswordSimple(ctx context.Context, dn DN, new string) (resp X500OpOutcome, result *x500.AdministerPasswordResultData, err error)

	// Entry Modifications

	// Add a new attribute to an entry, returning an X.500 attribute error if
	// the attribute already exists.
	AddAttribute(ctx context.Context, dn DN, attr x500.Attribute) (resp X500OpOutcome, result *x500.ModifyEntryResultData, err error)

	// Remove an attribute from an entry entirely, returning an X.500 attribute
	// error if the attribute does not exist.
	RemoveAttribute(ctx context.Context, dn DN, attr x500.AttributeType) (resp X500OpOutcome, result *x500.ModifyEntryResultData, err error)

	// Add new values to an entry, creating the attribute if it does not exist.
	// If the values already exist, a directory error is returned.
	AddValues(ctx context.Context, dn DN, values x500.Attribute) (resp X500OpOutcome, result *x500.ModifyEntryResultData, err error)

	// Remove values from an entry, returning an error if one or more do not
	// exist. If the last value is removed, the whole attribute is removed.
	RemoveValues(ctx context.Context, dn DN, values x500.Attribute) (resp X500OpOutcome, result *x500.ModifyEntryResultData, err error)

	// Add `addend` to the values of the attribute type. `addend` could be
	// negative, which would result in subtraction.
	AlterValues(ctx context.Context, dn DN, attrtype x500.AttributeType, addend int) (resp X500OpOutcome, result *x500.ModifyEntryResultData, err error)

	// Remove all values that have contexts for which fallback is FALSE.
	ResetValue(ctx context.Context, dn DN, attr x500.AttributeType) (resp X500OpOutcome, result *x500.ModifyEntryResultData, err error)

	// Replace an attribute entirely. If the supplied attribute is empty, the
	// existing attribute is deleted, if it exists, but no error is returned if
	// it does not.
	ReplaceValues(ctx context.Context, dn DN, attr x500.Attribute) (resp X500OpOutcome, result *x500.ModifyEntryResultData, err error)
}

type DN = x500.DistinguishedName

// Client for managing groups: X.500 directory entries of object class
// `groupOfNames` (`2.5.6.9`), `groupOfUniqueNames` (`2.5.6.17`), or any other
// object class that takes values of type `member` (`2.5.4.31`) or
// `uniqueMember` (`2.5.4.50`).
type DirectoryGroupClient interface {

	// Add a group member. If `uid` is `nil`, the `member` attribute is used,
	// otherwise, the `uniqueMember` attribute is used.
	GroupAdd(ctx context.Context, group, member DN, uid *asn1.BitString) (resp X500OpOutcome, result *x500.ModifyEntryResultData, err error)

	// Remove a group member. If `uid` is `nil`, the `member` attribute is used,
	// otherwise, the `uniqueMember` attribute is used.
	GroupRemove(ctx context.Context, group, member DN, uid *asn1.BitString) (resp X500OpOutcome, result *x500.ModifyEntryResultData, err error)

	// Check for the presence of a member in a group. If `uid` is `nil`, the
	// `member` attribute is used, otherwise, the `uniqueMember` attribute is used.
	GroupCheckMember(ctx context.Context, group, member DN, uid *asn1.BitString) (resp X500OpOutcome, result *x500.CompareResultData, err error)
}

const (
	CRIT_EXT_BIT_SUBENTRIES                            = 1
	CRIT_EXT_BIT_COPY_SHALL_DO                         = 2
	CRIT_EXT_BIT_ATTRIBUTE_SIZE_LIMIT                  = 3
	CRIT_EXT_BIT_EXTRA_ATTRIBUTES                      = 4
	CRIT_EXT_BIT_MODIFY_RIGHTS_REQUEST                 = 5
	CRIT_EXT_BIT_PAGED_RESULTS_REQUEST                 = 6
	CRIT_EXT_BIT_MATCHED_VALUES_ONLY                   = 7
	CRIT_EXT_BIT_EXTENDED_FILTER                       = 8
	CRIT_EXT_BIT_TARGET_SYSTEM                         = 9
	CRIT_EXT_BIT_USE_ALIAS_ON_UPDATE                   = 10
	CRIT_EXT_BIT_NEW_SUPERIOR                          = 11
	CRIT_EXT_BIT_MANAGE_DSA_IT                         = 12
	CRIT_EXT_BIT_USE_OF_CONTEXTS                       = 13
	CRIT_EXT_BIT_PARTIAL_NAME_RESOLUTION               = 14
	CRIT_EXT_BIT_OVERSPEC_FILTER                       = 15
	CRIT_EXT_BIT_SELECTION_ON_MODIFY                   = 16
	CRIT_EXT_BIT_SECURITY_OPERATION_CODE               = 18
	CRIT_EXT_BIT_SECURITY_ATTRIBUTE_CERTIFICATION_PATH = 19
	CRIT_EXT_BIT_SECURITY_ERROR_PROTECTION             = 20
	CRIT_EXT_BIT_SERVICE_ADMINISTRATION                = 25
	CRIT_EXT_BIT_ENTRY_COUNT                           = 26
	CRIT_EXT_BIT_HIERARCHY_SELECTIONS                  = 27
	CRIT_EXT_BIT_RELAXATION                            = 28
	CRIT_EXT_BIT_FAMILY_GROUPING                       = 29
	CRIT_EXT_BIT_FAMILY_RETURN                         = 30
	CRIT_EXT_BIT_DN_ATTRIBUTES                         = 31
	CRIT_EXT_BIT_FRIEND_ATTRIBUTES                     = 32
	CRIT_EXT_BIT_ABANDON_OF_PAGED_RESULTS              = 33
	CRIT_EXT_BIT_PAGED_RESULTS_ON_THE_DSP              = 34
	CRIT_EXT_BIT_REPLACE_VALUES                        = 35
)

// SetBit modifies a bit at the given index in an ASN.1 BitString.
// If set is true, the bit is set; if false, the bit is cleared.
// If the index is beyond the current length, the BitString is extended.
func setBit(bs *asn1.BitString, index int, set bool) {
	byteIndex := index / 8
	bitOffset := uint(7 - (index % 8))

	// Extend the bytes slice if needed
	if byteIndex >= len(bs.Bytes) {
		newBytes := make([]byte, byteIndex+1)
		copy(newBytes, bs.Bytes)
		bs.Bytes = newBytes
	}

	// Set or clear the bit
	if set {
		bs.Bytes[byteIndex] |= (1 << bitOffset) // Set bit
	} else {
		bs.Bytes[byteIndex] &^= (1 << bitOffset) // Clear bit
	}

	// Update bit length if needed
	if index+1 > bs.BitLength {
		bs.BitLength = index + 1
	}
}

// Set a bit in Critical Extensions
func setCritExtBit(bs *asn1.BitString, index int) {
	if index <= 0 {
		return
	}
	setBit(bs, index-1, true)
}
