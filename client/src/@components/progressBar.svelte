<script lang="ts">
    import { onMount } from 'svelte';

    export let progress: number;
    export let size = 150;
    export let strokeWidth = 10;
    export let circleOneStroke = '#7ea9e1';
    export let circleTwoStroke = '#7ea9e1';

    let offset: number;

    const center = size / 2;
    const radius = size / 2 - strokeWidth / 2;
    const circumference = 2 * Math.PI * radius;

    onMount(() => {
        offset = ((100 - progress) / 100) * circumference;
    });
</script>

<svg width={size} height={size}>
    <circle
        class="circleBg"
        stroke={circleOneStroke}
        cx={center}
        cy={center}
        r={radius}
        stroke-width={strokeWidth}
    />
    <circle
        class="progressCircle"
        stroke={circleTwoStroke}
        cx={center}
        cy={center}
        r={radius}
        stroke-width={strokeWidth}
        stroke-dasharray={circumference}
        stroke-dashoffset={offset}
    />
</svg>

<style>
    .circleBg {
        @apply stroke-slate-400;
        fill: none;
    }

    .progressCircle {
        @apply stroke-slate-600;
        fill: none;
        transition: stroke-dashoffset 850ms ease-in-out;
        stroke-linecap: round;

        transform: rotate(-90deg);
        transform-origin: 50% 50%;
    }
</style>
