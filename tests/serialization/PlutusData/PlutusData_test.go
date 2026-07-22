package plutusdata_test

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/big"
	"testing"

	"github.com/Salvionied/cbor/v2"
	"github.com/SundaeSwap-finance/apollo/serialization/PlutusData"
)

// func TestScriptDataHash(t *testing.T) {
// 	unit := new(PlutusData.PlutusData)
// 	redeemer := new(Redeemer.Redeemer)
// 	redeemer_cbor := "840000d87a80821a002fe1831a34801c64"
// 	datum_cbor := "d8799f581c98aeebe3627161246d1ba4444460f683e22e2e8621d10ed16452871a9fd8799fd8799fd8799f581ccab82fc1490cffe87333ebaeacac82b87e365b391c84e559a89fb76affd8799fd8799fd8799f581ca395cdb0f405e410af1219f0bbce853fcc8a04af3bf6db95dd3c6be9ffffffffa140d8799f00a1401a00fbc520ffffd8799fd8799fd8799f581c70e60f3b5ea7153e0acc7a803e4401d44b8ed1bae1c7baaad1a62a72ffd8799fd8799fd8799f581c1e78aae7c90cc36d624f7b3bb6d86b52696dc84e490f343eba89005fffffffffa140d8799f00a1401a0064b540ffffd8799fd8799fd8799f581c98aeebe3627161246d1ba4444460f683e22e2e8621d10ed16452871affd8799fd8799fd8799f581cf604d5a94dbd068f1bce584c0c6101ab56e2efd1c0cb365ecb4fb3fdffffffffa140d8799f00a1401a124aec20ffffffff"
// 	decoded_redeemer, _ := hex.DecodeString(redeemer_cbor)
// 	decoded_datum, _ := hex.DecodeString(datum_cbor)
// 	cbor.Unmarshal(decoded_redeemer, redeemer)
// 	cbor.Unmarshal(decoded_datum, unit)
// 	marshaled_redeemer, _ := cbor.Marshal(redeemer)
// 	marshaled_datum, _ := cbor.Marshal(unit)
// 	if hex.EncodeToString(marshaled_redeemer) != redeemer_cbor {
// 		t.Error("Invalid redeemer marshaling", hex.EncodeToString(marshaled_redeemer), "Expected", redeemer_cbor)
// 	}
// 	if hex.EncodeToString(marshaled_datum) != datum_cbor {
// 		t.Error("Invalid datum marshaling", hex.EncodeToString(marshaled_datum), "Expected", datum_cbor)
// 	}
// 	data_hash := TxBuilder.ScriptDataHash(nil, nil, nil, []Redeemer.Redeemer{*redeemer}, map[string]PlutusData.PlutusData{"l": *unit})
// 	if hex.EncodeToString(data_hash.Payload) != "8e191588a382fbf52c48e8956cca6a13fd5069038cb55bba30a22c310e517888" {
// 		t.Error("Invalid data hash", hex.EncodeToString(data_hash.Payload), "Expected, 8e191588a382fbf52c48e8956cca6a13fd5069038cb55bba30a22c310e517888")
// 	}
// }

// func TestScriptDataHashDatumOnly(t *testing.T) {
// 	unit := PlutusData.PlutusData{PlutusData.PlutusArray, 121, []any{}}
// 	tws := TransactionWitnessSet.TransactionWitnessSet{
// 		Redeemer:   []Redeemer.Redeemer{},
// 		PlutusData: PlutusData.PlutusIndefArray{unit},
// 	}
// 	data_hash := TxBuilder.ScriptDataHash(tws)
// 	if hex.EncodeToString(data_hash.Payload) != "2f50ea2546f8ce020ca45bfcf2abeb02ff18af2283466f888ae489184b3d2d39" {
// 		t.Error("Invalid data hash", hex.EncodeToString(data_hash.Payload), "Expected, 2f50ea2546f8ce020ca45bfcf2abeb02ff18af2283466f888ae489184b3d2d39")
// 	}
// }

