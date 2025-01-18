package x500

import (
	"context"
	"crypto"
	"crypto/tls"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/asn1"
	"encoding/pem"
	"errors"
	"net"
	"os"
	"sync"
	"testing"
	"time"
  "github.com/Wildboar-Software/x500-go/x500"
)

var sensibleTimeout = time.Duration(5) * time.Second

const keyPem string = `-----BEGIN PRIVATE KEY-----
  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
  yW0tzYNu6jXxgIzaaQ2T/ODhZR2iVd5B
  -----END PRIVATE KEY-----
  `

const certPem string = `-----BEGIN CERTIFICATE-----
  MIIFfzCCA2egAwIBAgIUGmlo7xlQ5Bd08eBgD+MD28XzVMowDQYJKoZIhvcNAQEL
  BQAwTzELMAkGA1UEBhMCVVMxCzAJBgNVBAgMAkZMMQ4wDAYDVQQHDAVUYW1wYTER
  MA8GA1UECgwIV2lsZGJvYXIxEDAOBgNVBAMMB21lZXJrYXQwHhcNMjIwNzI1MjAy
  NDM5WhcNMjQwNzI0MjAyNDM5WjBPMQswCQYDVQQGEwJVUzELMAkGA1UECAwCRkwx
  DjAMBgNVBAcMBVRhbXBhMREwDwYDVQQKDAhXaWxkYm9hcjEQMA4GA1UEAwwHbWVl
  cmthdDCCAiIwDQYJKoZIhvcNAQEBBQADggIPADCCAgoCggIBAPTK0uLlvhAg8e+6
  5o0utIjhmTv9sYlbtOu0FRPBgY81Qfb0JTPmotM8/5inK7cwoUxqiHXO2oc/4i6d
  U7KGHt1K66o/5xtAVjxTGuLI9WiMCT9485s7okxSk9tCpvwZ6i7y7cqxRe06QdDX
  VXiE1f3EdkcV6gbVOfJcFJKatdhCKGPKsxMKtBCTYyJ2Cft7Cy953ceV5rOMYcta
  VqlXr2dOyQYvwsGnZFPpe4mtyLR4wzWb84WIkGGmwFVaFUxWHn9B04l9G1eujGtQ
  PT+UeCpEFs1wjaMAufcDfxrocBJ/VJG/vBOmNCUelnh8B5DywtXD3RJmi7fqJK/z
  T8KgUTr6Ed8wLvnFfWi7iwsryNQv8CEeZeDbLIeDnEPc5XEXXXpcJm7Snf4nCfK/
  ofP2spCP0PI+60DzQac1+GMVmmY5GFX405DBQAwDL5dcJ1RZCVEKFdGuaDYCItvI
  7F9KRScbjUCbqJOD5B/GnmUoH0Xwxy14+LahMWCCKasA+9Y6QuUJ9fxMer3DhZ+U
  oI2Gk1LVZ99Wx2xer1n4/4pIYL0ZYcrBm57lNdJ6eVr8T8ZMgiSAcdkY5iCuu/7X
  e2zc1Km7RNSvJ6skF0YdYISCeImzecR8gNF8QU0h5ginaW2XPdMQsqOwwQaRyGcE
  NDAZCVuSwn3pNbd0gEpR7UG4eq4RAgMBAAGjUzBRMB0GA1UdDgQWBBRyPT6EMjCl
  cf4nE/u535EUzM3H1jAfBgNVHSMEGDAWgBRyPT6EMjClcf4nE/u535EUzM3H1jAP
  BgNVHRMBAf8EBTADAQH/MA0GCSqGSIb3DQEBCwUAA4ICAQDb+icQ06S74Bpc0n8t
  CXub0kNlENPdcVs/g1+nzbD9UgEi2PXuC0u9GLFuoFbe3i5Nn8Z4VlFo6/6rbvK2
  gVc7MD/Nfo4/Ohe/Tv8YPw8klGi4f1h0O4/beqEayyfgXI/pIPUHqGlbD9xp27D9
  f3stIjc7aGcoqX8N5tmamrKRM+/6p8IVKqqw+iVMXeLgTCZ/l1UkneJsRfz69m3z
  YSDxPWn0LaIdFhvUZizkhiDJcPeZ05dlNfhscDoJyz5G8w/c7aGFwD/qYtjelBFP
  hWV+9pfdSed6bnkStN6ztNPluJ7Vm2Mkw0aootc8HFxIfFPhgOElvv6XBMvZ8MzJ
  xtlo6sE8BeCusPX4rU63ZeoadE3NFUyI1BRFyXUg896dXhAVvvm45FWkBJdUcDaL
  F5bDAY41aZlmNi/nVtzwgk/ph6yLSp2NVkvT2wKPaimVcdsaXhlVtrRMLne6QiDm
  1Tql5+gTmQfc2APqNnL3Hp+xsm5DgccOR0hJucf7BWL3evu9gCijGJ5cnGTHviWY
  kczD5TXI9czgiEgEWID6qgCic3ZfNOXyhb5ZFrkVacWX2X76KVYffkW8NCFMkCjh
  nePOV/Y2kXvUoFNpCjcpaG4FgbQyMxC45jRq9znmCBa3y44BWKs+M2+sKk/z1hWT
  JDKTi1FovCNYrJ8YeKqWM+TMQA==
  -----END CERTIFICATE-----
  `

func getDN() x500.DistinguishedName {
	return x500.DistinguishedName{
		[]pkix.AttributeTypeAndValue{
			{
				Type: x500.Id_at_countryName,
				Value: asn1.RawValue{
					Tag:        asn1.TagPrintableString,
					Class:      asn1.ClassUniversal,
					IsCompound: false,
					Bytes:      []byte("US"),
				},
			},
		},
	}
}

func getDNWithManySubs() x500.DistinguishedName {
	return x500.DistinguishedName{
		[]pkix.AttributeTypeAndValue{
			{
				Type: x500.Id_at_countryName,
				Value: asn1.RawValue{
					Tag:        asn1.TagPrintableString,
					Class:      asn1.ClassUniversal,
					IsCompound: false,
					Bytes:      []byte("US"),
				},
			},
		},
		// c=US,st=FL,l=HIL,l=Tampa,l=Westchase
		[]pkix.AttributeTypeAndValue{
			{
				Type: x500.Id_at_stateOrProvinceName,
				Value: asn1.RawValue{
					Tag:        asn1.TagPrintableString,
					Class:      asn1.ClassUniversal,
					IsCompound: false,
					Bytes:      []byte("FL"),
				},
			},
		},
		[]pkix.AttributeTypeAndValue{
			{
				Type: x500.Id_at_localityName,
				Value: asn1.RawValue{
					Tag:        asn1.TagPrintableString,
					Class:      asn1.ClassUniversal,
					IsCompound: false,
					Bytes:      []byte("HIL"),
				},
			},
		},
		[]pkix.AttributeTypeAndValue{
			{
				Type: x500.Id_at_localityName,
				Value: asn1.RawValue{
					Tag:        asn1.TagPrintableString,
					Class:      asn1.ClassUniversal,
					IsCompound: false,
					Bytes:      []byte("Tampa"),
				},
			},
		},
		[]pkix.AttributeTypeAndValue{
			{
				Type: x500.Id_at_localityName,
				Value: asn1.RawValue{
					Tag:        asn1.TagPrintableString,
					Class:      asn1.ClassUniversal,
					IsCompound: false,
					Bytes:      []byte("Westchase"),
				},
			},
		},
	}
}

