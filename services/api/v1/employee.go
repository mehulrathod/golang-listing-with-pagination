package v1Service

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	u "testTask/apiHelpers"
	"testTask/models"
	resources "testTask/resources/api"
)

type EmployeeService struct {
	Product models.Employees
}

func GetAllEmployeeApi(c *gin.Context) map[string]interface{} {
	fmt.Println("service called")
	var v interface{}
	json.NewDecoder(c.Request.Body).Decode(&v)
	pa := v.(map[string]interface{})

	fmt.Println("=====================",c.Request.Body)
	fmt.Println(">>>>>", v)
	fmt.Println("pa    ------>>>>>", pa)

	p := int(pa["page"].(float64))
	pageSize := int(pa["pageSize"].(float64))

	fmt.Println("p ----->", p)
	fmt.Println("pageSize ----->", pageSize)

	var EmployeeData []resources.EmployeeResponse

	var queryMain = models.GetDB().Debug().Select("*").Table("employees").
		Limit(pageSize).Offset(p * pageSize).Order("employees.id")
	queryMain.Find(&EmployeeData)

	var queryTotal = models.GetDB().Table("employees")
	var total uint
	queryTotal.Count(&total)


	response := u.Message(0, "Employee Lists.")
	response["data"] = EmployeeData
	response["total"] = total
	return response
}
