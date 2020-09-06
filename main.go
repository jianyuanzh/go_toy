package main

import (
	"awesome-yuan/go_concurrent"
	"awesome-yuan/go_interface"
	"awesome-yuan/go_json"
	"awesome-yuan/go_pointer"
	"awesome-yuan/go_sugar"
	"fmt"
	"reflect"
	"runtime"
	"time"
)

/***
* 使用Make创建对象的时候返回的是引用类型，使用New创建对象的时候返回的是指针
* 使用make的目的是为了让go把这些对象初始化完成，避免使用的时候出错
*/
func main() {

	go_sugar.Sugar("test, ", "hellop", "fwold")
	go_sugar.Sugar2()

	jstr := go_json.SerializeServerStruct()
	go_json.DeserializeServer(jstr)
	jstr = go_json.SerializeMap()
	go_json.DeserializeMap(jstr)

	makeSlice()
	makeMap()

	newMap()

	appendSlice()
	copySlice()
	deleteFromMap()



	lenSlice()


	dog := varDog()
	dog.Run()
	newDg := newDog()
	newDg.Run()
	sDog := dogFromAnon()
	sDog.Run()
	sDog.Eat()

	doBehavior(&sDog)

	cat := Cat{Age: 4, Category: "梨花", Host: "元歌"}
	cat.Color = "灰色"
	doBehavior(&cat)


	fmt.Println("----------")



	go go_concurrent.Loop()
	go go_concurrent.Loop()

	go go_concurrent.Send()

	go go_concurrent.Recv()

	time.Sleep(time.Second * 4)
	// 设置GO的最大核心数
	runtime.GOMAXPROCS(runtime.NumCPU() - 1)
	fmt.Printf("\ncpu num: %d\n", runtime.NumCPU())

	// 携程同步
	go_concurrent.Read()
	go go_concurrent.Write()
	go_concurrent.WG.Wait()

	fmt.Println("All done")

	go_pointer.ShowPointer()

	go_pointer.TestPointerArray()
	defer recoverPanic()
	receivePanic()


}

func doBehavior(b go_interface.Behavior) {
	b.Eat()
	b.Run()
}

func deleteFromMap() {
	mMap := make(map[string]int)
	mMap["test"] = 12
	mMap["value"] = 1244

	fmt.Println(mMap)
	delete(mMap, "test")
	fmt.Println(mMap)

}

func newMap() {
	nMap := new(map[string]int)
	//*nMap["dog"] = 10
	//*nMap["cat"] = 20
	//*nMap["tiger"] = 100
	fmt.Println(nMap)
	fmt.Println(reflect.TypeOf(nMap))
}

func newSlice() {

}



// 尽量使用定长，append会进行扩容
func appendSlice() {
	nSlice := make([]string, 2)
	nSlice[0] ="12"
	fmt.Println(len(nSlice))
	nSlice = append(nSlice, "14")
	fmt.Println(len(nSlice))
	nSlice[1] = "13"
	fmt.Println(len(nSlice))
	fmt.Println(nSlice)
}

func copySlice() {
	source := make([]string ,2)
	dst := make([]string ,2)

	source[0] = "vincent"
	source[1] = "zhang"
	source = append(source, "duplicated")

	copy(dst, source)

	fmt.Println(source)
	fmt.Println(dst)
}


func makeSlice() {
	mSlice := make([]string, 3)
	mSlice[0] = "dog"
	mSlice[1] = "cat"
	mSlice[2] = "tiger"

	fmt.Println(mSlice)
}

func makeMap() {
	mMap := make(map[string]int)
	mMap["dog"] = 10
	mMap["cat"] = 20
	mMap["tiger"] = 100

	fmt.Println(mMap)
	fmt.Println(reflect.TypeOf(mMap))

}

//func makeChannel() {
//	mChan := make(chan int, 3)
//}