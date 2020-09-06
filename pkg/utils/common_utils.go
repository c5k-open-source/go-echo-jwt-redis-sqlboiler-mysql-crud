package utils

import (
	"encoding/json"
	"strings"

	"github.com/labstack/echo"
)

func PrettyString(data interface{}) string {
	s, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		panic(err)
	}
	return string(s)
}

func GetIP(c echo.Context) string {
	realIP := c.RealIP()
	if strings.Contains(realIP, ",") {
		realIP = strings.Split(realIP, ",")[0]
		if strings.Contains(realIP, ":") {
			realIP = c.Request().Header.Get("Cf-Pseudo-Ipv4")
		}
	}
	return realIP
}
