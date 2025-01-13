package x500

import (
	"testing"
)

func TestInterfaceImplementation(t *testing.T) {
	var ok bool

	// CommonArgumentsInterface

	var commonArguments interface{} = &CommonArguments{}
	_, ok = commonArguments.(CommonArgumentsInterface)
	if !ok {
		t.Error("interface not implemented")
	}
	var commonArgumentsSeq interface{} = &CommonArgumentsSeq{}
	_, ok = commonArgumentsSeq.(CommonArgumentsInterface)
	if !ok {
		t.Error("interface not implemented")
	}
	var readArgumentData interface{} = &ReadArgumentData{}
	_, ok = readArgumentData.(CommonArgumentsInterface)
	if !ok {
		t.Error("interface not implemented")
	}
	var compareArgumentData interface{} = &CompareArgumentData{}
	_, ok = compareArgumentData.(CommonArgumentsInterface)
	if !ok {
		t.Error("interface not implemented")
	}
	var listArgumentData interface{} = &ListArgumentData{}
	_, ok = listArgumentData.(CommonArgumentsInterface)
	if !ok {
		t.Error("interface not implemented")
	}
	var searchArgumentData interface{} = &SearchArgumentData{}
	_, ok = searchArgumentData.(CommonArgumentsInterface)
	if !ok {
		t.Error("interface not implemented")
	}
	var addEntryArgumentData interface{} = &AddEntryArgumentData{}
	_, ok = addEntryArgumentData.(CommonArgumentsInterface)
	if !ok {
		t.Error("interface not implemented")
	}
	var removeEntryArgumentData interface{} = &RemoveEntryArgumentData{}
	_, ok = removeEntryArgumentData.(CommonArgumentsInterface)
	if !ok {
		t.Error("interface not implemented")
	}
	var modifyEntryArgumentData interface{} = &ModifyEntryArgumentData{}
	_, ok = modifyEntryArgumentData.(CommonArgumentsInterface)
	if !ok {
		t.Error("interface not implemented")
	}
	var modifyDNArgumentData interface{} = &ModifyDNArgumentData{}
	_, ok = modifyDNArgumentData.(CommonArgumentsInterface)
	if !ok {
		t.Error("interface not implemented")
	}
	var ldapArgumentData interface{} = &LdapArgumentData{}
	_, ok = ldapArgumentData.(CommonArgumentsInterface)
	if !ok {
		t.Error("interface not implemented")
	}
	var linkedArgumentData interface{} = &LinkedArgumentData{}
	_, ok = linkedArgumentData.(CommonArgumentsInterface)
	if !ok {
		t.Error("interface not implemented")
	}

	// CommonResultsInterface

	var commonResults interface{} = &CommonResults{}
	_, ok = commonResults.(CommonResultsInterface)
	if !ok {
		t.Error("interface not implemented")
	}
	var commonResultsSeq interface{} = &CommonResultsSeq{}
	_, ok = commonResultsSeq.(CommonResultsInterface)
	if !ok {
		t.Error("interface not implemented")
	}
	var readResultData interface{} = &ReadResultData{}
	_, ok = readResultData.(CommonResultsInterface)
	if !ok {
		t.Error("interface not implemented")
	}
	var compareResultData interface{} = &CompareResultData{}
	_, ok = compareResultData.(CommonResultsInterface)
	if !ok {
		t.Error("interface not implemented")
	}
	var abandonResultData interface{} = &AbandonResultData{}
	_, ok = abandonResultData.(CommonResultsInterface)
	if !ok {
		t.Error("interface not implemented")
	}
	var addEntryResultData interface{} = &AddEntryResultData{}
	_, ok = addEntryResultData.(CommonResultsInterface)
	if !ok {
		t.Error("interface not implemented")
	}
	var removeEntryResultData interface{} = &RemoveEntryResultData{}
	_, ok = removeEntryResultData.(CommonResultsInterface)
	if !ok {
		t.Error("interface not implemented")
	}
	var modifyEntryResultData interface{} = &ModifyEntryResultData{}
	_, ok = modifyEntryResultData.(CommonResultsInterface)
	if !ok {
		t.Error("interface not implemented")
	}
	var modifyDNResultData interface{} = &ModifyDNResultData{}
	_, ok = modifyDNResultData.(CommonResultsInterface)
	if !ok {
		t.Error("interface not implemented")
	}
	var changePasswordResultData interface{} = &ChangePasswordResultData{}
	_, ok = changePasswordResultData.(CommonResultsInterface)
	if !ok {
		t.Error("interface not implemented")
	}
	var administerPasswordResultData interface{} = &AdministerPasswordResultData{}
	_, ok = administerPasswordResultData.(CommonResultsInterface)
	if !ok {
		t.Error("interface not implemented")
	}
	var ldapResultData interface{} = &LdapResultData{}
	_, ok = ldapResultData.(CommonResultsInterface)
	if !ok {
		t.Error("interface not implemented")
	}
	var abandonedData interface{} = &AbandonedData{}
	_, ok = abandonedData.(CommonResultsInterface)
	if !ok {
		t.Error("interface not implemented")
	}
	var abandonFailedData interface{} = &AbandonFailedData{}
	_, ok = abandonFailedData.(CommonResultsInterface)
	if !ok {
		t.Error("interface not implemented")
	}
	var attributeErrorData interface{} = &AttributeErrorData{}
	_, ok = attributeErrorData.(CommonResultsInterface)
	if !ok {
		t.Error("interface not implemented")
	}
	var nameErrorData interface{} = &NameErrorData{}
	_, ok = nameErrorData.(CommonResultsInterface)
	if !ok {
		t.Error("interface not implemented")
	}
	var referralData interface{} = &ReferralData{}
	_, ok = referralData.(CommonResultsInterface)
	if !ok {
		t.Error("interface not implemented")
	}
	var securityErrorData interface{} = &SecurityErrorData{}
	_, ok = securityErrorData.(CommonResultsInterface)
	if !ok {
		t.Error("interface not implemented")
	}
	var serviceErrorData interface{} = &ServiceErrorData{}
	_, ok = serviceErrorData.(CommonResultsInterface)
	if !ok {
		t.Error("interface not implemented")
	}
	var updateErrorData interface{} = &UpdateErrorData{}
	_, ok = updateErrorData.(CommonResultsInterface)
	if !ok {
		t.Error("interface not implemented")
	}
	var listResultData_listInfo interface{} = &ListResultData_listInfo{}
	_, ok = listResultData_listInfo.(CommonResultsInterface)
	if !ok {
		t.Error("interface not implemented")
	}
	var searchResultData_searchInfo interface{} = &SearchResultData_searchInfo{}
	_, ok = searchResultData_searchInfo.(CommonResultsInterface)
	if !ok {
		t.Error("interface not implemented")
	}
	var establishOperationalBindingResultData interface{} = &EstablishOperationalBindingResultData{}
	_, ok = establishOperationalBindingResultData.(CommonResultsInterface)
	if !ok {
		t.Error("interface not implemented")
	}
	var modifyOperationalBindingResultData interface{} = &ModifyOperationalBindingResultData{}
	_, ok = modifyOperationalBindingResultData.(CommonResultsInterface)
	if !ok {
		t.Error("interface not implemented")
	}
	var terminateOperationalBindingResultData interface{} = &TerminateOperationalBindingResultData{}
	_, ok = terminateOperationalBindingResultData.(CommonResultsInterface)
	if !ok {
		t.Error("interface not implemented")
	}
	var opBindingErrorParam interface{} = &OpBindingErrorParam{}
	_, ok = opBindingErrorParam.(CommonResultsInterface)
	if !ok {
		t.Error("interface not implemented")
	}
	var dsaReferralData interface{} = &DsaReferralData{}
	_, ok = dsaReferralData.(CommonResultsInterface)
	if !ok {
		t.Error("interface not implemented")
	}
	var requestShadowUpdateResultData interface{} = &RequestShadowUpdateResultData{}
	_, ok = requestShadowUpdateResultData.(CommonResultsInterface)
	if !ok {
		t.Error("interface not implemented")
	}
	var updateShadowResultData interface{} = &UpdateShadowResultData{}
	_, ok = updateShadowResultData.(CommonResultsInterface)
	if !ok {
		t.Error("interface not implemented")
	}
	var shadowErrorData interface{} = &ShadowErrorData{}
	_, ok = shadowErrorData.(CommonResultsInterface)
	if !ok {
		t.Error("interface not implemented")
	}

	// AccessPointInterface

	var accessPoint interface{} = &AccessPoint{}
	_, ok = accessPoint.(AccessPointInterface)
	if !ok {
		t.Error("interface not implemented")
	}
	var masterOrShadowAccessPoint interface{} = &MasterOrShadowAccessPoint{}
	_, ok = masterOrShadowAccessPoint.(AccessPointInterface)
	if !ok {
		t.Error("interface not implemented")
	}
	var accessPointInformation interface{} = &AccessPointInformation{}
	_, ok = accessPointInformation.(AccessPointInterface)
	if !ok {
		t.Error("interface not implemented")
	}
	var supplierOrConsumer interface{} = &SupplierOrConsumer{}
	_, ok = supplierOrConsumer.(AccessPointInterface)
	if !ok {
		t.Error("interface not implemented")
	}
	var supplierAndConsumers interface{} = &SupplierAndConsumers{}
	_, ok = supplierAndConsumers.(AccessPointInterface)
	if !ok {
		t.Error("interface not implemented")
	}

	// MasterOrShadowAccessPointInterface

	_, ok = masterOrShadowAccessPoint.(MasterOrShadowAccessPointInterface)
	if !ok {
		t.Error("interface not implemented")
	}
	_, ok = accessPointInformation.(MasterOrShadowAccessPointInterface)
	if !ok {
		t.Error("interface not implemented")
	}

	// AVMPcommonComponentsInterface

	var aVMPcommonComponents interface{} = &AVMPcommonComponents{}
	_, ok = aVMPcommonComponents.(AVMPcommonComponentsInterface)
	if !ok {
		t.Error("interface not implemented")
	}
	var certReq interface{} = &CertReq{}
	_, ok = certReq.(AVMPcommonComponentsInterface)
	if !ok {
		t.Error("interface not implemented")
	}
	var certRsp interface{} = &CertRsp{}
	_, ok = certRsp.(AVMPcommonComponentsInterface)
	if !ok {
		t.Error("interface not implemented")
	}
	var addAvlReq interface{} = &AddAvlReq{}
	_, ok = addAvlReq.(AVMPcommonComponentsInterface)
	if !ok {
		t.Error("interface not implemented")
	}
	var addAvlRsp interface{} = &AddAvlRsp{}
	_, ok = addAvlRsp.(AVMPcommonComponentsInterface)
	if !ok {
		t.Error("interface not implemented")
	}
	var replaceAvlReq interface{} = &ReplaceAvlReq{}
	_, ok = replaceAvlReq.(AVMPcommonComponentsInterface)
	if !ok {
		t.Error("interface not implemented")
	}
	var replaceAvlRsp interface{} = &ReplaceAvlRsp{}
	_, ok = replaceAvlRsp.(AVMPcommonComponentsInterface)
	if !ok {
		t.Error("interface not implemented")
	}
	var deleteAvlReq interface{} = &DeleteAvlReq{}
	_, ok = deleteAvlReq.(AVMPcommonComponentsInterface)
	if !ok {
		t.Error("interface not implemented")
	}
	var deleteAvlRsp interface{} = &DeleteAvlRsp{}
	_, ok = deleteAvlRsp.(AVMPcommonComponentsInterface)
	if !ok {
		t.Error("interface not implemented")
	}
	var rejectAVL interface{} = &RejectAVL{}
	_, ok = rejectAVL.(AVMPcommonComponentsInterface)
	if !ok {
		t.Error("interface not implemented")
	}

	// CASPcommonComponentsInterface

	var cASPcommonComponents interface{} = &CASPcommonComponents{}
	_, ok = cASPcommonComponents.(CASPcommonComponentsInterface)
	if !ok {
		t.Error("interface not implemented")
	}
	var certSubscribeReq interface{} = &CertSubscribeReq{}
	_, ok = certSubscribeReq.(CASPcommonComponentsInterface)
	if !ok {
		t.Error("interface not implemented")
	}
	var certSubscribeRsp interface{} = &CertSubscribeRsp{}
	_, ok = certSubscribeRsp.(CASPcommonComponentsInterface)
	if !ok {
		t.Error("interface not implemented")
	}
	var certUnsubscribeReq interface{} = &CertUnsubscribeReq{}
	_, ok = certUnsubscribeReq.(CASPcommonComponentsInterface)
	if !ok {
		t.Error("interface not implemented")
	}
	var certUnsubscribeRsp interface{} = &CertUnsubscribeRsp{}
	_, ok = certUnsubscribeRsp.(CASPcommonComponentsInterface)
	if !ok {
		t.Error("interface not implemented")
	}
	var certReplaceReq interface{} = &CertReplaceReq{}
	_, ok = certReplaceReq.(CASPcommonComponentsInterface)
	if !ok {
		t.Error("interface not implemented")
	}
	var certReplaceRsp interface{} = &CertReplaceRsp{}
	_, ok = certReplaceRsp.(CASPcommonComponentsInterface)
	if !ok {
		t.Error("interface not implemented")
	}
	var certUpdateReq interface{} = &CertUpdateReq{}
	_, ok = certUpdateReq.(CASPcommonComponentsInterface)
	if !ok {
		t.Error("interface not implemented")
	}
	var certUpdateRsp interface{} = &CertUpdateRsp{}
	_, ok = certUpdateRsp.(CASPcommonComponentsInterface)
	if !ok {
		t.Error("interface not implemented")
	}
	var rejectCAsubscribe interface{} = &RejectCAsubscribe{}
	_, ok = rejectCAsubscribe.(CASPcommonComponentsInterface)
	if !ok {
		t.Error("interface not implemented")
	}

	// SchemaElement

	var dITStructureRuleDescription interface{} = &DITStructureRuleDescription{}
	_, ok = dITStructureRuleDescription.(SchemaElement)
	if !ok {
		t.Error("interface not implemented")
	}
	var dITContentRuleDescription interface{} = &DITContentRuleDescription{}
	_, ok = dITContentRuleDescription.(SchemaElement)
	if !ok {
		t.Error("interface not implemented")
	}
	var matchingRuleDescription interface{} = &MatchingRuleDescription{}
	_, ok = matchingRuleDescription.(SchemaElement)
	if !ok {
		t.Error("interface not implemented")
	}
	var attributeTypeDescription interface{} = &AttributeTypeDescription{}
	_, ok = attributeTypeDescription.(SchemaElement)
	if !ok {
		t.Error("interface not implemented")
	}
	var objectClassDescription interface{} = &ObjectClassDescription{}
	_, ok = objectClassDescription.(SchemaElement)
	if !ok {
		t.Error("interface not implemented")
	}
	var nameFormDescription interface{} = &NameFormDescription{}
	_, ok = nameFormDescription.(SchemaElement)
	if !ok {
		t.Error("interface not implemented")
	}
	var matchingRuleUseDescription interface{} = &MatchingRuleUseDescription{}
	_, ok = matchingRuleUseDescription.(SchemaElement)
	if !ok {
		t.Error("interface not implemented")
	}
	var contextDescription interface{} = &ContextDescription{}
	_, ok = contextDescription.(SchemaElement)
	if !ok {
		t.Error("interface not implemented")
	}
	var dITContextUseDescription interface{} = &DITContextUseDescription{}
	_, ok = dITContextUseDescription.(SchemaElement)
	if !ok {
		t.Error("interface not implemented")
	}
	var friendsDescription interface{} = &FriendsDescription{}
	_, ok = friendsDescription.(SchemaElement)
	if !ok {
		t.Error("interface not implemented")
	}
	var searchRuleDescription interface{} = &SearchRuleDescription{}
	_, ok = searchRuleDescription.(SchemaElement)
	if !ok {
		t.Error("interface not implemented")
	}

	// ObjectIdentifierIdentified

	_, ok = dITContentRuleDescription.(ObjectIdentifierIdentified)
	if !ok {
		t.Error("interface not implemented")
	}
	_, ok = matchingRuleDescription.(ObjectIdentifierIdentified)
	if !ok {
		t.Error("interface not implemented")
	}
	_, ok = attributeTypeDescription.(ObjectIdentifierIdentified)
	if !ok {
		t.Error("interface not implemented")
	}
	_, ok = objectClassDescription.(ObjectIdentifierIdentified)
	if !ok {
		t.Error("interface not implemented")
	}
	_, ok = nameFormDescription.(ObjectIdentifierIdentified)
	if !ok {
		t.Error("interface not implemented")
	}
	_, ok = matchingRuleUseDescription.(ObjectIdentifierIdentified)
	if !ok {
		t.Error("interface not implemented")
	}
	_, ok = contextDescription.(ObjectIdentifierIdentified)
	if !ok {
		t.Error("interface not implemented")
	}
	_, ok = dITContextUseDescription.(ObjectIdentifierIdentified)
	if !ok {
		t.Error("interface not implemented")
	}
	_, ok = friendsDescription.(ObjectIdentifierIdentified)
	if !ok {
		t.Error("interface not implemented")
	}

	// WithSecurityParameters

	var establishOperationalBindingArgumentData interface{} = &EstablishOperationalBindingArgumentData{}
	_, ok = establishOperationalBindingArgumentData.(WithSecurityParameters)
	if !ok {
		t.Error("interface not implemented")
	}
	var modifyOperationalBindingArgumentData interface{} = &ModifyOperationalBindingArgumentData{}
	_, ok = modifyOperationalBindingArgumentData.(WithSecurityParameters)
	if !ok {
		t.Error("interface not implemented")
	}
	var terminateOperationalBindingArgumentData interface{} = &TerminateOperationalBindingArgumentData{}
	_, ok = terminateOperationalBindingArgumentData.(WithSecurityParameters)
	if !ok {
		t.Error("interface not implemented")
	}
	var chainingArguments interface{} = &ChainingArguments{}
	_, ok = chainingArguments.(WithSecurityParameters)
	if !ok {
		t.Error("interface not implemented")
	}
	var chainingResults interface{} = &ChainingResults{}
	_, ok = chainingResults.(WithSecurityParameters)
	if !ok {
		t.Error("interface not implemented")
	}
	var directoryBindError_OPTIONALLY_PROTECTED_Parameter1 interface{} = &DirectoryBindError_OPTIONALLY_PROTECTED_Parameter1{}
	_, ok = directoryBindError_OPTIONALLY_PROTECTED_Parameter1.(WithSecurityParameters)
	if !ok {
		t.Error("interface not implemented")
	}
	var coordinateShadowUpdateArgumentData interface{} = &CoordinateShadowUpdateArgumentData{}
	_, ok = coordinateShadowUpdateArgumentData.(WithSecurityParameters)
	if !ok {
		t.Error("interface not implemented")
	}
	var requestShadowUpdateArgumentData interface{} = &RequestShadowUpdateArgumentData{}
	_, ok = requestShadowUpdateArgumentData.(WithSecurityParameters)
	if !ok {
		t.Error("interface not implemented")
	}
	var updateShadowArgumentData interface{} = &UpdateShadowArgumentData{}
	_, ok = updateShadowArgumentData.(WithSecurityParameters)
	if !ok {
		t.Error("interface not implemented")
	}

	// WithTargetObject

	_, ok = readArgumentData.(WithTargetObject)
	if !ok {
		t.Error("interface not implemented")
	}
	_, ok = readResultData.(WithTargetObject)
	if !ok {
		t.Error("interface not implemented")
	}
	_, ok = compareArgumentData.(WithTargetObject)
	if !ok {
		t.Error("interface not implemented")
	}
	_, ok = compareResultData.(WithTargetObject)
	if !ok {
		t.Error("interface not implemented")
	}
	_, ok = listArgumentData.(WithTargetObject)
	if !ok {
		t.Error("interface not implemented")
	}
	_, ok = searchArgumentData.(WithTargetObject)
	if !ok {
		t.Error("interface not implemented")
	}
	var joinArgument interface{} = &JoinArgument{}
	_, ok = joinArgument.(WithTargetObject)
	if !ok {
		t.Error("interface not implemented")
	}
	_, ok = addEntryArgumentData.(WithTargetObject)
	if !ok {
		t.Error("interface not implemented")
	}
	_, ok = removeEntryArgumentData.(WithTargetObject)
	if !ok {
		t.Error("interface not implemented")
	}
	_, ok = modifyEntryArgumentData.(WithTargetObject)
	if !ok {
		t.Error("interface not implemented")
	}
	_, ok = modifyDNArgumentData.(WithTargetObject)
	if !ok {
		t.Error("interface not implemented")
	}
	var changePasswordArgumentData interface{} = &ChangePasswordArgumentData{}
	_, ok = changePasswordArgumentData.(WithTargetObject)
	if !ok {
		t.Error("interface not implemented")
	}
	var administerPasswordArgumentData interface{} = &AdministerPasswordArgumentData{}
	_, ok = administerPasswordArgumentData.(WithTargetObject)
	if !ok {
		t.Error("interface not implemented")
	}
	_, ok = ldapArgumentData.(WithTargetObject)
	if !ok {
		t.Error("interface not implemented")
	}
	_, ok = linkedArgumentData.(WithTargetObject)
	if !ok {
		t.Error("interface not implemented")
	}
	_, ok = attributeErrorData.(WithTargetObject)
	if !ok {
		t.Error("interface not implemented")
	}
	_, ok = nameErrorData.(WithTargetObject)
	if !ok {
		t.Error("interface not implemented")
	}
	_, ok = referralData.(WithTargetObject)
	if !ok {
		t.Error("interface not implemented")
	}
	_, ok = listResultData_listInfo.(WithTargetObject)
	if !ok {
		t.Error("interface not implemented")
	}
	_, ok = searchResultData_searchInfo.(WithTargetObject)
	if !ok {
		t.Error("interface not implemented")
	}
	var continuationReference interface{} = &ContinuationReference{}
	_, ok = continuationReference.(WithTargetObject)
	if !ok {
		t.Error("interface not implemented")
	}
	_, ok = chainingArguments.(WithTargetObject)
	if !ok {
		t.Error("interface not implemented")
	}
	var traceItem interface{} = &TraceItem{}
	_, ok = traceItem.(WithTargetObject)
	if !ok {
		t.Error("interface not implemented")
	}

	// WithProblemCode

	_, ok = abandonedData.(WithProblemCode)
	if !ok {
		t.Error("interface not implemented")
	}
	_, ok = abandonFailedData.(WithProblemCode)
	if !ok {
		t.Error("interface not implemented")
	}
	_, ok = nameErrorData.(WithProblemCode)
	if !ok {
		t.Error("interface not implemented")
	}
	_, ok = securityErrorData.(WithProblemCode)
	if !ok {
		t.Error("interface not implemented")
	}
	_, ok = serviceErrorData.(WithProblemCode)
	if !ok {
		t.Error("interface not implemented")
	}
	_, ok = updateErrorData.(WithProblemCode)
	if !ok {
		t.Error("interface not implemented")
	}
	_, ok = opBindingErrorParam.(WithProblemCode)
	if !ok {
		t.Error("interface not implemented")
	}
	_, ok = shadowErrorData.(WithProblemCode)
	if !ok {
		t.Error("interface not implemented")
	}

}
