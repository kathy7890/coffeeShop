package main

import (
	"coffeeShop/coffee"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"io"
	"net/http"
	"os"
)

// development tool. intellji with linking support, with function being able to call.
/*

 */
/*
application level: just build and run application
# update depedencies :
# go mod tidy

# go run coffeeShop  // coffeeShop is the package name on the top of go.mod file
# or: go run main.go

docker level: docker command
# docker build -t coffeeshop_app .  // build a image with image name "coffeeshop_app" with current dir.
# docker run --env-file ./config.env -p 8090:8085 coffeeshop_app   // coffeeshop_app

docker compose level: docker compose is a tool to run image(i.e. start container/service).
# docker-compose up


1. start the container with bash shell
$ docker run -it -v "$(pwd):/app" --env-file ./config.env -p 8090:8085 coffeeshop_app bash
# OR
$ docker run -it --mount "type=bind,source=$(pwd)/source_dir,target=/app/target_dir" coffeeshop_app bash


2. build the executible on the shell in the container. We don't build it locally on host machine, because maybe
different platforms between host machine and container. (mac vs linux).
$ go build -o coffeeShop .

3. start the executible in the container.
# ./coffeeShop

In the end, access the service through http://localhost:8090

After the executive has been build, you can run the command directly, this is based on executive has been built. Docker seems to have those volume as persistent data.
$ docker run -v "$(pwd):/app" --env-file ./config.env -p 8090:8085 coffeeshop_app

At the same time, the following command also works, just same as above command. all the info has been defined in
docker-composer.html.
$ docker-compose up


So, maybe a more concise approach is going to container shell + build the executive + run the executive, this approach
// goes to create a new container.  It requires killing process with pid 1 if using existing process. more twist involved.
// when we start the container, the command of starting the application gets running, due to the CMD command in dockerfile.

$ docker run -it -v "$(pwd):/app" -p 8090:8085 coffeeshop_app bash  // without "bash", it'll start the application. the argument "bash" replace CMD argument in makefile.
// so we can run the container without running the application.
// also, "docker run " will create a new container. if you want to base on existing container,
// commit changes as new image name and start container from new image, like the following:
(
$ docker commit 5a8f89adeead newimagename. // 5a8f89adeead is existing container id.
$ docker run -ti -v "$PWD/somedir":/somedir newimagename /bin/bash
)

$ go build -o coffeeShop .

$ ./coffeeShop

note:
docker run : will create a new container.  like docker-compose up
docker start: restart a stopped container with all its previous changes intact , like docker-compose start.



where docker caches comes in. just to test
*/

//Todo:
// docker image and docker run  # could not finish fetch dependency.
// docker production ready  # multi-stage build and from sctratch to reduce size.

// add testing starting from now: http testing
// other feature improvement like page and UI to make it more full. like adding database, picture, ordering,
// what other feature could be? like parallel computing of ordering. concurrency patterns.


func main() {
	// http.HandleFunc("/ping", ping)

	// fmt.Println("Starting the server at port 8081")
	// err := http.ListenAndServe(":8081", nil) // pass nill as handler, in which case the DefaultServeMux is used.
	// if err != nil { // ListenAndServe always returns non-nil error.
	// 	fmt.Println("Error while starting the server ", err)
	// }

	portNumber := os.Getenv("APP_PORT")
	r := gin.Default()
	r.LoadHTMLGlob("templates/*.html")

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Welcome to the Coffeeshop123!",
		})
	})
	r.GET("/home", getCoffeeListPage)
	r.GET("/getCoffees", getCoffees)

	log.Info().Msg("=====from zerolog")
	fmt.Println("======789 starting port portNumber=====")
	fmt.Println(portNumber)
	fmt.Println("Starting the server at port: ", portNumber)
	r.Run(fmt.Sprintf(":%s", portNumber))
}

func ping(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Welcome request")
	io.WriteString(w, "Welcome to the Coffeeshop!\n")
}

func pingGin(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Welcome to the Coffeeshop in Gin!",
	})
}

func getCoffeeListPage(c *gin.Context) {
	coffeelist, _ := coffee.GetCoffees()
	//c.String(http.StatusOK, " %s", coffeelist)
	// Call the HTML method of the Context to render a template
	c.HTML(
		http.StatusOK,
		"index.html",
		gin.H{
			"list": coffeelist.List,
		},
	)
}

func getCoffees(c *gin.Context) {
	coffeelist, _ := coffee.GetCoffees()
	//c.String(http.StatusOK, " %s", coffeelist)
	c.JSON(http.StatusOK, coffeelist)
	// c.JSON(http.StatusOK, gin.H{
	// 		"message": "Welcome to the Coffeeshop in Gin!",
	// })
}

func getCoffee(c *gin.Context) {
	coffeelist, _ := coffee.GetCoffees()
	c.String(http.StatusOK, " %s", coffeelist)
}
