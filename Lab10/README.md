# Comandos de configuração

```bash
# No Roteador

enable
configure terminal 
hostname ISP # Configurando o nome do roteador
int g0/1 # Configurando a interface 0/1
ip address 201.150.10.100 255.255.255.0 # Configurando o ip da interface
no shutdown # Habilitando a interface

exit 
ip route 0.0.0.0 0.0.0.0 g0/1 # Configurando a rota padrão

ctrl^C

copy run start # Salvando as configurações

# No Multilayer Switch

enable
configure terminal
hostname S1
vlan 10
vlan 20

ctrl^C

show vlan

# Configurando as interfaces que serão utilizadas como vlan
configure terminal
int f0/1
switchport mode access
switchport access vlan 10
int f0/10
switchport mode access
switchport access vlan 20

ctrl^C

show vlan

# Configurando o ip das vlans
configure terminal
int vlan 10
ip address 192.168.10.1 255.255.255.0 
no shutdown # Habilitando a interface

int vlan 20
ip address 192.168.20.1 255.255.255.0
no shutdown # Habilitando a interface

ip routing # Ativa o roteamento entre as vlans

# Configurando a interface que será conectada ao roteador
int g0/1
no switchport
ip address 201.150.10.101 255.255.255.0
no shutdown # Habilitando a interface

exit

ip default-gateway 201.150.10.100 # Configurando o gateway padrão

```

## Exercicio

- Desenvolver o DHCP no switch para ambas as redes

- Incluir dois hospedeiros em cada VLAM com DHCP

```bash

# No multilayer switch

# Configurando o novo range de IP das interfaces para permitir novos hosts
interface range f0/1-10
switchport mode access
switchport access vlan 10

interface range f0/11-20
switchport mode access
switchport access vlan 20


# Configurando o DHCP para cada rede (VLAN)
ip dhcp pool IP10 # Criando o pool de ip
network 192.168.10.0 255.255.255.0  # Configurando a rede
default-router 192.168.10.1  # Configurando o ip do gateway

ip dhcp pool IP20  # Criando o pool de ip
network 192.168.20.0 255.255.255.0 # Configurando a rede
default-router 192.168.20.1 # Configurando o ip do gateway


```

