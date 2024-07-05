1. `Compile Daemon`

- Deşiklikleri algılayıp restart atar

```shell
go get github.com/githubnemo/CompileDaemon
```

<br>

```shell
CompileDaemon -command="./go-http -my-options"
```

- `./go-http` dosya yolu
- `-my-options` alabileceği özel parametreler

<br>

2. `GoDotEnv`

- `.env` dosyasından değer okumayı sağlar

```shell
go get github.com/joho/godotenv
```

<br>

3. `Gin`

- Web framework'üdür

```shell
go get -u github.com/gin-gonic/gin
```

<br>

4. `Gorm`

- ORM

```shell
go get -u gorm.io/gorm
```

`-u` update -> modülün en son sürümünü getir demek

- Driver

```shell
go get -u gorm.io/driver/postgres
```

---

`migrate`

- Database tablosu oluşturur onun için projeden önce run edilmeli ve `main` paketi altında run edilebilir olmalıdır

```shell
go run migrate/migrate.go
```

---

`Swagger`

```shell
go get -u github.com/swaggo/swag/cmd/swag
```

```shell
go install github.com/swaggo/swag/cmd/swag@latest
```

```shell
swaggo -h
```

```shell
go get -u github.com/swaggo/files
```

```shell
go get -u github.com/swaggo/gin-swagger
```

```go
// Add Swagger
r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
```

```shell
swag init
```

`main.go`
```go
import (
	// Aşağıdakini ekle
	_ "go-http/docs"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go-http/controllers"
	"go-http/initializers"
)

func init() {
initializers.LoadEnvVeriables()
initializers.ConnectToDB()
}

// @title Tag Service API
// @version 1.0
// @description Description

// @host localhost:8888
// @BasePath /api
func main() {
```

```shell
swag init
```