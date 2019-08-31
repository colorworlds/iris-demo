package config

type params struct {
	LogPath string
}

func newParams() *params {
	return &params{
		LogPath: "runtime/logs/iris_web.%Y%m%d.log",
	}
}