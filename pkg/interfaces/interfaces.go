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
import "github.com/sriramy/vswitch-api/pkg/gen/go/managed_element/enums"

//export SetInterface
func SetInterface(ifName string, vlan uint64) {
	subinterfaces := make([]*vsApi.Interface_SubinterfaceKey, 0)
	subinterfaces = append(subinterfaces, &vsApi.Interface_SubinterfaceKey{
		Index: 0,
		Subinterface: &vsApi.Interface_Subinterface{
			AdminStatus: vsApi.Interface_Subinterface_ADMINSTATUS_UP,
			Name:        &ywrapper.StringValue{Value: ifName},
			Vlan: &vsApi.Interface_Subinterface_Vlan{
				VlanId: &vsApi.Interface_Subinterface_Vlan_VlanIdUint64{VlanIdUint64: vlan},
			},
		},
	})
	interfaces := make([]*vsApi.ManagedElement_InterfaceKey, 0)
	interfaces = append(interfaces, &vsApi.ManagedElement_InterfaceKey{
		Name: ifName,
		Interface: &vsApi.Interface{
			AdminStatus:  vsApi.Interface_ADMINSTATUS_UP,
			Enabled:      &ywrapper.BoolValue{Value: true},
			Subinterface: subinterfaces,
			Type:         enums.IETFInterfacesInterfaceType_IETFINTERFACESINTERFACETYPE_gigabitEthernet,
		},
	})

	me := &vsApi.ManagedElement{Interface: interfaces}

	b, err := proto.Marshal(me)
	if err != nil {
		log.Panicf("Error marshalling proto: %v", err)
	}

	fmt.Printf("%s\n", prototext.Format(me))
	fmt.Printf("Marshalled proto size in bytes: %d\n", len(b))
}

func main() {}
