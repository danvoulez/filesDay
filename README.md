# FlipApp

O Flip é o módulo de registro rápido, auditável e confiável de ações, decisões, dúvidas e eventos institucionais.

## Fluxo:

1. Usuário preenche rapidamente "Quem fez?", "O que foi feito?" e "Sobre o quê?" (opcional).
2. Sugestões inteligentes completam ou validam a entrada.
3. Registro é enviado como LogLine auditável (pode ser valid, ghost, pending).
4. Feedback imediato é exibido ao usuário.
5. Integração nativa com templates institucionais e flows.

## Como rodar

1. Instale dependências e rode o app:
   ```bash
   npm install
   npm run dev
   ```

2. Integre o endpoint `/api/logline` ao seu backend institucional.

---

Pronto para expandir com flows, IA, validação simbólica, e visualização de LogLines.