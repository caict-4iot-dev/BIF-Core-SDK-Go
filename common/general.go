package common

import (
	"fmt"
)

func AccountGetInfoURL(url string, address string) string {
	return url + "/getAccountBase?address=" + address
}

func AccountGetMetadataURL(url string, address string, key string) string {
	if key == "" {
		return url + "/getAccount?address=" + address
	}
	return url + "/getAccount?address=" + address + "&key=" + key
}

func BlockGetValidatorsURL(url string, blockNumber int64) string {
	return fmt.Sprintf("%s/getLedger?seq=%d&with_validator=true", url, blockNumber)
}

func BlockGetTransactionsURL(url string, blockNumber int64) string {
	return fmt.Sprintf("%s/getTransactionHistory?ledger_seq=%d", url, blockNumber)
}

func TransactionEvaluationFee(url string) string {
	return url + "/testTransaction"
}

func PriTxSend(url string) string {
	return url + "/priTxSend"
}

func TransactionSubmitURL(url string) string {
	return url + "/submitTransaction"
}

func ContractCallURL(url string) string {
	return url + "/callContract"
}

func PrivatecontractCallURL(url string) string {
	return url + "/callPrivateContract"
}

func TransactionGetInfoURL(url string, hash string) string {
	return url + "/getTransactionHistory?hash=" + hash
}

func BlockGetNumberURL(url string) string {
	return url + "/getLedger"
}

func BlockGetInfoURL(url string, blockNumber int64) string {
	return fmt.Sprintf("%s/getLedger?seq=%d", url, blockNumber)
}

func BlockGetLatestInfoURL(url string) string {
	return url + "/getLedger"
}

func BlockGetLatestValidatorsURL(url string) string {
	return url + "/getLedger?with_validator=true"
}

func PriTxStoreRaw(url string) string {
	return url + "/priTxStoreRaw"
}

func PriTxReceiveRaw(url string, priTxHash string) string {
	return url + "/priTxReceiveRaw?pri_tx_hash=" + priTxHash
}

func PriTxReceive(url string, priTxHash string) string {
	return url + "/priTxReceive?pri_tx_hash=" + priTxHash
}
