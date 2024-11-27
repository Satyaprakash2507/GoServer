package main
import(
	"fmt"
	"log"
	"net/http"
)
func helloHandler( w http.ResponseWriter , r *http.Request)

func formHandler( w http.ResponseWriter, r *http.Request)


func main(){
fileServer := http.FileServer(http.Dir("./static"))//it will look for static file and by default pick the index.html
http.Handle("/", fileServer) //handle root route which is just above 
http.HandleFunc("/form", formHandler) // handle formhandler
http.HandleFunc("/hello", helloHandler) //handle hellohandler

fmt.Printf("Starting server at port 8080 \n") 
//starting server on 8080 
if err:= http.ListenAndServer(":8080", nil); err != nil{
	log.fatal(err)
}

}