<script lang="ts">
    import type { Score } from 'src/types';
    export let score: Score[] = [];

    const sumOfScoresA = score.reduce((acc, curr) => acc + curr.teamClues[0].length, 0);
    const sumOfScoresB = score.reduce((acc, curr) => acc + curr.teamClues[1].length, 0);
</script>

<div class="text-center mb-2">
    <h3 class="mb-2">Score</h3>
    <div class="grid grid-cols-3">
        <div>Round</div>
        <div>Team A</div>
        <div>Team B</div>
        {#each [0, 1, 2] as s}
            <div>{s + 1}</div>
            <div>{score[s] ? score[s].teamClues[0].length : 0}</div>
            <div>{score[s] ? score[s].teamClues[1].length : 0}</div>
        {/each}
        <div class="border-t border-slate-300">Sum</div>
        <div class="border-t border-slate-300">{sumOfScoresA}</div>
        <div class="border-t border-slate-300">{sumOfScoresB}</div>
    </div>
</div>
