package apollotypes

import (
	"encoding/hex"
	"errors"

	"github.com/SundaeSwap-finance/apollo/serialization/PlutusData"
)

type AikenPlutusJSON struct {
	Preamble struct {
		Title         string `json:"title"`
		Description   string `json:"description"`
		Version       string `json:"version"`
		PlutusVersion string `json:"plutusVersion"`
		License       string `json:"license"`
	} `json:"preamble"`
	Validators []struct {
		Title string `json:"title"`
		Datum struct {
			Title  string `json:"title"`
			Schema struct {
				Ref string `json:"$ref"`
			} `json:"schema"`
		} `json:"datum"`
		Redeemer struct {
			Title  string `json:"title"`
			Schema struct {
				Ref string `json:"$ref"`
			} `json:"schema"`
		} `json:"redeemer"`
		CompiledCode string `json:"compiledCode"`
		Hash         string `json:"hash"`
	} `json:"validators"`
	Definitions struct {
	} `json:"definitions"`
}

func (apj *AikenPlutusJSON) GetScript(name string) (*PlutusData.PlutusV2Script, error) {
	for _, validator := range apj.Validators {
		if validator.Title == name {
			decoded_string, err := hex.DecodeString(validator.CompiledCode)
			if err != nil {
				return nil, err
			}
			p2Script := PlutusData.PlutusV2Script(decoded_string)
			return &p2Script, nil
		}
	}
	return nil, errors.New("script not found")
}
