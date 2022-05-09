package file

import "flag"

func init() {
	flag.BoolVar(&Options.FieldsFlag, "fields", false, "print file contents as fields")
	flag.BoolVar(&Options.LinesFlag, "lines", false, "print file contents as lines")

	flag.Parse()
}
