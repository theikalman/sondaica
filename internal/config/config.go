package config

type Config struct {
	RestAPI RestAPI `mapstructure:"rest_api"`
}

type Handler struct {
	OperationId string `mapstructure:"operationId"`
	HanlderName string `mapstructure:"handlerName"`
}

type Handlers []Handler

type RestAPI struct {
	OpenAPISpec     string   `mapstructure:"open_api_spec"`
	OutputServerDir string   `mapstructure:"output_server_dir"`
	OutputClientDir string   `mapstructure:"output_client_dir"`
	Handlers        Handlers `mapstructure:"handlers"`
}
