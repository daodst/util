package trerr

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

var Language = "EN"

type TrErr struct {
	RawMsg string
	TrMsg  string
	ErrMsg string
}

func (tr *TrErr) Error() string {
	return tr.ErrMsg
}

func (tr *TrErr) DumpError(errMsg string) {
	if errMsg != "" {
		r, _ := regexp.Compile(tr.RawMsg)
		params := r.FindStringSubmatch(errMsg)
		if len(params) <= 1 || !strings.Contains(tr.TrMsg, "%s") {
			return
		} else {
			tmp := make([]interface{}, 0)
			for _, param := range params[1:] {
				tmp = append(tmp, param)
			}
			tr.ErrMsg = fmt.Sprintf(tr.TrMsg, tmp...)
			return
		}
	}
	return
}

var trMap = make(map[string]TrErr)

func NewErr(unTrMsg, trMsg string) TrErr {
	te := TrErr{
		RawMsg: unTrMsg,
		TrMsg:  trMsg,
		ErrMsg: trMsg,
	}
	trMap[unTrMsg] = te
	return te
}

func TransError(unTrMsg string) error {
	for regC, te := range trMap {
		r, _ := regexp.Compile(regC)
		if r.Match([]byte(unTrMsg)) && Language == "CHC" {
			te.DumpError(unTrMsg)
			e := te
			return &e
		}
	}
	return errors.New(unTrMsg)
}

