import { serve } from "https://deno.land/std@0.168.0/http/server.ts";

serve(async (req) => {
  const body = await req.json();
  const { file_id, action, details } = body;

  // Aqui você integraria com Supabase Client (exemplo simplificado)
  // await supabase.from('file_log').insert({ file_id, action, details });

  return new Response(JSON.stringify({ status: "ok", received: body }));
});