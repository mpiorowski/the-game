<script>
    import { Toast } from '@mpiorowski/svelte-init';

    import '../app.css';
</script>

<Toast />
<div class="max-w-xl m-auto px-4 pt-8 pb-4 h-full fixed w-screen">
    <slot />
</div>
