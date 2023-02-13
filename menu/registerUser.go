package menu

import (
	"bufio"
	"fmt"
	"os"
	"project/config"
	"project/controllers"
	"project/models"
	"time"
)

func RegisterUser() {
	db, err := config.DBConn()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("\n")
	fmt.Println("Register Akun Baru")
	newUser := models.User{}
	fmt.Println("Nama Lengkap:")
	scanner.Scan()
	name := scanner.Text()
	newUser.Name = name
	fmt.Println("No Telepon:")
	fmt.Scanln(&newUser.Phone)
	fmt.Println("Password:")
	fmt.Scanln(&newUser.Password)
	fmt.Println("Jenis Kelamin (Pria/Wanita):")
	fmt.Scanln(&newUser.Sex)
	fmt.Println("Tanggal Lahir (d/m/y):")
	var layoutFormat, value string
	var date time.Time
	layoutFormat = "2/1/2006"
	fmt.Scanln(&value)
	date, _ = time.Parse(layoutFormat, value)
	newUser.DateOfBirth = date
	controllers.RegisterUser(db, newUser)
}
