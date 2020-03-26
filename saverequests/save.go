package saverequests

import (
	"fmt"
	"os"
)

func SaveLocationRequest(path []byte, payload []byte, rlocation string){
         request := fmt.Sprintf(`curl -X POST -H "Content-Type: application/json" --data '%s' $1%s`+"\n", payload, path)
         f, err := os.OpenFile("/tmp/" + rlocation +".sh", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
         if err != nil {
		 fmt.Printf("An error occurred: %+v", err)
         }
         defer f.Close()
         if _, err := f.WriteString(request); err != nil {
		 fmt.Printf("An error occurred: %+v", err)
         }
         return
}

