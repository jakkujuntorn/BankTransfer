package entries

import (
	"banktransfer/models"
	"gorm.io/gorm"
)

// เช็คพวก statement น่าจะอยู่ตรงนี้
// ดึงค่าตามช่วงเวลา ของ account
type repo_entries struct {
	db *gorm.DB
}

type I_Repo_Entries interface {
	GetStaement(accountId int, startTime, endTime string) ([]models.Entry, error)
}

func New_Repo_Entries(db *gorm.DB) I_Repo_Entries {
	return &repo_entries{
		db: db,
	}
}

// GetStaement implements I_Repo_Entries
func (re *repo_entries) GetStaement(accountId int, startTime string, endTime string) (dataStatement []models.Entry, err error) {
	startTime = `2023-11-01`
	endTime = `2023-12-30`

	queryGetStatement := `select * from entries where account_id = ? and created_at between  ? and  ? order by created_at desc`

	tx := re.db.Raw(queryGetStatement,accountId,startTime,endTime).Scan(&dataStatement)
	if tx.Error != nil {
		return dataStatement, tx.Error
	}

	return dataStatement, nil
}