func TestSerializeAndDeserializePlutusData(t *testing.T) {
	cborHex := "d8799fd8799fd8799f581c37dce7298152979f0d0ff71fb2d0c759b298ac6fa7bc56b928ffc1bcffd8799fd8799fd8799f581cf68864a338ae8ed81f61114d857cb6a215c8e685aa5c43bc1f879cceffffffffd8799fd8799f581c37dce7298152979f0d0ff71fb2d0c759b298ac6fa7bc56b928ffc1bcffd8799fd8799fd8799f581cf68864a338ae8ed81f61114d857cb6a215c8e685aa5c43bc1f879cceffffffffd87a80d8799fd8799f581c25f0fc240e91bd95dcdaebd2ba7713fc5168ac77234a3d79449fc20c47534f4349455459ff1b00002cc16be02b37ff1a001e84801a001e8480ff"
	decoded_cbor, _ := hex.DecodeString(cborHex)
	var pd PlutusData.PlutusData
	cbor.Unmarshal(decoded_cbor, &pd)
	marshaled, _ := cbor.Marshal(pd)
	if hex.EncodeToString(marshaled) != cborHex {
		t.Error("Invalid marshaling", hex.EncodeToString(marshaled), "Expected", cborHex)
	}
}

// func TestScriptDataHashRedeemerOnlyOnly2(t *testing.T) {
// 	data_hash := TxBuilder.ScriptDataHash(nil, nil, nil, []Redeemer.Redeemer{}, map[string]PlutusData.PlutusData{})
// 	if hex.EncodeToString(data_hash.Payload) != "a88fe2947b8d45d1f8b798e52174202579ecf847b8f17038c7398103df2d27b0" {
// 		t.Error("Invalid data hash", hex.EncodeToString(data_hash.Payload), "Expected, a88fe2947b8d45d1f8b798e52174202579ecf847b8f17038c7398103df2d27b0")
// 	}
// }

// func TestCostModels(t *testing.T) {
// 	final_cm := map[serialization.CustomBytes]PlutusData.CM{{Value: "00"}: PlutusData.PLUTUSV1COSTMODEL}
// 	fmt.Println(final_cm)
// 	new_cbor, err := cbor.Marshal(final_cm)
// 	if err != nil {
// 		t.Error(err)
// 	}
// 	fmt.Println(hex.EncodeToString(new_cbor))
// 	if hex.EncodeToString(new_cbor) != "a141005901d59f1a000302590001011a00060bc719026d00011a000249f01903e800011a000249f018201a0025cea81971f70419744d186419744d186419744d186419744d186419744d186419744d18641864186419744d18641a000249f018201a000249f018201a000249f018201a000249f01903e800011a000249f018201a000249f01903e800081a000242201a00067e2318760001011a000249f01903e800081a000249f01a0001b79818f7011a000249f0192710011a0002155e19052e011903e81a000249f01903e8011a000249f018201a000249f018201a000249f0182001011a000249f0011a000249f0041a000194af18f8011a000194af18f8011a0002377c190556011a0002bdea1901f1011a000249f018201a000249f018201a000249f018201a000249f018201a000249f018201a000249f018201a000242201a00067e23187600010119f04c192bd200011a000249f018201a000242201a00067e2318760001011a000242201a00067e2318760001011a0025cea81971f704001a000141bb041a000249f019138800011a000249f018201a000302590001011a000249f018201a000249f018201a000249f018201a000249f018201a000249f018201a000249f018201a000249f018201a00330da70101ff" {
// 		fmt.Println("THIS", hex.EncodeToString(new_cbor))
// 		t.Error("WRONG SERIALIZATIONOF THE COSTMODEL")
// 	}
// }

