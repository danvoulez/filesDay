import { serve } from "https://deno.land/std@0.168.0/http/server.ts";

serve(async (req) => {
  // Exemplo: registra uma LogLine no banco
  const body = await req.json();
  const { who, did, this: data } = body;

  // Aqui você integraria com Supabase Client (exemplo simplificado)
  // await supabase.from('logline_global').insert({ who, did, this: data });

  return new Response(JSON.stringify({ status: "ok", received: body }));
});