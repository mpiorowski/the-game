<script lang="ts">
    import { page } from '$app/stores';
    import { Input } from '@mpiorowski/svelte-init';
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
    <form id="nickname" class="w-full" on:submit|preventDefault={onCreateUser}>
        <Input type="text" label="Nickname" bind:value={nickname} />
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

{#each users as user}
    <div
        style="display: flex; gap: 4px; flex-direction: row; align-items: center;"
    >
        <div>Nickname: {user.nickname}</div>
        <div>Team: {user.team}</div>
        <div>Ready: {user.ready}</div>
    </div>
{/each}
