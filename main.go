package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

 func setupHeader(rw http.ResponseWriter, req *http.Request) {
        rw.Header().Set("Content-Type", "application/json")
        rw.Header().Set("Access-Control-Allow-Origin", "*")
        rw.Header().Set("Access-Control-Allow-Methods", "GET")
        rw.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}


func index(w http.ResponseWriter, r *http.Request) {
   setupHeader(w, r)

// 	postBody, _ := json.Marshal(map[string]string{
//       "dados":  "dados",
//    })
//    responseBody := bytes.NewBuffer(postBody)

// 	resp, err := http.Post("https://deolhonafila.prefeitura.sp.gov.br/processadores/dados.php","application/x-www-form-urlencoded; charset=UTF-8", responseBody)
// if err != nil {
//    log.Fatalln(err)
// } 
// //We Read the response body on the line below.
//    body, err := ioutil.ReadAll(resp.Body)
//    if err != nil {
//       log.Fatalln(err)
//    }
// //Convert the body to type string
//    sb := string(body)

// postBody, _ := json.Marshal(map[string]string{"dados":  "dados",})

client := http.Client{}
req , err := http.NewRequest("POST", "https://deolhonafila.prefeitura.sp.gov.br/processadores/dados.php", strings.NewReader("dados=dados"))
if err != nil {
    //Handle Error
}



req.Header = http.Header{
    "Host": []string{"deolhonafila.prefeitura.sp.gov.br"},
    "Content-Type": []string{"application/x-www-form-urlencoded; charset=UTF-8"},
    "Accept": []string{"application/json, text/javascript, */*; q=0.01"},
	
}

log.Println(req)

res , err := client.Do(req)
if err != nil {
    //Handle Error
}


   body, err := ioutil.ReadAll(res.Body)
   if err != nil {
      log.Fatalln(err)
   }
//Convert the body to type string
//    sb := string(body)

   // fmt.Fprint(w, string(body))

  


   w.Write(body)
}


func main() {
	// Upload route
	http.HandleFunc("/", index)


	//Listen on port 8080
	http.ListenAndServe(":8080", nil)
}