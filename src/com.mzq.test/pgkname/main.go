package main

import (
	"bytes"
	"com.mzq.test/user"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"reflect"
	"time"
)

type Sender interface {
	Send(to, mes string)  (string , error )
}

type DSender struct {
	sss string
}

func (dSender DSender)Send(to, mes string) (string , error ) {
	println(1111)
	return "2222222", nil
}

type Receiver interface {
	Receiver() error
}

type Connection interface {
	Sender
	Receiver
	Open() error
	Close() error
}

var ljh2 interface {
	Send(to, mes string)  (string , error )
}


func test() {
	defer func() {
		fmt.Println(recover())
	}()

	defer func() {
		panic("defer panic")
	}()

	panic("test panic")
}
func foo(a, b int) (i int, err error) {
	defer func() { fmt.Printf("second defer err %v\n", err) }()
	defer fmt.Printf("first defer err %v\n", err)

	defer func() { fmt.Printf("third defer err %v\n", err) }()

	fmt.Printf("third 333333333333333 err %v\n", err)
	if b == 0 {
		err = errors.New("divided by zero!")

		return 
	}

	i = a / b
	return
}



func main() {

	_, err2 := foo(2, 0)
	fmt.Printf("third 222222222222defer err %v\n", err2)
	//user.TestRunTime()

	//cobra.Execute();

	//cmd.TestCobra()
	//
	return
	
	var ssender Sender = DSender{sss: "22sss"}
	if sender04, err := ssender.(Sender); err{

		fmt.Printf("%T, %#v, %v\n", sender04,sender04,sender04)
		send, ccc := ssender.Send("111", "222")
		fmt.Printf("%T, %#v, %v, %#v, %T, %T\n", sender04,sender04,sender04,send, ccc, nil)
	}

	ljh2 = ssender
	ljh2.Send("|", "2")



	switch v := ssender.(type) {
	case Sender:
		fmt.Println(v)
	case DSender:
		fmt.Println(v.sss)
	}
	//fmt.Printf("%T, %#v, %v,%#v", ssender,ssender,ssender,ssender )

	ssss := "sssss";
	weight := 10;
	newUser := user.NewUser(1, "aaaa", "130", 180, &ssss, &weight)
	fmt.Println(newUser.String())

	(*newUser).SetId(2);
	fmt.Println(newUser.String())

	newUser.SetName("fffffff")
	fmt.Println(newUser.String())


	fmt.Println("------------------------------")

	newUser2 := user.NewUser2(1, "aaaa", "130", 180, &ssss, &weight)
	fmt.Println(newUser2.String())

	newUser2.SetId(3333);
	fmt.Println(newUser2.String())

	newUser2.SetName("33333")
	fmt.Println(newUser2.String())

	fmt.Println(newUser2)

	vars := make([]interface{}, 0, 20)
	var intV int = 1;
	var floatV float32 = 3.14
	var boolV bool = true
	var stringV string = "11111111111111111111"
	var arrayV [5]int =[...]int{1, 2, 3, 4, 5}
	var sliceV []int = make([]int, 3, 5)
	var funs []func() = make([]func(), 0)
	funs = append(funs, func() {
		fmt.Println(111)
	}, func() {
		fmt.Println(222222222)
	}, func() {
		fmt.Println(3333333333)
	})


	var mapV map[string]string = map[string]string{"name":"kk"}
	var funcV1 func(...interface{}) error = func(x ...interface{}) error {
		fmt.Println(x...);
		return nil
	}



	var funcV2 func(string, int) *user.Connection = user.NewConnection
	s := "1111111"
	sh := 50
	var userV *user.User = user.NewUser(1, "KKKK", "150",
		50, &s, &sh)

	var closeV user.Closer
	vars = append(vars, intV, &intV, floatV, boolV, stringV, arrayV, sliceV, mapV,funs,
		funcV1, funcV2, *userV, userV, closeV)
	//for _,v := range vars{
	//	fmt.Println("===========================")
	//	fmt.Println(v)
	//	//user.DisplayType(reflect.TypeOf(v), "")
	//	user.DisplayValue(reflect.ValueOf(v), "")
	//}

	//sum := func(parameters []reflect.Value)[]reflect.Value {
	//	var total int64
	//	for _,parameter := range parameters[:len(parameters) -1]{
	//		total += parameter.Int()
	//	}
	//
	//	last := parameters[len(parameters) -1]
	//	switch last.Kind() {
	//	case reflect.Int:
	//		total += last.Int();
	//	case reflect.Slice:
	//		for i := 0; i < last.Len(); i++ {
	//			total += last.Index(i).Int()
	//		}
	//
	//	}
	//	return []reflect.Value{reflect.ValueOf(total)}
	//}
	//
	//var add2 func(int, int) int64
	//
	//ref := reflect.ValueOf(&add2).Elem()
	//
	//makeFunc := reflect.MakeFunc(ref.Type(), sum)
	//
	//ref.Set(makeFunc)
	//
	//fmt.Println(add2(1, 2))
	//
	//var add3 func(int, int, int) int64
	//var add4 func(int, int, int, int) int64
	//makeFunc2 := func(fn interface{}) {
	//	fmt.Printf("111111111 %T\n", fn)
	//	fmt.Printf("不管 fn是函数实例还是指针， 这里都是 *interface {}  %T\n", &fn)
	//	ref := reflect.ValueOf(fn).Elem()
	//	makeFunc := reflect.MakeFunc(ref.Type(), sum)
	//	ref.Set(makeFunc)
	//}
	////makeFunc2(add3)
	//makeFunc2(&add3)
	//makeFunc2(&add4)
	//fmt.Println(add3(1, 2, 3))
	//fmt.Println(add4(1, 2, 4, 5))

	users_slice := make([]user.User2, 4, 20)
	users_map := make(map[int]user.User2)

	for i :=0; i<4; i++ {
		u := user.User2{
			Id: i,
			Name: "user." + string(i),
			Online: i%3 == 0,
			Register: time.Now(),
		}
		users_slice[i] = u
		users_map[i] = u
	}

	//jsonMarshal(users_slice)
	//var ljhg interface{}

	encoder := json.NewEncoder(user.User2{})
	encoder.SetIndent("", "\t")

	err := encoder.Encode(users_slice)
	if err != nil {
		fmt.Println(err)
	}


	//var out bytes.Buffer
	//newEncoder := json.NewEncoder(&out)
	//err = newEncoder.Encode(users_map)
	//if err == nil {
	//	fmt.Println(out.String())
	//}else {
	//	fmt.Println(err)
	//}



}

