package Utils

import (
	"encoding/hex"
	"fmt"
	"log"

	"github.com/SundaeSwap-finance/apollo/serialization"
	"github.com/SundaeSwap-finance/apollo/serialization/TransactionInput"
	"github.com/SundaeSwap-finance/apollo/serialization/TransactionOutput"
	"github.com/SundaeSwap-finance/apollo/serialization/UTxO"
	"github.com/SundaeSwap-finance/apollo/txBuilding/Backend/Base"

	"github.com/Salvionied/cbor/v2"
)

func Contains[T UTxO.Container[any]](container []T, contained T) bool {
	for _, c := range container {
		if c.EqualTo(contained) {
			return true
		}
	}
	return false
}

func MinLovelacePostAlonzo(output TransactionOutput.TransactionOutput, context Base.ChainContext) int64 {
	constantOverhead := 200
	amt := output.GetValue()
	if amt.Coin == 0 {
		amt.Coin = 1_000_000
	}
	tmp_out := TransactionOutput.TransactionOutput{
		IsPostAlonzo: true,
		PostAlonzo: TransactionOutput.TransactionOutputAlonzo{
			Address:   output.GetAddress(),
			Amount:    output.GetValue().ToAlonzoValue(),
			Datum:     output.GetDatumOption(),
			ScriptRef: output.GetScriptRef(),
		},
	}
	encoded, err := cbor.Marshal(tmp_out)
	if err != nil {
		log.Fatal(err)
	}
	return int64((constantOverhead + len(encoded)) * context.GetProtocolParams().GetCoinsPerUtxoByte())
}

func ToCbor(x interface{}) string {
	bytes, err := cbor.Marshal(x)
	if err != nil {
		log.Fatal(err)
	}
	return hex.EncodeToString(bytes)
}

func Fee(context Base.ChainContext, txSize int, steps int64, mem int64, references []TransactionInput.TransactionInput) int64 {
	pm := context.GetProtocolParams()
	refScriptsSize := 0
	for _, input := range references {
		utxo := context.GetUtxoFromRef(hex.EncodeToString(input.TransactionId), input.Index)
		script := utxo.Output.GetScriptRef()
		fmt.Printf("Utils.Fee: %s#%v script: %v\n",
			hex.EncodeToString(input.TransactionId),
			input.Index,
			script != nil)
		if script != nil {
			fmt.Printf("Utils.Fee: script len: %v\n", len(script.Script.Script))
			refScriptsSize += len(script.Script.Script)
		}
	}
	fee := int64(txSize*pm.MinFeeCoefficient+
		pm.MinFeeConstant+
		int(float32(steps)*pm.PriceStep)+
		int(float32(mem)*pm.PriceMem)+
		refScriptsSize*pm.MinFeeReferenceScripts) + 10_000
	return fee
}

func Copy[T serialization.Clonable[T]](input []T) []T {
	res := make([]T, 0)
	for _, value := range input {
		res = append(res, value.Clone())
	}
	return res
}
