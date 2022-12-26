<script lang="ts">
    import type { Clue, Score, User } from 'src/types';
    import Scoring from './scoring.svelte';

    export let users: User[] = [];
    export let score: Score[] = [];
    export let clues: Clue[] = [];

    const sumOfScoresA = score.reduce(
        (acc, curr) => acc + curr.teamClues[0].length,
        0
    );
    const sumOfScoresB = score.reduce(
        (acc, curr) => acc + curr.teamClues[1].length,
        0
    );
</script>

<div class="text-center">
    <h1 class="mb-6 font-bold">
        {#if sumOfScoresA > sumOfScoresB}
            Team A wins!
        {:else if sumOfScoresA < sumOfScoresB}
            Team B wins!
        {:else}
            It's a tie!
        {/if}
    </h1>
    <h3 class="mb-4">
        Want to play again? Just go to rooms and crate a new one. Hope You have
        fun! :)
    </h3>
    <Scoring {score} />
    <h3 class="mt-6">Guess words</h3>
    {#each users as user}
        <div>
            <h4 class="font-bold">{user.nickname}</h4>
            <ul class="list-decimal list-inside">
                {#each clues.filter((el) => el.userId === user.id) as clue}
                    <li class="text-center">{clue.word}</li>
                {/each}
            </ul>
        </div>
    {/each}
</div>