func getMeerkatDN() x500.DistinguishedName {
	return x500.DistinguishedName{
		[]pkix.AttributeTypeAndValue{
			{
				Type: x500.Id_at_countryName,
				Value: asn1.RawValue{
					Tag:        asn1.TagPrintableString,
					Class:      asn1.ClassUniversal,
					IsCompound: false,
					Bytes:      []byte("US"),
				},
			},
		},
		// c=US,st=FL,l=HIL,l=Tampa,l=Westchase
		[]pkix.AttributeTypeAndValue{
			{
				Type: x500.Id_at_stateOrProvinceName,
				Value: asn1.RawValue{
					Tag:        asn1.TagPrintableString,
					Class:      asn1.ClassUniversal,
					IsCompound: false,
					Bytes:      []byte("FL"),
				},
			},
		},
		[]pkix.AttributeTypeAndValue{
			{
				Type: x500.Id_at_localityName,
				Value: asn1.RawValue{
					Tag:        asn1.TagPrintableString,
					Class:      asn1.ClassUniversal,
					IsCompound: false,
					Bytes:      []byte("Tampa"),
				},
			},
		},
		[]pkix.AttributeTypeAndValue{
			{
				Type: x500.Id_at_organizationName,
				Value: asn1.RawValue{
					Tag:        asn1.TagPrintableString,
					Class:      asn1.ClassUniversal,
					IsCompound: false,
					Bytes:      []byte("Wildboar"),
				},
			},
		},
		[]pkix.AttributeTypeAndValue{
			{
				Type: x500.Id_at_commonName,
				Value: asn1.RawValue{
					Tag:        asn1.TagPrintableString,
					Class:      asn1.ClassUniversal,
					IsCompound: false,
					Bytes:      []byte("meerkat"),
				},
			},
		},
	}
}

func TestReadAnEntry(t *testing.T) {
	conn, err := net.Dial("tcp", "localhost:4632")
	if err != nil {
		t.Error(err.Error())
		return
	}
	errchan := make(chan error)
	idm := IDMClient(conn, &IDMClientConfig{
		StartTLSPolicy: StartTLSNever,
		Errchan:        errchan,
	})
	go func() {
		e := <-errchan
		t.Error(e)
	}()
	ctx, cancel := context.WithTimeout(context.Background(), sensibleTimeout)
	defer cancel()
	_, err = idm.BindAnonymously(ctx)
	if err != nil {
		t.Error(err.Error())
		return
	}
	invokeId := idm.GetNextInvokeId()
	dn := getDN()
	name_bytes, err := asn1.Marshal(dn)
	if err != nil {
		t.Error(err.Error())
		return
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
		t.Error(err.Error())
		return
	}
	iidBytes, err := asn1.Marshal(invokeId)
	if err != nil {
		t.Error(err.Error())
		return
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
		t.Error(err.Error())
		return
	}
	result := x500.ReadResultData{}
	rest, err := asn1.UnmarshalWithParams(outcome.Parameter.FullBytes, &result, "set")
	if err != nil {
		t.Error(err.Error())
		return
	}
	if len(rest) > 0 {
		t.Error(err.Error())
		return
	}
	for _, info := range result.Entry.Information {
		if info.Tag == asn1.TagOID { // AttributeType
			oid := asn1.ObjectIdentifier{}
			rest, err := asn1.Unmarshal(info.FullBytes, &oid)
			if err != nil {
				continue
			}
			if len(rest) > 0 {
				continue
			}
			t.Logf("Attribute Type: %s\n", oid.String())
		} else if info.Tag == asn1.TagSequence { // Attribute
			attr := x500.Attribute{}
			rest, err := asn1.Unmarshal(info.FullBytes, &attr)
			if err != nil {
				continue
			}
			if len(rest) > 0 {
				continue
			}
			t.Logf("Attribute Type: %s\n", attr.Type.String())
			for _, value := range attr.Values {
				str, err := x500.ASN1RawValueToStr(value)
				if err != nil {
					t.Log(err.Error())
					continue
				}
				if len(str) == 0 {
					t.Log("  <empty>")
				} else {
					t.Logf("  %s\n", str)
				}
			}
			for _, vwc := range attr.ValuesWithContext {
				str, err := x500.ASN1RawValueToStr(vwc.Value)
				if err != nil {
					t.Log(err.Error())
					continue
				}
				t.Logf("  %s (Has Contexts)\n", str)
			}
		} else { // Something else
			continue
		}
	}
}

func TestReadAnEntry2(t *testing.T) {
	conn, err := net.Dial("tcp", "localhost:4632")
	if err != nil {
		t.Error(err.Error())
		return
	}
	errchan := make(chan error)
	idm := IDMClient(conn, &IDMClientConfig{
		StartTLSPolicy: StartTLSNever,
		Errchan:        errchan,
	})
	go func() {
		e := <-errchan
		t.Error(e)
	}()
	ctx, cancel := context.WithTimeout(context.Background(), sensibleTimeout)
	defer cancel()
	_, err = idm.BindAnonymously(ctx)
	if err != nil {
		t.Error(err.Error())
		return
	}
	dn := getDN()
	name_bytes, err := asn1.Marshal(dn)
	if err != nil {
		t.Error(err.Error())
		return
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
		t.Error(err.Error())
		return
	}
	for _, info := range res.Entry.Information {
		if info.Tag == asn1.TagOID { // AttributeType
			oid := asn1.ObjectIdentifier{}
			rest, err := asn1.Unmarshal(info.FullBytes, &oid)
			if err != nil {
				continue
			}
			if len(rest) > 0 {
				continue
			}
			t.Logf("Attribute Type: %s\n", oid.String())
		} else if info.Tag == asn1.TagSequence { // Attribute
			attr := x500.Attribute{}
			rest, err := asn1.Unmarshal(info.FullBytes, &attr)
			if err != nil {
				continue
			}
			if len(rest) > 0 {
				continue
			}
			t.Logf("Attribute Type: %s\n", attr.Type.String())
			for _, value := range attr.Values {
				str, err := x500.ASN1RawValueToStr(value)
				if err != nil {
					t.Log(err.Error())
					continue
				}
				if len(str) == 0 {
					t.Log("  <empty>")
				} else {
					t.Logf("  %s\n", str)
				}
			}
			for _, vwc := range attr.ValuesWithContext {
				str, err := x500.ASN1RawValueToStr(vwc.Value)
				if err != nil {
					t.Log(err.Error())
					continue
				}
				t.Logf("  %s (Has Contexts)\n", str)
			}
		} else { // Something else
			continue
		}
	}
}

