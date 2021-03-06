package param

import (
	"github.com/aiot-network/aiotchain/common/private"
	"github.com/aiot-network/aiotchain/tools/arry"
	"time"
)

const (
	// Block interval period
	BlockInterval = uint64(15)
	// Re-election interval
	CycleInterval = 60 * 60 * 24
	//CycleInterval = 60
	// Maximum number of super nodes
	SuperSize = 9

	DPosSize = SuperSize*2/3 + 1
)

const (
	// Mainnet logo
	MainNet = "mainnet"
	// Testnet logo
	TestNet = "testnet"

	APPName = "AIOT_NETWORK"
)
const (
	MaxReadBytes = 1024 * 10
	MaxReqBytes  = MaxReadBytes * 1000
)

// AtomsPerCoin is the number of atomic units in one coin.
const AtomsPerCoin = 1e8

type Param struct {
	Name              string
	Data              string
	App               string
	RollBack          uint64
	PubKeyHashAddrID  [2]byte
	PubKeyHashTokenID [2]byte
	HDPrivateKeyID    [4]byte
	HDPublicKeyID     [4]byte
	Logging           bool
	PeerRequestChan   uint32
	*PrivateParam
	*TokenParam
	*P2pParam
	*RpcParam
	*DPosParam
	*PoolParam
	private.IPrivate
}

type TokenParam struct {
	PreCirculation        uint64
	Circulation           uint64
	CoinBaseOneDay        uint64
	EveryChangeCoinHeight uint64
	Proportion            uint64
	MinCoinCount          float64
	MaxCoinCount          float64
	MinimumTransfer       uint64
	MaximumTransfer       uint64
	RedemptionRate        uint64
	Consume               uint64
	MaximumReceiver       int
	MainToken             arry.Address
	EaterAddress          arry.Address
	PreCirculations       []PreCirculation
}

type PrivateParam struct {
	PrivateFile string
	PrivatePass string
}

type P2pParam struct {
	P2pPort    string
	ExternalIp string
	NetWork    string
	CustomBoot string
}

type RpcParam struct {
	RpcIp      string
	RpcPort    string
	RpcTLS     bool
	RpcCert    string
	RpcCertKey string
	RpcUser    string
	RpcPass    string
	HttpPort   string
}

type AddressInfo struct {
	Address string
	P2PId   string
}

type DPosParam struct {
	BlockInterval       uint64
	CycleInterval       uint64
	SuperSize           int
	DPosSize            int
	GenesisTime         uint64
	GenesisCycle        uint64
	WorkProofAddress    string
	GenesisSuperList    []AddressInfo
	CoinBaseAddressList *CoinBaseAddress
}

type PoolParam struct {
	MsgExpiredTime     int64
	MonitorMsgInterval time.Duration
	MaxPoolMsg         int
	MaxAddressMsg      uint64
}

