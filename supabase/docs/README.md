# Deploy Automático filesnight + Supabase

Este diretório contém tudo para instalar, atualizar e rodar o filesnight no Supabase, incluindo banco, funções, policies, triggers, templates, flows, consequences e documentação.

## Deploy Local

```bash
bash deploy/deploy_all.sh
```

## Deploy Automático

Push na branch main dispara o deploy completo via GitHub Actions.

## Estrutura

- migrations/: SQL do banco
- seeds/: dados iniciais
- functions/: edge functions TypeScript
- storage/: arquivos e buckets
- policies/: RLS/policies
- triggers/: triggers SQL
- templates/: contratos, flows, consequences
- docs/: documentação viva