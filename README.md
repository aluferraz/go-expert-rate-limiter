``` 
docker compose up --build
```

# README do Projeto

## Visão Geral

O código fornecido consiste em três pacotes principais do Go: `rate_limit`, `middleware` e `web`.

### Pacote rate_limit:

O pacote `rate_limit` oferece uma maneira de limitar o número de solicitações feitas por clientes individuais a uma aplicação. A ideia básica por trás da limitação de taxa é controlar com que frequência usuários autenticados, ou, neste caso, endereços IP (ou chaves de API), podem fazer solicitações ao seu servidor.

#### Como Funciona:

**Contadores de Solicitação:** Quando um cliente faz uma solicitação, a estrutura `RateLimitRepositoryRedis` verifica se já existe um contador para esse cliente específico no Redis. Se não existir, ele cria um com um valor inicial de `session.GetRequestsLimitInSeconds()`. Esse contador representa quantas solicitações ainda podem ser feitas antes que o cliente seja limitado pela taxa.

https://github.com/aluferraz/go-expert-rate-limiter/blob/f3e1443242d979640bd04b355d88a97bd1e65d6a/internal/infra/persistence/rate_limit/rate_limit_redis.go#L23-L32


**Diminuir o Token Bucket:** Cada vez que um cliente faz uma solicitação, o contador é diminuído. Se o contador atingir zero, significa que o cliente fez tantas solicitações quanto permitido dentro da janela de tempo atual (1 segundo neste caso), e solicitações adicionais serão limitadas até que a próxima janela de tempo comece.

https://github.com/aluferraz/go-expert-rate-limiter/blob/f3e1443242d979640bd04b355d88a97bd1e65d6a/internal/infra/persistence/rate_limit/rate_limit_redis.go#L46-L94

**Redefinir o Contador:** A cada segundo, uma tarefa em segundo plano é executada, redefinindo o contador de solicitações para cada cliente ao seu limite máximo.

**Limitar o Cliente:** Se um cliente fizer mais solicitações do que o permitido em um segundo, ele será limitado e não poderá fazer mais solicitações até que o tempo definido na variável de ambiente `EXPIRATION` termine. 


Com o [middleware](https://github.com/aluferraz/go-expert-rate-limiter/blob/master/internal/infra/web/middleware/rate_limiter.go) configurado o método `DecreaseTokenBucket` é chamado em cada solicitação e caso retorne um erro (significando que o cliente foi limitado), [retorna um código de status HTTP 429 (Too Many Requests) para o cliente.
](https://github.com/aluferraz/go-expert-rate-limiter/blob/f3e1443242d979640bd04b355d88a97bd1e65d6a/internal/infra/web/middleware/rate_limiter.go#L43)
Estratégia de rate limit aplicada:

https://en.wikipedia.org/wiki/Token_bucket

## Variáveis de Ambiente

A tabela abaixo lista as variáveis de ambiente utilizadas pelo projeto, junto com suas descrições e valores padrão:

| Variável de Ambiente | Descrição | Valor Padrão |
| -------------------- | --------- | ------------ |
| WEBSERVER_PORT | A porta na qual o servidor web escutará as conexões recebidas. | 8080 |
| REDIS_URI | A URI do servidor Redis usado para armazenar dados de limitação de taxa. Inclui host e porta. | redis:6379 |
| IP_THROTTLING | O número máximo de solicitações que um endereço IP pode fazer dentro de um determinado intervalo de tempo (especificado em EXPIRATION). | 5 |
| API_THROTTLING | O número máximo de solicitações para qualquer ponto de extremidade de API dentro de um determinado intervalo de tempo (especificado em EXPIRATION). | 10 |
| EXPIRATION | A duração do intervalo de tempo, em segundos, durante o qual os limites de taxa são aplicados. | 60 |

