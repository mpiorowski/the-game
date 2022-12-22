<script lang="ts">
    import { page } from "$app/stores";
    import { userId } from "src/store";
    import type { User } from "src/types";

    export let conn: WebSocket;
    export let users: User[] = [];

    const id = $page.params.id;
    let nickname = "";

    const onCreateUser = () => {
        const uuid = crypto.randomUUID();
        const data = JSON.stringify({
            id: uuid,
            nickname,
        });
        const request = JSON.stringify({
            type: "nickname",
            data,
            room: id,
        });
        userId.set(uuid);
        conn.send(request);
    };

    const onReady = () => {
        const request = JSON.stringify({
            type: "ready",
            data: $userId,
            room: id,
        });
        conn.send(request);
    };

    const goToClues = () => {
        const request = JSON.stringify({
            type: "go-to-clues",
            data: $userId,
            room: id,
        });
        conn.send(request);
    };
</script>

<form id="form" on:submit|preventDefault={onCreateUser}>
    <input type="submit" value="Send" />
    <input type="text" bind:value={nickname} size="64" />
</form>

<form id="form" on:submit|preventDefault={onReady}>
    <input type="submit" value="Ready" />
</form>

<form id="form" on:submit|preventDefault={goToClues}>
    <input type="submit" value="Clues" />
</form>

{#each users as user}
    <div
        style="display: flex; gap: 4px; flex-direction: row; align-items: center;"
    >
        <div>Nickname: {user.nickname}</div>
        <div>Team: {user.team}</div>
        <div>Ready: {user.ready}</div>
    </div>
{/each}
