package main

import (
	"C"
	"fmt"
	"log"

	"github.com/openconfig/ygot/proto/ywrapper"
	vsApi "github.com/sriramy/vswitch-api/pkg/gen/go/managed_element"
	"google.golang.org/protobuf/encoding/prototext"
	"google.golang.org/protobuf/proto"
)

//export SetHostName
func SetHostName(hostname string) {
	host := &vsApi.System{Hostname: &ywrapper.StringValue{Value: hostname}}
	me := &vsApi.ManagedElement{System: host}

	b, err := proto.Marshal(me)
	if err != nil {
		log.Panicf("Error marshalling proto: %v", err)
	}

	fmt.Printf("%s\n", prototext.Format(me))
	fmt.Printf("Marshalled proto size in bytes: %d\n", len(b))
}

func main() {}
