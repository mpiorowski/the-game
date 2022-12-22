<script lang="ts">
    import { Config } from "src/config";
    import { onMount } from "svelte";
    import { page } from "$app/stores";
    import type { Round, User } from "src/types";
    import Users from "./users.svelte";
    import Clues from "./clues.svelte";
    import Guesses from "./guesses.svelte";
    import { userId } from "src/store";

    const id = $page.params.id;
    let users: User[] = [];
    let round: Round | null = null;
    let user: User | null = null;
    let conn: WebSocket;
    let isLoading = true;

    type Response = {
        users: User[];
        round: Round;
    };

    onMount(async () => {
        // userId = localStorage.getItem("userId");
        conn = new WebSocket(Config.VITE_WS_URL + "/ws/" + id);
        conn.onclose = function (evt) {
            console.log(evt);
        };
        conn.onmessage = function (evt) {
            const json = JSON.parse(evt.data) as Response;
            round = json.round;
            users = json.users || [];
            user = users.find((u) => u.id === $userId) || null;
            isLoading = false;
        };
        conn.onopen = function () {
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
{:else if user.step === 3 || user.step === 4}
    <Guesses {conn} {user} {round} />
{:else}
    <div>Something went wrong. Please refresh the page.</div>
{/if}
