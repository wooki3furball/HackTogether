<script>
  import { navigate } from 'svelte-routing';

  let name = '';
  let username = '';
  let password = '';
  let confirmPassword = '';
  let error = '';

  const handleSignup = (event) => {
    event.preventDefault();

    // Basic validation
    if (!name || !username || !password || !confirmPassword) {
      error = 'Please fill in all required fields.';
      return;
    }

    if (password !== confirmPassword) {
      error = 'Passwords do not match.';
      return;
    }

    // Handle signup logic here (e.g., send data to server)
    console.log('Signing up...');
    error = ''; // Clear error on successful submission
  };
</script>

<style>
  *,
*:before,
*:after {
  padding: 0;
  margin: 0;
  box-sizing: border-box;
}

body {
  /* background-image: url("my-app/src/lib/assets/img/background.png"); */
  background-color: #ff512f;
  background-size: cover;
  background-position: center;
  background-repeat: no-repeat;
  height: 100vh;
  margin: 0;
  display: flex;
  align-items: center;
  justify-content: center;
}

.background {
  width: 430px;
  height: 520px;
  position: absolute;
  transform: translate(-50%, -50%);
  left: 50%;
  top: 50%;
}

.background .shape {
  height: 200px;
  width: 200px;
  position: absolute;
  border-radius: 50%;
}

.shape:first-child {
  background: linear-gradient(
    #1845ad,
    #23a2f6
  );
  left: -80px;
  top: -80px;
}

.shape:last-child {
  background: linear-gradient(
    to right,
    #ff512f,
    #f09819
  );
  right: -30px;
  bottom: -80px;
}

form {
  height: 600px;
  width: 400px;
  background-color: rgba(255, 255, 255, 0.13);
  position: absolute;
  transform: translate(-50%, -50%);
  top: 50%;
  left: 50%;
  border-radius: 10px;
  backdrop-filter: blur(10px);
  border: 2px solid rgba(255, 255, 255, 0.1);
  box-shadow: 0 0 40px rgba(8, 7, 16, 0.6);
  padding: 50px 35px;
}

form * {
  font-family: 'Poppins', sans-serif;
  color: #0c0000;
  letter-spacing: 0.5px;
  outline: none;
  border: none;
}

form h3 {
  font-size: 32px;
  font-weight: 500;
  line-height: 42px;
  text-align: center;
}

label {
  display: block;
  margin-top: 30px;
  font-size: 16px;
  font-weight: 500;
}

input {
  display: block;
  height: 50px;
}

button {
    margin-top: 10px;
    width: 100%;
    background-color: #4285f4; /* Set your desired background color */
    color: #fff; /* Set text color */
    padding: 15px 0;
    font-size: 18px;
    font-weight: 600;
    border-radius: 5px;
    cursor: pointer;
    text-align: center;
  }
</style>

<div class="background">
  <div class="shape"></div>
  <div class="shape"></div>
</div>

<form on:submit|preventDefault={handleSignup}>
  <h3>Sign Up Here</h3>

  {#if error}
    <p style="color: red;">{error}</p>
  {/if}

  <label for="name">First Name</label>
  <input type="text" placeholder="Enter Name" bind:value={name} id="name" required>

  <label for="username">Username</label>
  <input type="text" placeholder="Username" bind:value={username} id="username" required>

  <label for="password">Password</label>
  <input type="password" placeholder="Password" bind:value={password} id="password" required>

  <label for="confirmPassword">Confirm Password</label>
  <input type="password" placeholder="Confirm Password" bind:value={confirmPassword} id="confirmPassword" required>

  <button type="submit">Sign Up</button>


</form>
