# VoulezVous Trinity

## Princípios Fundamentais

1. **Realidade Auditável**: Tudo é registrado como minicontrato.
2. **Anonimato Digno**: Identidades preservadas, ações públicas.
3. **Consequência Simbólica**: Todo ato gera reação institucional.
4. **Orquestração Universal**: LogLine DSL para infra, fluxo, CI/CD, componentes e observabilidade.

## Núcleo do Ecossistema

- **/institutional/templates/minicontratos/**: 20+ templates canônicos
- **/institutional/templates/flows/**: Fluxos institucionais reais, encadeáveis
- **/institutional/consequences/dispatch_rules.json**: Regras automáticas de consequências
- **/institutional/tools/**: Digest, reconciler de ghosts, playground, logcard visual
- **/institutional/security/**: Assinatura, quarantine, hash
- **/docs/schema/**: minicontrato.schema.json, visualização e documentação viva

## Como Contribuir

- Crie minicontratos via API ou UI
- Visualize flows e consequences
- Consulte/resuma via digest.ts
- Edite templates e flows conforme sua operação

## Deploy

```bash
cd deploy && ./install.sh --env prod
```

---

**type**: *campo interpretativo, só pode ser atribuído por LLM autorizado (ver classify_minicontrato.ts e schema).  
Nunca manual, nunca obrigatório para validade.*

---

## LogLine DSL

- Orquestra migrations, seeds, infra, flows, CI/CD, componentes, workflows, testes, segurança e observabilidade em um único artefato versionado.
- Serve como "código-fonte" institucional e fonte de documentação viva.

---

## Próximos Passos

- Personalize templates conforme seus fluxos reais
- Defina consequences institucionais em dispatch_rules.json
- Rode reconciliador e digest periodicamente
- Amplie flows e playground conforme novas operações

---

## Referências

- [Manifesto LogLine DSL](#)
- [Schema minicontrato](./schema/minicontrato.schema.json)