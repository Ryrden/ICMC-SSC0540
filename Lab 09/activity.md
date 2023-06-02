Instituto de Ciências Matemáticas e de Computação - ICMC/USP

<center>

**Redes de Computadores**

**Atividade da Camada de Redes**

Prof. Jó Ueyama

</center>

Nome: Ryan Souza Sá Teles

No. USP: 12822062

1. Explique as funções chaves dos dois componentes do roteador: algoritmo de roteamento e a tabela de repasse.

    R: O algoritmo de roteamento é o que define qual o melhor caminho para o pacote seguir, enquanto a tabela de repasse é uma tabela interna do roteador armazena o endereço do destino para onde o pacote deve ser enviado.

---

2. Responda as duas perguntas abaixo:
   1. Qual a diferença entre roteamento e repasse implementados na camada de rede?
   2. Explique o problema conhecido como HOL (head of line blocking - bloqueio de cabeça de fila) existente nos roteadores

    (2.1) R: O roteamento é um processo mais complexo onde envolve algoritmos de roteamento, enquanto o repasse é um processo mais simples onde envolve apenas a tabela de repasse.

    (2.2) R: O HOL é um problema que ocorre quando um pacote fica enfileirado na frente da fila impedindo que os pacotes atrás dele sejam enviados.

---

3. Um determinado computador possui máscara 255.255.255.192. Qual o número máximo de hospedeiros que podem ser conectados nessa rede?

    R: um endereço IP tem 32 bits, como esta máscara ja ocupa 26 bits ativos, então sobram apenas 6 bits para os hosts, desta forma, temos 2^6 = 64 hospeiros. Vale salientar que existem dois endereços reservados, o endereço de rede e o endereço de broadcast, então o número máximo de hosts é 62.

---

4. O ICMP trata da comunicação de erros na camada de rede. Um dos erros tratados é o “estouro do TTL”. Explique o que significa o estouro do TTL na camada de rede.

    R: O "Estouro do TLL" ocorre quando o TTL (Time To Live) chega a zero, ou seja, o pacote não conseguiu chegar ao destino e foi descartado. Um dos exemplos em terminal é o comando "ping", onde o TTL é decrementado a cada salto que o pacote dá.

---

5. Qual o problema chave que levou ao surgimento do protocolo IPv6? Por que o IPv6 ainda não é largamente utilizado?

    R: O problema que levou o surgimento do protocolo IPv6 foi a necessidade de mais endereços IP, pois o IPv4 não suportava mais a demanda de endereços. O IPv6 ainda não é largamente utilizado pois a maioria dos dispositivos ainda utilizam o IPv4, e a transição de um protocolo para o outro é um processo lento e complexo.
