<script lang="ts">
    import { page } from '$app/stores';
    import { Button, Input, toast, ToastType } from '@mpiorowski/svelte-init';
    import { Team, type User } from 'src/types';
    import { fade } from 'svelte/transition';
    import { z } from 'zod';

    export let conn: WebSocket;
    export let users: User[] = [];
    export let user: User | null;

    const id = $page.params.id;
    let nickname = '';

    const onCreateUser = () => {
        const validation = z.object({
            nickname: z.string().min(1),
        }).safeParse({ nickname }).success;

        if (!validation) {
            toast('Invalid nickname', ToastType.ERROR);
            return;
        }

        const uuid = crypto.randomUUID();
        const data = JSON.stringify({
            id: uuid,
            nickname,
        });
        const request = JSON.stringify({
            type: 'nickname',
            data,
            room: id,
        });
        localStorage.setItem('userId', uuid);
        conn.send(request);
    };

    const onReady = () => {
        const request = JSON.stringify({
            type: 'ready',
            data: user?.id,
            room: id,
        });
        conn.send(request);
    };

    const onGoToClues = () => {
        const request = JSON.stringify({
            type: 'go-to-clues',
            data: user?.id,
            room: id,
        });
        conn.send(request);
    };
</script>

<div
    class="flex flex-col h-40 justify-center items-center border-b border-gray-300 mb-8"
    in:fade
>
    {#if !user?.id}
        <form
            id="nickname"
            class="w-full"
            on:submit|preventDefault={onCreateUser}
        >
            <Input type="text" label="Nickname" bind:value={nickname} />
            <Button type="primary" form="nickname">Create user</Button>
        </form>
    {:else if users.length < 4}
        <h2 class="text-center">Waiting for other players to join...</h2>
    {:else if !user?.ready}
        <form id="ready" class="w-full" on:submit|preventDefault={onReady}>
            <h2 class="text-center mb-4">
                {user?.nickname}, are you ready to play the best social game
                ever?
            </h2>
            <Button type="primary" form="ready">I am ready!</Button>
        </form>
    {:else if users.some((u) => !u.ready)}
        <h2 class="text-center">Waiting for other players to be ready...</h2>
    {:else if user.ready}
        <form id="clues" class="w-full" on:submit|preventDefault={onGoToClues}>
            <h2 class="text-center mb-4">
                {user?.nickname}, you are on team
                <span class="font-bold">
                    {Team[user?.team]}
                </span>
                .
                <br />
                Let's the game begin!
            </h2>

            <Button type="primary" form="clues">Start the fun</Button>
        </form>
    {/if}
</div>

<div
    class="grid grid-cols-[1fr_auto_1fr] w-full justify-center justify-items-center text-lg"
>
    <h3>Nickname</h3>
    <h3>Team</h3>
    <h3>Ready</h3>
    {#each users.sort((a, b) => a.team - b.team) as user}
        <div class="font-bold">{user.nickname}</div>
        <div class="font-bold">{user.team > -1 ? Team[user.team] : '-'}</div>
        <div>
            {#if user.ready}
                <span class="text-green-800">&#10004;</span>
            {:else}
                <span class="text-red-800">&#x2717;</span>
            {/if}
        </div>
    {/each}
</div>
