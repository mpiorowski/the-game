<script lang="ts">
    import { Config } from 'src/config';
    import { onMount } from 'svelte';
    import { page } from '$app/stores';
    import type { Clue, Round, Score, User } from 'src/types';
    import Users from './users.svelte';
    import Clues from './clues.svelte';
    import Guesses from './guesses.svelte';
    import { Button, Spinner } from '@mpiorowski/svelte-init';
    import { goto } from '$app/navigation';
    import Rounds from './rounds.svelte';
    import End from './end.svelte';

    const id = $page.params.id;

    let users: User[] = [];
    let round: Round | null = null;
    let score: Score[] = [];
    let clues: Clue[] = [];

    let user: User | null = null;
    let conn: WebSocket;
    let isLoading = true;

    type Response = {
        users: User[];
        round: Round;
        score: Score[];
        clues: Clue[];
    };

    onMount(async () => {
        conn = new WebSocket(Config.VITE_WS_URL + '/ws/' + id);
        conn.onclose = function (evt) {
            console.log(evt);
        };
        conn.onmessage = function (evt) {
            const userId = localStorage.getItem('userId');
            const json = JSON.parse(evt.data) as Response;
            round = json.round;
            users = json.users || [];
            score = json.score || [];
            clues = json.clues || [];
            user = users.find((u) => u.id === userId) || null;
            isLoading = false;
        };
        conn.onopen = function () {
            conn.send(JSON.stringify({ type: 'join', data: user, room: id }));
        };
    });

    const onReset = () => {
        goto('/');
    };
</script>

<div class="h-full grid grid-rows-[1fr_auto]">
    <div class="overflow-auto mb-4">
        {#if isLoading}
            <Spinner center />
        {:else if user?.step === 1 || !user}
            <Users {conn} {users} {user} />
        {:else if user.step === 2 || user.step === 3}
            <Clues {conn} {user} />
        {:else if user.step === 4 && round}
            <Rounds {conn} {user} {round} {score} />
        {:else if user.step === 5 && round}
            <Guesses {conn} {user} {round} />
        {:else if user.step === 6}
            <End {users} {score} {clues} />
        {:else}
            <div>Something went wrong. Please refresh the page.</div>
        {/if}
    </div>

    <Button on:click={onReset}>Go back to rooms</Button>
</div>
