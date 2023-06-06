# Comandos de configuração

```bash
# No Roteador

enable
configure terminal
hostname ISP
int g0/1
ip address 201.150.10.100 255.255.255.0
no shutdown

exit
ip route 0.0.0.0 0.0.0.0 g0/1

ctrl^C

copy run start

# No Switch

enable
configure terminal
hostname S1
vlan 10
vlan 20

ctrl^C

show vlan

configure terminal
int f0/1
switchport mode access
switchport access vlan 10
int f0/10
switchport mode access
switchport access vlan 20

ctrl^C

show vlan

configure terminal
int vlan 10
ip address 192.168.10.1 255.255.255.0
no shutdown

int vlan 20
ip address 192.168.20.1 255.255.255.0
no shutdown

ip routing # Ativa o roteamento

int g0/1
no switchport
ip address 201.150.10.101 255.255.255.0
no shutdown

exit

ip default-gateway 201.150.10.100

```
