# Documentação: Ferramenta de Stress Test em Go

Essa aplicação CLI foi desenvolvida para realizar testes de carga em serviços web. A ferramenta permite medir a capacidade do serviço em atender um número definido de requisições simultâneas e apresenta um relatório detalhado com métricas úteis após a execução.

## Como Funciona
1. O usuário fornece a URL do serviço, o número total de requisições e a quantidade de chamadas simultâneas através de parâmetros CLI.
2. A aplicação distribui as requisições de forma paralela entre múltiplos "workers" (goroutines).
3. Cada worker realiza um número específico de requisições HTTP para a URL fornecida.
4. Ao final da execução, a aplicação gera um relatório com:
   - Tempo total gasto na execução.
   - Total de requisições realizadas.
   - Número de requisições bem-sucedidas (status HTTP 200).
   - Distribuição de outros códigos de status HTTP.

## Parâmetros de Configuração
Os parâmetros podem ser configurados ao executar a aplicação via linha de comando:

- `--url`: (**Obrigatório**) Especifica a URL do serviço a ser testado.
- `--requests`: (**Opcional**) Número total de requisições a serem realizadas. O padrão é `100`.
- `--concurrency`: (**Opcional**) Número de chamadas simultâneas (goroutines). O padrão é `10`.

### Exemplos de Uso

#### Teste Básico
Para realizar um teste básico com 100 requisições, usando 10 chamadas simultâneas, execute:
```bash
go run main.go --url=https://example.com --requests=100 --concurrency=10
```

#### Teste com Mais Carga
Para realizar 1000 requisições com 50 chamadas simultâneas:
```bash
go run main.go --url=https://example.com --requests=1000 --concurrency=50
```

#### Teste com Docker
A ferramenta também pode ser executada em um container Docker. Primeiro, crie a imagem Docker:
```bash
docker build -t stress-tester .
```
Em seguida, execute o container:
```bash
docker run stress-tester --url=https://example.com --requests=1000 --concurrency=50
```

## Relatório Gerado
Após a execução do teste, será exibido um relatório no seguinte formato:

```plaintext
--- Resultado dos Testes ---
Tempo Total Gasto: 12.34 segundos
Total de Requests: 1000
Requests com Status 200: 980
Distribuição de Status HTTP:
200: 980
404: 10
500: 5
Erros: 5
```

### Descrição do Relatório
- **Tempo Total Gasto:** O tempo total necessário para executar todas as requisições.
- **Total de Requests:** O número total de requisições realizadas (incluindo sucessos e erros).
- **Requests com Status 200:** Quantidade de requisições com resposta bem-sucedida.
- **Distribuição de Status HTTP:** Contagem de respostas agrupadas por código de status.
- **Erros:** Número de requisições que falharam devido a problemas de rede ou outros erros.