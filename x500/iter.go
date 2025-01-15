package x500

import (
	"encoding/asn1"
	"errors"
	"fmt"
)

func DissectListResult(lr ListResult) (info *ListResultData_listInfo, subs []ListResult, err error) {
	data := lr
	if data.Class == asn1.ClassUniversal && data.Tag == asn1.TagSequence {
		signedResult := SIGNED{}
		rest, err := asn1.Unmarshal(data.FullBytes, &signedResult)
		if err != nil {
			return nil, make([]ListResult, 0), err
		}
		if len(rest) > 0 {
			return nil, make([]ListResult, 0), errors.New("trailing bytes")
		}
		data = signedResult.ToBeSigned
	}
	if data.Class == asn1.ClassContextSpecific && data.Tag == 0 {
		// This is the uncorrelatedListInfo
		sets := make([]ListResult, 0)
		rest, err := asn1.UnmarshalWithParams(data.Bytes, &sets, "set")
		if err != nil {
			return nil, make([]ListResult, 0), err
		}
		if len(rest) > 0 {
			return nil, make([]ListResult, 0), errors.New("trailing bytes")
		}
		return nil, sets, nil
	}
	if data.Class != asn1.ClassUniversal || data.Tag != asn1.TagSet {
		// This is some other syntax other than listInfo.
		errStr := fmt.Sprintf("unrecognized listresult syntax (class=%+v, tag=%d)", data.Class, data.Tag)
		return nil, make([]ListResult, 0), errors.New(errStr)
	}
	info = &ListResultData_listInfo{}
	rest, err := asn1.UnmarshalWithParams(data.FullBytes, info, "set")
	if err != nil {
		fmt.Println("still list info")
		return nil, make([]ListResult, 0), err
	}
	if len(rest) > 0 {
		return nil, make([]ListResult, 0), errors.New("trailing bytes")
	}
	return info, make([]ListResult, 0), err
}

// Get the number of entries returned in a ListResult.
//
// Once recursionLimit reaches (or falls below) 0, further recursion ceases and
// an error is returned.
//
// If you want it as a feature, I might be able to implement this without
// recursion, but it would not be very elegant.
func GetListResultEntriesReturnedCount(lr ListResult, recursionLimit int) (count int, err error) {
	info, sets, err := DissectListResult(lr)
	if err != nil {
		return 0, err
	}
	if info != nil {
		return len(info.Subordinates), nil
	}
	if recursionLimit <= 0 {
		return 0, errors.New("too much recursion in listresult")
	}
	sum := 0
	for _, s := range sets {
		subcount, err := GetListResultEntriesReturnedCount(s, recursionLimit-1)
		if err != nil {
			return 0, err
		}
		sum += subcount
	}
	return sum, nil
}

type ListResultIterator struct {
	// Stack to hold nodes for depth-first traversal
	stack []ListResult
	// Tracks the index of the current child for each node in the stack
	index []int
}

func NewListIter(lr ListResult) *ListResultIterator {
	return &ListResultIterator{
		stack: []ListResult{lr},
		index: []int{0},
	}
}

func (it *ListResultIterator) HasNext() (bool, error) {
	for len(it.stack) > 0 {
		current := it.stack[len(it.stack)-1]
		childIndex := it.index[len(it.index)-1]
		// I think in this algorithm, list info doesn't matter.
		info, subs, err := DissectListResult(current)
		if err != nil {
			// TODO: Should you just return the error?
			return false, nil
		}
		// If there are unprocessed children, or the current node has a value, return true
		if childIndex < len(subs) {
			return true, nil
		}
		// Pop the current node if all its children are processed
		it.stack = it.stack[:len(it.stack)-1]
		it.index = it.index[:len(it.index)-1]
    if info == nil {
      continue
    }
	}
	return false, nil
}

func (it *ListResultIterator) Next() (*ListResultData_listInfo, error) {
	for len(it.stack) > 0 {
		current := it.stack[len(it.stack)-1]
		childIndex := it.index[len(it.index)-1]

		info, subs, err := DissectListResult(current)
		if err != nil {
			return nil, err
		}

		// Process children
		if childIndex < len(subs) {
			it.index[len(it.index)-1]++
			it.stack = append(it.stack, subs[childIndex])
			it.index = append(it.index, 0)
			continue
		}

		// Pop the current node
		it.stack = it.stack[:len(it.stack)-1]
		it.index = it.index[:len(it.index)-1]
    if info == nil {
      continue
    }
		return info, nil
	}
	return nil, nil
}
