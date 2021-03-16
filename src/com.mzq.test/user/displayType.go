package user

import (
	"fmt"
	"reflect"
)

func DisplayType(t reflect.Type, tab string)  {
	if t ==nil{
		fmt.Printf("<nil>")
		return
	}
	switch t.Kind() {
	case reflect.Uint, reflect.Int, reflect.Bool, reflect.String:
		fmt.Printf("jibenshujuleixing %s%s\n", tab, t.Name())
		displayAgentType(t)
		fmt.Printf("<nilnilnilnilnilnilnilnilnilnilnilnilnil>\n")
	case reflect.Array,reflect.Slice:
		fmt.Printf("[array, slice]%s%s\n", tab, t)
		displayAgentType(t)
		fmt.Printf("<nilnilnilnilnilnilnilnilnilnilnilnilnil>\n")
	case reflect.Map:
		fmt.Printf("%s[map]{\n", tab)
		fmt.Printf("%s\tKey:  ", tab)
		fmt.Printf("%s%s:  ", tab, t.Key())
		fmt.Println()
		fmt.Printf("%s\tValue:  ", tab)
		fmt.Printf("%s%s:  ", tab, t.Elem())
		fmt.Println()
		fmt.Printf("%s}  \n", tab)
		displayAgentType(t)
		fmt.Printf("<nilnilnilnilnilnilnilnilnilnilnilnilnil>\n")
	case reflect.Func:
		fmt.Printf("%s[Func]\n", tab)
		fmt.Printf("%s, %d, %d:  \n",  t.IsVariadic(), t.NumIn(), t.NumOut())
		fmt.Printf("%s{  \n", tab)
		for i :=0; i < t.NumIn(); i++{
			fmt.Printf("%s, %s:  \n",i, t.In(i))
		}
		fmt.Printf("%s} \n ", tab)

		fmt.Printf("%s{  \n", tab)
		for i :=0; i < t.NumOut(); i++{
			fmt.Printf("%s, %s:  \n", i,  t.Out(i))
		}
		fmt.Printf("%s} \n ", tab)

		displayAgentType(t)
		fmt.Printf("<nilnilnilnilnilnilnilnilnilnilnilnilnil>\n")
	case reflect.Struct:
		fmt.Printf("%s[Struct]\n", tab)
		fmt.Printf("%s, %d:  \n", tab,  t.NumField())
		fmt.Printf("%s{  \n", tab)
		for i :=0; i < t.NumField(); i++{
			fmt.Printf("%s, %s, %t, %v:  \n", i,  t.Field(i).Name,
				t.Field(i).Anonymous, t.Field(i).Tag)
		}
		fmt.Printf("%s}  \n", tab)
		displayAgentType(t)
		fmt.Printf("<nilnilnilnilnilnilnilnilnilnilnilnilnil>\n")
	case reflect.Ptr:
		fmt.Printf("%s[Ptr]\n", tab)
		fmt.Printf("%s, %s:  \n", tab,  t.Elem())
		displayAgentType(t)
		DisplayType(t.Elem(), tab)
		fmt.Printf("<nilnilnilnilnilnilnilnilnilnilnilnilnil>\n")
	default:
		fmt.Printf("[%s]Unkonw[%s]\n", tab, t)
		fmt.Printf("<nilnilnilnilnilnilnilnilnilnilnilnilnil>\n")
	}
}

func displayAgentType(t reflect.Type)  {
	fmt.Printf("-------------\n")

	fmt.Printf("%s, %s, %s, %s, %s, %d:  \n",
		t.Name(), t.PkgPath(), t.Kind(), t.String(), t.Comparable(), t.NumMethod())
	//Implements(Type): 是否实现某接口
	//⚫ AssignableTo(Type): 是否可赋值给某类型
	//⚫ ConvertibleTo(Type): 是否可转换为某类型
	fmt.Printf("*************\n")
	for i :=0; i < t.NumMethod(); i++{
		fmt.Printf("%s, %s, %s:  \n",
			t.Method(i).Name, t.Method(i).Type, t.Method(i).Func)
	}

	fmt.Printf("-------------\n")
}
