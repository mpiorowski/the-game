<script>
    import { Toast } from '@mpiorowski/svelte-init';

    import '../app.css';
</script>

<Toast />
<div class="max-w-xl p-4 h-screen w-screen">
    <slot />
</div>
