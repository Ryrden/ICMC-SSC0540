# Laboratório de Redes - Aula 12

## Objetivos

- Configurar dispositivos IoT
- Configurar um servidor DNS em outra rede

## Passo a passo no Cisco Packet tracer

1. defina o servidor com rede X e habilite o serviço DHCP e IoT
2. defina um access point com a rede X e conecte hospedeiros IoT com a conexão do tipo remota
   1. o servidor é o endereço do servidor IoT da rede X
   2. o login e senha devem ser criados no primeiro acesso, por padrão o login é `admin` e a senha `admin`
3. crie um roteador que conecte a rede X com a rede Y que vai ser criada no passo 4
4. Crie a rede Y com o servidor habilitando o DNS
5. conecte o roteador com a rede Y e faça as devidas configurações de RIP e de portas com gateway padrão para ambas as redes

Teste tudo