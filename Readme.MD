# Golang Backend for Forbici project

**Strutcture of the backend console :**

        .
        └── ferro13console
            ├── Readme.MD
            ├── Script.sh
            ├── cmd
            │   └── ferro13console
            │       └── main.go
            ├── config
            │   ├── config.go
            │   └── values
            │       └── config.yaml
            ├── controllers
            │   ├── order_controller.go
            │   ├── product_controller.go
            │   └── user_controller.go
            ├── models
            │   ├── order.go
            │   ├── product.go
            │   └── user.go
            ├── repositories
            │   ├── order_repository.go
            │   ├── product_repository.go
            │   └── user_repository.go
            ├── services
            │   ├── order_service.go
            │   ├── product_service.go
            │   └── user_service.go
            └── utils
                ├── http_utils.go
                └── validation_utils.go

## ROADMAP:

- Studio delle funzioni controllers in laravel , capire cosa trattano (dati richiesti al db , api usate ) e replicare le chiamate al db.
- Studio delle strutture delle tabelle db(prodotti,utenti,) e creazione delle tali nel config/ 
- Creazione delle route che richiamino le funzioni di chiamata al db.