<script>
    import { fade } from 'svelte/transition';

    let projects = [
        { 
            id: 1, 
            name: "EcoFinder",
            description: "A platform to connect individuals with recycling centers and eco-friendly businesses.",
            languages: "JavaScript, HTML, CSS, Node.js"
        },
        { 
            id: 2, 
            name: "HealthTracker",
            description: "An app to monitor health parameters and suggest fitness plans.",
            languages: "Kotlin, Swift, Firebase"
        },
        { 
            id: 3, 
            name: "BudgetBuddy",
            description: "Personal finance manager to track and plan your budget.",
            languages: "Python, Django, React"
        },
        { 
            id: 4, 
            name: "CodeCollab",
            description: "A platform for developers to collaborate on projects in real-time.",
            languages: "JavaScript, Node.js, Socket.io"
        },
        { 
            id: 5, 
            name: "SafeRoute",
            description: "Application to find the safest routes for commuting based on real-time data.",
            languages: "Java, Google Maps API"
        },
        { 
            id: 6, 
            name: "ArtGallery",
            description: "Virtual gallery for artists to showcase and sell their artwork.",
            languages: "PHP, Laravel, Vue.js"
        },
        { 
            id: 7, 
            name: "StudySpace",
            description: "Platform for students to find study groups and resources.",
            languages: "Ruby on Rails, React"
        },
        { 
            id: 8, 
            name: "FreelanceHub",
            description: "Marketplace for freelancers to offer services and for clients to find quality talent.",
            languages: "Python, Flask, Angular"
        },
        { 
            id: 9, 
            name: "FoodieFinder",
            description: "App to discover local restaurants and cuisine based on user preferences.",
            languages: "Swift, Yelp API"
        },
        { 
            id: 10, 
            name: "EventMaster",
            description: "Event management tool for creating, promoting, and managing events.",
            languages: "C#, .NET, React"
        }
    ];

    let currentIndex = 0;
    let endOfPics = false; // Adjusted variable name for clarity

    function navigate(){
		currentIndex += 1;
        if (currentIndex >= projects.length - 1) {
            endOfPics = true;
        }
    }

    function like() {
        navigate();
    }

    function dislike() {
        navigate();
    }
</script>


<style>
    /* Adjusting container for overall theme */
    .container {
        display: flex;
        flex-direction: column;
        align-items: center;
        min-height: 100vh;
		min-width: 180vh;
        background-color: #181818; /* Dark background */
        color: #333;
        font-family: 'Helvetica Neue', Arial, sans-serif;
        margin: 0;
        padding: 0;
    }

    /* Styling for project details */
    .project-details {
        background-color: #ffffff; /* Slightly lighter than the container for depth */
        color: #181818; /* Soft green for text */
        padding: 20px;
        margin-bottom: 20px;
        border-radius: 10px;
        box-shadow: 0 2px 10px rgba(0, 0, 0, 0.2); /* Soft shadow for depth */
		height: 500px;
        width: 20%; /* Adjust based on preference */
    }

    h2 {
        color: #87b245; /* Brighter green for headings to stand out */
    }

    button {
        padding: 10px 20px;
        margin: 0 10px;
        border: none;
        border-radius: 5px;
        background-color: #0f0; /* Theme green */
        color: #181818; /* Dark text for contrast */
        font-size: 1rem;
        cursor: pointer;
        transition: background-color 0.3s, transform 0.2s;
    }

    button:hover {
        background-color: #006400; /* Darker green on hover */
    }

    button:active {
        transform: scale(0.98); /* Slightly scale down when clicked */
    }

    .button-group {
        display: flex;
    }
</style>


<div class="container">
    {#if projects.length > 0 && !endOfPics}
        <div class="project-details">
            <h2>{projects[currentIndex].name}</h2>
            <p>{projects[currentIndex].description}</p>
            <p>Technologies: {projects[currentIndex].languages}</p>
        </div>
        <div class="button-group">
            <button on:click={like}>Like</button>
            <button on:click={dislike}>Dislike</button>
        </div>
    {:else if endOfPics}
        <p style="color: #8FBC8F;">No more projects to show.</p>
    {/if}
</div>