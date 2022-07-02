package main

import (
	"fmt"
	"reflect"
)

type order struct {
	ordId      int
	customerId int
	name       string
}

func createQuery1(q interface{}) {
	t := reflect.TypeOf(q)
	k := t.Kind()
	v := reflect.ValueOf(q)
	fmt.Println("Type ", t)
	fmt.Println("Value ", v)
	fmt.Println("Kind ", k)

}

func createQuery(q interface{}) {
	if reflect.ValueOf(q).Kind() == reflect.Struct {
		t := reflect.TypeOf(q).Name()
		query := fmt.Sprintf("insert into %s values(", t)
		v := reflect.ValueOf(q)
		fmt.Println(v.NumField())
		for i := 0; i < v.NumField(); i++ {
			switch v.Field(i).Kind() {
			case reflect.Int:
				if i == 0 {
					query = fmt.Sprintf("%s%d", query, v.Field(i).Int())
				} else {
					query = fmt.Sprintf("%s, %d", query, v.Field(i).Int())
				}
			case reflect.String:
				if i == 0 {
					query = fmt.Sprintf("%s\"%s\"", query, v.Field(i).String())
				} else {
					query = fmt.Sprintf("%s, \"%s\"", query, v.Field(i).String())
				}
			default:
				fmt.Println("Unsupported type")
				return
			}
		}
		query = fmt.Sprintf("%s)", query)
		fmt.Println(query)
		return

	}
	fmt.Println("unsupported type")
}

func main() {
	// o := order{
	// 	ordId:      456,
	// 	customerId: 56,
	// 	name:       "ibnu",
	// }
	// createQuery(o)

	a := 1
	b := "hallo"
	c := order{
		ordId:      456,
		customerId: 56,
		name:       "ibnu",
	}

	at := reflect.TypeOf(a)
	bt := reflect.TypeOf(b)
	ct := reflect.TypeOf(c)

	fmt.Println("at.Name() : ", at.Name())
	fmt.Println("bt.Name() : ", bt.Name())
	fmt.Println("ct.Name() : ", ct.Name())
	fmt.Println("------------------------------------------------------------------")
	fmt.Println("at.Kind() : ", at.Kind())
	fmt.Println("bt.Kind() : ", bt.Kind())
	fmt.Println("ct.Kind() : ", ct.Kind())

	fmt.Println("------------------------------------------------------------------")
	fmt.Println("at.Kind() == reflect.Int : ", at.Kind() == reflect.Int)
	fmt.Println("bt.Kind() == reflect.String : ", bt.Kind() == reflect.String)
	fmt.Println("ct.Kind() == reflect.Struct : ", ct.Kind() == reflect.Struct)

	fmt.Println("valueof-------------------------------------------------------------------")

	dt := reflect.ValueOf(c)
	for i := 0; i < dt.NumField(); i++ {
		switch dt.Field(i).Kind() {
		case reflect.Int:
			fmt.Println("", i, dt.Field(i).Kind(), dt.Field(i).Int())
		case reflect.String:
			fmt.Println("", i, dt.Field(i).Kind(), dt.Field(i).String())
		default:
			fmt.Println("Unsupported type")
			return
		}
	}

	fmt.Println("-------------------------------------------------------------------")
	var e int
	et := reflect.TypeOf(&e)
	fmt.Println("dt.Name() :", et.Name())              // returns an empty string
	fmt.Println("et.Kind() :", et.Kind())              // returns reflect.Ptr
	fmt.Println("et.Elem().Name():", et.Elem().Name()) // returns "int"
	fmt.Println("et.Elem().Kind():", et.Elem().Kind()) // returns reflect.Int

}
