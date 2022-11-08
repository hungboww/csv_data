package CSV

import (
	"ads/database"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
)

type valueAccout struct {
	id           string
	password     string
	last_login   string
	is_superuser string
	email        string
	user_name    string
	first_name   string
	start_date   string
	about        string
	image        string
	is_active    string
	is_staff     string
}

func ReadCSVAcount() {
	roleData := []valueAccout{}
	file, err := os.Open("./account.csv")
	if err != nil {
		log.Fatal(err)
	}
	df := csv.NewReader(file)
	data, _ := df.ReadAll()

	for _, value := range data {
		roleData = append(roleData, valueAccout{id: value[0], password: value[1], last_login: value[2], is_superuser: value[3], email: value[4], user_name: value[5], first_name: value[6], start_date: value[7], about: value[8], image: value[9], is_active: value[10], is_staff: value[11]})
		fmt.Println("Read all data", value)
	}

	for i := 1; i < len(roleData); i++ {

		id, _ := strconv.Atoi(roleData[i].id)
		database.DB.Exec("insert into tbl_user * values(?,?)", id, roleData[i])
	}
	fmt.Println("All records inserted")
}
