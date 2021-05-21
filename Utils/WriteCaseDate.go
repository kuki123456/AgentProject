package Utils

import (
	"fmt"
	"os"
)

func WriteCaseDate(casename,moudle,casedata string){
	file,_:=os.OpenFile(fmt.Sprintf("./CaseData/%s.txt",moudle),os.O_APPEND|os.O_CREATE|os.O_WRONLY|os.O_RDONLY,os.ModeAppend|os.ModePerm )
	defer file.Close()
	switch moudle{
	case "CustomApi":
		_, _ = fmt.Fprintf(file, "%s:%s\n", casename, casedata)
	}
}