var (
	InvalidFromAddress = NewErr("^invalid from address: (.*)$", "發起帳號不合法")
	ParseError         = NewErr("failed to parse request", "格式化請求數據失敗")
	ChainIdError       = NewErr("chain-id required but not specified", "鏈編號必填")
	FeeGasError        = NewErr("cannot provide both fees and gas prices", "手續費和價格不能同時為零")
	FeeGasInvalid      = NewErr("invalid fees or gas prices provided", "手續費或價格不能同時為空")
	InvalidAddress     = NewErr("invalid address", "帳號地址不合法")
	InsufficientError  = NewErr("insufficient funds", "餘額不足")
	MarshalError       = NewErr("failed to marshal JSON bytes", "序列化失敗")
	UnMarshalError     = NewErr("failed to unmarshal JSON bytes", "序列化失敗")
	GasAdjustmentError = NewErr("invalid gas adjustment", "價格調整不合法")
	AccountNoExist     = NewErr("account not exist", "帳號不存在")
	AccountError       = NewErr("decoding Bech32 address failed: must provide an address", "解析帳號地址失敗")
	CoinError          = NewErr("coin can not be empty", "幣不能為空")
	BalanceError       = NewErr("Insufficient account balance", "帳號餘額不足")
	AccountEmpty       = NewErr("account can not be empty", "帳號不能為空")
	PageInvalid        = NewErr("page is invalid", "頁碼格式錯誤")
	PageSizeINvalid    = NewErr("pageSize is invalid", "頁面數據格式錯誤")
	TokenEmpty         = NewErr("token is empty", "幣種為空")
	QueryParamError    = NewErr("query param empty", "入參為空")
	NumError           = NewErr("query param num invalid", "數量參數為空")
	QueryExtError      = NewErr("query extype invalid", "")
	IdsError           = NewErr("query param idsString is empty", "")
	TxhashEmpty        = NewErr("query param txhash empty", "交易哈希為空")
	TxhashNotExist     = NewErr("txhash not exist", "交易哈希不存在")
	DeleAddressError   = NewErr("must use own delegator address", "POS抵押帳號不合法")
	QueryDeletorError  = NewErr("query delegators amount error", "査詢POS抵押餘額錯誤")
	UnbondInfuffiError = NewErr("unbond amount is not enough", "解綁餘額不足")
	DeleAddressEmpty   = NewErr("delegatorAddr can not be empty", "POS抵押帳號不能為空")
	_                  = NewErr("format gas error", "價格格式化失敗")
	_                  = NewErr("invalid gas adjustment", "價格調整參數不合法")
	_                  = NewErr("query data error", "査詢數據失敗")
	_                  = NewErr("current account exist block chain request ,please wait a minute", "當前帳號已存在上鏈請求，請稍後再試")
	_                  = NewErr("sign error", "錢包密碼驗證錯誤")
	_                  = NewErr("get accountManage error", "獲取帳號管理器失敗")
	_                  = NewErr("Entropy length must be \\[128, 256\\] and a multiple of 32", "生成助記詞位數錯誤")
	_                  = NewErr("account only exist", "帳號已存在")
	_                  = NewErr("encoding bech32 failed", "編碼類型失敗")
	_                  = NewErr("password verification failed", "帳號密碼錯誤")
	_                  = NewErr("account not exist", "帳號不存在")
	_                  = NewErr("account key not exist", "帳號秘鑰未找到")

	_ = NewErr("failed to decrypt private key", "私鑰解密失敗")
	_ = NewErr("invalid mnemonic", "助記詞格式錯誤")
	_ = NewErr("height must be equal or greater than zero", "塊高度不能小於零")
	_ = NewErr("empty delegator address", "委託帳號為空")
	_ = NewErr("empty validator address", "POS礦工器地址為空")
	_ = NewErr("invalid delegation amount", "委託金額不合法")
	_ = NewErr("invalid shares amount", "解綁金額不合法")
	_ = NewErr("validator does not exist", "POS礦工地址不存在")
	_ = NewErr("invalid coin denomination", "委託幣種錯誤")
	_ = NewErr("delegate progress error", "股權質押失敗")
	_ = NewErr("no validator distribution info", "沒有驗證程式分發資訊")
	_ = NewErr("no delegation distribution info", "無委派分發資訊")
	_ = NewErr("module account (.*)$ does not exist", "模塊帳號不存在")
	_ = NewErr("no validator commission to withdraw", "無驗證器傭金可選取")
	_ = NewErr("signature verification failed; verify correct account sequence and chain-id", "簽名驗證失敗,請檢查帳號序號和鏈ID")
	_ = NewErr("database error", "資料庫錯誤")
	_ = NewErr("parse account error", "帳號格式化錯誤")
	_ = NewErr("parse coin error", "價格錯誤")
	_ = NewErr("parse string to number error", "類型轉化失敗")
	_ = NewErr("query chain infor errors", "査詢鏈上數據失敗")
	_ = NewErr("valid chain request error", "請求參數驗證失敗")
	_ = NewErr("parse json error", "格式化失敗")
	_ = NewErr("parse byte to struct error", "類型轉化錯誤")
	_ = NewErr("format tx struct error", "交易數轉化失敗")
	_ = NewErr("broadcast error", "廣播失敗")
	_ = NewErr("format string to int error", "類型轉化失敗")
	_ = NewErr("parse valitor error", "驗證地址格式化失敗")
	_ = NewErr("parse time error", "時間格式化失敗")
	_ = NewErr("sign error", "錢包密碼驗證錯誤")
	_ = NewErr("operator address exist", "當前錢包地址已經申請過POS礦工了")
	_ = NewErr("pub key for validator exist", "當前公鑰已生成驗證器")
	_ = NewErr("current validator not jail", "POS礦工非監禁狀態")
	_ = NewErr("has no right to oprate validator", "無權解禁")
	_ = NewErr("validator description length error", "礦工資訊過長")
	_ = NewErr("current account does not have right to delete datahash ", "當前帳號無權删除")
	_ = NewErr("too many unbonding delegation entries for \\(delegator, validator\\) tuple", "當前申請OS贖回的記錄已達到上限,需等待到賬後再操作.")
	_ = NewErr("no delegation for \\(address, validator\\) tuple", "錢包地址下沒有POS抵押金額")
	_ = NewErr("parse validator address error", "解析POS礦工地址出錯")
	_ = NewErr("There is no reward to receive", "沒有獎勵可以領取")
	_ = NewErr("delegation does not exist", "沒有抵押數據")
	_ = NewErr("query sensitive words error", "査詢敏感詞錯誤")
	_ = NewErr("save sensitive words error", "保存敏感詞錯誤")
	_ = NewErr("sensitive status illegal", "敏感詞狀態不合法")
	_ = NewErr("contain sensitive words", "包含敏感詞")
	_ = NewErr("fee can not be zero", "手續費不能為零")
	_ = NewErr("fee is too less", "手續費太小")
	_ = NewErr("fee can not empty", "手續費不能為空")
	_ = NewErr("delegation coin less then min", "抵押金額太低")
	_ = NewErr("unbonding delegation shares less then min", "贖回投票權太低")
	_ = NewErr("delegation reward coin less then min", "POS獎勵金額單項太低")
	_ = NewErr("not enough delegation shares", "投票權不足")
	_ = NewErr("must use own validator address", "必須使用驗證器所屬的地址發起請求")
	_ = NewErr("private key is empty", "私鑰為空")
	_ = NewErr("^verification fail$", "驗簽失敗")
	_ = NewErr("pubkey not exist", "公鑰不存在")
	_ = NewErr("verification error", "驗簽失敗")
	_ = NewErr("con address is invalid", "共識地址不合法")
	_ = NewErr("con address can not empty", "共識地址不能為空")
	_ = NewErr("dir name has exist", "目錄已存在")
	_ = NewErr("delegator does not contain delegation", "未找到POS抵押資訊")
	_ = NewErr("The account to be unlocked must have a valid POS mortgage", "申請解禁的帳號必須存在有效的POS抵押")
	_ = NewErr("The mortgage amount is less than the self-mortgage amount", "申請解禁的帳號抵押值小於最小自抵押值")
	_ = NewErr("dir type illegal", "目錄類型不合法")
	_ = NewErr("dir not exist", "目錄不存在")
	_ = NewErr("lower min coin", "超出最小精度")
	_ = NewErr("SendCoins error", "轉帳失敗")
	_ = NewErr("signature verification failed, invalid chainid or account number", "驗簽失敗,無效的chainId或者帳號number")
	_ = NewErr("account serial number expired, the reason may be: node block behind or repeatedly sent messages", "帳號序號過期,原因可能是:節點區塊落後或者重複發送了消息.")
	_ = NewErr("Verifier information can only be changed once in 24 hours", "驗證器資訊24小時內僅可更改一次")
	_ = NewErr("Please use the main currency", "請使用主幣")
	_ = NewErr("chain error", "上鏈失敗")
	_ = NewErr("min self delegation cannot be zero", "最小自委託不能為零")

	_ = NewErr("commission must be positive", "傭金必須是整數")
	_ = NewErr("commission cannot be more than 100%", "傭金不能超過100%")
	_ = NewErr("commission cannot be more than the max rate", "傭金不能超過最高費率")
	_ = NewErr("commission cannot be changed more than once in 24h", "傭金修改24小時只能執行一次")
	_ = NewErr("commission change rate must be positive", "更改的傭金必須是整數")
	_ = NewErr("commission change rate cannot be more than the max rate", "更改的傭金不能超過最高費率")
	_ = NewErr("commission cannot be changed more than max change rate", "傭金變化不能超過最大變化率")
	_ = NewErr("validator's self delegation must be greater than their minimum self delegation", "DPOS礦工的抵押金額必須大於最小自抵押")
	_ = NewErr("minimum self delegation must be a positive integer", "最小自抵押必須是整數")
	_ = NewErr("minimum self delegation cannot be decrease", "不能减少最小自抵押")
	_ = NewErr("invalid amount", "無效金額")
	_ = NewErr("is not allowed to receive funds", "不允許的收款地址")
	_ = NewErr("There can only be one handling fee currency", "手續費只能有一項")
	_ = NewErr("fee amount is invalid", "手續費校驗失敗")
	_ = NewErr("Unauthorized account number", "未授權的帳號")
)
