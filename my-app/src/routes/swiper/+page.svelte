<!-- +page.svelte -->

<script>
    import Card from './Card.svelte';
  
    // Assuming you have imageUrl, onDismiss, onLike, onDislike props
    let imageUrl;
    let onDismiss;
    let onLike;
    let onDislike;
  </script>
  
  <Card {imageUrl} {onDismiss} {onLike} {onDislike} />
  