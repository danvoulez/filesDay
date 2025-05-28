# LogLine DSL – Linguagem Universal de Orquestração Institucional

## Visão Geral
- Orquestra migrations, infra, CI/CD, workflows, componentes, testes e observabilidade em uma linguagem única e auditável.
- Cada bloco é um comando simbólico, versionado, rastreável e integrável.

## Exemplos

```logline
migration 0001_core {
  CREATE TABLE users(id UUID PRIMARY KEY, name TEXT);
}

infra functions {
  deploy user-create from "users/create.ts";
}

workflow UserOnboarding {
  step 1: #1 user:open "GET /users";
  step 2: #2 user:input "name";
  step 3: execute user-create if_ok="refresh";
}
```

## Ferramentas

- logline init / compile / serve / test / docs
- Plugins para VSCode e JetBrains

## Adote: Use como fonte única de verdade para infra, operações e docs.