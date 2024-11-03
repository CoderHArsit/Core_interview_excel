package main

import (
	"interviewexcel-backend-go/config"
	"interviewexcel-backend-go/routes"

	"github.com/fatih/color"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	config.InitDB()

	routes.RegisterExpertRoutes(r)
	banner := `
		
		  ██╗     ███████╗████████╗███████╗    ███████╗██╗  ██╗███████╗██╗     
  ██║     ██╔════╝╚══██╔══╝██╔════╝    ██╔════╝██║  ██║██╔════╝██║     
  ██║     █████╗     ██║   █████╗      ███████╗███████║█████╗  ██║     
  ██║     ██╔══╝     ██║   ██╔══╝      ╚════██║██╔══██║██╔══╝  ██║     
  ███████╗███████╗   ██║   ███████╗    ███████║██║  ██║███████╗███████╗
  ╚══════╝╚══════╝   ╚═╝   ╚══════╝    ╚══════╝╚═╝  ╚═╝╚══════╝╚══════╝

		Welcome to InterviewExcel Backend - Powered by Go and PostgreSQL!
		`
	color.Green(banner)

	r.Run()

}
