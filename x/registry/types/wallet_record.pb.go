// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: mycel/registry/wallet_record.proto

package types

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type WalletRecordType int32

const (
	// BTC 0xxx
	WalletRecordType_BITCOIN_MAINNET WalletRecordType = 0
	WalletRecordType_BITCOIN_TESTNET WalletRecordType = 1
	// EVM 1xxxx
	WalletRecordType_ETHEREUM_MAINNET WalletRecordType = 10000
	WalletRecordType_ETHEREUM_GOERLI  WalletRecordType = 10001
	WalletRecordType_ETHEREUM_SEPOLIA WalletRecordType = 10002
	// Polygon
	WalletRecordType_POLYGON_MAINNET WalletRecordType = 10003
	WalletRecordType_POLYGON_MUMBAI  WalletRecordType = 10004
	// BNB Chain
	WalletRecordType_BNB_MAINNET WalletRecordType = 10005
	WalletRecordType_BNB_TESTNET WalletRecordType = 10006
	// Avalanche
	WalletRecordType_AVALANCHE_CCHAIN WalletRecordType = 10007
	WalletRecordType_AVALANCHE_FUJI   WalletRecordType = 10008
	// Gnosis
	WalletRecordType_GNOSIS_MAINNET WalletRecordType = 10009
	WalletRecordType_GNOSIS_CHIADO  WalletRecordType = 10010
	// Optimism
	WalletRecordType_OPTIMISM_MAINNET WalletRecordType = 10011
	WalletRecordType_OPTIMISM_GOERLI  WalletRecordType = 10012
	// Arbitrum
	WalletRecordType_ARBITRUM_MAINNET WalletRecordType = 10013
	WalletRecordType_ARBITRUM_GOERLI  WalletRecordType = 10014
	//Shardeum
	// SHARDEUM_MAINNET = 10015;
	// SHARDEUM_TESTNET = 10016;
	WalletRecordType_SHARDEUM_BETANET WalletRecordType = 10017
	// MOVE 2xxxx
	WalletRecordType_APTOS_MAINNET WalletRecordType = 20000
	WalletRecordType_APTOS_TESTNET WalletRecordType = 20001
	WalletRecordType_SUI_MAINNET   WalletRecordType = 20002
	WalletRecordType_SUI_TESTNET   WalletRecordType = 20003
	// SOLANA 3xxxx
	WalletRecordType_SOLANA_MAINNET WalletRecordType = 30000
	WalletRecordType_SOLANA_TESTNET WalletRecordType = 30001
)

var WalletRecordType_name = map[int32]string{
	0:     "BITCOIN_MAINNET",
	1:     "BITCOIN_TESTNET",
	10000: "ETHEREUM_MAINNET",
	10001: "ETHEREUM_GOERLI",
	10002: "ETHEREUM_SEPOLIA",
	10003: "POLYGON_MAINNET",
	10004: "POLYGON_MUMBAI",
	10005: "BNB_MAINNET",
	10006: "BNB_TESTNET",
	10007: "AVALANCHE_CCHAIN",
	10008: "AVALANCHE_FUJI",
	10009: "GNOSIS_MAINNET",
	10010: "GNOSIS_CHIADO",
	10011: "OPTIMISM_MAINNET",
	10012: "OPTIMISM_GOERLI",
	10013: "ARBITRUM_MAINNET",
	10014: "ARBITRUM_GOERLI",
	10017: "SHARDEUM_BETANET",
	20000: "APTOS_MAINNET",
	20001: "APTOS_TESTNET",
	20002: "SUI_MAINNET",
	20003: "SUI_TESTNET",
	30000: "SOLANA_MAINNET",
	30001: "SOLANA_TESTNET",
}