//	func TestPlutusScript(t *testing.T) {
//		plutusScript := PlutusData.PlutusV1Script([]byte("test_script"))
//		hash := PlutusData.PlutusScriptHash(plutusScript)
//		if hex.EncodeToString(hash[:]) != "36c198e1a9d05461945c1f1db2ffb927c2dfc26dd01b59ea93b678b2" {
//			t.Error("Invalid script hash", hex.EncodeToString(hash[:]), "Expected, 36c198e1a9d05461945c1f1db2ffb927c2dfc26dd01b59ea93b678b2")
//		}
//	}
func GetMinSwapPlutusData() PlutusData.PlutusData {
	// PkhStruct :=
	SkhStruct := PlutusData.PlutusData{
		PlutusData.PlutusArray,
		121,
		PlutusData.PlutusIndefArray{
			PlutusData.PlutusData{
				PlutusData.PlutusArray,
				121,
				PlutusData.PlutusIndefArray{
					PlutusData.PlutusData{
						PlutusData.PlutusArray,
						121,
						PlutusData.PlutusIndefArray{
							PlutusData.PlutusData{
								PlutusData.PlutusBytes,
								0,
								[]byte{}}},
					},
				},
			},
		},
	}

	pkhStruct := PlutusData.PlutusData{
		PlutusData.PlutusArray,
		121,
		PlutusData.PlutusIndefArray{
			PlutusData.PlutusData{
				PlutusData.PlutusArray,
				121,
				PlutusData.PlutusIndefArray{
					PlutusData.PlutusData{
						PlutusData.PlutusBytes,
						0,
						[]byte{},
					},
				},
			},
			SkhStruct,
		},
	}
	policy_bytes := []byte{}
	asset_bytes := []byte{}
	AssetStruct := PlutusData.PlutusData{
		PlutusDataType: PlutusData.PlutusArray,
		Value: PlutusData.PlutusIndefArray{
			PlutusData.PlutusData{
				PlutusData.PlutusBytes,
				0,
				policy_bytes,
			},
			PlutusData.PlutusData{
				PlutusData.PlutusBytes,
				0,
				asset_bytes,
			},
		},
		TagNr: 121,
	}
	BuyOrderStruct := PlutusData.PlutusData{
		PlutusDataType: PlutusData.PlutusArray,
		Value: PlutusData.PlutusIndefArray{
			AssetStruct,
			PlutusData.PlutusData{
				PlutusData.PlutusInt,
				0,
				0}},
		TagNr: 121,
	}

	Fee := PlutusData.PlutusData{
		PlutusData.PlutusInt,
		0,
		2_000_000,
	}
	Bribe := PlutusData.PlutusData{
		PlutusData.PlutusInt,
		0,
		0,
	}

	FullStruct := PlutusData.PlutusData{
		PlutusDataType: PlutusData.PlutusArray,
		Value: PlutusData.PlutusIndefArray{
			pkhStruct,
			pkhStruct,
			PlutusData.PlutusData{
				PlutusData.PlutusArray,
				122,
				[]PlutusData.PlutusData{},
			},
			BuyOrderStruct,
			Bribe,
			Fee,
		},
		TagNr: 121,
	}

	return FullStruct
}
func TestPlutusDataFromJson(t *testing.T) {
	PlutusJson := `{
		"fields": [
			{
				"constructor": 0,
				"fields": [
					{
						"constructor": 0,
						"fields": [
							{
								"bytes": "18d725dc0ac9223cac9e91946378dd11df46a58686c2e1e5d7f7eff2"
							}
						]
					},
					{
						"constructor": 0,
						"fields": [
							{
								"constructor": 0,
								"fields": [
									{
										"constructor": 0,
										"fields": [
											{
												"bytes": "a99a2b452e240cc8f538936c549d62813bcaba54f8e6334ee9436578"
											}
										]
									}
								]
							}
						]
					}
				]
			},
			{
				"constructor": 0,
				"fields": [
					{
						"constructor": 0,
						"fields": [
							{
								"bytes": "18d725dc0ac9223cac9e91946378dd11df46a58686c2e1e5d7f7eff2"
							}
						]
					},
					{
						"constructor": 0,
						"fields": [
							{
								"constructor": 0,
								"fields": [
									{
										"constructor": 0,
										"fields": [
											{
												"bytes": "a99a2b452e240cc8f538936c549d62813bcaba54f8e6334ee9436578"
											}
										]
									}
								]
							}
						]
					}
				]
			},
			{
				"constructor": 1,
				"fields": []
			},
			{
				"constructor": 0,
				"fields": [
					{
						"constructor": 0,
						"fields": [
							{
								"bytes": ""
							},
							{
								"bytes": ""
							}
						]
					},
					{
						"int": 18781299
					}
				]
			},
			{
				"int": 2000000
			},
			{
				"int": 2000000
			}
		]
	}
	`
	p := PlutusData.PlutusData{}
	err := json.Unmarshal([]byte(PlutusJson), &p)
	if err != nil {
		t.Error(err)
	}
	cborred, err := cbor.Marshal(p)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(hex.EncodeToString(cborred))
	datum := GetMinSwapPlutusData()
	receborred, err := cbor.Marshal(datum)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(hex.EncodeToString(receborred))
	if hex.EncodeToString(cborred) == hex.EncodeToString(receborred) {
		t.Error("Not the same")
	}
	//t.Error("test")

}

