package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/kataras/iris/v12"
)

func main() {
	app := iris.New()
	app.Logger().SetLevel("debug")

	app.Get("/", download)
	//app.Get("/download", downloadWithRateLimit)
	app.Get("/fileapi/{types:string}/{number:string}",func(ctx iris.Context) {
		types :=ctx.Params().Get("types")
		number :=ctx.Params().Get("number")
		log.Print("Type:"+types+"Number:"+number)

	}) //URL-API

	app.Get("/getdata",func(ctx iris.Context) {
		ctx.HTML("API：" + ListFile())  
	})
	app.Listen(":8080")
}

func download(ctx iris.Context) {
	src := "./files/first.zip"
	ctx.SendFile(src, "client.zip")
}


func ListFile() string {
    dir_list, e := ioutil.ReadDir("type")
    if e != nil {
        fmt.Println("read dir error")
        return "error"
    }
	/*
    for i, v := range dir_list {
        fmt.Println(i, "=", v.Name())
    }*/
	
	tatol := "<list>"
	for i,v := range dir_list{
		tatol = tatol+v.Name()+"</list>"
		//log.Print(i,string(rune(i)))
		fmt.Println(i,v.Name()+string(rune(i)))
	}
	log.Print(tatol)
	return tatol
}

/*
func downloadWithRateLimit(ctx iris.Context) {
	 REPLACE THAT WITH A BIG LOCAL FILE OF YOUR OWN.
	src := "./files/first.zip"
	dest := "" /* optionally, keep it empty to resolve the filename based on the "src" 

	 Limit download speed to ~50Kb/s with a burst of 100KB.
	limit := 50.0 * iris.KB
	burst := 100 * iris.KB
	ctx.SendFileWithRate(src, dest, limit, burst)
}
*/
