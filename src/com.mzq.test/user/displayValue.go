package user

import (
	"fmt"
	"math/rand"
	"reflect"
	"strconv"
)


func DisplayValue(t reflect.Value, tab string)  {
	//if t == nil{
	//	fmt.Printf("<nil>")
	//	return
	//}
	switch t.Kind() {
	case reflect.Uint:
		if t.CanSet() {
			fmt.Printf("可修改 %s %d, \n",   t.Type(), t.Uint())
			t.SetUint(t.Uint() + rand.Uint64())
		}else {
			fmt.Printf("不可修改 %s %d, \n",   t.Type(), t.Uint())
		}

		fmt.Printf("<nilnilnilnilnilnilnilnilnilnilnilnilnil>\n")
	case reflect.Int:
		if t.CanSet() {
			fmt.Printf("可修改 %s %d, \n",   t.Type(), t.Int())
			t.SetInt(t.Int() + rand.Int63n(1444))
		}else {
			fmt.Printf("不可修改 %s %d, \n",   t.Type(), t.Int())
		}

		fmt.Printf("<nilnilnilnilnilnilnilnilnilnilnilnilnil>\n")
	case reflect.String:
		if t.CanSet() {
			fmt.Printf("可修改 %s %s \n",  t.Type(), t)
			t.SetString("chanage:" + t.String())
		}else{
			fmt.Printf("不可修改 %s %s \n",  t.Type(), t)
		}

		fmt.Printf("<nilnilnilnilnilnilnilnilnilnilnilnilnil>\n")
	case reflect.Bool:
		if t.CanSet() {
			fmt.Printf("可修改 %s  %s \n",
				t.Type(),  strconv.FormatBool(t.Bool()))
			t.SetBool(!t.Bool())
		}else{
			fmt.Printf("不可修改 %s  %s \n",
				t.Type(),  strconv.FormatBool(t.Bool()))
		}


		fmt.Printf("<nilnilnilnilnilnilnilnilnilnilnilnilnil>\n")
	case reflect.Float32, reflect.Float64:


		if t.CanSet() {
			fmt.Printf("可修改 %s, %f \n", t.Type(), t.Float())
			t.SetFloat(t.Float()+rand.Float64())
		}else{
			fmt.Printf("不可修改 %s, %f \n", t.Type(), t.Float())
		}
		fmt.Printf("<nilnilnilnilnilnilnilnilnilnilnilnilnil>\n")
	case reflect.Complex64, reflect.Complex128:
		fmt.Printf("jibenshujuleixing %T, %v \n",  t.Complex(),  t.Complex())
		fmt.Printf("<nilnilnilnilnilnilnilnilnilnilnilnilnil>\n")
	case reflect.Array:
		if t.CanSet() {
			fmt.Printf("可修改 %s, %d, %#v \n", t.Type(), t.Len(), t)
		}else{
			fmt.Printf("不可修改 %s, %d, %#v \n", t.Type(), t.Len(), t)
		}
		for i:=0; i<t.Len();i++ {
			DisplayValue(t.Index(i), tab)
			fmt.Printf("\n")
		}
		fmt.Printf("<nilnilnilnilnilnilnilnilnilnilnilnilnil>\n")
	case reflect.Slice:
		if t.CanSet() {
			fmt.Printf("可修改 %s,, %d, %d\n", t.Type(), t.Len(), t.Cap())
		}else{
			fmt.Printf("不可修改 %s,, %d, %d\n", t.Type(), t.Len(), t.Cap())
		}

		for i:=0; i<t.Len();i++ {
			DisplayValue(t.Index(i), tab)
			fmt.Printf("\n")
		}
		fmt.Printf("<nilnilnilnilnilnilnilnilnilnilnilnilnil>\n")
	case reflect.Map:
		mapRange := t.MapRange()
		fmt.Printf("[Map]%v, %d, %v\n", t.IsNil(), t.Len(),
			t.MapKeys())
		for mapRange.Next() {
			DisplayValue(mapRange.Key(), tab)
			fmt.Printf(" : ")
			DisplayValue(mapRange.Value(), tab)
			fmt.Printf("\n")

			t.SetMapIndex(mapRange.Key(), reflect.ValueOf(mapRange.Value().String() + " change: "))
		}
		fmt.Printf("<nilnilnilnilnilnilnilnilnilnilnilnilnil>\n")


	case reflect.Func:
		fmt.Printf("%s[Func]\n", tab)
		fmt.Printf("%s, %d, %d:  \n",t.Type(), t.IsValid(), t.IsZero())
		fmt.Printf("%s{  \n", tab)
		fmt.Printf("%s}  \n", tab)
		callFunc(t)
		fmt.Printf("<nilnilnilnilnilnilnilnilnilnilnilnilnil>\n")
	case reflect.Struct:
		fmt.Printf("%s[Struct]\n", tab)
		fmt.Printf("%s, %d:  \n", tab,  t.NumField())
		fmt.Printf("%s{  \n", tab)
		t2 := t.Type()
		for i :=0; i < t2.NumField(); i++{
			fmt.Printf("%s: %v \n", t2.Field(i).Name,  t.Field(i))
			DisplayValue(t.Field(i), tab)

		}
		for i :=0; i < t.NumMethod(); i++{
			callMethod(t.Method(i), t.Type(), t.Type().Method(i))

		}
		fmt.Printf("%s}  \n", tab)
		fmt.Printf("<nilnilnilnilnilnilnilnilnilnilnilnilnil>\n")
	case reflect.Ptr:
		fmt.Printf("%s[Ptr]\n", tab)
		fmt.Printf("{ %s:  ",  t.Elem())
		fmt.Printf("%s}\n  ", tab)
		DisplayValue(t.Elem(), tab)
		for i :=0; i < t.NumMethod(); i++{
			callMethod(t.Method(i), t.Elem().Type(), t.Type().Method(i))

		}
		fmt.Printf("<nilnilnilnilnilnilnilnilnilnilnilnilnil>\n")
	default:
		fmt.Printf("[%s]Unkonw[%s]", tab, t)
	}
}

func callFunc(f reflect.Value)  {
	t := f.Type()
	fmt.Printf("&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&& %s, %v\n", t.String(), f)

	values := make([]reflect.Value, 0)
	for i := 0; i < t.NumIn(); i++ {
		values = append(values, reflect.Zero(t.In(i)))
	}
	if t.IsVariadic() {
		fmt.Printf("&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&& ----------1")
		fmt.Printf("%#v", f.CallSlice(values))
	}else {
		fmt.Printf("&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&& ----------2")
		fmt.Println(f.Call(values))
	}


}

func callMethod(value reflect.Value, p reflect.Type, method reflect.Method)  {
	ftype := value.Type()
	fmt.Printf("^^^^^^^^^^^^^^^^^^^^^^ " +
		"method: %s, %s => %s\n", p.Name(), method.Name, ftype.String())

	values := make([]reflect.Value, 0)
	for i := 0; i < ftype.NumIn(); i++ {
		elem := reflect.New(ftype.In(i)).Elem()
		fmt.Printf("%T, %v", elem, elem)
		values = append(values,elem)
	}
	if ftype.IsVariadic() {
		fmt.Printf("&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&& ----------1")
		fmt.Printf("%#v", value.CallSlice(values))
	}else {
		fmt.Printf("&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&& ----------2")
		fmt.Println(value.Call(values))
	}
}