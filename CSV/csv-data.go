package CSV

import (
	"ads/database"
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

type tbl_role struct {
	id   string
	name string
}

func (tbl_role) TableName() string {
	return "tbl_role"
}
func ReadCSV() {
	roleData := []tbl_role{}
	file, err := os.Open("./role.csv")
	if err != nil {
		log.Fatal(err)
	}
	df := csv.NewReader(file)
	data, _ := df.ReadAll()

	for _, value := range data {
		roleData = append(roleData, tbl_role{id: value[0], name: value[1]})
		fmt.Println("Read all data", value)
	}
	fmt.Println("Read all data", roleData[1])

	for i := 1; i < len(roleData); i++ {

		//id, _ := strconv.Atoi(roleData[i].id)
		//database.DB.Exec("insert into tbl_role(id,name) values(?,?)", id, roleData[i].name)
		data := tbl_role{
			id:   roleData[i].id,
			name: roleData[i].name,
		}
		fmt.Println("Read all data", roleData)

		fmt.Println(" roleData[i].id :roleData[i].name", roleData[i].id, roleData[i].name)

		fmt.Println(" roleData[i].name", roleData[i].name)
		database.DB.Create(&data)

	}

	fmt.Println("All records inserted")
}
