package EthereumAPI

import (
	"fmt"
)

//https://github.com/ethereum/wiki/wiki/JSON-RPC

func EthProtocolVersion() (string, error) {
	resp, err := Call("eth_protocolVersion", nil)
	if err != nil {
		return "", err
	}
	if resp.Error != nil {
		return "", fmt.Errorf(resp.Error.Message)
	}
	return resp.Result.(string), nil
}

func EthSyncing() (*EthSyncingResponse, error) {
	resp, err := Call("eth_syncing", nil)
	if err != nil {
		return nil, err
	}
	if resp.Error != nil {
		return nil, fmt.Errorf(resp.Error.Message)
	}
	answer := new(EthSyncingResponse)
	syncing, ok := resp.Result.(bool)
	if ok == true {
		answer.Syncing = syncing
		return answer, nil
	}
	err = MapToObject(resp.Result, answer)
	if err != nil {
		return nil, err
	}
	return answer, nil
}

func EthCoinbase() (string, error) {
	resp, err := Call("eth_coinbase", nil)
	if err != nil {
		return "", err
	}
	if resp.Error != nil {
		return "", fmt.Errorf(resp.Error.Message)
	}
	return resp.Result.(string), nil
}

func EthMining() (bool, error) {
	resp, err := Call("eth_mining", nil)
	if err != nil {
		return false, err
	}
	if resp.Error != nil {
		return false, fmt.Errorf(resp.Error.Message)
	}
	return resp.Result.(bool), nil
}

func EthHashrate() (int64, error) {
	resp, err := Call("eth_hashrate", nil)
	if err != nil {
		return 0, err
	}
	if resp.Error != nil {
		return 0, fmt.Errorf(resp.Error.Message)
	}
	return ParseQuantity(resp.Result.(string))
}

func EthGasPrice() (int64, error) {
	resp, err := Call("eth_gasPrice", nil)
	if err != nil {
		return 0, err
	}
	if resp.Error != nil {
		return 0, fmt.Errorf(resp.Error.Message)
	}
	return ParseQuantity(resp.Result.(string))
}

func EthAccounts() ([]string, error) {
	resp, err := Call("eth_accounts", nil)
	if err != nil {
		return nil, err
	}
	if resp.Error != nil {
		return nil, fmt.Errorf(resp.Error.Message)
	}
	answer := []string{}
	err = MapToObject(resp.Result, &answer)
	if err != nil {
		return nil, err
	}
	return answer, nil
}

func EthBlockNumber() (int64, error) {
	resp, err := Call("eth_blockNumber", nil)
	if err != nil {
		return 0, err
	}
	if resp.Error != nil {
		return 0, fmt.Errorf(resp.Error.Message)
	}
	return ParseQuantity(resp.Result.(string))
}

//TODO: test
func EthGetBalance(address string, blockNumberOrTag string) (int64, error) {
	resp, err := Call("eth_getBalance", []string{address, blockNumberOrTag})
	if err != nil {
		return 0, err
	}
	if resp.Error != nil {
		return 0, fmt.Errorf(resp.Error.Message)
	}
	return ParseQuantity(resp.Result.(string))
}

//TODO: test
func EthGetStorageAt(address, positionInTheStorage, blockNumberOrTag string) (string, error) {
	resp, err := Call("eth_getStorageAt", []string{address, positionInTheStorage, blockNumberOrTag})
	if err != nil {
		return "", err
	}
	if resp.Error != nil {
		return "", fmt.Errorf(resp.Error.Message)
	}
	return resp.Result.(string), nil
}

//TODO: test
func EthGetTransactionCount(address, blockNumberOrTag string) (int64, error) {
	resp, err := Call("eth_getTransactionCount", []string{address, blockNumberOrTag})
	if err != nil {
		return 0, err
	}
	if resp.Error != nil {
		return 0, fmt.Errorf(resp.Error.Message)
	}
	return ParseQuantity(resp.Result.(string))
}

//TODO: test
func EthGetBlockTransactionCountByHash(blockHash string) (int64, error) {
	resp, err := Call("eth_getBlockTransactionCountByHash", []string{blockHash})
	if err != nil {
		return 0, err
	}
	if resp.Error != nil {
		return 0, fmt.Errorf(resp.Error.Message)
	}
	return ParseQuantity(resp.Result.(string))
}

//TODO: test
func EthGetBlockTransactionCountByNumber(blockNumberOrTag string) (int64, error) {
	resp, err := Call("eth_getBlockTransactionCountByNumber", []string{blockNumberOrTag})
	if err != nil {
		return 0, err
	}
	if resp.Error != nil {
		return 0, fmt.Errorf(resp.Error.Message)
	}
	return ParseQuantity(resp.Result.(string))
}

//TODO: test
func EthGetUncleCountByBlockHash(blockHash string) (int64, error) {
	resp, err := Call("eth_getUncleCountByBlockHash", []string{blockHash})
	if err != nil {
		return 0, err
	}
	if resp.Error != nil {
		return 0, fmt.Errorf(resp.Error.Message)
	}
	return ParseQuantity(resp.Result.(string))
}

