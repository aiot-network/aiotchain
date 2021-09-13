package status

import (
	"github.com/aiot-network/aiotchain/tools/arry"
	"github.com/aiot-network/aiotchain/types"
)

type IStatus interface {
	InitRoots(actRoot, dPosRoot, tokenRoot arry.Hash) error
	Commit() (arry.Hash, arry.Hash, arry.Hash, error)
	SetConfirmed(confirmed uint64)
	CheckMsg(msg types.IMessage, strict bool, height uint64) error
	CheckBlockMsg(msg types.IMessage, strict bool) error
	Change(msgs []types.IMessage, block types.IBlock) error
	GenesisChange(msgs []types.IMessage, block types.IBlock) error
	Account(address arry.Address) types.IAccount
	Token(address arry.Address) (types.IToken, error)
	TokenList() []map[string]string
	Contract(address arry.Address) (types.IContract, error)
	ContractState(msgHash arry.Hash) types.IStatus
	Candidates() types.ICandidates
	CycleSupers(cycle uint64) types.ICandidates
	CycleReword(cycle uint64) []types.IReword
	CycleWork(cycle uint64, address arry.Address) (types.IWorks, error)
	SymbolContract(symbol string) (arry.Address, bool)
}
