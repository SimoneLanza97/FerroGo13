#!/bin/bash

########################################################
# SCRIPT PER LA CREAZIONE DELLA STRUTTURA DEL PROGETTO #
########################################################

# Nome del progetto
project_name="ferro13console"

# Creazione delle cartelle
mkdir -p $project_name/cmd/$project_name
mkdir -p $project_name/controllers
mkdir -p $project_name/models
mkdir -p $project_name/services
mkdir -p $project_name/repositories
mkdir -p $project_name/utils
mkdir -p $project_name/config

# Creazione dei file principali
touch $project_name/cmd/$project_name/main.go
touch $project_name/controllers/user_controller.go
touch $project_name/controllers/product_controller.go
touch $project_name/controllers/order_controller.go
touch $project_name/models/user.go
touch $project_name/models/product.go
touch $project_name/models/order.go
touch $project_name/services/user_service.go
touch $project_name/services/product_service.go
touch $project_name/services/order_service.go
touch $project_name/repositories/user_repository.go
touch $project_name/repositories/product_repository.go
touch $project_name/repositories/order_repository.go
touch $project_name/utils/http_utils.go
touch $project_name/utils/validation_utils.go
touch $project_name/config/config.go

echo "Struttura del progetto creata con successo."
