# tcazanga-estagio-2024-04
**Repostório dedicado ao Estágio de Programação na Zafir Tecnologia*

*Membros:*
1. Tito Fernando Cazanga
2. Silvano Paulino
3. Manuel Afonso

# Fitness CLI

Este é um CLI para gerenciamento de operações em um ambiente fitness. Ele oferece funcionalidades para abastecer expositores, consumir produtos, emitir faturas, registrar remessas e instalar expositores.

## Instalação

Certifique-se de ter o Go instalado em sua máquina. Para instalar este CLI, execute:

```   
go get -u github.com/seu-usuario/fitness-cli


## Uso

*abastecer-expositor*
Abastece um expositor com um determinado produto.
fitness-cli abastecer-expositor [expositorID] [produtoID] [quantidade]

*Exemplo*
fitness-cli abastecer-expositor expositor1 101 20  


*consumir-produto*
Consome um produto de um expositor.
fitness-cli consumir-produto [expositorID] [produtoID] [quantidade]

*Exemplo:*
fitness-cli consumir-produto expositor1 102 5


*emitir-fatura*
Emite uma fatura para um ginásio com base nos consumos de produtos. 
fitness-cli emitir-fatura [ginasioID] [consumos]

*Exemplo:* 
fitness-cli emitir-fatura ginasio1 expositor1:101:10,expositor2:102:5


*instalar-expositor*
Instala um novo expositor com produtos iniciais.
fitness-cli instalar-expositor [expositorID] [localizacao] [produtos]

*Exemplo:*
fitness-cli instalar-expositor expositor2 "Localização 2" "101:50,102:30"


*novo-guia-remessa*
Cria um novo Guia de Remessa.
fitness-cli novo-guia-remessa [origemID] [destinoID] [produtoID] [quantidade]

*Exemplo:*
fitness-cli novo-guia-remessa origem1 destino1 101 20


*registrar-remessa*
Registra uma remessa entre dois expositores. 
fitness-cli registrar-remessa [origemID] [destinoID] [produtoID] [quantidade]

*Exemplo:*
fitness-cli registrar-remessa origem1 destino1 101 20


Para obter mais detalhes sobre cada comando e suas opções, use a flag --help após o comando. Por exemplo:
fitness-cli abastecer-expositor --help
Isso exibirá informações detalhadas sobre como usar o comando abastecer-expositor.

_Divirta-se usando o Fitness CLI!_