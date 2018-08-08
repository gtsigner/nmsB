package execute

import (
	"../../memory"
	"log"
)

func Reader() error {

	pid := uint(24732)

	reader := memory.NewMemoryReader()
	err := reader.Open(pid)
	if err != nil {
		log.Panicln(err)
	}

	/*v, err := reader.PointerAt("nmsB-windows-amd64.exe", uintptr(0x00175328),uintptr(0x2FD))
	if err != nil {
		log.Panicln(err)
	}

	log.Printf("v: %X\n", v)*/

	baseAddress := reader.ModuleBase("nmsB-windows-amd64.exe")

	bb := baseAddress + uintptr(0x00177328)
	value1, _ := reader.ReadPtr(bb)
	log.Printf("bb: %X, value1: %X\n", bb, value1)

	bb2 := value1 + uintptr(0x90)
	value2, _ := reader.ReadPtr(bb2)
	log.Printf("bb2: %X, value2: %X\n", bb2, value2)

	bb3 := value2 + uintptr(0x40)
	value3, _ := reader.ReadPtr(bb3)
	log.Printf("bb3: %X, value3: %X\n", bb3, value3)

	bb4 := value3 + uintptr(0x70)
	value4, _ := reader.ReadPtr(bb4)
	log.Printf("bb4: %X, value4: %X\n", bb4, value4)

	bb5 := value4 + uintptr(0x88)
	value5, _ := reader.ReadPtr(bb5)
	log.Printf("bb5: %X, value5: %X\n", bb5, value5)

	v, _ := reader.ReadInt32(value5)

	ptr, _ := reader.PointerAt("nmsB-windows-amd64.exe", uintptr(0x00177328), uintptr(0x90), uintptr(0x40), uintptr(0x70), uintptr(0x88))
	v2, _ := reader.ReadInt32(ptr)

	log.Println(v, v2)

	return nil
}
