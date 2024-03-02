<script lang="js">
	import Navbar from '$lib/components/navbar.svelte';
  import Testimonial from '$lib/components/testimonial.svelte';
	import lady from '$lib/assets/img/lady-img.jpg';
	import demo from '$lib/assets/img/Demo.png';
	import iphone6 from '$lib/assets/img/iphone6.png';
	import dogImg from '$lib/assets/img/dog-img.jpg';
</script>

<Navbar />
<Testimonial />

<section id="pricing">
	<h2>A Plan for Every Dog's Needs</h2>
	<p>Simple and affordable price plans for your dog.</p>

	<div class="row">
		<div class="pricing-column col-lg-4 col-md-6">
			<div class="card">
				<div class="card-header">
					<h3>Chihuahua</h3>
				</div>
				<div class="card-body">
					<h2>Free</h2>
					<p>5 Matches per Day</p>
					<p>10 Matches per Day</p>
					<p>Unlimited App Usage</p>
					<button type="button" class="btn btn-lg btn-block btn-dark"> Sign Up </button>
				</div>
			</div>
		</div>

		<div class="pricing-column col-lg-4 col-md-6">
			<div class="card">
				<div class="card-header">
					<h3>Labrador</h3>
				</div>
				<div class="card-body">
					<h2>$49 / month</h2>
					<p>Unlimited Matches</p>
					<p>Unlimited Messages</p>
					<p>Unlimited App Usage</p>
					<button type="button" class="btn btn-lg btn-block btn-dark"> Sign Up </button>
				</div>
			</div>
		</div>

		<div class="pricing-column col-lg-4">
			<div class="card">
				<div class="card-header">
					<h3>Mastiff</h3>
				</div>
				<div class="card-body">
					<h2>$99 / month</h2>
					<p>Priority Listing</p>
					<p>Unlimited Messages</p>
					<p>Unlimited App Usage</p>
					<button type="button" class="btn btn-lg btn-block btn-dark"> Sign Up </button>
				</div>
			</div>
		</div>
	</div>
</section>

<section id="cta">
	<h4 id="footer-text">Find the True Love of Your Dog's Life Today.</h4>
	<button type="button" class="btn btn-dark btn-lg download-button">
		<i class="fab fa-apple"></i> Download
	</button>
	<button type="button" class="btn btn-light btn-lg download-button">
		<i class="fab fa-google-play"></i> Download
	</button>
</section>

<!-- Footer -->

<!-- <footer id="footer">
    <a target="_blank" href="#" style="color: #000;"><i class="fab footer-icons fa-twitter"></i></a>
    <a target="_blank" href="#" style="color: #000;"><i class="fab footer-icons fa-facebook-f"></i></a>
    <a target="_blank" href="#" style="color: #000;"><i class="fab footer-icons fa-instagram"></i></a>
    <a target="_blank" href="#" style="color: #000;"><i class="fab footer-icons fa-linkedin-in"></i></i></a>
    <p id="copyright-text">© Copyright 2023 TinDog</p>
    <p id="trademark-text">
      Made with <i class="fas heart-icon fa-heart"></i> and
      <a href="#"><i class="fab boot-icon fa-bootstrap"></i></a> by Niladri Chatterjee.
    </p>
  </footer> -->

<style>
	body {
		font-family: 'Montserrat', sans-serif;
	}

	/* section id=#title */
	#title {
		background-color: #ff4c68;
		color: #fff;
	}

	h1 {
		font-family: 'Montserrat', sans-serif;
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
		font-family: 'Ubuntu', sans-serif;
		font-size: 2.5rem;
		font-weight: bold;
	}

	.nav-item {
		padding: 0 18px;
	}

	.nav-link {
		font-size: 1.2rem;
		font-family: 'Montserrat', sans-serif;
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

	/* Feature Section */

	.features-icon {
		font-size: 45px;
		color: #ef8172;
		margin-bottom: 1rem;
	}

	#features {
		padding: 7% 15%;
		background-color: #fff;
		position: relative;
		/* z-index: 1; */
	}

	.feature-box {
		text-align: center;
		padding: 5%;
	}

	h3 {
		font-family: 'Montserrat', sans-serif;
		font-weight: 900;
		font-size: 1.45rem;
	}

	p {
		color: #8f8f8f;
	}

	.features-icon:hover {
		color: #ff4c68;
	}

	/* Pricing */
	#pricing {
		padding: 100px;
		text-align: center;
	}

	.pricing-column {
		padding: 3% 2%;
	}

	@media (max-width: 1028px) {
		#title {
			text-align: center;
		}

		.iphone-image {
			position: static;
			transform: rotate(0deg);
		}
	}

	/* CTA Section */
	#cta {
		background-color: #ef8172;
		text-align: center;
		color: #fff;
		padding: 4% 15%;
	}

	#footer-text {
		font-family: 'Montserrat', sans-serif;
		font-size: 3.5rem;
		font-weight: 900;
		line-height: 1.5;
	}

	#footer {
		text-align: center;
		padding: 8% 20% 3%;
	}

	p#copyright-text {
		color: #8f8f8f;
		padding: 1% 0;
		margin-bottom: 0;
	}

	p#trademark-text {
		color: #8f8f8f;
		padding: 0 2%;
	}

	.footer-icons {
		padding: 1% 1.4%;
	}

	.heart-icon {
		color: #ff4c68;
	}

	.boot-icon {
		color: #563d7c;
	}
</style>
