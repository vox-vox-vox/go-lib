package main

import (
	"fmt"
	"reflect"
)

func test(s reflect.Value){
	pf := fmt.Printf
    pf("KindType:%T, KindValue: %#v \n", s.Kind(), s.Kind())
    pf("TypeType:%T, TypeValue: %#v\n", s.Type(), s.Type())
    pf("NumFiled:%v \n", s.NumField())
    pf("Field0:%v, T:%T \n", s.Field(0), s.Field(0))
    pf("FieldTypeName:%v, T:%T \n", s.Field(0).Type().Name(), s.Field(0).Type().Name())
    pf("FieldByName:%v, T:%T \n", s.FieldByName("Age"), s.FieldByName("Age"))
    pf("---------------------\n")
}

func main() {
	//pf := fmt.Printf
    type B struct{
        Age int
        Name string
    }
    s := B{Age:100}
    s1:=&s
    rv := reflect.ValueOf(*s1)
    test(rv)
}
