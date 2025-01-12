# X.500-Related Golang Libraries

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

## X.500 Directory Implementation

If you are interested in working with X.500 directories, consider checking
out [Meerkat DSA](https://wildboar-software.github.io/directory/), which,
to my knowledge, is one of the two free and open source X.500 directory
implementations out there, written by yours truly.

## Developer Notes

`SET OF SEQUENCE` = Just use the `set` tag
`SEQUENCE OF SET` = Not supported by Golang. See: https://github.com/golang/go/issues/27426
`SET OF SET` = I am not sure what to do here.

## TODO

### MVP

- [x] Teletex Handling
- [ ] Separate NSAP library
- [x] `DirectoryString(s str)`
- [x] `FromDirectoryString(ds DirectoryString)`
- [ ] Test directory string encoding and decoding
- [ ] Separate DirectoryString library?
- [ ] Use `omitempty`
- [ ] Even higher-level API
- [ ] Change `int64` enums to `int`
- [ ] Define and implement interfaces
  - [ ] `CommonArguments`
  - [ ] `CommonResults`
  - [ ] `AccessPoint`
  - [ ] `AVMPcommonComponents`
  - [ ] `SchemaElement`
- [ ] Use `X500OperationError`
- [ ] List and Search Result Iterator
- [ ] Test signing
- [ ] Documentation

### Future

- [ ] Add types defined in newer X.500 specifications
  - I think I will wait until the newest version is released.
- [ ] Support other SASL methods:
  - [ ] `EXTERNAL`
  - [ ] `ANONYMOUS`
  - [ ] `OTP`
- [ ] Support more `DSAInfo` attributes

