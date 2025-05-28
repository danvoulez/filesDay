<script lang="ts">
  import { createClient } from '@supabase/supabase-js';
  const supabase = createClient(import.meta.env.VITE_SUPABASE_URL, import.meta.env.VITE_SUPABASE_ANON_KEY);

  export let user = "danvoulez";

  // Exemplo de dados fake. Troque por fetch real do Supabase!
  let urgencias = [
    { tipo: "emergência", titulo: "Assinatura urgente", quem: "dan", cor: "#dc2626" },
    { tipo: "urgência", titulo: "Minicontrato expira hoje", quem: "alice", cor: "#f43f5e" },
    { tipo: "pendência", titulo: "Testemunha não validou", quem: "dan", cor: "#f59e42" }
  ];

  // Função para buscar do Supabase (substituir o mock acima)
  /*
  async function loadUrgencias() {
    let { data, error } = await supabase.from('logline')
      .select('*')
      .or('status.eq.emergency,status.eq.urgent,status.eq.pending')
      .eq('who', user); // ou remover para ver tudo
    urgencias = data.map(l => ({
      tipo: l.status,
      titulo: l.this.did,
      quem: l.who,
      cor: l.status === "emergency" ? "#dc2626"
          : l.status === "urgent" ? "#f43f5e"
          : "#f59e42"
    }));
  }
  */

  // onMount(loadUrgencias);
</script>

<div class="painel-urgencia">
  <h3 class="painel-titulo">Radar de Pendências</h3>
  <div class="painel-cards">
    {#each urgencias as u}
      <div class="painel-card" style="background:{u.cor}">
        <span class="painel-icone">
          {u.tipo === "emergência" ? "🚨"
          : u.tipo === "urgência" ? "⚠️"
          : "⏳"}
        </span>
        <div>
          <strong>{u.titulo}</strong>
          <div class="painel-quem">({u.quem})</div>
        </div>
      </div>
    {/each}
  </div>
</div>

<style>
.painel-urgencia {
  padding: 0.8rem 1rem;
  background: #fff;
  border-radius: 1.2rem;
  box-shadow: 0 2px 12px #dc262633;
  margin-bottom: 2rem;
}
.painel-titulo {
  font-weight: 700;
  font-size: 1.14rem;
  margin-bottom: 1rem;
  color: #be185d;
}
.painel-cards {
  display: flex;
  flex-direction: column;
  gap: .7rem;
}
.painel-card {
  display: flex;
  align-items: center;
  gap: 0.7rem;
  padding: 0.85rem 1.1rem;
  border-radius: 1rem;
  color: #fff;
  font-weight: 600;
  font-size: 1.04rem;
  box-shadow: 0 1px 10px #0001;
}
.painel-icone {
  font-size: 1.3rem;
  margin-right: 0.4rem;
}
.painel-quem {
  font-size: 0.94rem;
  color: #f1f5f9;
}
</style>