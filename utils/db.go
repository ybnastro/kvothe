package utils

import (
	"fmt"

	"github.com/SurgicalSteel/kvothe/resources"
)

//GeneratePostgreURL generates postgresql URL from DB Account struct
func GeneratePostgreURL(dbAcc resources.DBAccount) string {
	return fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s  sslmode=disable extra_float_digits=-1 connect_timeout=%s", dbAcc.Username, dbAcc.Password, dbAcc.DBName, dbAcc.URL, dbAcc.Port, dbAcc.Timeout)
	//return "user=" + dbAcc.Username + " password=" + dbAcc.Password + " dbname=" + dbAcc.DBName + " host=" + dbAcc.URL + " port=" + dbAcc.Port + "  sslmode=disable extra_float_digits=-1 connect_timeout=" + dbAcc.Timeout
}
