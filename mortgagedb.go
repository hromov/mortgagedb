package mortgagedb

import (
	"errors"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var banks_data = []Bank{
	{Name: "Bank of America", Interest: 0.2, MaxLoan: 40000, MinDown: 0.3, Term: 12},
	{Name: "Bank of China", Interest: 0.2, MaxLoan: 2000000, MinDown: 0.25, Term: 14},
	{Name: "Bank of Ukraine", Interest: 0.2, MaxLoan: 20000, MinDown: 0.2, Term: 9},
	{Name: "Bank of Spain", Interest: 0.2, MaxLoan: 100000, MinDown: 0.4, Term: 16},
	{Name: "Bank of Italy", Interest: 0.2, MaxLoan: 300000, MinDown: 0.5, Term: 32},
}

type Bank struct {
	gorm.Model
	// ID       string  `json:"id"`
	Name     string  `json:"name"`
	Interest float32 `json:"interest"`
	MaxLoan  int     `json:"max_loan"`
	MinDown  float32 `json:"min_down"`
	Term     int32   `json:"term"`
}

var db *gorm.DB

func Init(dsn string) (err error) {
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return errors.New(fmt.Sprintf("failed to connect database error: %s", err.Error()))
	}

	// if table exist - do nothink, if not - create init structure with test data
	if db.Migrator().HasTable("banks") {
		return nil
	}

	if err := db.AutoMigrate(&Bank{}); err != nil {
		return err
	}

	for _, b := range banks_data {
		db.Create(&b)
	}
	return nil
}

//List - return all banks in the DB
func List() (banks []Bank) {
	db.Find(&banks)
	// log.Println("results = ", results)
	return banks
}
