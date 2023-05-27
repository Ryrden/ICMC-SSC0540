# V-lan 

## V-lan commands

```bash
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

show vlan

switchport mode trunk

interface g0/0.2
encapsulation dot1Q 2
ip address 192.168.1.1 255.255.255.0

interface g0/0.3
encapsulation dot1Q 3
ip address 192.168.2.1 255.255.255.0

interface g0/0/0
no shutdown

ip dhcp pool IP10
network 192.168.1.0 255.255.255.0
default-router 192.168.1.1
dns-server 192.168.1.10

do show vlan brief
```