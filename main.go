package main

import (
	"log"
	"net/http"
	"os/exec"

	"github.com/labstack/echo"
)

type File struct {
	FileName string `json:"file_name"`
	Statment string `json:"statment"`
}

func main() {

	e := echo.New()
	e.POST("/", func(c echo.Context) error {
		t := new(File)
		if err := c.Bind(t); err != nil {
			return err
		}
		arg1 := t.Statment
		arg2 := t.FileName
		cmd := " echo " + arg1 + ">" + arg2
		log.Println("cmd:", cmd)
		_, err := exec.Command("bash", "-c", cmd).Output()
		if err != nil {
			panic("some error found")
		}

		return c.String(http.StatusOK, "success")
	})
	e.Logger.Fatal(e.Start(":1323"))

}
