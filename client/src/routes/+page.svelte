<script lang="ts">
    import { goto } from '$app/navigation';
    import { Button, Input, toast, ToastType } from '@mpiorowski/svelte-init';
    import { useApi } from 'src/@utils/api.util';
    import type { Room } from 'src/types';

    let room = {
        name: 'Room 1',
        password: '1234',
    };

    const { request, response } = useApi<Room>();
    async function joinRoom() {
        toast('Joining room...', ToastType.INFO, 2000);
        await request({
            url: '/rooms',
            method: 'POST',
            body: JSON.stringify(room),
        });
        response.subscribe((data) => {
            if (data?.id) {
                goto(`/${data.id}`);
            }
        });
    }
</script>

<form class="w-full h-full flex flex-col justify-center items-center" id="room" on:submit|preventDefault={joinRoom}>
    <h1>The Game</h1>
    <Input bind:value={room.name} label="Room name" />
    <Input bind:value={room.password} label="Room password" />
    <Button type="ghost" form="room">Create or join room</Button>
</form>
