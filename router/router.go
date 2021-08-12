package router

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func readDir(dir string, res *[]string) {
	list, err := os.ReadDir(dir)

	if err != nil {
		log.Fatalln(err)
	}

	for _, f := range list {
		path := fmt.Sprintf("%s/%s", dir, f.Name())
		if f.IsDir() {
			readDir(path, res)
		} else {
			*res = append(*res, path)
		}
	}
}

func getTemplates() []string {
	files := []string{}
	readDir("templates", &files)
	return files
}

func InitRouter() *gin.Engine {
	r := gin.Default()

	files := getTemplates()
	fmt.Printf("%v", files)
	r.LoadHTMLFiles(files...)

	r.GET("/", showIndexPage)
	r.GET("/artical", showArticals)

	return r
}
