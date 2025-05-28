<script lang="ts">
  import { createClient } from '@supabase/supabase-js';
  import { fly, fade } from 'svelte/transition';

  const supabase = createClient(
    import.meta.env.VITE_SUPABASE_URL,
    import.meta.env.VITE_SUPABASE_ANON_KEY
  );

  let user = $page?.data?.user?.email || 'anon';
  let contractId = null;
  let contract = {
    who: user,
    did: "",
    when: "",
    witness: "",
    obs: ""
  };

  const questions = [
    {
      label: "Quem está registrando este acontecimento?",
      options: [user],
      field: "who"
    },
    {
      label: "O que aconteceu?",
      options: ["Venda", "Compra", "Pagamento", "Reunião", "Outro"],
      field: "did"
    },
    {
      label: "Quando aconteceu?",
      options: ["Agora", "Hoje", "Ontem", "Escolher Data"],
      field: "when"
    },
    {
      label: "Quem é a testemunha?",
      options: ["dan", "alice", "bob", "Outro"],
      field: "witness"
    },
    {
      label: "Alguma observação ou detalhe extra?",
      options: [],
      field: "obs"
    }
  ];

  let step = 0;
  let customInput = "";

  async function saveStep(field, value) {
    if (field === 'when') {
      if (value === "Agora") value = new Date().toISOString();
      else if (value === "Hoje") value = new Date().toISOString().split('T')[0];
      else if (value === "Ontem") {
        const d = new Date();
        d.setDate(d.getDate() - 1);
        value = d.toISOString().split('T')[0];
      }
    }
    contract[field] = value;
    if (!contractId) {
      const { data } = await supabase.from('logline').insert([{
        who: contract.who,
        did: contract.did,
        this: contract,
        when: contract.when || new Date().toISOString(),
        witness: contract.witness,
        status: 'ghost',
        valid: false
      }]).select();
      contractId = data?.[0]?.id;
    } else {
      await supabase.from('logline')
        .update({
          this: contract,
          [field]: value,
          status: 'ghost',
          valid: false
        })
        .eq('id', contractId);
    }
    customInput = "";
    step += 1;
  }

  function choose(option) {
    saveStep(questions[step].field, option);
  }
  function nextWithCustom() {
    if (customInput) choose(customInput);
  }

  async function submit() {
    await supabase.from('logline')
      .update({
        this: contract,
        status: 'pending',
        valid: true
      })
      .eq('id', contractId);
    step += 1;
  }
</script>

<div class="wallet-bg">
  <div class="wallet-stack">
    {#if step < questions.length}
      <div class="wallet-card" transition:fly="{{ x: 0, y: 40, duration: 350 }}" transition:fade>
        <h2 class="wallet-question">{questions[step].label}</h2>
        <div class="wallet-options">
          {#each questions[step].options as opt}
            <button class="wallet-chip" on:click={() => choose(opt)}>{opt}</button>
          {/each}
        </div>
        {#if !questions[step].options.length || questions[step].field === "obs" || (questions[step].field === "witness" && questions[step].options.includes("Outro"))}
          <input
            type="text"
            class="wallet-input"
            placeholder="Digite sua resposta..."
            bind:value={customInput}
            on:keydown={(e) => e.key === 'Enter' && nextWithCustom()}
          />
          <button class="wallet-chip" on:click={nextWithCustom} disabled={!customInput}>OK</button>
        {/if}
      </div>
    {:else if step === questions.length}
      <div class="wallet-card" transition:fly="{{ x: 0, y: 40, duration: 350 }}" transition:fade>
        <h2 class="wallet-question">Resumo</h2>
        <pre class="wallet-summary">{JSON.stringify(contract, null, 2)}</pre>
        <button class="wallet-final" on:click={submit}>Finalizar minicontrato</button>
      </div>
    {:else}
      <div class="wallet-card wallet-success" transition:fade>
        🎉 Minicontrato registrado e validado!
      </div>
    {/if}
  </div>
</div>

<style>
.wallet-bg {
  min-height: 100vh;
  background: linear-gradient(160deg,#f8fafc 0%,#e0e7ef 100%);
  display: flex;
  align-items: center;
  justify-content: center;
}
.wallet-stack {
  width: 100%;
  max-width: 420px;
  min-height: 450px;
  display: flex;
  flex-direction: column;
  align-items: center;
}
.wallet-card {
  background: #fff;
  border-radius: 24px;
  box-shadow: 0 16px 40px 0 #0001, 0 2px 8px #d1d5dbcc;
  padding: 2.2rem 2rem 2.2rem 2rem;
  width: 100%;
  margin-bottom: 1.4rem;
  display: flex;
  flex-direction: column;
  align-items: stretch;
  animation: walletCardPop 0.5s;
}
@keyframes walletCardPop {
  0% { transform: scale(0.95) translateY(32px); opacity: 0.3; }
  100% { transform: scale(1) translateY(0); opacity: 1; }
}
.wallet-question {
  font-size: 1.35rem;
  font-weight: 700;
  margin-bottom: 1.2rem;
  color: #22223b;
  letter-spacing: -0.01em;
}
.wallet-options {
  display: flex;
  flex-wrap: wrap;
  gap: 0.6rem;
  margin-bottom: 1.2rem;
}
.wallet-chip {
  background: #f1f5f9;
  border: none;
  border-radius: 9999px;
  padding: 0.7rem 1.4rem;
  font-size: 1.04rem;
  margin-right: 0.4rem;
  cursor: pointer;
  font-weight: 500;
  box-shadow: 0 1px 6px #0001;
  transition: background 0.2s;
}
.wallet-chip:hover, .wallet-chip:focus {
  background: #dbeafe;
}
.wallet-input {
  border: 1.5px solid #cbd5e1;
  border-radius: 12px;
  padding: 0.6rem 1rem;
  font-size: 1.08rem;
  margin-bottom: 1rem;
  outline: none;
}
.wallet-final {
  background: linear-gradient(90deg,#3b82f6 0%,#6366f1 100%);
  color: #fff;
  padding: 0.95rem 2.2rem;
  font-weight: bold;
  border-radius: 9999px;
  border: none;
  cursor: pointer;
  font-size: 1.14rem;
  margin-top: 0.7rem;
  box-shadow: 0 4px 18px #0001;
}
.wallet-summary {
  background: #f1f5f9;
  border-radius: 12px;
  padding: 1.4rem;
  margin-bottom: 1.2rem;
  font-size: 1.01rem;
  color: #374151;
}
.wallet-success {
  color: #16a34a;
  font-size: 1.27rem;
  margin-top: 2rem;
  text-align: center;
  background: linear-gradient(90deg,#dcfce7 10%,#f0fdf4 100%);
}
</style>