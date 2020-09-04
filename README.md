# Descritor

![Tests](https://github.com/candidatos-info/descritor/workflows/Tests/badge.svg?branch=master)

Estrutura de dados que descreve os dados armazenados pelo candidatos.info. Também exporta métodos que auxiliam na conversão entre os formatos. 

## Guia para alteraçōes de schema
### Pré-requisitos
+ [Go instalado](https://golang.org/doc/install);
+ [Git instalado](https://gist.github.com/derhuerst/1b15ff4652a867391f03);
+ [Compilador protoc instalado](http://google.github.io/proto-lens/installing-protoc.html);

Sempre que houver necessidade de alterar o schema da estrutura de candidatura siga os seguintes passos:

1. Clone o repositório;
2. Crie uma nova branch para fazer as novas alteraçōes;
3. Faça no arquivo [candidatura.proto](https://github.com/candidatos-info/descritor/blob/master/candidatura.proto) as alterações necessárias;
4. Compile o protocol buffer com o seguinte comando:
```
$ protoc --go_out=. --proto_path=. candidatura.proto
```
5. <b>Com cuidado</b> edite o arquivo gerado, candidatura.pb.go, e insira as devidas tags para formato JSON e CSV;
6. Rode os testes do pacote e eles devem quebrar por conta de suas alteraçōes;
7. Corrija os testes e quando tiver a suite de testes verde abra um [pull request](https://www.digitalocean.com/community/tutorials/como-criar-um-pull-request-no-github-pt) neste repositório e solicite avaliação de um dos mantenedores;
