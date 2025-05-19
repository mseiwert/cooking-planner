<script lang="ts">
  import { onMount } from 'svelte'

  let inputValue = "nix";
  
  onMount(async () => {
    const request = await fetch(`http://localhost:8080/api/v1/recipes`);

    if (request.status !== 200) {
      console.error("Error fetching data");
      return;
    }

    const recipe = await request.json();
    inputValue = recipe.name;
    
  });
</script>

<main>
  <input type='text' bind:value={inputValue} />
</main>