func TestManySimultaneousReads(t *testing.T) {
	conn, err := net.Dial("tcp", "localhost:4632")
	if err != nil {
		t.Error(err.Error())
		return
	}
	errchan := make(chan error)
	idm := IDMClient(conn, &IDMClientConfig{
		StartTLSPolicy: StartTLSNever,
		Errchan:        errchan,
	})
	go func() {
		e := <-errchan
		t.Error(e)
	}()
	ctx, cancel := context.WithTimeout(context.Background(), sensibleTimeout)
	defer cancel()
	_, err = idm.BindAnonymously(ctx)
	if err != nil {
		t.Error(err.Error())
		return
	}
	dn := getDN()
	name_bytes, err := asn1.Marshal(dn)
	if err != nil {
		t.Error(err.Error())
		return
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
	// Spam many requests all concurrently
	var wg sync.WaitGroup
	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			ctx, cancel = context.WithTimeout(context.Background(), sensibleTimeout)
			defer cancel()
			outcome, _, err := idm.Read(ctx, arg_data)
			if err != nil {
				t.Error(err.Error())
				return
			}
			if outcome.OutcomeType != OP_OUTCOME_RESULT {
				t.Error(errors.New("non-result-received"))
			}
		}()
	}
	wg.Wait()
}

func TestListAnEntry(t *testing.T) {
	conn, err := net.Dial("tcp", "localhost:4632")
	if err != nil {
		t.Error(err.Error())
		return
	}
	errchan := make(chan error)
	idm := IDMClient(conn, &IDMClientConfig{
		StartTLSPolicy: StartTLSNever,
		Errchan:        errchan,
	})
	go func() {
		e := <-errchan
		t.Error(e)
	}()
	ctx, cancel := context.WithTimeout(context.Background(), sensibleTimeout)
	defer cancel()
	_, err = idm.BindAnonymously(ctx)
	if err != nil {
		t.Error(err.Error())
		return
	}
	dn := getDN()
	name_bytes, err := asn1.Marshal(dn)
	if err != nil {
		t.Error(err.Error())
		return
	}
	name := asn1.RawValue{FullBytes: name_bytes}
	arg_data := x500.ListArgumentData{
		Object: asn1.RawValue{
			Tag:        0,
			Class:      asn1.ClassContextSpecific,
			IsCompound: true,
			Bytes:      name.FullBytes,
		},
		SecurityParameters: x500.SecurityParameters{
			// No signing so we get searchInfo instead of uncorrelatedSearchInfo
			Target:          x500.ProtectionRequest_None,
			ErrorProtection: x500.ErrorProtectionRequest_None,
		},
	}
	ctx, cancel = context.WithTimeout(context.Background(), sensibleTimeout)
	defer cancel()
	_, res, err := idm.List(ctx, arg_data)
	if err != nil {
		t.Error(err.Error())
		return
	}
	for _, sub := range res.Subordinates {
		t.Logf("%v\n", sub.Rdn)
	}
}

func TestTLS(t *testing.T) {
	errchan := make(chan error)
	stop := make(chan int)
	t.Cleanup(func() {
		stop <- 1
	})
	go func() {
		// We need to clean up because it will still receive the abandoned operation.
		select {
		case e := <-errchan:
			t.Error(e)
		case <-stop:
			return
		}
	}()
	conn, err := tls.Dial("tcp", "localhost:44632", &tls.Config{
		InsecureSkipVerify: true,
	})
	if err != nil {
		t.Error(err.Error())
		return
	}
	idm := IDMClient(conn, &IDMClientConfig{
		StartTLSPolicy: StartTLSNever,
		Errchan:        errchan,
	})
	ctx, cancel := context.WithTimeout(context.Background(), sensibleTimeout)
	defer cancel()
	_, err = idm.BindAnonymously(ctx)
	if err != nil {
		t.Error(err.Error())
		return
	}
	dn := getDN()
	name_bytes, err := asn1.Marshal(dn)
	if err != nil {
		t.Error(err.Error())
		return
	}
	name := asn1.RawValue{FullBytes: name_bytes}
	arg_data := x500.ListArgumentData{
		Object: asn1.RawValue{
			Tag:        0,
			Class:      asn1.ClassContextSpecific,
			IsCompound: true,
			Bytes:      name.FullBytes,
		},
		SecurityParameters: x500.SecurityParameters{
			// No signing so we get searchInfo instead of uncorrelatedSearchInfo
			Target:          x500.ProtectionRequest_None,
			ErrorProtection: x500.ErrorProtectionRequest_None,
		},
	}
	ctx, cancel = context.WithTimeout(context.Background(), sensibleTimeout)
	defer cancel()
	_, res, err := idm.List(ctx, arg_data)
	if err != nil {
		t.Error(err.Error())
		return
	}
	for _, sub := range res.Subordinates {
		t.Logf("%v\n", sub.Rdn)
	}
}

