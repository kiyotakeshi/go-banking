package service

import (
	"banking/domain"
	"banking/dto"
	"banking/errs"
	"time"
)

type AccountService interface {
	NewAccount(dto.NewAccountRequest) (*dto.NewAccountResponse, *errs.ApplicationError)
	MakeTransaction(request dto.TransactionRequest) (*dto.TransactionResponse, *errs.ApplicationError)
}

type DefaultAccountService struct {
	repository domain.AccountRepository
}

func (s DefaultAccountService) NewAccount(request dto.NewAccountRequest) (*dto.NewAccountResponse, *errs.ApplicationError) {
	err := request.Validate()
	if err != nil {
		return nil, err
	}

	account := domain.Account{
		AccountId:   "",
		CustomerId:  request.CustomerId,
		OpeningDate: time.Now().Format("2006-01-02 15:04:05"),
		AccountType: request.AccountType,
		Amount:      request.Amount,
		Status:      "1",
	}

	newAccount, err := s.repository.Save(account)
	if err != nil {
		return nil, err
	}

	response := newAccount.ToNewAccountResponseDto()

	return &response, nil
}

func (s DefaultAccountService) MakeTransaction(request dto.TransactionRequest) (*dto.TransactionResponse, *errs.ApplicationError) {
	err := request.Validate()
	if err != nil {
		return nil, err
	}

	if request.IsTransactionTypeWithdrawal() {
		account, err := s.repository.FindById(request.AccountId)
		if err != nil {
			return nil, err
		}

		if !account.CanWithdraw(request.Amount) {
			return nil, errs.NewValidationError("insufficient balance in the account")
		}
	}
	transaction := domain.Transaction{
		AccountId:       request.AccountId,
		Amount:          request.Amount,
		TransactionType: request.TransactionType,
		TransactionDate: time.Now().Format("2006-01-02 15:04:05"),
	}
	savedTransaction, appErr := s.repository.SaveTransaction(transaction)
	if appErr != nil {
		return nil, appErr
	}
	response := savedTransaction.ToDto()
	return &response, nil
}

func NewAccountService(repo domain.AccountRepository) DefaultAccountService {
	return DefaultAccountService{repo}
}