var TestNetParam = &Param{
	Name:              TestNet,
	Data:              "data",
	App:               "AIOT_NETWORK",
	RollBack:          0,
	PubKeyHashAddrID:  [2]byte{0x12, 0xfb},
	PubKeyHashTokenID: [2]byte{0x13, 0x14},
	HDPrivateKeyID:    [4]byte{0x02, 0xb7, 0xc3, 0x21},
	HDPublicKeyID:     [4]byte{0x02, 0xb7, 0xc3, 0x20},
	Logging:           true,
	PeerRequestChan:   1000,
	PrivateParam: &PrivateParam{
		PrivateFile: "key.json",
		PrivatePass: APPName,
	},
	TokenParam: &TokenParam{
		PreCirculation:  160000 * AtomsPerCoin,
		Circulation:     1 * 1e8 * AtomsPerCoin,
		CoinBaseOneDay:  27390 * AtomsPerCoin,
		Consume:         1e4 * AtomsPerCoin,
		MinCoinCount:    1 * 1e4,
		MaxCoinCount:    9 * 1e10,
		MinimumTransfer: 0.0001 * AtomsPerCoin,
		MaximumTransfer: 1 * 1e7 * AtomsPerCoin,
		MaximumReceiver: 1 * 1e4,
		RedemptionRate:  80,
		MainToken:       arry.StringToAddress("AIOT"),
		EaterAddress:    arry.StringToAddress("aiCoinEaterAddressDontSend000000000"),
		PreCirculations: []PreCirculation{{
			Address: "aiCSxRKuF8dYALbZ2av8gqcoVR34R4aecYX",
			Amount:  160000 * AtomsPerCoin,
		}},
	},
	P2pParam: &P2pParam{
		NetWork:    TestNet + "AIOT_NETWORK",
		P2pPort:    "13561",
		ExternalIp: "0.0.0.0",
		//aiTewnyK73P3chNgp7LgC8FoCHUhTe9cijZ
		CustomBoot: "/ip4/103.68.63.163/tcp/6008/ipfs/16Uiu2HAmJRKJkBvTxFEoSpQmvPaHZuVRrHBYVVKKetCMMBZ938ty",
	},
	RpcParam: &RpcParam{
		RpcIp:      "127.0.0.1",
		RpcPort:    "13562",
		HttpPort:   "13563",
		RpcTLS:     false,
		RpcCert:    "",
		RpcCertKey: "",
		RpcUser:    "",
		RpcPass:    "",
	},
	DPosParam: &DPosParam{
		BlockInterval:    BlockInterval,
		CycleInterval:    CycleInterval,
		SuperSize:        SuperSize,
		DPosSize:         DPosSize,
		GenesisTime:      1592268410,
		GenesisCycle:     1592268410 / CycleInterval,
		WorkProofAddress: "aiCSxRKuF8dYALbZ2av8gqcoVR34R4aecYX",
		GenesisSuperList: []AddressInfo{
			{
				Address: "aiMKrGcEGPFyRSW4WdM2ARY7kpc38EYpygy",
				P2PId:   "16Uiu2HAmMH8yCqrRvzyjEdcJ817pNQpcgrgZn95d5kn4LF2xm66L",
			},
			{
				Address: "aiAQ56a2nTmAxJ2Tycm8a3X8zYqs2NrDdUN",
				P2PId:   "16Uiu2HAmF6rDNZBympeEsDfoVTPPwtjRiizxDNmLxt5B7JAZTwdg",
			},
			{
				Address: "aiDvFTATGmmdcyD2trkrVcrf52QB2SXSEQf",
				P2PId:   "16Uiu2HAm851T4cCfRM7BvjgRdeirn4zeim8QUNAkmNW9A4sjWr9o",
			},
			{
				Address: "aiBntJ9itzbzT9R9Kpo2VJFwox6FfQV6UyM",
				P2PId:   "16Uiu2HAmAPg4wy9xHnrVedzgG6pkEhhvKRBw5ouhUPLYLi1GteKs",
			},
			{
				Address: "aiQ6KGjuZoJqabZqeVcv9iRLRL5D6iRmu5R",
				P2PId:   "16Uiu2HAm4po19Qm6gxVAQrP3ctmNgZdkgUYiLUtGXkKjAu1hzV9U",
			},
			{
				Address: "aiDU4CN68G3iMoR3iZL3oZFw1fdxE6sZYEg",
				P2PId:   "16Uiu2HAmLUKpbkYqWpFNDyxXwfaPg9BJ3AWP3gyidxhX6iGPw2dY",
			},
			{
				Address: "aiR264eSVnqwSPbiaRFp3oD66zi5Xu8MjzF",
				P2PId:   "16Uiu2HAmK5aMz7ah2edHnKHCnNFmkxPgJMgnrNBfgWgTLDtEtDbL",
			},
			{
				Address: "aiEiRsScq4rdTiBHc8XSTVmBzwgXYcFTZyS",
				P2PId:   "16Uiu2HAky7uuaiGx9cV3CD9WWUbobqHNsD9F4viQFzaYb3oB3fxj",
			},
			{
				Address: "aiChEPSNLznNv7hn3R2hEc98KbguLzoQQW1",
				P2PId:   "16Uiu2HAmUEgz6fTzy7KNaecXq4bfvaLymHwoMJoJ47KxANvp4YMe",
			},
		},
		CoinBaseAddressList: &CoinBaseAddress{
			{
				Address: "aiMKrGcEGPFyRSW4WdM2ARY7kpc38EYpygy",
				P2PId:   "16Uiu2HAmMH8yCqrRvzyjEdcJ817pNQpcgrgZn95d5kn4LF2xm66L",
			},
			{
				Address: "aiAQ56a2nTmAxJ2Tycm8a3X8zYqs2NrDdUN",
				P2PId:   "16Uiu2HAmF6rDNZBympeEsDfoVTPPwtjRiizxDNmLxt5B7JAZTwdg",
			},
			{
				Address: "aiDvFTATGmmdcyD2trkrVcrf52QB2SXSEQf",
				P2PId:   "16Uiu2HAm851T4cCfRM7BvjgRdeirn4zeim8QUNAkmNW9A4sjWr9o",
			},
			{
				Address: "aiBntJ9itzbzT9R9Kpo2VJFwox6FfQV6UyM",
				P2PId:   "16Uiu2HAmAPg4wy9xHnrVedzgG6pkEhhvKRBw5ouhUPLYLi1GteKs",
			},
			{
				Address: "aiQ6KGjuZoJqabZqeVcv9iRLRL5D6iRmu5R",
				P2PId:   "16Uiu2HAm4po19Qm6gxVAQrP3ctmNgZdkgUYiLUtGXkKjAu1hzV9U",
			},
			{
				Address: "aiDU4CN68G3iMoR3iZL3oZFw1fdxE6sZYEg",
				P2PId:   "16Uiu2HAmLUKpbkYqWpFNDyxXwfaPg9BJ3AWP3gyidxhX6iGPw2dY",
			},
			{
				Address: "aiR264eSVnqwSPbiaRFp3oD66zi5Xu8MjzF",
				P2PId:   "16Uiu2HAmK5aMz7ah2edHnKHCnNFmkxPgJMgnrNBfgWgTLDtEtDbL",
			},
			{
				Address: "aiEiRsScq4rdTiBHc8XSTVmBzwgXYcFTZyS",
				P2PId:   "16Uiu2HAky7uuaiGx9cV3CD9WWUbobqHNsD9F4viQFzaYb3oB3fxj",
			},
			{
				Address: "aiChEPSNLznNv7hn3R2hEc98KbguLzoQQW1",
				P2PId:   "16Uiu2HAmUEgz6fTzy7KNaecXq4bfvaLymHwoMJoJ47KxANvp4YMe",
			},
		},
	},
	PoolParam: &PoolParam{
		MaxPoolMsg:         100000,
		MsgExpiredTime:     60 * 60 * 3,
		MonitorMsgInterval: 10,
		MaxAddressMsg:      1000,
	},
}

