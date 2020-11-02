package rpcserver

import (
	"errors"
	"fmt"

	"github.com/incognitochain/incognito-chain/blockchain"
	"github.com/incognitochain/incognito-chain/common"
	"github.com/incognitochain/incognito-chain/common/base58"
	"github.com/incognitochain/incognito-chain/incognitokey"
	"github.com/incognitochain/incognito-chain/metadata"
	"github.com/incognitochain/incognito-chain/rpcserver/bean"
	"github.com/incognitochain/incognito-chain/rpcserver/jsonresult"
	"github.com/incognitochain/incognito-chain/rpcserver/rpcservice"
	"github.com/incognitochain/incognito-chain/wallet"
)

// handleCreateUnstakeTransaction - RPC create and send unstake tx to network
func (httpServer *HttpServer) handleCreateUnstakeTransaction(params interface{}, closeChan <-chan struct{}) (interface{}, *rpcservice.RPCError) {
	var err error
	data, err := httpServer.handleCreateRawUnstakeTransaction(params)
	if err.(*rpcservice.RPCError) != nil {
		return nil, rpcservice.NewRPCError(rpcservice.CreateTxDataError, err)
	}
	tx := data.(jsonresult.CreateTransactionResult)
	base58CheckData := tx.Base58CheckData

	newParam := make([]interface{}, 0)
	newParam = append(newParam, base58CheckData)
	sendResult, err := httpServer.handleSendRawTransaction(newParam, closeChan)
	if err.(*rpcservice.RPCError) != nil {
		return nil, rpcservice.NewRPCError(rpcservice.SendTxDataError, err)
	}
	result := jsonresult.NewCreateTransactionResult(nil, sendResult.(jsonresult.CreateTransactionResult).TxID, nil, tx.ShardID)
	return result, nil
}

// handleUnstakeRawTx - handle unstake tx request
// and return raw data for making raw transaction
func (httpServer *HttpServer) handleCreateRawUnstakeTransaction(params interface{}) (interface{}, *rpcservice.RPCError) {

	paramsArray := common.InterfaceSlice(params)
	if paramsArray == nil || len(paramsArray) < 5 {
		return nil, rpcservice.NewRPCError(rpcservice.RPCInvalidParamsError, errors.New("param must be an array at least 5 element"))
	}

	createRawTxParam, errNewParam := bean.NewCreateRawTxParam(params)
	if errNewParam != nil {
		return nil, rpcservice.NewRPCError(rpcservice.RPCInvalidParamsError, errNewParam)
	}

	keyWallet := new(wallet.KeyWallet)
	keyWallet.KeySet = *createRawTxParam.SenderKeySet

	//Get data to create meta data
	data, ok := paramsArray[4].(map[string]interface{})
	if !ok {
		return nil, rpcservice.NewRPCError(rpcservice.RPCInvalidParamsError, fmt.Errorf("Invalid unstaking metadata %+v", paramsArray[4]))
	}

	//Get staking type
	unStakingType, ok := data["UnStakingType"].(float64)
	if !ok {
		return nil, rpcservice.NewRPCError(rpcservice.RPCInvalidParamsError, fmt.Errorf("Invalid UnStaking Type For Staking Transaction %+v", data["UnStakingType"]))
	}

	//Get Candidate Payment Address
	candidatePaymentAddress, ok := data["CandidatePaymentAddress"].(string)
	if !ok {
		return nil, rpcservice.NewRPCError(rpcservice.RPCInvalidParamsError, fmt.Errorf("Invalid Producer Payment Address for Staking Transaction %+v", data["CandidatePaymentAddress"]))
	}
	// Get private seed, a.k.a mining key
	privateSeed, ok := data["PrivateSeed"].(string)
	if !ok {
		return nil, rpcservice.NewRPCError(rpcservice.RPCInvalidParamsError, fmt.Errorf("Invalid Private Seed for Staking Transaction %+v", data["PrivateSeed"]))
	}
	privateSeedBytes, ver, err := base58.Base58Check{}.Decode(privateSeed)
	if (err != nil) || (ver != common.ZeroByte) {
		return nil, rpcservice.NewRPCError(rpcservice.UnexpectedError, errors.New("Decode privateseed failed"))
	}

	// Get candidate publickey
	candidateWallet, err := wallet.Base58CheckDeserialize(candidatePaymentAddress)
	if err != nil || candidateWallet == nil {
		return nil, rpcservice.NewRPCError(rpcservice.RPCInvalidParamsError, errors.New("Base58CheckDeserialize candidate Payment Address failed"))
	}
	pk := candidateWallet.KeySet.PaymentAddress.Pk

	committeePK, err := incognitokey.NewCommitteeKeyFromSeed(privateSeedBytes, pk)
	if err != nil {
		return nil, rpcservice.NewRPCError(rpcservice.RPCInvalidParamsError, err)
	}

	committeePKBytes, err := committeePK.Bytes()
	if err != nil {
		return nil, rpcservice.NewRPCError(rpcservice.RPCInvalidParamsError, err)
	}

	unStakingMetadata, err := metadata.NewUnStakingMetadata(int(unStakingType), base58.Base58Check{}.Encode(committeePKBytes, common.ZeroByte))
	if err != nil {
		return nil, rpcservice.NewRPCError(rpcservice.RPCInvalidParamsError, err)
	}

	beaconview := httpServer.blockService.BlockChain.BeaconChain.GetFinalView()
	beaconFinalView := beaconview.(*blockchain.BeaconBestState)
	check, ok := beaconFinalView.GetAutoStaking()[unStakingMetadata.CommitteePublicKey]
	if !ok {
		return nil, rpcservice.NewRPCError(rpcservice.RPCInvalidParamsError, errors.New("Committee Public Key has not staked yet"))
	}
	if !check {
		return nil, rpcservice.NewRPCError(rpcservice.RPCInvalidParamsError, errors.New("Committee Public Key AutoStaking has been already false"))
	}

	txID, txBytes, txShardID, err := httpServer.txService.CreateRawTransaction(createRawTxParam, unStakingMetadata)
	if err.(*rpcservice.RPCError) != nil {
		return nil, rpcservice.NewRPCError(rpcservice.CreateTxDataError, err)
	}

	result := jsonresult.CreateTransactionResult{
		TxID:            txID.String(),
		Base58CheckData: base58.Base58Check{}.Encode(txBytes, common.ZeroByte),
		ShardID:         txShardID,
	}

	return result, nil
}