<script>
    import { Toast } from '@mpiorowski/svelte-init';

    import '../app.css';
</script>

<Toast />
<div class="max-w-xl px-4 pt-8 pb-4 h-full fixed w-screen left-1/2 -translate-x-1/2 ">
    <slot />
</div>
