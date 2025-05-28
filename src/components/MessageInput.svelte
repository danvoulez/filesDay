<script lang="ts">
  import { createEventDispatcher } from 'svelte';
  let input = '';
  const dispatch = createEventDispatcher();

  function send() {
    if (input.trim()) {
      dispatch('send', { text: input });
      input = '';
    }
  }
</script>
<input bind:value={input} on:keydown={(e) => e.key === 'Enter' && send()} />
<button on:click={send}>Enviar</button>