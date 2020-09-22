package main
import (
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"fmt" 
	//"log"
  )

  type Product struct {
	Id  int
	Code  string
	Price uint
  }


  func main() {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	//dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := "docker:docker@tcp(db:3306)/test_db?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := sql.Open("mysql",dsn)
	if err != nil {
		panic("failed to connect database")
	}
	// query all data
	rows, e := db.Query("select id, code from products")
	if e != nil {
		panic("failed to products table")
	}
	/*for rows.Next() {
		var (
			id   int64
			code string
		)
		if err := rows.Scan(&id, &code); err != nil {
			panic(err)
		}
		fmt.Printf("id %d code is %s\n", id, code)
	}*/

	contactos := []Product{}
	// iterate over rows
	for rows.Next() {
		c2 := Product{}
		if err := rows.Scan(&c2.Id, &c2.Code); err != nil {
			panic(err)
		}
		fmt.Printf("id %d codex is %s\n", c2.Id, c2.Code)
		contactos = append(contactos, c2)
	}

	for _, contacto := range contactos {
		fmt.Printf("%v\n", contacto.Code)
	}


	//panic(db)
	/* Migrate the schema
	db.AutoMigrate(&{})

	// Create
	db.Create(&Product{Code: "D42", Price: 100})

	// Read
	var product Product
	db.First(&product, 1) // find product with integer primary key
	db.First(&product, "code = ?", "D42") // find product with code D42

	// Update - update product's price to 200
	db.Model(&product).Update("Price", 200)
	// Update - update multiple fields
	db.Model(&product).Updates(Product{Price: 200, Code: "F42"}) // non-zero fields
	db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})
  
	// Delete - delete product
	db.Delete(&product{}, 1)
	db.Unscoped().Delete(&product, 1)
	*/



  }