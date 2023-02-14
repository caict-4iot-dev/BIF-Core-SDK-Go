# BIF-Core-SDK-Go

# 1. BIF-Core-SDK-Go使用说明

​		本节详细说明SDK常用接口文档。星火链提供 GO SDK供开发者使用。

​        **github**代码库地址：github.com/caict-4iot-dev/BIF-Core-SDK-Go

## 1.1 SDK概述

### 1.1.1 名词解析

+ 账户服务： 提供账户相关的有效性校验、创建与查询接口

+ 合约服务： 提供合约相关的有效性校验、创建与查询接口

+ 交易服务： 提供构建交易及交易查询接口

+ 区块服务： 提供区块的查询接口

+ 账户nonce值： 每个账户都维护一个序列号，用于用户提交交易时标识交易执行顺序的

### 1.1.2 请求参数与相应数据格式

+ **请求参数**

​		接口的请求参数的类名，是\[服务名][方法名]Request，例如: 账户服务下的getAccount接口的请求参数格式是BIFAccountGetInfoRequest。

​		请求参数的成员，是各个接口的入参的成员。例如：账户服务下的getAccount接口的入参成员是address，那么该接口的请求参数的完整结构如下：

```go
    // BIFAccountGetInfoRequest 获取指定的账户信息请求体
    type BIFAccountGetInfoRequest struct {
        Address string `json:"address"` // 必填，待查询的区块链账户地址
    }
```

+ **响应数据**

​		接口的响应数据的类名，是\[服务名][方法名]Response，例如：账户服务下的getNonce接口的响应数据格式是BIFAccountGetNonceResponse。

​		响应数据的成员，包括错误码、错误描述和返回结果。响应数据的成员如下：

```go
    // BIFAccountGetNonceResponse 获取指定账户的nonce值返回体
    type BIFAccountGetNonceResponse struct {
        BIFBaseResponse
        Result BIFAccountGetNonceResult `json:"result"`
    }
    
    // BIFBaseResponse SDK响应返回结构体
    type BIFBaseResponse struct {
        ErrorCode int    `json:"error_code"`
        ErrorDesc string `json:"error_desc"`
    }

```