var WalletRecordType_value = map[string]int32{
	"BITCOIN_MAINNET":  0,
	"BITCOIN_TESTNET":  1,
	"ETHEREUM_MAINNET": 10000,
	"ETHEREUM_GOERLI":  10001,
	"ETHEREUM_SEPOLIA": 10002,
	"POLYGON_MAINNET":  10003,
	"POLYGON_MUMBAI":   10004,
	"BNB_MAINNET":      10005,
	"BNB_TESTNET":      10006,
	"AVALANCHE_CCHAIN": 10007,
	"AVALANCHE_FUJI":   10008,
	"GNOSIS_MAINNET":   10009,
	"GNOSIS_CHIADO":    10010,
	"OPTIMISM_MAINNET": 10011,
	"OPTIMISM_GOERLI":  10012,
	"ARBITRUM_MAINNET": 10013,
	"ARBITRUM_GOERLI":  10014,
	"SHARDEUM_BETANET": 10017,
	"APTOS_MAINNET":    20000,
	"APTOS_TESTNET":    20001,
	"SUI_MAINNET":      20002,
	"SUI_TESTNET":      20003,
	"SOLANA_MAINNET":   30000,
	"SOLANA_TESTNET":   30001,
}

func (x WalletRecordType) String() string {
	return proto.EnumName(WalletRecordType_name, int32(x))
}

func (WalletRecordType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_163ce144c4dda9d7, []int{0}
}

type WalletAddressFormat int32

const (
	WalletAddressFormat_BITCOIN  WalletAddressFormat = 0
	WalletAddressFormat_ETHEREUM WalletAddressFormat = 1
	WalletAddressFormat_MOVE     WalletAddressFormat = 2
	WalletAddressFormat_SOLANA   WalletAddressFormat = 3
)

var WalletAddressFormat_name = map[int32]string{
	0: "BITCOIN",
	1: "ETHEREUM",
	2: "MOVE",
	3: "SOLANA",
}

var WalletAddressFormat_value = map[string]int32{
	"BITCOIN":  0,
	"ETHEREUM": 1,
	"MOVE":     2,
	"SOLANA":   3,
}

func (x WalletAddressFormat) String() string {
	return proto.EnumName(WalletAddressFormat_name, int32(x))
}

func (WalletAddressFormat) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_163ce144c4dda9d7, []int{1}
}

func init() {
	proto.RegisterEnum("mycel.registry.WalletRecordType", WalletRecordType_name, WalletRecordType_value)
	proto.RegisterEnum("mycel.registry.WalletAddressFormat", WalletAddressFormat_name, WalletAddressFormat_value)
}

func init() {
	proto.RegisterFile("mycel/registry/wallet_record.proto", fileDescriptor_163ce144c4dda9d7)
}

