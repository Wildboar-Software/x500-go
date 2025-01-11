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
