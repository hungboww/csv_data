package CSV

import (
	"ads/database"
	"fmt"
	"strconv"
)

type tbl_role struct {
	id   string `csv:"id"`
	name string `csv:"name"`
}

func (tbl_role) TableName() string {
	return "tbl_role"
}

func InsertDataCSV(rows [][]string) {
	roleData := []tbl_role{}

	for _, value := range rows {
		roleData = append(roleData, tbl_role{id: value[0], name: value[1]})
		fmt.Println("Read all data", value)
	}
	for i := 1; i < len(roleData); i++ {
		id, _ := strconv.Atoi(roleData[i].id)
		database.DB.Exec("insert into tbl_role(id,name) values(?,?)", id, roleData[i].name)
	}

	fmt.Println("All records inserted")
}
