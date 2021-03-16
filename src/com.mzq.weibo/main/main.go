package main

import (
	"io/ioutil"
	"net/http"
)

func main() {
	//for true {
	//	go func() {
			url := "https://m.weibo.cn/profile/info?uid=1987241375"

			request, err := http.NewRequest("GET", url, nil)
			if err != nil {
				return
			}
			request.Header.Add("Accept", "application/json, text/plain, */*")
			//cookie3 := &http.Cookie{Name: "WEIBOCN_FROM",Value: "1110006030"}
			//request.AddCookie(cookie3)
			cookie := &http.Cookie{Name: "MLOGIN",Value: "1"}
			request.AddCookie(cookie)
			//cookie5 := &http.Cookie{Name: "M_WEIBOCN_PARAMS",Value: "oid%3D4538243879540070%26luicode%3D20000061%26lfid%3D4538243879540070"}
			//request.AddCookie(cookie5)
			cookie = &http.Cookie{Name: "SUB",Value: "_2A25yM6KSDeRhGeVL7FsT8irMyD2IHXVR387arDV6PUJbkdAKLU7lkW1NTD9Yt3SUG2xQeSdkSp9-EAi4dMx3kjJ8"}
			request.AddCookie(cookie)
			//cookie7 := &http.Cookie{Name: "SUHB",Value: "00xpQeZBu0KlRg"}
			//request.AddCookie(cookie7)
			client := &http.Client{}

			response, _ := client.Do(request)


			body, err1 := ioutil.ReadAll(response.Body)
			if err1 != nil {
				// handle error
			}
			println(string(body))

			defer response.Body.Close()
			//http.Get(url)
		//}()
		//time.Sleep(time.Minute)
	//}
}