var MainNetParam = &Param{
	Name:              MainNet,
	Data:              "data",
	App:               APPName,
	RollBack:          0,
	PubKeyHashAddrID:  [2]byte{0x5, 0x78},
	PubKeyHashTokenID: [2]byte{0x5, 0x91},
	HDPrivateKeyID:    [4]byte{0x01, 0xb7, 0xc3, 0x21},
	HDPublicKeyID:     [4]byte{0x01, 0xb7, 0xc3, 0x20},
	Logging:           true,
	PeerRequestChan:   1000,
	PrivateParam: &PrivateParam{
		PrivateFile: "key.json",
		PrivatePass: "AIOT_NETWORK",
	},
	TokenParam: &TokenParam{
		PreCirculation:  160000 * AtomsPerCoin,
		Circulation:     1 * 1e8 * AtomsPerCoin,
		CoinBaseOneDay:  27390 * AtomsPerCoin,
		Consume:         1e4 * AtomsPerCoin,
		MinCoinCount:    1 * 1e4,
		MaxCoinCount:    9 * 1e10,
		MinimumTransfer: 0.0001 * AtomsPerCoin,
		MaximumTransfer: 1 * 1e7 * AtomsPerCoin,
		MaximumReceiver: 1 * 1e4,
		RedemptionRate:  80,
		MainToken:       arry.StringToAddress("AIOT"),
		EaterAddress:    arry.StringToAddress("AiCoinEaterAddressDontSend000000000"),
		PreCirculations: []PreCirculation{{
			Address: "AijJxc2QqB32QmF83WZrLT5TFz61web6bRw",
			Amount:  160000 * AtomsPerCoin,
		}},
	},
	P2pParam: &P2pParam{
		NetWork:    MainNet + "AIOT_NETWORK",
		P2pPort:    "23561",
		ExternalIp: "0.0.0.0",
		//Aib7bDUym4KWu9painU1NHYi9w2dVAfe3ok
		CustomBoot: "/ip4/180.188.198.241/tcp/29564/ipfs/16Uiu2HAmBA6XR8USSaN4SWPwaha5gYcADSBu2hfA7rWtVE9NN44m",
	},
	RpcParam: &RpcParam{
		RpcIp:      "127.0.0.1",
		RpcPort:    "23562",
		HttpPort:   "23563",
		RpcTLS:     false,
		RpcCert:    "",
		RpcCertKey: "",
		RpcUser:    "",
		RpcPass:    "",
	},
	DPosParam: &DPosParam{
		BlockInterval:    BlockInterval,
		CycleInterval:    CycleInterval,
		SuperSize:        SuperSize,
		DPosSize:         DPosSize,
		GenesisTime:      1624083180,
		GenesisCycle:     1624083180 / CycleInterval,
		WorkProofAddress: "AiZ3V77E7S5jA8afLVrS3eDYoNFKXQYSRo1",
		GenesisSuperList: []AddressInfo{
			{
				Address: "AiXNaeAq1uDTEeCzFM7E3SUyt9QAgWfFY2W",
				P2PId:   "16Uiu2HAm5wMadephoZ6rB7kTGQihPxJiMp6fbMwSgyAN11TEiMDS",
			},
			{
				Address: "AigRAj8FbhmpYvUrLHGcXHJpc2FSfaou93s",
				P2PId:   "16Uiu2HAm7WtyPvenAqgbmrMnyCTnAeSSvWEmB2BrYSBuJheUUPes",
			},
			{
				Address: "AiLDAse8pXRbuvCfQccs58KKW9QXTyW1MjU",
				P2PId:   "16Uiu2HAm6RebrSYtH5XnuHe7U6ba5q766bfAVeu5k3gd2gUKePsJ",
			},
			{
				Address: "Aie5jBLzfCGytjPww74gUVmRUZ4DAxo8PSj",
				P2PId:   "16Uiu2HAm9SqiY4vdFgv245ksscq9pXpZBavWMR4GDZdaZBu3W6Fk",
			},
			{
				Address: "AiYN8StRg2T9mW8wqyiY6nGhCHqGB1sBziQ",
				P2PId:   "16Uiu2HAm1fY5f7HTq2Y2PH5fbRHNzKRzYxLMvaDH3VQ2LJiUdAEB",
			},
			{
				Address: "Aibj78ho6oLWXxqWB52GiWQT7PZfjhcmvZ1",
				P2PId:   "16Uiu2HAm9d4PQMLK3XLyZgKmRKPgFGg9LmRudFwKqXovhcA1u5iP",
			},
			{
				Address: "AiaCTpxmHbMGWrG73J4Qo8dVmd36Kkoewai",
				P2PId:   "16Uiu2HAmU5EzM2dDKeWFdYSZETMixru5hxcbE1tT8yx8yzFvubPK",
			},
			{
				Address: "AihNXLpZxsAsVMhYsNx3Dbdr11LeZYUjxzk",
				P2PId:   "16Uiu2HAm3fu1EWk5VQGg3DT4ih38BodMaHxhC3CGHdYM9eAUkNM7",
			},
			{
				Address: "AiYyPZGDwictAPoddqBDzV9Drc3LGc3xKrH",
				P2PId:   "16Uiu2HAmFdmQdg4TjUNLfYbavVENU74WCCrUnNqfXCP7n1M3DjRy",
			},
		},
		CoinBaseAddressList: &CoinBaseAddress{
			{
				Address: "AiXNaeAq1uDTEeCzFM7E3SUyt9QAgWfFY2W",
				P2PId:   "16Uiu2HAm5wMadephoZ6rB7kTGQihPxJiMp6fbMwSgyAN11TEiMDS",
			},
			{
				Address: "AigRAj8FbhmpYvUrLHGcXHJpc2FSfaou93s",
				P2PId:   "16Uiu2HAm7WtyPvenAqgbmrMnyCTnAeSSvWEmB2BrYSBuJheUUPes",
			},
			{
				Address: "AiLDAse8pXRbuvCfQccs58KKW9QXTyW1MjU",
				P2PId:   "16Uiu2HAm6RebrSYtH5XnuHe7U6ba5q766bfAVeu5k3gd2gUKePsJ",
			},
			{
				Address: "Aie5jBLzfCGytjPww74gUVmRUZ4DAxo8PSj",
				P2PId:   "16Uiu2HAm9SqiY4vdFgv245ksscq9pXpZBavWMR4GDZdaZBu3W6Fk",
			},
			{
				Address: "AiYN8StRg2T9mW8wqyiY6nGhCHqGB1sBziQ",
				P2PId:   "16Uiu2HAm1fY5f7HTq2Y2PH5fbRHNzKRzYxLMvaDH3VQ2LJiUdAEB",
			},
			{
				Address: "Aibj78ho6oLWXxqWB52GiWQT7PZfjhcmvZ1",
				P2PId:   "16Uiu2HAm9d4PQMLK3XLyZgKmRKPgFGg9LmRudFwKqXovhcA1u5iP",
			},
			{
				Address: "AiaCTpxmHbMGWrG73J4Qo8dVmd36Kkoewai",
				P2PId:   "16Uiu2HAmU5EzM2dDKeWFdYSZETMixru5hxcbE1tT8yx8yzFvubPK",
			},
			{
				Address: "AihNXLpZxsAsVMhYsNx3Dbdr11LeZYUjxzk",
				P2PId:   "16Uiu2HAm3fu1EWk5VQGg3DT4ih38BodMaHxhC3CGHdYM9eAUkNM7",
			},
			{
				Address: "AiYyPZGDwictAPoddqBDzV9Drc3LGc3xKrH",
				P2PId:   "16Uiu2HAmFdmQdg4TjUNLfYbavVENU74WCCrUnNqfXCP7n1M3DjRy",
			},
		},
	},
	PoolParam: &PoolParam{
		MaxPoolMsg:         100000,
		MsgExpiredTime:     60 * 60 * 3,
		MonitorMsgInterval: 10,
		MaxAddressMsg:      1000,
	},
}

type PreCirculation struct {
	Address string
	Note    string
	Amount  uint64
}

type CoinBaseAddress []AddressInfo

func (s *CoinBaseAddress) CurrentAddress(height uint64) arry.Address {
	if len(*s) == 0 || height == 0 {
		return arry.Address{}
	}

	index := height % uint64(len(*s))
	if index == 0 {
		return arry.StringToAddress((*s)[len(*s)-1].Address)
	} else {
		return arry.StringToAddress((*s)[index-1].Address)
	}
}

func (s *CoinBaseAddress) Len() int {
	return len(*s)
}
