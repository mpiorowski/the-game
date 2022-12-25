<script lang="ts">
    import { page } from '$app/stores';
    import { Button } from '@mpiorowski/svelte-init';
    import { Teams, type Round, type User } from 'src/types';

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

{#if round.time === -1}
    <div class="text-center flex flex-col gap-6">
        {#if round.game === 0}
            <h2>Round 1 - Taboo</h2>
            <h3>
                The first round is a Taboo game, You need to explain given
                sentence without using any of the words it contains.
            </h3>
        {:else if round.game === 1}
            <h2>Round 2 - Charades</h2>
            <h3>
                The second round is a Charades game, You need to explain given
                sentence without using any words.
            </h3>
        {:else if round.game === 2}
            <h2>Round 3 - Stone</h2>
            <h3>
                The third and final round is a little different. You look at
                Your sentence, <b>REMEMBER IT</b>
                , and then
                <b>SAY ONE WORD</b>
                (not included in the sentence). After that You turn
                <b>INTO STONE</b>
                , no move, no mimick, nothing, not until Your team says the FULL,
                CORRECT sentence. If you at any point as much as blink, timer stops,
                and next time starts.
            </h3>
        {/if}
        {#if round.user.id === user.id}
            <Button on:click={onStartRound}>Start round!</Button>
        {:else}
            <h3>
                Starting person is <b>{round.user.nickname}</b>
                from team
                <b>{Teams[round.team]}</b>
            </h3>
        {/if}
    </div>
{:else}
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
