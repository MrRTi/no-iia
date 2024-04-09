package cli

import "github.com/hellflame/argparse"

type Opts struct {
	Args     []string
	FilePath string
}

func CliOpts() (*Opts, error) {
	parser := argparse.NewParser("no-iia", "Get all values", &argparse.ParserConfig{DisableDefaultShowHelp: true})

	args := parser.Strings("a", "args", &argparse.Option{
		Positional: true,
		Required:   false,
		Default:    "",
	})

	filePath := parser.String("f", "filepath", &argparse.Option{
		Required: true,
		Default:  "",
	})
	err := parser.Parse(nil)
	if err != nil {
		return nil, err
	}

	return &Opts{
		Args:     *args,
		FilePath: *filePath,
	}, nil
}
