# X.500-Related Golang Libraries

[![teletex Go Reference](https://pkg.go.dev/badge/github.com/Wildboar-Software/x500-go/teletex.svg)](https://pkg.go.dev/github.com/Wildboar-Software/x500-go/teletex)
[![nsap_address Go Reference](https://pkg.go.dev/badge/github.com/Wildboar-Software/x500-go/nsap-address.svg)](https://pkg.go.dev/github.com/Wildboar-Software/x500-go/nsap-address)
[![x500 Go Reference](https://pkg.go.dev/badge/github.com/Wildboar-Software/x500-go/x500.svg)](https://pkg.go.dev/github.com/Wildboar-Software/x500-go/x500)
[![x500-dap-client Go Reference](https://pkg.go.dev/badge/github.com/Wildboar-Software/x500-go/x500-dap-client.svg)](https://pkg.go.dev/github.com/Wildboar-Software/x500-go/x500-dap-client)

This is a workspace for several X.500-directory and X.509 PKI-related
libraries (and perhaps even tools, one day) written in Go. They are:

- `nsap-address`: ITU-T Rec. X.213 NSAP decoding and encoding
- `teletex`: ITU-T Rec. T.61 Teletex string decoding
- `x500`: Data types and related functionality from the ITU-T X.500 series of
  recommendations describing X.500 directories, including decoding and encoding
  using Golang's standard `encoding/asn1`.
- `x500-dap-client`: A fully-featured X.500 Directory Access Protocol (DAP)
  over Internet Directly Mapped (IDM) protocol client, as described in ITU-T
  Recommendation X.519.

## Warnings

Almost all types in `SchemaAdministration.asn1` will decode with errors if the
`UniversalString` alternative is used for the `description` fields. I kind of
had no choice to but to do it like this. Blame Go's `encodings/asn1` module.

## X.500 Directory Implementation

If you are interested in working with X.500 directories, consider checking
out [Meerkat DSA](https://wildboar-software.github.io/directory/), which,
to my knowledge, is one of the two free and open source X.500 directory
implementations out there, written by yours truly.

## How to Publish New Versions

(It wasn't obvious to me how to do this, so I am documenting it here.)
If you upgrade `x500`, create a Git tag `x500/v#.#.#`. Then run
`GOPROXY=proxy.golang.org go list -m github.com/Wildboar-Software/x500-go/x500@v#.#.#`.
This will cause the Go packages index to update... but it will take a half an
hour for it to show up. It seems to take even longer for the search index to
update with tags, text, keywords, etc.

## Developer Notes

`SET OF SEQUENCE` = Just use the `set` tag
`SEQUENCE OF SET` = Not supported by Golang. See: https://github.com/golang/go/issues/27426
`SET OF SET` = I am not sure what to do here.

## TODO

### MVP

- [x] Even higher-level API
- [ ] Use `X500OperationError` (I might not do this...)
- [ ] Fill in critical extensions appropriately
- [x] Documentation
- [x] Reflection API (See Below)

### Future

- [ ] Complete NSAP library
- [ ] Add types defined in newer X.500 specifications
  - I think I will wait until the newest version is released.
- [ ] Support other SASL methods:
  - [ ] `EXTERNAL`
  - [ ] `ANONYMOUS`
  - [ ] `OTP`

