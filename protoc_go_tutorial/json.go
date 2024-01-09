package main

import (
	"log"

	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

func toJSON(pb proto.Message) string {
	//マルチラインでjsonを表示するオプション
	option := protojson.MarshalOptions{
		Multiline: true,
	}
	out, err := option.Marshal(pb)

	if err != nil {
		log.Fatalln("Can't convert to JSON", err)
		return ""
	}

	return string(out)
}

//go to goの通信と違いjson to go ではdefault値は適応されない
func fromJSON(in string, pb proto.Message) {
	//第二引数の構造体のkeyに存在しない値は入れないオプション
	//このオプションがなかったらエラーをはく
	option := protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}
	if err := option.Unmarshal([]byte(in), pb); err != nil {
		log.Fatalln("Couldn't unarshal from JSON", err)
	}
}