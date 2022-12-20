<script lang="ts">
    import { Config } from "src/config";
    import { onMount } from "svelte";
    import { page } from "$app/stores";

    const id = $page.params.id;

    let nickname = "";

    let users: {
        id: string;
        nickname: string;
        team: string;
        ready: boolean;
    }[] = [];

    let userId = "";

    let conn: WebSocket;
    onMount(async () => {
        conn = new WebSocket(Config.VITE_WS_URL + "/ws/" + id);
        conn.onclose = function (evt) {
            console.log(evt);
        };
        conn.onmessage = function (evt) {
            const json = JSON.parse(evt.data);
            if (json.type === "users") {
                users = json.data;
            }
            if (json.type === "user") {
                userId = json.data.id;
            }
        };
    });

    const onSubmit = () => {
        if (conn) {
            conn.send(
                JSON.stringify({ type: "nickname", data: nickname, room: id })
            );
        }
    };

    const onReady = () => {
        if (conn) {
            conn.send(JSON.stringify({ type: "ready", data: userId }));
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

{#each users as user}
    <div style="display: flex; gap: 4px; flex-direction: row; align-items: center;">
        <div>Nickname: {user.nickname}</div>
        <div>Team: {user.team}</div>
        <div>Ready: {user.ready}</div>
    </div>
{/each}
