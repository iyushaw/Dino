package dynowebportal

import (
	"fmt"
	"net/http"
)

//Main RunWebPortal
func RunWebPortal(addr string)error{
	http.HandleFunc("/",roothandler)
	err := http.ListenAndServe(addr,nil)
	return err
}

func roothandler(w http.ResponseWriter,r *http.Request){
	fmt.Fprintf(w,"Welcome to Dino web portal %s", r.RemoteAddr)
}