<script lang="js">
	import bizinsider from '$lib/assets/img/bizinsider.png';
	import mashable from '$lib/assets/img/mashable.png';
	import TechCrunch from '$lib/assets/img/TechCrunch.png';
	import tnw from '$lib/assets/img/tnw.png';
</script>

<div class="testimonialBS">
	<section id="testimonials">
		<div id="testimonial-carousel" class="testimonialBS" data-ride="false">
			<h2 class="mx-5 pt-3 fs-2 z-index:2">
				As a developer who has always believed in the power of collaboration, finding 
				Hack Together was a game-changer for me. It's not just a platform; 
				it's a community where creativity meets technology. 
			</h2>
			<img class="testimonial-image" src={bizinsider} alt="dog-profile" />
			<em>Peebles, New York</em>
		</div>
	</section>
</div>

<section id="press">
	<img class="press-logo" src={TechCrunch} alt="tc-logo" />
	<img class="press-logo" src={tnw} alt="tnw-logo" />
	<img class="press-logo" src={bizinsider} alt="biz-insider-logo" />
	<img class="press-logo" src={mashable} alt="mashable-logo" />
</section>

<style>
	.testimonialBS {
		z-index: 4;
		position: relative;
	}

	#testimonials {
		text-align: center;
		background-color: #0f0;
		color: #000000;
		z-index: 4;
		position: relative;
		
	}

	h2 {
		font-size: 3rem;
		font-family: 'Montserrat', sans-serif;
		font-weight: 900;
		line-height: 1.5;
	}

	.testimonial-image {
		width: 10%;
		border-radius: 100%;
		margin: 20px;
	}

	#press {
		background-color: #0f0;
		text-align: center;
		padding-bottom: 3%;
		position: relative;
		z-index: 4;
	}

	.press-logo {
		width: 15%;
		margin: 20px 20px 50px;
	}

	.carousel-item {
		padding: 7% 15%;
	}
</style>
