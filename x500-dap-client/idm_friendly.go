package x500_dap_client

import (
	"context"
	"encoding/asn1"

	"github.com/Wildboar-Software/x500-go/x500"
)

func getMemberAttr(member DN, uid *asn1.BitString) (asn1.RawValue, error) {
	if uid == nil || uid.BitLength == 0 {
		fullBytes, err := asn1.Marshal(member)
		if err != nil {
			return asn1.RawValue{}, err
		}
		return asn1.RawValue{FullBytes: fullBytes}, nil
	}
	name := x500.NameAndOptionalUID{
		Dn:  member,
		Uid: *uid,
	}
	fullBytes, err := asn1.Marshal(name)
	if err != nil {
		return asn1.RawValue{}, err
	}
	return asn1.RawValue{FullBytes: fullBytes}, nil
}

// If uid is not nil, the uniqueMember attribute will be used.
// Otherwise, the member attribute will be used.
func (stack *IDMProtocolStack) GroupAdd(ctx context.Context, group, member DN, uid *asn1.BitString) (resp X500OpOutcome, result *x500.ModifyEntryResultData, err error) {
	attr := x500.Attribute{
		Type:   x500.Id_at_member,
		Values: make([]asn1.RawValue, 0, 1),
	}
	value, err := getMemberAttr(member, uid)
	if err != nil {
		return X500OpOutcome{}, nil, err
	}
	attr.Values = append(attr.Values, value)
	if uid != nil && uid.BitLength > 0 {
		attr.Type = x500.Id_at_uniqueMember
	}
	return stack.AddValues(ctx, group, attr)
}

// If uid is not nil, the uniqueMember attribute will be used.
// Otherwise, the member attribute will be used.
func (stack *IDMProtocolStack) GroupRemove(ctx context.Context, group, member DN, uid *asn1.BitString) (resp X500OpOutcome, result *x500.ModifyEntryResultData, err error) {
	attr := x500.Attribute{
		Type:   x500.Id_at_member,
		Values: make([]asn1.RawValue, 0, 1),
	}
	value, err := getMemberAttr(member, uid)
	if err != nil {
		return X500OpOutcome{}, nil, err
	}
	attr.Values = append(attr.Values, value)
	if uid != nil && uid.BitLength > 0 {
		attr.Type = x500.Id_at_uniqueMember
	}
	return stack.RemoveValues(ctx, group, attr)
}

// If uid is not nil, the uniqueMember attribute will be used.
// Otherwise, the member attribute will be used.
func (stack *IDMProtocolStack) GroupCheckMember(ctx context.Context, group, member DN, uid *asn1.BitString) (resp X500OpOutcome, result *x500.CompareResultData, err error) {
	value, err := getMemberAttr(member, uid)
	if err != nil {
		return X500OpOutcome{}, nil, err
	}
	ava := x500.AttributeValueAssertion{
		Type:      x500.Id_at_member,
		Assertion: value,
	}
	if uid != nil && uid.BitLength > 0 {
		ava.Type = x500.Id_at_uniqueMember
	}
	return stack.CompareSimple(ctx, group, ava)
}