func jsonMarshalIndent(users_map map[int]user.User2)  {
	indent, err := json.MarshalIndent(users_map, "", "\t")
	if err == nil {
		fmt.Println(string(indent))
	}
}





func jsonMarshal(users_slice []user.User2)  {
	b, err := json.Marshal(users_slice)
	if err == nil {
		fmt.Println(reflect.TypeOf(b))
		fmt.Println(string(b))
		var buffer bytes.Buffer
		json.Indent(&buffer, b, "", "\t")
		//fmt.Println(buffer)
		buffer.WriteTo(os.Stdout)
	}



	var users []user.User2
	err = json.Unmarshal(b, &users)
	if err == nil {
		fmt.Printf("%T, %v", users, users)
	}
}

func reflectCreateStruct()  {
	fmt.Printf("-------------------")

	fields := []reflect.StructField{
		{
			Name: "Name",
			Type: reflect.TypeOf(""),
		},
		{
			Name: "Score",
			Type: reflect.TypeOf(int64(0)),
		},
	}
	fmt.Printf("111111111 %T\n", fields)

	UserType := reflect.StructOf(fields)
	fmt.Printf("22222222222 %T\n", UserType)

	user := reflect.New(UserType).Elem()

	fmt.Printf("33333333333 %T\n", reflect.New(UserType))
	fmt.Printf("4444444444 %T\n", user)



	user.Field(0).Set(reflect.ValueOf("sssssss"))
	user.Field(1).Set(reflect.ValueOf(int64(1111)))

	fmt.Printf("55555555555    %#v\n", user.Interface())
	fmt.Printf("6666666666 %#v\n", user.Addr().Interface())
}

