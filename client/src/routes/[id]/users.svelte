<script lang="ts">
    import { page } from '$app/stores';
    import { Button, Input } from '@mpiorowski/svelte-init';
    import type { User } from 'src/types';

    export let conn: WebSocket;
    export let users: User[] = [];
    export let user: User | null;

    const id = $page.params.id;
    let nickname = '';

    const onCreateUser = () => {
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

    const goToClues = () => {
        const request = JSON.stringify({
            type: 'go-to-clues',
            data: user?.id,
            room: id,
        });
        conn.send(request);
    };
</script>

{#if !user?.id}
    <form
        id="nickname"
        class="flex flex-col h-40"
        on:submit|preventDefault={onCreateUser}
    >
        <Input type="text" label="Nickname" bind:value={nickname} />
        <Button type="ghost" form="nickname">Create user</Button>
    </form>
{:else if !user?.ready}
    <form id="form" on:submit|preventDefault={onReady}>
        <input type="submit" value="Ready" />
    </form>
{:else if user.ready}
    <form id="form" on:submit|preventDefault={goToClues}>
        <input type="submit" value="Clues" />
    </form>
{/if}

<div class="grid grid-cols-[1fr_auto_1fr] w-full justify-center justify-items-center text-lg">
    <h3>Nickname</h3>
    <h3>Team</h3>
    <h3>Ready</h3>
    {#each users as user}
        <div>{user.nickname}</div>
        <div>{user.team > -1 ? user.team : '-'}</div>
        <div>
            {#if user.ready}
                <span class="text-green-800">✅</span>
            {:else}
                <span class="text-red-800">❌</span>
            {/if}
        </div>
    {/each}
</div>
