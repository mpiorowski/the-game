<script lang="ts">
    import { page } from '$app/stores';
    import { Button } from '@mpiorowski/svelte-init';
    import ProgressCircle from 'src/@components/progressCircle.svelte';
    import type { Round, User } from 'src/types';

    const id = $page.params.id;
    export let conn: WebSocket;
    export let user: User;
    export let round: Round;

    const onStartRound = () => {
        const request = JSON.stringify({
            type: 'start-round',
            data: user.id,
            room: id,
        });
        conn.send(request);
    };

    const onSendGuess = () => {
        if (conn && round) {
            conn.send(
                JSON.stringify({
                    type: 'send-guess',
                    data: 'correct',
                    room: id,
                    userId: user.id,
                })
            );
        }
    };
</script>

<div class="text-center flex flex-col items-center gap-6">
    <ProgressCircle
        size={200}
        progress={round.time === -1 ? 0 : 60 - round.time}
        progressNumber={round.time === -1 ? 60 : round.time}
    />
    {#if round.user.id === user.id && round.time > 0}
        <h2 class="font-bold">It's your turn!</h2>

        <h3>
            Category: <b>{round.clue.type}</b>
        </h3>
        <h3>Sentence:</h3>
        <h3>
            {#if round.clue.type === 'movie'}
                <b>"{round.clue.word}"</b>
            {:else}
                <b>{round.clue.word}</b>
            {/if}
        </h3>
        <Button on:click={onSendGuess}>Correct!</Button>
    {:else if round.user.id === user.id}
        <h2 class="font-bold">It's Your turn!</h2>
        <Button on:click={onStartRound}>Start Your team turn!</Button>
    {/if}
    {#if round.nextUser.id === user.id}
        <h2 class="font-bold">You're next!</h2>
    {/if}
</div>
