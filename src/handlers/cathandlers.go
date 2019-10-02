package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"vo"

	"github.com/labstack/echo"
)

//http://localhost:8000/cats/json?name=arnold&type=fluffy
func GetCats(c echo.Context) error {
	catName := c.QueryParam("name")
	catType := c.QueryParam("type")
	dataType := c.Param("data")
	if dataType == "string" {
		return c.String(http.StatusOK, fmt.Sprintf("your cat name is : %s\nand cat type is : %s\n", catName, catType))
	} else if dataType == "json" {
		cat := vo.Cat{
			Name: catName,
			Type: catType,
		}
		return c.JSON(http.StatusOK, cat)
	} else {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Please specify the data type as Sting or JSON",
		})
	}

}

func AddCat(c echo.Context) error {
	cat := vo.Cat{}
	defer c.Request().Body.Close()

	err := json.NewDecoder(c.Request().Body).Decode(&cat)
	if err != nil {
		log.Fatalf("Failed reading the request body %s", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error)
	}
	log.Printf("this is yout cat %#v", cat)
	return c.String(http.StatusOK, "We got your Cat!!!")
}
