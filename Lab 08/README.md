# V-LAN commands

## Objetivo desta aula

- criar duas vlans em um switch
- configurar uma subinterface em um router
- configurar um pool de ip em um router
- configurar um servidor dhcp e dns em um router

## O que é uma V-LAN?

Uma vlan é uma rede virtual, ou seja, uma rede que não existe fisicamente, mas que é criada logicamente. Ela é criada para separar redes em um mesmo switch, ou seja, é possível criar uma vlan para o departamento de marketing e outra para o departamento de vendas, por exemplo.

Uma vlan é útil porque permite que um switch seja dividido em várias redes, sem a necessidade de comprar um switch para cada rede.

```bash

## NO COMPUTADOR CONECTADO AO SWITCH

enable
configure terminal

vlan 10 # Criando ou acessando a vlan 10
name departamentoX # Nomeando a vlan 10

vlan 20 # Criando ou acessando a vlan 20
name departamentoY # Nomeando a vlan 20

interface range f0/1-10 # Selecionando a interface no range de 1 a 10 
switchport mode access # Configurando a interface como acesso
switchport access vlan 10 # Configurando a interface para a vlan 10

interface range f0/11-20 # Selecionando a interface no range de 11 a 20
switchport mode access # Configurando a interface como acesso
switchport access vlan 20 # Configurando a interface para a vlan 20

show vlan # Comando para verificar as vlans criadas

## NO CLI DO SWITCH

enable
configure terminal
inter fastEthernet 0/24 # Selecionando a interface 24
switchport mode trunk # Configurando a interface como trunk (para conectar ao router)

## NO COMPUTADOR CONECTADO NO ROTEADOR

interface g0/0.10 # Criando a subinterface 10
encapsulation dot1Q 10 # Configurando a subinterface como dot1Q
ip address 192.168.1.1 255.255.255.0 # Configurando o ip da subinterface

interface g0/0.20 # Criando a subinterface 20
encapsulation dot1Q 20 # Configurando a subinterface como dot1Q
ip address 192.168.2.1 255.255.255.0 # Configurando o ip da subinterface

inter gigabitEthernet 0/0   # Configurando a interface 0/0
no shutdown # Habilitando a interface


## NO COMPUTADOR CONECTADO AO ROUTER OU NO CLI DO ROUTER

ip dhcp pool IP10 # Criando o pool de ip
network 192.168.1.0 255.255.255.0 # Configurando a rede
default-router 192.168.1.1 # Configurando o ip do gateway
dns-server 192.168.1.10 # Configurando o ip do servidor dns

ip dhcp pool IP20 # Criando o pool de ip
network 192.168.2.0 255.255.255.0
default-router 192.168.2.1
dns-server 192.168.1.10

show vlan brief # Comando para verificar as vlans criadas
```
