package Utils

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

func UPLOADData(adress string,body io.Reader){
	client := &http.Client{}
	requestPost, err := http.NewRequest("POST", adress,body)
	resp, err := client.Do(requestPost)
	//u2:= uuid.NewV4()
	//requestPost.Header.Set("ProtoType","json")
	//requestPost.Header.Set("Br-Content-Encoding","gzip")
	requestPost.Header.Set("Content-Type","application/json")
	//requestPost.Header.Set("brkey",u2.String())
	//requestPost.Header.Add("Accept-Encoding", "gzip")
	//resp, err:=http.Post(adress,"application/x-www-form-urlencoded",body)
	//resp, err:=http.PostForm(adress,url.Values{"ProtoType":{"json"},"brkey":{u2.String()},"body":{body}})
	if err != nil {
		fmt.Printf("get request failed, err:[%s]", err.Error())
		return
	}
	defer resp.Body.Close()
	bodyContent, err := ioutil.ReadAll(resp.Body)
	fmt.Printf("resp status code:[%d]\n", resp.StatusCode)
	fmt.Printf("resp body data:[%s]\n", string(bodyContent))
}
func CONFIGData(adress string,body io.Reader) string{
	client := &http.Client{}
	requestPost, err := http.NewRequest("POST", adress,body)
	resp, err := client.Do(requestPost)
	requestPost.Header.Set("Content-Type","application/x-www-form-urlencoded")
	requestPost.Header.Set("ProtoType","json")
	if err != nil {
		fmt.Printf("get request failed, err:[%s]", err.Error())
		return "get request failed!"
	}
	defer resp.Body.Close()
	bodyContent, err := ioutil.ReadAll(resp.Body)
	fmt.Printf("resp status code:[%d]\n", resp.StatusCode)
	fmt.Printf("resp body data:[%s]\n", string(bodyContent))
	return string(bodyContent)
}