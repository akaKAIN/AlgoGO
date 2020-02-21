package LoggerParser

import (
"fmt"
"log"
"os"
"time"
)

var(
	outFile, _ = os.OpenFile("errors.log", os.O_CREATE|os.O_RDWR|os.O_APPEND, os.ModePerm)
	LogFile = log.New(outFile, fmt.Sprintf(time.Now().Format("02.01.2006-15:04:05  ")), 0)
)

func ForError(err error){
	if err != nil{
		LogFile.Fatalln(err)
	}
}
