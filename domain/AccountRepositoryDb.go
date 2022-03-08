package domain

import (
	"banking/errs"
	"banking/logger"
	"github.com/jmoiron/sqlx"
	"strconv"
)

type AccountRepositoryDb struct {
	client *sqlx.DB
}

func (d AccountRepositoryDb) Save(account Account) (*Account, *errs.ApplicationError) {
	sqlInsert := "INSERT INTO ACCOUNTS (customer_id, opening_date, account_type, amount, status) VALUES (?, ?, ?, ?, ?)"

	result, err := d.client.Exec(sqlInsert, account.CustomerId, account.OpeningDate, account.AccountType, account.Amount, account.Status)
	if err != nil {
		logger.Error("error while creating new account: " + err.Error())
		return nil, errs.NewUnexpectedError("unexpected error from database")
	}

	id, err := result.LastInsertId()
	if err != nil {
		logger.Error("error while getting last insert id for new account: " + err.Error())
		return nil, errs.NewUnexpectedError("unexpected error from database")
	}
	account.AccountId = strconv.FormatInt(id, 10)
	return &account, nil
}

func (d AccountRepositoryDb) FindById(accountId string) (*Account, *errs.ApplicationError) {
	sqlGetAccount := "SELECT account_id, customer_id, opening_date, account_type, amount FROM accounts WHERE account_id = ?"

	var account Account
	err := d.client.Get(&account, sqlGetAccount, accountId)
	if err != nil {
		logger.Error("error while fetching account information: " + err.Error())
		return nil, errs.NewUnexpectedError("unexpected database error")
	}
	return &account, nil
}

func (d AccountRepositoryDb) SaveTransaction(transaction Transaction) (*Transaction, *errs.ApplicationError) {
	tx, err := d.client.Beginx()
	if err != nil {
		logger.Error("error while starting a new transaction for bank account transaction: " + err.Error())
		return nil, errs.NewUnexpectedError("unexpected database error")
	}

	result, _ := tx.Exec(`INSERT INTO transactions(account_id, amount, transaction_type, transaction_date) VALUES (?, ?, ?, ?)`,
		transaction.AccountId, transaction.Amount, transaction.TransactionType, transaction.TransactionDate)

	if transaction.IsWithdrawal() {
		_, err = tx.Exec(`UPDATE accounts SET amount = amount - ? WHERE account_id = ?`,
			transaction.Amount, transaction.AccountId)
	} else {
		_, err = tx.Exec(`UPDATE accounts SET amount = amount + ? WHERE account_id = ?`,
			transaction.Amount, transaction.AccountId)
	}

	if err != nil {
		tx.Rollback()
		logger.Error("error while saving transaction: " + err.Error())
		return nil, errs.NewUnexpectedError("unexpected database error")
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		logger.Error("error while committing transaction for bank account: " + err.Error())
		return nil, errs.NewUnexpectedError("unexpected database error")
	}

	transactionId, err := result.LastInsertId()
	if err != nil {
		logger.Error("error while getting the last transaction id: " + err.Error())
		return nil, errs.NewUnexpectedError("unexpected database error")
	}

	account, appErr := d.FindById(transaction.AccountId)
	if appErr != nil {
		return nil, appErr
	}
	transaction.TransactionId = strconv.FormatInt(transactionId, 10)

	transaction.Amount = account.Amount
	return &transaction, nil
}

func NewAccountRepositoryDb(dbClient *sqlx.DB) AccountRepositoryDb {
	return AccountRepositoryDb{dbClient}
}
