package cli

import (
	"os"
)

// GetFileList returns a list of files in the specified path
// that match any of the specified patterns while excluding
// those that match the specified exclude list.
func GetFileList(path string, pattern, exclude []string, recurse bool) ([]string, error) {
	return nil, errNotImplemented("GetFileList()")
}

// StatCli returns the os.FileInfo from filename.
// In the Cli version, any error results in log.Fatal().
func StatCli(filename string) os.FileInfo {
	fi, err := os.Stat(filename)
	if err != nil {
		Log.Fatal(osErr(err, "StatCli()"))
	}
	return fi
}

// GetDataCli gets the contents of filename and returns
// the string version.
// In the Cli version, any error results in log.Fatal().
func ReadFileCli(filename string) string {
	data, err := os.ReadFile(filename)
	if err != nil {
		Log.Fatal(osErr(err, "ReadFileCli()"))
	}

	return B2S(data)
}
