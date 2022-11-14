package CSV

import (
	"ads/database"
	"fmt"
	"strconv"
	"strings"
)

type valueAccount struct {
	id           string `csv:"id"`
	password     string `csv:"password"`
	last_login   string `csv:"last_login"`
	is_superuser bool   `csv:"is_superuser"`
	email        string `csv:"email"`
	user_name    string `csv:"user_name"`
	first_name   string `csv:"first_name"`
	start_date   string `csv:"start_date"`
	about        string `csv:"about"`
	image        string `csv:"image"`
	is_active    bool   `csv:"is_active"`
	is_staff     bool   `csv:"is_staff"`
	role_id      string `csv:"role_id"`
}

func (valueAccount) TableName() string { return "tbl_user" }

func InsertDataAccount(rows [][]string) {
	roleData := []valueAccount{}
	for _, value := range rows {
		roleData = append(roleData, valueAccount{
			id:           value[0],
			password:     value[1],
			last_login:   value[2],
			is_superuser: convertStringToBool(value[3]),
			email:        value[4],
			user_name:    value[5],
			first_name:   value[6],
			start_date:   value[7],
			about:        value[8],
			image:        value[9],
			is_active:    convertStringToBool(value[10]),
			is_staff:     convertStringToBool(value[11]),
			role_id:      value[12],
		})
	}
	s := ""
	for i := 1; i < len(roleData); i++ {
		id, _ := strconv.Atoi(roleData[i].id)
		s += fmt.Sprintf("(%d, %s, %s, %t,%s, %s,%s, %s,%s, %t,%t, %s, %s),", id, roleData[i].password, roleData[i].last_login, roleData[i].is_superuser, roleData[i].email, roleData[i].user_name, roleData[i].first_name, roleData[i].about, roleData[i].image, roleData[i].is_active, roleData[i].is_staff, roleData[i].start_date, roleData[i].role_id)
	}
	sql := fmt.Sprintf(`INSERT INTO tbl_user(id,password,last_login, is_superuser,email , user_name, first_name,about, image, is_active, is_staff, start_date, role_id) values  %s`, strings.TrimSuffix(s, ","))
	err := database.DB.Exec(sql)
	if err != nil {
		fmt.Println("Can't insert CSV file into database. Please try again!!!", err)
	}
	fmt.Println("成功した挿入")
}

func convertStringToBool(value string) bool {
	value = strings.ToLower(value)
	a, _ := strconv.Atoi(value)
	if a == 0 {
		return false
	}
	return true
}
