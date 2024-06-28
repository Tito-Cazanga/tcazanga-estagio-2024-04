CVL - Aplicação CLI para Internar Paciente
=====================================

Descrição
-----------

CVL é uma aplicação CLI desenvolvida em Go que permite internar pacientes.

Funcionalidades
---------------

* Internar paciente
* Fazer as Rondas

Uso
----
Certifique-se de que você tem o Go instalado em sua máquina.

Instalação
------------

Para instalar o CVL, deves clonar o repositório:
```
git clone github.com/Tito-Cazanga/tcazanga-estagio-2024-04
```
Certifique-se de entrar no Directorio *centro-veterinaria-luanda*

```
cd centro-veterinario-luanda/
```

Depois, compile o código com o comando:
```
go build -o bin/cvl cmd/cvl.go
```

Execute o comando o seguinte para iniciar a aplicação. Onde terá as instruões
```
bin/cvl
```
Para internar paciente execute o comando seguinte:
```
bin/cvl internar <id_do_paciente> <nome_do_paciente> <raça_do_paciente>
```
Para consultar paciente execute o comando seguinte:
```
bin/cvl consultar <id_do_paciente> 
```
Autor
------
1. Denilson Simões
2. Tito Cazanga
