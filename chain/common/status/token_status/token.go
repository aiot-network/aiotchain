package token_status

import (
	"errors"
	"fmt"
	"github.com/aiot-network/aiotchain/chain/db/status/token_db"
	chaintypes "github.com/aiot-network/aiotchain/chain/types"
	"github.com/aiot-network/aiotchain/common/config"
	"github.com/aiot-network/aiotchain/tools/arry"
	"github.com/aiot-network/aiotchain/types"
	"sync"
)

const tokenDB = "token_db"

type TokenStatus struct {
	db    ITokenDB
	mutex sync.RWMutex
}

func NewTokenStatus() (*TokenStatus, error) {
	db, err := token_db.Open(config.Param.Data + "/" + tokenDB)
	if err != nil {
		return nil, err
	}
	return &TokenStatus{db: db}, nil
}

func (t *TokenStatus) SetTrieRoot(hash arry.Hash) error {
	return t.db.SetRoot(hash)
}

func (t *TokenStatus) TrieRoot() arry.Hash {
	return t.db.Root()
}

func (t *TokenStatus) Commit() (arry.Hash, error) {
	return t.db.Commit()
}

func (t *TokenStatus) CheckMessage(msg types.IMessage) error {
	t.mutex.RLock()
	defer t.mutex.RUnlock()

	switch chaintypes.MessageType(msg.Type()) {
	case chaintypes.Token:
		body, ok := msg.MsgBody().(*chaintypes.TokenBody)
		if !ok {
			return errors.New("incorrect message type and message body")
		}
		token := t.db.Token(body.TokenAddress)
		if token != nil {
			return token.CheckToken(msg)
		}
	case chaintypes.TokenV2:
		body, ok := msg.MsgBody().(*chaintypes.TokenV2Body)
		if !ok {
			return errors.New("incorrect message type and message body")
		}
		token := t.db.Token(body.TokenAddress)
		if token != nil {
			return token.CheckTokenV2(msg)
		}
	case chaintypes.Redemption:
		body, ok := msg.MsgBody().(*chaintypes.RedemptionBody)
		if !ok {
			return errors.New("incorrect message type and message body")
		}
		token := t.db.Token(body.TokenAddress)
		if token != nil {
			return token.CheckRedemption(msg)
		}
	}
	return nil
}

// Update contract status
func (t *TokenStatus) UpdateToken(msg types.IMessage, height uint64) error {
	t.mutex.Lock()
	defer t.mutex.Unlock()

	switch chaintypes.MessageType(msg.Type()) {
	case chaintypes.Token:
		msgBody, ok := msg.MsgBody().(*chaintypes.TokenBody)
		if !ok {
			return errors.New("wrong message type")
		}
		record := &chaintypes.Record{
			Height:   height,
			Type:     "Token",
			MsgHash:  msg.Hash(),
			Receiver: msgBody.Receiver,
			Time:     msg.Time(),
			Amount:   msgBody.MsgAmount(),
		}
		tokenAddr := msgBody.TokenAddress
		token := t.db.Token(tokenAddr)
		if token != nil && token.IncreaseIssues {
			token.IncreaseRecord(record)
			token.Name = msgBody.Name
		} else {
			token = &chaintypes.TokenRecord{
				Address:        tokenAddr,
				Sender:         msg.From(),
				Name:           msgBody.Name,
				Shorthand:      msgBody.Shorthand,
				IncreaseIssues: msgBody.IncreaseIssues,
				Records: &chaintypes.RecordList{
					record,
				},
			}
		}
		t.db.SetToken(token)
	case chaintypes.TokenV2:
		msgBody, ok := msg.MsgBody().(*chaintypes.TokenV2Body)
		if !ok {
			return errors.New("wrong message type")
		}
		record := &chaintypes.Record{
			Height:   height,
			Type:     "Token",
			MsgHash:  msg.Hash(),
			Receiver: msgBody.Receiver,
			Time:     msg.Time(),
			Amount:   msgBody.MsgAmount(),
		}
		tokenAddr := msgBody.TokenAddress
		token := t.db.Token(tokenAddr)
		if token != nil && token.IncreaseIssues {
			token.IncreaseRecord(record)
			token.Name = msgBody.Name
		} else {
			token = &chaintypes.TokenRecord{
				Address:        tokenAddr,
				Sender:         msg.From(),
				Name:           msgBody.Name,
				Shorthand:      msgBody.Shorthand,
				IncreaseIssues: false,
				PledgeRate:     msgBody.PledgeRate,
				PledgeAmount:   msgBody.PledgeAmount(),
				Records: &chaintypes.RecordList{
					record,
				},
			}
		}
		t.db.SetToken(token)
	case chaintypes.Redemption:
		msgBody, ok := msg.MsgBody().(*chaintypes.RedemptionBody)
		if !ok {
			return errors.New("wrong message type")
		}
		record := &chaintypes.Record{
			Height:   height,
			Type:     "Redemption",
			MsgHash:  msg.Hash(),
			Receiver: msg.From(),
			Time:     msg.Time(),
			Amount:   msgBody.MsgAmount(),
		}
		tokenAddr := msgBody.TokenAddress
		token := t.db.Token(tokenAddr)
		if token == nil {
			return fmt.Errorf("token %s is not exist", msgBody.MsgToken().String())
		}
		reAmount := msgBody.RedemptionAmount()
		if token.PledgeAmount < reAmount {
			return fmt.Errorf("insufficient to redeem")
		}
		token.PledgeAmount -= reAmount
		token.IncreaseRecord(record)
		t.db.SetToken(token)
	}

	return nil
}

func (t *TokenStatus) Token(address arry.Address) (types.IToken, error) {
	token := t.db.Token(address)
	if token == nil {
		return nil, errors.New("not found")
	}
	return token, nil
}
