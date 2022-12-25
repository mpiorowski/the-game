<script lang="ts">
    import { Button } from '@mpiorowski/svelte-init';
    import { page } from '$app/stores';
    import { Team, type Round, type Score, type User } from 'src/types';
    import Scoring from './scoring.svelte';

    const id = $page.params.id;
    export let conn: WebSocket;
    export let round: Round;
    export let user: User;
    export let score: Score[];

    const onStartRound = () => {
        const request = JSON.stringify({
            type: 'start-round',
            data: user.id,
            room: id,
        });
        conn.send(request);
    };
</script>

<div class="grid grid-rows-[1fr_auto] h-full">
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
            <div class="text-lg">
                The third and final round is a little different. You've managed
                to get here, so You should know at least some of the crazy
                sentences that were created. So now, when Your turn comes, You
                <b>LOOK</b>
                at Your sentence with silence,
                <b>REMEMBER IT</b>
                , and then
                <b>SAY ONE WORD</b>
                (not included in the sentence). After that You turn
                <b>INTO STONE</b>
                , no move, no mimick, nothing, not until Your team says the
                <b>FULL, CORRECT</b>
                sentence. If you at any point as much as yawn, the oposing team can
                stop the timer.
            </div>
        {/if}
        {#if round.user.id === user.id}
            <h2 class="font-bold">It's Your turn!</h2>
            <Button on:click={onStartRound}>Start Your team turn</Button>
        {:else}
            <h3>
                Starting person is <b>{round.user.nickname}</b>
                from team
                <b>{Team[round.team]}</b>
            </h3>
        {/if}
        {#if round.nextUser.id === user.id}
            <h2 class="font-bold">You are next!</h2>
        {/if}
    </div>
    <Scoring {score} />
</div>