func TestRoundTripNestedConstrDatum(t *testing.T) {
	var pd PlutusData.PlutusData
	datumBytes, err := hex.DecodeString("d8799fd8799f581c4a9f827cdfe7036bf395c7d4304ea3fc1050c420282ac68cb8a1a7b5ffd8799fd8799f581c4a9f827cdfe7036bf395c7d4304ea3fc1050c420282ac68cb8a1a7b5ffd8799fd8799fd8799f581c1e376f232be5a1fa620c5a4b109b8587ec3142dd4876b6bd32ad0beeffffffffd87980d8799fd8799f581c4a9f827cdfe7036bf395c7d4304ea3fc1050c420282ac68cb8a1a7b5ffd8799fd8799fd8799f581c1e376f232be5a1fa620c5a4b109b8587ec3142dd4876b6bd32ad0beeffffffffd87980d8799f581cf5808c2c990d86da54bfc97d89cee6efa20cd8461616359478d96b4c58208996cebb9511cf1e6584abd92667e26cf7034bfe805ac618e2e6a14992198873ffd905029f9fd8799fd8799f581cf5808c2c990d86da54bfc97d89cee6efa20cd8461616359478d96b4c58208996cebb9511cf1e6584abd92667e26cf7034bfe805ac618e2e6a14992198873ffd87980ffd8799fd8799f581cf5808c2c990d86da54bfc97d89cee6efa20cd8461616359478d96b4c5820e74c52975908a612d5ce68327040d449aae99f8b463bb6de046a1b23c5713169ffd87a80ffffd87a9f00ff1a00030b8bff1a001e8480d87a80ff")
	if err != nil {
		t.Fatalf("couldn't decode hex: %v", err)
	}
	if err := cbor.Unmarshal(datumBytes, &pd); err != nil {
		t.Fatalf("couldn't decode: %v", err)
	}
	newBytes, err := cbor.Marshal(pd)
	if err != nil {
		t.Fatalf("couldn't encode: %v", err)
	}
	if !bytes.Equal(datumBytes, newBytes) {
		t.Errorf("failed roundtrip:\n got  %s\n want %s", hex.EncodeToString(newBytes), hex.EncodeToString(datumBytes))
	}
}

func TestRoundTripDefiniteDatum(t *testing.T) {
	var pd PlutusData.PlutusData
	datumBytes, err := hex.DecodeString("d879844100d87982d87982d87982d87981581c49ce0fc15732f1bb8c9c82f2224329a49cbb41e81c52e8a7fce5cf98d87a80d87a80d87a801a002625a0d87983d879801903e8d879811903e8")
	if err != nil {
		t.Error("couldn't decode hex")
	}
	if err := cbor.Unmarshal(datumBytes, &pd); err != nil {
		t.Error("couldn't decode")
	}
	newBytes, err := cbor.Marshal(pd)
	if err != nil {
		t.Error("couldn't encode")
	}
	if !bytes.Equal(datumBytes, newBytes) {
		fmt.Println(hex.EncodeToString(datumBytes))
		fmt.Println(hex.EncodeToString(newBytes))
		t.Error("failed roundtrip")
	}
}

