-- Cria tabela de arquivos e logs para filesnight
create table if not exists file_entry (
  id uuid primary key default gen_random_uuid(),
  filename text not null,
  url text,
  uploaded_by uuid references auth.users(id),
  uploaded_at timestamptz default now(),
  metadata jsonb
);

create table if not exists file_log (
  id uuid primary key default gen_random_uuid(),
  file_id uuid references file_entry(id),
  action text not null,
  performed_by uuid references auth.users(id),
  performed_at timestamptz default now(),
  details jsonb
);