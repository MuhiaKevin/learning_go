package main


import(
	"io/ioutil"
	"fmt"
	"log"
)


func main(){
	contents, error:= ioutil.ReadFile("/home/muhia/Downloads/torrents/FreeTutorialsUS.com-Udemy-Build-Realtime-Apps-React-Js-Golang-and-RethinkDB.torrent")
	if error != nil{
		log.Fatal(error)
	}

	fmt.Printf("File contents: %s", contents)

}