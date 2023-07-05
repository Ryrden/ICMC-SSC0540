# SSC0540 - Redes de Computadores (2023)

## Informações Gerais

A disciplina foi ministrada pelo Prof. Dr. Jó Ueyama, no 1º semestre de 2023. Mais informações podem ser encontradas na página da disciplina no [JupiterWeb](https://uspdigital.usp.br/jupiterweb/obterDisciplina?sgldis=SSC0540&codcur=55051&codhab=4).

Durante o curso foi utilizado as video-aulas da [Univesp](https://www.youtube.com/playlist?list=PLxI8Can9yAHc-_dZ6nsfoon08i2-4OvEk) e o livro *Redes de Computadores e a Internet - Uma Abordagem Top-Down* (Kurose, J. F. e Ross, K. W., 6ª edição, Pearson, 2013).

Para as atividades práticas foi utilizado o software [Cisco Packet Tracer](https://www.netacad.com/pt-br/courses/packet-tracer) e o software [Wireshark](https://www.wireshark.org/).

o Wireshark foi utilizado apenas nas aulas de questões utilizando o material [Wireshark Labs](http://www-net.cs.umass.edu/wireshark-labs/).

## Por que este repositório?

Este repositório foi criado para armazenar os códigos, atividades e anotações feitas durante a disciplina de Redes de Computadores (SSC0540) do curso de Bacharelado em Sistema de Informações da Universidade de São Paulo (USP).

Ao longo do semestre, foram dados 14 Laboratórios de atividades práticas, 2 Trabalhos de implementação e nenhuma Prova. Todos os Laboratórios e Trabalhos foram feitos de forma individual.

## Estrutura do repositório e conteúdo

O curso seguiu de maneira Top-Down o conteúdo, desta forma, a cada conjunto de pastas uma camada do modelo OSI é abordada. A seguir, o conteúdo de cada pasta:

### **Camada de Aplicação**

- [Lab 1](./Lab01/) - Introdução ao Cisco Packet Tracer e criação de uma rede local simples com dispositivos básicos
- [Lab 2](./Lab02) - Uso de servidores DHCP, DNS e HTTP
- [Lab 3](./Lab03) - Uso de servidor SMTP e FTP
- [Lab 4](./Lab04) - Introdução ao Wireshark e análise de mensagens HTTP com uso do material [Wireshark Lab: HTTP v8](http://www-net.cs.umass.edu/wireshark-labs/Wireshark_HTTP_v8.0.pdf)

### **Camada de Transporte**

- [Lab 5](./Lab05) - Uso de um roteador para criar 3 redes distintas e e separar responsabilidade de servidores DNS e HTTP. Também foi abordado uma exemplificação em código de conexão TCP e UDP.
- [Lab 6](./Lab06) - Análise de segmentos TCP com uso do material [Wireshark Lab: TCP v8](http://www-net.cs.umass.edu/wireshark-labs/Wireshark_TCP_v8.0.pdf)
- [Lab 7](./Lab07) - Criação de 5 redes distintas com uso de 3 roteadores, separar responsabilidade de servidores DNS, HTTP e FTP
- [Lab 8](./Lab08) - Configuração de uma VLAN usando roteadores e switches, foi configurado também um servidor DHCP [(video apoio)](https://youtu.be/wmFxjvH1lyo)

### **Camada de Rede**

- [Lab 9](./Lab09) - Perguntas sobre a camada de rede
- [Lab 10](./Lab10) - Configuração de uma VLAN usando um Switch Multilayer, foi configurado também um servidor DHCP através do switch [(video apoio 1)](https://youtu.be/EjOEdqkWPvE) [(video apoio 2)](https://youtu.be/ED6hVImjs2A)
- [Lab 11](./Lab11) - Análise de datagramas IPv4 com uso do material [Wireshark Lab: IP v8](http://www-net.cs.umass.edu/wireshark-labs/Wireshark_IP_v8.0.pdf)
- [Lab 12](./Lab12) - Configuração de um servidor IoT e DNS para mapeamento do endereço HTTP do servidor IoT [(video apoio)](https://youtu.be/JEITMT-XcjU)
- [Lab 13](./Lab13) - Configuração de um Firewall para bloqueio de acesso HTTP, FTP e DNS [(video apoio)](https://youtu.be/-iNnNf5tSi0)

### **Canada de Enlace**

- [Lab 14](./Lab14) - Perguntas sobre a camada de enlace - [(Aula 13 Univesp)](https://youtu.be/zp9mQbGvndw) e [(Aula 14 Univesp)](https://youtu.be/l93e4P4d9mk)
