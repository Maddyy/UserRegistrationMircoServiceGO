package dbconnectionApi
import (
	
    "log"
    "database/sql"
  _ "github.com/go-sql-driver/mysql"
)
func GetConnection()(db *sql.DB,err error){
	dbDriver :="mysql"
	dbUName :="root"
	dbPassword :="root"
	dbName :="userinfo"
	db,err = sql.Open(dbDriver,dbUName+":"+dbPassword+"@tcp(127.0.0.1:3306)/"+dbName)
	return db,err
}
func GetDBUser(db *sql.DB)(rows1 *sql.Rows){
	rows1,_ =db.Query("select email,password from userinfo")
	return rows1
}
func RegistrationUser(db *sql.DB,ids string,email string, name string,password string,mobile string,city string)(rows1 string,err error){
	insForm, err := db.Prepare("INSERT INTO UserRegistration(id,email,name,password,mobile,city) VALUES(?,?,?,?,?,?)")
	if err != nil {
		panic(err.Error())
	}
	insForm.Exec(ids,email,name,password,mobile,city)
	log.Println("INSERT: Name: " + name + " | City: " + city)

     defer db.Close()   
  return  "success",err
}