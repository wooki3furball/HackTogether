
<script lang="js">
     import iphone6 from "$lib/assets/img/iphone6.png";
</script>
<section id="title">
    <div class="container-fluid">
      <!-- Nav Bar -->
      <nav class="navbar navbar-expand-lg navbar-dark">
        <!-- Hamburger Menu Button -->
        <button class="navbar-toggler" type="button" data-toggle="collapse" data-target="#navbarSupportedContent"
          aria-controls="navbarSupportedContent" aria-expanded="false" aria-label="Toggle navigation">
          <span class="navbar-toggler-icon"></span>
        </button>
        <a class="navbar-brand" href="">tindog</a>

        <div class="collapse navbar-collapse" id="navbarSupportedContent">
          <ul class="navbar-nav ml-auto">
            <li class="nav-item">
              <a class="nav-link" href="#footer">Contact</a>
            </li>
            <li class="nav-item">
              <a class="nav-link" href="#pricing">Pricing</a>
            </li>
            <li class="nav-item">
              <a class="nav-link" href="#cta">Download</a>
            </li>
          </ul>
        </div>
      </nav>

      <!-- Title -->
      <div class="row">
        <div class="col-lg-6">
          <h1>Meet new and interesting dogs nearby.</h1>
          <button type="button" class="btn btn-dark btn-lg download-button">
            <i class="fab fa-apple"></i> Download
          </button>
          <button type="button" class="btn btn-outline-light btn-lg download-button">
            <i class="fab fa-google-play"></i> Download
          </button>
        </div>

      </div>
    </div>
  </section>

  <style>
    #title {
        background-color: #ff4c68;
        color: #fff;
        }

        h1 {
        font-family: "Montserrat", sans-serif;
        font-size: 3.5rem;
        font-weight: 900;
        line-height: 1.5;
        }

        .container-fluid {
        padding: 3% 15% 7%;
        }

        /* Navigation Bar */
        .navbar {
        padding: 0 0 4.5rem;
        }

        .navbar-brand {
        font-family: "Ubuntu", sans-serif;
        font-size: 2.5rem;
        font-weight: bold;
        }

        .nav-item {
        padding: 0 18px;
        }

        .nav-link {
        font-size: 1.2rem;
        font-family: "Montserrat", sans-serif;
        font-weight: 300;
        }

        /* Download Buttons */
        .download-button {
        margin: 5% 3% 5% 0;
        }

        .iphone-image {
        width: 60%;
        transform: rotate(25deg);
        position: absolute;
        right: 30%;
        }
  </style>