<script lang="ts">
  import { onMount } from 'svelte';
  import { createClient } from '@supabase/supabase-js';

  const supabase = createClient(import.meta.env.VITE_SUPABASE_URL, import.meta.env.VITE_SUPABASE_ANON_KEY);
  export let user: string = "danvoulez";

  // Mock de urgência master
  let urgenciaMaster = {
    existe: true,
    titulo: "Servidor financeiro FORA DO AR",
    quem: "danvoulez",
    detalhe: "Nenhuma operação será registrada até resolver",
    cor: "#ff002f",
    icone: "🛑",
    acao: "Acesse o dashboard de incidentes imediatamente!"
  };

  let radar = [
    { tipo: "emergência", titulo: "Assinatura urgente", quem: "danvoulez", cor: "#dc2626", icone: "🚨" },
    { tipo: "urgência", titulo: "Minicontrato expira hoje", quem: "alice", cor: "#f43f5e", icone: "⚡" }
  ];
  let alertas = [
    { msg: "Você tem 2 contratos aguardando testemunha", cor: "#fde68a", icone: "🕒" }
  ];
  let kpis = [
    { label: "Minicontratos nos últimos 7 dias", valor: 12, cor: "#38bdf8", icone: "📈" },
  ];
</script>

<div class="painel-lateral-ia">
  <!-- URGÊNCIA MASTER -->
  {#if urgenciaMaster.existe}
    <div class="urgencia-master" style="background:{urgenciaMaster.cor}; animation: shake 0.7s infinite alternate;">
      <span class="urgencia-master-icone">{urgenciaMaster.icone}</span>
      <div class="urgencia-master-conteudo">
        <h2 class="urgencia-master-titulo">{urgenciaMaster.titulo}</h2>
        <div class="urgencia-master-detalhe">{urgenciaMaster.detalhe}</div>
        <div class="urgencia-master-quem">Responsável: <b>{urgenciaMaster.quem}</b></div>
        <button class="urgencia-master-acao">{urgenciaMaster.acao}</button>
      </div>
    </div>
  {/if}

  <!-- RESTANTE DO PAINEL (normal) -->
  <section class="quente">
    <h3 class="painel-titulo">Radar de Emergência</h3>
    <div class="painel-cards">
      {#each radar as u}
        <div class="painel-card" style="background:{u.cor}">
          <span class="painel-icone">{u.icone}</span>
          <div>
            <strong>{u.titulo}</strong>
            <div class="painel-quem">({u.quem})</div>
          </div>
        </div>
      {/each}
    </div>
  </section>
  <section class="moderado">
    <h4 class="painel-subtitulo">Alertas & Sugestões</h4>
    <div class="alertas-cards">
      {#each alertas as a}
        <div class="alerta-card" style="background:{a.cor}">
          <span>{a.icone}</span> {a.msg}
        </div>
      {/each}
    </div>
  </section>
  <section class="frio">
    <h4 class="painel-subtitulo">Métricas</h4>
    <div class="kpi-cards">
      {#each kpis as k}
        <div class="kpi-card" style="background:{k.cor}">
          <span>{k.icone}</span>
          <div>
            <div class="kpi-valor">{k.valor}</div>
            <div class="kpi-label">{k.label}</div>
          </div>
        </div>
      {/each}
    </div>
  </section>
</div>

<style>
.painel-lateral-ia {
  width: 340px;
  max-width: 100vw;
  padding: 1.2rem 0.3rem 1.2rem 1rem;
  background: linear-gradient(185deg, #fff 60%, #f1f5f9 100%);
  border-left: 2px solid #f1f1f1;
  box-shadow: 0 8px 32px #dc262622;
  border-radius: 24px 0 0 24px;
  overflow-y: auto;
  min-height: 100vh;
  display: flex;
  flex-direction: column;
  gap: 2.2rem;
}

.urgencia-master {
  display: flex;
  align-items: flex-start;
  gap: 1.1rem;
  padding: 2rem 1.2rem 1.6rem 1.2rem;
  border-radius: 1.5rem;
  box-shadow: 0 6px 40px 0 #dc262644, 0 2px 8px #f43f5ecc;
  margin-bottom: 1.7rem;
  color: #fff;
  position: relative;
  z-index: 100;
  border: 3px solid #fff;
  animation: blink 1.2s infinite alternate;
}
.urgencia-master-icone {
  font-size: 3.2rem;
  margin-top: 0.5rem;
  filter: drop-shadow(0 0 12px #fff6);
}
.urgencia-master-conteudo {
  flex: 1;
  display: flex;
  flex-direction: column;
  gap: 0.5rem;
}
.urgencia-master-titulo {
  font-size: 1.6rem;
  font-weight: 900;
  letter-spacing: -0.01em;
  margin-bottom: 0.18rem;
  text-shadow: 1px 2px 10px #ffec;
}
.urgencia-master-detalhe {
  font-size: 1.07rem;
  font-weight: 500;
  margin-bottom: 0.15rem;
}
.urgencia-master-quem {
  font-size: 1.04rem;
  color: #fff9;
  margin-bottom: 0.32rem;
}
.urgencia-master-acao {
  background: linear-gradient(90deg, #fff 20%, #ffb3c6 100%);
  color: #dc2626;
  font-weight: 800;
  border: none;
  border-radius: 999px;
  padding: 0.65rem 1.7rem;
  font-size: 1.1rem;
  margin-top: 0.5rem;
  box-shadow: 0 2px 16px #fff6;
  cursor: pointer;
  text-shadow: 0px 1px 5px #fff7;
  animation: pulse 0.9s infinite alternate;
}
@keyframes shake {
  0% { transform: translateX(0px); }
  30% { transform: translateX(-3px); }
  60% { transform: translateX(4px);}
  100% { transform: translateX(0px);}
}
@keyframes blink {
  from { box-shadow: 0 0 18px 8px #fff4, 0 6px 40px #dc262644; }
  to   { box-shadow: 0 0 38px 18px #fff9, 0 6px 60px #ff002f55; }
}
@keyframes pulse {
  from { filter: brightness(1.1) drop-shadow(0 0 8px #fff6);}
  to   { filter: brightness(1.35) drop-shadow(0 0 28px #fff9);}
}

/* ...restante do CSS igual ao exemplo anterior para radar, alertas, kpis... */
.painel-titulo { color: #dc2626; font-size: 1.12rem; font-weight: 700; margin-bottom: 0.7rem; letter-spacing: -0.01em; }
.painel-subtitulo { color: #f43f5e; font-size: 1.01rem; font-weight: 600; margin-bottom: 0.5rem; }
.painel-cards { display: flex; flex-direction: column; gap: 0.55rem; }
.painel-card { display: flex; align-items: center; gap: 0.7rem; padding: 0.9rem 1.1rem; border-radius: 1rem; color: #fff; font-weight: 600; font-size: 1.05rem; box-shadow: 0 1px 10px #0002; transition: transform 0.2s; }
.painel-card:hover { transform: scale(1.03);}
.painel-icone { font-size: 1.38rem; margin-right: 0.24rem; }
.painel-quem { font-size: 0.97rem; color: #fef9c3; }
.alertas-cards { display: flex; flex-direction: column; gap: 0.42rem; }
.alerta-card { display: flex; align-items: center; gap: 0.7rem; background: #fde68a; color: #92400e; border-radius: 0.9rem; padding: 0.65rem 1rem; font-size: 1.01rem; font-weight: 500; box-shadow: 0 1px 6px #0001; }
.kpi-cards { display: flex; flex-direction: column; gap: 0.42rem; }
.kpi-card { display: flex; align-items: center; gap: 0.7rem; background: #e0e7ef; color: #22223b; border-radius: 0.95rem; padding: 0.64rem 1rem; font-size: 1.04rem; box-shadow: 0 1px 7px #0001; }
.kpi-valor { font-size: 1.15rem; font-weight: 700; color: #0ea5e9; }
.kpi-label { font-size: 0.96rem; color: #64748b; }
@media (max-width: 600px) {
  .painel-lateral-ia { width: 100vw; border-radius: 0; }
}
</style>