1. ErrorCode: 错误码。错误码为0表示响应正常，其他错误码请查阅[错误码详情](# 4.7 错误码)。
2. ErrorDesc: 错误描述。
3. Result: 返回结果。一个结构体，其类名是\[服务名][方法名]Result，其成员是各个接口返回值的成员，例如：账户服务下的getNonce接口的结果类名是BIFAccountGetNonceResult，成员有nonce, 完整结构如下：

```go
    type BIFAccountGetNonceResult struct {
        Nonce int64 `json:"nonce"`
    }
```

## 1.2 SDK使用方法

​		本节介绍SDK的使用流程。

​		首先需要生成SDK实现，然后调用相应服务的接口，其中服务包括账户服务、合约服务、交易服务和区块服务。

### 1.2.1 生成SDK实例

​		调用SDK的接口getInstance来实现，调用如下：

```go
    url := "http://test-bif-core.xinghuo.space"
    sdk, err := GetInstance(url)
```

### 1.2.2 生成公私钥地址

+ **Ed25519算法生成**

```go
    privateKeyManager, err := key.GetPrivateKeyManager(key.ED25519)
    if err != nil {
        return err
    }
    publicKeyManager, err := key.GetPublicKeyManager([]byte(privateKeyManager.EncPrivateKey), key.ED25519)
    if err != nil {
        return err
    }
    
    encAddress := publicKeyManager.EncAddress
    encPublicKey := publicKeyManager.EncPublicKey
    rawPublicKey := privateKeyManager.RawPublicKey
    encPrivateKey := privateKeyManager.EncPrivateKey
    rawPrivateKey := privateKeyManager.RawPrivateKey
```

+ **SM2算法生成**

```go
    keyPair, err := key.GetBidAndKeyPairBySM2()
    if err != nil {
        return err
    }
    encAddress := keyPair.GetEncAddress()
    encPublicKey := keyPair.GetEncPublicKey()
    encPrivateKey := keyPair.GetEncPrivateKey()
    rawPublicKey := keyPair.GetRawPublicKey()
    rawPrivateKey := keyPair.GetRawPrivateKey()
```

### 1.2.3 私钥对象使用

+ **构造对象**

```go
    // 私钥对象构造
    privateKeyManager, err := key.GetPrivateKeyManager(key.ED25519)
    if err != nil {
        return err
    }
```

+ **解析对象**

```go
    // 私钥对象构造
    privateKeyManager, err := key.GetPrivateKeyManager(key.ED25519)
    if err != nil {
        return err
    }
   
    // 星火私钥
    encPrivateKey := privateKeyManager.EncPrivateKey
    // 原生私钥
    rawPrivateKey := privateKeyManager.RawPrivateKey
```

+ **根据私钥获取公钥**

```go
    encPrivateKey := "priSrrk31MhNGEGAmnmZPH5K8fnuqTKLuLMvWd6E7TEdEjWkcQ"
    publicKeyManager, err := key.GetPublicKeyManager([]byte(encPrivateKey), key.ED25519)
    if err != nil {
        t.Error(err)
    }
    // 星火公钥
    encPublicKey := publicKeyManager.EncPublicKey
```

+ **原生私钥转星火私钥**

```go
    encPrivateKey := key.GetEncPrivateKey(rawPrivateKey, key.ED25519)
```

+ **原生公钥转星火公钥**

```go
    encPublicKey := key.EncPublicKey(rawPublicKey, key.ED25519)
```

+ **签名**

```go
    encPrivateKey := "priSPKhTMRa7SsQLc4wXUDrEZW5wSeKN68Xy5LuCYQmndS75SZ"
    msg := "hello word"
    // 签名
    signMsg, err := key.Sign([]byte(encPrivateKey), []byte(msg))
    if err != nil {
        return err
    }
```

### 1.2.4 公钥对象使用

+ **构造对象**

```go
    // 公钥对象构造
    publicKeyManager, err := key.GetPublicKeyManager([]byte(privateKeyManager.EncPrivateKey), key.ED25519)
    if err != nil {
        return err
    }
```

+ **获取账号地址**

```go
    encPrivateKey := "priSrrk31MhNGEGAmnmZPH5K8fnuqTKLuLMvWd6E7TEdEjWkcQ"
    publicKeyManager, err := key.GetPublicKeyManager([]byte(encPrivateKey), key.ED25519)
    if err != nil {
        return err
    }
    // 原生公钥
    rawPublicKey := publicKeyManager.RawPublicKey
    // 星火公钥
    encPublicKey := publicKeyManager.EncPublicKey
    // 星火地址
    encAddress := publicKeyManager.EncAddress
```

+ **账号地址校验**

```go
    address := "did:bid:efLrFZCn3wqSrozTG9MkxXbriRmwUHs5"
    isAddress := key.IsAddressValid(address)
```

+ **验签**

```go
    encPrivateKey := "priSPKhTMRa7SsQLc4wXUDrEZW5wSeKN68Xy5LuCYQmndS75SZ"
    msg := "hello word"
    // 签名
    signMsg, err := key.Sign([]byte(encPrivateKey), []byte(msg))
    if err != nil {
        return err
    }
    publicKeyManager, err := key.GetPublicKeyManager([]byte(encPrivateKey), key.ED25519)
    if err != nil {
        return err
    }

    // 验签
    isOK := key.Verify([]byte(publicKeyManager.EncPublicKey), []byte(msg), signMsg, ED25519)
    if !isOK {
        return errors.New("verify sign message is failed")
    }
```

### 1.2.5 密钥存储器

+ **生成密钥存储器**

```go
    key.GenerateKeyStore(encPrivateKey, password, n, r, p, version)
```

>  请求参数

| 参数          | 类型    | 描述                     |
| ------------- | ------- | ------------------------ |
| encPrivateKey | string  | 待存储的密钥，可为null   |
| password      | string  | 口令                     |
| n             | int | CPU消耗参数，必填且大于1 |
| r             | int | 内存消息参数，必填       |
| p             | int | 并行化参数，必填         |
| version       | int | 版本号，必填             |

> 响应数据

| 参数     | 类型        | 描述             |
| -------- | ----------- | ---------------- |
| keyStore | KeyStore | 存储密钥的存储器 |

> 示例

```go
    encPrivateKey := "priSPKrR4w6H89kRXaC2XZT5Lmj7XoCoBdvTuv7ySXSCDDGaZZ"
    password := "123456"
    n := 16384
    r := 8
    p := 1
    version := 32
    _, keyStore := key.GenerateKeyStore(encPrivateKey, password, n, r, p, version)
```

+ **解析密钥存储器**

```
    key.DecipherKeyStore(keyStore, password)
```

>  请求参数

| 参数     | 类型   | 描述             |
| -------- | ------ | ---------------- |
| password | string | 口令             |
| keyStore | string | 存储密钥的存储器 |

> 响应数据

| 参数          | 类型   | 描述           |
| ------------- | ------ | -------------- |
| encPrivateKey | string | 解析出来的密钥 |

> 示例

```go
    keyStore := key.KeyStore{
        Address:    "did:bid:ef24oGV9p46o1uwm2aFgZpScwW13r8nbu",
        AesctrIv:   "978071003b6d24f0b861048cbd4c008b",
        CypherText: "978071003b6d24f0b861048cbd4c008bdcc773c2dfbf58ab5e9fd01fc0893b342e387d701058d9f7f9e769ee5f36da750c598213b46b1fa6157e9e80e55a33e75a83",
        ScryptParams: key.ScryptParams{
            N:    16384,
            P:    1,
            R:    8,
            Salt: "8498081",
        },
        Version: 32,
    }
    password := "123456"
    encPrivateKey := key.DecipherKeyStore(keyStore, password)
    fmt.Println("encPrivateKey:", encPrivateKey)
```

### 1.2.6 助记词

+ **生成助记词**

> 响应数据

| 参数          | 类型         | 描述   |
| ------------- | ------------ | ------ |
| mnemonic | string | 助记词 |

> 示例

```go
    mnemonic, err := mnemonic.GenerateMnemonicCode()
    if err != nil {
        return err
    }
```

+ **根据助记词生成私钥**

> 请求参数

| 参数          | 类型         | 描述                      |
| ------------- | ------------ | ------------------------- |
| keyType       | int       | 选填，加密类型ED25519/SM2 |
| mnemonic | string | 必填，助记词              |
| hdPaths       | string | 必填，路径                |

> 响应数据

| 参数        | 类型         | 描述 |
| ----------- | ------------ | ---- |
| privateKeys | string | 私钥 |

> 示例

```go
    mnemonicStr := "style orchard science puppy place differ benefit thing wrap type build scare"
    hdPaths := "m/44'/526'/1'/0/0"
    keyType := key.ED25519
    encPrivateKey, err := mnemonic.GeneratePrivateKeys(mnemonicStr, hdPaths, keyType)
    if err != nil {
        return err
    }
    
    fmt.Println("encPrivateKey:", encPrivateKey)
```

## 1.3 账户服务接口列表

​		账户服务接口主要是账户相关的接口，目前有8个接口：

| 序号 | 接口                | 说明                                  |
| ---- | ------------------- | ------------------------------------- |
| 1    | CreateAccount       | 生成主链数字身份                      |
| 2    | GetAccount          | 该接口用于获取指定的账户信息          |
| 3    | GetNonce            | 该接口用于获取指定账户的nonce值       |
| 4    | GetAccountBalance   | 该接口用于获取指定账户的星火令的余额  |
| 5    | GetMetadatas        | 设置metadatas                         |
| 6    | GetAccountMetadatas | 该接口用于获取指定账户的metadatas信息 |
| 7    | SetPrivilege        | 设置权限                              |
| 8    | GetAccountPriv      | 获取账户权限                          |

### 1.3.1 CreateAccount

> 接口说明

```
该接口用于生成主链数字身份。
```

> 调用方法

```go
    CreateAccount(r request.BIFCreateAccountRequest) response.BIFCreateAccountResponse
```

> 请求参数

| 参数          | 类型   | 描述                                                         |
| ------------- | ------ | ------------------------------------------------------------ |
| senderAddress | string | 必填，交易源账号，即交易的发起方                             |
| privateKey    | string | 必填，交易源账户私钥                                         |
| ceilLedgerSeq | int64   | 选填，区块高度限制, 如果大于0，则交易只有在该区块高度之前（包括该高度）才有效 |
| remarks       | string | 选填，用户自定义给交易的备注                                 |
| destAddress   | string | 必填，目标账户地址                                           |
| initBalance   | int64   | 必填，初始化星火令，单位PT，1 星火令 = 10^8 PT, 大小(0, int64.MAX_VALUE] |
| gasPrice      | int64   | 选填，打包费用 (单位是PT)，默认100L                          |
| feeLimit      | int64   | 选填，交易花费的手续费(单位是PT)，默认1000000L               |

> 响应数据

| 参数 | 类型   | 描述     |
| ---- | ------ | -------- |
| hash | string | 交易hash |


> 错误码

| 异常                      | 错误码 | 描述                                             |
| ------------------------- | ------ | ------------------------------------------------ |
| INVALID_ADDRESS_ERROR     | 11006  | Invalid address                                  |
| REQUEST_NULL_ERROR        | 12001  | Request parameter cannot be null                 |
| PRIVATEKEY_NULL_ERROR     | 11057  | PrivateKeys cannot be empty                      |
| INVALID_DESTADDRESS_ERROR | 11003  | Invalid destAddress                              |
| INVALID_INITBALANCE_ERROR | 11004  | InitBalance must be between 1 and int64.MAX_VALUE |
| SYSTEM_ERROR              | 20000  | System error                                     |


> 示例

```go
    as := GetAccountInstance(SDK_INSTANCE_URL)
    senderAddress := "did:bid:efzewQxg38x2Tmb1cpxSC1ZWwMZUxUeV"
    senderPrivateKey := "priSPKhTMRa7SsQLc4wXUDrEZW5wSeKN68Xy5LuCYQmndS75SZ"
    destAddress := "did:bid:zf2AoXhJsmr1aaUMxhnKeMAX42G9Ck526"
    r := request.BIFCreateAccountRequest{
        SenderAddress: senderAddress,
        DestAddress:   destAddress,
        PrivateKey:    senderPrivateKey,
        InitBalance:   1000000,
        Remarks:       "init account",
    }
    
    res := as.CreateAccount(r)
    if res.ErrorCode != 0 {
        t.Error(res.ErrorDesc)
    }
    
    dataByte, err := json.Marshal(res)
    if err != nil {
        t.Error(err)
    }
    
    fmt.Println("res: ", string(dataByte))
```

### 1.3.2 GetAccount

> 接口说明

   	该接口用于获取指定的账户信息。

> 调用方法

```go
    GetAccount(r request.BIFAccountGetInfoRequest) response.BIFAccountGetInfoResponse
```

> 请求参数

| 参数    | 类型   | 描述                         |
| ------- | ------ | ---------------------------- |
| address | string | 必填，待查询的区块链账户地址 |

> 响应数据

| 参数    | 类型   | 描述                                           |
| ------- | ------ | ---------------------------------------------- |
| address | string | 账户地址                                       |
| balance | int64   | 账户余额，单位PT，1 星火令 = 10^8 PT, 必须大于0 |
| nonce   | int64   | 账户交易序列号，必须大于0                      |

> 错误码

| 异常                  | 错误码 | 描述                             |
| --------------------- | ------ | -------------------------------- |
| INVALID_ADDRESS_ERROR | 11006  | Invalid address                  |
| REQUEST_NULL_ERROR    | 12001  | Request parameter cannot be null |
| CONNECTNETWORK_ERROR  | 11007  | Failed to connect to the network |
| SYSTEM_ERROR          | 20000  | System error                     |

> 示例

```go
    as := GetAccountInstance(SDK_INSTANCE_URL)
    // 初始化请求参数
    accountAddress := "did:bid:ef21AHDJWnFfYQ3Qs3kMxo64jD2KATwBz"
    r := request.BIFAccountGetInfoRequest{
        Address: accountAddress,
    }
    res := as.GetAccount(r)
    if res.ErrorCode != 0 {
        t.Error(res.ErrorDesc)
    }
    
    dataByte, err := json.Marshal(res)
    if err != nil {
        t.Error(err)
    }
    
    fmt.Println("res: ", string(dataByte))
    }
```

### 1.3.3 GetNonce

> 接口说明

   	该接口用于获取指定账户的nonce值。

> 调用方法

```go
    GetNonce(r request.BIFAccountGetNonceRequest) response.BIFAccountGetNonceResponse
```

> 请求参数

| 参数    | 类型   | 描述                         |
| ------- | ------ | ---------------------------- |
| address | string | 必填，待查询的区块链账户地址 |

> 响应数据

| 参数  | 类型 | 描述           |
| ----- | ---- | -------------- |
| nonce | int64 | 账户交易序列号 |

> 错误码

| 异常                  | 错误码 | 描述                             |
| --------------------- | ------ | -------------------------------- |
| INVALID_ADDRESS_ERROR | 11006  | Invalid address                  |
| REQUEST_NULL_ERROR    | 12001  | Request parameter cannot be null |
| CONNECTNETWORK_ERROR  | 11007  | Failed to connect to the network |
| SYSTEM_ERROR          | 20000  | System error                     |

> 示例

```go
    as := GetAccountInstance(SDK_INSTANCE_URL)
    // 初始化请求参数
    accountAddress := "did:bid:ef21AHDJWnFfYQ3Qs3kMxo64jD2KATwBz"
    r := request.BIFAccountGetNonceRequest{
        Address: accountAddress,
    }
    res := as.GetNonce(r)
    if res.ErrorCode != 0 {
        t.Error(res.ErrorDesc)
    }
    
    dataByte, err := json.Marshal(res)
    if err != nil {
        t.Error(err)
    }
    
    fmt.Println("res: ", string(dataByte))
```

### 1.3.4 GetAccountBalance

> 接口说明

  	该接口用于获取指定账户的余额。

> 调用方法

```go
    GetAccountBalance(r request.BIFAccountGetBalanceRequest) response.BIFAccountGetBalanceResponse
```

> 请求参数

| 参数    | 类型   | 描述                         |
| ------- | ------ | ---------------------------- |
| address | string | 必填，待查询的区块链账户地址 |

> 响应数据

| 参数    | 类型 | 描述 |
| ------- | ---- | ---- |
| balance | int64 | 余额 |

> 错误码

| 异常                  | 错误码 | 描述                             |
| --------------------- | ------ | -------------------------------- |
| INVALID_ADDRESS_ERROR | 11006  | Invalid address                  |
| REQUEST_NULL_ERROR    | 12001  | Request parameter cannot be null |
| CONNECTNETWORK_ERROR  | 11007  | Failed to connect to the network |
| SYSTEM_ERROR          | 20000  | System error                     |

> 示例

```go
    as := GetAccountInstance(SDK_INSTANCE_URL)
    // 初始化请求参数
    accountAddress := "did:bid:ef21AHDJWnFfYQ3Qs3kMxo64jD2KATwBz"
    r := request.BIFAccountGetBalanceRequest{
        Address: accountAddress,
    }
    res := as.GetAccountBalance(r)
    if res.ErrorCode != 0 {
        t.Error(res.ErrorDesc)
    }
    
    dataByte, err := json.Marshal(res)
    if err != nil {
        t.Error(err)
    }
    
    fmt.Println("res: ", string(dataByte))
```

### 1.3.5 SetMetadatas

> 接口说明

   	该接口用于修改账户的metadatas信息。

> 调用方法

```go
    SetMetadatas(r request.BIFAccountSetMetadatasRequest) response.BIFAccountSetMetadatasResponse
```

> 请求参数

| 参数          | 类型    | 描述                                                         |
| ------------- | ------- | ------------------------------------------------------------ |
| senderAddress | string  | 必填，交易源账号，即交易的发起方                             |
| privateKey    | string  | 必填，交易源账户私钥                                         |
| ceilLedgerSeq | int64    | 选填，区块高度限制, 如果大于0，则交易只有在该区块高度之前（包括该高度）才有效 |
| remarks       | string  | 选填，用户自定义给交易的备注                                 |
| key           | string  | 必填，metadatas的关键词，长度限制[1, 1024]                   |
| value         | string  | 必填，metadatas的内容，长度限制[0, 256000]                   |
| version       | int64    | 选填，metadatas的版本                                        |
| deleteFlag    | bool | 选填，是否删除remarks                                        |
| gasPrice      | int64    | 选填，打包费用 (单位是PT)，默认100L                          |
| feeLimit      | int64    | 选填，交易花费的手续费(单位是PT)，默认1000000L               |

> 响应数据

| 参数 | 类型   | 描述     |
| ---- | ------ | -------- |
| hash | string | 交易hash |


> 错误码

| 异常                    | 错误码 | 描述                                             |
| ----------------------- | ------ | ------------------------------------------------ |
| INVALID_ADDRESS_ERROR   | 11006  | Invalid address                                  |
| REQUEST_NULL_ERROR      | 12001  | Request parameter cannot be null                 |
| PRIVATEKEY_NULL_ERROR   | 11057  | PrivateKeys cannot be empty                      |
| INVALID_DATAKEY_ERROR   | 11011  | The length of key must be between 1 and 1024     |
| INVALID_DATAVALUE_ERROR | 11012  | The length of value must be between 0 and 256000 |
| SYSTEM_ERROR            | 20000  | System error                                     |


> 示例

```go
    as := GetAccountInstance(SDK_INSTANCE_URL)
    // 初始化请求参数
    r := request.BIFAccountSetMetadatasRequest{
        SenderAddress: "did:bid:effMzw4pjqgVxpFZCQ3fVWN5n7USpRYu",
        PrivateKey:    "priSPKe86UJsnJ3WTDtLViP5ii8WTZKCXRMJmmqkDBWHq1eyMy",
        Remarks:       "set remarks",
        Key:           "20220101-01",
        Value:         "metadata-20220101-01",
    }
    res := as.SetMetadatas(r)
    if res.ErrorCode != 0 {
        t.Error(res.ErrorDesc)
    }
    
    dataByte, err := json.Marshal(res)
    if err != nil {
        t.Error(err)
    }
    
    fmt.Println("res: ", string(dataByte))
```

### 1.3.6 GetAccountMetadatas

> 接口说明

   	该接口用于获取指定账户的metadatas信息。

> 调用方法

```go
    GetAccountMetadatas(r request.BIFAccountGetMetadatasRequest) response.BIFAccountGetMetadatasResponse
```

> 请求参数

| 参数    | 类型   | 描述                                                         |
| ------- | ------ | ------------------------------------------------------------ |
| address | string | 必填，待查询的账户地址                                       |
| key     | string | 选填，metadatas关键字，长度限制[1, 1024]，有值为精确查找，无值为全部查找 |

> 响应数据

| 参数                 | 类型     | 描述              |
| -------------------- | -------- | ----------------- |
| metadatas            | object[] | 账户              |
| metadatas[i].key     | string   | metadatas的关键词 |
| metadatas[i].value   | string   | metadatas的内容   |
| metadatas[i].version | int64     | metadatas的版本   |


> 错误码

| 异常                  | 错误码 | 描述                                         |
| --------------------- | ------ | -------------------------------------------- |
| INVALID_ADDRESS_ERROR | 11006  | Invalid address                              |
| REQUEST_NULL_ERROR    | 12001  | Request parameter cannot be null             |
| CONNECTNETWORK_ERROR  | 11007  | Failed to connect to the network             |
| NO_METADATAS_ERROR    | 11010  | The account does not have the metadatas      |
| INVALID_DATAKEY_ERROR | 11011  | The length of key must be between 1 and 1024 |
| SYSTEM_ERROR          | 20000  | System error                                 |


> 示例

```go
    as := GetAccountInstance(SDK_INSTANCE_URL)
    // 初始化请求参数
    r := request.BIFAccountGetMetadatasRequest{
        Address: "did:bid:effMzw4pjqgVxpFZCQ3fVWN5n7USpRYu",
        Key:     "20220101-01",
    }
    res := as.GetAccountMetadatas(r)
    if res.ErrorCode != 0 {
        t.Error(res.ErrorDesc)
    }
    
    dataByte, err := json.Marshal(res)
    if err != nil {
        t.Error(err)
    }
    
    fmt.Println("res: ", string(dataByte))
```

### 1.3.7 SetPrivilege

> 接口说明

   	该接口用于设置权限。

> 调用方法

```go
    SetPrivilege(r request.BIFAccountSetPrivilegeRequest) response.BIFAccountSetPrivilegeResponse
```

> 请求参数

| 参数                    | 类型   | 描述                                                         |
| ----------------------- | ------ | ------------------------------------------------------------ |
| senderAddress           | string | 必填，交易源账号，即交易的发起方                             |
| privateKey              | string | 必填，交易源账户私钥                                         |
| ceilLedgerSeq           | int64   | 选填，区块高度限制, 如果大于0，则交易只有在该区块高度之前（包括该高度）才有效 |
| remarks                 | string | 选填，用户自定义给交易的备注                                 |
| signers                 | []BIFSigner   | 选填，签名者权重列表                                         |
| signers.address         | string | 签名者区块链账户地址                                         |
| signers.weight          | int64   | 为签名者设置权重值                                           |
| txThreshold             | string | 选填，交易门限，大小限制[0, int64.MAX_VALUE]                  |
| typeThreshold           | []BIFTypeThreshold   | 选填，指定类型交易门限                                       |
| typeThreshold.type      | int64   | 操作类型，必须大于0                                          |
| typeThreshold.threshold | int64   | 门限值，大小限制[0, int64.MAX_VALUE]                          |
| masterWeight            | string | 选填                                                         |
| gasPrice                | int64   | 选填，打包费用 (单位是PT)，默认100L                          |
| feeLimit                | int64   | 选填，交易花费的手续费(单位是PT)，默认1000000L               |

> 响应数据

| 参数 | 类型   | 描述     |
| ---- | ------ | -------- |
| hash | string | 交易hash |


> 错误码

| 异常                  | 错误码 | 描述                             |
| --------------------- | ------ | -------------------------------- |
| INVALID_ADDRESS_ERROR | 11006  | Invalid address                  |
| REQUEST_NULL_ERROR    | 12001  | Request parameter cannot be null |
| PRIVATEKEY_NULL_ERROR | 11057  | PrivateKeys cannot be empty      |
| SYSTEM_ERROR          | 20000  | System error                     |


> 示例

```go
    as := GetAccountInstance(SDK_INSTANCE_URL)
    // 初始化请求参数
    r := request.BIFAccountSetPrivilegeRequest{
        SenderAddress: "did:bid:effMzw4pjqgVxpFZCQ3fVWN5n7USpRYu", 
        PrivateKey:    "priSPKe86UJsnJ3WTDtLViP5ii8WTZKCXRMJmmqkDBWHq1eyMy",
        Remarks:       "set privilege",
        TxThreshold:   "0",
        //MasterWeight:  "1",
    }
    res := as.SetPrivilege(r)
    if res.ErrorCode != 0 {
        t.Error(res.ErrorDesc)
    }
    
    dataByte, err := json.Marshal(res)
    if err != nil {
        t.Error(err)
    }
    
    fmt.Println("res: ", string(dataByte))  
```

### 1.3.8 GetAccountPriv

> 接口说明

   	该接口用于获取指定的账户权限信息。

> 调用方法

```go
    GetAccountPriv(r request.BIFAccountPrivRequest) response.BIFAccountPrivResponse
```

> 请求参数

| 参数    | 类型   | 描述                         |
| ------- | ------ | ---------------------------- |
| address | string | 必填，待查询的区块链账户地址 |

> 响应数据

| 参数                                        | 类型     | 描述                   |
| ------------------------------------------- | -------- | ---------------------- |
| address                                     | string   | 账户地址               |
| priv                                        | Object   | 账户权限               |
| Priv.masterWeight                           | Object   | 账户自身权重，大小限制 |
| Priv.signers                                | Object   | 签名者权重列表         |
| Priv.signers[i].address                     | string   | 签名者区块链账户地址   |
| Priv.signers[i].weight                      | int64     | 签名者权重，大小限制   |
| Priv.Thresholds                             | Object   |                        |
| Priv.Thresholds.txThreshold                 | int64     | 交易默认门限，大小限制 |
| Priv.Thresholds.typeThresholds              | Object[] | 不同类型交易的门限     |
| Priv.Thresholds.typeThresholds[i].type      | int64     | 操作类型，必须大于0    |
| Priv.Thresholds.typeThresholds[i].threshold | int64     | 门限值，大小限制       |

> 错误码

| 异常                  | 错误码 | 描述                             |
| --------------------- | ------ | -------------------------------- |
| INVALID_ADDRESS_ERROR | 11006  | Invalid address                  |
| REQUEST_NULL_ERROR    | 12001  | Request parameter cannot be null |
| CONNECTNETWORK_ERROR  | 11007  | Failed to connect to the network |
| SYSTEM_ERROR          | 20000  | System error                     |

> 示例

```go
    as := GetAccountInstance(SDK_INSTANCE_URL)
    // 初始化请求参数
    r := request.BIFAccountPrivRequest{
        Address: "did:bid:effMzw4pjqgVxpFZCQ3fVWN5n7USpRYu",
    }
    res := as.GetAccountPriv(r)
    if res.ErrorCode != 0 {
        t.Error(res.ErrorDesc)
    }
    
    dataByte, err := json.Marshal(res)
    if err != nil {
        t.Error(err)
    }
    
    fmt.Println("res: ", string(dataByte))
```

## 1.4 合约服务接口列表

​		合约服务接口主要是合约相关的接口，目前有6个接口：

| 序号 | 接口                 | 说明                               |
| ---- | -------------------- | ---------------------------------- |
| 1    | CheckContractAddress | 该接口用于检测合约账户的有效性     |
| 2    | ContractCreate       | 创建合约                           |
| 3    | GetContractInfo      | 该接口用于查询合约代码             |
| 4    | GetContractAddress   | 该接口用于根据交易Hash查询合约地址 |
| 5    | ContractQuery        | 该接口用于调试合约代码             |
| 6    | ContractInvoke       | 合约调用                           |

### 1.4.1 CheckContractAddress

> 接口说明

   	该接口用于检测合约账户的有效性。

> 调用方法

```go
    CheckContractAddress(r request.BIFContractCheckValidRequest) response.BIFContractCheckValidResponse
```

> 请求参数

| 参数            | 类型   | 描述                 |
| --------------- | ------ | -------------------- |
| contractAddress | string | 待检测的合约账户地址 |

> 响应数据

| 参数    | 类型    | 描述     |
| ------- | ------- | -------- |
| isValid | bool | 是否有效 |

> 错误码

| 异常                          | 错误码 | 描述                             |
| ----------------------------- | ------ | -------------------------------- |
| INVALID_CONTRACTADDRESS_ERROR | 11037  | Invalid contract address         |
| REQUEST_NULL_ERROR            | 12001  | Request parameter cannot be null |
| SYSTEM_ERROR                  | 20000  | System error                     |

> 示例

```go
    bs := GetContractInstance(SDK_INSTANCE_URL)
    var r request.BIFContractCheckValidRequest
    r.ContractAddress = "did:bid:efWVypEKTQoVTunsdBDw8rp4uoG5Lsy5"
    res := bs.CheckContractAddress(r)
    if res.ErrorCode != 0 {
        t.Error(res.ErrorDesc)
    }
    
    dataByte, err := json.Marshal(res)
    if err != nil {
        t.Error(err)
    }
    
    fmt.Println("res: ", string(dataByte))
```

### 1.4.2 ContractCreate

> 接口说明

   	该接口用于创建合约。

> 调用方法

```go
    ContractCreate(r request.BIFContractCreateRequest) response.BIFContractCreateResponse
```

> 请求参数

| 参数          | 类型    | 描述                                                         |
| ------------- | ------- | ------------------------------------------------------------ |
| senderAddress | string  | 必填，交易源账号，即交易的发起方                             |
| gasPrice      | int64    | 选填，打包费用 (单位是PT)默认，默认100L                      |
| feeLimit      | int64    | 选填，交易花费的手续费(单位是PT)，默认1000000L               |
| privateKey    | string  | 必填，交易源账户私钥                                         |
| ceilLedgerSeq | int64    | 选填，区块高度限制, 如果大于0，则交易只有在该区块高度之前（包括该高度）才有效 |
| remarks       | string  | 选填，用户自定义给交易的备注                                 |
| initBalance   | int64    | 选填，给合约账户的初始化星火令，单位PT，1 星火令 = 10^8 PT, 大小限制[1, int64.MAX_VALUE] |
| type          | int | 选填，合约的类型，默认是0 , 0: javascript，1 :evm 。         |
| payload       | string  | 必填，对应语种的合约代码                                     |
| initInput     | string  | 选填，合约代码中init方法的入参                               |

> 响应数据

| 参数 | 类型   | 描述     |
| ---- | ------ | -------- |
| hash | string | 交易hash |


> 错误码

| 异常                      | 错误码 | 描述                                             |
| ------------------------- | ------ | ------------------------------------------------ |
| INVALID_ADDRESS_ERROR     | 11006  | Invalid address                                  |
| REQUEST_NULL_ERROR        | 12001  | Request parameter cannot be null                 |
| PRIVATEKEY_NULL_ERROR     | 11057  | PrivateKeys cannot be empty                      |
| INVALID_INITBALANCE_ERROR | 11004  | InitBalance must be between 1 and int64.MAX_VALUE |
| PAYLOAD_EMPTY_ERROR       | 11044  | Payload cannot be empty                          |
| INVALID_FEELIMIT_ERROR    | 11050  | FeeLimit must be between 0 and int64.MAX_VALUE    |
| SYSTEM_ERROR              | 20000  | System error                                     |


> 示例

```go
    bs := GetContractInstance(SDK_INSTANCE_URL)
    var r request.BIFContractCreateRequest
    senderAddress := "did:bid:efzewQxg38x2Tmb1cpxSC1ZWwMZUxUeV"
    senderPrivateKey := "priSPKhTMRa7SsQLc4wXUDrEZW5wSeKN68Xy5LuCYQmndS75SZ"
    payload := "\"use strict\"; function init(bar){return;} function main(input){let para = JSON.parse(input);if(para.do_foo){let x = {'hello' : 'world'};}} function query(input){return input;}"
    r.SenderAddress = senderAddress
    r.PrivateKey = senderPrivateKey
    r.Metadata = "create contract"
    r.Payload = payload
    r.InitBalance = 1
    r.Type = 0
    r.InitBalance = 1
    r.FeeLimit = 10000000000
    
    res := bs.ContractCreate(r)
    if res.ErrorCode != 0 {
        t.Error(res.ErrorDesc)
    }
    
    dataByte, err := json.Marshal(res)
    if err != nil {
        t.Error(err)
    }
    
    fmt.Println("res: ", string(dataByte))
```

### 1.4.3 GetContractInfo

> 接口说明

   	该接口用于查询合约代码。

> 调用方法

```go
    GetContractInfo(r request.BIFContractGetInfoRequest) response.BIFContractGetInfoResponse
```

> 请求参数

| 参数            | 类型   | 描述                 |
| --------------- | ------ | -------------------- |
| contractAddress | string | 待查询的合约账户地址 |

> 响应数据

| 参数                 | 类型    | 描述            |
| -------------------- | ------- | --------------- |
| contract             | object  | 合约信息        |
| contractInfo.type    | int | 合约类型，默认0 |
| contractInfo.payload | string  | 合约代码        |

> 错误码

| 异常                                      | 错误码 | 描述                                      |
| ----------------------------------------- | ------ | ----------------------------------------- |
| INVALID_CONTRACTADDRESS_ERROR             | 11037  | Invalid contract address                  |
| CONTRACTADDRESS_NOT_CONTRACTACCOUNT_ERROR | 11038  | contractAddress is not a contract account |
| NO_SUCH_TOKEN_ERROR                       | 11030  | No such token                             |
| GET_TOKEN_INFO_ERROR                      | 11066  | Failed to get token info                  |
| REQUEST_NULL_ERROR                        | 12001  | Request parameter cannot be null          |
| SYSTEM_ERROR                              | 20000  | System error                              |

> 示例

```go
    bs := GetContractInstance(SDK_INSTANCE_URL)
    var r request.BIFContractGetInfoRequest
    r.ContractAddress = "did:bid:efWVypEKTQoVTunsdBDw8rp4uoG5Lsy5"
    res := bs.GetContractInfo(r)
    if res.ErrorCode != 0 {
        t.Error(res.ErrorDesc)
    }
    
    dataByte, err := json.Marshal(res)
    if err != nil {
        t.Error(err)
    }
    
    fmt.Println("res: ", string(dataByte))
```

### 1.4.4 GetContractAddress

> 接口说明

```
该接口用于根据交易Hash查询合约地址。
```

> 调用方法

```go
    GetContractAddress(r request.BIFContractGetAddressRequest) response.BIFContractGetAddressResponse
```

> 请求参数

| 参数 | 类型   | 描述               |
| ---- | ------ | ------------------ |
| hash | string | 创建合约交易的hash |

> 响应数据

| 参数                                                        | 类型                      | 描述           |
| ----------------------------------------------------------- | ------------------------- | -------------- |
| contractAddressInfos                                        | []ContractAddressInfo | 合约地址列表   |
| contractAddressInfos[i].ContractAddressInfo                 | object                    | 成员           |
| contractAddressInfos[i].ContractAddressInfo.contractAddress | string                    | 合约地址       |
| contractAddressInfos[i].ContractAddressInfo.operationIndex  | int                   | 所在操作的下标 |

> 错误码

| 异常                 | 错误码 | 描述                             |
| -------------------- | ------ | -------------------------------- |
| INVALID_HASH_ERROR   | 11055  | Invalid transaction hash         |
| CONNECTNETWORK_ERROR | 11007  | Failed to connect to the network |
| REQUEST_NULL_ERROR   | 12001  | Request parameter cannot be null |
| SYSTEM_ERROR         | 20000  | System error                     |

> 示例

```go
    bs := GetContractInstance(SDK_INSTANCE_URL)
    var r request.BIFContractGetAddressRequest
    r.Hash = "ff6a9d1a0c0011fbb9f51cfb99e4cd5e7c31380046fda3fd6e0daae44d1d4648"
    res := bs.GetContractAddress(r)
    if res.ErrorCode != 0 {
        t.Error(res.ErrorDesc)
    }
    
    dataByte, err := json.Marshal(res)
    if err != nil {
        t.Error(err)
    }
    
    fmt.Println("res: ", string(dataByte))

```

### 1.4.5 ContractQuery

> 接口说明

   	该接口用于调用合约查询接口。

> 调用方法

```go
    ContractQuery(r request.BIFContractCallRequest) response.BIFContractCallResponse
```

> 请求参数

| 参数            | 类型   | 描述                                           |
| --------------- | ------ | ---------------------------------------------- |
| sourceAddress   | string | 选填，合约触发账户地址                         |
| contractAddress | string | 必填，合约账户地址                             |
| input           | string | 选填，合约入参                                 |
| gasPrice        | int64   | 选填，打包费用 (单位是PT)默认，默认100L        |
| feeLimit        | int64   | 选填，交易花费的手续费(单位是PT)，默认1000000L |


> 响应数据

| 参数      | 类型      | 描述       |
| --------- | --------- | ---------- |
| queryRets | JSONArray | 查询结果集 |

> 错误码

| 异常                                      | 错误码 | 描述                                             |
| ----------------------------------------- | ------ | ------------------------------------------------ |
| INVALID_SOURCEADDRESS_ERROR               | 11002  | Invalid sourceAddress                            |
| INVALID_CONTRACTADDRESS_ERROR             | 11037  | Invalid contract address                         |
| SOURCEADDRESS_EQUAL_CONTRACTADDRESS_ERROR | 11040  | SourceAddress cannot be equal to contractAddress |
| REQUEST_NULL_ERROR                        | 12001  | Request parameter cannot be null                 |
| CONNECTNETWORK_ERROR                      | 11007  | Failed to connect to the network                 |
| SYSTEM_ERROR                              | 20000  | System error                                     |

> 示例

```go
    bs := GetContractInstance(SDK_INSTANCE_URL)
    var r request.BIFContractCallRequest
    r.ContractAddress = "did:bid:efWVypEKTQoVTunsdBDw8rp4uoG5Lsy5"
    res := bs.ContractQuery(r)
    if res.ErrorCode != 0 {
        t.Error(res.ErrorDesc)
    }
    
    dataByte, err := json.Marshal(res)
    if err != nil {
        t.Error(err)
    }
    
    fmt.Println("res: ", string(dataByte))
```

### 1.4.6 ContractInvoke

> 接口说明

   	该接口用于合约调用。

> 调用方法

```go
    ContractInvoke(r request.BIFContractInvokeRequest) response.BIFContractInvokeResponse
```

> 请求参数

| 参数            | 类型   | 描述                                                         |
| --------------- | ------ | ------------------------------------------------------------ |
| senderAddress   | string | 必填，交易源账号，即交易的发起方                             |
| gasPrice        | int64   | 选填，打包费用 (单位是PT)默认，默认100L                      |
| feeLimit        | int64   | 选填，交易花费的手续费(单位是PT)，默认1000000L               |
| privateKey      | string | 必填，交易源账户私钥                                         |
| ceilLedgerSeq   | int64   | 选填，区块高度限制, 如果大于0，则交易只有在该区块高度之前（包括该高度）才有效 |
| remarks         | string | 选填，用户自定义给交易的备注                                 |
| contractAddress | string | 必填，合约账户地址                                           |
| BIFAmount       | int64   | 必填，转账金额                                               |
| input           | string | 选填，待触发的合约的main()入参                               |

> 响应数据

| 参数 | 类型   | 描述     |
| ---- | ------ | -------- |
| hash | string | 交易hash |


> 错误码

| 异常                          | 错误码 | 描述                                          |
| ----------------------------- | ------ | --------------------------------------------- |
| INVALID_ADDRESS_ERROR         | 11006  | Invalid address                               |
| REQUEST_NULL_ERROR            | 12001  | Request parameter cannot be null              |
| PRIVATEKEY_NULL_ERROR         | 11057  | PrivateKeys cannot be empty                   |
| INVALID_CONTRACTADDRESS_ERROR | 11037  | Invalid contract address                      |
| INVALID_AMOUNT_ERROR          | 11024  | Amount must be between 0 and int64.MAX_VALUE   |
| INVALID_FEELIMIT_ERROR        | 11050  | FeeLimit must be between 0 and int64.MAX_VALUE |
| SYSTEM_ERROR                  | 20000  | System error                                  |


> 示例

```go
    bs := GetContractInstance(SDK_INSTANCE_URL)
    var r request.BIFContractInvokeRequest
    senderAddress := "did:bid:efzewQxg38x2Tmb1cpxSC1ZWwMZUxUeV"
    contractAddress := "did:bid:efWVypEKTQoVTunsdBDw8rp4uoG5Lsy5"
    senderPrivateKey := "priSPKhTMRa7SsQLc4wXUDrEZW5wSeKN68Xy5LuCYQmndS75SZ"
    
    r.SenderAddress = senderAddress
    r.PrivateKey = senderPrivateKey
    r.ContractAddress = contractAddress
    r.BIFAmount = 1
    r.Metadata = "contract invoke"
    
    res := bs.ContractInvoke(r)
    if res.ErrorCode != 0 {
        t.Error(res.ErrorDesc)
    }
    
    dataByte, err := json.Marshal(res)
    if err != nil {
        t.Error(err)
    }
    
    fmt.Println("res: ", string(dataByte))
```

## 1.5 交易服务接口列表

​		交易服务接口主要是交易相关的接口，目前有6个接口：

| 序号 | 接口                  | 说明                               |
| ---- | --------------------- | ---------------------------------- |
| 1    | GasSend               | 交易                               |
| 2    | PrivateContractCreate | 私有化交易-合约创建                |
| 3    | PrivateContractCall   | 私有化交易-合约调用                |
| 4    | GetTransactionInfo    | 该接口用于实现根据交易hash查询交易 |
| 5    | EvaluateFee           | 该接口实现交易的费用评估           |
| 6    | BIFSubmit             | 提交交易                           |

### 1.5.1 GasSend

> 接口说明

   	该接口用于发起交易。

> 调用方法

```go
    GasSend(r request.BIFTransactionGasSendRequest) response.BIFTransactionGasSendResponse
```

> 请求参数

| 参数          | 类型   | 描述                                                         |
| ------------- | ------ | ------------------------------------------------------------ |
| senderAddress | string | 必填，交易源账号，即交易的发起方                             |
| privateKey    | string | 必填，交易源账户私钥                                         |
| ceilLedgerSeq | int64   | 选填，区块高度限制, 如果大于0，则交易只有在该区块高度之前（包括该高度）才有效 |
| remarks       | string | 选填，用户自定义给交易的备注                                 |
| destAddress   | string | 必填，发起方地址                                             |
| amount        | int64   | 必填，转账金额                                               |
| gasPrice      | int64   | 选填，打包费用 (单位是PT)，默认100L                          |
| feeLimit      | int64   | 选填，交易花费的手续费(单位是PT)，默认1000000L               |

> 响应数据

| 参数 | 类型   | 描述     |
| ---- | ------ | -------- |
| hash | string | 交易hash |


> 错误码

| 异常                      | 错误码 | 描述                                           |
| ------------------------- | ------ | ---------------------------------------------- |
| INVALID_ADDRESS_ERROR     | 11006  | Invalid address                                |
| REQUEST_NULL_ERROR        | 12001  | Request parameter cannot be null               |
| PRIVATEKEY_NULL_ERROR     | 11057  | PrivateKeys cannot be empty                    |
| INVALID_DESTADDRESS_ERROR | 11003  | Invalid destAddress                            |
| INVALID_GAS_AMOUNT_ERROR  | 11026  | BIFAmount must be between 0 and int64.MAX_VALUE |
| SYSTEM_ERROR              | 20000  | System error                                   |


> 示例

```go
    ts := GetTransactionInstance(SDK_INSTANCE_URL)
    var r request.BIFTransactionGasSendRequest
    r.SenderAddress = "did:bid:zf2AoXhJsmr1aaUMxhnKeMAX42G9Ck526"
    r.PrivateKey = "priSrrk31MhNGEGAmnmZPH5K8fnuqTKLuLMvWd6E7TEdEjWkcQ"
    r.DestAddress = "did:bid:efzewQxg38x2Tmb1cpxSC1ZWwMZUxUeV"
    r.Metadata = "gas send"
    r.Amount = 100000
    
    res := ts.GasSend(r)
    if res.ErrorCode != 0 {
        t.Error(res.ErrorDesc)
    }
    
    dataByte, err := json.Marshal(res)
    if err != nil {
        t.Error(err)
    }
    
    fmt.Println("res: ", string(dataByte))
```

### 1.5.2 PrivateContractCreate(Deprecated)

> 接口说明

```
该接口用于私有化交易的合约创建。
```

> 调用方法

```go
    PrivateContractCreate(r request.BIFTransactionPrivateContractCreateRequest) response.BIFTransactionPrivateContractCreateResponse
```

> 请求参数

| 参数          | 类型     | 描述                                                         |
| ------------- | -------- | ------------------------------------------------------------ |
| senderAddress | string   | 必填，交易源账号，即交易的发起方                             |
| privateKey    | string   | 必填，交易源账户私钥                                         |
| ceilLedgerSeq | int64     | 选填，区块高度限制, 如果大于0，则交易只有在该区块高度之前（包括该高度）才有效 |
| remarks       | string   | 选填，用户自定义给交易的备注                                 |
| type          | int  | 选填，合约的语种                                             |
| payload       | string   | 必填，对应语种的合约代码                                     |
| from          | string   | 必填，发起方加密机公钥                                       |
| to            | []string | 必填，接收方加密机公钥                                       |
| gasPrice      | int64     | 选填，打包费用 (单位是PT)默认，默认100L                      |
| feeLimit      | int64     | 选填，交易花费的手续费(单位是PT)，默认1000000L               |

> 响应数据

| 参数 | 类型   | 描述     |
| ---- | ------ | -------- |
| hash | string | 交易hash |


> 错误码

| 异常                        | 错误码 | 描述                             |
| --------------------------- | ------ | -------------------------------- |
| INVALID_ADDRESS_ERROR       | 11006  | Invalid address                  |
| REQUEST_NULL_ERROR          | 12001  | Request parameter cannot be null |
| PRIVATEKEY_NULL_ERROR       | 11057  | PrivateKeys cannot be empty      |
| INVALID_CONTRACT_TYPE_ERROR | 11047  | Invalid contract type            |
| PAYLOAD_EMPTY_ERROR         | 11044  | Payload cannot be empty          |
| SYSTEM_ERROR                | 20000  | System error                     |


> 示例

```go
    ts := GetTransactionInstance(SDK_INSTANCE_URL)
    var r request.BIFTransactionPrivateContractCreateRequest
    r.SenderAddress = "did:bid:efnVUgqQFfYeu97ABf6sGm3WFtVXHZB2"
    r.PrivateKey = "priSPKkWVk418PKAS66q4bsiE2c4dKuSSafZvNWyGGp2sJVtXL"
    r.Payload = "\"use strict\";function queryBanance(address)\r\n{return \" test query private contract sdk_3\";}\r\nfunction sendTx(to,amount)\r\n{return Chain.payCoin(to,amount);}\r\nfunction init(input)\r\n{return;}\r\nfunction main(input)\r\n{let args=JSON.parse(input);if(args.method===\"sendTx\"){return sendTx(args.params.address,args.params.amount);}}\r\nfunction query(input)\r\n{let args=JSON.parse(input);if(args.method===\"queryBanance\"){return queryBanance(args.params.address);}}"
    r.From = "sX46dMvKzKgH/SByjBs0uCROD9paCc/tF6WwcgUx3nA="
    r.To = []string{"Pz8tQqi4DZcL5Vrh/GXS20vZ4oqaiNyFxG0B9xAJmhw="}
    r.Metadata = "init account"
    r.Type = 0
    
    res := ts.PrivateContractCreate(r)
    if res.ErrorCode != 0 {
    t.Error(res.ErrorDesc)
    }
    
    dataByte, err := json.Marshal(res)
    if err != nil {
    t.Error(err)
    }
    
    // "hash":"278a1f189e235fee846baf22b4ff699e702f9dd407d2361bbfb41159d57f4a2f"
    fmt.Println("res: ", string(dataByte))
```

### 1.5.3 PrivateContractCall(Deprecated)

> 接口说明

   	该接口用于私有化交易的合约调用。

> 调用方法

```go
    PrivateContractCall(r request.BIFTransactionPrivateContractCallRequest) response.BIFTransactionPrivateContractCallResponse
```

> 请求参数

| 参数          | 类型     | 描述                                                         |
| ------------- | -------- | ------------------------------------------------------------ |
| senderAddress | string   | 必填，交易源账号，即交易的发起方                             |
| privateKey    | string   | 必填，交易源账户私钥                                         |
| ceilLedgerSeq | int64     | 选填，区块高度限制, 如果大于0，则交易只有在该区块高度之前（包括该高度）才有效 |
| remarks       | string   | 选填，用户自定义给交易的备注                                 |
| destAddress   | string   | 必填，发起方地址                                             |
| type          | int  | 选填，合约的语种（待用）                                     |
| input         | string   | 必填，待触发的合约的main()入参                               |
| from          | string   | 必填，发起方加密机公钥                                       |
| to            | []string | 必填，接收方加密机公钥                                       |
| gasPrice      | int64     | 选填，打包费用 (单位是PT)默认，默认100L                      |
| feeLimit      | int64     | 选填，交易花费的手续费(单位是PT)，默认1000000L               |

> 响应数据

| 参数 | 类型   | 描述     |
| ---- | ------ | -------- |
| hash | string | 交易hash |


> 错误码

| 异常                        | 错误码 | 描述                             |
| --------------------------- | ------ | -------------------------------- |
| INVALID_ADDRESS_ERROR       | 11006  | Invalid address                  |
| REQUEST_NULL_ERROR          | 12001  | Request parameter cannot be null |
| PRIVATEKEY_NULL_ERROR       | 11057  | PrivateKeys cannot be empty      |
| INVALID_CONTRACT_TYPE_ERROR | 11047  | Invalid contract type            |
| SYSTEM_ERROR                | 20000  | System error                     |


> 示例

```go
    ts := GetTransactionInstance(SDK_INSTANCE_URL)
    var r request.BIFTransactionPrivateContractCallRequest
    r.SenderAddress = "did:bid:efnVUgqQFfYeu97ABf6sGm3WFtVXHZB2"
    r.PrivateKey = "priSPKkWVk418PKAS66q4bsiE2c4dKuSSafZvNWyGGp2sJVtXL"
    r.Input = "{\"method\":\"queryBanance\",\"params\":{\"address\":\"567890哈哈=======\"}}"
    r.From = "sX46dMvKzKgH/SByjBs0uCROD9paCc/tF6WwcgUx3nA="
    r.To = []string{"Pz8tQqi4DZcL5Vrh/GXS20vZ4oqaiNyFxG0B9xAJmhw="}
    r.DestAddress = "did:bid:efTuswkPE1HP9Uc7vpNbRVokuQqhxaCE"
    r.Metadata = "init account"
    r.Type = 0
    
    res := ts.PrivateContractCall(r)
    if res.ErrorCode != 0 {
        t.Error(res.ErrorDesc)
    }
    
    dataByte, err := json.Marshal(res)
    if err != nil {
        t.Error(err)
    }
    
    // "hash":"278a1f189e235fee846baf22b4ff699e702f9dd407d2361bbfb41159d57f4a2f"
    fmt.Println("res: ", string(dataByte))
```

### 1.5.4 GetTransactionInfo

> 接口说明

   	该接口用于实现根据交易hash查询交易。

> 调用方法

```go
    GetTransactionInfo(r request.BIFTransactionGetInfoRequest) response.BIFTransactionGetInfoResponse
```

> 请求参数

| 参数 | 类型   | 描述     |
| ---- | ------ | -------- |
| hash | string | 交易hash |

> 响应数据

| 参数                              | 类型               | 描述           |
| --------------------------------- | ------------------ | -------------- |
| totalCount                        | int64               | 返回的总交易数 |
| transactions                      | TransactionHistory | 交易内容       |
| transactions.fee                  | string             | 交易实际费用   |
| transactions.confirmTime          | int64               | 交易确认时间   |
| transactions.errorCode            | int64               | 交易错误码     |
| transactions.errorDesc            | string             | 交易描述       |
| transactions.hash                 | string             | 交易hash       |
| transactions.ledgerSeq            | int64               | 区块序列号     |
| transactions.transaction          | TransactionInfo    | 交易内容列表   |
| transactions.signatures           | Signature          | 签名列表       |
| transactions.signatures.signData  | int64               | 签名后数据     |
| transactions.signatures.publicKey | int64               | 公钥           |
| transactions.txSize               | int64               | 交易大小       |

> 错误码

| 异常                 | 错误码 | 描述                             |
| -------------------- | ------ | -------------------------------- |
| INVALID_HASH_ERROR   | 11055  | Invalid transaction hash         |
| REQUEST_NULL_ERROR   | 12001  | Request parameter cannot be null |
| CONNECTNETWORK_ERROR | 11007  | Failed to connect to the network |
| SYSTEM_ERROR         | 20000  | System error                     |

> 示例

```go
    ts := GetTransactionInstance(SDK_INSTANCE_URL)
    var r request.BIFTransactionGetInfoRequest
    r.Hash = "2c0a445f603bdef7e4cfe5f63650f201cda3315b7c560edb79e3fcef611c5f8e"
    res := ts.GetTransactionInfo(r)
    if res.ErrorCode != 0 {
        t.Error(res.ErrorDesc)
    }
    
    dataByte, err := json.Marshal(res)
    if err != nil {
        t.Error(err)
    }
    
    fmt.Println("res: ", string(dataByte))
```

### 1.5.5 EvaluateFee

> 接口说明

   	该接口实现交易的费用评估。

> 调用方法

```go
    EvaluateFee(r request.BIFTransactionEvaluateFeeRequest) response.BIFTransactionEvaluateFeeResponse
```

> 请求参数

| 参数            | 类型                            | 描述                                                         |
| --------------- | ------------------------------- | ------------------------------------------------------------ |
| signatureNumber | int                         | 选填，待签名者的数量，默认是1，大小限制[1, int.MAX_VALUE] |
| remarks         | string                          | 选填，用户自定义给交易的备注                                 |
| Operations       | []BIFBaseOperation | 必填，待提交的操作，不能为空                                 |
| gasPrice        | int64                            | 选填，打包费用 (单位是PT)                                    |
| feeLimit        | int64                            | 选填，交易花费的手续费(单位是PT)                             |

#### BIFBaseOperation

| 序号 | 操作                              | 描述                             |
| ---- | --------------------------------- | -------------------------------- |
| 1    | BIFAccountActivateOperation       | 生成主链数字身份                 |
| 2    | BIFAccountSetMetadataOperation    | 修改账户的metadatas信息          |
| 3    | BIFAccountSetPrivilegeOperation   | 设置权限                         |
| 4    | BIFContractCreateOperation        | 创建合约（暂不支持EVM 合约）     |
| 5    | BIFContractInvokeOperation        | 合约调用（暂不支持EVM 合约）     |
| 6    | BIFGasSendOperation               | 发起交易                         |
| 7    | BIFPrivateContractCallOperation   | 私有化交易的合约创建--Deprecated |
| 8    | BIFPrivateContractCreateOperation | 私有化交易的合约调用--Deprecated |

> 响应数据

| 参数 | 类型                | 描述       |
| ---- | ------------------- | ---------- |
| txs  | []BIFTestTx | 评估交易集 |

#### BIFTestTx

| 成员变量       | 类型                                        | 描述         |
| -------------- | ------------------------------------------- | ------------ |
| transactionEnv | [TestTransactionFees](#TestTransactionFees) | 评估交易费用 |

#### TestTransactionFees

| 成员变量        | 类型                                | 描述     |
| --------------- | ----------------------------------- | -------- |
| transactionFees | [TransactionFees](#TransactionFees) | 交易费用 |

#### TransactionFees

| 成员变量 | 类型 | 描述               |
| -------- | ---- | ------------------ |
| feeLimit | int64 | 交易要求的最低费用 |
| gasPrice | int64 | 交易燃料单价       |

> 错误码

| 异常                          | 错误码 | 描述                                                    |
| ----------------------------- | ------ | ------------------------------------------------------- |
| INVALID_SOURCEADDRESS_ERROR   | 11002  | Invalid sourceAddress                                   |
| OPERATIONS_EMPTY_ERROR        | 11051  | Operations cannot be empty                              |
| OPERATIONS_ONE_ERROR          | 11053  | One of the operations cannot be resolved                |
| INVALID_SIGNATURENUMBER_ERROR | 11054  | SignagureNumber must be between 1 and int.MAX_VALUE |
| REQUEST_NULL_ERROR            | 12001  | Request parameter cannot be null                        |
| SYSTEM_ERROR                  | 20000  | System error                                            |

> 示例

```go
    ts := GetTransactionInstance(SDK_INSTANCE_URL)
    var r request.BIFTransactionEvaluateFeeRequest
    senderAddresss := "did:bid:zf2AoXhJsmr1aaUMxhnKeMAX42G9Ck526"
    destAddress := "did:bid:efzewQxg38x2Tmb1cpxSC1ZWwMZUxUeV"
    bifAmount := 10
    
    operation := request.BIFGasSendOperation{
    DestAddress: destAddress,
    Amount:      int64(bifAmount),
    BIFBaseOperation: request.BIFBaseOperation{
    OperationType: common.GAS_SEND,
    },
    }
    r.SourceAddress = senderAddresss
    r.Operation = operation
    r.SignatureNumber = 1
    
    res := ts.EvaluateFee(r)
    if res.ErrorCode != 0 {
    t.Error(res.ErrorDesc)
    }
    
    dataByte, err := json.Marshal(res)
    if err != nil {
    t.Error(err)
    }
    
    fmt.Println("res: ", string(dataByte))
```



### 1.5.6 BIFSubmit

> 接口说明

   	该接口用于交易提交。

> 调用方法

```go
BIFSubmit(r request.BIFTransactionSubmitRequest) response.BIFTransactionSubmitResponse
```

> 请求参数

| 参数          | 类型   | 描述               |
| ------------- | ------ | ------------------ |
| serialization | string | 必填，交易序列化值 |
| signData      | string | 必填，签名数据     |
| privateKey    | string | 必填，签名者私钥   |

> 响应数据

| 参数 | 类型   | 描述     |
| ---- | ------ | -------- |
| hash | string | 交易hash |

> 错误码

| 异常                        | 错误码 | 描述                             |
| --------------------------- | ------ | -------------------------------- |
| INVALID_SERIALIZATION_ERROR | 11056  | Invalid serialization            |
| SIGNATURE_EMPTY_ERROR       | 11067  | The signatures cannot be empty   |
| SIGNDATA_NULL_ERROR         | 11059  | SignData cannot be empty         |
| PUBLICKEY_NULL_ERROR        | 11061  | PublicKey cannot be empty        |
| REQUEST_NULL_ERROR          | 12001  | Request parameter cannot be null |
| SYSTEM_ERROR                | 20000  | System error                     |

> 示例

```go
    submitRequest := request.BIFTransactionSubmitRequest{
        Serialization: hex.EncodeToString(blob),
        SignData:      hex.EncodeToString(signData),
        PublicKey:     pubKey,
    }
    
    res := ts.BIFSubmit(submitRequest)
    if res.ErrorCode != 0 {
        t.Error(res.ErrorDesc)
    }
    
    dataByte, err = json.Marshal(res)
    if err != nil {
        t.Error(err)
    }
    
    fmt.Println("res: ", string(dataByte))
```


### 1.5.7 getTxCacheSize

> 接口说明

   	该接口用于获取交易池中交易条数。

> 调用方法

```go
GetTxCacheSize() response.BIFTransactionGetTxCacheSizeResponse
```

> 响应数据

| 参数       | 类型 | 描述                 |
| ---------- | ---- | -------------------- |
| queue_size | int64 | 返回交易池中交易条数 |

> 错误码

| 异常                 | 错误码 | 描述                             |
| -------------------- | ------ | -------------------------------- |
| CONNECTNETWORK_ERROR | 11007  | Failed to connect to the network |
| SYSTEM_ERROR         | 20000  | System error                     |

> 示例

```go
    res := ts.GetTxCacheSize()
    if res.ErrorCode != 0 {
        t.Error(res.ErrorDesc)
    }
    
    dataByte, err := json.Marshal(res)
    if err != nil {
        t.Error(err)
    }
    
    fmt.Println("res: ", string(dataByte))
```

### 1.5.8 getTxCacheData

> 接口说明

   	该接口用于获取交易池中交易数据。

> 调用方法

```go
GetTxCacheData(r request.BIFTransactionCacheRequest) response.BIFTransactionCacheResponse {
```

> 请求参数

| 参数 | 类型   | 描述           |
| ---- | ------ | -------------- |
| hash | string | 选填，交易hash |

> 响应数据

| 参数                           | 类型     | 描述                 |
| ------------------------------ | -------- | -------------------- |
| transactions                   | Object[] | 返回交易池中交易数据 |
| transactionsp[i].hash          | String   | 交易hash             |
| transactionsp[i].incoming_time | String   | 进入时间             |
| transactionsp[i].status        | String   | 状态                 |
| transactionsp[i].transaction   | Object   |                      |

> 错误码

| 异常                 | 错误码 | 描述                             |
| -------------------- | ------ | -------------------------------- |
| CONNECTNETWORK_ERROR | 11007  | Failed to connect to the network |
| SYSTEM_ERROR         | 20000  | System error                     |
| INVALID_HASH_ERROR   | 11055  | Invalid transaction hash         |

> 示例
```go
r := request.BIFTransactionCacheRequest{
    Hash: "",
}
res := ts.GetTxCacheData(r)
if res.ErrorCode != 0 {
    t.Error(res.ErrorDesc)
}
```

## 1.6 区块服务接口列表

​		区块服务接口主要是区块相关的接口，目前有6个接口：

| 序号 | 接口                | 说明                                   |
| ---- | ------------------- | -------------------------------------- |
| 1    | GetBlockNumber      | 该接口用于查询最新的区块高度           |
| 2    | GetTransactions     | 该接口用于查询指定区块高度下的所有交易 |
| 3    | GetBlockInfo        | 该接口用于获取区块信息                 |
| 4    | GetBlockLatestInfo  | 该接口用于获取最新区块信息             |
| 5    | GetValidators       | 该接口用于获取指定区块中所有验证节点数 |
| 6    | GetLatestValidators | 该接口用于获取最新区块中所有验证节点数 |

### 1.6.1 GetBlockNumber

> 接口说明

   	该接口用于查询最新的区块高度。

> 调用方法

```go
    GetBlockNumber() response.BIFBlockGetNumberResponse
```

> 响应数据

| 参数               | 类型        | 描述                            |
| ------------------ | ----------- | ------------------------------- |
| header             | BlockHeader | 区块头                          |
| header.blockNumber | int64        | 最新的区块高度，对应底层字段seq |

> 错误码

| 异常                 | 错误码 | 描述                             |
| -------------------- | ------ | -------------------------------- |
| CONNECTNETWORK_ERROR | 11007  | Failed to connect to the network |
| SYSTEM_ERROR         | 20000  | System error                     |

> 示例

```go
    bs := GetBlockInstance(SDK_INSTANCE_URL)
    res := bs.GetBlockNumber()
    if res.ErrorCode != 0 {
        t.Error(res.ErrorDesc)
    }
    
    fmt.Println("blockNumber:", res.Result.Header.BlockNumber)
```

### 1.6.2 GetTransactions

> 接口说明

   	该接口用于查询指定区块高度下的所有交易。

> 调用方法

```go
    GetTransactions(r request.BIFBlockGetTransactionsRequest) response.BIFBlockGetTransactionsResponse
```

> 请求参数

| 参数        | 类型 | 描述                                  |
| ----------- | ---- | ------------------------------------- |
| blockNumber | int64 | 必填，最新的区块高度，对应底层字段seq |

> 响应数据

| 参数         | 类型                    | 描述           |
| ------------ | ----------------------- | -------------- |
| totalCount   | int64                    | 返回的总交易数 |
| transactions | []BIFTransactionHistory | 交易内容       |

> 错误码

| 异常                      | 错误码 | 描述                             |
| ------------------------- | ------ | -------------------------------- |
| INVALID_BLOCKNUMBER_ERROR | 11060  | BlockNumber must bigger than 0   |
| REQUEST_NULL_ERROR        | 12001  | Request parameter cannot be null |
| CONNECTNETWORK_ERROR      | 11007  | Failed to connect to the network |
| SYSTEM_ERROR              | 20000  | System error                     |

> 示例

```go
    bs := GetBlockInstance(SDK_INSTANCE_URL)
    var r request.BIFBlockGetTransactionsRequest
    r.BlockNumber = 617247
    res := bs.GetTransactions(r)
    if res.ErrorCode != 0 {
        t.Error(res.ErrorDesc)
    }
    
    fmt.Printf("result: %+v \n", res.Result)
```

### 1.6.3 GetBlockInfo

> 接口说明

   	该接口用于获取指定区块信息。

> 调用方法

```go
    GetBlockInfo(r request.BIFBlockGetInfoRequest) response.BIFBlockGetInfoResponse
```

> 请求参数

| 参数        | 类型 | 描述                   |
| ----------- | ---- | ---------------------- |
| blockNumber | int64 | 必填，待查询的区块高度 |

> 响应数据

| 参数               | 类型           | 描述         |
| ------------------ | -------------- | ------------ |
| header             | BIFBlockHeader | 区块信息     |
| header.confirmTime | int64           | 区块确认时间 |
| header.number      | int64           | 区块高度     |
| header.txCount     | int64           | 交易总量     |
| header.version     | string         | 区块版本     |

> 错误码

| 异常                      | 错误码 | 描述                             |
| ------------------------- | ------ | -------------------------------- |
| INVALID_BLOCKNUMBER_ERROR | 11060  | BlockNumber must bigger than 0   |
| REQUEST_NULL_ERROR        | 12001  | Request parameter cannot be null |
| CONNECTNETWORK_ERROR      | 11007  | Failed to connect to the network |
| SYSTEM_ERROR              | 20000  | System error                     |

> 示例

```go
    bs := GetBlockInstance(SDK_INSTANCE_URL)
    var r request.BIFBlockGetInfoRequest
    r.BlockNumber = 617247
    res := bs.GetBlockInfo(r)
    if res.ErrorCode != 0 {
        t.Error(res.ErrorDesc)
    }
    
    fmt.Printf("result: %+v \n", res.Result)
```

### 1.6.4 GetBlockLatestInfo

> 接口说明

```
该接口用于获取最新区块信息。
```

> 调用方法

```go
    GetBlockLatestInfo() response.BIFBlockGetLatestInfoResponse
```

> 响应数据

| 参数               | 类型           | 描述                      |
| ------------------ | -------------- | ------------------------- |
| header             | BIFBlockHeader | 区块信息                  |
| header.confirmTime | int64           | 区块确认时间              |
| header.number      | int64           | 区块高度，对应底层字段seq |
| header.txCount     | int64           | 交易总量                  |
| header.version     | string         | 区块版本                  |


> 错误码

| 异常                 | 错误码 | 描述                             |
| -------------------- | ------ | -------------------------------- |
| CONNECTNETWORK_ERROR | 11007  | Failed to connect to the network |
| SYSTEM_ERROR         | 20000  | System error                     |

> 示例

```go
    bs := GetBlockInstance(SDK_INSTANCE_URL)
    res := bs.GetBlockLatestInfo()
    if res.ErrorCode != 0 {
        t.Error(res.ErrorDesc)
    }
    
    fmt.Printf("result: %+v \n", res.Result)
```

### 1.6.5 GetValidators

> 接口说明

   	该接口用于获取指定区块中所有验证节点数。

> 调用方法

```go
    GetValidators(r request.BIFBlockGetValidatorsRequest) response.BIFBlockGetValidatorsResponse
```

> 请求参数

| 参数        | 类型 | 描述                              |
| ----------- | ---- | --------------------------------- |
| blockNumber | int64 | 必填，待查询的区块高度，必须大于0 |

> 响应数据

| 参数               | 类型     | 描述         |
| ------------------ | -------- | ------------ |
| validators         | []string | 验证节点列表 |
| validators.address | string   | 共识节点地址 |

> 错误码

| 异常                      | 错误码 | 描述                             |
| ------------------------- | ------ | -------------------------------- |
| INVALID_BLOCKNUMBER_ERROR | 11060  | BlockNumber must bigger than 0   |
| REQUEST_NULL_ERROR        | 12001  | Request parameter cannot be null |
| CONNECTNETWORK_ERROR      | 11007  | Failed to connect to the network |
| SYSTEM_ERROR              | 20000  | System error                     |

> 示例

```go
    bs := GetBlockInstance(SDK_INSTANCE_URL)
    var r request.BIFBlockGetValidatorsRequest
    r.BlockNumber = 617247
    res := bs.GetValidators(r)
    if res.ErrorCode != 0 {
        t.Error(res.ErrorDesc)
    }

    fmt.Printf("result: %+v \n", res.Result)
```

### 1.6.6 GetLatestValidators

> 接口说明

   	该接口用于获取最新区块中所有验证节点数。

> 调用方法

```go
    GetLatestValidators() response.BIFBlockGetLatestValidatorsResponse
```

> 响应数据

| 参数               | 类型     | 描述         |
| ------------------ | -------- | ------------ |
| validators         | []string | 验证节点列表 |
| validators.address | string   | 共识节点地址 |

> 错误码

| 异常                 | 错误码 | 描述                             |
| -------------------- | ------ | -------------------------------- |
| CONNECTNETWORK_ERROR | 11007  | Failed to connect to the network |
| SYSTEM_ERROR         | 20000  | System error                     |

> 示例

```go
    bs := GetBlockInstance(SDK_INSTANCE_URL)
    res := bs.GetLatestValidators()
    if res.ErrorCode != 0 {
        t.Error(res.ErrorDesc)
    }
    
    fmt.Printf("result: %+v \n", res.Result)
```