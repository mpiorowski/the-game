<script lang="ts">
    import { goto } from '$app/navigation';
    import { Button, Input, toast, ToastType } from '@mpiorowski/svelte-init';
    import { useApi } from 'src/@utils/api.util';
    import type { Room } from 'src/types';
    import { fade } from 'svelte/transition';

    let room = {
        name: '',
        password: '',
    };

    const { request, response, error } = useApi<Room>();
    async function joinRoom() {
        await request({
            url: '/rooms',
            method: 'POST',
            body: JSON.stringify(room),
        });
        response.subscribe((data) => {
            if (data?.id) {
                toast('Joining room...', ToastType.INFO, 2000);
                goto(`/${data.id}`);
            }
        });
    }
    $: if ($error) {
        toast(String($error), ToastType.ERROR);
    }
</script>

<form
    in:fade
    class="w-full h-full flex flex-col gap-2 justify-center items-center"
    id="room"
    on:submit|preventDefault={joinRoom}
>
    <h1 class="mb-6">The Game</h1>
    <Input bind:value={room.name} label="Room name" />
    <Input bind:value={room.password} label="Room password" />
    <Button type="primary" form="room">Create or join room</Button>
    <Button
        type="ghost"
        on:click={() => {
            goto('/rules');
        }}
    >
        Rules
    </Button>
</form>
