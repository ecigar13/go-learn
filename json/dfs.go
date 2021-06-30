package main

import (
	"errors"
	"reflect"
)

func dfsKey(item interface{}, searchKey string) (interface{}, error){
	valType:=reflect.TypeOf(item)
	
	switch valType.Kind(){
	case reflect.Slice, reflect.Array:
		for _, data := range item.([]interface{}){
			ans, error := dfsKey(data, searchKey)
			if error == nil{
				return ans, error
			}
		}
		return "",errors.New("can't find the key")

	case reflect.Map:
		if val, ok:= item.(map[string]interface{})[searchKey]; ok{
			return val, nil
		}else{
			for _, val := range item.(map[string]interface{}){
				ans, error:=dfsKey(val, searchKey)
				if error == nil{
					return ans, error
				}
			}
			return "",errors.New("can't find the key")
		}

	default:
		return "", errors.New("can't find the key")
	}
}

func dfsValue(item interface{}, searchValue string) (bool, errors){
	valType:=reflect.TypeOf(item)
	
	switch valType.Kind(){
	case reflect.Slice, reflect.Array:
		for _, val:= range item.([]interface{}){
			if val.(string) == searchValue{
				return true, nil
			}else {
				_, err := dfsKey(val, searchValue)
				if err == nil{
					return true, err
				}
			}
		}
		return false, errors.New("can't find value")

	case reflect.Map:
		for _, val := range item.(map[string]interface{}){
			if val.(string) == searchValue{
				return true, nil
			}else{
				_, err := dfsValue(val, searchValue)
				if err == nil{
					return true, err
				}
			}
		}


	default:
		return false, errors.New("can't find the value")
	}
}
