<script lang="ts">
  export let message: {
    who: string,
    text: string,
    status?: string, // 'loading', 'valid', 'error'
    isUser?: boolean
  };
</script>
<div class="bubble" class:user={message.isUser} class:error={message.status === 'error'} class:loading={message.status === 'loading'}>
  <b>{message.who}</b>: {message.text}
</div>