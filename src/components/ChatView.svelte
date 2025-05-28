<script lang="ts">
  import { onMount, afterUpdate } from 'svelte';
  import MessageBubble from './MessageBubble.svelte';
  export let messages: any[] = [];
  let chatContainer: HTMLDivElement;

  afterUpdate(() => {
    if (chatContainer) chatContainer.scrollTop = chatContainer.scrollHeight;
  });
</script>
<div bind:this={chatContainer} class="chat-view">
  {#each messages as message}
    <MessageBubble {message} />
  {/each}
</div>