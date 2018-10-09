package signature

import (
	"../modules"
)

func ScanModule(signature string, moduleName string) (uintptr, error) {
	// get info about the module
	mod, err := modules.GetProcessModule(moduleName)
	if err != nil {
		return 0, err
	}

	// make the pattern for the signature
	p, err := FromString(signature)
	if err != nil {
		return 0, err
	}

	length := int64(mod.Size)
	start := uintptr(mod.Handle)
	// scan for the signature
	offset := Scan(start, length, p)

	if offset < 0 {
		return 0, nil
	}

	// calculate pointer
	ptr := start + uintptr(offset)
	return ptr, nil
}
