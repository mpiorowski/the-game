<script lang="ts">
    import { goto } from "$app/navigation";
    import { useApi } from "src/@utils/api.util";
    import type { Room } from "./room.type";

    let room = {
        name: "Room 1",
        password: "1234",
    };

    const { request, response } = useApi<Room>();
    async function joinRoom() {
        await request({
            url: "/rooms",
            method: "POST",
            body: JSON.stringify(room),
        });
        response.subscribe((data) => {
            if (data?.id) {
                goto(`/${data.id}`);
            }
        });
    }
</script>

<h1>Create room</h1>
<form on:submit|preventDefault={joinRoom}>
    <input type="text" bind:value={room.name} />
    <input type="password" bind:value={room.password} />
    <button type="submit">Create</button>
</form>
