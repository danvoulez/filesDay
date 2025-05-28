-- Exemplo de Row Level Security para file_entry
alter table file_entry enable row level security;

create policy "Users can insert their own files"
  on file_entry for insert
  using (auth.uid() is not null);

create policy "Users can view their own files"
  on file_entry for select
  using (auth.uid() is not null);