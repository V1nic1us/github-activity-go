# Atividade do Usuário no GitHub

Neste projeto, você irá construir uma interface de linha de comando (CLI) simples para buscar a atividade recente de um usuário do GitHub e exibi-la no terminal. Este projeto ajudará você a praticar suas habilidades de programação, incluindo trabalhar com APIs, manipular dados em JSON e construir uma aplicação CLI simples.

## Requisitos

A aplicação deve rodar a partir da linha de comando, aceitar o nome de usuário do GitHub como argumento, buscar a atividade recente do usuário usando a API do GitHub e exibi-la no terminal. O usuário deve ser capaz de:

- Fornecer o nome de usuário do GitHub como argumento ao rodar a CLI.
  
```bash
  github-activity <username>
```
Buscar a atividade recente do usuário especificado usando a API do GitHub. Você pode usar o seguinte endpoint para buscar a atividade do usuário:

```shell
# https://api.github.com/users/<username>/events
# Exemplo: https://api.github.com/users/kamranahmedse/events
```

Exibir a atividade buscada no terminal.

```
Saída:
- Enviou 3 commits para kamranahmedse/developer-roadmap
- Abriu um novo issue em kamranahmedse/developer-roadmap
- Deu um star em kamranahmedse/developer-roadmap
- ...
```

Você pode [aprender mais sobre a API do GitHub aqui.](https://docs.github.com/en/rest/activity/events?apiVersion=2022-11-28)

Tratar erros de forma adequada, como nomes de usuários inválidos ou falhas na API.
Use a linguagem de programação de sua escolha para construir este projeto.
Não use nenhuma biblioteca ou framework externo para buscar a atividade do GitHub.
Se você deseja construir uma versão mais avançada deste projeto, pode considerar adicionar funcionalidades como filtrar a atividade por tipo de evento, exibir a atividade em um formato mais estruturado ou armazenar em cache os dados buscados para melhorar o desempenho. Você também pode explorar outros endpoints da API do GitHub para buscar informações adicionais sobre o usuário ou seus repositórios.

URL
https://roadmap.sh/projects/github-user-activity

Build
Execute o comando para gerar o CLI para Windows e Linux.

bash
./build.sh
