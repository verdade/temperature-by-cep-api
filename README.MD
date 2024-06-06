
### Subindo o Projeto

Para subir o projeto, você tem duas opções:


1. **Construa a imagem e suba o projeto:**

   Siga os comandos abaixo para construir a imagem e subir o projeto:

   ```bash
   make docker-build-image
   make docker-up
   ```

## Executando Testes

Para executar os testes da aplicação, utilize o comando:

```bash
make test
```

Para gerar a cobertura de código e um relatório em `coverage.html`, execute:

```bash
make test-coverage
```

Após a execução, o arquivo `coverage.html` será gerado na raiz do projeto, permitindo que você visualize detalhadamente a cobertura de testes.

## Uso da API

A Temperature API está disponível no seguinte link: http://localhost:8080/temperature

Para utilizar, passe o CEP desejado via body e faça chamada via POST, por exemplo:
```bash
curl --location 'http://localhost:8080/temperature' \
--header 'Content-Type: application/json' \
--data '{
    "cep": "01153000"
}'
```

Para acessar os traces da aplicação acesse: http://localhost:9411/zipkin/
