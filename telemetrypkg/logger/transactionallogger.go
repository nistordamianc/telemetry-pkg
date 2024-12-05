package logger

import "github.com/google/uuid"

type TransactionalLogger struct {
	TransactionID string
	Logger
}

func NewTransactionalLogger(logger Logger) *TransactionalLogger {
	return &TransactionalLogger{
		TransactionID: uuid.NewString(),
		Logger:        logger,
	}
}

func (t *TransactionalLogger) log(log Log) {
	transactionMetadata := Metadata{
		Key:   "transaction_id",
		Value: t.TransactionID,
	}
	log.Metadata = append(log.Metadata, transactionMetadata)

	t.Logger.log(log)
}

func (t *TransactionalLogger) LogW(message string, metadata []Metadata) {
	transactionMetadata := Metadata{
		Key:   "transaction_id",
		Value: t.TransactionID,
	}
	metadata = append(metadata, transactionMetadata)

	t.Logger.LogW(message, metadata)
}

func (t *TransactionalLogger) LogE(message string, metadata []Metadata) {
	transactionMetadata := Metadata{
		Key:   "transaction_id",
		Value: t.TransactionID,
	}
	metadata = append(metadata, transactionMetadata)

	t.Logger.LogE(message, metadata)
}

func (t *TransactionalLogger) LogI(message string, metadata []Metadata) {
	transactionMetadata := Metadata{
		Key:   "transaction_id",
		Value: t.TransactionID,
	}
	metadata = append(metadata, transactionMetadata)

	t.Logger.LogI(message, metadata)
}
