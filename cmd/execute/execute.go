package execute

import (
	"../../cmd"
)

func Execute(args *cmd.Arguments) error {
	if *args.Pointer {
		err := Pointer(args.ProcessId, args.Address)
		return err
	}
	
	return nil
}
