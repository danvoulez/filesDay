<script lang="ts">
  export let log: any;
</script>

<div class="message" class:ghost={log.status === 'ghost'}>
  <b>{log.who}</b>: {log.this?.text}
</div>