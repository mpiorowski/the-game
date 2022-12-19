<script lang="ts">
    import { Config } from "src/config";
    import { onMount } from "svelte";
    import { page } from "$app/stores";

    const id = $page.params.id;

    let conn: WebSocket;
    onMount(async () => {
        conn = new WebSocket(Config.VITE_WS_URL + "/ws/" + id);
        conn.onclose = function (evt) {
            console.log(evt);
        };
        conn.onmessage = function (evt) {
            console.log(evt);
        };
    });

    const onSubmit = () => {
        if (conn) {
            conn.send(JSON.stringify({ type: "message", data: "Hello" }));
            conn.send(JSON.stringify({ type: "start-timer", data: "Hello" }));
        }
    };
</script>

<form id="form" on:submit|preventDefault={onSubmit}>
    <input type="submit" value="Send" />
    <input type="text" id="msg" size="64" />
</form>
