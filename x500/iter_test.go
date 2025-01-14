package x500

import (
	"crypto/x509/pkix"
	"encoding/asn1"
	"fmt"
	"testing"
)

func getListSubordinate(cn string) ListResultData_listInfo_subordinates_Item {
	var rdn RelativeDistinguishedName = []pkix.AttributeTypeAndValue{
		{
			Type: Id_at_commonName,
			Value: asn1.RawValue{
				Class:      asn1.ClassUniversal,
				Tag:        asn1.TagUTF8String,
				IsCompound: false,
				Bytes:      []byte(cn),
			},
		},
	}
	return ListResultData_listInfo_subordinates_Item{
		Rdn: rdn,
	}
}

func getListInfo(size int) ListResultData_listInfo {
	subordinates := []ListResultData_listInfo_subordinates_Item{}
	for i := range size {
		subordinates = append(subordinates, getListSubordinate(fmt.Sprintf("%d", i)))
	}
	return ListResultData_listInfo{
		Subordinates: subordinates,
	}
}

func getComplexListResults() ListResult {
	info1 := getListInfo(3)
	info2 := getListInfo(5)
	info3 := getListInfo(7)
	info4 := getListInfo(11)
	info5 := getListInfo(13)
	info6 := getListInfo(17)
	info7 := getListInfo(19)
	info8 := getListInfo(23)
	info9 := getListInfo(0)

	resultBytes1, _ := asn1.MarshalWithParams(info1, "set")
	resultBytes2, _ := asn1.MarshalWithParams(info2, "set")
	resultBytes3, _ := asn1.MarshalWithParams(info3, "set")
	resultBytes4, _ := asn1.MarshalWithParams(info4, "set")
	resultBytes5, _ := asn1.MarshalWithParams(info5, "set")
	resultBytes6, _ := asn1.MarshalWithParams(info6, "set")
	resultBytes7, _ := asn1.MarshalWithParams(info7, "set")
	resultBytes8, _ := asn1.MarshalWithParams(info8, "set")
	resultBytes9, _ := asn1.MarshalWithParams(info9, "set")

	result1 := asn1.RawValue{FullBytes: resultBytes1}
	result2 := asn1.RawValue{FullBytes: resultBytes2}
	result3 := asn1.RawValue{FullBytes: resultBytes3}
	result4 := asn1.RawValue{FullBytes: resultBytes4}
	result5 := asn1.RawValue{FullBytes: resultBytes5}
	result6 := asn1.RawValue{FullBytes: resultBytes6}
	result7 := asn1.RawValue{FullBytes: resultBytes7}
	result8 := asn1.RawValue{FullBytes: resultBytes8}
	result9 := asn1.RawValue{FullBytes: resultBytes9}

	ucli1 := []ListResult{result1, result2}
	ucli2 := []ListResult{result3, result7}

	ub1, _ := asn1.MarshalWithParams(ucli1, "set")
	ub2, _ := asn1.MarshalWithParams(ucli2, "set")

	u1 := asn1.RawValue{
		Class:      asn1.ClassContextSpecific,
		Tag:        0,
		IsCompound: true,
		Bytes:      ub1,
	}
	u2 := asn1.RawValue{
		Class:      asn1.ClassContextSpecific,
		Tag:        0,
		IsCompound: true,
		Bytes:      ub2,
	}

	ucli3 := []ListResult{result4, result5}
	ucli4 := []ListResult{result6, result8, result9, u1, u2}
	ub3, _ := asn1.MarshalWithParams(ucli3, "set")
	ub4, _ := asn1.MarshalWithParams(ucli4, "set")

	u3 := asn1.RawValue{
		Class:      asn1.ClassContextSpecific,
		Tag:        0,
		IsCompound: true,
		Bytes:      ub3,
	}

	u4 := asn1.RawValue{
		Class:      asn1.ClassContextSpecific,
		Tag:        0,
		IsCompound: true,
		Bytes:      ub4,
	}

	ucli5 := []ListResult{u3, u4}
	ub5, err := asn1.MarshalWithParams(ucli5, "set")
	if err != nil {
		panic(err)
	}
	u5 := asn1.RawValue{
		Class:      asn1.ClassContextSpecific,
		Tag:        0,
		IsCompound: true,
		Bytes:      ub5,
	}
	// Serialize, then de-serialize so the RawValue fields are all filled in.
	retbytes, err := asn1.Marshal(u5)
	if err != nil {
		fmt.Println(err)
		panic("could not serialize")
	}
	ret := asn1.RawValue{}
	_, _ = asn1.Unmarshal(retbytes, &ret)
	return ret
}

func TestListResultsCount(t *testing.T) {
	input := getComplexListResults()
	count, err := GetListResultEntriesReturnedCount(input, 5)
	if err != nil {
		t.Error(err)
		return
	}
	if count != 98 {
		t.Errorf("actual count was %d", count)
		return
	}
}
