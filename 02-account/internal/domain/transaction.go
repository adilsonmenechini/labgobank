package domain

import "errors"

type TransactionType string
type Success int

const (
	Deposit  TransactionType = "Deposit"
	Transfer TransactionType = "Transfer"
	Withdraw TransactionType = "Withdraw"
	Payment  TransactionType = "Payment"
	Refund   TransactionType = "Refund"
	Reversal TransactionType = "Reversal"
)

// Custom error types
var (
	ErrInsufficientFunds      = errors.New("insufficient funds")
	ErrDepositLimitExceeded   = errors.New("deposit limit exceeded")
	ErrCreditLimitExceeded    = errors.New("credit card limit exceeded")
	ErrCreditCardExpired      = errors.New("credit card expired")
	ErrDebitInsufficient      = errors.New("debit failed - insufficient funds")
	ErrPaymentInsufficient    = errors.New("insufficient funds for payment")
	ErrPaymentLimitExceeded   = errors.New("payment limit exceeded")
	ErrRefundInsufficient     = errors.New("refund failed - insufficient funds")
	ErrWithdrawalInsufficient = errors.New("withdrawal failed - insufficient funds")
	ErrTransferInsufficient   = errors.New("transfer failed - insufficient funds")
	ErrTransactionNotFound    = errors.New("transaction not found")
	ErrInvalidCVV             = errors.New("invalid CCV")
)
