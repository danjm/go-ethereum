// Copyright 2015 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package tests

import (
	// "bytes"
	"fmt"
	"math/big"
	"strings"
	// "log"
	"encoding/json"
	"reflect"
	"errors"
	// "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/params"
	"github.com/ethereum/go-ethereum/rlp"
)

// TransactionTest checks RLP decoding and sender derivation of transactions.

type TransactionTest struct {
	Byzantium struct {
		Sender      string       `json:"sender"`
		Hash        string      	 `json:"hash"`
	} `json:"Byzantium"`
	Constantinople struct {
		Sender      string       `json:"sender"`
		Hash        string      	 `json:"hash"`
	} `json:"Constantinople"`
	EIP150         struct {
		Sender      string       `json:"sender"`
		Hash        string      	 `json:"hash"`
	} `json:"EIP150"`
	EIP158         struct {
		Sender      string       `json:"sender"`
		Hash        string      	 `json:"hash"`
	} `json:"EIP158"`
	Frontier       struct {
		Sender      string       `json:"sender"`
		Hash        string      	 `json:"hash"`
	} `json:"Frontier"`
	Homestead      struct {
		Sender      string       `json:"sender"`
		Hash        string      	 `json:"hash"`
	} `json:"Homestead"`
	RLP            hexutil.Bytes  `json:"rlp"`
}

//go:generate gencodec -type ttTransaction -field-override ttTransactionMarshaling -out gen_tttransaction.go

type ttTransaction struct {
	Sender      hexutil.Bytes       `json:"sender"`
	Hash        hexutil.Bytes       `json:"has"`
}

type ttTransactionMarshaling struct {
	Sender     hexutil.Bytes
	Hash       hexutil.Bytes
}

func (tt *TransactionTest) Run(config *params.ChainConfig) error {
	// log.Print("MainnetChainConfig- ", params.MainnetChainConfig)
	// var cases map[string]interface{}

	// if err := json.Unmarshal(tt, &cases); err != nil {
	//   fmt.Errorf("Unmarshalling JSON failed: %v", err)
	// }
	// fmt.Println(cases)
	var configInterface map[string]interface{}
  conf, _ := json.Marshal(params.MainnetChainConfig)
  json.Unmarshal(conf, &configInterface)
  // log.Print("configInterface- ", configInterface, " | ", reflect.TypeOf(tt))

	var testDataInterface map[string]interface{}
  td, _ := json.Marshal(tt)
  json.Unmarshal(td, &testDataInterface)
  // log.Print("testDataInterface- ", testDataInterface, " | ", reflect.TypeOf(tt))

	var forkName string
	var fieldBlockNumber float64
	testData := reflect.Indirect(reflect.ValueOf(tt))
	for i := 0; i < testData.NumField(); i++ {
		forkName = testData.Type().Field(i).Name
		// log.Print("forkName- ", forkName)
		if forkName != "RLP" {
			if configInterface[strings.ToLower(forkName) + "Block"] == nil {
				fieldBlockNumber = 0
			} else {
				fieldBlockNumber = configInterface[strings.ToLower(forkName) + "Block"].(float64)
			}
			// log.Print("fieldBlockNumber- ", fieldBlockNumber)
			tx := new(types.Transaction)
			// log.Print("tt- ", tt.RLP)
			forkTestExpectation := testDataInterface[forkName].(map[string]interface{})
			// log.Print("forkTestExpectation- ", forkTestExpectation)

			expectedSender := forkTestExpectation["sender"]
			noSenderExpected := forkTestExpectation["sender"] == ""
			// log.Print("####### expectedSender- ", reflect.TypeOf(expectedSender))
			if err := rlp.DecodeBytes(tt.RLP, tx); err != nil {
				if noSenderExpected {
					return nil
				}
				return fmt.Errorf("RLP decoding failed: %v", err)
			}
			// log.Print("tx- ", reflect.TypeOf(tx))
			signer := types.MakeSigner(config, new(big.Int).SetUint64(uint64(fieldBlockNumber)))
			sender, err := types.Sender(signer, tx)
			// log.Print("sender- ", sender, " ", err)
			if err != nil || sender != expectedSender {
				return err
			}
			// err = tt.Transaction.verify(signer, tx)
			if noSenderExpected && err == nil {
				return errors.New("field validations succeeded but should fail")
			}
			if !noSenderExpected && err != nil {
				return fmt.Errorf("field validations failed after RLP decoding: %s", err)
			}
		}
	}
	// if sender != common.BytesToAddress(tt.Sender) {
	// 	return fmt.Errorf("Sender mismatch: got %x, want %x", sender, tt.Sender)
	// }
	// Check decoded fields.
	// log.Print("err--- ", err)
	return nil
}

func (tt *ttTransaction) verify(signer types.Signer, tx *types.Transaction) error {
	// if !bytes.Equal(tx.Data(), tt.Data) {
	// 	return fmt.Errorf("Tx input data mismatch: got %x want %x", tx.Data(), tt.Data)
	// }
	// if tx.Gas() != tt.GasLimit {
	// 	return fmt.Errorf("GasLimit mismatch: got %d, want %d", tx.Gas(), tt.GasLimit)
	// }
	// if tx.GasPrice().Cmp(tt.GasPrice) != 0 {
	// 	return fmt.Errorf("GasPrice mismatch: got %v, want %v", tx.GasPrice(), tt.GasPrice)
	// }
	// if tx.Nonce() != tt.Nonce {
	// 	return fmt.Errorf("Nonce mismatch: got %v, want %v", tx.Nonce(), tt.Nonce)
	// }
	// v, r, s := tx.RawSignatureValues()
	// if r.Cmp(tt.R) != 0 {
	// 	return fmt.Errorf("R mismatch: got %v, want %v", r, tt.R)
	// }
	// if s.Cmp(tt.S) != 0 {
	// 	return fmt.Errorf("S mismatch: got %v, want %v", s, tt.S)
	// }
	// if v.Cmp(tt.V) != 0 {
	// 	return fmt.Errorf("V mismatch: got %v, want %v", v, tt.V)
	// }
	// if tx.To() == nil {
	// 	if tt.To != (common.Address{}) {
	// 		return fmt.Errorf("To mismatch when recipient is nil (contract creation): %x", tt.To)
	// 	}
	// } else if *tx.To() != tt.To {
	// 	return fmt.Errorf("To mismatch: got %x, want %x", *tx.To(), tt.To)
	// }
	// if tx.Value().Cmp(tt.Value) != 0 {
	// 	return fmt.Errorf("Value mismatch: got %x, want %x", tx.Value(), tt.Value)
	// }
	return nil
}