// TestBigIntRoundTrip tests that big integers (> uint64 max) are correctly
// marshaled and unmarshaled. This is critical for stableswap SumInvariant values.
func TestBigIntRoundTrip(t *testing.T) {
	// Create a big integer larger than uint64 max (18446744073709551615)
	// This is the SumInvariant value from production: 1000831573897326959589221
	bigVal := new(big.Int)
	bigVal.SetString("1000831573897326959589221", 10)

	// Encode as CBOR
	encoded, err := cbor.Marshal(bigVal)
	if err != nil {
		t.Fatalf("Failed to marshal big int: %v", err)
	}

	// Verify it's encoded as a CBOR bignum (tag 2)
	// c2 = tag 2 (positive bignum), followed by byte string
	if encoded[0] != 0xc2 {
		t.Errorf("Expected CBOR tag 2 (0xc2), got 0x%02x", encoded[0])
	}

	t.Logf("Encoded big int: %s", hex.EncodeToString(encoded))

	// Unmarshal into PlutusData
	var pd PlutusData.PlutusData
	if err := cbor.Unmarshal(encoded, &pd); err != nil {
		t.Fatalf("Failed to unmarshal into PlutusData: %v", err)
	}

	// Verify the type is PlutusBigInt
	if pd.PlutusDataType != PlutusData.PlutusBigInt {
		t.Errorf("Expected PlutusBigInt type, got %v", pd.PlutusDataType)
	}

	// Verify the value is correct
	decodedBigInt, ok := pd.Value.(big.Int)
	if !ok {
		t.Fatalf("Value is not big.Int, got %T", pd.Value)
	}

	if decodedBigInt.Cmp(bigVal) != 0 {
		t.Errorf("Big int value mismatch: got %s, expected %s", decodedBigInt.String(), bigVal.String())
	}

	// Re-encode and verify round-trip
	reencoded, err := cbor.Marshal(&pd)
	if err != nil {
		t.Fatalf("Failed to re-marshal PlutusData: %v", err)
	}

	if !bytes.Equal(encoded, reencoded) {
		t.Errorf("Round-trip failed: original %s, reencoded %s",
			hex.EncodeToString(encoded), hex.EncodeToString(reencoded))
	}

	t.Logf("Successfully round-tripped big int: %s", bigVal.String())
}

// TestNegativeBigIntRoundTrip tests that negative big integers are correctly handled.
// CBOR uses tag 3 for negative bignums.
func TestNegativeBigIntRoundTrip(t *testing.T) {
	// Create a negative big integer larger than int64 min
	bigVal := new(big.Int)
	bigVal.SetString("-1000831573897326959589221", 10)

	// Encode as CBOR
	encoded, err := cbor.Marshal(bigVal)
	if err != nil {
		t.Fatalf("Failed to marshal negative big int: %v", err)
	}

	// Verify it's encoded as a CBOR negative bignum (tag 3)
	// c3 = tag 3 (negative bignum), followed by byte string
	if encoded[0] != 0xc3 {
		t.Errorf("Expected CBOR tag 3 (0xc3), got 0x%02x", encoded[0])
	}

	t.Logf("Encoded negative big int: %s", hex.EncodeToString(encoded))

	// Unmarshal into PlutusData
	var pd PlutusData.PlutusData
	if err := cbor.Unmarshal(encoded, &pd); err != nil {
		t.Fatalf("Failed to unmarshal into PlutusData: %v", err)
	}

	// Verify the type is PlutusBigInt
	if pd.PlutusDataType != PlutusData.PlutusBigInt {
		t.Errorf("Expected PlutusBigInt type, got %v", pd.PlutusDataType)
	}

	// Verify the value is correct
	decodedBigInt, ok := pd.Value.(big.Int)
	if !ok {
		t.Fatalf("Value is not big.Int, got %T", pd.Value)
	}

	if decodedBigInt.Cmp(bigVal) != 0 {
		t.Errorf("Negative big int value mismatch: got %s, expected %s", decodedBigInt.String(), bigVal.String())
	}

	// Re-encode and verify round-trip
	reencoded, err := cbor.Marshal(&pd)
	if err != nil {
		t.Fatalf("Failed to re-marshal PlutusData: %v", err)
	}

	if !bytes.Equal(encoded, reencoded) {
		t.Errorf("Round-trip failed: original %s, reencoded %s",
			hex.EncodeToString(encoded), hex.EncodeToString(reencoded))
	}

	t.Logf("Successfully round-tripped negative big int: %s", bigVal.String())
}

