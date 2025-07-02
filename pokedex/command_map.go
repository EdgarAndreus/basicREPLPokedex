package main

import (
	"fmt"
	"io"
	"net/http"
	"encoding/json"
	"bytes"
)

func commandMap(config *Config, args []string)error{
	var locationArea LocationArea
	fullUrl := baseUrl + "/location-area/"
	cacheData, exist := config.cache.Get(fullUrl)

	if config.Next != ""{
		fullUrl = config.Next
	}
	if exist {
		r := bytes.NewReader(cacheData)
		//c := io.Reader(cacheData)
		decoder := json.NewDecoder(r)
		if err := decoder.Decode(&locationArea); err != nil {
			return err
	}

		config.Next = locationArea.Next
		if locationArea.Previous == nil{
			config.Previous = ""
		}else{
			config.Previous = locationArea.Previous.(string)
	}
	

		for _, name := range locationArea.Results{
			fmt.Println(name.Name)
		}
		return nil
	}else{
		res, err := http.Get(fullUrl)
		if err != nil {
			return err
		}
		defer res.Body.Close()
		b, errr := io.ReadAll(res.Body)
		c := bytes.NewReader(b)
		if errr != nil {
			return errr
		}
		config.cache.Add(fullUrl, b)
	
		decoder := json.NewDecoder(c)
		if err := decoder.Decode(&locationArea); err != nil {
			return err
		}
	
		config.Next = locationArea.Next
		if locationArea.Previous == nil{
			config.Previous = ""
		}else{
			config.Previous = locationArea.Previous.(string)
		}
		
	
		for _, name := range locationArea.Results{
			fmt.Println(name.Name)
		}
	
		return nil
	}
	


	

}