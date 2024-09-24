package di

import (
	"errors"
	"sync"
)

// Container armazena as dependências
type Container struct {
	services map[string]interface{}
}

// Instância global do container (singleton)
var instance *Container
var once sync.Once

// GetInstance retorna a instância singleton do container
func GetInstance() *Container {
	once.Do(func() {
		instance = &Container{
			services: make(map[string]interface{}),
		}
	})
	return instance
}

// Register adiciona uma dependência ao container
func (c *Container) Register(name string, service interface{}) {
	c.services[name] = service
}

// Resolve retorna a dependência registrada
func (c *Container) Resolve(name string) (interface{}, error) {
	service, found := c.services[name]
	if !found {
		return nil, errors.New("serviço não encontrado no container: " + name)
	}
	return service, nil
}


/*** Exemplo de uso

// ------------- Exemplo de registro da dependência no método MAIN, por exemplo.
// Obter o container singleton de DI
container := di.GetInstance()

// Registrar o DB, por exemplo, no container de DI
container.Register("db", &db)

// ------------- Exemplo de uso depois de registrado
// Obter o container singleton de DI
container := di.GetInstance()

// Resolver o serviço de DB para fazer a DI
database, err := container.Resolve("db")
if err != nil {
	fmt.Println("Erro ao resolver o DI db:", err)
	return
}

// Convertendo de interface{} para *sql.DB para viabilizar o uso do db
db := database.(**sql.DB)

*/