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
	// Zeta
	// ZETA_MAINNET = 10018;
	WalletRecordType_ZETA_TESTNET WalletRecordType = 10019
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
	10019: "ZETA_TESTNET",
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
	"ZETA_TESTNET":     10019,
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
	// 469 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x92, 0xcf, 0x6e, 0xd3, 0x30,
	0x1c, 0xc7, 0x63, 0x98, 0xc6, 0xf4, 0xeb, 0xd6, 0x79, 0xee, 0xb8, 0xe6, 0xc0, 0x71, 0x12, 0xed,
	0x81, 0x27, 0x70, 0x32, 0xb7, 0x31, 0x4a, 0xec, 0x2a, 0x71, 0x87, 0xd8, 0xa5, 0xea, 0xda, 0x68,
	0x54, 0x6a, 0x49, 0x95, 0x06, 0x41, 0xdf, 0x62, 0xfc, 0xff, 0x57, 0x10, 0x03, 0x1e, 0x00, 0xde,
	0x82, 0xe3, 0x8e, 0x1c, 0x51, 0x7b, 0xdb, 0x53, 0xa0, 0xba, 0x89, 0xd3, 0x1d, 0xf3, 0xd1, 0x47,
	0x5f, 0x7f, 0xed, 0x7c, 0xe1, 0xde, 0x78, 0xd6, 0x8f, 0x47, 0x8d, 0x34, 0x3e, 0x1f, 0x4e, 0xb3,
	0x74, 0xd6, 0x78, 0xde, 0x1b, 0x8d, 0xe2, 0xac, 0x9b, 0xc6, 0xfd, 0x24, 0x1d, 0xd4, 0x27, 0x69,
	0x92, 0x25, 0xa4, 0xaa, 0x9d, 0x7a, 0xe1, 0x1c, 0xfd, 0xdc, 0x02, 0xfc, 0x48, 0x7b, 0xa1, 0xd6,
	0xd4, 0x6c, 0x12, 0x93, 0x1a, 0xec, 0x3b, 0x5c, 0xb9, 0x92, 0x8b, 0x6e, 0x40, 0xb9, 0x10, 0x4c,
	0x61, 0x6b, 0x13, 0x2a, 0x16, 0xa9, 0x15, 0x44, 0xe4, 0x2e, 0x60, 0xa6, 0x3c, 0x16, 0xb2, 0x4e,
	0x60, 0xd4, 0x0b, 0x41, 0x0e, 0x61, 0xdf, 0xe0, 0x96, 0x64, 0xa1, 0xcf, 0xf1, 0x4b, 0x71, 0x43,
	0x8e, 0x58, 0x5b, 0xfa, 0x9c, 0xe2, 0x57, 0x5a, 0x6e, 0x4b, 0xff, 0x71, 0x4b, 0x96, 0xa7, 0xbd,
	0x16, 0xa4, 0x06, 0x55, 0x43, 0x3b, 0x81, 0x43, 0x39, 0x7e, 0x23, 0x08, 0x86, 0x8a, 0x23, 0x1c,
	0xa3, 0xbd, 0x35, 0xa4, 0x68, 0xf4, 0x4e, 0x9f, 0x42, 0x4f, 0xa8, 0x4f, 0x85, 0xeb, 0xb1, 0xae,
	0xeb, 0x7a, 0x94, 0x0b, 0xfc, 0x5e, 0xe7, 0x95, 0xb8, 0xd9, 0x79, 0xc8, 0xf1, 0x07, 0x0d, 0x5b,
	0x42, 0x46, 0x3c, 0x32, 0x91, 0x1f, 0x05, 0x21, 0xb0, 0x97, 0x43, 0xd7, 0xe3, 0xf4, 0x58, 0xe2,
	0x4f, 0x3a, 0x54, 0xb6, 0x15, 0x0f, 0x78, 0x54, 0xde, 0xf3, 0xb3, 0xae, 0x6e, 0x70, 0x7e, 0xcf,
	0xf9, 0xba, 0x41, 0xe8, 0x70, 0x15, 0x6e, 0x3c, 0xca, 0x17, 0x2d, 0x1b, 0x9c, 0xcb, 0x5f, 0xb5,
	0x1c, 0x79, 0x34, 0x3c, 0x5e, 0x3d, 0x8a, 0xc3, 0x14, 0x5d, 0xc9, 0x97, 0x82, 0x1c, 0xc0, 0xee,
	0x29, 0x53, 0xd4, 0x5c, 0xec, 0xc7, 0xaa, 0xec, 0x1e, 0x6d, 0x2b, 0x59, 0x76, 0xfd, 0x36, 0x47,
	0x25, 0x2c, 0xc4, 0xcb, 0x39, 0x22, 0x07, 0x50, 0x89, 0x3a, 0xdc, 0x78, 0xdf, 0x4b, 0x64, 0xe2,
	0xe6, 0x88, 0x1c, 0x42, 0x35, 0x92, 0x3e, 0x15, 0xd4, 0x88, 0xbf, 0xae, 0x37, 0x69, 0xe1, 0xfe,
	0xbe, 0x46, 0x47, 0x4d, 0xa8, 0xad, 0x57, 0x42, 0x07, 0x83, 0x34, 0x9e, 0x4e, 0x9b, 0x49, 0x3a,
	0xee, 0x65, 0xa4, 0x02, 0x77, 0xf2, 0x4d, 0x60, 0x8b, 0xec, 0xc2, 0x4e, 0xf1, 0x7b, 0x31, 0x22,
	0x3b, 0xb0, 0x15, 0xc8, 0x13, 0x86, 0x6f, 0x11, 0x80, 0xed, 0x75, 0x22, 0xbe, 0xed, 0x78, 0x7f,
	0x16, 0x36, 0xba, 0x5a, 0xd8, 0xe8, 0xdf, 0xc2, 0x46, 0x17, 0x4b, 0xdb, 0xba, 0x5a, 0xda, 0xd6,
	0xdf, 0xa5, 0x6d, 0x9d, 0xd6, 0xcf, 0x87, 0xd9, 0x93, 0x67, 0x67, 0xf5, 0x7e, 0x32, 0x6e, 0xe8,
	0x8d, 0xde, 0x1f, 0x24, 0xe3, 0xde, 0xf0, 0xe9, 0xfa, 0xa3, 0xf1, 0xa2, 0x9c, 0x75, 0x36, 0x9b,
	0xc4, 0xd3, 0xb3, 0x6d, 0xbd, 0xe7, 0x07, 0xff, 0x03, 0x00, 0x00, 0xff, 0xff, 0x38, 0x89, 0x35,
	0x14, 0xf5, 0x02, 0x00, 0x00,
}