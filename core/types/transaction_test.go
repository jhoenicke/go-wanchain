// Copyright 2014 The go-ethereum Authors
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

package types

import (
	"bytes"
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	"math/big"
	"testing"

	"github.com/wanchain/go-wanchain/common"
	"github.com/wanchain/go-wanchain/crypto"
	"github.com/wanchain/go-wanchain/rlp"
)

// The values in those tests are from the Transaction Tests
// at github.com/ethereum/tests.
var (
	emptyTx = NewTransaction(
		0,
		common.HexToAddress("095e7baea6a6c7c4c2dfeb977efac326af552d87"),
		big.NewInt(0), big.NewInt(0), big.NewInt(0),
		nil,
	)

	rightvrsTx, _ = NewTransaction(
		4,
		common.HexToAddress("b94f5374fce5edbc8e2a8697c15331677e6ebf0b"),
		big.NewInt(100),
		big.NewInt(200000),
		big.NewInt(1),
		common.FromHex("5544"),
	).WithSignature(
		NewEIP155Signer(big.NewInt(5201314)),
		// HomesteadSigner{},
		common.Hex2Bytes("477e2517143dd329f2f6f5a2d554f9b51e306ebe11afc046cc528ec067e4fada756767362196f120957ce96dd2d8b664ee69b83bc1bed077d82ed8dfe03bd7f400"),
	)
)

func TestTransactionSigHash(t *testing.T) {
	// var homestead HomesteadSigner
	// @anson
	// make an eip155 signer
	eip155 := NewEIP155Signer(big.NewInt(5201314))

	if eip155.Hash(emptyTx) != common.HexToHash("0x0e9969d6dc776ab0cae59d85e794bc81c06517787c78ea35f78685953c51ec96") {
		t.Errorf("empty transaction hash mismatch, got %x", emptyTx.Hash())
	}

	if eip155.Hash(rightvrsTx) != common.HexToHash("0xf72c163df11788d7db9680b91e8dee7ca2188953404523a0bf8f509d962cc680") {
		t.Errorf("RightVRS transaction hash mismatch, got %x", rightvrsTx.Hash())
	}
}

func TestTransactionEncode(t *testing.T) {
	txb, err := rlp.EncodeToBytes(rightvrsTx)
	if err != nil {
		t.Fatalf("encode error: %v", err)
	}
	should := common.FromHex("f86601040183030d4094b94f5374fce5edbc8e2a8697c15331677e6ebf0b64825544839ebb67a0477e2517143dd329f2f6f5a2d554f9b51e306ebe11afc046cc528ec067e4fadaa0756767362196f120957ce96dd2d8b664ee69b83bc1bed077d82ed8dfe03bd7f4")
	if !bytes.Equal(txb, should) {
		t.Errorf("encoded RLP mismatch, got %x", txb)
	}
}

func decodeTx(data []byte) (*Transaction, error) {
	var tx Transaction
	t, err := &tx, rlp.Decode(bytes.NewReader(data), &tx)
	return t, err
}

func defaultTestKey() (*ecdsa.PrivateKey, common.Address) {
	// key, _ := crypto.HexToECDSA("45a915e4d060149eb4365960e6a7a45f334393093061116b197e3240065ff2d8")
	key, _ := crypto.HexToECDSA("a4369e77024c2ade4994a9345af5c47598c7cfb36c65e8a4a3117519883d9014")
	addr := crypto.PubkeyToAddress(key.PublicKey)
	fmt.Println("addr is: ", addr.Hex())
	return key, addr
}

//0x010683030d40830f4240940000000000000000000000000000000000000064884563918244f40000b901043f8582d700000000000000000000000000000000000000000000000000000000000000400000000000000000000000000000000000000000000000004563918244f40000000000000000000000000000000000000000000000000000000000000000008630783032376338333431303436353463346362663763346435306330313862303734353736653862356364643839623932623130373037373935613362303830333231363032336133613762383761383137343038636232346462333762636163326662646465323466313039623636333933353963643334313939303762396239643438360000000000000000000000000000000000000000000000000000839ebb68a0ef4208b101e4c29d20ffc1b5f56659eb6cde16a96bc10d9fa64215e4e79cd1f6a00ff87b2336958a1c7eb9046dd4a3aae3a54affce7af4b3e84d863fff201559f2

// @anson
// This testing-case is out-of-date due to ring-signature
// func TestRecipientEmpty(t *testing.T) {
// 	_, addr := defaultTestKey()
// 	tx, err := decodeTx(common.Hex2Bytes("f8498080808080011ca09b16de9d5bdee2cf56c28d16275a4da68cd30273e2525f3959f5d62557489921a0372ebd8fb3345f7db7b5a86d42e24d36e983e259b0664ceb8c227ec9af572f3d"))
// 	if err != nil {
// 		t.Error(err)
// 		t.FailNow()
// 	}

// 	from, err := Sender(HomesteadSigner{}, tx)
// 	if err != nil {
// 		t.Error(err)
// 		t.FailNow()
// 	}
// 	if addr != from {
// 		t.Error("derived address doesn't match")
// 	}
// }

// @anson
// This testing-case is out-of-date due to ring-signature
// func TestRecipientNormal(t *testing.T) {
// 	_, addr := defaultTestKey()
// txb, err := rlp.EncodeToBytes(rightvrsTx)
// fmt.Println("txb: ", common.Bytes2Hex(txb))

