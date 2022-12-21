<script lang="ts">
    import { page } from "$app/stores";
    import type { User } from "src/types";

    export let conn: WebSocket;
    export let users: User[] = [];

    const id = $page.params.id;
    let nickname = "";

    const onSubmit = () => {
        if (conn) {
            conn.send(
                JSON.stringify({ type: "nickname", data: nickname, room: id })
            );
        }
    };

    const onReady = () => {
        const userId = localStorage.getItem("userId");
        if (conn && userId) {
            conn.send(JSON.stringify({ type: "ready", data: userId }));
        }
    };

    const goToClues = () => {
        const userId = localStorage.getItem("userId");
        if (conn && userId) {
            conn.send(JSON.stringify({ type: "clues", data: userId }));
        }
    };
</script>

<form id="form" on:submit|preventDefault={onSubmit}>
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
