package Certificate

import "github.com/SundaeSwap-finance/apollo/serialization"

type StakeCredential struct {
	_          struct{} `cbor:"toarray"`
	_CODE      int      `cbor:",omitempty"`
	Credential serialization.ConstrainedBytes
}

// TODO
type Certificate struct {
	_               struct{} `cbor:"toarray"`
	_CODE           int
	StakeCredential *StakeCredential
}

type Certificates []*Certificate
