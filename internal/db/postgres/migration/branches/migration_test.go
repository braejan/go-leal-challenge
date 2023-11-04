package branches_test

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/braejan/go-leal-challenge/internal/db/postgres/migration/branches"
	branchModel "github.com/braejan/go-leal-challenge/internal/domain/branches/model"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Test_BranchMigrationData_Err_Open(t *testing.T) {

	mockDb, mock, _ := sqlmock.New()
	defer mockDb.Close()
	dialector := postgres.New(postgres.Config{
		Conn:       mockDb,
		DriverName: "postgres",
	})
	db, _ := gorm.Open(dialector, &gorm.Config{PrepareStmt: true})
	result := sqlmock.NewResult(0, 1)
	mock.ExpectPrepare("CREATE TABLE \"branches\" (.+)").ExpectExec().WillReturnResult(result).WillReturnError(nil)
	mock.ExpectPrepare("CREATE INDEX IF NOT EXISTS .+").ExpectExec().
		WillReturnResult(sqlmock.NewResult(0, 0))
	err := db.AutoMigrate(
		&branchModel.Branch{},
	)
	if err != nil {
		t.Fatal(err)
	}
	mock.ExpectQuery("SELECT (.+) FROM branches").WillReturnError(gorm.ErrRecordNotFound)
	err = branches.BranchMigrationData(db)
	assert.NotNil(t, err)
}