//TODO: test
func EthGetUncleCountByBlockNumber(blockNumberOrTag string) (interface{}, error) {
	resp, err := Call("eth_getUncleCountByBlockNumber", []string{blockNumberOrTag})
	if err != nil {
		return 0, err
	}
	if resp.Error != nil {
		return 0, fmt.Errorf(resp.Error.Message)
	}
	return ParseQuantity(resp.Result.(string))
}

//TODO: test
func EthGetCode(address, blockNumberOrTag string) (string, error) {
	resp, err := Call("eth_getCode", []string{address, blockNumberOrTag})
	if err != nil {
		return "", err
	}
	if resp.Error != nil {
		return "", fmt.Errorf(resp.Error.Message)
	}
	return resp.Result.(string), nil
}

//TODO: test
func EthSign(address, hashOfDataToSign string) (string, error) {
	resp, err := Call("eth_sign", []string{address, hashOfDataToSign})
	if err != nil {
		return "", err
	}
	if resp.Error != nil {
		return "", fmt.Errorf(resp.Error.Message)
	}
	return resp.Result.(string), nil
}

//TODO: test
func EthSendTransaction(tx *TransactionObject) (string, error) {
	resp, err := Call("eth_sendTransaction", []interface{}{tx})
	if err != nil {
		return "", err
	}
	if resp.Error != nil {
		return "", fmt.Errorf(resp.Error.Message)
	}
	return resp.Result.(string), nil
}

//TODO: test
func EthSendRawTransaction(txData string) (interface{}, error) {
	resp, err := Call("eth_sendRawTransaction", []string{txData})
	if err != nil {
		return "", err
	}
	if resp.Error != nil {
		return "", fmt.Errorf(resp.Error.Message)
	}
	return resp.Result.(string), nil
}

//TODO: test
func EthCall(tx *TransactionObject, blockNumberOrTag string) (interface{}, error) {
	resp, err := Call("eth_call", []interface{}{tx, blockNumberOrTag})
	if err != nil {
		return "", err
	}
	if resp.Error != nil {
		return "", fmt.Errorf(resp.Error.Message)
	}
	return resp.Result.(string), nil
}

//TODO: test
func EthEstimateGas(tx *TransactionObject, blockNumberOrTag string) (int64, error) {
	resp, err := Call("eth_estimateGas", []interface{}{tx, blockNumberOrTag})
	if err != nil {
		return 0, err
	}
	if resp.Error != nil {
		return 0, fmt.Errorf(resp.Error.Message)
	}
	return ParseQuantity(resp.Result.(string))
}

//TODO: test
func EthGetBlockByHash(blockHash string, fullTransaction bool) (*BlockObject, error) {
	resp, err := Call("eth_getBlockByHash", []interface{}{blockHash, fullTransaction})
	if err != nil {
		return nil, err
	}
	if resp.Error != nil {
		return nil, fmt.Errorf(resp.Error.Message)
	}
	answer := new(BlockObject)
	err = MapToObject(resp.Result, answer)
	if err != nil {
		return nil, err
	}
	return answer, nil
}

//TODO: test
func EthGetBlockByNumber(blockNumberOrTag string, fullTransaction bool) (*BlockObject, error) {
	resp, err := Call("eth_getBlockByNumber", []interface{}{blockNumberOrTag, fullTransaction})
	if err != nil {
		return nil, err
	}
	if resp.Error != nil {
		return nil, fmt.Errorf(resp.Error.Message)
	}
	answer := new(BlockObject)
	err = MapToObject(resp.Result, answer)
	if err != nil {
		return nil, err
	}
	return answer, nil
}

/*
//TODO: finish
func EthGetTransactionByHash(txHash string) (interface{}, error) {
	resp, err := Call("eth_getTransactionByHash", nil)
	if err != nil {
		return nil, err
	}
	if resp.Error != nil {
		return nil, fmt.Errorf(resp.Error.Message)
	}
	return resp.Result, nil
}

//TODO: finish
func EthGetTransactionByBlockHashAndIndex() (interface{}, error) {
	resp, err := Call("eth_getTransactionByBlockHashAndIndex", nil)
	if err != nil {
		return nil, err
	}
	if resp.Error != nil {
		return nil, fmt.Errorf(resp.Error.Message)
	}
	return resp.Result, nil
}

//TODO: finish
func EthGetTransactionByBlockNumberAndIndex() (interface{}, error) {
	resp, err := Call("eth_getTransactionByBlockNumberAndIndex", nil)
	if err != nil {
		return nil, err
	}
	if resp.Error != nil {
		return nil, fmt.Errorf(resp.Error.Message)
	}
	return resp.Result, nil
}

//TODO: finish
func EthGetTransactionReceipt() (interface{}, error) {
	resp, err := Call("eth_getTransactionReceipt", nil)
	if err != nil {
		return nil, err
	}
	if resp.Error != nil {
		return nil, fmt.Errorf(resp.Error.Message)
	}
	return resp.Result, nil
}

//TODO: finish
func EthGetUncleByBlockHashAndIndex() (interface{}, error) {
	resp, err := Call("eth_getUncleByBlockHashAndIndex", nil)
	if err != nil {
		return nil, err
	}
	if resp.Error != nil {
		return nil, fmt.Errorf(resp.Error.Message)
	}
	return resp.Result, nil
}

//TODO: finish
func EthGetUncleByBlockNumberAndIndex() (interface{}, error) {
	resp, err := Call("eth_getUncleByBlockNumberAndIndex", nil)
	if err != nil {
		return nil, err
	}
	if resp.Error != nil {
		return nil, fmt.Errorf(resp.Error.Message)
	}
	return resp.Result, nil
}
*/

