package common

// OperationType
const (
	// UNKNOWN Unknown operation
	UNKNOWN = 0

	// ACCOUNT_ACTIVATE Activate an account
	ACCOUNT_ACTIVATE = 1

	// ACCOUNT_SET_METADATA Set metadata
	ACCOUNT_SET_METADATA = 2

	// ACCOUNT_SET_PRIVILEGE Set privilege
	ACCOUNT_SET_PRIVILEGE = 3

	// GAS_SEND Send gas
	GAS_SEND = 6

	// CONTRACT_CREATE Create contract
	CONTRACT_CREATE = 13

	// CONTRACT_INVOKE Invoke contract by sending
	CONTRACT_INVOKE = 15

	// PRIVATE_CONTRACT_CREATE Create Private Contract
	PRIVATE_CONTRACT_CREATE = 17

	// PRIVATE_CONTRACT_CALL Call Private Contract
	PRIVATE_CONTRACT_CALL = 18
)
