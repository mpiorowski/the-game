<script lang="ts">
    import { page } from "$app/stores";
    import type { Round, User } from "src/types";

    const id = $page.params.id;
    export let conn: WebSocket;
    export let user: User;
    export let round: Round | null;

    const onStartRound = () => {
        const request = JSON.stringify({
            type: "start-round",
            data: user.id,
            room: id,
        });
        conn.send(request);
    };

    const onSendGuess = () => {
        if (conn && round) {
            conn.send(
                JSON.stringify({
                    type: "send-guess",
                    data: "correct",
                    room: id,
                    userId: user.id,
                })
            );
        }
    };
</script>

{#if user.step === 3}
    <h1>Waiting for other players to submit clues</h1>
{/if}
{#if user.step === 4 && round}
    <h1>Guesses</h1>
    <h2>
        Game: {round.game}
        Team: {round.team}
    </h2>
    {#if round.user.id === user.id && round.time === -1}
        <h2>Start Timer</h2>
        <button on:click={onStartRound}>Start round</button>
    {/if}
    {#if round.user.id === user.id && round.time > 0}
        <h2>It's your turn</h2>
        <h3>
            {round.clue.word}
        </h3>
        <h1>
            {round.time}
        </h1>
        <button on:click={onSendGuess}>Submit</button>
    {/if}
    {#if round.nextUser.id === user.id}
        <h1>
            {round.time}
        </h1>
        <h2>You're next</h2>
    {:else if round.user.id !== user.id}
        <h1>
            {round.time}
        </h1>
    {/if}
{/if}
