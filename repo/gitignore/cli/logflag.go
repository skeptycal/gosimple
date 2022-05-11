package cli

// Set sets the logging level from a string value.
// Allowed values: Panic, Fatal, Error, Warn, Info, Debug, Trace
func (l *logLevelFlag) Set(lvl string) error {
	err := Log.SetLogLevel(lvl)
	if err != nil {
		return err
	}
	l.level = Log.GetLevel()
	return nil
}

func (l logLevelFlag) Get() any       { return Log.GetLevel() }
func (l logLevelFlag) String() string { return l.level.String() }
