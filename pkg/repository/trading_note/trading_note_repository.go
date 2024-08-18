package trading_note_repository

import (
	"os"
	"pdf_balance_parser/pkg/model/trading_note"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type TradingNoteRepository struct {
	connection gorm.DB
}

func (r *TradingNoteRepository) New() TradingNoteRepository {
	var err error
	conn, err := gorm.Open(sqlite.Open(os.Getenv("DB_DATABASE")), &gorm.Config{})
	r.connection = *conn
	if err != nil {
		panic("failed to connect on database")
	}
	return *r
}

func (repository TradingNoteRepository) Store(trading_note_summary *trading_note.TradingNoteSummaries) {
	repository.connection.Create(trading_note_summary)
}