var fileDescriptor_163ce144c4dda9d7 = []byte{
	// 461 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x92, 0xcb, 0x6e, 0xd3, 0x40,
	0x14, 0x86, 0x3d, 0x50, 0x95, 0xea, 0x84, 0xa6, 0xc3, 0xa4, 0x6c, 0xbd, 0x60, 0x59, 0x89, 0x64,
	0xc1, 0x13, 0x8c, 0xdd, 0x49, 0x3c, 0xc8, 0x9e, 0x89, 0x3c, 0x93, 0x22, 0xd8, 0x44, 0x69, 0x62,
	0x95, 0x48, 0x09, 0x8e, 0x1c, 0x23, 0xc8, 0x5b, 0xb4, 0xdc, 0x2f, 0x01, 0x51, 0x78, 0x01, 0x78,
	0x0b, 0x96, 0x5d, 0xb2, 0x44, 0xc9, 0xae, 0x4f, 0x81, 0x32, 0x89, 0xc7, 0x61, 0xe9, 0x4f, 0x9f,
	0xce, 0xfc, 0xe7, 0xf8, 0x87, 0x7b, 0xe3, 0x59, 0x3f, 0x19, 0x35, 0xb2, 0xe4, 0x6c, 0x38, 0xcd,
	0xb3, 0x59, 0xe3, 0x45, 0x6f, 0x34, 0x4a, 0xf2, 0x6e, 0x96, 0xf4, 0xd3, 0x6c, 0x50, 0x9f, 0x64,
	0x69, 0x9e, 0x92, 0xaa, 0x71, 0xea, 0x85, 0x73, 0x74, 0xb1, 0x03, 0xf8, 0x91, 0xf1, 0x62, 0xa3,
	0xe9, 0xd9, 0x24, 0x21, 0x35, 0x38, 0xf0, 0xb8, 0xf6, 0x25, 0x17, 0xdd, 0x88, 0x72, 0x21, 0x98,
	0xc6, 0xce, 0x36, 0xd4, 0x4c, 0xe9, 0x15, 0x44, 0xe4, 0x2e, 0x60, 0xa6, 0x03, 0x16, 0xb3, 0x4e,
	0x64, 0xd5, 0x73, 0x41, 0x0e, 0xe1, 0xc0, 0xe2, 0x96, 0x64, 0x71, 0xc8, 0xf1, 0x85, 0xf8, 0x4f,
	0x56, 0xac, 0x2d, 0x43, 0x4e, 0xf1, 0x2b, 0x23, 0xb7, 0x65, 0xf8, 0xb8, 0x25, 0xcb, 0xd7, 0x5e,
	0x0b, 0x52, 0x83, 0xaa, 0xa5, 0x9d, 0xc8, 0xa3, 0x1c, 0xbf, 0x11, 0x04, 0x43, 0xc5, 0x13, 0x9e,
	0xd5, 0xde, 0x5a, 0x52, 0x24, 0x7a, 0x67, 0x5e, 0xa1, 0x27, 0x34, 0xa4, 0xc2, 0x0f, 0x58, 0xd7,
	0xf7, 0x03, 0xca, 0x05, 0x7e, 0x6f, 0xe6, 0x95, 0xb8, 0xd9, 0x79, 0xc8, 0xf1, 0x07, 0x03, 0x5b,
	0x42, 0x2a, 0xae, 0xec, 0xc8, 0x8f, 0x82, 0x10, 0xd8, 0xdf, 0x40, 0x3f, 0xe0, 0xf4, 0x58, 0xe2,
	0x4f, 0x66, 0xa8, 0x6c, 0x6b, 0x1e, 0x71, 0x55, 0xee, 0xf9, 0xd9, 0x44, 0xb7, 0x78, 0xb3, 0xe7,
	0x7c, 0x9d, 0x20, 0xf6, 0xb8, 0x8e, 0xb7, 0x8e, 0xf2, 0xc5, 0xc8, 0x16, 0x6f, 0xe4, 0xaf, 0x46,
	0x56, 0x01, 0x8d, 0x8f, 0x57, 0x47, 0xf1, 0x98, 0xa6, 0x2b, 0xf9, 0x72, 0x95, 0x6c, 0x9f, 0xb6,
	0xb5, 0x2c, 0x83, 0x7d, 0x9b, 0xa3, 0x12, 0x16, 0xeb, 0x5e, 0xce, 0x11, 0xb9, 0x03, 0x15, 0xd5,
	0xe1, 0xd6, 0xfb, 0x5e, 0xa2, 0xc2, 0xfa, 0x31, 0x47, 0xe4, 0x10, 0xaa, 0x4a, 0x86, 0x54, 0x50,
	0x2b, 0xfe, 0xbc, 0xde, 0xa6, 0x85, 0xfb, 0xeb, 0x1a, 0x1d, 0x35, 0xa1, 0xb6, 0xae, 0x04, 0x1d,
	0x0c, 0xb2, 0x64, 0x3a, 0x6d, 0xa6, 0xd9, 0xb8, 0x97, 0x93, 0x0a, 0xdc, 0xda, 0x14, 0x00, 0x3b,
	0xe4, 0x36, 0xec, 0x15, 0xff, 0x12, 0x23, 0xb2, 0x07, 0x3b, 0x91, 0x3c, 0x61, 0xf8, 0x06, 0x01,
	0xd8, 0x5d, 0x4f, 0xc4, 0x37, 0xbd, 0xe0, 0xf7, 0xc2, 0x45, 0x57, 0x0b, 0x17, 0xfd, 0x5d, 0xb8,
	0xe8, 0x7c, 0xe9, 0x3a, 0x57, 0x4b, 0xd7, 0xf9, 0xb3, 0x74, 0x9d, 0x27, 0xf5, 0xb3, 0x61, 0xfe,
	0xf4, 0xf9, 0x69, 0xbd, 0x9f, 0x8e, 0x1b, 0xa6, 0x90, 0xf7, 0x07, 0xe9, 0xb8, 0x37, 0x7c, 0xb6,
	0xfe, 0x68, 0xbc, 0x2c, 0x3b, 0x9c, 0xcf, 0x26, 0xc9, 0xf4, 0x74, 0xd7, 0x94, 0xf7, 0xc1, 0xbf,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x9e, 0x72, 0x2d, 0x15, 0xe2, 0x02, 0x00, 0x00,
}
