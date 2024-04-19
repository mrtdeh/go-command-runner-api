package main

import (
	"flag"
	"io"
	"log"
	api_server "mrtdeh/exec-api/pkg/api"
	"os/exec"

	"github.com/gin-gonic/gin"
)

func main() {

	host := flag.String("n", "0.0.0.0", "api host")
	port := flag.Uint("p", 8080, "api port")
	flag.Parse()

	log.Println("start service")
	gin.SetMode(gin.ReleaseMode)
	gin.DefaultWriter = io.Discard

	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.POST("/run", Exec)
	httpServer := api_server.HttpServer{
		Host:   *host,
		Port:   *port,
		Router: r,
	}
	log.Fatal(httpServer.Serve())
}

type ExecRequest struct {
	Command string `json:"cmd"`
}

func Exec(c *gin.Context) {
	var execReq ExecRequest
	if err := c.ShouldBind(&execReq); err != nil {
		c.String(400, "invalid input data : ", err)
		return
	}

	cmd := exec.Command("sh", "-c", execReq.Command)
	out, err := cmd.CombinedOutput()

	if err != nil {
		log.Printf("error '%s' : %v", cmd, err)
		c.String(500, "%s", string(out))
		return
	}
	c.String(200, "%s", string(out))
}
