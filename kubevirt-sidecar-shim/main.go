package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"os"

	"github.com/spf13/pflag"
	v1 "kubevirt.io/api/core/v1"
	"kubevirt.io/kubevirt/pkg/virt-launcher/virtwrap/api"
)

func onDefineDomain(vmiJSON, domainXML []byte) string {
	vmiSpec := v1.VirtualMachineInstance{}
	if err := json.Unmarshal(vmiJSON, &vmiSpec); err != nil {
		panic(err)
	}

	domainSpec := api.DomainSpec{}
	if err := xml.Unmarshal(domainXML, &domainSpec); err != nil {
		panic(err)
	}

	// Find and update existing product and family entries
	for i, entry := range domainSpec.SysInfo.System {
		switch entry.Name {
		case "product":
			domainSpec.SysInfo.System[i].Value = "KVM"
		case "family":
			domainSpec.SysInfo.System[i].Value = "Virtual Machine"
		}
	}

	if newDomainXML, err := xml.Marshal(domainSpec); err != nil {
		panic(err)
	} else {
		return string(newDomainXML)
	}
}

func main() {
	var vmiJSON, domainXML string
	pflag.StringVar(&vmiJSON, "vmi", "", "VMI to change in JSON format")
	pflag.StringVar(&domainXML, "domain", "", "Domain spec in XML format")
	pflag.Parse()

	if vmiJSON == "" || domainXML == "" {
		os.Exit(1)
	}
	fmt.Println(onDefineDomain([]byte(vmiJSON), []byte(domainXML)))
}
