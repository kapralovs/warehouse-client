package requests

func ConstructURL(host, port, path string) string {
	return host + port + path
}
