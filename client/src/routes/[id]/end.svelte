<script lang="ts">
    import { goto } from '$app/navigation';
    import { Button } from '@mpiorowski/svelte-init';
    import type { Clue, Score, User } from 'src/types';
    import { fade } from 'svelte/transition';
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

    const onReset = () => {
        goto('/');
    };
</script>

<div class="text-center" in:fade>
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
    <Button on:click={onReset}>Go back to rooms</Button>
</div>
