<script lang="ts">
    import { Config } from "src/config";
    import { onMount } from "svelte";
    import { page } from "$app/stores";
    import type { User } from "src/types";
    import Users from "./users.svelte";
    import Clues from "./clues.svelte";
    import Guesses from "./guesses.svelte";

    const id = $page.params.id;
    let users: User[] = [];
    let user: User | null = null;
    let conn: WebSocket;
    let isLoading = true;

    type Response = {
        users: User[];
    };

    type ClientResponse = {
        user: User;
    };

    onMount(async () => {
        conn = new WebSocket(Config.VITE_WS_URL + "/ws/" + id);
        conn.onclose = function (evt) {
            console.log(evt);
        };
        conn.onmessage = function (evt) {
            const json = JSON.parse(evt.data) as Response | ClientResponse;
            if ("users" in json) {
                users = json.users || [];
                const userId = localStorage.getItem("userId");
                user = users.find((u) => u.id === userId) || null;
            } else if ("user" in json) {
                localStorage.setItem("userId", json.user.id);
            }
            isLoading = false;
        };
        conn.onopen = function () {
            const user = localStorage.getItem("user") || "{}";
            conn.send(JSON.stringify({ type: "join", data: user, room: id }));
        };
    });
</script>

{#if isLoading}
    <div>Loading...</div>
{:else if user?.step === 1 || !user}
    <Users {conn} {users} />
{:else if user.step === 2}
    <Clues {conn} />
{:else if user.step === 3}
    <Guesses {conn} />
{/if}
