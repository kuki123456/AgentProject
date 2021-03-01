package main

import (
	"AgentProject/Utils"
	"fmt"
	"github.com/tidwall/gjson"
	"time"
)

func main() {
	da:=time.Now().Format("2006-01-02 15")
	fmt.Println(da)
	networkbody:="{\"ent\":1612148754575438,\"k\":\"network\",\"sin\":[\"\",\"1611910357403D0EE4B9C-CA5F-438F-941F-18A7069AB3F4\",\"f0f4e7e1045399eb3313a880297874b3\"],\"v\":{\"eop\":3,\"ep\":\"NSURLErrorDomain\",\"dt\":0,\"tp\":8080,\"ct\":7000,\"ec\":-1005,\"rt\":2298,\"art\":10,\"sslt\":0,\"rti\":0,\"ru\":\"http:\\/\\/192.168.1.152:8080\\/upload?v=2020122401&a=6d97637b-ab2c-479a-90fa-59d2417ba4a0&d=82AD80A2-A4A1-4C43-B945-54C3A8336FF6\",\"pt\":1,\"ti\":\"192.168.1.152\",\"m\":\"POST\",\"ic\":false,\"dti\":0,\"ds\":0}}"
	fmt.Printf("%v",gjson.Get(networkbody,"v.ic").Type)
	asd:=[]float64{1,2,3,4,10}
	fmt.Println(gjson.Get(networkbody,"v.art").Num)
	fmt.Println(Utils.Assertin(gjson.Get(networkbody,"v.art").Num,asd))
	var testinterface interface{}
	testinterface=12
	fmt.Println(testinterface.(int))
	v,ok:=testinterface.(int)
	fmt.Println(v)
	fmt.Println(ok)
}
