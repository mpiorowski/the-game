<script>
    import { goto } from '$app/navigation';
    import { Button } from '@mpiorowski/svelte-init';
    import { fade } from 'svelte/transition';
</script>

<div class="h-full grid grid-rows-[1fr_auto]" in:fade>
    <div class="flex flex-col gap-4 items-center px-6 overflow-auto mb-4">
        <h2>Rules</h2>
        <h3>What we need?</h3>
        <ul class="list-disc text-lg">
            <li>
                At least 4 players, min 6 players recommended, the best
                experience with 8-12 players (even number recommended)
            </li>
            <li>
                Some time (2-4 hours, depends on number of players and game
                knowledge)
            </li>
            <li>Some alcohol (but not too much, game is not easy)</li>
        </ul>
        <h3>How to play?</h3>
        <div class="text-lg font-bold">
            Everything will be explained during the game, so You can jump into
            it.....or if You are still here, some basics for You:
        </div>
        <ul class="list-disc text-lg">
            <li>Players are divided into two teams</li>
            <li>Game consists of 3 rounds</li>
            <li>
                Each round has a different goal, but the main goal is to guess
                as many words as possible
            </li>
            <li>
                Each round will be explained just before it starts, to maximize
                the fun
            </li>
            <li>
                Each team will have 1 minute to guess as many words as possible
            </li>
            <li>The round ends when all words are guessed</li>
            <li>After all 3 rounds, the team with the most guessed words wins</li>
        </ul>
    </div>
    <Button type="primary" on:click={() => goto('/')}>Return to rooms</Button>
</div>
