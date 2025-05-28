<script lang="ts">
  export let loglines: any[] = [];
</script>

<div class="timeline">
  {#each loglines as log}
    <!-- Renderização do log -->
  {/each}
</div>