<script lang="ts">
  import { onMount } from 'svelte';
  import Timeline from '$components/Timeline.svelte';
  let loglines: any[] = [];
  onMount(async () => {
    const res = await fetch('/api/log?tenant=main');
    loglines = await res.json();
  });
</script>