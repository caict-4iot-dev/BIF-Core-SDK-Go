package exception

import "github.com/caict-4iot-dev/BIF-Core-SDK-Go/types/response"

var (
	// SUCCESS ...
	SUCCESS = response.BIFBaseResponse{0, "Success"}

	// ACCOUNT_CREATE_ERROR ...
	ACCOUNT_CREATE_ERROR = response.BIFBaseResponse{11001, "Failed to create the account"}

	// INVALID_AMOUNT_ERROR ...
	INVALID_AMOUNT_ERROR = response.BIFBaseResponse{11024, "Amount must be between 0 and Long.MAX_VALUE"}

	// INVALID_SOURCEADDRESS_ERROR ...
	INVALID_SOURCEADDRESS_ERROR = response.BIFBaseResponse{11002, "Invalid sourceAddress"}

	// INVALID_DESTADDRESS_ERROR ...
	INVALID_DESTADDRESS_ERROR = response.BIFBaseResponse{11003, "Invalid destAddress"}

	// INVALID_INITBALANCE_ERROR ...
	INVALID_INITBALANCE_ERROR = response.BIFBaseResponse{11004, "InitBalance must be between 1 and Long.MAX_VALUE"}

	// SOURCEADDRESS_EQUAL_DESTADDRESS_ERROR ...
	SOURCEADDRESS_EQUAL_DESTADDRESS_ERROR = response.BIFBaseResponse{11005, "SourceAddress cannot be equal to destAddress"}

	// INVALID_ADDRESS_ERROR ...
	INVALID_ADDRESS_ERROR = response.BIFBaseResponse{11006, "Invalid address"}

	// CONNECTNETWORK_ERROR ...
	CONNECTNETWORK_ERROR = response.BIFBaseResponse{11007, "Failed to connect to the network"}

	// NO_METADATAS_ERROR ...
	NO_METADATAS_ERROR = response.BIFBaseResponse{11010, "The account does not have this metadatas"}

	// INVALID_DATAKEY_ERROR ...
	INVALID_DATAKEY_ERROR = response.BIFBaseResponse{11011, "The length of key must be between 1 and 1024"}

	// INVALID_DATAVALUE_ERROR ...
	INVALID_DATAVALUE_ERROR = response.BIFBaseResponse{11012, "The length of value must be between 0 and 256000"}

	// INVALID_DATAVERSION_ERROR ...
	INVALID_DATAVERSION_ERROR = response.BIFBaseResponse{11013, "The version must be equal to or greater than 0"}

	// INVALID_MASTERWEIGHT_ERROR ...
	INVALID_MASTERWEIGHT_ERROR = response.BIFBaseResponse{11015, "MasterWeight must be between 0 and = response.BIFBaseResponse{Integer.MAX_VALUE // 2L + 1)"}

	// INVALID_SIGNER_ADDRESS_ERROR ...
	INVALID_SIGNER_ADDRESS_ERROR = response.BIFBaseResponse{11016, "Invalid signer address"}

	// INVALID_SIGNER_WEIGHT_ERROR ...
	INVALID_SIGNER_WEIGHT_ERROR = response.BIFBaseResponse{11017, "Signer weight must be between 0 and = response.BIFBaseResponse{Integer.MAX_VALUE // 2L + 1)"}

	// INVALID_TX_THRESHOLD_ERROR ...
	INVALID_TX_THRESHOLD_ERROR = response.BIFBaseResponse{11018, "TxThreshold must be between 0 and Long.MAX_VALUE"}

	// INVALID_TYPETHRESHOLD_TYPE_ERROR ...
	INVALID_TYPETHRESHOLD_TYPE_ERROR = response.BIFBaseResponse{11019, "Type of TypeThreshold is invalid"}

	// INVALID_TYPE_THRESHOLD_ERROR ...
	INVALID_TYPE_THRESHOLD_ERROR = response.BIFBaseResponse{11020, "TypeThreshold must be between 0 and Long.MAX_VALUE"}

	// INVALID_CONTRACT_HASH_ERROR ...
	INVALID_CONTRACT_HASH_ERROR = response.BIFBaseResponse{11025, "Invalid transaction hash to create contract"}

	// INVALID_GAS_AMOUNT_ERROR ...
	INVALID_GAS_AMOUNT_ERROR = response.BIFBaseResponse{11026, "bifAmount must be between 0 and Long.MAX_VALUE"}

	// INVALID_CONTRACTADDRESS_ERROR ...
	INVALID_CONTRACTADDRESS_ERROR = response.BIFBaseResponse{11037, "Invalid contract address"}

	// CONTRACTADDRESS_NOT_CONTRACTACCOUNT_ERROR ...
	CONTRACTADDRESS_NOT_CONTRACTACCOUNT_ERROR = response.BIFBaseResponse{11038, "contractAddress is not a contract account"}

	// SOURCEADDRESS_EQUAL_CONTRACTADDRESS_ERROR ...
	SOURCEADDRESS_EQUAL_CONTRACTADDRESS_ERROR = response.BIFBaseResponse{11040, "SourceAddress cannot be equal to contractAddress"}

	// INVALID_FROMADDRESS_ERROR ...
	INVALID_FROMADDRESS_ERROR = response.BIFBaseResponse{11041, "Invalid fromAddress"}

	// FROMADDRESS_EQUAL_DESTADDRESS_ERROR ...
	FROMADDRESS_EQUAL_DESTADDRESS_ERROR = response.BIFBaseResponse{11042, "FromAddress cannot be equal to destAddress"}

	// INVALID_SPENDER_ERROR ...
	INVALID_SPENDER_ERROR = response.BIFBaseResponse{11043, "Invalid spender"}

	// PAYLOAD_EMPTY_ERROR ...
	PAYLOAD_EMPTY_ERROR = response.BIFBaseResponse{11044, "Payload cannot be empty"}

	// INVALID_CONTRACT_TYPE_ERROR ...
	INVALID_CONTRACT_TYPE_ERROR = response.BIFBaseResponse{11047, "Invalid contract type"}

	// INVALID_NONCE_ERROR ...
	INVALID_NONCE_ERROR = response.BIFBaseResponse{11048, "Nonce must be between 1 and Long.MAX_VALUE"}

	// INVALID_GASPRICE_ERROR ...
	INVALID_GASPRICE_ERROR = response.BIFBaseResponse{11049, "GasPrice must be between 0 and Long.MAX_VALUE"}

	// INVALID_FEELIMIT_ERROR ...
	INVALID_FEELIMIT_ERROR = response.BIFBaseResponse{11050, "FeeLimit must be between 0 and Long.MAX_VALUE"}

	// OPERATIONS_EMPTY_ERROR ...
	OPERATIONS_EMPTY_ERROR = response.BIFBaseResponse{11051, "Operations cannot be empty"}

	// INVALID_CEILLEDGERSEQ_ERROR ...
	INVALID_CEILLEDGERSEQ_ERROR = response.BIFBaseResponse{11052, "CeilLedgerSeq must be equal to or greater than 0"}

	// OPERATIONS_ONE_ERROR ...
	OPERATIONS_ONE_ERROR = response.BIFBaseResponse{11053, "One of the operations cannot be resolved"}

	// INVALID_SIGNATURENUMBER_ERROR ...
	INVALID_SIGNATURENUMBER_ERROR = response.BIFBaseResponse{11054, "SignagureNumber must be between 1 and Integer.MAX_VALUE"}

	// INVALID_HASH_ERROR ...
	INVALID_HASH_ERROR = response.BIFBaseResponse{11055, "Invalid transaction hash"}

	// INVALID_SERIALIZATION_ERROR ...
	INVALID_SERIALIZATION_ERROR = response.BIFBaseResponse{11056, "Invalid serialization"}

	// PRIVATEKEY_NULL_ERROR ...
	PRIVATEKEY_NULL_ERROR = response.BIFBaseResponse{11057, "PrivateKeys cannot be empty"}

	// PRIVATEKEY_ONE_ERROR ...
	PRIVATEKEY_ONE_ERROR = response.BIFBaseResponse{11058, "One of privateKeys is invalid"}

	// SIGNDATA_NULL_ERROR ...
	SIGNDATA_NULL_ERROR = response.BIFBaseResponse{11059, "SignData cannot be empty"}

	// INVALID_BLOCKNUMBER_ERROR ...
	INVALID_BLOCKNUMBER_ERROR = response.BIFBaseResponse{11060, "BlockNumber must be bigger than 0"}

	// PUBLICKEY_NULL_ERROR ...
	PUBLICKEY_NULL_ERROR = response.BIFBaseResponse{11061, "PublicKey cannot be empty"}

	// URL_EMPTY_ERROR ...
	URL_EMPTY_ERROR = response.BIFBaseResponse{11062, "Url cannot be empty"}

	// CONTRACTADDRESS_CODE_BOTH_NULL_ERROR ...
	CONTRACTADDRESS_CODE_BOTH_NULL_ERROR = response.BIFBaseResponse{11063, "ContractAddress and code cannot be empty at the same time"}

	// INVALID_OPTTYPE_ERROR ...
	INVALID_OPTTYPE_ERROR = response.BIFBaseResponse{11064, "OptType must be between 0 and 2"}

	// GET_ALLOWANCE_ERROR ...
	GET_ALLOWANCE_ERROR = response.BIFBaseResponse{11065, "Failed to get allowance"}

	// SIGNATURE_EMPTY_ERROR ...
	SIGNATURE_EMPTY_ERROR = response.BIFBaseResponse{11067, "The signatures cannot be empty"}

	// OPERATIONS_INVALID_ERROR ...
	OPERATIONS_INVALID_ERROR = response.BIFBaseResponse{11068, "Operations length must be between 1 and 100"}

	// OPERATION_TYPE_ERROR 操作类型为空
	OPERATION_TYPE_ERROR = response.BIFBaseResponse{11077, "Operation type cannot be empty"}

	// CONNECTN_BLOCKCHAIN_ERROR ...
	CONNECTN_BLOCKCHAIN_ERROR = response.BIFBaseResponse{19999, "Failed to connect blockchain"}

	// SYSTEM_ERROR ...
	SYSTEM_ERROR = response.BIFBaseResponse{20000, "System error"}

	// REQUEST_NULL_ERROR ...
	REQUEST_NULL_ERROR = response.BIFBaseResponse{12001, "Request parameter cannot be null"}

	// INVALID_CONTRACTBALANCE_ERROR ...
	INVALID_CONTRACTBALANCE_ERROR = response.BIFBaseResponse{12002, "ContractBalance must be between 1 and Long.MAX_VALUE"}

	// INVALID_PRITX_FROM_ERROR ...
	INVALID_PRITX_FROM_ERROR = response.BIFBaseResponse{12003, "Invalid Private Transaction Sender"}

	// INVALID_PRITX_PAYLAOD_ERROR ...
	INVALID_PRITX_PAYLAOD_ERROR = response.BIFBaseResponse{12004, "Invalid Private Transaction payload"}

	// INVALID_PRITX_TO_ERROR ...
	INVALID_PRITX_TO_ERROR = response.BIFBaseResponse{12005, "Invalid Private Transaction recipient list"}

	// INVALID_PRITX_HASH_ERROR ...
	INVALID_PRITX_HASH_ERROR = response.BIFBaseResponse{12006, "Invalid Private Transaction Hash"}
)
