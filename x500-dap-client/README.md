# X.500 Directory Access Protocol (DAP) Client in Go

This is an implementation an X.500 Directory Access Protocol (DAP) client in
Go as described in ITU-T Recommendation X.519. It currently only supports the
use of the Internet Directly-Mapped (IDM) protocol also described in the same
standard, but this library was implemented in such a way as to facilitate easy
implementation of OSI transport later on, if desired.

This library was developed and tested against
[Meerkat DSA](https://wildboar-software.github.io/directory/), which, to my
knowledge, is one of two free and open source X.500 directories. If you are
interested in working with X.500 directories, consider checking it out!

## Usage

This library exposes multiple layers of complexity to the user, which are, in
order of increasing abstraction:

1. The Remote Operation Service Element (ROSE) Layer (implementing `RemoteOperationServiceElement`)
2. The Directory Access Protocol (DAP) Layer (implementing `DirectoryAccessClient`)
3. The Simple Directory Access Protocol Layer (implementing `SimpleDirectoryAccessClient`)
4. The group management interface (implementing `DirectoryGroupClient`)

With increasing abstraction, you lose a little bit of control and possibly
performance too. The simplest, more abstracted interface is designed for
common use cases, like adding a single attribute to an entry named by a
distinguished name, or checking if a person named by a distinguished name is a
member of the group. If you need more specialized features, you may have to use
the lower-level APIs.

You are likely to never need to use the ROSE API, but one interesting
consequence of exposing this is that you could use a different programming
language to parse and create the payloads for directory operations, and just use
this Go library for handling the IDM and ROSE layers.

### Binding

To use these APIs, the first thing you will need to do is bind to the directory
server.

#### Anonymously

```go
conn, err := net.Dial("tcp", "localhost:4632") // replace with your DSA address
if err != nil {
    return err
}
errchan := make(chan error)
idm := x500_dap_client.IDMClient(conn, &x500_dap_client.IDMClientConfig{
    StartTLSPolicy: x500_dap_client.StartTLSNever,
    Errchan:        errchan,
})
go func() {
    e := <-errchan
    fmt.Println(e)
}()
ctx, cancel := context.WithTimeout(context.Background(), sensibleTimeout)
defer cancel()
_, err = idm.BindAnonymously(ctx) // Yes, it is that simple!
if err != nil {
    return err
}
```

#### Simple Credentials (Username and Password)

```go
conn, err := net.Dial("tcp", "localhost:4632") // replace with your DSA address
if err != nil {
    return err
}
idm := IDMClient(conn, &IDMClientConfig{
    StartTLSPolicy: StartTLSNever,
})
ctx, cancel := context.WithTimeout(context.Background(), sensibleTimeout)
defer cancel()
dn := getDN()
simpleCreds := x500.SimpleCredentials{
    Name: dn,
    Password: asn1.RawValue{
        Class:      asn1.ClassContextSpecific,
        Tag:        2,
        IsCompound: true,
        Bytes:      []byte{4, 4, 'a', 's', 'd', 'f'},
    },
}
credsEncoded, err := asn1.Marshal(simpleCreds)
if err != nil {
    t.Error(err)
    return
}
arg := X500AssociateArgument{
    V1: true,
    V2: true,
    Credentials: &asn1.RawValue{
        Class:      asn1.ClassContextSpecific,
        Tag:        0,
        IsCompound: true,
        Bytes:      credsEncoded,
    },
}
response, err := idm.Bind(ctx, arg)
if err != nil {
    t.Error(err.Error())
    return
}
if response.OutcomeType != OP_OUTCOME_RESULT {
    fmt.Printf("Authentication failure: outcome type %v\n", response.OutcomeType)
    return
}
```

#### Binding with Strong Authentication

Strong authentication is fairly simple: just configure a signing certificate and
signing key, then invoke `BindStrongly()`.

```go
keyBlock, _ := pem.Decode([]byte(keyPem))
if keyBlock == nil || keyBlock.Type != "PRIVATE KEY" {
    return errors.New("Failed to decode PEM block containing the private key")
}

// Parse the certificate
privKey, err := x509.ParsePKCS8PrivateKey(keyBlock.Bytes)
if err != nil {
    return err
}

signer, is_signer := privKey.(crypto.Signer)
if !is_signer {
    return errors.New("not an signing key")
}

idm := x500_dap_client.IDMClient(conn, &x500_dap_client.IDMClientConfig{
    StartTLSPolicy: StartTLSNever,
    Errchan:        errchan,
    SigningCert: &x500.CertificationPath{
        UserCertificate:   *cert,
        TheCACertificates: make([]x500.CertificatePair, 0),
    },
    SigningKey: &signer,
})

dn := getDistinguishedNameOfTheDSA()
_, err = idm.BindStrongly(ctx, dn, dn, nil)
if err != nil {
    t.Error(err.Error())
    return
}
```

### ROSE-Layer Interface

Example usage:

```go
invokeId := idm.GetNextInvokeId()
dn := getDN()
name_bytes, err := asn1.Marshal(dn)
if err != nil {
    return err
}
name := asn1.RawValue{FullBytes: name_bytes}
arg_data := x500.ReadArgumentData{
    Object: asn1.RawValue{
        Tag:        0,
        Class:      asn1.ClassContextSpecific,
        IsCompound: true,
        Bytes:      name.FullBytes,
    },
}
arg_bytes, err := asn1.MarshalWithParams(arg_data, "set")
if err != nil {
    return err
}
iidBytes, err := asn1.Marshal(invokeId)
if err != nil {
    return err
}
req := X500Request{
    InvokeId: asn1.RawValue{
        FullBytes: iidBytes,
    },
    OpCode: asn1.RawValue{
        Tag:        asn1.TagInteger,
        Class:      asn1.ClassUniversal,
        IsCompound: false,
        Bytes:      []byte{0x01}, // Read operation
    },
    Argument: asn1.RawValue{FullBytes: arg_bytes},
}
ctx, cancel = context.WithTimeout(context.Background(), sensibleTimeout)
defer cancel()
outcome, err := idm.Request(ctx, req)
if err != nil {
    return err
}
result := x500.ReadResultData{}
rest, err := asn1.UnmarshalWithParams(outcome.Parameter.FullBytes, &result, "set")
if err != nil {
    return err
}
if len(rest) > 0 {
    return err
}
for _, info := range result.Entry.Information {
    // Print out the attributes, use them for access control decisions, etc.
}
```

### Directory Access Protocol (DAP) Interface

The bare-bones read operation API:

```go
dn := getDN()
name_bytes, err := asn1.Marshal(dn)
if err != nil {
    return err
}
name := asn1.RawValue{FullBytes: name_bytes}
arg_data := x500.ReadArgumentData{
    Object: asn1.RawValue{
        Tag:        0,
        Class:      asn1.ClassContextSpecific,
        IsCompound: true,
        Bytes:      name.FullBytes,
    },
    SecurityParameters: x500.SecurityParameters{
        Target:          idm.ResultsSigning,
        ErrorProtection: idm.ErrorSigning,
    },
}
ctx, cancel = context.WithTimeout(context.Background(), sensibleTimeout)
defer cancel()
_, res, err := idm.Read(ctx, arg_data)
if err != nil {
    return err
}
for _, info := range res.Entry.Information {
    // Print out the attributes, use them for access control decisions, etc.
}
```


### Simple Directory Access Protocol Interface

Even simpler read operation API:

```go
dn := getDN()
attrs := make([]asn1.ObjectIdentifier, 1)
attrs[0] = x500.Id_at_countryName
ctx, cancel = context.WithTimeout(context.Background(), sensibleTimeout)
defer cancel()
_, res, err := idm.ReadSimple(ctx, dn, attrs)
if err != nil {
    t.Error(err.Error())
    return
}
for _, info := range res.Entry.Information {
    // Print out the attributes, use them for access control decisions, etc.
}
```

### Group Management

To check if a user is in a group:

```go
member := getPersonDN()
adminsGroup := getGroupDN()
ctx, cancel = context.WithTimeout(context.Background(), sensibleTimeout)
defer cancel()
_, res, err := idm.GroupCheckMember(ctx, adminsGroup, member, nil)
if res != nil && res.Matched {
    // Access granted
}
```

### Unbind

```go
ctx, cancel = context.WithTimeout(context.Background(), sensibleTimeout)
req := x500_dap_client.X500UnbindRequest{}
defer cancel()
_, err = idm.Unbind(ctx, req)
if err != nil {
    return err
}
```

### Socket Closure

To close the underlying TCP or TLS socket, or any other underlying protocol
layers, simply call `CloseTransport()`.

```go
err := idm.CloseTransport()
```

### TLS and StartTLS

Using TLS is straight-forward. Just create a TLS connection, then pass it in
to `IDMClient()`.

```go
conn, err := tls.Dial("tcp", "localhost:44632", &tls.Config{
    InsecureSkipVerify: true,
})
if err != nil {
    return err
}
idm := x500_dap_client.IDMClient(conn, &IDMClientConfig{
    StartTLSPolicy: StartTLSNever,
    Errchan:        errchan,
})
```

By default, if TLS is not used, StartTLS will be _demanded_; the connection will
fail with an error if StartTLS does not succeed. You MUST disable the StartTLS
upgrade if you plan to connect over cleartext. You can do this by setting
`IDMClientConfig.StartTLSPolicy` to either `StartTLSNever` (never perform
StartTLS), or `StartTLSPrefer` (try StartTLS, but keep going if it fails).

If StartTLS is to be used, this will be done when you invoke `.Bind()` or one of
the higher-level APIs that themselves perform a bind operation. This is the one
thing you don't get control over in this library: I had to do it this way for
annoying technical reasons.

### Signing Requests

To produce signed requests, all you have to do is configure a signing key and
certificate like so:

```go
keyBlock, _ := pem.Decode([]byte(keyPem))
if keyBlock == nil || keyBlock.Type != "PRIVATE KEY" {
    return errors.New("Failed to decode PEM block containing the private key")
}

// Parse the certificate
privKey, err := x509.ParsePKCS8PrivateKey(keyBlock.Bytes)
if err != nil {
    return err
}

signer, is_signer := privKey.(crypto.Signer)
if !is_signer {
    return errors.New("not an signing key")
}

idm := x500_dap_client.IDMClient(conn, &x500_dap_client.IDMClientConfig{
    StartTLSPolicy: StartTLSNever,
    Errchan:        errchan,
    SigningCert: &x500.CertificationPath{
        UserCertificate:   *cert,
        TheCACertificates: make([]x500.CertificatePair, 0),
    },
    SigningKey: &signer,
})
```

Any requests generated will silently have signatures applied. Only Ed25519,
ECDSA, and RSA signatures are supported, and all of them only use SHA-256.
There is no support for customizing the hash algorithm right now. If you attempt
to use a key that is designed for some other algorithm, I don't know what will
happen, actually.

If an RSA key is supplied, the `rSASSA-PSS` algorithm will be used, because it
is more secure than the old `RSASSA-PKCS1-v1.5` algorithm. If this is a problem
for you, let me know.

### Signed Results and Errors

This library does not implement any verification of signed results or errors.
You get the full payload back with most APIs and it is expected that you verify
them.

### Error Handling

This library tries to avoid doing any logging to the console; we don't want to
circumvent the logging middleware you've chosen for your application. Instead,
the IDM stack takes an error channel, which you can supply when creating the
stack, and use that to handle errors however you see fit, like so:

```go
errchan := make(chan error)
idm := x500_dap_client.IDMClient(conn, &x500_dap_client.IDMClientConfig{
    StartTLSPolicy: x500_dap_client.StartTLSNever,
    Errchan:        errchan,
})
go func() {
    e := <-errchan
    fmt.Println(e)
}()
```

## Tests

The tests are a great way to see examples for usage.

## Notes for Users

You will need to set `priority` in the service controls. Golang defaults enums
to 0, even if you use the `default:1` tag on a struct member.

The `timeLimit` field of the service controls is populated by the timeout
specified with the `Context` object. If there is no such timeout specified
either in the service controls, or in the context object, the request will not
timeout unless configured to do so at the TCP or TLS layers. The `sizeLimit` and
`attributeSizeLimit` fields get populated automatically with sensible defaults
unless you supply your own values.

This library supports requesting an attribute certificate from the DSA per the
[private extension used by Meerkat DSA](https://wildboar-software.github.io/directory/docs/attr-cert).

Known issues: https://github.com/golang/go/issues/27426
Any `SEQUENCE OF SET` type will fail to be unmarshalled.
