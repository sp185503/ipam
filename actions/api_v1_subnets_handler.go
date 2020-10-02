package actions

import (
	"net/http"

	"github.com/gobuffalo/buffalo"
)

// apiV1SubnetHandler we fetch the subnets for the corresponding BU
func apiV1SubnetHandler(c buffalo.Context) error {
	buName := c.Param("buName")
	return c.Render(http.StatusOK, r.JSON(map[string]string{"message": "fetching subnets for " + buName}))
}
