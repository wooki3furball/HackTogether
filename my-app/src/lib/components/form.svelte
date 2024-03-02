<script>
    import { onMount } from 'svelte';
    import { navigate } from 'svelte-routing';
    import Form from '$lib/components/form.svelte';
  
    let size = '';
    let fragrance = '';
    let length = '';
    let width = '';
    let calluses = false;
    let callusCount = '';
    let age = '';
    let bunion = false;
    let preferences = false;
  
    onMount(() => {
      // You can perform any initialization here if needed
    });
  
    const handleSubmit = () => {
      // Perform actions when the form is submitted
      // For now, you can log the input values
      console.log({
        size,
        fragrance,
        length,
        width,
        calluses,
        callusCount,
        age,
        bunion,
        preferences
      });
  
      // Redirect or perform other actions as needed
      navigate('/success'); // Redirect to a success page, for example
    };
</script>
  
  <style>
    body {
    background-color: #ce6565;
    font-family: 'Arial', sans-serif;
    display: flex;
    align-items: center;
    justify-content: center;
    height: 100vh;
    margin: 0;
  }

  form {
    background-color: #ffc0cb; /* Pink background color */
    border-radius: 8px;
    box-shadow: 0 0 20px rgba(0, 0, 0, 0.2); /* Subtle box shadow */
    padding: 20px;
    width: 1525px;
    height: 1000px;
  }

  h3 {
    font-size: 20px;
    text-align: center;
    margin-bottom: 20px;
  }

  label {
    display: block;
    font-size: 14px;
    margin-top: 10px;
  }

  input,
  select,
  button {
    width: 100%;
    padding: 10px;
    margin-top: 5px;
    margin-bottom: 15px;
    font-size: 14px;
    border: 1px solid #ddd;
    border-radius: 4px;
    box-sizing: border-box;
  }

  button {
    background-color: #4caf50;
    color: #fff;
    cursor: pointer;
  }

  button:hover {
    background-color: #45a049;
  }
  </style>
  
  <form on:submit|preventDefault={handleSubmit}>
    <h3>Feet Information Form</h3>
  
    <label for="size">Size</label>
    <input type="text" bind:value={size} id="size" required>
  
    <label for="fragrance">Fragrance</label>
    <select bind:value={fragrance} id="fragrance" required>
      <option value="Cheesy">Cheesy</option>
      <option value="Fungus Amongus">Fungus Amongus</option>
      <option value="Sweet">Sweet</option>
      <option value="Floral">Floral</option>
    </select>
  
    <label for="length">Length</label>
    <input type="text" bind:value={length} id="length" required>
  
    <label for="width">Width</label>
    <input type="text" bind:value={width} id="width" required>
  
    <label for="calluses">Calluses</label>
    <input type="checkbox" bind:checked={calluses} id="calluses">
  
    {#if calluses}
      <label for="callusCount">Callus Count</label>
      <input type="text" bind:value={callusCount} id="callusCount">
    {/if}
  
    <label for="age">Age</label>
    <select bind:value={age} id="age" required>
      <option value="18-25">18-25</option>
      <option value="26-30">26-30</option>
      <option value="31-36">31-36</option>
      <option value="36-40">36-40</option>
      <option value="41-50">41-50</option>
      <option value="51-59">51-59</option>
      <option value="60+">60+</option>
    </select>
<!--   
    <label for="bunion">Bunion</label>
    <input type="checkbox" bind:checked={bunion} id="bunion">
  
    <label for="preferences">Preferences</label>
    <input type="checkbox" bind:checked={preferences} id="preferences"> -->
  
    <button type="submit">Submit</button>
  </form>
  