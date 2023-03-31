# Laborat�rio de Redes - Aula 2

## Conte�do

- Laborat�rio de Redes
    - Conte�do
    - Objetivo
    - DNS
    - Passo a passo

## Objetivo

Separar a responsabilidade de servidores.

Criar um servidor para DNS, DHCP e Email


## DNS

DNS (Domain Name System), � um sistema hier�rquico e distribu�do de gest�o de nomes para computadores, servi�os ou qualquer m�quina conectada � Internet ou a uma rede privada.

## Passo a passo


O servidor DNS é o pai de todos, ele redireciona para outros como email e o proprio site.


Cada servidor deve ter um email statico bem definido e o DNS aponta para eles

Toda comunicação feita passa pelo DNS antes de ir ao servidor ideal.

O servidor de DNS deve apontar a outros e definir o nome do dominio do servidor de emails como por exemplo "serveremail.com" (toda a comunicação de emails deve passar por esse servidor)

O servidor de DHCP Aponta ao DNS e inicia em uma porta distinta aos servidores para que os aparelhos conectados a rede seja configurado automaticamente com DHCP

O servidor de EMAIL configura as contas como "Remetente" e "Destinatario" além de definir qual será dominio do email como "meuemail.com.br"

-------------

Cada maquina terá um DHCP para configurar seu IP automaticamente e deve configurar seu email sendo uma para "remetente" ou "destinatario"

