### PROJECTO ACME

*Membros:*
1. Tito Fernando Cazanga
2. Silvano Paulino
3. Manuel Afonso

# Fitness CLI

Este é um CLI para gerenciamento de operações em um ambiente fitness. Ele oferece funcionalidades para abastecer expositores, consumir produtos, emitir faturas, registrar remessas e instalar expositores.

## Instalação

Certifique-se de ter o Go instalado em sua máquina. Para instalar este CLI, execute:
  


## Uso

*abastecer-expositor*
Abastece um expositor com um determinado produto.
```
  go run main.go abastecer-expositor --expositorID[expositorID] --produtoID[produtoID] --quantidade[quantidade]
```

*Exemplo*
```
  go run main.go abastecer-expositor expositor1 101 20  
```


*consumir-produto*
Consome um produto de um expositor.
```
  go run main.go consumir-produto expositorID[expositorID] produtoID[produtoID] quantidade[quantidade]
```
*Exemplo:*
go run main.go consumir-produto expositor1 102 5


*emitir-fatura*
Emite uma fatura para um ginásio com base nos consumos de produtos. 

```
  go run main.go emitir-fatura ginasioID[ginasioID] consumos[consumos]
```

*Exemplo:* 
go run main.go emitir-fatura ginasio1 expositor1:101:10,expositor2:102:5


*instalar-expositor*
Instala um novo expositor com produtos iniciais.
go run main.go instalar-expositor expositorID[expositorID] a[localizacao] produtos[produtos]
```
```

*Exemplo:*
go run main.go instalar-expositor expositor2 "Localização 2" "101:50,102:30"
```
```

*novo-guia-remessa*
Cria um novo Guia de Remessa.
go run main.go novo-guia-remessa origemID[origemID] destinoID[destinoID] produtoID[produtoID] quantidade[quantidade]
```
```

*Exemplo:*
go run main.go novo-guia-remessa origem1 destino1 101 20
```
```

*registrar-remessa*
Registra uma remessa entre dois expositores. 
go run main.go registrar-remessa origemID[origemID] destinoID[destinoID] produtoID[produtoID] quantidade[quantidade]
```
```

*Exemplo:*
go run main.go registrar-remessa origem1 destino1 101 20
```
```

Para obter mais detalhes sobre cada comando e suas opções, use a flag --help após o comando. Por exemplo:
go run main.go abastecer-expositor --help
```
```
Isso exibirá informações detalhadas sobre como usar o comando abastecer-expositor.

_Divirta-se usando o Fitness CLI!_
