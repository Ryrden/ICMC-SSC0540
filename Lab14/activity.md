# Laboratório de Redes - Aula 15

Questões da Camada de Enlace baseadas nas duas últimas aulas da univesp (aulas 13 e 14).

1. Explique duas funções importantes desempenhadas pela camada de enlace na Internet.

    **R:** A camada de enlace é responsável por transmitir os dados entre dois nós da rede local. Ela também é responsável por detectar e corrigir erros que possam ocorrer durante a transmissão dos dados, tudo isso a nível de hardware.

2. Defina o conceito de acesso múltiplo e identifique as três formas de acesso múltiplo discutidas na vídeo-aula.

    **R:** Acesso múltiplo é quando vários nós da rede compartilham o mesmo meio de transmissão. As três formas de acesso múltiplo são: Divisão de canal, Acesso aleatório e Passagem de revezamento. Para a Divisão de canal, o canal é dividido em várias partes (tempo e frequência), cada uma delas é alocada para um nó. Para o Acesso aleatório, o canal não é dividido o que pode causar colisões, desta forma, os nós transmitem os dados de forma aleatória, sem nenhuma ordem. Para a Passagem de revezamento, existe uma seleção (Pooling) para determinar qual nó irá transmitir os dados, desta forma, os nós transmitem os dados de forma ordenada por meio de um token.

3. Destaque a diferença fundamental entre o CSMA (Acesso Múltiplo por Detecção de Portadora) e o CSMA/CD (Acesso Múltiplo por Detecção de Colisão).

    **R:** A diferença fundamental entre o CSMA e o CSMA/CD está no tratamento de colisões. O CSMA apenas verifica se o meio está ocupado antes de iniciar a transmissão, enquanto o CSMA/CD adiciona a detecção de colisão e a recuperação rápida em caso de colisões.

4. Descreva o método de endereçamento utilizado na camada de enlace.

    **R:** O endereçamento da camada de enlace é feito por meio do endereço MAC (Media Access Control), que é um endereço físico único para cada placa de rede. Este endereço é composto por 48 bits, sendo os primeiros 24 bits referentes ao fabricante e os últimos 24 bits referentes ao número de série da placa. O endereço MAC é gravado na memória ROM da placa de rede e não pode ser alterado. Sua notação é feita por meio de 6 pares de números hexadecimais separados por dois pontos, por exemplo: 00:0A:95:9D:68:16.

5. Identifique a função principal do protocolo ARP (Protocolo de Resolução de Endereço) na camada de enlace.

    **R:** O protocolo ARP é responsável por mapear o endereço IP para o endereço MAC. Ele é utilizado quando um nó precisa enviar um pacote para outro nó da rede local, mas não sabe o endereço MAC do nó de destino. Sua estrutura é dada por uma tupla (IP, MAC, TTL), onde o TTL é o tempo de vida da tupla, ou seja, o tempo que o mapeamento deve ficar armazenado na tabela ARP.