func TestAbandon(t *testing.T) {
	errchan := make(chan error)
	stop := make(chan int)
	t.Cleanup(func() {
		stop <- 1
	})
	go func() {
		// We need to clean up because it will still receive the abandoned operation.
		select {
		case e := <-errchan:
			t.Error(e)
		case <-stop:
			return
		}
	}()
	conn, err := net.Dial("tcp", "localhost:4632")
	if err != nil {
		t.Error(err.Error())
		return
	}
	idm := IDMClient(conn, &IDMClientConfig{
		StartTLSPolicy: StartTLSNever,
		Errchan:        errchan,
	})
	ctx, cancel := context.WithTimeout(context.Background(), sensibleTimeout)
	defer cancel()
	_, err = idm.BindAnonymously(ctx)
	if err != nil {
		t.Error(err.Error())
		return
	}
	dn := getDNWithManySubs()
	name_bytes, err := asn1.Marshal(dn)
	if err != nil {
		t.Error(err.Error())
		return
	}
	name := asn1.RawValue{FullBytes: name_bytes}
	arg_data := x500.ListArgumentData{
		Object: asn1.RawValue{
			Tag:        0,
			Class:      asn1.ClassContextSpecific,
			IsCompound: true,
			Bytes:      name.FullBytes,
		},
		SecurityParameters: x500.SecurityParameters{
			// No signing so we get searchInfo instead of uncorrelatedSearchInfo
			Target:          x500.ProtectionRequest_None,
			ErrorProtection: x500.ErrorProtectionRequest_None,
		},
	}
	abandon_arg_data := x500.AbandonArgumentData{
		InvokeID: asn1.RawValue{
			Class:      asn1.ClassContextSpecific,
			Tag:        0,
			IsCompound: true,
			Bytes:      []byte{2, 1, byte(1)},
		},
	}
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		ctx, cancel = context.WithTimeout(context.Background(), sensibleTimeout)
		defer cancel()
		outcome, _, err := idm.List(ctx, arg_data)
		if err != nil {
			t.Error(err.Error())
			return
		}
		if outcome.OutcomeType != OP_OUTCOME_ERROR {
			t.Errorf("Did not receive error. Outcome type=%d", outcome.OutcomeType)
			return
		}
		wg.Done()
	}()
	go func() {
		// Abandon MUST run AFTER the list operation.
		time.Sleep(time.Duration(500) * time.Millisecond)
		ctx, cancel = context.WithTimeout(context.Background(), sensibleTimeout)
		defer cancel()
		outcome, _, err := idm.Abandon(ctx, abandon_arg_data)
		if err != nil {
			panic(err)
		}
		if outcome.OutcomeType != OP_OUTCOME_RESULT {
			t.Errorf("Did not receive abort result. Outcome type=%d", outcome.OutcomeType)
		}
		wg.Done()
	}()
	wg.Wait()
}

func TestStartTLS(t *testing.T) {
	errchan := make(chan error)
	stop := make(chan int)
	t.Cleanup(func() {
		stop <- 1
	})
	go func() {
		// We need to clean up because it will still receive the abandoned operation.
		select {
		case e := <-errchan:
			t.Error(e)
		case <-stop:
			return
		}
	}()
	conn, err := net.Dial("tcp", "localhost:4632")
	if err != nil {
		t.Error(err.Error())
		return
	}
	idm := IDMClient(conn, &IDMClientConfig{
		StartTLSPolicy: StartTLSDemand,
		TlsConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
		Errchan: errchan,
	})
	ctx, cancel := context.WithTimeout(context.Background(), sensibleTimeout)
	defer cancel()
	_, err = idm.BindAnonymously(ctx)
	if err != nil {
		t.Error(err.Error())
		return
	}
	dn := getDN()
	name_bytes, err := asn1.Marshal(dn)
	if err != nil {
		t.Error(err.Error())
		return
	}
	name := asn1.RawValue{FullBytes: name_bytes}
	arg_data := x500.ListArgumentData{
		Object: asn1.RawValue{
			Tag:        0,
			Class:      asn1.ClassContextSpecific,
			IsCompound: true,
			Bytes:      name.FullBytes,
		},
		SecurityParameters: x500.SecurityParameters{
			// No signing so we get searchInfo instead of uncorrelatedSearchInfo
			Target:          x500.ProtectionRequest_None,
			ErrorProtection: x500.ErrorProtectionRequest_None,
		},
	}
	ctx, cancel = context.WithTimeout(context.Background(), sensibleTimeout)
	defer cancel()
	_, res, err := idm.List(ctx, arg_data)
	if err != nil {
		t.Error(err.Error())
		return
	}
	for _, sub := range res.Subordinates {
		t.Logf("%v\n", sub.Rdn)
	}
}

func TestIDMv1(t *testing.T) {
	errchan := make(chan error)
	stop := make(chan int)
	t.Cleanup(func() {
		stop <- 1
	})
	go func() {
		// We need to clean up because it will still receive the abandoned operation.
		select {
		case e := <-errchan:
			t.Error(e)
		case <-stop:
			return
		}
	}()
	conn, err := net.Dial("tcp", "localhost:4632")
	if err != nil {
		t.Error(err.Error())
		return
	}
	idm := IDMClient(conn, &IDMClientConfig{
		StartTLSPolicy: StartTLSNever,
		Errchan:        errchan,
		UseIDMv1:       true,
	})
	ctx, cancel := context.WithTimeout(context.Background(), sensibleTimeout)
	defer cancel()
	_, err = idm.BindAnonymously(ctx)
	if err != nil {
		t.Error(err.Error())
		return
	}
	dn := getDN()
	name_bytes, err := asn1.Marshal(dn)
	if err != nil {
		t.Error(err.Error())
		return
	}
	name := asn1.RawValue{FullBytes: name_bytes}
	arg_data := x500.ListArgumentData{
		Object: asn1.RawValue{
			Tag:        0,
			Class:      asn1.ClassContextSpecific,
			IsCompound: true,
			Bytes:      name.FullBytes,
		},
		SecurityParameters: x500.SecurityParameters{
			// No signing so we get searchInfo instead of uncorrelatedSearchInfo
			Target:          x500.ProtectionRequest_None,
			ErrorProtection: x500.ErrorProtectionRequest_None,
		},
	}
	ctx, cancel = context.WithTimeout(context.Background(), sensibleTimeout)
	defer cancel()
	_, res, err := idm.List(ctx, arg_data)
	if err != nil {
		t.Error(err.Error())
		return
	}
	for _, sub := range res.Subordinates {
		t.Logf("%v\n", sub.Rdn)
	}
}

func TestBindTimeout(t *testing.T) {
	errchan := make(chan error)
	stop := make(chan int)
	t.Cleanup(func() {
		stop <- 1
	})
	go func() {
		// We need to clean up because it will still receive the abandoned operation.
		select {
		case e := <-errchan:
			t.Error(e)
		case <-stop:
			return
		}
	}()
	conn, err := net.Dial("tcp", "localhost:4632")
	if err != nil {
		t.Error(err.Error())
		return
	}
	idm := IDMClient(conn, &IDMClientConfig{
		StartTLSPolicy: StartTLSNever,
		Errchan:        errchan,
	})
	impossibleTimeout := time.Duration(1) * time.Microsecond
	ctx, cancel := context.WithTimeout(context.Background(), impossibleTimeout)
	defer cancel()
	_, err = idm.BindAnonymously(ctx)
	if err == nil {
		t.FailNow()
	}
}

