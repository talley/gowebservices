package main

import (
	"fmt"
	"math"

	"github.com/gin-gonic/gin"
	"github.com/jung-kurt/gofpdf"
	"github.com/talley/gowebservices/interfaces"
	"github.com/talley/gowebservices/models"
	"gorm.io/driver/sqlite"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

const name = "Abibou"
const (
	age      = "48"
	location = "Raleigh"
)

func gorm_sqlite_all_products() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	//Migrate the schema
	db.AutoMigrate(&models.Product{})

	// Create
	db.Create(&models.Product{Code: "D42", Price: 100})
	db.Create(&models.Product{Code: "D43", Price: 101})
	db.Create(&models.Product{Code: "D44", Price: 102})
	db.Create(&models.Product{Code: "D45", Price: 103})

	var products []models.Product
	var result = db.Find(&products)
}
func gorm_sqlite() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	//Migrate the schema
	db.AutoMigrate(&models.Product{})

	// Create
	db.Create(&models.Product{Code: "D42", Price: 100})

	// Read
	var product models.Product
	db.First(&product, 1)                 // find product with integer primary key
	db.First(&product, "code = ?", "D42") // find product with code D42

	// Update - update product's price to 200
	db.Model(&product).Update("Price", 200)
	// Update - update multiple fields
	db.Model(&product).Updates(models.Product{Price: 200, Code: "F42"}) // non-zero fields
	db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})

	// Delete - delete product
	//db.Delete(&product, 1)

}
func gorm_sqlserver() {
	// github.com/denisenkom/go-mssqldb
	dsn := "sqlserver://gorm:LoremIpsum86@localhost:9930?database=gorm"
	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})

}
func main() {
	//pointers()
	//maps()
	//gofpdfdemos()

	/*	var u models.User
			fmt.Println(u)
			u.Name = "Abibou"
			u.ID = 1
			u.Location = "Raleigh"
			fmt.Println(u)

			u2 := models.User{
				ID:       1,
				Location: "New York",
				Name:     "John",
			}
			fmt.Println(u2)

		var arr []float64
		arr[0] = 12.5
		arr[1] = 3.4
		//fmt.Println(totalArea2(arr))
		numbers := [7]int{1, 2, 3, 4, 5, 6, 7}
		fmt.Println(Sum(numbers))
	*/

	router := gin.Default()
	router.SetTrustedProxies([]string{"localhost"})

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "hello gin",
		})
	})

	router.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})
	router.Run(":9097")

}

type MultiShape struct {
	shapes []interfaces.IShape
}

func Sum(numbers [7]int) int {
	var sum int = 0
	for i := 0; i <= len(numbers); i++ {
		sum = sum + numbers[i]
	}
	return sum
}
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

type Vertex struct {
	X, Y float64
}

func totalArea(shapes ...interfaces.IShape) float64 {
	var area float64
	for _, s := range shapes {
		area += s.Area()
	}
	return area
}

func totalArea2(shapes []float64) float64 {
	var area float64
	for _, s := range shapes {
		area += s
	}
	return area
}

func gofpdfdemos() {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 16)
	pdf.Cell(40, 10, "Hello, world")
	pdf.OutputFileAndClose("hello.pdf")
}
func maps() {
	ranks := map[string]int{"USA": 1}

	ranks["China"] = 2
	ranks["Japan"] = 3
	ranks["Germany"] = 4

	fmt.Println(ranks)
	fmt.Println(ranks["Germany"])

}
func pointers() {
	var firstname *string = new(string)
	*firstname = "Abibou"

	fmt.Println(*firstname)
	fmt.Println(location)
}
func start() {
	fmt.Println("Hello World")
	var myint int = 34
	fmt.Println(myint)
	var myfloat float32 = 12.4
	fmt.Println(myfloat)

	name := "Abibou"
	fmt.Println(name)
	number := 100
	fmt.Println(number)

	cp := complex(3, 4)
	fmt.Println(cp)

	a, b := 3, 4
	fmt.Println(a, b)
}
