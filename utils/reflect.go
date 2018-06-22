package utils

import (
	"fmt"
	"reflect"
	"github.com/vmihailenco/msgpack"
)

func setFiled(obj interface{},name string,value interface{})error{
	structData:=reflect.ValueOf(obj).Elem();
	fieldValue:=structData.FieldByName(name);
	if !fieldValue.IsValid() {
		return fmt.Errorf("utils.setField() not such field:%s in obj",name)
	}
	if !fieldValue.CanSet() {
		return fmt.Errorf("Canot set %s field value",name)
	}

	fieldType:=fieldValue.Type()
	val:=reflect.ValueOf(value)

	valTypeStr:=val.Type().String()
	fieldTypeStr:=fieldType.String()

	if valTypeStr=="float64" && fieldTypeStr=="int" {
		val=val.Convert(fieldType)
	}else if fieldType!=val.Type() {
		return fmt.Errorf("provided value type"+valTypeStr+"did't match obj field type "+fieldTypeStr)
	}
	fieldValue.Set(val)
	return nil
}


func SetStructByJSON(obj interface{},mapData map[string]interface{}) error {
	for key,value:=range mapData {
		if err:=setFiled(obj,key,value);err!=nil {
			fmt.Println(err.Error())
			return err;
		}
	}
	return nil
}

func TransformToOther(in interface{}, out interface{}) error {
	// 序列化参数
	data, err := msgpack.Marshal(in)
	if err != nil {
		return err
	}
	// 反序列化参数
	err = msgpack.Unmarshal(data, out)
	if err != nil {
		return err
	}
	return nil
}