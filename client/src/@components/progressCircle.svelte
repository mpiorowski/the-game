<script lang="ts">
    export let progress = 0;
    export let progressNumber: number | null = null;
    export let size = 160;
    export let strokeWidth = 10;
    export let maxProgress = 60;

    const circumference = 2 * Math.PI * (size / 2 - strokeWidth / 2);
</script>

<div id="progress-container" style="width: {size}px; height: {size}px;">
    <svg viewBox="0 0 {size} {size}" width={size} height={size}>
        <circle
            class="progress-background"
            cx={size / 2}
            cy={size / 2}
            r={size / 2 - strokeWidth / 2}
            fill="transparent"
            stroke-width={strokeWidth}
        />
        <circle
            class="progress"
            cx={size / 2}
            cy={size / 2}
            r={size / 2 - strokeWidth / 2}
            fill="transparent"
            stroke-linecap="round"
            stroke-width={strokeWidth}
            stroke-dasharray={circumference}
            stroke-dashoffset={circumference -
                (circumference * progress) / maxProgress}
            transform="rotate(-90, {size / 2}, {size / 2})"
        />
    </svg>
    <div class="progress-counter">
        <span class="progress-value">{progressNumber ?? progress}</span>
    </div>
</div>

<style>
    #progress-container {
        position: relative;
        display: flex;
        align-items: center;
        justify-content: center;
    }
    .progress-background {
        @apply stroke-slate-400;
    }
    .progress {
        @apply stroke-slate-600;
        transition: stroke-dashoffset 100ms ease-in-out;
    }
    .progress-counter {
        font-family: 'Roboto', sans-serif;
        position: absolute;
        top: 50%;
        left: 50%;
        transform: translate(-50%, -50%);
    }
    .progress-value {
        @apply text-5xl font-bold text-slate-600;
    }
</style>

