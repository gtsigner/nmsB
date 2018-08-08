package execute

import (
	"../../cmd"
)

func Execute(args *cmd.Arguments) error {
	if *args.Pointer {
		err := Pointer(args.ProcessId, args.Address)
		return err
	} else if *args.Reader {
		err := Reader()
		return err
	}

	return nil
}