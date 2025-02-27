package domain

import (
	"time"
	"xyz-finance-api/internal/transaction/entity"
)

type Transaction struct {
	ID          string
	LoanID      string
	AssetName   string
	TotalAmount int
	Tenor       int
	OTRPrice    int
	AdminFee    int
	Interest    int
	Status      string
	CreatedAt   time.Time
}

// mapper
func TransactionDomainToTransactionEntity(transactionDomain Transaction) entity.Transaction {
	return entity.Transaction{
		ID:          transactionDomain.ID,
		LoanID:      transactionDomain.LoanID,
		AssetName:   transactionDomain.AssetName,
		TotalAmount: transactionDomain.TotalAmount,
		Tenor:       transactionDomain.Tenor,
		OTRPrice:    transactionDomain.OTRPrice,
		AdminFee:    transactionDomain.AdminFee,
		Interest:    transactionDomain.Interest,
		Status:      transactionDomain.Status,
		CreatedAt:   transactionDomain.CreatedAt,
	}
}

func TransactionEntityToTransactionDomain(transactionEntity entity.Transaction) Transaction {
	return Transaction{
		ID:          transactionEntity.ID,
		LoanID:      transactionEntity.LoanID,
		AssetName:   transactionEntity.AssetName,
		TotalAmount: transactionEntity.TotalAmount,
		Tenor:       transactionEntity.Tenor,
		OTRPrice:    transactionEntity.OTRPrice,
		AdminFee:    transactionEntity.AdminFee,
		Interest:    transactionEntity.Interest,
		Status:      transactionEntity.Status,
		CreatedAt:   transactionEntity.CreatedAt,
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
