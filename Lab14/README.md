# Laboratório de Redes - Aula 14

## Objetivos

- Configurar um Firewall para bloquear acesso HTTP

```bash
# No Firewall

enable
configure terminal
interface g1/1
ip address 192.168.2.2 255.255.255.252 #Máximo de 2 hosts
no shutdown

nameif cliente
security-level 0

exit

interface g1/2
ip address 192.168.3.2 255.255.255.252 #Máximo de 2 hosts
no shutdown

nameif servidor
security-level 0

router rip

network 192.168.2.0
network 192.168.3.0

access-list permitir-http permit tcp 192.168.1.0 255.255.255.0 host 192.168.3.2 eq 80
access-group permitir-http in interface cliente

```

**Note:** para configurar o DNS é necessario mapear tanto a Ida quando a Volta com `permitir-dns` e `permitir-dns-retorno`
