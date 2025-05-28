-- Exemplo de Row Level Security para logline_global
alter table logline_global enable row level security;

create policy "Users can insert their own logs"
  on logline_global for insert
  using (auth.uid() is not null);

create policy "Users can view their own logs"
  on logline_global for select
  using (auth.uid() is not null);