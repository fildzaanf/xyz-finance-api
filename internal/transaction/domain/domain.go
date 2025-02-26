package domain

import (
	"time"
	// el "xyz-finance-api/internal/loan/entity"
	// dl "xyz-finance-api/internal/loan/domain"
	// ep "xyz-finance-api/internal/payment/entity"
	"xyz-finance-api/internal/transaction/entity"
)

type Transaction struct {
	ID                string
	LoanID            string
	AssetName         string
	TotalAmount       int
	OTRPrice          int
	AdminFee          int
	Interest          int
	InstallmentAmount int
	Status            string
	CreatedAt         time.Time
	// Loan              el.Loan
	// Payments          []ep.Payment `gorm:"foreignKey:TransactionID"`
}

// mapper
func TransactionDomainToTransactionEntity(transactionDomain Transaction) entity.Transaction {
	return entity.Transaction{
		ID:                transactionDomain.ID,
		LoanID:            transactionDomain.LoanID,
		AssetName:         transactionDomain.AssetName,
		TotalAmount:       transactionDomain.TotalAmount,
		OTRPrice:          transactionDomain.OTRPrice,
		AdminFee:          transactionDomain.AdminFee,
		Interest:          transactionDomain.Interest,
		InstallmentAmount: transactionDomain.InstallmentAmount,
		Status:            transactionDomain.Status,
		CreatedAt:         transactionDomain.CreatedAt,
		// Loan:              dl.LoanDomainToLoanEntity(),
		// Payments:          transactionDomain.Payments,
	}
}

func TransactionEntityToTransactionDomain(transactionEntity entity.Transaction) Transaction {
	return Transaction{
		ID:                transactionEntity.ID,
		LoanID:            transactionEntity.LoanID,
		AssetName:         transactionEntity.AssetName,
		TotalAmount:       transactionEntity.TotalAmount,
		OTRPrice:          transactionEntity.OTRPrice,
		AdminFee:          transactionEntity.AdminFee,
		Interest:          transactionEntity.Interest,
		InstallmentAmount: transactionEntity.InstallmentAmount,
		Status:            transactionEntity.Status,
		CreatedAt:         transactionEntity.CreatedAt,
		// Loan:              transactionEntity.Loan,
		// Payments:          transactionEntity.Payments,
	}
}

func ListTransactionDomainToTransactionEntity(transactionDomains []Transaction) []entity.Transaction {
	listTransactionEntities := []entity.Transaction{}
	for _, transaction := range transactionDomains {
		transactionEntity := TransactionDomainToTransactionEntity(transaction)
		listTransactionEntities = append(listTransactionEntities, transactionEntity)
	}
	return listTransactionEntities
}

func ListTransactionEntityToTransactionDomain(transactionEntities []entity.Transaction) []Transaction {
	listTransactionDomains := []Transaction{}
	for _, transaction := range transactionEntities {
		transactionDomain := TransactionEntityToTransactionDomain(transaction)
		listTransactionDomains = append(listTransactionDomains, transactionDomain)
	}
	return listTransactionDomains
}
