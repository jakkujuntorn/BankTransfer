package models

import (
	"time"
)

// Repo ******************************

type Create_Deposit_and_Withdraw struct {
	Owner     string `json:"owner"`
	AccountID int64  `json:"from_account_id"`
	Amount    int64  `json:"amount"`
}

type Currency struct {
	Currency string `json:"currency"`
}

type CreateTransferParams struct {
	Owner         string `json:"owner"`
	FromAccountID int64  `json:"from_account_id"`
	ToAccountID   int64  `json:"to_account_id"`
	Amount        int64  `json:"amount"`
	CreatedAt time.Time `json:"created_at"`
}

type ListTransfersParams struct {
	FromAccountID int64 `json:"from_account_id"`
	ToAccountID   int64 `json:"to_account_id"`
	Limit         int32 `json:"limit"`
	Offset        int32 `json:"offset"`
}

type ListTransfers_ForOwner struct {
	Owner         string `json:"owner"`
	FromAccountID int64  `json:"from_account_id"`
	ToAccountID   int64  `json:"to_account_id"`
	Amount        int64  `json:"amount"`
	CreatedAt time.Time `json:"created_at"`
	Limit         int32 `json:"limit"`
	Offset        int32 `json:"offset"`
}

// *****************************************

type Transfer struct {
	ID            int64 `json:"id"`
	FromAccountID int64 `json:"from_account_id"`
	ToAccountID   int64 `json:"to_account_id"`
	// must be positive
	Amount    int64     `json:"amount"`
	CreatedAt time.Time `json:"created_at"`
}

type TransferTxParams struct {
	FromAccountID int64 `json:"from_account_id"`
	ToAccountID   int64 `json:"to_account_id"`
	Amount        int64 `json:"amount"`
}

// TransferTxResult is the result of the transfer transaction
type TransferTxResult struct {
	Transfer    Transfer `json:"transfer"`
	FromAccount Account  `json:"from_account"`
	ToAccount   Account  `json:"to_account"`
	FromEntry   Entry    `json:"from_entry"`
	ToEntry     Entry    `json:"to_entry"`
}

type GetTransfer_ById struct{
AccountID int64
// StartTime 
// EndTime
}

type GetTransfer_ByOwner struct{
Owner string
// StartTime 
// EndTime
}
