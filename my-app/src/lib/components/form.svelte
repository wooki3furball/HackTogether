<script>
    let projectName = "";
    let projectDescription = "";
    let interestList = "";
    let numberOfSpans = 64;
  
    const handleSubmit = async () => {
      // Implement your login logic here (e.g., sending data to server)
      console.log("Project Name:", projectName);
      console.log("Project Description:", projectDescription);
      console.log("List of Interests:", interestList);
  
      // Replace this with your actual login functionality (e.g., sending data to server)
      alert("Form submitted!");
    };
  </script>
  
  <style>
    @import url('https://fonts.googleapis.com/css2?family=Quicksand:wght@300;400;500;600;700&display=swap');
  
    * {
      margin: 0;
      padding: 0;
      box-sizing: border-box;
      font-family: 'Quicksand', sans-serif;
    }
  
    body {
      display: flex;
      justify-content: center;
      align-items: center;
      min-height: 100vh;
      background: #000;
    }
  
    section {
      position: absolute;
      width: 100vw;
      height: 100vh;
      display: flex;
      justify-content: center;
      align-items: center;
      gap: 2px;
      flex-wrap: wrap;
      overflow: hidden;
      z-index: 1; /* Set a higher z-index for the login form to appear on top */
    }
  
    section::before {
      content: '';
      position: absolute;
      width: 100%;
      height: 100%;
      background: linear-gradient(#000, #0f0, #000);
      animation: animate 5s linear infinite;
    }
  
    @keyframes animate {
      0% {
        transform: translateY(-100%);
      }
      100% {
        transform: translateY(100%);
      }
    }
  
    section span {
      position: relative;
      display: block;
      width: calc(6.25vw - 2px);
      height: calc(6.25vw - 2px);
      background: #181818;
      z-index: 2;
      transition: 1.5s;
    }
  
    section span:hover {
      background: #0f0;
      transition: 0s;
    }
  
    form {
      z-index: 3; 
    }
  
    form {
        position: absolute;
        top: 50%;
        left: 50%;
        transform: translate(-50%, -50%); /* Center the form relative to its position */
        width: 1000px;
        height: 700px;
        background: #222;
        z-index: 1000;
        display: flex;
        justify-content: center;
        align-items: center;
        padding: 40px;
        border-radius: 4px;
        box-shadow: 0 15px 35px rgba(0, 0, 0, 0.9);
        flex-direction: column;
        gap: 25px;
    }

  
    form h3 {
      font-size: 2em;
      color: #0f0;
      text-transform: uppercase;
    }
  
    form label {
      color: #fff;
    }
  
    form input {
      position: relative;
      width: 100%;
      background: #333;
      border: none;
      outline: none;
      padding: 15px 10px;
      border-radius: 4px;
      color: #fff;
      font-weight: 500;
      font-size: 1em;
    }
  
    form button {
      padding: 10px;
      background: #0f0;
      color: #000;
      font-weight: 600;
      font-size: 1.35em;
      letter-spacing: 0.05em;
      cursor: pointer;
    }
  
    form button:active {
      opacity: 0.6;
    }
  
    @media (max-width: 900px) {
      section span {
        width: calc(10vw - 2px);
        height: calc(10vw - 2px);
      }
    }
  
    @media (max-width: 600px) {
      section span {
        width: calc(20vw - 2px);
        height: calc(20vw - 2px);
      }
    }
  </style>
  
  <section>
    {#each Array(numberOfSpans) as _, index (index)}
        <span></span>
    {/each}
  </section>
  
  <form on:submit|preventDefault={handleSubmit}>
    <h3>Hack Together Form</h3>
    
    <label for="projectName">Project Name</label>
    <input type="text" placeholder="Project Name" bind:value={projectName} id="projectName" required>

    <label for="projectDescription">Project Description</label>
    <input type="text" placeholder="Enter Project Description" bind:value={projectDescription} id="projectDescription" required>

    <label for="interestList">Interest List</label>
    <input type="text" placeholder="List your interests" bind:value={interestList} id="interestList">

    
    <button type="submit">Log In</button>
  </form>
  