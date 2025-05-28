-- Seeds de templates básicos (exemplo)
insert into flow_template (id, steps, required_roles, max_duration, final_conditions)
values (
  'venda_simples',
  '["iniciar_venda","selecionar_produto","efetuar_pagamento","finalizar_venda"]',
  '["user","caixa"]',
  '2 hours',
  '{"status":"executed"}'
)
on conflict (id) do nothing;