# X.500 Directory ASN.1 Data Structures in Go

This is a library containing Go equivalents of most of the data structures
described in the International Telecommuncations Union's X.500 series of
specifications for use in X.500 directories. They are decoded and encoded
using Golang's standard library's `encoding/asn1` module.

You can create attribute certificates as described in ITU-T Recommendation
X.509 by using this library. I might provide an example later on, but it's
pretty straight-forward.

## X.500 Directory Implementation

If you are interested in working with X.500 directories, consider checking
out [Meerkat DSA](https://wildboar-software.github.io/directory/), which,
to my knowledge, is one of the two free and open source X.500 directory
implementations out there, written by yours truly.

