<script lang="js">
    import bizinsider from '$lib/assets/img/bizinsider.png';
    import mashable from '$lib/assets/img/mashable.png';
	import TechCrunch from '$lib/assets/img/TechCrunch.png';
	import tnw from '$lib/assets/img/tnw.png';
</script>

<section id="testimonials">
	<div id="testimonial-carousel" class="carousel slide" data-ride="false">
		<h2 class="mx-5 pt-3">
			I no longer have to sniff other dogs for love. I've found the hottest Corgi on TinDog. Woof.
		</h2>
		<img class="testimonial-image" src={bizinsider} alt="dog-profile" />
		<em>Peebles, New York</em>
	</div>
</section>

<section id="press">
	<img class="press-logo" src={TechCrunch} alt="tc-logo" />
	<img class="press-logo" src={tnw} alt="tnw-logo" />
	<img class="press-logo" src={bizinsider} alt="biz-insider-logo" />
	<img class="press-logo" src={mashable} alt="mashable-logo" />
</section>

<style>
    	#testimonials {
		text-align: center;
		background-color: #ef8172;
		color: #fff;
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
		background-color: #ef8172;
		text-align: center;
		padding-bottom: 3%;
	}

	.press-logo {
		width: 15%;
		margin: 20px 20px 50px;
	}

	.carousel-item {
		padding: 7% 15%;
	}
</style>