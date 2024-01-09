package main

import (
	"fmt"
	"reflect"

	pb "github.com/Clement-Jean/protoc_go_tutorial/proto"
	"google.golang.org/protobuf/proto"
)

func doSimple() *pb.Sample {
	return &pb.Sample{
		Id: 42,
		IsSimple: true,
		Name: "A name",
		SampleLists: []int32{1, 2, 3, 4, 5, 6},
	}
}

func doComplex() *pb.Complex {
	return &pb.Complex{
		OneDummy: &pb.Dummy{Id: 42, Name: "My name"},
		MultipleDummies: []*pb.Dummy{
			{Id: 43, Name: "My name 2"},
			{Id: 44, Name: "My name 3"},
		},
	}
}

func doEnum() *pb.Enumerations {
	return &pb.Enumerations{
		EyeColor: 1, //pb.EyeColor_EYE_COLOR_GREEN,
	}
}

func doOneOf(message interface{}) {
	switch x := message.(type) {
	case *pb.Result_Id:
		fmt.Println(message.(*pb.Result_Id).Id)
	case *pb.Result_Message:
		fmt.Println(message.(*pb.Result_Message).Message)
	default:
		fmt.Errorf("message has unexpected type: %v", x)
	}
}

func doMap() *pb.MapExample {
	return &pb.MapExample{
		Ids: map[string]*pb.IdWrapper{
			"myid": {Id: 42},
			"myid2": {Id: 43},
			"myid3": {Id: 44},
		},
	}
}

func doFile(p proto.Message) {
	path := "simple.bin"

	writeToFile(path, p)
	// message := &pb.Sample{}
	message := &pb.Result{}
	readFromFile(path, message)
	fmt.Println(message)
	// fmt.Println(message.Result)
	// fmt.Println(message.Yamamoto, "111")
}

func doToJSON(p proto.Message) string {
	jsonString := toJSON(p)
	// fmt.Println(jsonString)
	return jsonString
}

func doFromJSON(jsonString string, t reflect.Type) proto.Message {
	message := reflect.New(t).Interface().(proto.Message)
	fromJSON(jsonString, message)
	return message
}

func main() {
	// fmt.Println(doSimple())
	// fmt.Println(doComplex())
	// fmt.Println(doEnum())
	// fmt.Println("This should be an Id:")
	// doOneOf(&pb.Result_Id{Id: 42})
	// fmt.Println("This should be a message")
	// doOneOf(&pb.Result_Message{Message: "a message"})
	// fmt.Println(doMap())

	// //DoFile
	// a := &pb.Result{Result: &pb.Result_Id{Id: 1}, Test: &pb.Result_Test1{Test1: "1"}, Yamamoto: "1"}
	// // typeを固定することでResultのタイプを固定してIdを使えるようにする
	// fmt.Println(a)
	// doFile(&pb.Result{Result: &pb.Result_Id{Id: 1}, Test: &pb.Result_Test1{Test1: "1"}, Yamamoto: 2})
	doFile(&pb.Result2{Result: &pb.Result2_Id{Id: 1}, Test: &pb.Result2_Test1{Test1: "1"}, Yamamoto: 2, Ascii: "sss"})
	// fmt.Println(a.Result.(*pb.Result_Id).Id)
	// // doFile(doSimple())

	
	//from to json
	// jsonString := doToJSON(doSimple())
	// message := doFromJSON(jsonString, reflect.TypeOf(pb.Sample{}))
	// fmt.Println(jsonString)
	// fmt.Println(message)

	// jsonString = doToJSON(doComplex())
	// message = doFromJSON(jsonString, reflect.TypeOf(pb.Complex{}))
	// fmt.Println(jsonString)
	// fmt.Println(message)

	// fmt.Println(doFromJSON(`{"id": 42, "unkown": "test"}`, reflect.TypeOf(pb.Sample{})))
}

