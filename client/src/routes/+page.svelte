<script lang="ts">
    import { useApi } from "src/@utils/api.util";

    let room = {
        name: "Room 1",
        password: "",
    };

    const { request } = useApi();
    function joinRoom() {
        request({
            url: "/rooms",
            method: "POST",
            body: JSON.stringify(room),
        });
    }

</script>

<h1>Create room</h1>
<form on:submit|preventDefault={joinRoom}>
    <input type="text" bind:value={room.name} />
    <input type="password" bind:value={room.password} />
    <button type="submit">Create</button>
</form>
