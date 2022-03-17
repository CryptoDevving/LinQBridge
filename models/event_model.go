/*
 * Copyright (C) 2022 The LinQ Authors
 * This file is part of The LinQ library.
 *
 * The  LinQ  is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Lesser General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * The  LinQ  is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Lesser General Public License for more details.
 * You should have received a copy of the GNU Lesser General Public License
 * along with The LinQ.  If not, see <http://www.gnu.org/licenses/>.
 */

package models

import "math/big"

type ECCMLockEvent struct {
	Method   string
	Txid     string
	TxHash   string
	User     string
	Tchain   uint32
	Contract string
	Height   uint64
	Value    []byte
	Fee      uint64
}
type ECCMUnlockEvent struct {
	Method   string
	TxHash   string
	RTxHash  string
	FChainID uint32
	Contract string
	Height   uint64
	Fee      uint64
}
type ProxyLockEvent struct {
	Method        string
	TxHash        string
	FromAddress   string
	FromAssetHash string
	ToChainID     uint32
	ToAssetHash   string
	ToAddress     string
	TokenID       *big.Int
	DstUser       string
}
type ProxyUnlockEvent struct {
	Method      string
	TxHash      string
	ToAssetHash string
	ToAddress   string
	TokenID     *big.Int
}

type SwapLockEvent struct {
	Type          uint64
	TxHash        string
	FromAddress   string
	FromAssetHash string
	ToChainID     uint64
	ToPoolID      uint64
	ToAddress     string
	Amount        *big.Int
	FeeAssetHash  string
	Fee           *big.Int
	ServerID      *big.Int
}

type SwapUnlockEvent struct {
	Type         uint64
	TxHash       string
	ToPoolID     uint64
	InAssetHash  string
	InAmount     *big.Int
	OutAssetHash string
	OutAmount    *big.Int
	ToChainID    uint64
	ToAssetHash  string
	ToAddress    string
}

type SwapEvent struct {
	Type          uint64
	TxHash        string
	FromAddress   string
	FromAssetHash string
	ToChainID     uint64
	ToPoolID      uint64
	ToAddress     string
	Amount        *big.Int
	FeeAssetHash  string
	Fee           *big.Int
	ServerID      *big.Int
}
type UnSwapEvent struct {
	Type         uint64
	TxHash       string
	ToPoolID     uint64
	InAssetHash  string
	InAmount     *big.Int
	OutAssetHash string
	OutAmount    *big.Int
	ToChainID    uint64
	ToAssetHash  string
	ToAddress    string
}