func EthGetCompilers() ([]string, error) {
	resp, err := Call("eth_getCompilers", nil)
	if err != nil {
		return nil, err
	}
	if resp.Error != nil {
		return nil, fmt.Errorf(resp.Error.Message)
	}
	return resp.Result.([]string), nil
}

/*
//TODO: finish
func EthCompileLLL() (interface{}, error) {
	resp, err:=Call("eth_compileLLL", nil)
	if err!=nil {
		return nil, err
	}
	if resp.Error!=nil {
		return nil, fmt.Errorf(resp.Error.Message)
	}
	return resp.Result, nil
}



//TODO: finish
func EthCompileSolidity() (interface{}, error) {
	resp, err:=Call("eth_compileSolidity", nil)
	if err!=nil {
		return nil, err
	}
	if resp.Error!=nil {
		return nil, fmt.Errorf(resp.Error.Message)
	}
	return resp.Result, nil
}



//TODO: finish
func EthCompileSerpent() (interface{}, error) {
	resp, err:=Call("eth_compileSerpent", nil)
	if err!=nil {
		return nil, err
	}
	if resp.Error!=nil {
		return nil, fmt.Errorf(resp.Error.Message)
	}
	return resp.Result, nil
}



//TODO: finish
func EthNewFilter() (interface{}, error) {
	resp, err:=Call("eth_newFilter", nil)
	if err!=nil {
		return nil, err
	}
	if resp.Error!=nil {
		return nil, fmt.Errorf(resp.Error.Message)
	}
	return resp.Result, nil
}
*/

//TODO: finish
func EthNewBlockFilter() (interface{}, error) {
	resp, err := Call("eth_newBlockFilter", nil)
	if err != nil {
		return nil, err
	}
	if resp.Error != nil {
		return nil, fmt.Errorf(resp.Error.Message)
	}
	return resp.Result, nil
}

//TODO: finish
func EthNewPendingTransactionFilter() (interface{}, error) {
	resp, err := Call("eth_newPendingTransactionFilter", nil)
	if err != nil {
		return nil, err
	}
	if resp.Error != nil {
		return nil, fmt.Errorf(resp.Error.Message)
	}
	return resp.Result, nil
}

/*
//TODO: finish
func EthUninstallFilter() (interface{}, error) {
	resp, err:=Call("eth_uninstallFilter", nil)
	if err!=nil {
		return nil, err
	}
	if resp.Error!=nil {
		return nil, fmt.Errorf(resp.Error.Message)
	}
	return resp.Result, nil
}



//TODO: finish
func EthGetFilterChanges() (interface{}, error) {
	resp, err:=Call("eth_getFilterChanges", nil)
	if err!=nil {
		return nil, err
	}
	if resp.Error!=nil {
		return nil, fmt.Errorf(resp.Error.Message)
	}
	return resp.Result, nil
}



//TODO: finish
func EthGetFilterLogs() (interface{}, error) {
	resp, err:=Call("eth_getFilterLogs", nil)
	if err!=nil {
		return nil, err
	}
	if resp.Error!=nil {
		return nil, fmt.Errorf(resp.Error.Message)
	}
	return resp.Result, nil
}



//TODO: finish
func EthGetLogs() (interface{}, error) {
	resp, err:=Call("eth_getLogs", nil)
	if err!=nil {
		return nil, err
	}
	if resp.Error!=nil {
		return nil, fmt.Errorf(resp.Error.Message)
	}
	return resp.Result, nil
}
*/

//TODO: finish
func EthGetWork() (interface{}, error) {
	resp, err := Call("eth_getWork", nil)
	if err != nil {
		return nil, err
	}
	if resp.Error != nil {
		return nil, fmt.Errorf(resp.Error.Message)
	}
	return resp.Result, nil
}

/*
//TODO: finish
func EthSubmitWork() (interface{}, error) {
	resp, err:=Call("eth_submitWork", nil)
	if err!=nil {
		return nil, err
	}
	if resp.Error!=nil {
		return nil, fmt.Errorf(resp.Error.Message)
	}
	return resp.Result, nil
}



//TODO: finish
func EthSubmitHashrate() (interface{}, error) {
	resp, err:=Call("eth_submitHashrate", nil)
	if err!=nil {
		return nil, err
	}
	if resp.Error!=nil {
		return nil, fmt.Errorf(resp.Error.Message)
	}
	return resp.Result, nil
}

*/