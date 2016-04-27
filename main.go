package main

import "github.com/hyperhq/runv/hypervisor"
import "github.com/hyperhq/runv/driverloader"
import "github.com/hyperhq/runv/hypervisor/types"
import "fmt"
func  main() {
	var (
		err error
		cpu = 1
		mem = 128
	)

	driver := "qemu"
	if hypervisor.HDriver, err = driverloader.Probe(driver); err != nil {
		fmt.Println("%s\n", err.Error())
	}


	b := &hypervisor.BootConfig{
		Kernel: "vmlinuz",
		Initrd: "initrd",
		Bios:   "",
		Cbfs:   "",
		Vbox:   "",
		CPU:    cpu,
		Memory: mem,
	}

	vm := hypervisor.NewVm("vm-123456", cpu, mem, false)
	fmt.Println("gao jianqiao")
	err = vm.Launch(b)
	if err != nil {
		fmt.Printf("%s\n", err.Error())
	}
	
	fmt.Printf("gao jianqiao")
	Status, err := vm.GetResponseChan()
	if err != nil {
		fmt.Printf("gao jianqiao15")
		vm.Kill()
		return 
	}
	
	defer func() {
		
		vm.ReleaseResponseChan(Status)
		vm.Kill()
	} ()
	for {
		fmt.Printf("gao j2ianqiao")
		vmResponse, ok := <-Status
		
		if !ok || vmResponse.Code == types.E_VM_RUNNING {
			break
		}
	}
	
}
