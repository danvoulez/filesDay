#!/bin/bash
set -e

echo "🚀 Deploy automático filesnight + Supabase"

# 1. Deploy migrations e seeds
supabase db push
supabase db seed

# 2. Deploy edge functions (TypeScript)
supabase functions deploy --project supabase/functions

# 3. Deploy storage/buckets (opcional)
if [ -d "supabase/storage" ]; then
  supabase storage migrate --project supabase/storage
fi

# 4. Deploy policies (RLS)
if ls supabase/policies/*.sql 1> /dev/null 2>&1; then
  supabase db execute supabase/policies/*.sql
fi

# 5. Deploy triggers
if ls supabase/triggers/*.sql 1> /dev/null 2>&1; then
  supabase db execute supabase/triggers/*.sql
fi

# 6. Deploy templates, flows, consequences
if [ -d "supabase/templates" ]; then
  npx vv-cli sync-templates supabase/templates/
  npx vv-cli sync-flows supabase/templates/
  npx vv-cli sync-consequences supabase/templates/
fi

# 7. Gera documentação viva
if [ -d "supabase/docs" ]; then
  npx vv-cli docs supabase/docs/
fi

echo "✅ Deploy completo!"