// tx, err := decodeTx(common.Hex2Bytes("f85d80808094000000000000000000000000000000000000000080011ca0527c0d8f5c63f7b9f41324a7c8a563ee1190bcbf0dac8ab446291bdbf32f5c79a0552c4ef0a09a04395074dab9ed34d3fbfb843c2f2546cc30fe89ec143ca94ca6"))
// 	if err != nil {
// 		t.Error(err)
// 		t.FailNow()
// 	}

// 	// from, err := Sender(HomesteadSigner{}, tx)
// 	from, err := Sender(EIP155Signer{chainId: big.NewInt(5201314)}, tx)
// 	if err != nil {
// 		t.Error(err)
// 		t.FailNow()
// 	}

// 	if addr != from {
// 		t.Error("derived address doesn't match")
// 	}
// }

// Tests that transactions can be correctly sorted according to their price in
// decreasing order, but at the same time with increasing nonces when issued by
// the same account.
func TestTransactionPriceNonceSort(t *testing.T) {
	// Generate a batch of accounts to start with
	keys := make([]*ecdsa.PrivateKey, 25)
	for i := 0; i < len(keys); i++ {
		keys[i], _ = crypto.GenerateKey()
	}

	signer := HomesteadSigner{}
	// Generate a batch of transactions with overlapping values, but shifted nonces
	groups := map[common.Address]Transactions{}
	for start, key := range keys {
		addr := crypto.PubkeyToAddress(key.PublicKey)
		for i := 0; i < 25; i++ {
			tx, _ := SignTx(NewTransaction(uint64(start+i), common.Address{}, big.NewInt(100), big.NewInt(100), big.NewInt(int64(start+i)), nil), signer, key)
			groups[addr] = append(groups[addr], tx)
		}
	}
	// Sort the transactions and cross check the nonce ordering
	txset := NewTransactionsByPriceAndNonce(signer, groups)

	txs := Transactions{}
	for tx := txset.Peek(); tx != nil; tx = txset.Peek() {
		txs = append(txs, tx)
		txset.Shift()
	}
	if len(txs) != 25*25 {
		t.Errorf("expected %d transactions, found %d", 25*25, len(txs))
	}
	for i, txi := range txs {
		fromi, _ := Sender(signer, txi)

		// Make sure the nonce order is valid
		for j, txj := range txs[i+1:] {
			fromj, _ := Sender(signer, txj)

			if fromi == fromj && txi.Nonce() > txj.Nonce() {
				t.Errorf("invalid nonce ordering: tx #%d (A=%x N=%v) < tx #%d (A=%x N=%v)", i, fromi[:4], txi.Nonce(), i+j, fromj[:4], txj.Nonce())
			}
		}
		// Find the previous and next nonce of this account
		prev, next := i-1, i+1
		for j := i - 1; j >= 0; j-- {
			if fromj, _ := Sender(signer, txs[j]); fromi == fromj {
				prev = j
				break
			}
		}
		for j := i + 1; j < len(txs); j++ {
			if fromj, _ := Sender(signer, txs[j]); fromi == fromj {
				next = j
				break
			}
		}
		// Make sure that in between the neighbor nonces, the transaction is correctly positioned price wise
		for j := prev + 1; j < next; j++ {
			fromj, _ := Sender(signer, txs[j])
			if j < i && txs[j].GasPrice().Cmp(txi.GasPrice()) < 0 {
				t.Errorf("invalid gasprice ordering: tx #%d (A=%x P=%v) < tx #%d (A=%x P=%v)", j, fromj[:4], txs[j].GasPrice(), i, fromi[:4], txi.GasPrice())
			}
			if j > i && txs[j].GasPrice().Cmp(txi.GasPrice()) > 0 {
				t.Errorf("invalid gasprice ordering: tx #%d (A=%x P=%v) > tx #%d (A=%x P=%v)", j, fromj[:4], txs[j].GasPrice(), i, fromi[:4], txi.GasPrice())
			}
		}
	}
}

// TestTransactionJSON tests serializing/de-serializing to/from JSON.
func TestTransactionJSON(t *testing.T) {
	key, err := crypto.GenerateKey()
	if err != nil {
		t.Fatalf("could not generate key: %v", err)
	}
	signer := NewEIP155Signer(common.Big1)

	for i := uint64(0); i < 25; i++ {
		var tx *Transaction
		switch i % 2 {
		case 0:
			tx = NewTransaction(i, common.Address{1}, common.Big0, common.Big1, common.Big2, []byte("abcdef"))
		case 1:
			tx = NewContractCreation(i, common.Big0, common.Big1, common.Big2, []byte("abcdef"))
		}

		tx, err := SignTx(tx, signer, key)
		if err != nil {
			t.Fatalf("could not sign transaction: %v", err)
		}

		data, err := json.Marshal(tx)
		if err != nil {
			t.Errorf("json.Marshal failed: %v", err)
		}

		var parsedTx *Transaction
		if err := json.Unmarshal(data, &parsedTx); err != nil {
			t.Errorf("json.Unmarshal failed: %v", err)
		}

		// compare nonce, price, gaslimit, recipient, amount, payload, V, R, S
		if tx.Hash() != parsedTx.Hash() {
			t.Errorf("parsed tx differs from original tx, want %v, got %v", tx, parsedTx)
		}
		if tx.ChainId().Cmp(parsedTx.ChainId()) != 0 {
			t.Errorf("invalid chain id, want %d, got %d", tx.ChainId(), parsedTx.ChainId())
		}
	}
}
