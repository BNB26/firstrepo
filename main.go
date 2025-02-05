package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Employee struct {
	Id   int
	Name string
}

var EmployeeList = []Employee{
	{1, "Basavaraj"},
	{2, "Bhavikatti"},
}

// handler func
func getEmployee(c *gin.Context) {
	c.JSON(http.StatusOK, &EmployeeList)
}

func createEmployee(c *gin.Context) {
	var newEmployee Employee
	err := c.BindJSON(&newEmployee)
	if err != nil {
		log.Fatal("error bionding new employee")
		return
	}
	EmployeeList = append(EmployeeList, newEmployee)
	c.JSON(http.StatusCreated, newEmployee)
}

func updateEmployee(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var updateEmployee Employee
	err := c.BindJSON(&updateEmployee)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "error binding json"})
		return
	}
	for i, employee := range EmployeeList {
		if id == employee.Id {
			EmployeeList[i].Name = updateEmployee.Name
			//employee.Id = updateEmployee.Id
			c.JSON(http.StatusAccepted, EmployeeList[i])
			return
		}
	}
	c.JSON(http.StatusBadRequest, gin.H{"error": "Id not found"})
}

func deleteEmployee(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	for i, employee := range EmployeeList {
		if id == employee.Id {
			EmployeeList = append(EmployeeList[:i], EmployeeList[i+1:]...)
			c.JSON(http.StatusOK, EmployeeList)
			return
		}
	}
	c.JSON(http.StatusBadRequest, gin.H{"error": "employee not found"})
}

func getSpecificEmployee(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	for i, employee := range EmployeeList {
		if id == employee.Id {
			c.JSON(http.StatusOK, EmployeeList[i])
			return
		}
	}
	c.JSON(http.StatusBadRequest, gin.H{"error": "id not found"})
}

func main() {
	fmt.Println("hi")
	//route
	r := gin.Default()

	//endpoint
	r.GET("/getEmployee", getEmployee)
	r.POST("/createEmployee", createEmployee)
	r.PUT("/updateEmployee/:id", updateEmployee)
	r.DELETE("/deleteEmployee/:id", deleteEmployee)
	r.GET("/getSpecificEmployee/:id", getSpecificEmployee)

	//port
	r.Run(":8088")
}

// Among all most popular frameworks mostly used for building a RestAPIs Gin is most popular because of following features it brings in:

// Gin framework makes use of HTTP routers.
// Gin framework have rich documentation provided on Github.
// Gin framework supports most essential libraries and features.
// Gin framework is similar but 40 times faster than Martini.
// Gin frameworks have well tested middleware
