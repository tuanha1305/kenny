package config

//Namespace is the name for application instance
const Namespace = "Kenny"

//nolint:lll,gochecknoglobals,gomnd
var def = Config{
	Debug: true,

	Logger: Logger{
		Level:   5,
		Enabled: true,
	},

	Recorder: Recorder{
		NumberOfChannels: 1,
		SampleRate:       44100,
		FramesPerBuffer:  64,
	},
}
