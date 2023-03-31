# Laboratório de Redes - Aula 2

## Conteúdo

- Laboratório de Redes
    - Conteúdo
    - Objetivo
    - DNS
    - Passo a passo

## Objetivo

Separar a responsabilidade de servidores.

Criar um servidor para DNS, DHCP e Email


## DNS

DNS (Domain Name System), é um sistema hierárquico e distribuído de gestão de nomes para computadores, serviços ou qualquer máquina conectada à Internet ou a uma rede privada.

## Passo a passo


O servidor DNS Ã© o pai de todos, ele redireciona para outros como email e o proprio site.


Cada servidor deve ter um email statico bem definido e o DNS aponta para eles

Toda comunicaÃ§Ã£o feita passa pelo DNS antes de ir ao servidor ideal.

O servidor de DNS deve apontar a outros e definir o nome do dominio do servidor de emails como por exemplo "serveremail.com" (toda a comunicaÃ§Ã£o de emails deve passar por esse servidor)

O servidor de DHCP Aponta ao DNS e inicia em uma porta distinta aos servidores para que os aparelhos conectados a rede seja configurado automaticamente com DHCP

O servidor de EMAIL configura as contas como "Remetente" e "Destinatario" alÃ©m de definir qual serÃ¡ dominio do email como "meuemail.com.br"

-------------

Cada maquina terÃ¡ um DHCP para configurar seu IP automaticamente e deve configurar seu email sendo uma para "remetente" ou "destinatario"

