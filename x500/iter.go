package x500

import (
	"encoding/asn1"
	"errors"
	"fmt"
)

func dissectResult[T any](r asn1.RawValue) (info *T, signed *SIGNED, subs []ListResult, err error) {
	data := r
	if data.Class == asn1.ClassUniversal && data.Tag == asn1.TagSequence {
		signedResult := SIGNED{}
		rest, err := asn1.Unmarshal(data.FullBytes, &signedResult)
		if err != nil {
			return nil, nil, subs, err
		}
		if len(rest) > 0 {
			return nil, &signedResult, subs, errors.New("trailing bytes")
		}
		signed = &signedResult
		data = signedResult.ToBeSigned
	}
	if data.Class == asn1.ClassContextSpecific && data.Tag == 0 {
		// This is the uncorrelatedListInfo or uncorrelatedSearchInfo
		sets := make([]ListResult, 0)
		rest, err := asn1.UnmarshalWithParams(data.Bytes, &sets, "set")
		if err != nil {
			return nil, signed, subs, err
		}
		if len(rest) > 0 {
			return nil, signed, subs, errors.New("trailing bytes")
		}
		return nil, signed, sets, nil
	}
	if data.Class != asn1.ClassUniversal || data.Tag != asn1.TagSet {
		// This is some other syntax other than listInfo or searchInfo
		errStr := fmt.Sprintf("unrecognized list or search result syntax (class=%+v, tag=%d)", data.Class, data.Tag)
		return nil, signed, subs, errors.New(errStr)
	}
	info = new(T)
	rest, err := asn1.UnmarshalWithParams(data.FullBytes, info, "set")
	if err != nil {
		return nil, signed, subs, err
	}
	if len(rest) > 0 {
		return nil, signed, subs, errors.New("trailing bytes")
	}
	return info, signed, subs, err
}

// Convenience function for breaking a ListResult into the underlying listInfo, its signature, or substituent sets
func DissectListResult(lr ListResult) (info *ListResultData_listInfo, signed *SIGNED, subs []ListResult, err error) {
	return dissectResult[ListResultData_listInfo](lr)
}

// Convenience function for breaking a SearchResult into the underlying searchInfo, its signature, or substituent sets
func DissectSearchResult(sr SearchResult) (info *SearchResultData_searchInfo, signed *SIGNED, subs []SearchResult, err error) {
	return dissectResult[SearchResultData_searchInfo](sr)
}

// Count the number of list result sets and entries
func CountListResult(lr ListResult) (sets int, entries int, err error) {
	iter := NewListIter(lr)
	set, _, err := iter.Next()
	if err != nil {
		return 0, 0, err
	}
	for set != nil {
		sets++
		entries += len(set.Subordinates)
		set, _, err = iter.Next()
		if err != nil {
			return sets, entries, err
		}
	}
	return sets, entries, err
}

// State for iteration over list results
type ListResultIterator struct {
	// Stack to hold nodes for depth-first traversal
	stack []ListResult
	// Tracks the index of the current child for each node in the stack
	index []int
}

// Create a new ListResultIterator
func NewListIter(lr ListResult) *ListResultIterator {
	return &ListResultIterator{
		stack: []ListResult{lr},
		index: []int{0},
	}
}

// Return the next listInfo and its signature (if present)
func (it *ListResultIterator) Next() (*ListResultData_listInfo, *SIGNED, error) {
	for len(it.stack) > 0 {
		current := it.stack[len(it.stack)-1]
		childIndex := it.index[len(it.index)-1]

		info, signed, subs, err := DissectListResult(current)
		if err != nil {
			return nil, signed, err
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
		return info, signed, nil
	}
	return nil, nil, nil
}

// Count the number of search result sets and entries
func CountSearchResult(sr SearchResult) (sets int, entries int, err error) {
	iter := NewSearchIter(sr)
	set, _, err := iter.Next()
	if err != nil {
		return 0, 0, err
	}
	for set != nil {
		sets++
		entries += len(set.Entries)
		set, _, err = iter.Next()
		if err != nil {
			return sets, entries, err
		}
	}
	return sets, entries, err
}

// State for iteration over list results
type SearchResultIterator struct {
	// Stack to hold nodes for depth-first traversal
	stack []SearchResult
	// Tracks the index of the current child for each node in the stack
	index []int
}

// Create a new ListResultIterator
func NewSearchIter(sr SearchResult) *SearchResultIterator {
	return &SearchResultIterator{
		stack: []SearchResult{sr},
		index: []int{0},
	}
}

// Return the next searchInfo and its signature (if present)
func (it *SearchResultIterator) Next() (*SearchResultData_searchInfo, *SIGNED, error) {
	for len(it.stack) > 0 {
		current := it.stack[len(it.stack)-1]
		childIndex := it.index[len(it.index)-1]

		info, signed, subs, err := DissectSearchResult(current)
		if err != nil {
			return nil, signed, err
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
		return info, signed, nil
	}
	return nil, nil, nil
}

