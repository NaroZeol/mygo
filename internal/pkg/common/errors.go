package common

import "errors"

var (
	ErrorEmpty              = errors.New("empty parameter")
	ErrorUnknownUsername    = errors.New("unknown username")
	ErrorUnknownUserId      = errors.New("unknown user id")
	ErrorWrongPassword      = errors.New("wrong password")
	ErrorRepeatUsername     = errors.New("repeat username")
	ErrorUnknownRole        = errors.New("unknown role, only has 'Old', 'Volunteer', 'Admin'")
	ErrorUnknownTransaction = errors.New("unknown transaction")
	ErrorMatchTransaction   = errors.New("user id do not match the transaction's user id")
	ErrorInvalidParam       = errors.New("invalid parameter")
	ErrorNoPermission       = errors.New("no permission")
	// blockchain error
	ErrorNoWallet         = errors.New("user do not have wallet")
	ErrorBalanceNotEnough = errors.New("balance is not enough")
	ErrorInvalidAmount    = errors.New("invalid amount")
	// token error
	ErrorInvalidToken     = errors.New("invalid token")
	ErrorExpiredToken     = errors.New("expired token")
	ErrorGetInfoFromToken = errors.New("failed to get info from token")
	// internal error
	ErrorOperateDatabase      = errors.New("failed to operate database")
	ErrorBlockchainDisconnect = errors.New("blockchain is disconnected")
	ErrorTokenContract        = errors.New("failed to get token contract")
	// transfer error
	ErrorSameID              = errors.New("sender id and receiver id are the same")
	ErrorNegativeAmount      = errors.New("amount is negative")
	ErrorInsufficientBalance = errors.New("insufficient balance")
	ErrorWrongPassphrase     = errors.New("wrong passphrase")
)

var internalErrors = []error{
	ErrorOperateDatabase,
	ErrorBlockchainDisconnect,
	ErrorTokenContract,
}

func CheckInternalError(err error) bool {
	if err == nil {
		return false
	}

	for _, e := range internalErrors {
		if err == e {
			return true
		}
	}

	return false
}
