package main 

import (
	"fmt"
	"io"
	"bytes"
	"encoding/json"
	"net/http"
)

func commandMapBack(config *Config, args []string)error{
	var locationArea LocationArea
	if config.Previous == ""{
		return fmt.Errorf("you're on the first page")
	}
	cacheData, exist := config.cache.Get(config.Previous)
	if exist {
		r := bytes.NewReader(cacheData)
		decoder := json.NewDecoder(r)
		if err := decoder.Decode(&locationArea); err != nil {
			return err
		}
		config.Next = locationArea.Next
		if locationArea.Previous == nil {
			config.Previous = ""
		}else{
			config.Previous = locationArea.Previous.(string)
		}

		for _, location := range locationArea.Results {
			fmt.Println(location.Name)
		}
		return nil

	}else{
		res, err := http.Get(config.Previous)
		if err != nil {
			return err
		}
		defer res.Body.Close()

		
		b, errr := io.ReadAll(res.Body)
		c := bytes.NewReader(b)
		if errr != nil {
			return errr
		}
		config.cache.Add(config.Previous, b)
		decoder := json.NewDecoder(c)
		if err := decoder.Decode(&locationArea); err != nil {
			return err
		}
		
		config.Next = locationArea.Next
		if locationArea.Previous == nil {
			config.Previous = ""
		}else{
			config.Previous = locationArea.Previous.(string)
		}

		for _, location := range locationArea.Results {
			fmt.Println(location.Name)
		}
		return nil
	}
	
}