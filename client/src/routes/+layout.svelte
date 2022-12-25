<script>
    import { Toast } from '@mpiorowski/svelte-init';

    import '../app.css';
</script>

<Toast />
<div class="max-w-xl m-auto px-4 py-8 h-screen w-screen">
    <slot />
</div>
