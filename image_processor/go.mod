module image_processor

go 1.23.0

toolchain go1.24.2

require (
	github.com/disintegration/imaging v1.6.2
	github.com/streadway/amqp v1.1.0
	pkg v0.0.0-00010101000000-000000000000
	user_backend v0.0.0-00010101000000-000000000000
)

require (
	github.com/BurntSushi/toml v1.2.1 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/ilyakaznacheev/cleanenv v1.5.0 // indirect
	github.com/joho/godotenv v1.5.1 // indirect
	golang.org/x/image v0.0.0-20191009234506-e7c1f5e7dbb8 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
	olympos.io/encoding/edn v0.0.0-20201019073823-d3554ca0b0a3 // indirect
)

replace user_backend => ../user_backend

replace pkg => ../pkg
