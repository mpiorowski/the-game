<script lang="ts">
    import { Config } from 'src/config';
    import { onMount } from 'svelte';
    import { page } from '$app/stores';
    import {
        Team,
        type Clue,
        type Round,
        type Score,
        type User,
    } from 'src/types';
    import Users from './users.svelte';
    import Clues from './clues.svelte';
    import Guesses from './guesses.svelte';
    import { Spinner } from '@mpiorowski/svelte-init';
    import Rounds from './rounds.svelte';
    import End from './end.svelte';
    import { fade } from 'svelte/transition';

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
</script>

<div class="h-full grid grid-rows-[1fr_auto]" in:fade>
    <div class="overflow-auto p-1">
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

    {#if user?.id}
        <div class="flex gap-1 justify-center mb-1">
            You are <b>{user.nickname}</b>
            {#if user.team >= 0}
                in team
                <b>{Team[user.team]}</b>
            {/if}
        </div>
    {/if}
</div>
