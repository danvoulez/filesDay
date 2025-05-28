<script lang="ts">
  import Layout from '$components/Layout.svelte';
  import LeftMenu from '$components/LeftMenu.svelte';
  import RightMenu from '$components/RightMenu.svelte';
</script>

<Layout>
  <LeftMenu />
  <slot />
  <RightMenu />
</Layout>