func TestRequestTimeout(t *testing.T) {
	errchan := make(chan error)
	stop := make(chan int)
	t.Cleanup(func() {
		stop <- 1
	})
	go func() {
		// We need to clean up because it will still receive the abandoned operation.
		select {
		case e := <-errchan:
			t.Error(e)
		case <-stop:
			return
		}
	}()
	conn, err := net.Dial("tcp", "localhost:4632")
	if err != nil {
		t.Error(err.Error())
		return
	}
	idm := IDMClient(conn, &IDMClientConfig{
		StartTLSPolicy: StartTLSNever,
		Errchan:        errchan,
	})
	ctx, cancel := context.WithTimeout(context.Background(), sensibleTimeout)
	defer cancel()
	_, err = idm.BindAnonymously(ctx)
	if err != nil {
		t.Error(err.Error())
		return
	}
	dn := getDN()
	name_bytes, err := asn1.Marshal(dn)
	if err != nil {
		t.Error(err.Error())
		return
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
	impossibleTimeout := time.Duration(1) * time.Millisecond
	ctx, cancel = context.WithTimeout(context.Background(), impossibleTimeout)
	defer cancel()
	_, _, err = idm.Read(ctx, arg_data)
	if err == nil {
		t.FailNow()
	}
}

func TestUnbind(t *testing.T) {
	conn, err := net.Dial("tcp", "localhost:4632")
	if err != nil {
		t.Error(err.Error())
		return
	}
	errchan := make(chan error)
	idm := IDMClient(conn, &IDMClientConfig{
		StartTLSPolicy: StartTLSNever,
		Errchan:        errchan,
	})
	stop := make(chan int)
	t.Cleanup(func() {
		stop <- 1
	})
	go func() {
		select {
		case e := <-errchan:
			t.Error(e)
		case <-stop:
			return
		}
	}()
	ctx, cancel := context.WithTimeout(context.Background(), sensibleTimeout)
	defer cancel()
	_, err = idm.BindAnonymously(ctx)
	if err != nil {
		t.Error(err.Error())
		return
	}
	ctx, cancel = context.WithTimeout(context.Background(), sensibleTimeout)
	req := X500UnbindRequest{}
	defer cancel()
	_, err = idm.Unbind(ctx, req)
	if err != nil {
		t.FailNow()
	}
}

func TestBindError(t *testing.T) {
	conn, err := net.Dial("tcp", "localhost:4632")
	if err != nil {
		t.Error(err.Error())
		return
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
			Bytes:      []byte{4, 1, 'a'},
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
	if response.OutcomeType != OP_OUTCOME_ERROR {
		t.Logf("Outcome type: %v\n", response.OutcomeType)
		t.FailNow()
	}
}

// TODO: Also test closing after sending the request
func TestSocketClosure1(t *testing.T) {
	conn, err := net.Dial("tcp", "localhost:4632")
	if err != nil {
		t.Error(err.Error())
		return
	}
	errchan := make(chan error)
	idm := IDMClient(conn, &IDMClientConfig{
		StartTLSPolicy: StartTLSNever,
		Errchan:        errchan,
	})
	go func() {
		e := <-errchan
		t.Error(e)
	}()
	ctx, cancel := context.WithTimeout(context.Background(), sensibleTimeout)
	defer cancel()
	response, err := idm.BindAnonymously(ctx)
	if err != nil {
		t.Error(err.Error())
		return
	}
	if response.OutcomeType != OP_OUTCOME_RESULT {
		t.Logf("Outcome type: %v\n", response.OutcomeType)
		t.FailNow()
	}
	dn := getDN()
	name_bytes, err := asn1.Marshal(dn)
	if err != nil {
		t.Error(err.Error())
		return
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
	idm.socket.Close()
	_, _, err = idm.Read(ctx, arg_data)
	if err == nil {
		t.FailNow()
	}
}

// This test differs from the above by closing the socket shortly after sending
// the request.
func TestSocketClosure2(t *testing.T) {
	conn, err := net.Dial("tcp", "localhost:4632")
	if err != nil {
		t.Error(err.Error())
		return
	}
	errchan := make(chan error)
	idm := IDMClient(conn, &IDMClientConfig{
		StartTLSPolicy: StartTLSNever,
		Errchan:        errchan,
	})
	go func() {
		e := <-errchan
		t.Error(e)
	}()
	ctx, cancel := context.WithTimeout(context.Background(), sensibleTimeout)
	defer cancel()
	response, err := idm.BindAnonymously(ctx)
	if err != nil {
		t.Error(err.Error())
		return
	}
	if response.OutcomeType != OP_OUTCOME_RESULT {
		t.Logf("Outcome type: %v\n", response.OutcomeType)
		t.FailNow()
	}
	dn := getDN()
	name_bytes, err := asn1.Marshal(dn)
	if err != nil {
		t.Error(err.Error())
		return
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
	go func() {
		<-time.After(time.Duration(5) * time.Millisecond)
		idm.socket.Close()
	}()
	outcome, _, err := idm.Read(ctx, arg_data)
	if err == nil {
		t.FailNow()
	}
	// We don't want the request to timeout. We want it to abort immediately
	// if the socket closes. So a timeout should be considered a failure.
	// Since a timeout is carried by the context object, we just check if the
	// error is the same as that of the
	if err == ctx.Err() {
		t.Error("Failed due to timeout, but should have failed due to socket closure")
		return
	}
	if outcome.OutcomeType != OP_OUTCOME_FAILURE {
		t.Errorf("Outcome type should have been failure, but it was %d", outcome.OutcomeType)
		return
	}
	if !errors.Is(err, net.ErrClosed) {
		t.FailNow()
	}
}

func TestSocketClosure3(t *testing.T) {
	conn, err := net.Dial("tcp", "localhost:4632")
	if err != nil {
		t.Error(err.Error())
		return
	}
	errchan := make(chan error)
	idm := IDMClient(conn, &IDMClientConfig{
		StartTLSPolicy: StartTLSNever,
		Errchan:        errchan,
	})
	go func() {
		e := <-errchan
		t.Error(e)
	}()
	ctx, cancel := context.WithTimeout(context.Background(), sensibleTimeout)
	defer cancel()
	idm.socket.Close()
	_, err = idm.BindAnonymously(ctx)
	if err == nil {
		t.Error("bind should have failed, since the socket was closed")
		return
	}
	if !errors.Is(err, net.ErrClosed) {
		t.Errorf("bind should have failed due to socket being closed, but instead got %v", err)
		return
	}
}

func TestSocketClosure4(t *testing.T) {
	t.SkipNow() // This test is so flaky and the timing is hard to get right, but it works.
	conn, err := net.Dial("tcp", "localhost:4632")
	if err != nil {
		t.Error(err.Error())
		return
	}
	errchan := make(chan error)
	idm := IDMClient(conn, &IDMClientConfig{
		StartTLSPolicy: StartTLSNever,
		Errchan:        errchan,
	})
	// go func() {
	// 	e := <-errchan
	// 	t.Error(e)
	// }()
	ctx, cancel := context.WithTimeout(context.Background(), sensibleTimeout)
	defer cancel()
	go func() {
		<-time.After(time.Duration(1) * time.Second)
		idm.socket.Close()
	}()
	outcome, err := idm.BindAnonymously(ctx)
	if err == nil {
		t.FailNow()
	}
	// We don't want the request to timeout. We want it to abort immediately
	// if the socket closes. So a timeout should be considered a failure.
	// Since a timeout is carried by the context object, we just check if the
	// error is the same as that of the
	if err == ctx.Err() {
		t.Error("Failed due to timeout, but should have failed due to socket closure")
		return
	}
	if outcome.OutcomeType != OP_OUTCOME_FAILURE {
		t.Errorf("Outcome type should have been failure, but it was %d", outcome.OutcomeType)
		return
	}
	if !errors.Is(err, net.ErrClosed) {
		t.FailNow()
	}
}

func TestSocketClosure5(t *testing.T) {
	errchan := make(chan error)
	stop := make(chan int)
	t.Cleanup(func() {
		stop <- 1
	})
	go func() {
		// We need to clean up because it will still receive the abandoned operation.
		select {
		case e := <-errchan:
			t.Error(e)
		case <-stop:
			return
		}
	}()
	conn, err := net.Dial("tcp", "localhost:4632")
	if err != nil {
		t.Error(err.Error())
		return
	}
	idm := IDMClient(conn, &IDMClientConfig{
		StartTLSPolicy: StartTLSDemand,
		TlsConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
		Errchan: errchan,
	})
	ctx, cancel := context.WithTimeout(context.Background(), sensibleTimeout)
	defer cancel()
	idm.socket.Close()
	_, err = idm.BindAnonymously(ctx)
	if err == nil {
		t.FailNow()
	}
	if !errors.Is(err, net.ErrClosed) {
		t.FailNow()
	}
}

func TestSocketClosure6(t *testing.T) {
	errchan := make(chan error)
	stop := make(chan int)
	t.Cleanup(func() {
		stop <- 1
	})
	go func() {
		// We need to clean up because it will still receive the abandoned operation.
		select {
		case e := <-errchan:
			t.Error(e)
		case <-stop:
			return
		}
	}()
	conn, err := net.Dial("tcp", "localhost:4632")
	if err != nil {
		t.Error(err.Error())
		return
	}
	idm := IDMClient(conn, &IDMClientConfig{
		StartTLSPolicy: StartTLSDemand,
		TlsConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
		Errchan: errchan,
	})
	ctx, cancel := context.WithTimeout(context.Background(), sensibleTimeout)
	defer cancel()
	go func() {
		// Give it just enough time for the request to be sent.
		<-time.After(time.Duration(1) * time.Millisecond)
		idm.socket.Close()
	}()
	_, err = idm.BindAnonymously(ctx)
	if err == nil {
		t.FailNow()
	}
	if !errors.Is(err, net.ErrClosed) {
		t.FailNow()
	}
}

// If the library user makes a mistake in tagging an asn1.RawValue-typed field,
// does this library automatically correct that mistake correctly?
// Search for "This is the mistake" to find the mistake.
func TestTagCorrection(t *testing.T) {
	conn, err := net.Dial("tcp", "localhost:4632")
	if err != nil {
		t.Error(err.Error())
		return
	}
	errchan := make(chan error)
	idm := IDMClient(conn, &IDMClientConfig{
		StartTLSPolicy: StartTLSNever,
		Errchan:        errchan,
	})
	go func() {
		e := <-errchan
		t.Error(e)
	}()
	ctx, cancel := context.WithTimeout(context.Background(), sensibleTimeout)
	defer cancel()
	_, err = idm.BindAnonymously(ctx)
	if err != nil {
		t.Error(err.Error())
		return
	}
	dn := getDN()
	name_bytes, err := asn1.Marshal(dn)
	if err != nil {
		t.Error(err.Error())
		return
	}
	name := asn1.RawValue{FullBytes: name_bytes}
	arg_data := x500.ReadArgumentData{
		Object: name, // This is the mistake.
		SecurityParameters: x500.SecurityParameters{
			Target:          idm.ResultsSigning,
			ErrorProtection: idm.ErrorSigning,
		},
	}
	ctx, cancel = context.WithTimeout(context.Background(), sensibleTimeout)
	defer cancel()
	outcome, _, err := idm.Read(ctx, arg_data)
	if err != nil {
		t.Error(err.Error())
		return
	}
	if outcome.OutcomeType != OP_OUTCOME_RESULT {
		t.Error("Non-result received")
		t.FailNow()
	}
}

func TestReadSimple(t *testing.T) {
	conn, err := net.Dial("tcp", "localhost:4632")
	if err != nil {
		t.Error(err.Error())
		return
	}
	errchan := make(chan error)
	idm := IDMClient(conn, &IDMClientConfig{
		StartTLSPolicy: StartTLSNever,
		Errchan:        errchan,
	})
	go func() {
		e := <-errchan
		t.Error(e)
	}()
	ctx, cancel := context.WithTimeout(context.Background(), sensibleTimeout)
	defer cancel()
	_, err = idm.BindAnonymously(ctx)
	if err != nil {
		t.Error(err.Error())
		return
	}
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
		if info.Tag == asn1.TagSequence { // Attribute
			attr := x500.Attribute{}
			rest, err := asn1.Unmarshal(info.FullBytes, &attr)
			if err != nil {
				continue
			}
			if len(rest) > 0 {
				continue
			}
			if !attr.Type.Equal(x500.Id_at_countryName) {
				t.FailNow()
			}
			t.Logf("Attribute Type: %s\n", attr.Type.String())
			for _, value := range attr.Values {
				str, err := x500.ASN1RawValueToStr(value)
				if err != nil {
					t.Log(err.Error())
					continue
				}
				if len(str) == 0 {
					t.Log("  <empty>")
				} else {
					t.Logf("  %s\n", str)
				}
			}
			for _, vwc := range attr.ValuesWithContext {
				str, err := x500.ASN1RawValueToStr(vwc.Value)
				if err != nil {
					t.Log(err.Error())
					continue
				}
				t.Logf("  %s (Has Contexts)\n", str)
			}
		} else { // Something else
			t.FailNow()
		}
	}
}

func TestListSimple(t *testing.T) {
	conn, err := net.Dial("tcp", "localhost:4632")
	if err != nil {
		t.Error(err.Error())
		return
	}
	errchan := make(chan error)
	idm := IDMClient(conn, &IDMClientConfig{
		StartTLSPolicy: StartTLSNever,
		Errchan:        errchan,
	})
	go func() {
		e := <-errchan
		t.Error(e)
	}()
	ctx, cancel := context.WithTimeout(context.Background(), sensibleTimeout)
	defer cancel()
	_, err = idm.BindAnonymously(ctx)
	if err != nil {
		t.Error(err.Error())
		return
	}
	dn := getDN()
	ctx, cancel = context.WithTimeout(context.Background(), sensibleTimeout)
	defer cancel()
	outcome, _, err := idm.ListByDN(ctx, dn, 0)
	if err != nil {
		t.Error(err.Error())
		return
	}
	if outcome.OutcomeType != OP_OUTCOME_RESULT {
		t.FailNow()
	}
}

func TestAddEntrySimple(t *testing.T) {
	conn, err := net.Dial("tcp", "localhost:4632")
	if err != nil {
		t.Error(err.Error())
		return
	}
	errchan := make(chan error)
	idm := IDMClient(conn, &IDMClientConfig{
		StartTLSPolicy: StartTLSNever,
		Errchan:        errchan,
	})
	go func() {
		e := <-errchan
		t.Error(e)
	}()
	ctx, cancel := context.WithTimeout(context.Background(), sensibleTimeout)
	defer cancel()
	_, err = idm.BindAnonymously(ctx)
	if err != nil {
		t.Error(err.Error())
		return
	}
	dn := getDNWithManySubs()
	rdn := []pkix.AttributeTypeAndValue{
		{
			Type: x500.Id_at_commonName,
			Value: asn1.RawValue{
				Class:      asn1.ClassUniversal,
				Tag:        asn1.TagPrintableString,
				IsCompound: false,
				Bytes:      []byte("sillybilly"),
			},
		},
	}
	dn = append(dn, rdn)
	ctx, cancel = context.WithTimeout(context.Background(), sensibleTimeout)
	defer cancel()

	attrs := make([]x500.Attribute, 0)
	objectClassAttr := x500.Attribute{
		Type: x500.Id_at_objectClass,
		Values: []asn1.RawValue{
			{
				Class:      asn1.ClassUniversal,
				Tag:        asn1.TagOID,
				IsCompound: false,
				Bytes:      []byte{0x55, 6, 6}, // (id-oc-person / 2.5.6.6)
			},
		},
	}
	commonNameAttr := x500.Attribute{
		Type: x500.Id_at_commonName,
		Values: []asn1.RawValue{
			{
				Class:      asn1.ClassUniversal,
				Tag:        asn1.TagPrintableString,
				IsCompound: false,
				Bytes:      []byte("sillybilly"),
			},
		},
	}
	surnameAttr := x500.Attribute{
		Type: x500.Id_at_surname,
		Values: []asn1.RawValue{
			{
				Class:      asn1.ClassUniversal,
				Tag:        asn1.TagPrintableString,
				IsCompound: false,
				Bytes:      []byte("billy"),
			},
		},
	}
	attrs = append(attrs, objectClassAttr)
	attrs = append(attrs, commonNameAttr)
	attrs = append(attrs, surnameAttr)

	outcome, _, err := idm.AddEntrySimple(ctx, dn, attrs)
	if err != nil {
		t.Error(err.Error())
		return
	}
	// We probably don't have permission to do this. I just want to make
	// sure the request is well-formed.
	if outcome.OutcomeType == OP_OUTCOME_ABORT || outcome.OutcomeType == OP_OUTCOME_FAILURE || outcome.OutcomeType == OP_OUTCOME_REJECT {
		t.FailNow()
	}
}

// This test will delete the entry created from the prior test, if it succeeded.
func TestRemoveEntrySimple(t *testing.T) {
	conn, err := net.Dial("tcp", "localhost:4632")
	if err != nil {
		t.Error(err.Error())
		return
	}
	errchan := make(chan error)
	idm := IDMClient(conn, &IDMClientConfig{
		StartTLSPolicy: StartTLSNever,
		Errchan:        errchan,
	})
	go func() {
		e := <-errchan
		t.Error(e)
	}()
	ctx, cancel := context.WithTimeout(context.Background(), sensibleTimeout)
	defer cancel()
	_, err = idm.BindAnonymously(ctx)
	if err != nil {
		t.Error(err.Error())
		return
	}
	dn := getDNWithManySubs()
	rdn := []pkix.AttributeTypeAndValue{
		{
			Type: x500.Id_at_commonName,
			Value: asn1.RawValue{
				Class:      asn1.ClassUniversal,
				Tag:        asn1.TagPrintableString,
				IsCompound: false,
				Bytes:      []byte("sillybilly"),
			},
		},
	}
	dn = append(dn, rdn)
	ctx, cancel = context.WithTimeout(context.Background(), sensibleTimeout)
	defer cancel()
	outcome, _, err := idm.RemoveEntryByDN(ctx, dn)
	if err != nil {
		t.Error(err.Error())
		return
	}
	// We probably don't have permission to do this. I just want to make
	// sure the request is well-formed.
	if outcome.OutcomeType == OP_OUTCOME_ABORT || outcome.OutcomeType == OP_OUTCOME_FAILURE || outcome.OutcomeType == OP_OUTCOME_REJECT {
		t.FailNow()
	}
}

func TestChangePasswordSimple(t *testing.T) {
	conn, err := net.Dial("tcp", "localhost:4632")
	if err != nil {
		t.Error(err.Error())
		return
	}
	errchan := make(chan error)
	idm := IDMClient(conn, &IDMClientConfig{
		StartTLSPolicy: StartTLSNever,
		Errchan:        errchan,
	})
	go func() {
		e := <-errchan
		t.Error(e)
	}()
	ctx, cancel := context.WithTimeout(context.Background(), sensibleTimeout)
	defer cancel()
	_, err = idm.BindAnonymously(ctx)
	if err != nil {
		t.Error(err.Error())
		return
	}
	dn := getDN()
	ctx, cancel = context.WithTimeout(context.Background(), sensibleTimeout)
	defer cancel()
	outcome, _, err := idm.ChangePasswordSimple(ctx, dn, "asdf", "zxcv")
	if err != nil {
		t.Error(err.Error())
		return
	}
	// We probably don't have permission to do this. I just want to make
	// sure the request is well-formed.
	if outcome.OutcomeType == OP_OUTCOME_ABORT || outcome.OutcomeType == OP_OUTCOME_FAILURE || outcome.OutcomeType == OP_OUTCOME_REJECT {
		t.FailNow()
	}
}

func TestAdministerPasswordSimple(t *testing.T) {
	conn, err := net.Dial("tcp", "localhost:4632")
	if err != nil {
		t.Error(err.Error())
		return
	}
	errchan := make(chan error)
	idm := IDMClient(conn, &IDMClientConfig{
		StartTLSPolicy: StartTLSNever,
		Errchan:        errchan,
	})
	go func() {
		e := <-errchan
		t.Error(e)
	}()
	ctx, cancel := context.WithTimeout(context.Background(), sensibleTimeout)
	defer cancel()
	_, err = idm.BindAnonymously(ctx)
	if err != nil {
		t.Error(err.Error())
		return
	}
	dn := getDN()
	ctx, cancel = context.WithTimeout(context.Background(), sensibleTimeout)
	defer cancel()
	outcome, _, err := idm.AdministerPasswordSimple(ctx, dn, "qwer")
	if err != nil {
		t.Error(err.Error())
		return
	}
	// We probably don't have permission to do this. I just want to make
	// sure the request is well-formed.
	if outcome.OutcomeType == OP_OUTCOME_ABORT || outcome.OutcomeType == OP_OUTCOME_FAILURE || outcome.OutcomeType == OP_OUTCOME_REJECT {
		t.FailNow()
	}
}

func TestStrongAuth(t *testing.T) {
	conn, err := net.Dial("tcp", "localhost:4632")
	if err != nil {
		t.Error(err.Error())
		return
	}
	errchan := make(chan error)

	go func() {
		e := <-errchan
		t.Error(e)
	}()
	ctx, cancel := context.WithTimeout(context.Background(), sensibleTimeout)
	defer cancel()

	// Decode the PEM block
	block, _ := pem.Decode([]byte(certPem))
	if block == nil || block.Type != "CERTIFICATE" {
		t.Error("Failed to decode PEM block containing the certificate")
		return
	}

	// Parse the certificate
	cert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		t.Errorf("Error parsing certificate: %v\n", err)
		return
	}

	// Decode the PEM block
	keyBlock, _ := pem.Decode([]byte(keyPem))
	if keyBlock == nil || keyBlock.Type != "PRIVATE KEY" {
		t.Error("Failed to decode PEM block containing the private key")
		return
	}

	// Parse the certificate
	privKey, err := x509.ParsePKCS8PrivateKey(keyBlock.Bytes)
	if err != nil {
		t.Errorf("Error parsing certificate: %v\n", err)
		return
	}

	signer, is_signer := privKey.(crypto.Signer)
	if !is_signer {
		t.Errorf("not an signing key")
		return
	}

	idm := IDMClient(conn, &IDMClientConfig{
		StartTLSPolicy: StartTLSNever,
		Errchan:        errchan,
		SigningCert: &x500.CertificationPath{
			UserCertificate:   *cert,
			TheCACertificates: make([]x500.CertificatePair, 0),
		},
		SigningKey: &signer,
	})

	// C = US, ST = FL, L = Tampa, O = Wildboar, CN = meerkat
	dn := getMeerkatDN()
	_, err = idm.BindStrongly(ctx, dn, dn, nil)

	if err != nil {
		t.Error(err.Error())
		return
	}
}

func TestInterfaceImplementation(t *testing.T) {
	conn, err := net.Dial("tcp", "localhost:4632")
	if err != nil {
		t.Error(err.Error())
		return
	}
	errchan := make(chan error)
	var idm interface{} = IDMClient(conn, &IDMClientConfig{
		StartTLSPolicy: StartTLSNever,
		Errchan:        errchan,
	})
	_, ok1 := idm.(RemoteOperationServiceElement)
	_, ok2 := idm.(DirectoryAccessClient)
	if !ok1 {
		t.Error("IDM does not implement RemoteOperationServiceElement")
		return
	}
	if !ok2 {
		t.Error("IDM does not implement DirectoryAccessClient")
		return
	}
}

func TestSignedRequest(t *testing.T) {
	conn, err := net.Dial("tcp", "localhost:4632")
	if err != nil {
		t.Error(err.Error())
		return
	}
	errchan := make(chan error)


	// Decode the PEM block
	block, _ := pem.Decode([]byte(certPem))
	if block == nil || block.Type != "CERTIFICATE" {
		t.Error("Failed to decode PEM block containing the certificate")
		return
	}

	// Parse the certificate
	cert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		t.Errorf("Error parsing certificate: %v\n", err)
		return
	}

	// Decode the PEM block
	keyBlock, _ := pem.Decode([]byte(keyPem))
	if keyBlock == nil || keyBlock.Type != "PRIVATE KEY" {
		t.Error("Failed to decode PEM block containing the private key")
		return
	}

	// Parse the certificate
	privKey, err := x509.ParsePKCS8PrivateKey(keyBlock.Bytes)
	if err != nil {
		t.Errorf("Error parsing certificate: %v\n", err)
		return
	}

	signer, is_signer := privKey.(crypto.Signer)
	if !is_signer {
		t.Errorf("not an signing key")
		return
	}

	idm := IDMClient(conn, &IDMClientConfig{
		StartTLSPolicy: StartTLSNever,
		Errchan:        errchan,
		SigningCert: &x500.CertificationPath{
			UserCertificate:   *cert,
			TheCACertificates: make([]x500.CertificatePair, 0),
		},
		SigningKey: &signer,
	})
	go func() {
		e := <-errchan
		t.Error(e)
	}()
	ctx, cancel := context.WithTimeout(context.Background(), sensibleTimeout)
	defer cancel()
	_, err = idm.BindAnonymously(ctx)
	if err != nil {
		t.Error(err.Error())
		return
	}
	dn := getDN()
	name_bytes, err := asn1.Marshal(dn)
	if err != nil {
		t.Error(err.Error())
		return
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
		t.Error(err.Error())
		return
	}
	for _, info := range res.Entry.Information {
		if info.Tag == asn1.TagOID { // AttributeType
			oid := asn1.ObjectIdentifier{}
			rest, err := asn1.Unmarshal(info.FullBytes, &oid)
			if err != nil {
				continue
			}
			if len(rest) > 0 {
				continue
			}
			t.Logf("Attribute Type: %s\n", oid.String())
		} else if info.Tag == asn1.TagSequence { // Attribute
			attr := x500.Attribute{}
			rest, err := asn1.Unmarshal(info.FullBytes, &attr)
			if err != nil {
				continue
			}
			if len(rest) > 0 {
				continue
			}
			t.Logf("Attribute Type: %s\n", attr.Type.String())
			for _, value := range attr.Values {
				str, err := x500.ASN1RawValueToStr(value)
				if err != nil {
					t.Log(err.Error())
					continue
				}
				if len(str) == 0 {
					t.Log("  <empty>")
				} else {
					t.Logf("  %s\n", str)
				}
			}
			for _, vwc := range attr.ValuesWithContext {
				str, err := x500.ASN1RawValueToStr(vwc.Value)
				if err != nil {
					t.Log(err.Error())
					continue
				}
				t.Logf("  %s (Has Contexts)\n", str)
			}
		} else { // Something else
			continue
		}
	}
}