// TestSmallNegativeInt tests that small negative integers (fitting in int64) work correctly.
func TestSmallNegativeInt(t *testing.T) {
	// Small negative number that fits in int64
	val := int64(-12345678)

	// Encode as CBOR
	encoded, err := cbor.Marshal(val)
	if err != nil {
		t.Fatalf("Failed to marshal negative int: %v", err)
	}

	t.Logf("Encoded small negative int: %s", hex.EncodeToString(encoded))

	// Unmarshal into PlutusData
	var pd PlutusData.PlutusData
	if err := cbor.Unmarshal(encoded, &pd); err != nil {
		t.Fatalf("Failed to unmarshal into PlutusData: %v", err)
	}

	t.Logf("PlutusDataType: %v, Value: %v (%T)", pd.PlutusDataType, pd.Value, pd.Value)

	// Re-encode and verify round-trip
	reencoded, err := cbor.Marshal(&pd)
	if err != nil {
		t.Fatalf("Failed to re-marshal PlutusData: %v", err)
	}

	if !bytes.Equal(encoded, reencoded) {
		t.Errorf("Round-trip failed: original %s, reencoded %s",
			hex.EncodeToString(encoded), hex.EncodeToString(reencoded))
	}

	t.Logf("Successfully round-tripped small negative int: %d", val)
}

// TestStablePoolDatumWithLargeSumInvariant tests parsing a real stableswap datum
// with a SumInvariant that exceeds uint64 max.
func TestStablePoolDatumWithLargeSumInvariant(t *testing.T) {
	// This is the datum from production that caused the bug
	// SumInvariant is 1000831573897326959589221 (encoded as c24ad3ef3036c2ae0594f765)
	dataHex := "d8799f581c227d7fdbb88788d4140c997682e51c9d968fc211ef742d7b09e216f89f9f581cc48cbb3d5e57ed56e276bc45f99ab39abe94e6cd7ac39fb402da47ad480014df105553444dff9f581cfe7c786ab321f41c654ef6c1af7b3250a613c24e4213e0425a7ae4564455534441ffff1b000000a6734ea3fd9f0505ff9f0000ffd8799fd8799f581c55bf4118b01e1c794647db9375ffc873e435d737007b2adbc48cdbaaffff009f1a03b14a060000ff1901f4c24ad3ef3036c2ae0594f765d8799fd8799f581c55bf4118b01e1c794647db9375ffc873e435d737007b2adbc48cdbaaffffff"
	data, err := hex.DecodeString(dataHex)
	if err != nil {
		t.Fatalf("Failed to decode hex: %v", err)
	}

	// Unmarshal into PlutusData
	var pd PlutusData.PlutusData
	if err := cbor.Unmarshal(data, &pd); err != nil {
		t.Fatalf("Failed to unmarshal: %v", err)
	}

	// Re-encode
	reencoded, err := cbor.Marshal(&pd)
	if err != nil {
		t.Fatalf("Failed to re-marshal: %v", err)
	}

	// Verify round-trip
	if !bytes.Equal(data, reencoded) {
		t.Logf("Original:  %s", hex.EncodeToString(data))
		t.Logf("Reencoded: %s", hex.EncodeToString(reencoded))
		t.Error("Round-trip failed for stableswap datum with large SumInvariant")
	} else {
		t.Log("Successfully round-tripped stableswap datum with large SumInvariant")
	}
}

// decodeCostModelV1 parses the output of CostModelV1Encoding.MarshalCBOR
// back into the integer sequence it wrapped.
func decodeCostModelV1(t *testing.T, encoded []byte) []int {
	t.Helper()
	var inner []byte
	if err := cbor.Unmarshal(encoded, &inner); err != nil {
		t.Fatalf("outer bytestring unmarshal failed: %v", err)
	}
	if len(inner) < 2 || inner[0] != 0x9f || inner[len(inner)-1] != 0xff {
		t.Fatalf("inner bytes not an indefinite-length CBOR array: %x", inner)
	}
	// Re-wrap as a definite-length array decode by swapping the indefinite
	// markers for a definite-length header computed from the integer count,
	// or just decode directly — the Salvionied/cbor library accepts
	// indefinite arrays.
	var out []int
	if err := cbor.Unmarshal(inner, &out); err != nil {
		t.Fatalf("inner array unmarshal failed: %v", err)
	}
	return out
}

