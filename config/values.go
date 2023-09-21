package config

func TestLoad() bool {
	return vip.IsSet("test_load")
}
