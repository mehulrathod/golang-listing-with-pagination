package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	u "testTask/apiHelpers"
	v1sv "testTask/services/api/v1"
)

func EmployeesList(c *gin.Context) {
	fmt.Println("controller called")
	resp := v1sv.GetAllEmployeeApi(c)
	u.Respond(c.Writer, resp)
}