func TestCostModelV1EncodingSmall(t *testing.T) {
	// <24 entries: the outer bytestring header fits in one byte.
	in := PlutusData.CostModel{1, 2, 3}
	got, err := PlutusData.CostModelV1Encoding(in).MarshalCBOR()
	if err != nil {
		t.Fatalf("MarshalCBOR: %v", err)
	}
	// Bytestring of length 5 (9f 01 02 03 ff).
	want, _ := hex.DecodeString("459f010203ff")
	if !bytes.Equal(got, want) {
		t.Fatalf("small case bytes mismatch:\n got  %x\n want %x", got, want)
	}
	if decoded := decodeCostModelV1(t, got); !equalInts(decoded, []int(in)) {
		t.Fatalf("round-trip mismatch: got %v want %v", decoded, []int(in))
	}
}

func TestCostModelV1EncodingMedium(t *testing.T) {
	// 24..255 entries: header is two bytes (0x58 <len>).
	in := make(PlutusData.CostModel, 30)
	for i := range in {
		in[i] = i
	}
	got, err := PlutusData.CostModelV1Encoding(in).MarshalCBOR()
	if err != nil {
		t.Fatalf("MarshalCBOR: %v", err)
	}
	if got[0] != 0x58 {
		t.Fatalf("expected 0x58-prefixed bytestring, got first byte %#x", got[0])
	}
	if decoded := decodeCostModelV1(t, got); !equalInts(decoded, []int(in)) {
		t.Fatalf("round-trip mismatch: got %v want %v", decoded, []int(in))
	}
}

func TestCostModelV1EncodingLarge(t *testing.T) {
	// >=256 entries: header is three bytes (0x59 <len-hi> <len-lo>).
	// The pre-patch CM.MarshalCBOR corrupted the length prefix at this size
	// because it unconditionally overwrote partial[1] with 0x9f; this test
	// guards against a regression to that behaviour.
	in := make(PlutusData.CostModel, 332)
	for i := range in {
		in[i] = 100788 + i
	}
	got, err := PlutusData.CostModelV1Encoding(in).MarshalCBOR()
	if err != nil {
		t.Fatalf("MarshalCBOR: %v", err)
	}
	if got[0] != 0x59 {
		t.Fatalf("expected 0x59-prefixed bytestring for 332-entry cost model, got first byte %#x", got[0])
	}
	if decoded := decodeCostModelV1(t, got); !equalInts(decoded, []int(in)) {
		t.Fatalf("round-trip mismatch: len(got)=%d len(want)=%d", len(decoded), len(in))
	}
}

func TestCostModelV2Passthrough(t *testing.T) {
	// V2 and V3 now both encode as a plain definite-length CBOR array of
	// ints in the order the chain context supplied. This test asserts the
	// exact bytes for a small array and the round-trip for a large one.
	in := PlutusData.CostModel{1, 2, 3}
	got, err := PlutusData.CostModelV2(in).MarshalCBOR()
	if err != nil {
		t.Fatalf("MarshalCBOR: %v", err)
	}
	want, _ := hex.DecodeString("83010203")
	if !bytes.Equal(got, want) {
		t.Fatalf("V2 small bytes mismatch:\n got  %x\n want %x", got, want)
	}

	large := make(PlutusData.CostModel, 332)
	for i := range large {
		large[i] = 100788 + i
	}
	gotLarge, err := PlutusData.CostModelV2(large).MarshalCBOR()
	if err != nil {
		t.Fatalf("MarshalCBOR large: %v", err)
	}
	var round []int
	if err := cbor.Unmarshal(gotLarge, &round); err != nil {
		t.Fatalf("round-trip unmarshal failed: %v", err)
	}
	if !equalInts(round, []int(large)) {
		t.Fatalf("V2 large round-trip mismatch: len(got)=%d len(want)=%d", len(round), len(large))
	}
}

func equalInts(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
