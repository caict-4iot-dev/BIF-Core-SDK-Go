package common

const (
	METADATA_KEY_MIN   int = 1
	METADATA_KEY_MAX   int = 1024
	METADATA_VALUE_MAX int = 256000

	HASH_HEX_LENGTH int = 64
	OPT_TYPE_MIN    int = 0
	OPT_TYPE_MAX    int = 2
	// 交易默认值
	GAS_PRICE int64 = 100
	FEE_LIMIT int64 = 1000000
	// 合约查询类型
	CONTRACT_QUERY_OPT_TYPE int = 2

	// 账号参数
	VERSION    int64 = 0
	SUCCESS    int   = 0
	ERRORCODE  int   = 4
	INIT_NONCE int64 = 0
	INIT_ZERO  int64 = 0
	INIT_ONE   int64 = 1

	INIT_ZERO_L int64 = 0
	INIT_ONE_L  int64 = 1
)
